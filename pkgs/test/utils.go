package test_utils

import (
	"OnlineStoreBackend/db"
	"OnlineStoreBackend/pkgs/config"

	"gorm.io/gorm"
)

func InitTestDB(cfg *config.Config) *gorm.DB {
	db := db.Init(cfg)

	return db
}

func PrepareAllConfiguration(path string) *config.Config {
	cfg, err := config.Load([]string{path}, true, nil)
	if err != nil {
		panic(err)
	}

	return cfg
}
