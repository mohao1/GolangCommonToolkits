// Package users Implementation of Extended Self Query （扩展自query的实现）
package users

import (
	"gorm.io/gorm"
	"sync"
)

// Single instance storage （单例存储）
var query *UserQuery
var one sync.Once

type UserQuery struct {
	*userQuery
}

func NewUserQuery(db *gorm.DB) *UserQuery {
	one.Do(func() {
		query = &UserQuery{
			newQuery(db),
		}
	})
	return query
}

// The extension method can be implemented here （扩展方法可以在这实现）
