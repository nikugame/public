// Copyright 2016 Nikugame. All Rights Reserved

package sms

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"net/url"
	"sort"
	"strings"

	"github.com/nikgame/public/config"
	"github.com/nikgame/public/tools"

	"fmt"

	"github.com/axgle/mahonia"
)

//SendMessage func(phone, code string) error Send SMS function
func SendMessage(phone, code string) error {

	conf, _ := config.NewConfig("ini", "conf/settings.conf")
	key := conf.DefaultInt("SMS::type", 1)

	switch key {
	case 1:
		return beiwei(phone, code)
	case 2:
		return dayu(phone, code)
	default:
		return errors.New("error configure")
	}
}

//beiwei func(phone, code string) error send sms use beiwei
func beiwei(phone, code string) error {
	conf, _ := config.NewConfig("ini", "conf/settings.conf")

	u, _ := url.Parse(conf.String("BEIWEI::url"))
	q := u.Query()

	decoder := mahonia.NewEncoder("GBK")
	rrid := tools.BuilderRandomString(tools.NIKRandSringTypeAllCharacter, 8)

	q.Set("sn", conf.String("BEIWEI::sn"))
	q.Set("pwd", strings.ToUpper(tools.BuilderMD5(tools.StringJoin("", conf.String("BEIWEI::sn"), conf.String("BEIWEI::pwd")))))
	q.Set("mobile", phone)
	q.Set("content", decoder.ConvertString(strings.Replace(conf.String("BEIWEI::content"), "*", code, -1)))
	q.Set("ext", conf.String("BEIWEI::ext"))
	q.Set("stime", conf.DefaultString("BEIWEI::stime", ""))
	q.Set("rrid", rrid)

	u.RawQuery = q.Encode()
	res, err := http.Get(u.String())
	if err != nil {
		fmt.Printf("[%s] beiwei can't connect %s \n", tools.TimeStampString(), err)
		return err
	}
	body, err := ioutil.ReadAll(res.Body)
	defer res.Body.Close()
	if err != nil {
		fmt.Printf("[%s] body error: %s \n", tools.TimeStampString(), err)
		return err
	}
	if !strings.Contains(string(body), rrid) {
		return errors.New("system error")
	}
	return nil
}

//dayu func(phone, code string) error send sms use dayu
func dayu(phone, code string) error {
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
	m["sms_template_code"] = conf.String("ALIDAYU::template")
	m["sign"] = sign(m, conf.String("ALIDAYU::secert"))

	v := url.Values{}
	for k, values := range m {
		v.Set(k, values)
	}

	body := ioutil.NopCloser(strings.NewReader(v.Encode()))
	client := http.Client{}

	request, err := http.NewRequest("POST", conf.String("ALIDAYU::url"), body)
	if err != nil {
		fmt.Printf("[%s]dayu request error: %s \n", tools.TimeStampString(), err)
		return err
	}
	request.Header.Set("Content-Type", "application/x-www-form-urlencoded;charset=utf-8")
	response, err := client.Do(request)
	if err != nil {
		fmt.Printf("[%s]dayu response error: %s \n", tools.TimeStampString(), err)
		return err
	}
	b, err := ioutil.ReadAll(response.Body)
	defer response.Body.Close()
	if err != nil {
		fmt.Printf("[%s]dayu response body error: %s \n", tools.TimeStampString(), err)
		return err
	}
	var alim map[string]interface{}
	if err = json.Unmarshal(b, &alim); err != nil {
		fmt.Printf("[%s]dayu json error: %s \n", tools.TimeStampString(), err)
		return err
	}
	if _, found := alim["alibaba_aliqin_fc_sms_num_send_response"]; found {
		var res struct {
			Key struct {
				Code    string `json:"err_code"`
				Model   string `json:"model"`
				Success string `json:"success"`
				Message string `json:"msg"`
			} `json:"alibaba_aliqin_fc_sms_num_send_response"`
		}
		if err := json.Unmarshal(b, &res); err != nil {
			fmt.Printf("[%s]dayu success json error: %s, json: %s \n", tools.TimeStampString(), err, b)
			return err
		}
		switch res.Key.Code {
		case "0":
			return nil
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
			fmt.Printf("[%s]dayu error json error: %s, json: %s \n", tools.TimeStampString(), err, b)
			return err
		}
		return errors.New(res.Key.Message)
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
