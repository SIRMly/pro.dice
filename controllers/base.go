/*
* @Author               : ruoli
* @Email                : ruooooooli@gmail.com
* @Date                 : 2016-06-20 15:06:44
* @Last Modified by     : ruoli
* @Last Modified time   : 2016-06-20 16:50:56
 */

package controllers

import (
	"github.com/astaxie/beego"
	"strings"
)

type BaseController struct {
	beego.Controller
}

func (b *BaseController) json(data interface{}) {
	b.Data["json"] = data
	b.ServeJSON()
	b.StopRun()
}

func (b *BaseController) respSuccess(message string, data ...interface{}) {
	out := make(map[string]interface{})
	out["code"] = "success"
	out["message"] = message
	if len(data) == 0 {
		out["data"] = ""
	} else {
		out["data"] = data
	}

	b.json(out)
}

func (b *BaseController) respError(message string, data ...interface{}) {
	out := make(map[string]interface{})
	out["code"] = "error"
	out["message"] = message
	if len(data) == 0 {
		out["data"] = ""
	} else {
		out["data"] = data
	}

	b.json(out)
}

func (b *BaseController) url(str string) string {
	rooturl := beego.AppConfig.String("rooturl")
	return rooturl + strings.TrimSpace(str)
}
