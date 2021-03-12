package root

import (
	"fmt"
	"golang.org/x/text/encoding/simplifiedchinese"
	"os/exec"
)
type Charset string
const (
	UTF8    = Charset("UTF-8")
	GB18030 = Charset("GB18030")
)

func ConvertByte2String(byte []byte, charset Charset) string {
	var str string
	switch charset {
	case GB18030:
		var decodeBytes, _ = simplifiedchinese.GB18030.NewDecoder().Bytes(byte)
		str = string(decodeBytes)
	case UTF8:
		fallthrough
	default:
		str = string(byte)
	}
	return str
}
func Cmd(command string) (out string) {   //windows下执行
	cmd := exec.Command("cmd.exe", "/c", command)
	o, err := cmd.CombinedOutput()
	if err != nil {
		out = fmt.Sprintf("[info] shell run error: \n%s\n", err)
	} else {
		cmdRe := ConvertByte2String(o, "GB18030")
		out = fmt.Sprintf(cmdRe)
		//fmt.Println(cmdRe)
	}
	return
}
func LinuxCmd(command string) (out string) {   //linux下执行
	cmd := exec.Command("/bin/bash", "-c", command)
	o, err := cmd.CombinedOutput()
	if err != nil {
		out = fmt.Sprintf("[info] shell run error: \n%s\n", err)
	} else {
		cmdRe := ConvertByte2String(o, "GB18030")
		out = fmt.Sprintf(cmdRe)
		//fmt.Println(cmdRe)
	}
	return
}