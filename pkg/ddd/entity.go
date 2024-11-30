package ddd

import "github.com/google/uuid"

// IEntity 接口定义了实体的行为
type IEntity interface {
	Equal(other IEntity) bool
}

type TKey interface {
	int | int64 | string | uuid.UUID
}

// Entity 结构体实现了 IEntity 接口，可以用于具有唯一标识符的实体
type Entity[T TKey] struct {
	ID T
}

// Equal 方法用于比较两个实体是否相等
func (e Entity[TKey]) Equal(other IEntity) bool {
	// 检查 other 是否为 Entity 类型
	if otherEntity, ok := other.(Entity[TKey]); ok {
		// 使用 ID 比较实体
		return e.ID == otherEntity.ID
	}
	return false
}
