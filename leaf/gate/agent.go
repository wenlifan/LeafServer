package gate

import (
	"net"
)

type Agent interface {
	WriteMsg(msg interface{})
	WriteMsgBase(msg interface{})
	LocalAddr() net.Addr
	RemoteAddr() net.Addr
	Close()
	Destroy()
	UserData() interface{}
	SetUserData(data interface{})
}
