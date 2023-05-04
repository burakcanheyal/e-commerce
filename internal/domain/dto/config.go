package dto

type Config struct {
	DBURL        string `mapstructure:"DB_URL"`
	Secret       string `mapstructure:"SECRET"`
	Secret2      string `mapstructure:"SECRET2"`
	WkHtmlPath   string `mapstructure:"WKHTMLTOPDF_PATH"`
	MailAddress  string `mapstructure:"MAIL_ADDRESS"`
	MailPassword string `mapstructure:"MAIL_PASS"`
}
