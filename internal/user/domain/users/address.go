package users

import "github.com/zhangxianchengvip/go-ddd-dapr-k8s-microservices/pkg/ddd"

type Address struct {
	ddd.ValueObject        //值对象
	Street          string //街道
	City            string //城市
}
