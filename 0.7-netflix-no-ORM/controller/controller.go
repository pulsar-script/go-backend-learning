package controller

import (
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	models "main.go/model"
)

// DB connections

const connectionString = "mongodb://localhost:27017/"
const dbName = "netflix"
const colName = "watchlist"

// * MOST IMP
var collection *mongo.Collection

// connect with mongoDB
func init() {

	//client option
	clientOption := options.Client().ApplyURI(connectionString)

	// connect to mongoDB
	client, err := mongo.Connect(context.TODO(), clientOption)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("MongoDB connections success")

	collection = client.Database(dbName).Collection(colName)

	// collection instance
	fmt.Println("Collection insatance is ready")
}

// MONGODB Helper functions - file

// Insert One Movie

// * here our helper function is not started with Capital letter, beacuse these helper functions are specifices to MONGODB, That's why we dont want to use out of this package
func insertOneMovie(movie models.Netflix) { //* that movie is coming from models package (models file)

	insertedMovie, err := collection.InsertOne(context.Background(), movie)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Insterted One Movie in DB : ", insertedMovie)
}

// Update One Movie

func updateOneMovie(movieId string) {

	// we get movieId in string , but mongoDB only store and understand data in BSON
	_id, err := primitive.ObjectIDFromHex(movieId)

	if err != nil {
		log.Fatal(err)
	}

	// As you know we are using directly mongoDB with any ORM
	filter := bson.M{"_id": _id}
	update := bson.M{"$set": bson.M{"watched": true}}

	// update document
	result, err := collection.UpdateOne(context.Background(), filter, update)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("modified movie result: ", result)
}

// Delete one Movie
func deleteOneMovie(movieId string) {

	_id, err := primitive.ObjectIDFromHex(movieId)

	if err != nil {
		log.Fatal(err)
	}

	filter := bson.M{"_id": _id}

	deletedOne, err := collection.DeleteOne(context.Background(), filter)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Delete One movie : ", deletedOne)
}

// Delete all Movies
func deleteAllMovies() *mongo.DeleteResult {

	// filter := bson.D{{}}  //* Senior Tip = insted direcly put it
	deletedAll, err := collection.DeleteMany(context.Background(), bson.D{{}}, nil) // putting "nil" , see here we can put other options , but when you have no options then you can put "nil"

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Deleted All Movies : ", deletedAll)
	return deletedAll
}

// Get all Movies

func getAllMovies() []primitive.M {

	// here mongoDB does not return all movies directly
	cursor, err := collection.Find(context.Background(), bson.D{{}})

	if err != nil {
		log.Fatal(err)
	}

	//TEST
	fmt.Printf(" how cursor look in getAllMovies ? : %v ", cursor)

	// var to store our movies from cursor

	var movies []primitive.M //* primitive.M is basically bson.M , hitesh said bson.M crate errors

	for cursor.Next(context.Background()) {
		var movie bson.M
		err := cursor.Decode(&movie)

		if err != nil {
			log.Fatal(err)
		}

		movies = append(movies, movie)
	}

	defer cursor.Close(context.Background())
	return movies
}

// Controllers
