package utils

import "log"

//日志控制，根据开发环境和生产环境自己设置RUN_MODE
//修改此处的RUN_MODE控制自定义日志是否打印
//定义RUN_MODE==1，表示dev
//定义RUN_MODE==2，表示prod
//……

const (
	RUN_MODE int = 2
)

func Println(v ...interface{}) {
	if RUN_MODE == 1 {
		log.Println(v)
	}
}

func Print(v ...interface{}) {
	if RUN_MODE == 1 {
		log.Print(v)
	}
}

func Fatal(v ...interface{}) {
	if RUN_MODE == 1 {
		log.Fatal(v)
	}
}

//项目只添加上述几个日志的控制，其它函数可以自己定义
//func FuncName(params) {
//	if RUN_MODE == "dev" {
//		log.FuncName(params)
//	}
//}
