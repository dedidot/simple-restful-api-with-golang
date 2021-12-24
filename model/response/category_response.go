package response

import "restful-api/model/domain"

type CategoryResponse struct {
	Id int `json:"id"`
	Name string `json:"name"`
}

func DetailCategoryResponse(category domain.Category) CategoryResponse {
	return CategoryResponse{
		Id: category.Id,
		Name: category.Name,
	}
}

func ListCategoryResponse(categories []domain.Category) []CategoryResponse {
	var categoriesResponse []CategoryResponse
	for _, category := range categories {
		categoriesResponse = append(categoriesResponse, DetailCategoryResponse(category))
	}

	return categoriesResponse
}