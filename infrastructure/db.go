package infrastructure

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"os"
	"path"
)

var DB *gorm.DB

func OpenDbConnection() *gorm.DB {
	dialect := os.Getenv("DB_DIALECT")
	username := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")
	host := os.Getenv("DB_HOST")
	tablePrefix := os.Getenv("TABLE_PREFIX")
	var db *gorm.DB
	var err error
	if dialect == "sqlite3" {
		db, err = gorm.Open("sqlite3", path.Join(".", "app.db"))
	} else {
		// db, err := gorm.Open("mysql", "root:root@localhost/go_api_shop_gonc?charset=utf8")
		databaseUrl := fmt.Sprintf("%s:%s@(%s)/%s?charset=utf8&parseTime=True&loc=Local", username, password,host, dbName)
		db, err = gorm.Open(dialect, databaseUrl)
	}

	if err != nil {
		fmt.Println("db err: ", err)
		os.Exit(-1)
	}
	db.DB().SetMaxIdleConns(10)
	db.LogMode(true)
	db.SingularTable(true)
	//设置默认表名前缀
	gorm.DefaultTableNameHandler = func(db *gorm.DB, defaultTableName string) string {
		return tablePrefix + defaultTableName
	}

	DB = db
	return DB
}


// Using this function to get a connection, you can create your connection pool here.
func GetDb() *gorm.DB {
	return DB
}