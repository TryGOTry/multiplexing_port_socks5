package root

import "fmt"

func Del() { //删除转发规则
	Cmd("netsh interface portproxy reset")
	fmt.Println("[Info] 转发规则清空成功！")
}