package gate

import (
	"encoding/binary"
	"github.com/zhanglifan/leaf_server/leaf/chanrpc"
	"github.com/zhanglifan/leaf_server/leaf/log"
	"github.com/zhanglifan/leaf_server/leaf/network"
	"net"
	"reflect"
	"time"
)

type Gate struct {
	MaxConnNum      int
	PendingWriteNum int
	MaxMsgLen       uint32
	Processor       network.Processor
	AgentChanRPC    *chanrpc.Server

	// websocket
	WSAddr      string
	HTTPTimeout time.Duration
	CertFile    string
	KeyFile     string

	// tcp
	TCPAddr      string
	LenMsgLen    int
	LittleEndian bool
}

func (gate *Gate) Run(closeSig chan bool) {
	var wsServer *network.WSServer
	if gate.WSAddr != "" {
		wsServer = new(network.WSServer)
		wsServer.Addr = gate.WSAddr
		wsServer.MaxConnNum = gate.MaxConnNum
		wsServer.PendingWriteNum = gate.PendingWriteNum
		wsServer.MaxMsgLen = gate.MaxMsgLen
		wsServer.HTTPTimeout = gate.HTTPTimeout
		wsServer.CertFile = gate.CertFile
		wsServer.KeyFile = gate.KeyFile
		wsServer.NewAgent = func(conn *network.WSConn) network.Agent {
			a := &agent{conn: conn, gate: gate}
			if gate.AgentChanRPC != nil {
				gate.AgentChanRPC.Go("NewAgent", a)
			}
			return a
		}
	}

	var tcpServer *network.TCPServer
	if gate.TCPAddr != "" {
		tcpServer = new(network.TCPServer)
		tcpServer.Addr = gate.TCPAddr
		tcpServer.MaxConnNum = gate.MaxConnNum
		tcpServer.PendingWriteNum = gate.PendingWriteNum
		tcpServer.LenMsgLen = gate.LenMsgLen
		tcpServer.MaxMsgLen = gate.MaxMsgLen
		tcpServer.LittleEndian = gate.LittleEndian
		tcpServer.NewAgent = func(conn *network.TCPConn) network.Agent {
			a := &agent{conn: conn, gate: gate}
			if gate.AgentChanRPC != nil {
				gate.AgentChanRPC.Go("NewAgent", a)
			}
			return a
		}
	}

	if wsServer != nil {
		wsServer.Start()
	}
	if tcpServer != nil {
		tcpServer.Start()
	}
	<-closeSig
	if wsServer != nil {
		wsServer.Close()
	}
	if tcpServer != nil {
		tcpServer.Close()
	}
}

func (gate *Gate) OnDestroy() {}

type agent struct {
	conn           network.Conn
	gate           *Gate
	userData       interface{}
	receivedHeader bool // 首条消息验证
}

func int32ToBytes(i int32) []byte {
	buf := make([]byte, 4)
	binary.BigEndian.PutUint32(buf, uint32(i))
	return buf
}

func reverse(s []byte) []byte {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
	return s
}

func bytesToInt32(buf []byte) int32 {
	return int32(binary.BigEndian.Uint32(buf))
}

func (a *agent) checkReceivedHeader(data []byte) (bool, error) {
	//var , nVersion, nNodeIDA, nNodeIDB, nNodeIDC, nNodeIDD int32
	nMagicNumber := bytesToInt32(reverse(data[0:4]))
	nVersion := bytesToInt32(reverse(data[4:8]))
	// 验证 参数
	log.Release("checkReceivedHeader Client nMagicNumber: %d", nMagicNumber)
	log.Release("checkReceivedHeader Client nVersion: %d", nVersion)

	//retData := make([]byte, 24)
	var key int32 = 0
	rM := reverse(int32ToBytes(nMagicNumber))
	for _, value := range rM {
		data[key] = value
		key += 1
	}
	rV := reverse(int32ToBytes(nVersion))
	for _, value := range rV {
		data[key] = value
		key += 1
	}
	a.conn.WriteMsg(data)

	a.receivedHeader = true
	return true, nil
}

func (a *agent) Run() {
	for {
		data, err := a.conn.ReadMsg()
		if err != nil {
			log.Debug("read message: %v", err)
			break
		}
		if !a.receivedHeader {
			//	首条消息验证
			result, err := a.checkReceivedHeader(data)
			if !result {
				log.Debug("check received header message error: %v", err)
				break
			}
		} else {
			if a.gate.Processor != nil {
				msg, err := a.gate.Processor.UnmarshalBase(data)
				if err != nil {
					log.Debug("unmarshal message error: %v", err)
					break
				}
				err = a.gate.Processor.Route(msg, a)
				if err != nil {
					log.Debug("route message error: %v", err)
					break
				}
			}
		}
	}
}

func (a *agent) OnClose() {
	if a.gate.AgentChanRPC != nil {
		err := a.gate.AgentChanRPC.Call0("CloseAgent", a)
		if err != nil {
			log.Error("chanrpc error: %v", err)
		}
	}
}

func (a *agent) WriteMsg(msg interface{}) {
	if a.gate.Processor != nil {
		data, err := a.gate.Processor.Marshal(msg)
		if err != nil {
			log.Error("marshal message %v error: %v", reflect.TypeOf(msg), err)
			return
		}
		err = a.conn.WriteMsg(data...)
		if err != nil {
			log.Error("write message %v error: %v", reflect.TypeOf(msg), err)
		}
	}
}

func (a *agent) WriteMsgBase(msg interface{}) {
	if a.gate.Processor != nil {
		data, err := a.gate.Processor.MarshalBase(msg)
		if err != nil {
			log.Error("marshal message %v error: %v", reflect.TypeOf(msg), err)
			return
		}
		err = a.conn.WriteMsg(data...)
		if err != nil {
			log.Error("write message %v error: %v", reflect.TypeOf(msg), err)
		}
	}
}

func (a *agent) LocalAddr() net.Addr {
	return a.conn.LocalAddr()
}

func (a *agent) RemoteAddr() net.Addr {
	return a.conn.RemoteAddr()
}

func (a *agent) Close() {
	a.conn.Close()
}

func (a *agent) Destroy() {
	a.conn.Destroy()
}

func (a *agent) UserData() interface{} {
	return a.userData
}

func (a *agent) SetUserData(data interface{}) {
	a.userData = data
}
