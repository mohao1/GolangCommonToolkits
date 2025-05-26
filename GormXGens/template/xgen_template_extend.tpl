// 扩展自query的实现
package {{.QueryPackage}}

import (
	"gorm.io/gorm"
	"sync"
)

// 单例存储
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

// 扩展方法可以在这实现