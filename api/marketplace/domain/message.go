package marketplace_domain

type Message struct {
	Data string `json:"data"`
}

type MessageMongo struct {
	Data string `bson:"data,omitempty"`
}
