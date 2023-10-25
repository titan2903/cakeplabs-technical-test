package config

import (
	"cakeplabs-technical-test/entity"
	"cakeplabs-technical-test/pkg/utils"
	"cakeplabs-technical-test/repositories"
	"cakeplabs-technical-test/services"
	"log"
	"sync"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	c            services.Services
	onController sync.Once
)

func GetController() services.Services {
	onController.Do(func() {
		c = services.NewServices(GetRepository())
	})

	return c
}

var (
	repo    repositories.Repositories
	oneRepo sync.Once
)

func GetRepository() repositories.Repositories {
	oneRepo.Do(func() {
		repo = repositories.NewRepositories(GetQuery())
	})

	return repo
}

var (
	qry     *gorm.DB
	oneSync sync.Once
)

func GetQuery() *gorm.DB {

	oneSync.Do(func() {
		// ! Connect to PostgreSQL database
		dsnMaster := utils.GoDotEnvVariable("DATABASE_URL")
		dbMaster, errMaster := gorm.Open(postgres.Open(dsnMaster), &gorm.Config{})
		if errMaster != nil {
			log.Panic(errMaster)
		}

		//! ----------------------------------------------------------------------------------------------

		if errMaster = dbMaster.AutoMigrate(
			&entity.Menu{},
			&entity.Addition{},
			&entity.Order{},
		); errMaster != nil {
			log.Fatal(errMaster)
		}

		log.Println("success connect to database")
		qry = dbMaster
	})

	return qry
}
