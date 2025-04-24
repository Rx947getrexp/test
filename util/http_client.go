package util

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"go-speed/global"
	"io/ioutil"
	"net/http"
	"strconv"
	"time"
)

func HttpClientPost(url string, bodyParam interface{}, response interface{}) error {
	paramBytes, err := json.Marshal(bodyParam)
	if err != nil {
		global.Logger.Err(err).Msg("json解析出错")
		return err
	}
	body := bytes.NewReader(paramBytes)
	httpReq, err := http.NewRequest("POST", url, body)
	if err != nil {
		global.Logger.Err(err).Msg("创建http请求失败")
		return err
	}
	httpReq.Header.Add("Content-Type", "application/json")
	httpReq.Close = true
	respData, err := http.DefaultClient.Do(httpReq)
	if err != nil {
		global.Logger.Err(err).Msg("发送http请求失败")
		return err
	}
	respBytes, err := ioutil.ReadAll(respData.Body)
	if err != nil {
		global.Logger.Err(err).Msg("读取响应数据失败")
		return err
	}
	fmt.Println("res响应：", string(respBytes))
	defer respData.Body.Close()
	switch response.(type) {
	case string:
		response = string(respBytes)
		fmt.Println("res---kkk响应：", response)
	default:
		if err = json.Unmarshal(respBytes, response); err != nil {
			global.Logger.Err(err).Msg("解析数据")
			return err
		}
	}
	return nil
}

func HttpClientPostV2(url string, headerParam map[string]string, bodyParam interface{}, response interface{}) error {
	paramBytes, err := json.Marshal(bodyParam)
	if err != nil {
		global.Logger.Err(err).Msg("json解析出错")
		return err
	}
	body := bytes.NewReader(paramBytes)

	httpReq, err := http.NewRequest("POST", url, body)
	if err != nil {
		global.Logger.Err(err).Msg("创建http请求失败")
		return err
	}
	httpReq.Header.Add("Content-Type", "application/json")
	for k, v := range headerParam {
		httpReq.Header.Add(k, v)
	}
	httpReq.Close = true
	client := http.Client{
		Timeout: time.Second * 60,
	}
	respData, err := client.Do(httpReq)
	if err != nil {
		global.Logger.Err(err).Msg("发送http请求失败")
		return err
	}
	respBytes, err := ioutil.ReadAll(respData.Body)
	if err != nil {
		global.Logger.Err(err).Msg("读取响应数据失败")
		return err
	}
	//fmt.Println("res响应：", string(respBytes))
	defer respData.Body.Close()
	switch response.(type) {
	case string:
		response = string(respBytes)
		fmt.Println("res---kkk响应：", response)
	default:
		if err = json.Unmarshal(respBytes, response); err != nil {
			global.Logger.Err(err).Msg("解析数据")
			return err
		}
	}
	return nil
}

func HttpClientPostReturnStr(url string, bodyParam interface{}) (string, error) {
	paramBytes, err := json.Marshal(bodyParam)
	if err != nil {
		global.Logger.Err(err).Msg("json解析出错")
		return "", err
	}
	body := bytes.NewReader(paramBytes)
	httpReq, err := http.NewRequest("POST", url, body)
	if err != nil {
		global.Logger.Err(err).Msg("创建http请求失败")
		return "", err
	}
	httpReq.Header.Add("Content-Type", "application/json")
	httpReq.Close = true
	respData, err := http.DefaultClient.Do(httpReq)
	if err != nil {
		global.Logger.Err(err).Msg("发送http请求失败")
		return "", err
	}
	respBytes, err := ioutil.ReadAll(respData.Body)
	if err != nil {
		global.Logger.Err(err).Msg("读取响应数据失败")
		return "", err
	}
	defer respData.Body.Close()
	return string(respBytes), nil
}

func HttpClientGet(url string, params map[string]interface{}, response interface{}) error {
	request, err := http.NewRequest("GET", url, nil)
	if err != nil {
		global.Logger.Err(err).Msg("创建http请求失败")
		return err
	}
	if params != nil {
		query := request.URL.Query()

		for name, value := range params {
			switch value.(type) {
			case string:
				query.Add(name, value.(string))

			case int:
				query.Add(name, strconv.Itoa(value.(int)))

			case float64:
				query.Add(name, strconv.FormatFloat(value.(float64), 'f', -1, 64))

			default:
				return errors.New("params type only support string, int and float64")
			}
		}

		request.URL.RawQuery = query.Encode()
	}
	request.Close = true
	respData, err := http.DefaultClient.Do(request)
	if err != nil {
		global.Logger.Err(err).Msg("发送http请求失败")
		return err
	}
	respBytes, err := ioutil.ReadAll(respData.Body)
	if err != nil {
		global.Logger.Err(err).Msg("读取响应数据失败")
		return err
	}
	defer respData.Body.Close()
	if err = json.Unmarshal(respBytes, response); err != nil {
		global.Logger.Err(err).Msg("解析数据")
		return err
	}
	return nil
}
