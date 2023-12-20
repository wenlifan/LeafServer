package gate

import (
	"net"
)

type Agent interface {
	WriteMsg(msg interface{})
	WriteMsgBase(msg interface{}, msgName string)
	LocalAddr() net.Addr
	RemoteAddr() net.Addr
	Close()
	Destroy()
	UserData() interface{}
	SetUserData(data interface{})
}
