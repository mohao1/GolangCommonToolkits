// Package apply_link_doctor Implementation of Extended Self Query （扩展自query的实现）
package apply_link_doctor

import (
	"gorm.io/gorm"
	"sync"
)

// Single instance storage （单例存储）
var query *ApplyLinkDoctorQuery
var one sync.Once

type ApplyLinkDoctorQuery struct {
	*applyLinkDoctorQuery
}

func NewApplyLinkDoctorQuery(db *gorm.DB) *ApplyLinkDoctorQuery {
	one.Do(func() {
		query = &ApplyLinkDoctorQuery{
			newQuery(db),
		}
	})
	return query
}

// The extension method can be implemented here （扩展方法可以在这实现）
