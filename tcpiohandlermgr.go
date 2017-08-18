package netlib

import "net"

type TcpIoHandlerMgr interface {
	OnAccept(net.Conn, *NetlibTcpServer)
	OnClose(net.Conn, *NetlibTcpServer)
}
