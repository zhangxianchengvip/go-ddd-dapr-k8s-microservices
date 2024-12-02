package users

import "github.com/zhangxianchengvip/go-ddd-dapr-k8s-microservices/pkg/ddd"

type PasswordChangedEvent struct {
	ddd.DomainEvent
	UserID string
}
