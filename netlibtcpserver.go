package netlib

import "net"
import "fmt"

type NetlibTcpServer struct {
	listener   net.Listener
	handlerMgr TcpIoHandlerMgr
}

func (s *NetlibTcpServer) Init(addr string, handlerMgr TcpIoHandlerMgr) bool {

	var err error
	s.listener, err = net.Listen("tcp", addr)
	if err != nil {
		return false
	}

	s.handlerMgr = handlerMgr
	go s.connAccept()

	return true
}

func (s *NetlibTcpServer) connAccept() {

	for {
		conn, err := s.listener.Accept()
		if err != nil {
			fmt.Println("Accept error!")
		}

		s.handlerMgr.OnAccept(conn, s)
	}
}

func (s *NetlibTcpServer) Close() {

	s.listener.Close()

}

func (s *NetlibTcpServer) CreateSession(conn net.Conn, ioHandler TcpIoHandler) {

	go s.HandleNetEvent(conn, ioHandler)
}

func (s *NetlibTcpServer) HandleNetEvent(conn net.Conn, ioHandler TcpIoHandler) {

	buf := make([]byte, 1024)
	for {
		len, err := conn.Read(buf)
		if err != nil {
			ioHandler.OnError(err)
			s.handlerMgr.OnClose(conn, s)
			break
		} else {
			ioHandler.OnRecv(buf, len)
		}
	}
}
