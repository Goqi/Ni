/*
0 error(s),0 warning(s)
Team:0e0w Security Team
Author:0e0wTeam[at]gmail.com
Datetime:2022/11/22 7:14
*/

package main

import (
	"Ernuclei/cmd"
	"fmt"
	"time"
)

func main() {
	StartTime := time.Now()
	cmd.Nuclei()
	EndTime := time.Now()
	Time := EndTime.Sub(StartTime)
	fmt.Printf("程序运行完成！运行时间:%s\n", Time)
}
