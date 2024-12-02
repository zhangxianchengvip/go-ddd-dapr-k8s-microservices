package roles

import (
	"errors"

	"github.com/google/uuid"
	"github.com/zhangxianchengvip/go-ddd-dapr-k8s-microservices/pkg/ddd"
)

type Role struct {
	ddd.AggregateRoot[uuid.UUID]            // 聚合根
	Code                         string     // 编码
	Name                         string     // 名称
	Order                        int        // 排序
	ParentId                     *uuid.UUID // 父级ID
	Description                  *string    // 描述
}

func NewRole(id uuid.UUID, name, code string, desc *string, order int, parentId *uuid.UUID) (*Role, error) {

	if name == "" {
		return nil, errors.New("name is required")
	}

	if code == "" {
		return nil, errors.New("code is required")
	}

	return &Role{
		AggregateRoot: ddd.NewAggregateRoot(id),
		Name:          name,
		Description:   desc,
		Code:          code,
		Order:         order,
		ParentId:      parentId,
	}, nil
}

func (r *Role) UpdateName(name string) error {
	if name == "" {
		return errors.New("name is required")
	}

	r.Name = name
	return nil
}
