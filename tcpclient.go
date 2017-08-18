package netlib

import "net"

type TcpClient struct {
	conn       net.Conn
	disconnect bool
	serverAddr string
	reader     TcpClientReader
}

func (t *TcpClient) Connect(addr string, reader TcpClientReader) bool {

	t.serverAddr = addr
	t.disconnect = true
	t.reader = reader
	var err error = nil
	t.conn, err = net.Dial("tcp", addr)
	if err != nil {
		return false
	}

	t.disconnect = false

	go t.ReadFun()

	return true
}

func (t *TcpClient) Close() {
	t.disconnect = true
	t.conn.Close()
}

func (t *TcpClient) Reconnect() bool {
	var err error
	t.conn, err = net.Dial("tcp", t.serverAddr)
	if err != nil {
		return false
	}

	t.disconnect = false

	go t.ReadFun()

	return true
}

func (t *TcpClient) IsDisconnect() bool {
	return t.disconnect
}

func (t *TcpClient) ReadFun() {
	var readBuf []byte = make([]byte, 1024)

	for {
		n, err := t.conn.Read(readBuf)
		if err != nil {
			t.disconnect = true
			t.reader.OnError(err)
			break
		}

		t.reader.OnRecv(readBuf, n)
	}
}

func (t *TcpClient) SendData(b []byte) (int, error) {

	return t.conn.Write(b)
}
