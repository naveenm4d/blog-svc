package config

import (
	"os"

	flag "github.com/spf13/pflag"
)

type Configuration struct {
	GrpcPort *string

	MongoDBEndpoint      *string
	MongoDBDatabase      *string
	MongoPostsCollection *string
}

var (
	Config *Configuration

	grpcPort = flag.String(
		"grpc-port",
		"5007",
		"the port to serve on")

	mongoDBEndpoint = flag.String(
		"mongodb-endpoint",
		"mongodb://root:password@127.0.0.1:27017",
		"Mongodb Endpoint.")

	mongoDBDatabase = flag.String(
		"mongodb-database",
		"blog_svc",
		"Mongodb Database.")

	mongoPostsCollection = flag.String(
		"posts-svc-collection",
		"posts",
		"Posts Mongodb Collection")
)

func updateStringEnvVariable(defValue *string, key string) *string {
	val := os.Getenv(key)

	if val == "" {
		return defValue
	}

	return &val
}

func init() {
	flag.Parse()

	grpcPort = updateStringEnvVariable(grpcPort, "GRPC_PORT")

	mongoDBEndpoint = updateStringEnvVariable(mongoDBEndpoint, "MONGODB_ENDPOINT")
	mongoDBDatabase = updateStringEnvVariable(mongoDBDatabase, "MONGODB_DATABASE")

	Config = &Configuration{
		GrpcPort: grpcPort,

		// MongoDB
		MongoDBEndpoint:      mongoDBEndpoint,
		MongoDBDatabase:      mongoDBDatabase,
		MongoPostsCollection: mongoPostsCollection,
	}
}
