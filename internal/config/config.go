package config

import "os"

type config struct {
	MongoDBURI           string
	Port                 string
	JWTSecret            string
	ImgBBAPI             string
	StripeKey            string
	StripeEndpointSecret string
	SecretKey            string
	MailUsername         string
	MailPassword         string
}

func getEnv(key, def string) string {
	val := os.Getenv(key)
	if val == "" {
		return def
	}
	return val
}

func LoadConfig() *config {
	return &config{
		MongoDBURI:           getEnv("MONGO_URI", "mongodb://localhost:27017"),
		Port:                 ":" + getEnv("PORT", "8080"),
		JWTSecret:            getEnv("JWT_SECRET", ""),
		ImgBBAPI:             getEnv("IMG_API", ""),
		StripeKey:            getEnv("STRIPE_SECRET_KEY", ""),
		StripeEndpointSecret: getEnv("ENDPOINT_SECRET", ""),
		SecretKey:            getEnv("SECRET_KEY", ""),
		MailUsername:         getEnv("MAIL_USERNAME", ""),
		MailPassword:         getEnv("MAIL_PASSWORD", ""),
	}
}
