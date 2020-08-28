package models

type Config struct {
	DBHost            string `json:"DB_HOST"`
	DBUsername        string `json:"DB_USERNAME"`
	DBPassword        string `json:"DB_PASSWORD"`
	DBPort            string `json:"DB_PORT"`
	DBName            string `json:"DB_NAME"`
	Port              string `json:"PORT"`
	Debug             bool   `json:"DEBUG"`
	JwtSecret         string `json:"JWT_SECRET"`
	MailgunApiKey     string `json:"MAILGUN_API"`
	MailDomain        string `json:"MAIL_DOMAIN"`
	BaseURL           string `json:"BASE_URL"`
	NaverClientID     string `json:"NAVER_CLIENT_ID"`
	NaverClientSecret string `json:"NAVER_CLIENT_SECRET"`
}
