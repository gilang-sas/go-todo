package model

import (
	"time"
	"context"
	"github.com/gilang-sas/todo-app/db"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/bson"
)

type ToDo struct{
	ID			primitive.ObjectID	`json:"_id,omitempty" bson:"_id,omitempty"`
	Task		string				`json:"task,omitempty"`
	Status		bool				`json:"status,omitempty"`
	CreatedAt 	*time.Time			`json:"createdAt,omitempty" bson:"createdAt"`
	LastUpdate	*time.Time			`json:"lastUpdate,omitempty" bson:"last_update"`
}



func GetAllTask() ([]*ToDo, error) {
	var todos []*ToDo

	collection := db.Collection
	cur, err := collection.Find(context.Background(), bson.D{{}})
	if err != nil {
		
		return  todos, err
	}
	
	for cur.Next(context.Background()) {
		var todo ToDo
		err := cur.Decode(&todo)
		if err != nil {
			return todos, err
		}
		todos = append(todos, &todo)
	}

	return todos, nil
}

func InsertTask(task *ToDo) error {
	collection := db.Collection
	_, err := collection.InsertOne(context.Background(), task)
	if err != nil {
		return err
	}

	return nil
}

func TaskComplete(id string) error {
	collection := db.Collection
	filter := bson.M{"_id":id}
	update := bson.M{"$set": bson.M{"status": true}}
	_, err := collection.UpdateOne(context.Background(), filter, update)
	if err != nil {
		return err
	}

	return nil

}

func UndoTask(id string) error {
	collection := db.Collection
	filter := bson.M{"_id":id}
	update := bson.M{"$set": bson.M{"status": false}}
	_, err := collection.UpdateOne(context.Background(), filter, update)
	if err != nil {
		return err
	}

	return nil

}

func DeleteTask(id string) error {
	collection := db.Collection
	filter := bson.M{"_id":id}
	_, err := collection.DeleteOne(context.Background(), filter)
	if err != nil {
		return err
	}

	return nil
}


func DeleteAllTask() error {
	collection := db.Collection
	_, err := collection.DeleteMany(context.Background(), bson.D{{}}, nil)
	if err != nil {
		return err
	}

	return nil
}
