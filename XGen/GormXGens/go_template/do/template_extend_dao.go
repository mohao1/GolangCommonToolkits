package do

import (
	"gorm.io/gorm"
	"sync"
)

// Single instance storage （单例存储）
var do *UserDo
var one sync.Once

type UserDo struct {
	*userDo
}

func NewUserDo(db *gorm.DB) *UserDo {
	one.Do(func() {
		do = &UserDo{
			newUserDao(db),
		}
	})
	return do
}

// The extension method can be implemented here （扩展方法可以在这实现）
