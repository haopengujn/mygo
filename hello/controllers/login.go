package controllers

import (
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

func init() {
	orm.RegisterDriver("mysql", orm.DR_MySQL)
	orm.RegisterDataBase("default", "mysql", "moster:shzygjrmdwg@tcp(192.168.0.8:3306)/zt?charset=utf8")
	orm.RegisterModel(new(Feedback))
}

type Login struct {
	beego.Controller
}

type Feedback struct {
	Id      int
	Site    string
	Content string
	Ip      string
}

func (this *Login) GoLogin() {
	this.TplNames = "login.tpl"
}

func (this *Login) DoLogin() {
	// username := this.Ctx.Input.Params("username")
	// password := this.Ctx.Input.Params("password")
	username := this.GetString("username")
	password := this.GetString("password")
	fmt.Println("username: " + username)
	fmt.Println("password: " + password)
	msg := ""
	if len(username) == 0 {
		msg += "用户名不可以为空。<br/>"
	}
	if len(password) == 0 {
		msg += "密码不可以为空。<br/>"
	}
	if len(msg) == 0 && password == "123456" {
		msg += "登录成功。"
	} else {
		msg += "登录失败。"
	}
	this.Data["Msg"] = msg
	this.TplNames = "msg.tpl"
}

func (this *Login) Show() {
	o := orm.NewOrm()
	var feedbacks []Feedback
	num, err := o.Raw("select ID, SITE, CONTENT, IP from NEWS_FEEDBACK limit 5").QueryRows(&feedbacks)
	fmt.Println(num)
	fmt.Println(err)
	if err == nil {
		for _, feedback := range feedbacks {
			fmt.Println(feedback.Content)
		}
	}
	this.Data["Msg"] = "success"
	this.TplNames = "msg.tpl"
}
