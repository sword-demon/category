package repository

import (
	"github.com/sword-demon/category/domain/model"
	"github.com/jinzhu/gorm"
)

type ICategoryRepository interface {
	InitTable() error
	FindCategoryByID(int64) (*model.Category, error)
	CreateCategory(*model.Category) (int64, error)
	DeleteCategoryByID(int64) error
	UpdateCategory(*model.Category) error
	FindAll() ([]model.Category, error)
}

// NewCategoryRepository 创建categoryRepository
func NewCategoryRepository(db *gorm.DB) ICategoryRepository {
	return &CategoryRepository{mysqlDb: db}
}

type CategoryRepository struct {
	mysqlDb *gorm.DB
}

// InitTable 初始化表
func (u *CategoryRepository) InitTable() error {
	return u.mysqlDb.CreateTable(&model.Category{}).Error
}

// FindCategoryByID 根据ID查找Category信息
func (u *CategoryRepository) FindCategoryByID(categoryID int64) (category *model.Category, err error) {
	category = &model.Category{}
	return category, u.mysqlDb.First(category, categoryID).Error
}

// CreateCategory 创建Category信息
func (u *CategoryRepository) CreateCategory(category *model.Category) (int64, error) {
	return category.ID, u.mysqlDb.Create(category).Error
}

// DeleteCategoryByID 根据ID删除Category信息
func (u *CategoryRepository) DeleteCategoryByID(categoryID int64) error {
	return u.mysqlDb.Where("id = ?", categoryID).Delete(&model.Category{}).Error
}

// UpdateCategory 更新Category信息
func (u *CategoryRepository) UpdateCategory(category *model.Category) error {
	return u.mysqlDb.Model(category).Update(category).Error
}

// FindAll 获取结果集
func (u *CategoryRepository) FindAll() (categoryAll []model.Category, err error) {
	return categoryAll, u.mysqlDb.Find(&categoryAll).Error
}
