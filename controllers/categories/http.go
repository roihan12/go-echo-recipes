package categories

import (
	"echo-recipe/businesses/categories"
	controller "echo-recipe/controllers"
	"echo-recipe/controllers/categories/request"
	"echo-recipe/controllers/categories/response"
	"net/http"

	"github.com/labstack/echo/v4"
)

type CategoryController struct {
	categoryUseCase categories.Usecase
}

func NewCategoryController(categoryUC categories.Usecase) *CategoryController {
	return &CategoryController{
		categoryUseCase: categoryUC,
	}
}

func (ctrl *CategoryController) GetAllCategories(c echo.Context) error {
	categoriesData := ctrl.categoryUseCase.GetAll()

	categories := []response.Category{}

	for _, category := range categoriesData {
		categories = append(categories, response.FromDomain(category))
	}

	return controller.NewResponse(c, http.StatusOK, "success", "all categories", categories)
}

func (ctrl *CategoryController) GetByID(c echo.Context) error {

	var id string = c.Param("id")
	category := ctrl.categoryUseCase.GetByID(id)

	if category.ID == 0 {
		return controller.NewResponse(c, http.StatusNotFound, "failed", "category not found", "")
	}
	return controller.NewResponse(c, http.StatusOK, "success", "category found", response.FromDomain(category))
}

func (ctrl *CategoryController) CreateCategory(c echo.Context) error {
	input := request.Category{}

	if err := c.Bind(&input); err != nil {
		return controller.NewResponse(c, http.StatusBadRequest, "failed", "validation failed", "")
	}

	if err := c.Validate(&input); err != nil {
		return err
	}

	category := ctrl.categoryUseCase.Create(input.ToDomain())

	return controller.NewResponse(c, http.StatusCreated, "success", "category created", response.FromDomain(category))
}

func (ctrl *CategoryController) UpdateCategory(c echo.Context) error {
	input := request.Category{}

	if err := c.Bind(&input); err != nil {
		return controller.NewResponse(c, http.StatusBadRequest, "failed", "validation failed", "")
	}

	var id string = c.Param("id")

	if err := c.Validate(&input); err != nil {
		return err
	}

	category := ctrl.categoryUseCase.Update(id, input.ToDomain())

	return controller.NewResponse(c, http.StatusOK, "success", "category updated", response.FromDomain(category))
}

func (ctrl *CategoryController) DeleteCategory(c echo.Context) error {
	var id string = c.Param("id")

	isDeleted := ctrl.categoryUseCase.Delete(id)

	if !isDeleted {
		return controller.NewResponse(c, http.StatusNotFound, "failed", "category not found", "")
	}

	return controller.NewResponse(c, http.StatusOK, "success", "category deleted", "")
}
