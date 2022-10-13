package config

type Config struct {
	Sys SysConfig `value:"${db}"`
	DB  DBConfig  `value:"${db}"`
}

type DBConfig struct {
	UserName     string `value:"${username}"`
	Password     string `value:"${password}"`
	Url          string `value:"${url}"`
	Port         string `value:"${port}"`
	DBName       string `value:"${db}"`
	Config       string `value:"${config}"`
	Type         string `value:"${type}"`
	MaxIdleConns int    `mapstructure:"max-idle-conns" json:"max-idle-conns" yaml:"max-idle-conns"` // 空闲中的最大连接数
	MaxOpenConns int    `mapstructure:"max-open-conns" json:"max-open-conns" yaml:"max-open-conns"` // 打开到数据库的最大连接数
}

type SysConfig struct {
	//TO distinguish current running environment is dev or prod
	Env  string `mapstructure:"env" json:"env" yaml:"env"`
	Port int    `mapstructure:"port" json:"port" yaml:"port"`
}
