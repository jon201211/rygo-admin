package cfg

import (
	"sync"

	"github.com/BurntSushi/toml"
)

var (
	instance *config
	once     sync.Once
)

//获取配置文档实例
func Instance() *config {
	once.Do(func() {
		var conf config
		filePath := "./config.toml"
		if _, err := toml.DecodeFile(filePath, &conf); err != nil {
			return
		}
		instance = &conf
	})

	return instance
}

type config struct {
	Status   status
	Admin    admin
	JWT      jwt
	Task     task
	Database database
	Logger   logger
	Gen      gen
}

type status struct {
	Admin bool
	Api   bool
}

type admin struct {
	Address    string
	ServerRoot string
	Swagger    string
}

type jwt struct {
	Timeout    int
	Refresh    int
	EncryptKey string
}

type task struct {
	WorkPoolSize int
}

type database struct {
	Driver string
	Master string
	Slave  string
	Debug  bool
	Log    string
}

type logger struct {
	Path   string
	Level  uint32
	Stdout bool
}

type gen struct {
	Author      string
	ModuleName  string
	PackageName string
}
