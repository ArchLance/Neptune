package config

type ServerConfig struct {
	//Name      string      `mapstructure:"name" json:"name" yaml:"name"`
	//Port      int         `mapstructure:"port" json:"port" yaml:"port"`
	//Mysqlinfo MysqlConfig `mapstructure:"mysql" json:"mysql" yaml:"mysql"`
	//Admininfo AdminConfig `mapstructure:"admin" json:"admin" yaml:"admin"`
	JWTKey   JWTConfig  `mapstructure:"jwt" json:"jwt" yaml:"jwt"`
	DBconfig DBConfig   `mapstructure:"mysql" json:"mysql" yaml:"mysql"`
	BaseConf BaseConfig `mapstructure:"base" json:"base" yaml:"base"`
}

type JWTConfig struct {
	SigningKey string `mapstructure:"key" json:"key" yaml:"key"`
	ExpireTime int    `mapstructure:"exp" json:"exp" yaml:"exp"`
}

type DBConfig struct {
	Host     string `mapstructure:"host" json:"host" yaml:"host"`
	Port     string `mapstructure:"port" json:"port" yaml:"port"`
	User     string `mapstructure:"user" json:"user" yaml:"user"`
	Password string `mapstructure:"password" json:"password" yaml:"password"`
	DbName   string `mapstructure:"dbname" json:"dbname" yaml:"dbname"`
}

type BaseConfig struct {
	Upload Upload `mapstructure:"upload" json:"upload" yaml:"upload"`
}

type Upload struct {
	Avatar string `mapstructure:"avatar" json:"avatar" yaml:"avatar"`
}
