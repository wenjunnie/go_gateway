package main

import (
	"fmt"
	"github.com/e421083458/golang_common/lib"
	"go_gateway/controller"
	"go_gateway/router"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	lib.InitModule("./conf/dev/", []string{"base", "mysql", "redis"})
	defer lib.Destroy()
	err := controller.InitClient()
	if err != nil {
		fmt.Printf("connect failed, err:%v\n", err)
		return
	}
	router.HttpServerRun()

	quit := make(chan os.Signal)
	signal.Notify(quit, syscall.SIGKILL, syscall.SIGQUIT, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	router.HttpServerStop()
}
