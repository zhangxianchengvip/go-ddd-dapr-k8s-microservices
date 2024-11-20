package main

import (
	"fmt"

	"github.com/zhangxianchengvip/go-ddd-dapr-k8s-microservices/user/cmd/user/webapp"
)

func main() {
	fmt.Println("Hello User Service!")
	webapp.Run()
}
