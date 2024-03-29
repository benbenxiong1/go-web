package sms

import (
	"encoding/json"
	aliyunsmsclient "github.com/KenmyZhang/aliyun-communicate"
	"go-web/pkg/logger"
)

type Aliyun struct {
}

func (a *Aliyun) Send(phone string, msg Message, config map[string]string) bool {
	smsClient := aliyunsmsclient.New("http://dysmsapi.aliyuncs.com/")

	templateParam, err := json.Marshal(msg.Data)
	if err != nil {
		logger.ErrorString("短信[阿里云]", "解析绑定错误", err.Error())
		return false
	}

	logger.DebugJSON("短信[阿里云]", "配置信息", config)

	result, err := smsClient.Execute(
		config["access_key_id"],
		config["access_key_secret"],
		phone,
		config["sign_name"],
		msg.Template,
		string(templateParam),
	)

	logger.DebugJSON("短信[阿里云]", "请求内容", smsClient.Request)
	logger.DebugJSON("短信[阿里云]", "接口响应", result)

	if err != nil {
		logger.ErrorString("短信[阿里云]", "发信失败", err.Error())
		return false
	}

	resultJson, err := json.Marshal(result)
	if err != nil {
		logger.ErrorString("短信[阿里云]", "解析响应 JSON 错误", err.Error())
		return false
	}

	if result.IsSuccessful() {
		logger.DebugString("短信[阿里云]", "发信成功", "")
		return true
	} else {
		logger.ErrorString("短信[阿里云]", "服务商返回错误", string(resultJson))
		return false
	}

}
