package handler

import (
	"context"
	log "github.com/micro/go-micro/v2/logger"
	"github.com/sword-demon/category/common"
	"github.com/sword-demon/category/domain/model"
	"github.com/sword-demon/category/domain/service"
	category "github.com/sword-demon/category/proto/category"
)

type Category struct {
	CategoryDataService service.ICategoryDataService
}

// CreateCategory 提供创建分类的服务
func (c *Category) CreateCategory(ctx context.Context, request *category.CategoryRequest,
	response *category.CreateCategoryResponse) error {
	categoryInfo := &model.Category{}
	// 赋值
	err := common.SwapTo(request, categoryInfo)
	if err != nil {
		return err
	}
	categoryId, err := c.CategoryDataService.AddCategory(categoryInfo)
	if err != nil {
		return err
	}
	response.Message = "分类添加成功"
	response.CategoryId = categoryId
	return nil
}

// UpdateCategory 提供更新分类的服务
func (c *Category) UpdateCategory(ctx context.Context, request *category.CategoryRequest,
	response *category.UpdateCategoryResponse) error {
	categoryInfo := &model.Category{}
	err := common.SwapTo(request, categoryInfo)
	if err != nil {
		return err
	}
	err = c.CategoryDataService.UpdateCategory(categoryInfo)
	response.Message = "分类更新成功"
	return nil
}

// DeleteCategory 提供删除分类的服务
func (c *Category) DeleteCategory(ctx context.Context, request *category.DeleteCategoryRequest,
	response *category.DeleteCategoryResponse) error {
	err := c.CategoryDataService.DeleteCategory(request.CategoryId)
	if err != nil {
		return err
	}
	response.Message = "分类删除成功"
	return nil
}

// FindCategoryByName 提供根据分类名称查找分类的服务
func (c *Category) FindCategoryByName(ctx context.Context, request *category.FindByNameRequest,
	response *category.CategoryResponse) error {
	categoryInfo, err := c.CategoryDataService.FindCategoryByName(request.CategoryName)
	if err != nil {
		return err
	}
	// 反向映射
	return common.SwapTo(categoryInfo, response)
}

// FindCategoryByID 提供根据分类ID查找分类的服务
func (c *Category) FindCategoryByID(ctx context.Context, request *category.FindByIdRequest,
	response *category.CategoryRequest) error {
	categoryInfo, err := c.CategoryDataService.FindCategoryByID(request.CategoryId)
	if err != nil {
		return err
	}
	// 反向映射
	return common.SwapTo(categoryInfo, response)
}

// FindCategoryByLevel 根据层级查找所有的分类服务
func (c *Category) FindCategoryByLevel(ctx context.Context, request *category.FindByLevelRequest,
	response *category.FindAllResponse) error {
	categorySlice, err := c.CategoryDataService.FindCategoryByLevel(request.Level)
	if err != nil {
		return err
	}
	categoryToResponse(categorySlice, response)
	return nil
}

// FindCategoryByParent 根据父分类查找所有的子分类
func (c *Category) FindCategoryByParent(ctx context.Context, request *category.FindByParentRequest,
	response *category.FindAllResponse) error {
	categorySlice, err := c.CategoryDataService.FindCategoryByParent(request.ParentId)
	if err != nil {
		return err
	}
	categoryToResponse(categorySlice, response)
	return nil
}

// FindAllCategory 查找所有的分类
func (c *Category) FindAllCategory(ctx context.Context, request *category.FindAllRequest,
	response *category.FindAllResponse) error {
	categorySlice, err := c.CategoryDataService.FindAllCategory()
	if err != nil {
		return err
	}
	categoryToResponse(categorySlice, response)
	return nil
}

func categoryToResponse(categorySlice []model.Category, response *category.FindAllResponse) {
	for _, cg := range categorySlice {
		cr := &category.CategoryResponse{}
		err := common.SwapTo(cg, cr)
		if err != nil {
			log.Error(err)
			break
		}
		response.Category = append(response.Category, cr)
	}
}
