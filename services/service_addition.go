package services

import (
	"cakeplabs-technical-test/pkg/constant"
	"cakeplabs-technical-test/transport"
	"net/http"

	"github.com/labstack/gommon/log"
)

func (s *services) GetAdditionById(id int) (*transport.Response, error) {
	addition, err := s.repo.GetAdditionById(id)
	if err != nil {
		log.Error("GetAdditionById", err)
		return nil, err
	}

	additionMapping := &transport.AdditionMapping{
		ID:        addition.ID,
		Name:      addition.Name,
		Price:     addition.Price,
		CreatedAt: addition.CreatedAt,
		UpdatedAt: addition.UpdatedAt,
	}

	result := transport.ApiResponse("Success get addition", http.StatusOK, constant.StatusSuccess, additionMapping)
	return &result, nil
}

func (s *services) GetAllAdditionsByKeyword(keywords string) *transport.Response {
	additions := s.repo.GetAllAdditionsByKeyword(keywords)

	var arrAdditions []transport.AdditionMapping
	for _, addition := range additions {
		additionMapping := transport.AdditionMapping{
			ID:        addition.ID,
			Name:      addition.Name,
			Price:     addition.Price,
			CreatedAt: addition.CreatedAt,
			UpdatedAt: addition.UpdatedAt,
		}

		arrAdditions = append(arrAdditions, additionMapping)
	}

	result := transport.ApiResponse("Success get all addition data", http.StatusOK, constant.StatusSuccess, arrAdditions)

	if len(arrAdditions) == 0 {
		result = transport.ApiResponse("There is no addition list yet", http.StatusOK, constant.StatusSuccess, len(arrAdditions))
	}

	return &result
}
