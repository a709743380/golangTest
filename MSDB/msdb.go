package msdb

import (
	"database/sql"
	"fmt"
	"log"
	"os" // 导入 os 包

	_ "github.com/microsoft/go-mssqldb" // SQL Server 驱动
	"gopkg.in/yaml.v3"
)

// 配置结构体，用于存储数据库连接信息
type Config struct {
	DB_SERVER   string `yaml:"DB_SERVER"`
	DB_PORT     string `yaml:"DB_PORT"`
	DB_USER     string `yaml:"DB_USER"`
	DB_PASSWORD string `yaml:"DB_PASSWORD"`
	DB_DATABASE string `yaml:"DB_DATABASE"`
}

// 读取配置文件并解析成 Config 结构体
func LoadConfig(filename string) (Config, error) {
	var config Config
	// 读取配置文件
	file, err := os.ReadFile(filename)
	if err != nil {
		return config, err
	}
	// 解析 YAML 配置文件
	err = yaml.Unmarshal(file, &config)
	if err != nil {
		return config, err
	}
	return config, nil
}

// 初始化数据库连接并返回 *sql.DB 对象
func InitDB() (*sql.DB, error) {

	config, err := LoadConfig("msconfig.yaml")
	// 使用配置中的信息构建连接字符串
	connString := fmt.Sprintf("sqlserver://%s:%s@%s:%s?database=%s",
		config.DB_USER, config.DB_PASSWORD, config.DB_SERVER, config.DB_PORT, config.DB_DATABASE)

	// 打开数据库连接
	db, err := sql.Open("sqlserver", connString)
	if err != nil {
		return nil, fmt.Errorf("数据库连接失败: %v", err)
	}

	// 测试数据库连接
	if err := db.Ping(); err != nil {
		return nil, fmt.Errorf("无法连接到数据库: %v", err)
	}

	log.Println("数据库连接成功")
	return db, nil
}
