package main

import (
	"fmt"

	"github.com/zhangxianchengvip/go-ddd-dapr-k8s-microservices/cmd/user/app"
)

func main() {
	fmt.Println("Hello User Service!")
	app.Run()
}
