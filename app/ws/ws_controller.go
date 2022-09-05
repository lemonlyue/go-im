package ws

import (
	"gin-skeleton/pkg/console"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"net/http"
)

type WsController struct {

}

func (ws *WsController) WsClient(c *gin.Context) {
	upGrande := websocket.Upgrader{
		CheckOrigin: func(r *http.Request) bool {
			return true
		},
		Subprotocols: []string{
			c.GetHeader("Sec-WebSocket-Protocol"),
		},
	}

	conn, err := upGrande.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		console.Error("websocket connection error: " + err.Error())
		return
	}
	defer conn.Close()
	for {
		_, message, err := conn.ReadMessage()
		if err != nil {
			console.Error("websocket read message error: " + err.Error())
			break
		}
		err = conn.WriteMessage(websocket.TextMessage, message)
		if err != nil {
			console.Error("websocket write message error: " + err.Error())
			break
		}
	}
}