package socket5

import (
	socks5 "github.com/armon/go-socks5"
	"github.com/hashicorp/yamux"
	"log"
)

var session *yamux.Session

func CreateForwardSocks(address string) error {
	server, err := socks5.New(&socks5.Config{})
	if err != nil {
		return err
	}
	log.Println("[Info] socks5开启成功！端口:",address)
	if err := server.ListenAndServe("tcp", "0.0.0.0:"+address); err != nil {
		return err
	}
	return nil
}
