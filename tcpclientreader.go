package netlib

type TcpClientReader interface {
	OnRecv([]byte, int)
	OnError(error)
}
