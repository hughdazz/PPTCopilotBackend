package websocket

import (
	"github.com/beego/beego/logs"
	beego "github.com/beego/beego/v2/server/web"
	"github.com/gorilla/websocket"
)

type Controller struct {
	beego.Controller
}

var upgrader = websocket.Upgrader{} // WebSocket升级器

type WSRequest struct {
	Type string                 `json:"type"`
	Data map[string]interface{} `json:"data"`
}

func (this *Controller) Join() {
	conn, err := upgrader.Upgrade(this.Ctx.ResponseWriter, this.Ctx.Request, nil)
	if err != nil {
		logs.Error("升级为WebSocket失败:", err)
		return
	}
	defer conn.Close()

	for {
		var req WSRequest
		// 读取客户端消息
		err := conn.ReadJSON(&req)
		if err != nil {
			logs.Error("读取客户端消息失败:", err)
		}

		if req.Type == "gen_slide" {
			logs.Info("收到gen_slide消息:", req.Data)
		}
		conn.WriteJSON(map[string]interface{}{
			"msg": "ok",
		})
	}
}
