package main
// by :Try
import (
	"fmt"
	"netport/root"
	"os"
	"runtime"
)

func main() {
	sysType := runtime.GOOS
	if sysType =="windows"{
		root.WinRun()   //windows执行这个
	}else if sysType=="linux" {
		root.LinuxRun()		//Linux执行这个
	}else {
		fmt.Println("[Info] 系统识别错误,请联系作者")
		os.Exit(1)
	}
}
