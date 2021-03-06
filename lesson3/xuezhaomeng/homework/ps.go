package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strconv"
)

func main() {
	// linux: /proc, 其他系统自己创建个模拟目录
	path := os.Args[1]
	f, _ := os.Open(path)

	infos, err := f.Readdir(-2)
	if err != nil {
		log.Fatal(err)
	}

	for _, info := range infos {
		// 判断是否是pid
		_, err := strconv.Atoi(info.Name())
		if err != nil {
			continue
		}

		// 判断是否是pid目录
		if info.IsDir() {
			// 拼接类似/proc/1877/cmdline的路径
			filename := path + "/" + info.Name() + "/cmdline"
			// 作业: 读取filename里的命令, 打印 pid, cmd
			buf, err := ioutil.ReadFile(filename)
			if err != nil {
				fmt.Println(err)
				return

			}
			//fmt.Println(string(buf))
			//fmt.Println(filename)
			cmd := string(buf)
			fmt.Printf("%v\t%v \n", info.Name(), cmd)
		}
	}
}
