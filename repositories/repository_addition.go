package repositories

import "cakeplabs-technical-test/entity"

func (r *repositories) GetAllAdditionsByKeyword(keywords string) []entity.Addition {
	var additions []entity.Addition
	_ = r.db.Where("name LIKE ?", "%"+keywords+"%").Find(&additions)

	return additions
}

func (r *repositories) GetAdditionById(id int) (entity.Addition, error) {
	var addition entity.Addition
	result := r.db.Where("id = ?", id).First(&addition)
	return addition, result.Error
}

func (r *repositories) GetAllAdditions() []entity.Addition {
	var additions []entity.Addition
	_ = r.db.Find(&additions)

	return additions
}
