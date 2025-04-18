package ctx

import (
	"gorm.io/gorm"
)

type AppDbContext struct {
	DB *gorm.DB
}

func NewAppDbContext(db *gorm.DB) *AppDbContext {
	return &AppDbContext{DB: db}
}

func (ctx *AppDbContext) Create(model interface{}) error {
	return ctx.DB.Create(model).Error
}

func (ctx *AppDbContext) Find(dest interface{}, conds ...interface{}) error {
	return ctx.DB.Find(dest, conds...).Error
}

func (ctx *AppDbContext) First(dest interface{}, conds ...interface{}) error {
	return ctx.DB.First(dest, conds...).Error
}

func (ctx *AppDbContext) Update(model interface{}, values interface{}) error {
	err := ctx.DB.
		Omit("created_by", "created_at").
		Model(model).
		Updates(values).Error
	if err != nil {
		return err
	}
	return ctx.DB.Find(model, model).Error
}

func (ctx *AppDbContext) Delete(model interface{}) error {
	return ctx.DB.Delete(model).Error
}

func (ctx *AppDbContext) Raw() *gorm.DB {
	return ctx.DB
}
