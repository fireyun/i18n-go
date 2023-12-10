package main

import (
	"fmt"

	"github.com/fireyun/i18n-go/pkg/i18n/localizer"
)

func main() {
	languages := []string{"zh", "zh-CN", "cmn", "en", "en-US", "bad-lang"}
	var name = "fireyun"
	var age = 20
	for _, lang := range languages {
		l := localizer.Get(lang)
		fmt.Println("lang:", lang)
		fmt.Println(l.Translate("Welcome!"))
		fmt.Println(l.Translate("Who are you? How old are you?"))
		fmt.Println(l.Translate("My name is %s, I'm %d years old.", name, age))
		fmt.Println("*****************************")
	}
	/*
		程序运行输出结果如下:
		lang: zh
		欢迎！
		你是谁？你多大了？
		我的名字叫fireyun，我20岁了。
		*****************************
		lang: zh-CN
		欢迎！
		你是谁？你多大了？
		我的名字叫fireyun，我20岁了。
		*****************************
		lang: cmn
		欢迎！
		你是谁？你多大了？
		我的名字叫fireyun，我20岁了。
		*****************************
		lang: en
		Welcome!
		Who are you? How old are you?
		My name is fireyun, I'm 20 years old.
		*****************************
		lang: en-US
		Welcome!
		Who are you? How old are you?
		My name is fireyun, I'm 20 years old.
		*****************************
		lang: bad-lang
		Welcome!
		Who are you? How old are you?
		My name is fireyun, I'm 20 years old.
		*****************************
	*/
}
