package conf

import (
	"encoding/json"
	"fmt"
	"sync"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// 文件内容说明：
// 一个Config对象定义，Config对象有一个方法。
// 一个MySQL对象定义，MySQL对象有两个方法。MySQL对象组合到Config对象中。

// 这里不采用直接暴露变量的方式，比较好的方式是使用函数
var config *Config

// 这里就可以补充逻辑。这是conf包的普通函数。
func C() *Config {
	// sync.Lock
	if config == nil {
		// 给个默认值
		config = DefaultConfig()
	}
	return config
}

func DefaultConfig() *Config {
	return &Config{
		Application: &Application{
			Domain: "127.0.0.1",
		},
		MySQL: &MySQL{
			Host:     "192.168.0.77",
			Port:     3306,
			DB:       "vblog",
			Username: "root",
			Password: "123456",
			Debug:    true,
			// Debug打开的目的：希望GORM的时候把信息打印出来
		},
	}
}

// 程序配置对象，启动时会读取配置，并且为程序提供需要的全局变量
// 把配置对象做成全局变量（单例模式）
// toml
/*
[mysql]
host="127.0.0.1"
port=3306
...
*/
type Config struct {
	Application *Application `json:"app" yaml:"app" toml:"app"`
	MySQL *MySQL `json:"mysql" yaml:"mysql" toml:"mysql"`
}

type Application struct {
	Domain string `json:"domain" yaml:"domain" toml:"domain" env:"APP_DOMAIN"`
}

// fmt.Stringer
// 如果你想要自定义 对象fmt.PrintXXX() 打印的值
// String() string
// &{0xc0000263c0} ----> JSON {}
func (c *Config) String() string {
	// jd ----> json data
	jd, err := json.MarshalIndent(c, "", "  ")
	if err != nil {
		return fmt.Sprintf("%p", c)
	}
	return string(jd)
}

// db对象也是一个单例模式
// 这是一个精巧的结构体设计
// 这里把关于数据库连接字符串的各种配置参数和实际连接数据库的句柄、同步锁等组合在一起，形成一个结构体
// 然后通过结构体方法巧妙地把从配置文件中读取的参数进行处理后，去实际连接数据库，拿到数据库连接池中的一个数据连接
// 这里所有的字段都和MySQL有关系，所以放在一个结构体内没有违和感
type MySQL struct {
	Host     string `json:"host" yaml:"host" toml:"host" env:"DATASOURCE_HOST"`
	Port     int    `json:"port" yaml:"port" toml:"port" env:"DATASOURCE_PORT"`
	DB       string `json:"database" yaml:"database" toml:"database" env:"DATASOURCE_DB"`
	Username string `json:"username" yaml:"username" toml:"username" env:"DATASOURCE_USERNAME"`
	Password string `json:"password" yaml:"password" toml:"password" env:"DATASOURCE_PASSWORD"`
	Debug    bool   `json:"debug" yaml:"debug" toml:"debug" env:"DATASOURCE_DEBUG"`

	// 判断这个私有属性，来判断是否返回已有的对象
	db *gorm.DB
	l  sync.Mutex
}

// dsn := "user:pass@tcp(127.0.0.1:3306)/dbname?charset=utf8mb4&parseTime=True&loc=Local"
func (m *MySQL) DSN() string {
	return fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		m.Username,
		m.Password,
		m.Host,
		m.Port,
		m.DB,
	)
}

// 通过配置就能获取一个DB实例
func (m *MySQL) GetDB() *gorm.DB {
	// 避免多个Goroutine，同时执行打开操作，加互斥锁。把并行变成串行。
	m.l.Lock()
	defer m.l.Unlock()

	if m.db == nil {
		db, err := gorm.Open(mysql.Open(m.DSN()), &gorm.Config{})
		if err != nil {
			panic(err)
		}
		m.db = db

		// 补充Debug配置
		if m.Debug {
			m.db = db.Debug()
		}
	}

	return m.db
}

// 配置对象提供全局单例配置
func (c *Config) DB() *gorm.DB {
	return c.MySQL.GetDB()
}
