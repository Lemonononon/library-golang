package db

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"library/utils/yamlparse"
	"log"
)

type Conf struct {
	DBName       string `yaml:"DBName"`
	Timeout      string `yaml:"Timeout"`
	ReadTimeout  string `yaml:"ReadTimeout"`
	WriteTimeout string `yaml:"WriteTimeout"`
	MaxIdleConns int    `yaml:"MaxIdleConns"`
	MaxOpenConns int    `yaml:"MaxOpenConns"`
	Collation    string `yaml:"Collation"`
	Host         string `yaml:"Host"`
	Username     string `yaml:"Username"`
	Password     string `yaml:"Password"`
}

func readConfig(configName string) (map[string]Conf, error) {
	confMap := make(map[string]Conf)

	confFilePath := "./config/" + configName + ".yaml"
	err := yamlparse.LoadYamlFile(confFilePath, &confMap)
	if err != nil {
		return nil, err
	}
	return confMap, nil
}
func initMySQL() {
	confMap, err := readConfig("rdb")
	conf, _ := confMap["Library"]
	if err != nil {
		panic(err)
	}
	dsn := fmt.Sprintf(`%s:%s@tcp(%s)/%s?collation=%s&parseTime=True&loc=Local&timeout=%s&readTimeout=%s&writeTimeout=%s`,
		conf.Username,
		conf.Password,
		conf.Host,
		conf.DBName,
		conf.Collation,
		conf.Timeout,
		conf.ReadTimeout,
		conf.WriteTimeout)
	MySQL, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("db init failed")
	}
}
