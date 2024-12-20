package database

import "go.mongodb.org/mongo-driver/mongo/options"

func Connect() err {
	opts := options.Client().ApplyURI("//URI AQ").SetServerAPIOptions(serverAPI)

}
