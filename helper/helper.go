package helper

import (
	"ASimpleProgram/model"
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const connectionString = "mongodb://localhost:27017"
const colName = "Employee"
const dbName = "EmployeeRepo"

var collection *mongo.Collection

func init() {
	clientOptions := options.Client().ApplyURI(connectionString)
	client, err := mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		panic(err)
	}

	collection = client.Database(dbName).Collection(colName)
	fmt.Println("MongoDB ready")
}

func InsertOneRecord(employee model.Employee) string {
	inserted, err := collection.InsertOne(context.Background(), employee)
	if err != nil {
		panic(err)
	}
	fmt.Println("Inserted:", inserted)
	return "Inserted"
}

func GetOneRecord(id string) *model.Employee {
	var emp model.Employee
	empId, _ := primitive.ObjectIDFromHex(id)
	filter := bson.M{"_id": empId}
	collection.FindOne(context.Background(), filter).Decode(&emp)
	fmt.Println("Got employe", emp)
	return &emp
}

func GetAllRecords() []bson.M {
	filter := bson.D{{}}
	cursor, err := collection.Find(context.Background(), filter)
	if err != nil {
		panic(err)
	}
	var empArray []bson.M
	for cursor.Next(context.Background()) {
		var emp bson.M
		err := cursor.Decode(&emp)
		if err != nil {
			panic(err)
		}
		empArray = append(empArray, emp)
	}
	return empArray

}

func DeleteRecord(id string) string {
	empId, _ := primitive.ObjectIDFromHex(id)
	filter := bson.M{"_id": empId}
	deleteres, err := collection.DeleteOne(context.Background(), filter)
	if err != nil {
		panic(err)
	}
	fmt.Println("Deleted:", deleteres)
	return "deleted"
}

func UpdateRecord(emp model.Employee) string {
	empId, _ := primitive.ObjectIDFromHex(emp.ID.Hex())
	filter := bson.M{"_id": empId}
	update := bson.M{"$set": bson.M{"name": emp.Name, "company": emp.Company}}
	updateres, err := collection.UpdateOne(context.Background(), filter, update)
	if err != nil {
		panic(err)
	}
	fmt.Println("Updated", updateres)
	return "Updated"
}

func DeleteAllRecords() string {
	filter := bson.D{{}}
	res, err := collection.DeleteMany(context.Background(), filter)
	if err != nil {
		panic(err)
	}
	fmt.Println(res)
	return "Deleted all"
}
