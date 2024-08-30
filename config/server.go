package config

type ServerConfig struct {
	//Name      string      `mapstructure:"name" json:"name" yaml:"name"`
	//Port      int         `mapstructure:"port" json:"port" yaml:"port"`
	//Mysqlinfo MysqlConfig `mapstructure:"mysql" json:"mysql" yaml:"mysql"`
	//Admininfo AdminConfig `mapstructure:"admin" json:"admin" yaml:"admin"`
	JWTKey JWTConfig `mapstructure:"jwt" json:"jwt" yaml:"jwt"`
}

type JWTConfig struct {
	SigningKey string `mapstructure:"key" json:"key" yaml:"key"`
	ExpireTime int    `mapstructure:"exp" json:"exp" yaml:"exp"`
}
