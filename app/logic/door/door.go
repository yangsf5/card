// Author: sheppard(ysf1026@gmail.com) 2013-12-13

package door

import (
	"time"
	"code.google.com/p/go.net/websocket"
	"github.com/yangsf5/card/app/logic/hall"
	"github.com/yangsf5/card/app/logic/user"
)

func init() {
	hall.ReadConfig("conf/room.xml")
}

func Login(conn *websocket.Conn, userName string) {
	offline := make(chan error)

	recvMsg := make(chan string)
	go func() {
		var msg string
		for {
			err := websocket.Message.Receive(conn, &msg)
			if err != nil {
				close(recvMsg)
				return
			}
			recvMsg <- msg
		}
	}()

	sendMsg := make(chan string)
	go func() {
		for {
			select {
			case msg, ok := <-sendMsg:
				// If the channel is closed, they disconnected.
				if !ok {
					return
				}

				if err := websocket.Message.Send(conn, msg); err != nil {
					// Disconneted.
					offline <- err
					return
				}
			}
		}
	}()

	u := user.NewUser(userName, recvMsg, sendMsg, offline)
	ok := hall.Enter(u.Name, u)
	if !ok {
		return
	}
	u.Login()
	go u.Tick()

	for {
		time.Sleep(100 * time.Millisecond)
	}
}

