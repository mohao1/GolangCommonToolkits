package do

import (
	"common-toolkits-v1/GormXGens/config"
	"common-toolkits-v1/GormXGens/go_template/entity"
	"context"
	"database/sql"
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

// GetList 获取批量数据
func (u userDo) GetList(ctx context.Context, opts ...config.DBFunc) ([]*entity.User, error) {
	db := u.db.WithContext(ctx)
	for _, f := range opts {
		err := f(db)
		if err != nil {
			return nil, err
		}
	}
	var models []*entity.User
	db.Find(&models)
	return models, nil
}

// GetRow 获取Row的数据
func (u userDo) GetRow(ctx context.Context, opts ...config.DBFunc) (*sql.Row, error) {
	db := u.db.WithContext(ctx)
	for _, f := range opts {
		err := f(db)
		if err != nil {
			return nil, err
		}
	}
	row := db.Row()
	return row, nil
}

// GetRows 获取Rows的数据
func (u userDo) GetRows(ctx context.Context, opts ...config.DBFunc) (*sql.Rows, error) {
	db := u.db.WithContext(ctx)
	for _, f := range opts {
		err := f(db)
		if err != nil {
			return nil, err
		}
	}
	return db.Rows()
}

// =====主键 Start=====

// GetUserByID 主键查询
func (u userDo) GetUserByID(ctx context.Context, ID int32, opts ...config.DBFunc) (*entity.User, error) {
	db := u.db.WithContext(ctx)
	for _, f := range opts {
		err := f(db)
		if err != nil {
			return nil, err
		}
	}
	db.Where("id = ?", ID)
	var models *entity.User
	db.First(&models)
	return models, nil
}

// DeleteByID 主键删除数据
func (u userDo) DeleteByID(ctx context.Context, ID int32, opts ...config.DBFunc) error {
	db := u.db.WithContext(ctx)
	for _, f := range opts {
		err := f(db)
		if err != nil {
			return err
		}
	}
	tx := db.Delete(&entity.User{
		ID: ID,
	})
	if tx.Error != nil {
		return tx.Error
	}
	return nil
}

// UpdateByID 主键修改数据
func (u userDo) UpdateByID(ctx context.Context, ID int32, model *any, opts ...config.DBFunc) error {
	db := u.db.WithContext(ctx)
	for _, f := range opts {
		err := f(db)
		if err != nil {
			return err
		}
	}
	db.Where("id = ?", ID)
	result := db.Save(model)
	if err := result.Error; err != nil {
		return err
	}
	return nil
}

// =====主键 End=====

// =====索引 Start=====

// GetUserByName 索引查询
func (u userDo) GetUserByName(ctx context.Context, Name string, opts ...config.DBFunc) (*entity.User, error) {
	db := u.db.WithContext(ctx)
	for _, f := range opts {
		err := f(db)
		if err != nil {
			return nil, err
		}
	}
	db.Where("name = ?", Name)
	var models *entity.User
	db.First(&models)
	return models, nil
}

// DeleteByName 索引删除数据
func (u userDo) DeleteByName(ctx context.Context, Name string, opts ...config.DBFunc) error {
	db := u.db.WithContext(ctx)
	for _, f := range opts {
		err := f(db)
		if err != nil {
			return err
		}
	}
	tx := db.Delete(&entity.User{
		Name: Name,
	})
	if tx.Error != nil {
		return tx.Error
	}
	return nil
}

// UpdateByName 主键修改数据
func (u userDo) UpdateByName(ctx context.Context, Name string, model *any, opts ...config.DBFunc) error {
	db := u.db.WithContext(ctx)
	for _, f := range opts {
		err := f(db)
		if err != nil {
			return err
		}
	}
	db.Where("name = ?", Name)
	result := db.Save(model)
	if err := result.Error; err != nil {
		return err
	}
	return nil
}

// GetListByName 获取批量数据
func (u userDo) GetListByName(ctx context.Context, Name string, opts ...config.DBFunc) ([]*entity.User, error) {
	db := u.db.WithContext(ctx)
	for _, f := range opts {
		err := f(db)
		if err != nil {
			return nil, err
		}
	}
	db.Where("name = ?", Name)
	var models []*entity.User
	db.Find(&models)
	return models, nil
}

// GetRowByName 获取Row的数据
func (u userDo) GetRowByName(ctx context.Context, Name string, opts ...config.DBFunc) (*sql.Row, error) {
	db := u.db.WithContext(ctx)
	for _, f := range opts {
		err := f(db)
		if err != nil {
			return nil, err
		}
	}
	db.Where("name = ?", Name)
	row := db.Row()
	return row, nil
}

// GetRowsByName 获取Rows的数据
func (u userDo) GetRowsByName(ctx context.Context, Name string, opts ...config.DBFunc) (*sql.Rows, error) {
	db := u.db.WithContext(ctx)
	for _, f := range opts {
		err := f(db)
		if err != nil {
			return nil, err
		}
	}
	db.Where("name = ?", Name)
	return db.Rows()
}

// GetListCountByName Count数量统计
func (u userDo) GetListCountByName(ctx context.Context, Name string, opts ...config.DBFunc) (int64, error) {
	db := u.db.WithContext(ctx)
	for _, f := range opts {
		err := f(db)
		if err != nil {
			return 0, err
		}
	}
	db.Where("name = ?", Name)
	var count int64
	db.Count(&count)
	if db.Error != nil {
		return 0, db.Error
	}
	return count, nil
}

// =====索引 End=====

// Create 创建数据
func (u userDo) Create(ctx context.Context, model *any) error {
	db := u.db.WithContext(ctx)
	result := db.Create(model)
	if err := result.Error; err != nil {
		return err
	}
	return nil
}

// Update 修改数据
func (u userDo) Update(ctx context.Context, model *any, opts ...config.DBFunc) error {
	db := u.db.WithContext(ctx)
	for _, f := range opts {
		err := f(db)
		if err != nil {
			return err
		}
	}
	result := db.Updates(model)
	if err := result.Error; err != nil {
		return err
	}
	return nil
}

// GetListCount Count数量统计
func (u userDo) GetListCount(ctx context.Context, opts ...config.DBFunc) (int64, error) {
	db := u.db.WithContext(ctx)
	for _, f := range opts {
		err := f(db)
		if err != nil {
			return 0, err
		}
	}
	var count int64
	db.Count(&count)
	if db.Error != nil {
		return 0, db.Error
	}
	return count, nil
}

// ....需要增加模板函数
