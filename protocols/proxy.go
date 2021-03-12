package protocols


import (
	"net"
)

func proxy(from net.Conn, to net.Conn, closed chan bool) {
	data := make([]byte, 4096) // 4KiB buffer

	for {
		n, err := from.Read(data)
		if err != nil {
			closed <- true
			return
		}
		_, err = to.Write(data[:n])
		if err != nil {
			closed <- true
			return
		}
	}
}
