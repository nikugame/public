// Copyright 2016 Nikugame. All Rights Reserved

package tools

import (
	"crypto/md5"
	"encoding/hex"
	"math/rand"
	"strings"
	"time"
)

const (
	//NIKRandSringTypeOnlyNumber 纯数字模式
	NIKRandSringTypeOnlyNumber int = 1
	//NIKRandSringTypeLowCharacter 数字与小写字母组合模式
	NIKRandSringTypeLowCharacter int = 2
	//NIKRandSringTypeAllCharacter 数字与大小写字母组合模式
	NIKRandSringTypeAllCharacter int = 3
)

//BuilderRandomString func(stringType, length int) string 根据生成模式与长度构建随机字符串
func BuilderRandomString(stringType, length int) string {

	typeNumber := "0123456789"
	typeLowCharacter := "abcdefghijklmnopqrstuvwxyz"
	typeUpCharacter := "ABCDEFGHIJKLMNOPQRSTUVWXYZ"

	random := ""

	switch stringType {
	case NIKRandSringTypeOnlyNumber:
		random = typeNumber
	case NIKRandSringTypeLowCharacter:
		random = StringJoin("", typeNumber, typeLowCharacter)
	case NIKRandSringTypeAllCharacter:
		random = StringJoin("", typeNumber, typeLowCharacter, typeUpCharacter)
	default:
		random = typeNumber
	}

	bytes := []byte(random)
	res := []byte{}
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < length; i++ {
		res = append(res, bytes[r.Intn(len(bytes))])
	}
	return string(res)
}

//StringJoin func(split string, str...string) string  字符串拼接封装方法
func StringJoin(split string, str ...string) string {
	list := []string(str)
	return strings.Join(list, split)
}

//TimeStampString func() string 获取当前精确到秒的时间字符串,样式：2016-10-10 15:45:30
func TimeStampString() string {
	return time.Now().Format("2006-01-02 15:04:05")
}

//BuilderMD5 func(str string) string MD5加密方法
func BuilderMD5(str string) string {
	ctx := md5.New()
	ctx.Write([]byte(str))
	s := ctx.Sum(nil)
	return hex.EncodeToString(s)
}
