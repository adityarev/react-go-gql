package mongo

import (
	"context"
	"github.com/mongodb/mongo-go-driver/mongo"
)

var Client, err = mongo.Connect(context.Background(), "nsert your MongoDB URI here!", nil)
