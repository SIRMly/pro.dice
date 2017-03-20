/*
* @Author               : ruoli
* @Email                : ruooooooli@gmail.com
* @Date                 : 2016-06-20 15:17:06
* @Last Modified by     : ruooooooli
* @Last Modified time   : 2016-06-22 19:10:42
 */

package controllers

import (
	"fmt"
	"github.com/astaxie/beego/cache"
	"github.com/chanxuehong/wechat.v2/mp/core"
	"github.com/chanxuehong/wechat.v2/mp/jssdk"
	"github.com/gorilla/websocket"
	"github.com/ruooooooli/utils"
	"github.com/satori/go.uuid"
	qrcode "github.com/skip2/go-qrcode"
	"strconv"
	"time"
)

const (
	StatusIn = iota
	StatusReady
	StatusBegin
	StatusEnd
)

const (
	TheUserOne = "userone"
	TheUserTwo = "usertwo"
)

type Room struct {
	UserOne   *websocket.Conn
	UserTwo   *websocket.Conn
	StatusOne int
	StatusTwo int
}

var (
	appid             = "wxd8da84ed2a26aa06"
	appsecret         = "00e6fd3ce1151e3d2bd0e01c98c925d3"
	roomlist, _       = cache.NewCache("memory", `{"interval":60}`)
	accessTokenServer = core.NewDefaultAccessTokenServer(appid, appsecret, nil)
	client            = core.NewClient(accessTokenServer, nil)
	tickerServer      = jssdk.NewDefaultTicketServer(client)
)

type IndexController struct {
	BaseController
}

// 用户打开二维码的页面
func (i *IndexController) Index() {
	name := fmt.Sprintf("%s", uuid.NewV4())
	room := Room{}
	setRoom(name, &room)
	i.SetSession(TheUserOne, TheUserOne)

	jsapiticket, _ := tickerServer.Ticket()
	nonceStr := utils.RandomString(16)
	timestamp := fmt.Sprintf("%d", time.Now().Unix())
	url := i.Ctx.Input.Site() + i.Ctx.Input.URI()
	signature := jssdk.WXConfigSign(jsapiticket, nonceStr, timestamp, url)

	i.Data["Jsapiticket"] = jsapiticket
	i.Data["Timestamp"] = timestamp
	i.Data["Signature"] = signature
	i.Data["NonceStr"] = nonceStr
	i.Data["Appid"] = appid
	i.Data["Name"] = name
	i.TplName = "index.tpl"
}

// 新建游戏
func (i *IndexController) Open() {
	name := i.Ctx.Input.Param(":name")
	room := getRoom(name)

	ws, err := websocket.Upgrade(i.Ctx.ResponseWriter, i.Ctx.Request, nil, 1024, 1024)
	if err != nil {
		i.respError(err.Error())
	}

	room.UserOne = ws
	room.StatusOne = StatusReady
	setRoom(name, room)
	defer ws.Close()

	for {
		_, _, err := ws.ReadMessage()
		if err != nil {
			return
		}
	}
}

// 显示游戏页面
func (i *IndexController) Game() {
	name := i.Ctx.Input.Param(":name")
	room := getRoom(name)
	if room == nil {
		i.Redirect("/", 302)
	}

	user := i.GetString("user")
	if user == TheUserTwo {
		i.SetSession(TheUserOne, TheUserTwo)
	}

	jsapiticket, _ := tickerServer.Ticket()
	nonceStr := utils.RandomString(16)
	timestamp := fmt.Sprintf("%d", time.Now().Unix())
	url := i.Ctx.Input.Site() + i.Ctx.Input.URI()
	signature := jssdk.WXConfigSign(jsapiticket, nonceStr, timestamp, url)

	i.Data["Jsapiticket"] = jsapiticket
	i.Data["Timestamp"] = timestamp
	i.Data["Signature"] = signature
	i.Data["NonceStr"] = nonceStr
	i.Data["Appid"] = appid

	i.Data["Name"] = name
	i.TplName = "game.tpl"
}

// 用户加入
func (i *IndexController) Join() {
	name := i.GetString(":name")
	room := getRoom(name)

	ws, err := websocket.Upgrade(i.Ctx.ResponseWriter, i.Ctx.Request, nil, 1024, 1024)
	if err != nil {
		fmt.Println(err.Error())
	}

	if i.isUserOne() {
		room.UserOne = ws
		room.StatusOne = StatusReady
	} else {
		room.UserTwo = ws
		room.StatusTwo = StatusReady
	}
	setRoom(name, room)

	for {
		messageType, p, err := ws.ReadMessage()
		if err != nil {
			return
		}

		status, _ := strconv.Atoi(string(p))
		if i.isUserOne() {
			switch status {
			case StatusIn:
				room.StatusOne = StatusIn
				// 页面加载完成
			case StatusReady:
				room.StatusOne = StatusReady
			case StatusBegin:
				room.StatusOne = StatusBegin
				if room.StatusOne == StatusBegin && room.StatusTwo == StatusBegin {
					if err := room.UserOne.WriteMessage(messageType, []byte(strconv.Itoa(StatusEnd))); err != nil {
						fmt.Println(err.Error())
					} else {
						room.StatusOne = StatusEnd
					}

					if err := room.UserTwo.WriteMessage(messageType, []byte(strconv.Itoa(StatusEnd))); err != nil {
						fmt.Println(err.Error())
					} else {
						room.StatusTwo = StatusEnd
					}
				}
			case StatusEnd:
				room.StatusOne = StatusReady

			}
		} else {
			switch status {
			//用户二进入 给用户一发送进入消息
			case StatusIn:
				room.StatusTwo = StatusIn
				if err := room.UserOne.WriteMessage(messageType, []byte(strconv.Itoa(StatusIn))); err != nil {
					fmt.Println(166, err.Error())
				}
				// 页面加载完成
			case StatusReady:
				room.StatusTwo = StatusReady
			case StatusBegin:
				room.StatusTwo = StatusBegin
				if room.StatusOne == StatusBegin && room.StatusTwo == StatusBegin {
					if err := room.UserOne.WriteMessage(messageType, []byte(strconv.Itoa(StatusEnd))); err != nil {
						fmt.Println(err.Error())
					} else {
						room.StatusOne = StatusEnd
					}

					if err := room.UserTwo.WriteMessage(messageType, []byte(strconv.Itoa(StatusEnd))); err != nil {
						fmt.Println(err.Error())
					} else {
						room.StatusTwo = StatusEnd
					}
				}
			case StatusEnd:
				room.StatusTwo = StatusEnd
			}
		}

	}
}

// 获取二维码图片
func (i *IndexController) Image() {
	name := i.Ctx.Input.Param(":name")
	png, err := qrcode.Encode(i.url("game/"+name+"?user=usertwo"), qrcode.Highest, 400)
	if err != nil {
		i.Ctx.WriteString(err.Error())
		i.StopRun()
	}
	i.Ctx.WriteString(string(png))
}

// 获取房间
func getRoom(name string) *Room {
	cachedata := roomlist.Get(name)
	room, ok := cachedata.(*Room)
	if ok {
		return room
	}
	return nil
}

// 设置房间
func setRoom(name string, room *Room) {
	roomlist.Put(name, room, 3600*time.Second)
}

func (i *IndexController) isUserOne() bool {
	session := i.GetSession(TheUserOne)
	if session == nil {
		return false
	}

	key := session.(string)
	return key == TheUserOne
}
