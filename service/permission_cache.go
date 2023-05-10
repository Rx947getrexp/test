package service

import (
	"context"
	"encoding/json"
	"go-speed/constant"
	"go-speed/global"
	"go-speed/model"
	"strconv"
	"strings"
)

type NodeTree struct {
	Id         int        `json:"id"`
	Pid        int        `json:"pid"`
	Name       string     `json:"name"`
	ResType    int        `json:"res_type"`
	ResTypeStr string     `json:"res_type_str"`
	Url        string     `json:"url"`
	Sort       int        `json:"sort"`
	Icon       string     `json:"icon"`
	Show       bool       `json:"show"`
	Disabled   bool       `json:"disabled"`
	Children   []NodeTree `json:"children,omitempty"`
}

type Node struct {
	Resource *model.AdminRes
	Children []int
}

func UpdateCachePermission() {
	UpdateCacheFullTree()
	UpdateCacheMenuTree()
	UpdateCacheRoleTree()
	UpdateCacheRoleUrl()
}

func UpdateCacheMenuTree() {
	menuFullTree := getFullMenuTree()
	menuFullTreeBytes, _ := json.Marshal(menuFullTree)
	global.Redis.Set(context.Background(), constant.MenuFullTreeKey, menuFullTreeBytes, 0).Err()
}

func UpdateCacheRoleTree() {
	roleTreeMap := getRoleTree()
	roleTreeMapBytes, _ := json.Marshal(roleTreeMap)
	global.Redis.Set(context.Background(), constant.RoleTreeMapKey, roleTreeMapBytes, 0).Err()
}

func UpdateCacheRoleUrl() {
	roleUrlMap := getRoleUrl()
	roleUrlMapBytes, _ := json.Marshal(roleUrlMap)
	global.Redis.Set(context.Background(), constant.RoleUrlMapKey, roleUrlMapBytes, 0).Err()
}

func UpdateCacheFullTree() {
	fullTree := getFullTree()
	fullTreeBytes, _ := json.Marshal(fullTree)
	global.Redis.Set(context.Background(), constant.FullTreeKey, fullTreeBytes, 0).Err()
}

func GetCacheMenuTree() []NodeTree {
	menuFullTreeBytes, err := global.Redis.Get(context.Background(), constant.MenuFullTreeKey).Bytes()
	if err != nil {
		global.Logger.Err(err).Msg("redis读取MenuFullTree出错")
		return getFullMenuTree()
	}
	var menuFullTree []NodeTree
	err = json.Unmarshal(menuFullTreeBytes, &menuFullTree)
	if err != nil {
		global.Logger.Err(err).Msg("解析json出错")
		return getFullMenuTree()
	}
	return menuFullTree
}

func GetCacheRoleTree(roleId int) []NodeTree {
	roleTreeMapBytes, err := global.Redis.Get(context.Background(), constant.RoleTreeMapKey).Bytes()
	if err != nil {
		global.Logger.Err(err).Msg("redis读取RoleTreeMapKey出错")
		return getRoleTree()[roleId]
	}
	var roleTreeMap map[int][]NodeTree
	err = json.Unmarshal(roleTreeMapBytes, &roleTreeMap)
	if err != nil {
		global.Logger.Err(err).Msg("解析json出错")
		return getRoleTree()[roleId]
	}
	if _, ok := roleTreeMap[roleId]; !ok {
		return []NodeTree{}
	}
	return roleTreeMap[roleId]
}

func GetCacheFullTree() []NodeTree {
	fullTreeBytes, err := global.Redis.Get(context.Background(), constant.FullTreeKey).Bytes()
	if err != nil {
		global.Logger.Err(err).Msg("redis读取FullTreeKey出错")
		return getFullTree()
	}
	var fullTree []NodeTree
	err = json.Unmarshal(fullTreeBytes, &fullTree)
	if err != nil {
		global.Logger.Err(err).Msg("解析json出错")
		return getFullTree()
	}
	return fullTree
}

func GetCacheRoleUrlHasPermission(roleId int, url string) bool {
	var roleUrlMap = make(map[int]map[string]bool)
	roleUrlMapBytes, err := global.Redis.Get(context.Background(), constant.RoleUrlMapKey).Bytes()
	if err != nil {
		global.Logger.Err(err).Msg("redis读取RoleTreeMapKey出错")
		roleUrlMap = getRoleUrl()
	} else {
		err = json.Unmarshal(roleUrlMapBytes, &roleUrlMap)
		if err != nil {
			global.Logger.Err(err).Msg("解析json出错")
			roleUrlMap = getRoleUrl()
		}
	}
	if _, ok := roleUrlMap[roleId]; !ok {
		return false
	}
	if _, ok := roleUrlMap[roleId][url]; !ok {
		return false
	}
	return roleUrlMap[roleId][url]
}

func getFullMenuTree() []NodeTree {
	treeList, _ := getResourceList()
	return generateTreeByResourceList(treeList)
}

func getFullTree() []NodeTree {
	treeList, _ := getAllResourceList()
	return generateTreeByResourceList(treeList)
}

func getRoleTree() map[int][]NodeTree {
	var resultMap = make(map[int][]NodeTree)
	treeList, _ := getResourceList()
	fullTree := generateTreeByResourceList(treeList)
	roleList, _ := getRoleResList()
	for _, item := range roleList {
		var roleId int
		has, _ := global.Db.Cols("id").Table("admin_role").Where("id= ? and is_del = 0", item.RoleId).Get(&roleId)
		if !has {
			//如果角色已删则跳过
			continue
		}
		var roleBoolMap = make(map[int]bool)
		var tmpIds []int
		resIds := strings.Split(item.ResIds, ",")
		for _, resId := range resIds {
			id, _ := strconv.Atoi(resId)
			tmpIds = append(tmpIds, id)
			roleBoolMap[id] = true
		}
		roleBoolMap = generateChildMap(tmpIds, roleBoolMap)
		roleBoolMap = generateDirectMap(tmpIds, roleBoolMap)
		resultMap[item.RoleId] = generateRoleTree(roleBoolMap, fullTree)
	}
	return resultMap
}

func getRoleUrl() map[int]map[string]bool {
	var resultMap = make(map[int]map[string]bool)
	roleList, _ := getRoleResList()
	for _, item := range roleList {
		var roleId int
		has, _ := global.Db.Cols("id").Table("admin_role").Where("id= ? and is_del = 0", item.RoleId).Get(&roleId)
		if !has {
			//如果角色已删则跳过
			continue
		}
		var roleBoolMap = make(map[string]bool)
		var tmpIds []int
		resIds := strings.Split(item.ResIds, ",")
		for _, resId := range resIds {
			id, _ := strconv.Atoi(resId)
			tmpIds = append(tmpIds, id)
		}
		roleBoolMap = generateChildUrl(tmpIds, roleBoolMap)
		resultMap[item.RoleId] = roleBoolMap
	}
	return resultMap
}

func getResourceList() ([]*model.AdminRes, error) {
	var list []*model.AdminRes
	err := global.Db.Where("is_del = 0 and res_type = 1").OrderBy("id asc").Find(&list)
	if err != nil {
		return nil, err
	}
	return list, nil
}

func getAllResourceList() ([]*model.AdminRes, error) {
	var list []*model.AdminRes
	err := global.Db.Where("is_del = 0").OrderBy("id asc").Find(&list)
	if err != nil {
		return nil, err
	}
	return list, nil
}

func getRoleResList() ([]*model.AdminRoleRes, error) {
	var list []*model.AdminRoleRes
	err := global.Db.Find(&list)
	if err != nil {
		return nil, err
	}
	return list, nil
}

func generateChildUrl(resources []int, roleBoolMap map[string]bool) map[string]bool {
	var urlResources []string
	global.Db.Cols("url").Table(model.AdminRes{}).Where("is_del = 0 and (res_type = 2 or res_type = 1)").In("pid", resources).Find(&urlResources)
	if len(urlResources) > 0 {
		for _, item := range urlResources {
			roleBoolMap[item] = true
		}
	}
	var childResources []int
	global.Db.Cols("id").Table(model.AdminRes{}).Where("is_del = 0 and res_type = 1").In("pid", resources).Find(&childResources)
	if len(childResources) > 0 {
		generateChildUrl(childResources, roleBoolMap)
	}
	return roleBoolMap
}

func generateChildMap(resources []int, roleBoolMap map[int]bool) map[int]bool {
	var childResources []int
	global.Db.Cols("id").Table(model.AdminRes{}).Where("is_del = 0 and res_type = 1").In("pid", resources).Find(&childResources)
	if len(childResources) > 0 {
		for _, item := range childResources {
			roleBoolMap[item] = true
		}
		generateChildMap(childResources, roleBoolMap)
	}
	return roleBoolMap
}

func generateDirectMap(resources []int, roleBoolMap map[int]bool) map[int]bool {

	var directResources []int
	global.Db.Cols("pid").Table(model.AdminRes{}).Where("is_del = 0 and res_type = 1").In("id", resources).Find(&directResources)
	if len(directResources) > 0 {
		for _, item := range directResources {
			roleBoolMap[item] = true
		}
		generateDirectMap(directResources, roleBoolMap)
	}
	return roleBoolMap
}

func generateTreeByResourceList(treeList []*model.AdminRes) []NodeTree {
	if len(treeList) == 0 {
		return []NodeTree{}
	}
	var nodeMap = make(map[int]*model.AdminRes)
	var nodeChildMap = make(map[int][]int)
	for _, item := range treeList {
		nodeMap[item.Id] = item
		nodeChildMap[item.Pid] = append(nodeChildMap[item.Pid], item.Id)
	}
	//遍历Pid为0下面的树
	return getTreeByPid(0, nodeMap, nodeChildMap)
}

func getTreeByPid(pid int, nodeMap map[int]*model.AdminRes, nodeChildMap map[int][]int) []NodeTree {
	var result []NodeTree
	childs := nodeChildMap[pid]
	if len(childs) == 0 {
		return []NodeTree{}
	}
	for _, child := range childs {
		var tree NodeTree
		children := getTreeByPid(nodeMap[child].Id, nodeMap, nodeChildMap)
		tree.Id = nodeMap[child].Id
		tree.Pid = nodeMap[child].Pid
		tree.Name = nodeMap[child].Name
		tree.ResType = nodeMap[child].ResType
		tree.Url = nodeMap[child].Url
		tree.Icon = nodeMap[child].Icon
		tree.Sort = nodeMap[child].Sort
		if tree.ResType == 1 {
			tree.ResTypeStr = "菜单"
		} else {
			tree.ResTypeStr = "接口"
		}
		if nodeMap[child].Id == 34 || nodeMap[child].Id == 41 || nodeMap[child].Id == 44 {
			tree.Disabled = true
		}
		if len(children) > 0 {
			tree.Children = children
		}
		result = append(result, tree)
	}
	return result
}

func generateRoleTree(roleBoolMap map[int]bool, fullTree []NodeTree) []NodeTree {
	var result []NodeTree
	for _, item := range fullTree {
		if !roleBoolMap[item.Id] {
			continue
		}
		var tree NodeTree
		tree.Id = item.Id
		tree.Pid = item.Pid
		tree.Name = item.Name
		tree.ResType = item.ResType
		tree.Url = item.Url
		tree.Icon = item.Icon
		tree.Sort = item.Sort
		if tree.ResType == 1 {
			tree.ResTypeStr = "菜单"
		} else {
			tree.ResTypeStr = "接口"
		}
		if len(item.Children) > 0 {
			tree.Children = generateRoleTree(roleBoolMap, item.Children)
		}
		result = append(result, tree)
	}
	if len(result) == 0 {
		return []NodeTree{}
	}
	return result
}
