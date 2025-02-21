package roles

import (
	"github.com/google/uuid"
	"github.com/zhangxianchengvip/go-ddd-dapr-k8s-microservices/pkg/ddd"
	"github.com/zhangxianchengvip/go-ddd-dapr-k8s-microservices/pkg/tools"
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

	if tools.StringIsEmplyOrWhiteSpace(name) {
		return nil, ErrNameEmpty
	}

	if tools.StringIsEmplyOrWhiteSpace(code) {
		return nil, ErrCodeEmpty
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
		return ErrNameEmpty
	}

	r.Name = name
	return nil
}
