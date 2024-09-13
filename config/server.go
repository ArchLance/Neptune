package config

type ServerConfig struct {
	//Name      string      `mapstructure:"name" json:"name" yaml:"name"`
	//Port      int         `mapstructure:"port" json:"port" yaml:"port"`
	//Mysqlinfo MysqlConfig `mapstructure:"mysql" json:"mysql" yaml:"mysql"`
	//Admininfo AdminConfig `mapstructure:"admin" json:"admin" yaml:"admin"`
	JWTConfig   JWTConfig   `mapstructure:"jwt" json:"jwt" yaml:"jwt"`
	MysqlConfig MysqlConfig `mapstructure:"mysql" json:"mysql" yaml:"mysql"`
	BaseConfig  BaseConfig  `mapstructure:"base" json:"base" yaml:"base"`
	RedisConfig RedisConfig `mapstructure:"redis" json:"redis" yaml:"redis"`
}

type JWTConfig struct {
	SigningKey string `mapstructure:"key" json:"key" yaml:"key"`
	ExpireTime int    `mapstructure:"exp" json:"exp" yaml:"exp"`
}

type MysqlConfig struct {
	Host     string `mapstructure:"host" json:"host" yaml:"host"`
	Port     string `mapstructure:"port" json:"port" yaml:"port"`
	User     string `mapstructure:"user" json:"user" yaml:"user"`
	Password string `mapstructure:"password" json:"password" yaml:"password"`
	DbName   string `mapstructure:"dbname" json:"dbname" yaml:"dbname"`
}

type RedisConfig struct {
	Host     string `mapstructure:"host" json:"host" yaml:"host"`
	Port     string `mapstructure:"port" json:"port" yaml:"port"`
	Password string `mapstructure:"password" json:"password" yaml:"password"`
	DbName   int    `mapstructure:"dbname" json:"dbname" yaml:"dbname"`
}

type BaseConfig struct {
	Upload Upload `mapstructure:"upload" json:"upload" yaml:"upload"`
}

type Upload struct {
	Avatar string `mapstructure:"avatar" json:"avatar" yaml:"avatar"`
}

// TODO: 邮箱配置
