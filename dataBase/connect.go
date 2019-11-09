package dataBase

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"shSpider_plus/model"

	//必须在这里手动导入mysql驱动 前面要加_
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/joho/godotenv"
	"os"
	"shSpider_plus/Vcb/vcbModel"
	"time"
)

func Connect()  {
	err := godotenv.Load(".env")
	fmt.Println("读取.env文件作为环境变量：...")
	if err!=nil {
		fmt.Printf("%s",err)
		panic(err)
	}

	mysqlDsn := os.Getenv("MYSQL_DSN")
	fmt.Println("从环境变量中读取到mysqlDsn：",mysqlDsn)
	//连接数据库
	Database(mysqlDsn)
}



// DB 数据库链接单例
var DB *gorm.DB

// Database 在中间件中初始化mysql链接
func Database(connString string) {
	db, err := gorm.Open("mysql", connString)
	// Error
	if err != nil {
		fmt.Println("连接数据库不成功", err)
	}

	db.LogMode(true)

	//设置连接池
	//空闲
	db.DB().SetMaxIdleConns(50)
	//打开
	db.DB().SetMaxOpenConns(100)
	//超时
	db.DB().SetConnMaxLifetime(time.Second * 30)

	DB = db

	migration()
}


// 自动迁移模式
func migration() {
	DB.AutoMigrate(&vcbModel.Section{})
	DB.AutoMigrate(&vcbModel.Invitation{})
	DB.AutoMigrate(&model.Profile{})
}