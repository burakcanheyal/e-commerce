package dto

type Config struct {
	DBURL   string `mapstructure:"DB_URL"`
	Secret  string `mapstructure:"SECRET"`
	Secret2 string `mapstructure:"SECRET2"`
}
