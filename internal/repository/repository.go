package repository

import (
	"github.com/getclasslabs/course/internal/config"
	"github.com/getclasslabs/go-tools/pkg/db"
	_ "github.com/go-sql-driver/mysql"
)

func Start() {
	Db = &db.MySQL{}
	Db.Connect(config.Config.Mysql)
}

var Db db.Database
