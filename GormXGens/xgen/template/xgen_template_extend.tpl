// Package query Implementation of Extended Self Query （扩展自query的实现）
package {{.QueryPackage}}

import (
	"gorm.io/gorm"
	"sync"
)

// Single instance storage （单例存储）
var query *{{.ExtendName}}
var one sync.Once

type {{.ExtendName}} struct {
	*{{.ParentName}}
}

func New{{.ExtendName}}(db *gorm.DB) *{{.ExtendName}} {
	one.Do(func() {
		query = &{{.ExtendName}}{
			newQuery(db),
		}
	})
	return query
}

// The extension method can be implemented here （扩展方法可以在这实现）