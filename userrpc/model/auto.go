package model

import "zg5/z304/framework/mysql"

func AutoTable() error {
	return mysql.DB.AutoMigrate(new(User))
}
