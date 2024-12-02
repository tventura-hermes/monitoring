package envs_helpers

// import (
// 	validations_helpers "demo/helpers/validations"
// 	"encoding/json"
// 	"log"
// 	"os"
// )

// type envStruct struct {
// 	CACHE_TYPE          string `json:"CACHE_TYPE" validate:"required"`
// 	CACHE_HOST          string `json:"CACHE_HOST" validate:"required"`
// 	MONGO_URI           string `json:"MONGO_URI" validate:"required"`
// 	PUBSUB_TOPIC        string `json:"PUBSUB_TOPIC" validate:"required"`
// 	PUBSUB_SUBSCRIPTION string `json:"PUBSUB_SUBSCRIPTION" validate:"required"`
// }

// func SetupEnvs() (envStruct, error) {
// 	env := os.Getenv("ENV_JSON")
// 	var data envStruct

// 	if err := json.Unmarshal([]byte(env), &data); err != nil {
// 		return envStruct{}, err
// 	}

// 	if err := validations_helpers.StructValidator.Struct(data); err != nil {
// 		log.Printf("Could setup env vars: %v", err)
// 		return envStruct{}, nil
// 	}

// 	return data, nil
// }
