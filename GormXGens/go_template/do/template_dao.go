package do

import (
	"common-toolkits-v1/GormXGens/config"
	"common-toolkits-v1/GormXGens/go_template/entity"
	"gorm.io/gorm"
)

type userDo struct {
	db        gorm.DB
	tableName string
}

func newUserDao(db *gorm.DB) *userDo {
	model := entity.User{}
	tableName := model.TableName()
	tableDB := db.Table(tableName)
	return &userDo{
		db:        *tableDB,
		tableName: tableName,
	}
}

// GetTableName table名称
func (u userDo) GetTableName() string {
	return u.tableName
}

func (u userDo) GetList(opts ...config.DBFunc) ([]*entity.User, error) {
	db := u.db
	for _, f := range opts {
		f(&db)
	}
	var models []*entity.User
	db.Find(&models)
	return models, nil
}

// ....需要增加模板函数
