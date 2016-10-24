// Copyright 2016 Nikugame. All Rights Reserved

package sms

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"sort"
	"strings"

	"github.com/axgle/mahonia"
	"github.com/nikugame/public/config"
	"github.com/nikugame/public/tools"
)

//SendMessage func(phone, code string) error Send SMS function
// func SendMessage(phone, code string) error {

// 	conf, _ := config.NewConfig("ini", "conf/settings.conf")
// 	key := conf.DefaultInt("SMS::type", 1)

// 	switch key {
// 	case 1:
// 		return beiwei(phone, code)
// 	case 2:
// 		return dayu(phone, code)
// 	case 3:
// 		return xiao(phone, code)
// 	default:
// 		return errors.New("error configure")
// 	}
// }

//SendMessage func(str ..string) error phone, code, channel Send SMS function
func SendMessage(str ...string) error {

	list := []string(str)
	if len(list) < 2 {
		return errors.New("param error : 1:phone, 2:code, must, 3:channel, can use default")
	}
	phone := list[0]
	code := list[1]
	channel := "C0"
	if len(list) > 2 {
		channel = tools.StringJoin("", "C", list[2])
	}

	conf, _ := config.NewConfig("ini", "conf/settings.conf")
	key := conf.DefaultInt("SMS::type", 1)

	switch key {
	case 1:
		return beiwei(phone, code, channel)
	case 2:
		return dayu(phone, code, channel)
	case 3:
		return xiao(phone, code, channel)
	default:
		return errors.New("error configure")
	}
}

//beiwei func(phone, code, channel string) error send sms use beiwei
func xiao(phone, code, channel string) error {
	conf, _ := config.NewConfig("ini", "conf/settings.conf")

	u, _ := url.Parse(conf.String("XIAO::url"))
	q := u.Query()

	q.Set("uid", conf.String("XIAO::uid"))
	q.Set("auth", tools.BuilderMD5(tools.StringJoin("", conf.String("XIAO::cid"), conf.String("XIAO::pwd"))))
	q.Set("expid", "0")
	q.Set("encode", "utf-8")
	decoder := mahonia.NewEncoder("utf-8")
	q.Set("msg", decoder.ConvertString(strings.Replace(conf.String(tools.StringJoin("::", "XIAO", channel)), "*", code, -1)))
	q.Set("mobile", phone)

	u.RawQuery = q.Encode()
	response, err := http.Get(u.String())
	if err != nil {
		return err
	}
	defer response.Body.Close()
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return err
	}
	res := strings.Split(string(body), ",")
	if strings.EqualFold(res[0], "0") {
		return nil
	}
	return errors.New("sms error")
}

//beiwei func(phone, code, channel string) error send sms use beiwei
func beiwei(phone, code, channel string) error {
	conf, _ := config.NewConfig("ini", "conf/settings.conf")

	u, _ := url.Parse(conf.String("BEIWEI::url"))
	q := u.Query()

	decoder := mahonia.NewEncoder("GBK")
	rrid := tools.BuilderRandomString(tools.NIKRandSringTypeAllCharacter, 8)

	q.Set("sn", conf.String("BEIWEI::sn"))
	q.Set("pwd", strings.ToUpper(tools.BuilderMD5(tools.StringJoin("", conf.String("BEIWEI::sn"), conf.String("BEIWEI::pwd")))))
	q.Set("mobile", phone)
	q.Set("content", decoder.ConvertString(strings.Replace(conf.String(tools.StringJoin("::", "BEIWEI", channel)), "*", code, -1)))
	q.Set("ext", conf.String("BEIWEI::ext"))
	q.Set("stime", conf.DefaultString("BEIWEI::stime", ""))
	q.Set("rrid", rrid)

	u.RawQuery = q.Encode()
	res, err := http.Get(u.String())
	if err != nil {
		return err
	}
	body, err := ioutil.ReadAll(res.Body)
	defer res.Body.Close()
	if err != nil {
		return err
	}
	if !strings.Contains(string(body), rrid) {
		return errors.New("system error")
	}
	return nil
}

//dayu func(phone, code, channel string) error send sms use dayu
func dayu(phone, code, channel string) error {
	conf, _ := config.NewConfig("ini", "conf/settings.conf")

	var param struct {
		Code    string `json:"code"`
		Product string `json:"product"`
	}
	param.Code = code
	param.Product = conf.String("ALIDAYU::name")
	str, _ := json.Marshal(param)

	m := make(map[string]string)
	m["app_key"] = conf.String("ALIDAYU::key")
	m["format"] = "json"
	m["timestamp"] = tools.TimeStampString()
	m["v"] = "2.0"
	m["method"] = "alibaba.aliqin.fc.sms.num.send"
	m["sign_method"] = "md5"
	m["sms_type"] = "normal"
	m["sms_free_sign_name"] = conf.String("ALIDAYU::sign")
	m["sms_param"] = string(str)
	m["rec_num"] = phone
	m["sms_template_code"] = conf.String(tools.StringJoin("::", "ALIDAYU", channel))
	m["sign"] = sign(m, conf.String("ALIDAYU::secert"))

	v := url.Values{}
	for k, values := range m {
		v.Set(k, values)
	}

	body := ioutil.NopCloser(strings.NewReader(v.Encode()))
	client := http.Client{}

	request, err := http.NewRequest("POST", conf.String("ALIDAYU::url"), body)
	if err != nil {
		return err
	}
	request.Header.Set("Content-Type", "application/x-www-form-urlencoded;charset=utf-8")
	response, err := client.Do(request)
	if err != nil {
		return err
	}
	b, err := ioutil.ReadAll(response.Body)
	defer response.Body.Close()
	if err != nil {
		return err
	}
	var alim map[string]interface{}
	if err = json.Unmarshal(b, &alim); err != nil {
		return err
	}
	// fmt.Printf("ALIM : %v \n", alim)
	if _, found := alim["alibaba_aliqin_fc_sms_num_send_response"]; found {
		var res struct {
			Key struct {
				Result struct {
					Code    string `json:"err_code"`
					Model   string `json:"model"`
					Success bool   `json:"success"`
				} `json:"result"`
				ID string `json:"request_id"`
			} `json:"alibaba_aliqin_fc_sms_num_send_response"`
		}
		if err := json.Unmarshal(b, &res); err != nil {
			return err
		}

		switch res.Key.Result.Code {
		case "0":
			return nil
		default:
			return errors.New(res.Key.Result.Code)
		}
	}
	if _, found := alim["error_response"]; found {
		var res struct {
			Key struct {
				Code       int    `json:"code"`
				Message    string `json:"msg"`
				SubCode    string `json:"sub_code"`
				SubMessage string `json:"sub_msg"`
			} `json:"error_response"`
		}
		if err := json.Unmarshal(b, &res); err != nil {
			return err
		}
		switch res.Key.SubCode {
		case "isv.OUT_OF_SERVICE":
			fmt.Println("短信业务停机，要充值！")
			return errors.New("System ERROR, Try later!")
		case "isv.PRODUCT_UNSUBSCRIBE":
			fmt.Println("产品服务未开通，参数配置错误")
			return errors.New("System ERROR, Try later!")
		case "isv.ACCOUNT_NOT_EXISTS":
			fmt.Println("账户信息不存在，参数配置错误")
			return errors.New("System ERROR, Try later!")
		case "isv.ACCOUNT_ABNORMAL":
			fmt.Println("账户信息异常，联系大于客服")
			return errors.New("System ERROR, Try later!")
		case "isv.SMS_TEMPLATE_ILLEGAL":
			fmt.Println("模板不合法， 参数配置错误")
			return errors.New("System ERROR, Try later!")
		case "isv.SMS_SIGNATURE_ILLEGAL":
			fmt.Println("签名不合法，参数配置错误")
			return errors.New("System ERROR, Try later!")
		case "isv.MOBILE_NUMBER_ILLEGAL":
			fmt.Println("手机号码格式错误，检查是否存在漏洞，被非法调用 ")
			return errors.New("Bad Request, Must phone number!")
		case "isv.MOBILE_COUNT_OVER_LIMIT":
			fmt.Println("手机号码数量超过限制, 不能超过200个号码")
			return errors.New("Too many phone number!")
		case "isv.TEMPLATE_MISSING_PARAMETERS":
			fmt.Println("短信模板变量缺少参数，配置错误")
			return errors.New("System ERROR, Try later!")
		case "isv.INVALID_PARAMETERS":
			fmt.Println("参数异常，配置错误")
			return errors.New("System ERROR, Try later!")
		case "isv.BUSINESS_LIMIT_CONTROL":
			fmt.Printf("%s : 超过1分钟1条 或者 1分钟7条 1天50条限制 \n", phone)
			return errors.New("Send too often, Please try 1 min later")
		case "isv.INVALID_JSON_PARAM":
			fmt.Println("JSON参数不合法，代码被改了？")
			return errors.New("System ERROR, Try later!")
		case "isp.SYSTEM_ERROR":
			fmt.Println("大于挂了！改用其他通道发送")
			return beiwei(phone, code, channel)
		case "isv.BLACK_KEY_CONTROL_LIMIT":
			fmt.Println("模板变量中存在黑名单关键字，大于后台配置错")
			return errors.New("System ERROR, Try later!")
		case "isv.PARAM_NOT_SUPPORT_URL":
			fmt.Println("不支持url为变量，配置错误")
			return errors.New("System ERROR, Try later!")
		case "isv.PARAM_LENGTH_LIMIT":
			fmt.Println("变量长度受限，配置错误")
			return errors.New("System ERROR, Try later!")
		case "isv.AMOUNT_NOT_ENOUGH":
			fmt.Println("余额不足，要充值！")
			return errors.New("System ERROR, Try later!")
		}
	}
	return nil
}

//sign func(map[string]string, string) string
func sign(m map[string]string, secret string) string {

	list := []string{}
	for k := range m {
		list = append(list, k)
	}
	sort.Strings(list)

	param := ""
	for i := 0; i < len(list); i++ {
		param = tools.StringJoin("", param, list[i], m[list[i]])
	}
	str := tools.StringJoin("", secret, param, secret)
	return strings.ToUpper(tools.BuilderMD5(str))

}
