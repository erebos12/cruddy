package configuration

import "os"

func GetPort() string {
	return os.Getenv("GOAPP_PORT")
}

func GetMongoUrl() string {
	return os.Getenv("MONGO_URL")
}

func GetMongoDB() string {
	return os.Getenv("MONGO_DB")
}

func GetMongoCollection() string {
	return os.Getenv("MONGO_COLLECTION")
}
