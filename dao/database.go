package dao

import (
	"fmt"
	"io/ioutil"
	"time"

	"gopkg.in/yaml.v2"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

var Db *gorm.DB

type conf struct {
	Ip       string `yaml:"ip"`
	UserName string `yaml:"userName"`
	Password string `yaml:"password"`
	DbName   string `yaml:"dbname"`
	Port     int64  `yaml:"port"`
}

func InitSql() (err error) {

	var c conf

	conf := c.getConf()

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local", conf.UserName, conf.Password, conf.Ip, conf.Port, conf.DbName)

	Db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
	})

	if err != nil {
		fmt.Println(err)
	}

	sqlDB, err := Db.DB()

	sqlDB.SetMaxIdleConns(10)

	sqlDB.SetMaxOpenConns(100)

	sqlDB.SetConnMaxLifetime(time.Hour)

	return err
}
func Close() {
	sqlDB, _ := Db.DB()
	sqlDB.Close()
}
func (c *conf) getConf() *conf {
	//讀取resources/application.yaml檔案
	yamlFile, err := ioutil.ReadFile("config.yaml")
	//若出現錯誤，列印錯誤提示
	if err != nil {
		fmt.Println(err.Error())
	}
	//將讀取的字串反序列化
	err = yaml.Unmarshal(yamlFile, c)
	if err != nil {
		fmt.Println(err.Error())
	}
	return c
}
