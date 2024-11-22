package roles

import (
	"errors"

	"github.com/google/uuid"
)

type Role struct {
	ID          uuid.UUID  `gorm:"type:uuid;primary_key" json:"id"` // 主键
	Code        string     `gorm:"type:varchar(255)" json:"code"`   // 编码
	Name        string     `gorm:"type:varchar(255)" json:"name"`   // 名称
	Order       int        `gorm:"type:int" json:"order"`           // 排序
	ParentId    *uuid.UUID `gorm:"type:uuid" json:"parent_id"`      // 父级ID
	Description *string    `gorm:"type:text" json:"description"`    // 描述
}

func NewRole(id uuid.UUID, name, code string, desc *string, order int, parentId *uuid.UUID) (*Role, error) {

	if name == "" {
		return nil, errors.New("name is required")
	}

	if code == "" {
		return nil, errors.New("code is required")
	}

	return &Role{
		ID:          id,
		Name:        name,
		Description: desc,
		Code:        code,
		Order:       order,
		ParentId:    parentId,
	}, nil
}

func (r *Role) UpdateName(name string) error {
	if name == "" {
		return errors.New("name is required")
	}

	r.Name = name
	return nil
}
