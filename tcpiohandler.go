package netlib

type TcpIoHandler interface {
	OnRecv(b []byte, len int)
	OnError(err error)
}
