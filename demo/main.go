package main

import (
	"fmt"
	"regexp"
)

func main() {
	phone := "13800138000"
	fmt.Println("原始手机号：", phone)
	fmt.Println("脱密后手机号：", maskPhone(phone))
}

func maskPhone(phone string) string {
	// 使用正则表达式匹配手机号
	reg := regexp.MustCompile(`^(1[3-9]\d{9})$`)
	if !reg.MatchString(phone) {
		return ""
	}

	// 将手机号脱密
	maskedPhone := phone[:3] + "****" + phone[7:]

	return maskedPhone
}
