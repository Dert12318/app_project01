package models
//Model to run server

//Server config
type ServerConfig struct {
	Name 				string `env:"NAME_SERVER"`
	Port                string `env:"PORT_SERVER,required"`
	Host                string `env:"HOST_SERVER,required"`
	JSONPathFile        string `env:"JSON_PATHFILE,required"`
	PostgresConfig     	PostgresConfig
}

//Postgres Config
type PostgresConfig struct {
	Name     string `env:"NAME_POSTGRES,required"`
	Host     string `env:"HOST_POSTGRES,required"`
	Port     string `env:"PORT_POSTGRES,required"`
	User     string `env:"USER_POSTGRES"`
	Password string `env:"PASS_POSTGRES"`
}
