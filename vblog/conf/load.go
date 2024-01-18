package conf

import (
	"github.com/BurntSushi/toml"
	"github.com/caarlos0/env/v6"
)

// 该文件定义配置对象的加载方法
// github.com/BurntSushi/toml Go中使用比较广泛的toml格式解析库
// https://github.com/BurntSushi/toml 查看该库的基本用法
// 将toml文件和对象进行映射 object <---> toml 配置文件
func LoadFromFile(filepath string) error {
	c := DefaultConfig()
	if _, err := toml.DecodeFile(filepath, c); err != nil {
		return err
	}
	config = c
	return nil
}

// "github.com/caarlos0/env/v6" 读取环境变量
// env ---> object
func LoadFromEnv() error {
	c := DefaultConfig()
	// env Tag
	if err := env.Parse(c); err != nil {
		return err
	}
	config = c
	// c.MySQL.Host = os.Getenv("DATASOURCE_HOST")
	return nil
}
