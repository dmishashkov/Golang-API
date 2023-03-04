package schemas

type JWTConfig struct {
	TokenDuration int
	SecretString  string
}

type DatabaseConfig struct {
	User     string
	Password string
	Host     string
	Port     string
	DBName   string
}
type Config struct {
	DB  DatabaseConfig
	JWT JWTConfig
}
