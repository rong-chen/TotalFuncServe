package chat

import (
	"ChatServe/src/model/chatContent"
	"ChatServe/src/model/friends"
	userInfo "ChatServe/src/model/user"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	uuid "github.com/satori/go.uuid"
	"net/http"
	"sync"
)

var (
	offline   = "该用户不在线"
	paramsErr = "参数错误"
	uuidErr   = "uuid格式错误"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

type Node struct {
	conn    *websocket.Conn
	msgChan chan []byte
}

// websocket 参数结构体
type Params struct {
	FromUuid string      `json:"fromUuid" form:"fromUuid"`
	ToUuid   string      `json:"toUuid" form:"toUuid"`
	SendMsg  string      `json:"sendMsg"  form:"sendMsg"`
	MsgType  string      `json:"msgType"  form:"msgType"`
	ATList   []uuid.UUID `json:"at_list"  form:"at_list"`
}
type FirsConn struct {
	Token string `json:"token" form:"token"`
}

var rwLocker sync.RWMutex
var ConnectMap = make(map[uuid.UUID]*Node)

func Connect(c *gin.Context) {
	conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		return
	}
	node := Node{
		conn,
		make(chan []byte),
	}
	useID := c.Query("Id")
	uuidUserID, err := uuid.FromString(useID)
	if err != nil {
		conn.Close()
	}
	rwLocker.Lock()
	ConnectMap[uuidUserID] = &node
	rwLocker.Unlock()
	// 推送广播给所有已加好友的用户
	go PushAllFriends(uuidUserID)
	go recvMsg(node)
	go sendMsg(node)
}
func PushAllFriends(uuid uuid.UUID) {
	friendList := friends.FindAllFriends(uuid, "1")
	user := userInfo.FindUserInfoByUUID(uuid)
	strUser, err := json.Marshal(user)
	if err != nil {
		return
	}
	for _, f := range friendList {
		if nodes, ok := ConnectMap[f.FriendsId]; ok {
			nodes.conn.WriteMessage(websocket.TextMessage, strUser)
		}
	}
}

func recvMsg(node Node) {
	var params Params
	for {
		msgType, msg, err := node.conn.ReadMessage()
		if err != nil {
			node.conn.Close()
			return
		}
		errs := json.Unmarshal(msg, &params)
		if errs != nil {
			e := node.conn.WriteMessage(msgType, []byte(paramsErr))
			if e != nil {
				node.conn.Close()
			}
		}
		sendUuid := params.ToUuid
		uuids, err := uuid.FromString(sendUuid)
		if err != nil {
			e := node.conn.WriteMessage(msgType, []byte(uuidErr))
			if e != nil {
				node.conn.Close()
			}
		}
		//
		fid, errs2 := uuid.FromString(params.FromUuid)
		tid, errs2 := uuid.FromString(params.ToUuid)
		if errs2 != nil {
			node.conn.WriteMessage(msgType, []byte(offline))
		}
		//消息存储到数据库中
		var cc chatContent.ChatContent
		cc.ID = uuid.NewV4()
		cc.FriendId = tid
		cc.Content = params.SendMsg
		cc.UserId = fid
		err3 := chatContent.CreateChatContent(&cc)
		if err3 != nil {
			node.conn.WriteMessage(msgType, []byte("消息存储失败"))
		}
		if nodes, ok := ConnectMap[uuids]; ok {
			//用户在线
			nodes.msgChan <- msg
		} else {
			//否则 //用户不在线
			//e := node.conn.WriteMessage(msgType, []byte(offline))
		}
	}
}

func sendMsg(node Node) {
	for {
		select {
		case datas := <-node.msgChan:
			var params Params
			json.Unmarshal(datas, &params)
			switch params.MsgType {
			case "1":
				//私聊
				go privateSend(params)
				break
			case "2":
				go massHair(params)
				//群聊
				break
			case "3":
				//所有人公告
				break
			}
		}
	}
}

// 私人
func privateSend(params Params) {
	strUUid := params.ToUuid
	uuids, _ := uuid.FromString(strUUid)
	type BackParams struct {
		SendUuid string `json:"sendUuid" form:"sendUuid"`
		Msg      string `json:"msg" form:"msg"`
	}
	var bp BackParams
	bp.SendUuid = params.FromUuid
	bp.Msg = params.SendMsg
	if mapItem, ok := ConnectMap[uuids]; ok {
		byteBP, _ := json.Marshal(bp)
		mapItem.conn.WriteMessage(websocket.TextMessage, byteBP)
	}
}

// 群聊
func massHair(params Params) {
	//群聊
}
