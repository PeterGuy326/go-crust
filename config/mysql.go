package config

import (
	"fmt"
	"strings"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"github.com/spf13/viper"
)

// Db数据库连接池
var DB *gorm.DB

// 注意方法名大写，就是public
func InitMysqlDB() {

	// 数据库参数拉取配置
	userName := viper.Get("mysql.userName").(string)
	password := viper.Get("mysql.password").(string)
	ip := viper.Get("mysql.ip").(string)
	port := viper.Get("mysql.port").(string)
	dbName := viper.Get("mysql.dbName").(string)

	// 构建连接："用户名:密码@tcp(IP:端口)/数据库?charset=utf8"
	path := strings.Join([]string{userName, ":", password, "@tcp(", ip, ":", port, ")/", dbName, "?charset=utf8&parseTime=true"}, "")

	// 打开数据库,前者是驱动名，所以要导入： _ "github.com/go-sql-driver/mysql"
	DB, _ = gorm.Open("mysql", path)

	fmt.Println("connect success")
}
