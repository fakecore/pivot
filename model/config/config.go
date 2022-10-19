package config

type Config struct {
	Sys   SysConfig   `mapstructure:"sys" json:"sys" yaml:"sys"`
	DB    DBConfig    `mapstructure:"db" json:"db" yaml:"db"`
	Jwt   JwtConfig   `mapstructure:"jwt" json:"jwt" yaml:"jwt"`
	Redis RedisConfig `mapstructure:"redis" json:"redis" yaml:"redis"`
}

type MysqlConfig struct {
	Url          string `mapstructure:"url" json:"url" yaml:"url"`                                // 服务器地址:端口
	Config       string `mapstructure:"config" json:"config" yaml:"config"`                       // 高级配置
	DbName       string `mapstructure:"db-name" json:"dbname" yaml:"db-name"`                     // 数据库名
	UserName     string `mapstructure:"user-name" json:"username" yaml:"username"`                // 数据库用户名
	Password     string `mapstructure:"password" json:"password" yaml:"password"`                 // 数据库密码
	MaxIdleConns int    `mapstructure:"max-idle-conns" json:"maxIdleConns" yaml:"max-idle-conns"` // 空闲中的最大连接数
	MaxOpenConns int    `mapstructure:"max-open-conns" json:"maxOpenConns" yaml:"max-open-conns"` // 打开到数据库的最大连接数
}
type DBConfig struct {
	Type  string      `mapstructure:"type" json:"type" yaml:"type"`
	Mysql MysqlConfig `mapstructure:"mysql" json:"mysql" yaml:"mysql"`
}

type SysConfig struct {
	//TO distinguish current running environment is dev or prod
	Env  string `mapstructure:"env" json:"env" yaml:"env"`
	Port int    `mapstructure:"port" json:"port" yaml:"port"`
}

type JwtConfig struct {
	//TO distinguish current running environment is dev or prod
	Key         string `mapstructure:"key" json:"key" yaml:"key"`
	ExpiredTime int64  `mapstructure:"expiredTime" json:"expiredTime" yaml:"expiredTime" comment:"expired_time hour"`
	BufferTime  int64  `mapstructure:"bufferTime" json:"bufferTime" yaml:"bufferTime" comment:" hour"`
	Issuer      string `mapstructure:"issuer" json:"issuer" yaml:"issuer" `
}

type RedisConfig struct {
	Addr     string `mapstrcuture:"addr" json:"addr" yaml:"addr" comment:"ip:port"`
	Password string `mapstrcuture:"password" json:"password" yaml:"password" comment:"ip:port"`
}
