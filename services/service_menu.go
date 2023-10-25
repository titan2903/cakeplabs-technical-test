package services

import (
	"cakeplabs-technical-test/pkg/constant"
	"cakeplabs-technical-test/transport"
	"net/http"

	"github.com/labstack/gommon/log"
)

func (s *services) GetMenuById(id int) (*transport.Response, error) {
	menu, err := s.repo.GetMenuById(id)
	if err != nil {
		log.Error("GetMenuById", err)
		return nil, err
	}

	MenuMapping := &transport.MenuMapping{
		ID:        menu.ID,
		Name:      menu.Name,
		Price:     menu.Price,
		CreatedAt: menu.CreatedAt,
		UpdatedAt: menu.UpdatedAt,
	}

	result := transport.ApiResponse("Success get menu", http.StatusOK, constant.StatusSuccess, MenuMapping)
	return &result, nil
}

func (s *services) GetAllMenusByKeyword(keywords string) *transport.Response {
	menus := s.repo.GetAllMenusByKeyword(keywords)

	var arrMenus []transport.MenuMapping
	for _, menu := range menus {
		menuMapping := transport.MenuMapping{
			ID:        menu.ID,
			Name:      menu.Name,
			Price:     menu.Price,
			CreatedAt: menu.CreatedAt,
			UpdatedAt: menu.UpdatedAt,
		}

		arrMenus = append(arrMenus, menuMapping)
	}

	result := transport.ApiResponse("Success get all menu data", http.StatusOK, constant.StatusSuccess, arrMenus)

	if len(arrMenus) == 0 {
		result = transport.ApiResponse("There is no menu list yet", http.StatusOK, constant.StatusSuccess, len(arrMenus))
	}

	return &result
}
