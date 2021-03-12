package root

import (
	"log"
	"os"
	"strings"
)

func Add(listenport string, address string, connectport string, connectaddress string) { //添加转发规则
	cmdstr := Cmd("netsh interface portproxy show all")
	if strings.Contains(cmdstr, connectaddress+"    "+connectport) == true && strings.Contains(cmdstr, connectaddress+"    "+listenport) == true {
		log.Println("[Info] 监听ip:", connectaddress, " 端口:", listenport)
		log.Println("[Info] 转发ip:", connectaddress, " 转发端口:", connectport)
		log.Println("[Info] 该规则已经添加过了！")
		return
	} else {
		Cmd("netsh interface portproxy add v4tov4 listenport=" + listenport + " listenaddress=" + connectaddress + " connectport=" + connectport + " connectaddress=" + connectaddress)
		c1 := Cmd("netsh interface portproxy show all")
		if strings.Contains(c1, connectaddress+"    "+connectport) == true {
			log.Println("[Info] 监听ip:", address, " 端口:", listenport)
			log.Println("[Info] 转发ip:", connectaddress, " 转发端口:", connectport)
			log.Println("[Info] 添加成功！")
		} else {
			log.Println("[Info] 添加失败！")
			os.Exit(1)
		}
	}
}
