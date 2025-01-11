package service

import (
	"fmt"
	"go-speed/global"
	"go-speed/model"
	"time"
)

func DepositOrder(order *model.TOrder) {
	var goods model.TGoods
	has, err := global.Db.Where("id = ?", order.GoodsId).Get(&goods)
	if err != nil || !has {
		global.Logger.Err(err).Msg("goods信息不正确")
		return
	}
	var sendDay int
	nowTime := time.Now()
	startSec := nowTime.Unix()
	goodsDay := goods.Period
	if goods.IsDiscount == 1 {
		//随机事件
		sendDay = GenerateRangeNum(goods.Low, goods.High+1)
	}
	totalSec := (goodsDay + sendDay) * 24 * 3600
	endSec := startSec + int64(totalSec)

	//开启事务
	sess := global.Db.NewSession()
	defer sess.Close()
	sess.Begin()

	//1.添加success-record记录
	successRecord := &model.TSuccessRecord{
		UserId:     order.UserId,
		OrderId:    order.Id,
		StartTime:  startSec,
		EndTime:    endSec,
		SurplusSec: int64(totalSec),
		TotalSec:   int64(totalSec),
		GoodsDay:   goodsDay,
		SendDay:    sendDay,
		PayType:    1, //1-银行卡；2-支付宝；3-微信支付
		Status:     1, //1-using使用中；2-wait等待; 3-end已结束
		CreatedAt:  nowTime,
		UpdatedAt:  nowTime,
		Comment:    "",
	}
	rows, err := sess.Insert(successRecord)
	if err != nil || rows < 1 {
		global.Logger.Err(err).Msg("添加successRecord记录失败")
		sess.Rollback()
		return
	}

	//2.如有赠送记录，添加赠送记录
	if sendDay > 0 {
		gift := &model.TGift{
			UserId:    order.UserId,
			OpId:      fmt.Sprint(order.Id),
			OpUid:     order.UserId,
			Title:     fmt.Sprintf("充值v%v套餐赠送", goods.MType),
			GiftSec:   sendDay * 24 * 3600,
			GType:     4, //赠送类别（1-注册；2-推荐；3-日常活动；4-充值）
			CreatedAt: nowTime,
			UpdatedAt: nowTime,
			Comment:   "",
		}
		rows, err = sess.Insert(gift)
		if err != nil || rows < 1 {
			global.Logger.Err(err).Msg("添加gift记录失败")
			sess.Rollback()
			return
		}
	}

	//3.修改用户表会员属性及到期时间
	var user model.TUser
	has, err = global.Db.Where("id = ?", order.UserId).Get(&user)
	if err != nil || !has {
		global.Logger.Err(err).Msg("用户不存在")
		sess.Rollback()
		return
	}
	if user.ExpiredTime < nowTime.Unix() {
		user.ExpiredTime = nowTime.Unix() + int64(totalSec)
	} else {
		user.ExpiredTime += int64(totalSec)
	}
	user.UpdatedAt = nowTime
	user.Level = goods.MType
	user.Kicked = 0
	rows, err = sess.Cols("level", "updated_at", "expired_time").Where("id = ?", user.Id).Update(&user)
	if err != nil || !has {
		global.Logger.Err(err).Msg("更新用户信息出错")
		sess.Rollback()
		return
	}

	//4.若存在推荐人，则赠送20%奖励; 并修改推荐人的会员到期时间
	var userTeam model.TUserTeam
	has, err = global.Db.Where("user_id = ? and direct_id > 0", user.Id).Get(&userTeam)
	if err != nil {
		global.Logger.Err(err).Msg("获取用户team信息出错")
		sess.Rollback()
		return
	}
	if has {
		var directUser model.TUser
		has, err = global.Db.Where("id = ?", userTeam.DirectId).Get(&directUser)
		if err != nil || !has {
			global.Logger.Err(err).Msg("获取directUser信息出错")
			sess.Rollback()
			return
		}
		awardTime := goodsDay * 24 * 3600 * 20 / 100 //赠送20%
		if directUser.ExpiredTime < nowTime.Unix() {
			directUser.ExpiredTime = nowTime.Unix() + int64(awardTime)
		} else {
			directUser.ExpiredTime += int64(awardTime)
		}
		directUser.UpdatedAt = nowTime
		directUser.Kicked = 0
		rows, err = sess.Cols("updated_at", "expired_time").Where("id = ?", directUser.Id).Update(&directUser)
		if err != nil || !has {
			global.Logger.Err(err).Msg("更新用户信息出错")
			sess.Rollback()
			return
		}

		directGift := &model.TGift{
			UserId:    directUser.Id,
			OpId:      fmt.Sprint(order.Id),
			OpUid:     order.UserId,
			Title:     fmt.Sprintf("粉丝v%v套餐推荐奖励", goods.MType),
			GiftSec:   awardTime,
			GType:     2, //赠送类别（1-注册；2-推荐；3-日常活动；4-充值）
			CreatedAt: nowTime,
			UpdatedAt: nowTime,
			Comment:   "",
		}
		rows, err = sess.Insert(directGift)
		if err != nil || rows < 1 {
			global.Logger.Err(err).Msg("添加gift记录失败")
			sess.Rollback()
			return
		}

	}

	sess.Commit()
}

func Change() {

}
