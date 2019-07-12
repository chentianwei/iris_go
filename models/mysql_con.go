package models

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"log"
	"sync"
	"os"
)

/*
* MysqlConnectiPool
* 数据库连接操作库
* 基于gorm封装开发
 */
type MysqlConnectiPool struct {}

var instance *MysqlConnectiPool
var once sync.Once

var db *gorm.DB
var err_db error

func GetInstance() *MysqlConnectiPool {
	once.Do(func() {
		instance = &MysqlConnectiPool{}
	})
	return instance
}

/*
* @fuc 初始化数据库连接
*/
func (m *MysqlConnectiPool) InitDataPool() (issucc bool) {
	mysql_config := os.Getenv("DB_USERNAME") + ":" + os.Getenv("DB_PASSWORD") + "@tcp(" + os.Getenv("DB_HOST") + ":3306)/" + os.Getenv("DB_DATABASE") + "?charset=" + os.Getenv("DB_CHARSET") + "&parseTime=True&loc=Local"
	//"user:password@tcp(192.168.1.4:3306)/dbname?charset=utf8&parseTime=True&loc=Local"
	db, err_db = gorm.Open("mysql", mysql_config)
	fmt.Println(err_db)
	if err_db != nil {
		log.Fatal(err_db)
		return false
	}
	//SetMaxOpenConns用于设置最大打开的连接数
	//SetMaxIdleConns用于设置闲置的连接数
	db.DB().SetMaxIdleConns(10)
	db.DB().SetMaxOpenConns(100)
	//关闭数据库，db会被多个goroutine共享，可以不调用
	// defer db.Close()
	db.SingularTable(true)
	db.LogMode(true)
	return true
}

/*
* @fuc  对外获取数据库连接对象db
*/
func (m *MysqlConnectiPool) GetMysqlDB() (db_con *gorm.DB) {
	return db
}

