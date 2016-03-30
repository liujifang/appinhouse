package main

import (
	. "appinhouse/service/constants"

	_ "appinhouse/service/routers"

	"appinhouse/service/controllers"

	"strings"

	"github.com/astaxie/beego"
)

func main() {
	setParam()
	controllers.InitDirectory()
	beego.SetLevel(beego.LevelInformational)
	beego.SetLogFuncCall(true)
	beego.SetLogger("file", `{"filename":"appinhouse.log"}`)
	beego.Run()
}
func setParam() {
	Root_Dir = beego.AppConfig.String("users::root_dir")
	beego.Info("app.conf-> Root_Dir:", Root_Dir)
	if Root_Dir == "" {
		panic("app.conf not have users::root_dir ")
	}
	Ios_Channel = beego.AppConfig.String("users::ios_channel")
	beego.Info("app.conf-> Ios_Channel:", Ios_Channel)
	if Ios_Channel == "" {
		panic("app.conf not have users::ios_channel ")
	}
	Domain = beego.AppConfig.String("users::domain")
	beego.Info("app.conf-> Domain:", Domain)
	if Domain == "" {
		panic("app.conf not have users::domain ")
	}
	apps := beego.AppConfig.String("users::apps")
	beego.Info("app.conf-> apps:", apps)
	if apps == "" {
		panic("app.conf not have users::apps ")
	}

	AddApps(strings.Split(apps, ";"))
	Full_Size_Image = beego.AppConfig.String("users::full_size_image")
	beego.Info("app.conf-> Full_Size_Image:", Full_Size_Image)
	if Full_Size_Image == "" {
		panic("app.conf not have users::full_size_image ")
	}
	Display_Image = beego.AppConfig.String("users::display_image")
	beego.Info("app.conf-> Display_Image:", Display_Image)
	if Display_Image == "" {
		panic("app.conf not have users::display_image ")
	}
	Min_Residue, _ = beego.AppConfig.Int("users::min_residue")
	beego.Info("app.conf-> Min_Residue:", Min_Residue)
	if Min_Residue == 0 {
		panic("app.conf not have users::min_residue or not int")
	}
	Page_Size, _ = beego.AppConfig.Int("users::page_size")
	beego.Info("app.conf-> Page_Size:", Page_Size)
	if Page_Size == 0 {
		panic("app.conf not have users::page_size or not int")
	}
	Max_Page, _ = beego.AppConfig.Int("users::max_page")
	beego.Info("app.conf-> Max_Page:", Max_Page)
	if Max_Page == 0 {
		panic("app.conf not have users::max_page or not int")
	}
}
