package repositories

import "cakeplabs-technical-test/entity"

func (r *repositories) GetAllMenusByKeyword(keywords string) []entity.Menu {
	var menus []entity.Menu
	_ = r.db.Where("name LIKE ?", "%"+keywords+"%").Find(&menus)

	return menus
}

func (r *repositories) GetMenuById(id int) (entity.Menu, error) {
	var menu entity.Menu
	result := r.db.Where("id = ?", id).First(&menu)
	return menu, result.Error
}

func (r *repositories) GetAllMenus() []entity.Menu {
	var menus []entity.Menu
	_ = r.db.Find(&menus)

	return menus
}
