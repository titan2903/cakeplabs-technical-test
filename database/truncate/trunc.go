package truncate

import (
	"cakeplabs-technical-test/database/config"
	"log"

	"gorm.io/gorm"
)

type trunc struct {
	DB *gorm.DB
}

func NewTrunc() *trunc {
	return &trunc{config.GetQuery()}
}

func (t *trunc) DeleteDataMenus() {
	log.Println("success delete data menus")
	t.DB.Exec("DELETE FROM menus")
}

func (t *trunc) DeleteDataAdditions() {
	log.Println("success delete data additions")
	t.DB.Exec("DELETE FROM additions")
}
