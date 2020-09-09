package main

import (
	"bytes"
	"fmt"
	"log"
	"os/exec"
	"regexp"
	"strings"
)

func main() {
	var stdout bytes.Buffer
	// log.SetFlags(log.Lshortfile | log.LstdFlags)
	log.SetOutput(&stdout)

	log.Println("111111111")
	fmt.Println("222222222")
	log.Println("333333")

	str := `
2020/05/30 17:13:06 do.go:18: There are 22 customers who need to upgrade!
2020/05/30 17:13:42 do.go:36: Releated tasks 5ZPerY1Ky3uHGITQ, 5ZPerY1KjQ5rFalX, 5ZPerY1K2tihedEd, 5ZPerY1KIC9k3nD4
2020/05/30 17:14:18 do.go:48: Task uuid[DgdrtjQHBSNoIXSR] name[机甲大师],template uuid[8uJU1VNk8dZIMJFj] name[机甲大师私有部署「高可用版」升级申请]
2020/05/30 17:14:56 do.go:48: Task uuid[DgdrtjQHjVaR4HYi] name[上海伊邦医药],template uuid[8uJU1VNkLndo2e5X] name[上海伊邦私有部署升级申请]
2020/05/30 17:15:34 do.go:48: Task uuid[DgdrtjQHSJNZnlqR] name[销售演示],template uuid[8uJU1VNkZjnFMQ24] name[销售演示私有部署「高可用版」升级申请]
2020/05/30 17:16:08 do.go:48: Task uuid[8uJU1VNkGBBuRk34] name[国家超级计算无锡中心],template uuid[8uJU1VNkILCVIbDM] name[国家超算无锡中心私有部署升级申请]
2020/05/30 17:16:43 do.go:48: Task uuid[DgdrtjQHDGkmuPlN] name[文轩在线],template uuid[8uJU1VNke4VcTa4J] name[文轩在线私有部署升级申请]
2020/05/30 17:17:18 do.go:48: Task uuid[DgdrtjQHflTG2O6E] name[广州猫玩网络科技有限公司],template uuid[8uJU1VNkRofSgIBl] name[猫玩私有部署升级申请]
2020/05/30 17:17:57 do.go:48: Task uuid[8uJU1VNkrFU3nlNp] name[创梦天地（乐逗游戏）],template uuid[8uJU1VNkPS42shPC] name[创梦天地（乐逗）私有部署升级申请]
2020/05/30 17:18:34 do.go:48: Task uuid[8uJU1VNkKcPgAtWM] name[上海万物工场],template uuid[8uJU1VNkUZ6a8EBP] name[万物工场私有部署升级申请]
2020/05/30 17:19:19 do.go:48: Task uuid[8uJU1VNkU86vf8Bm] name[上海荣数],template uuid[8uJU1VNk4Y7fWiOA] name[上海荣数私有部署升级申请]
2020/05/30 17:19:57 do.go:48: Task uuid[8uJU1VNkj2X4VTrB] name[医惠科技],template uuid[8uJU1VNkGsYj1RdO] name[医惠科技私有部署升级申请]
2020/05/30 17:20:35 do.go:48: Task uuid[8uJU1VNkZwPkdNxQ] name[日电（中国）有限公司],template uuid[8uJU1VNkUmFx6Enk] name[日电中国（NEC）私有部署升级申请]
2020/05/30 17:21:13 do.go:48: Task uuid[8uJU1VNkJxsFhZWl] name[热云科技],template uuid[8uJU1VNk55px3hoD] name[热云科技私有部署升级申请]
2020/05/30 17:21:51 do.go:48: Task uuid[DgdrtjQHWcelXZDW] name[招商基金],template uuid[8uJU1VNkWmEdC7Vk] name[招商基金私有部署升级申请]
2020/05/30 17:22:26 do.go:48: Task uuid[DgdrtjQHq9phcous] name[北大英华],template uuid[8uJU1VNkQQE9uSEG] name[北大英华私有部署升级申请]
2020/05/30 17:23:04 do.go:48: Task uuid[8uJU1VNk5vZvCKSp] name[鸿合科技],template uuid[8uJU1VNkfZH9n6vE] name[鸿合科技私有部署升级申请]
2020/05/30 17:23:42 do.go:48: Task uuid[8uJU1VNk1g4uO8NV] name[中福彩科技],template uuid[8uJU1VNkq7agQEY5] name[中福彩科技私有部署升级申请]
2020/05/30 17:24:16 do.go:48: Task uuid[DgdrtjQH15udoCNA] name[玩蟹科技],template uuid[8uJU1VNk8QJmeaw1] name[玩蟹科技私有部署升级申请]
2020/05/30 17:24:54 do.go:48: Task uuid[8uJU1VNkBKiRtjvp] name[上证信息],template uuid[8uJU1VNkuxD6U9rj] name[上证信息私有部署升级申请]
2020/05/30 17:25:33 do.go:48: Task uuid[8uJU1VNkS5GHTtrD] name[盐城一之来],template uuid[8uJU1VNks9rbAgET] name[盐城一之来私有部署升级申请]
2020/05/30 17:26:08 do.go:48: Task uuid[DgdrtjQHnSwcWEQq] name[掌游],template uuid[8uJU1VNktQbm3FqZ] name[掌游私有部署升级申请]
2020/05/30 17:26:43 do.go:48: Task uuid[DgdrtjQHWUbeVGkY] name[华发集团],template uuid[8uJU1VNkInt6HVo3] name[华发集团私有部署升级申请]
2020/05/30 17:27:22 do.go:48: Task uuid[DgdrtjQHckCHTeqd] name[云象安全],template uuid[8uJU1VNkPy8qaRiy] name[云象安全私有部署升级申请]
	`
	re3, _ := regexp.Compile("\\d\\d\\d\\d/\\d\\d/\\d\\d \\d\\d:\\d\\d:\\d\\d")
	rep := re3.ReplaceAllStringFunc(str, repl)
	fmt.Println(rep)

	strings.Replace(stdout.String(), "", "", -1)

	// RunCommand(fmt.Sprintf(`curl 'https://qyapi.weixin.qq.com/cgi-bin/webhook/send?key=8b076bf1-edbd-4dc8-af27-e1ee0f92e096' \
	// -H 'Content-Type: application/json' \
	// -d '
	// {
	// 	 "msgtype": "text",
	// 	 "text": {
	// 		 "content": "%s"
	// 	 }
	// }'`, stdout.String()))

	log.Println("over")
	fmt.Println(stdout.String())

}

func repl(a string) string {
	return ""
}
func RunCommand(command string) (string, error) {
	cmd := exec.Command("/bin/bash", "-c", command)

	stdout, err := cmd.Output()
	if err != nil {
		return string(stdout), err
	}
	return string(stdout), nil
}
