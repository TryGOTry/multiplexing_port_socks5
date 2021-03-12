package root

import (
	"flag"
	"fmt"
	"github.com/rs/zerolog"
	"log"
	"netport/protocols"
	"netport/socket5"
	"os"
)

func WinRun() {
	logger := zerolog.New(os.Stdout).With().Timestamp().Logger()
	p := make([]*protocols.Protocol, 0, 7)
	socketport := flag.String("socks", "", "socks5开启端口")
	cmdtype := flag.String("t", "", "执行类型(add/del)")
	listenport := flag.String("p", "", "本地监听的端口")
	address := flag.String("h", "", "本地监听的ip地址")
	connectport := flag.String("cp", "", "需要转发到的端口")
	connectaddress := flag.String("ch", "", "需要转发到的地址")
	flag.Parse()
	if *cmdtype == "add" {
		if *listenport == "" {
			log.Println("[Info] 请填写需要转发的端口！")
		} else if *socketport ==""{
			log.Println("[Info] 请填写socks5开启的端口！")
		}else if *address == "" {
			log.Println("[Info] 请填写需要转发的ip地址！")
		} else if *connectport == "" {
			log.Println("[Info] 请填写需要转发的目标端口！")
		} else if *connectaddress == "" {
			log.Println("[Info] 请填写需要转发的目标ip地址！")
		} else {
			p = append(p, protocols.NewHTTPProtocol(*address+":"+*listenport))
			p = append(p, protocols.NewSOCKS5Protocol(*connectaddress+":"+*socketport))
			Add(*listenport, *address, *connectport, *connectaddress)
			go socket5.CreateForwardSocks(*socketport)
			protocols.RunServer(*connectaddress+":"+*connectport, p, logger)
		}
	} else if *cmdtype == "del" {
			Del()
	} else {
		fmt.Println("[Info] Socks5代理以及端口复用小工具  by:T00ls.Net   Try")
		fmt.Println("[Info] Add_Example:open_port_socks5.exe -h 127.0.0.1 -p 80 -t add -ch 10.10.10.241(内网ipd) -cp 899 -socks 789")
		fmt.Println("[Info] 解释:本地开启一个socks5代理,端口789.将80端口的流量转发到899端口,然后在899端口进行流量分流,http请求就走127.0.0.1:80端口,socks5请求就走10.10.10.241的789端口.")
		fmt.Println("[Info] Del_Example:open_port_socks5.exe -t del")
		fmt.Println("[Info] 解释:删除所有转发规则.")
		fmt.Println("[Info] 查看帮助:open_port_socks5.exe --help")
		//flag.Usage()
		os.Exit(1)
	}
}
