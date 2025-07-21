package config

import (
	"fmt"
	"os"
	"strings"
)

var (
	Projects           []string
	Loglevel           string
	HarborUrl          string
	HarborUser         string
	HarborUserPassword string
	ClearFlag          bool
	DelTagPrefix       []string
	RepoNamePrefix     []string
)

func init() {
	ps := os.Getenv("harborClear_Projects")
	for _, v := range strings.Split(ps, ",") {
		Projects = append(Projects, v)
	}
	if Projects[0] == "" {
		fmt.Println("请设置 harborClear_Projects 环境变量")
		os.Exit(-1)
	}

	//monthStr := os.Getenv("harborClear_KeepMonth")
	//if monthStr == "" {
	//	fmt.Println("请设置 harborClear_KeepMonth 环境变量")
	//	os.Exit(-1)
	//}
	//MonthInt, err := strconv.Atoi(monthStr)
	//if err != nil {
	//	fmt.Println("harborClear_KeepMonth 必须为数字字符")
	//	os.Exit(-1)
	//}
	//Month = MonthInt

	level := os.Getenv("harborClear_Loglevel")
	if level == "" {
		fmt.Println("请设置 harborClear_Loglevel 环境变量")
		os.Exit(-1)
	}
	Loglevel = level

	Url := os.Getenv("harborClear_HarborUrl")
	if Url == "" {
		fmt.Println("请设置 harborClear_HarborUrl 环境变量")
		os.Exit(-1)
	}
	HarborUrl = Url

	user := os.Getenv("harborClear_HarborUser")
	if user == "" {
		fmt.Println("请设置 harborClear_HarborUser 环境变量")
		os.Exit(-1)
	}
	HarborUser = user

	pwd := os.Getenv("harborClear_UserPassword")
	if pwd == "" {
		fmt.Println("请设置 harborClear_UserPassword 环境变量")
		os.Exit(-1)
	}
	HarborUserPassword = pwd

	flag := os.Getenv("harborClear_ClearFlag")
	if flag == "" {
		fmt.Println("请设置 harborClear_ClearFlag 环境变量")
		os.Exit(-1)
	}
	if strings.ToLower(flag) == "true" {
		ClearFlag = true
	} else {
		ClearFlag = false
	}

	//ks := os.Getenv("harborClear_KeepSave")
	//if ks == "" {
	//	fmt.Println("请设置 harborClear_KeepSave 环境变量")
	//	os.Exit(-1)
	//}
	//if strings.ToLower(ks) == "true" {
	//	KeepSave = true
	//} else {
	//	KeepSave = false
	//}
	dl := os.Getenv("harborClear_DelTagPrefix")
	if dl == "" {
		fmt.Println("请设置 harborClear_DelTagPrefix 环境变量")
		os.Exit(-1)
	}
	DelTagPrefix = strings.Split(dl, ",")

	ns := os.Getenv("harborClear_RepoNamePrefix")
	if ns == "" {
		fmt.Println("请设置 harborClear_RepoNamePrefix 环境变量")
		os.Exit(-1)
	}
	RepoNamePrefix = strings.Split(ns, ",")

}
