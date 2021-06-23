package util

type RedisConfig struct {
	Host      string `ini:"host"`
	Port      string `ini:"port"`
	DefaultDB string `ini:"default_db"`
	Password  string `ini:"password"`
}
