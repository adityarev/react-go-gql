package queries

import (
	"context"

	"github.com/graphql-go/graphql"
	"github.com/mongodb/mongo-go-diver/bson"

	"app/data"
	types "app/types"
)

type todoStruct struct {
	name				string `json: "name"`
	description	string `json: "description"`
}

var GetNotTodos = &graphql.Field {
	Type: graphql.NewList(types.NotTodo),
	Description: "Get all not todos",
	Resolve: func(params graphql.ResolveParams) (interface {}, error) {
		notTodoCollection := mongo.Client.Database("medium-app").Collection("not_todos")
		todos, err := notTodoCollection.Find(context.Background(), nil)
		if err != nil {
			panic(err)
		}

		var todoList []todoStruct
		for todos.Next(context.Background()) {
			doc := bson.NewDocument()
			err := todos.Decode(doc)
			if err != nil {
				panic(err)
			}

			keys, err := doc.Keys(false)
			if err != nil {
				panic(err)
			}

			// Convert BSON to struct
			todo := todoStruct{}
			for _, key := range keys {
				keyString := key.String()
				elm, err := doc.Lookup(keyString)
				if err != nil {
					panic(err)
				}

				switch(keyString) {
				case "name":
					todo.name = elm.Value().StringValue()
				case "description":
					todo.description = elm.Value().StringValue()
				default:
				}
			}
			todoList = append(todoList, todo)
		}

		return todoList, nil
	},
}
