package services

import (
	"cakeplabs-technical-test/pkg/constant"
	"cakeplabs-technical-test/transport"
	"net/http"

	"github.com/labstack/gommon/log"
)

func (s *services) SeedData() (*transport.Response, error) {
	menus := s.repo.GetAllMenus()
	additions := s.repo.GetAllAdditions()

	if len(menus) == 0 {
		err := s.repo.MenusSeeder()
		if err != nil {
			log.Error("MenusSeeder", err)
			return nil, err
		}

	}

	if len(additions) == 0 {
		err := s.repo.AdditionsSeeder()
		if err != nil {
			log.Error("AdditionsSeeder", err)
			return nil, err
		}
	}

	result := transport.ApiResponse("Success seed data menus, toppings and fillings", http.StatusOK, constant.StatusSuccess, nil)
	return &result, nil
}
