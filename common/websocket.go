package common

import (
	"github.com/google/uuid"
	"github.com/gorilla/websocket"
	"github.com/pkg/errors"
	"strings"
)

type Connection struct {
	Conn      *websocket.Conn
	SessionId string
	UserId    int64
	RoleId    int64
	inchan    chan []byte
	outchan   chan []byte
	exitchan  chan bool
	IsClose   bool
}

func NewConnection(conn *websocket.Conn, uid, rid int64) (*Connection, error) {
	sessionStr := uuid.New().String()
	sessionId := strings.Replace(sessionStr, "-", "", -1)
	wsConn := &Connection{
		Conn:      conn,
		SessionId: sessionId,
		UserId:    uid,
		RoleId:    rid,
		inchan:    make(chan []byte, 1000),
		outchan:   make(chan []byte, 1000),
		exitchan:  make(chan bool, 1),
	}
	return wsConn, nil
}

func (conn *Connection) Start() {
	// 不停的读长连接的消息，放到inchan中
	go conn.readLoop()
	// 不停的从outchan,读取消息，发到客户端
	go conn.writeLoop()
}

func (conn *Connection) ReadMessage() ([]byte, error) {
	if conn.IsClose {
		return nil, errors.New("connection is closed")
	}
	data := <-conn.inchan
	return data, nil
}
func (conn *Connection) WriteMessage(data []byte) error {
	if conn.IsClose {
		return errors.New("connection is closed")
	}
	conn.outchan <- data
	return nil
}

func (conn *Connection) Close() {
	if conn.IsClose == true {
		return
	}
	conn.IsClose = true
	// 线程安全，可重入的close
	conn.Conn.Close()
	close(conn.exitchan)
	close(conn.inchan)
	close(conn.outchan)
}

func (conn *Connection) readLoop() {
	defer conn.Close()
	for {
		_, data, err := conn.Conn.ReadMessage()
		if err != nil {
			conn.exitchan <- true
			conn.IsClose = true
			return
		}
		conn.inchan <- data
	}
}

func (conn *Connection) writeLoop() {
	for {
		select {
		case data := <-conn.outchan:
			if err := conn.Conn.WriteMessage(websocket.TextMessage, data); err != nil {
				conn.exitchan <- true
				return
			}
		case <-conn.exitchan:
			return
		}
	}
}
