package config

type Config struct {
	Server  Server  `mapstructure:"server"`
	DB      DBConf  `mapstructure:"psql"`
	Session Session `mapstructure:"session"`
}

type Server struct {
	Port string `mapstructure:"port"`
}

type DBConf struct {
	Host     string `mapstructure:"host"`
	Port     int    `mapstructure:"port"`
	User     string `mapstructure:"user"`
	Password string `mapstructure:"password"`
	DBName   string `mapstructure:"dbname"`
}

type Redis struct {
	ConnectionTimeoutSeconds string `mapstructure:"connection_timeout_seconds"`
	NetworkType              string `mapstructure:"network_type"`
	Host                     string `mapstructure:"host"`
	Port                     string `mapstructure:"port"`
	Password                 string `mapstructure:"password"`
}

type Session struct {
	SecretKey   string `mapstructure:"session_key"`
	SessionName string `mapstructure:"sessionName"`
	SessionKey  string `mapstructure:"sessionKey"`
	Redis       Redis  `mapstructure:"redis"`
}
