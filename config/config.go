package config

import "os"

type Env struct {
	Port                         string
	FirestoreProjectID           string
	FirebaseAPIKey               string
	GoogleApplicationCredentials string
	JWTSecret                    string
	RedisAddress                 string
	RedisPassword                string
}

func LoadEnv() Env {
	return Env{
		Port:                         getEnv("PORT", "3000"),
		FirestoreProjectID:           getEnv("FIRESTORE_PROJECT_ID", ""),
		FirebaseAPIKey:               getEnv("FIREBASE_API_KEY", ""),
		GoogleApplicationCredentials: getEnv("GOOGLE_APPLICATION_CREDENTIALS", ""),
		RedisAddress:                 getEnv("REDIS_ADDRESS", ""),
		RedisPassword:                getEnv("REDIS_PASSWORD", ""),
	}
}

func getEnv(key, fallback string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return fallback
}
