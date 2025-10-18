package controllers

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	models "main.go/models"
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

	fmt.Println("\nMongoDB connections success")

	collection = client.Database(dbName).Collection(colName)

	// collection instance
	fmt.Println("\nCollection instance is ready")
}

// MONGODB Helper functions - file

// Insert One Movie

// * here our helper function is not started with Capital letter, beacuse these helper functions are specifices to MONGODB, That's why we dont want to use out of this package
func insertOneMovie(movie models.Netflix) *mongo.InsertOneResult { //* that movie is coming from models package (models file)

	// generate Object ID before inserting into DB
	movie.ID = primitive.NewObjectID()

	insertedMovieId, err := collection.InsertOne(context.Background(), movie)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("\nInsterted One Movie in DB : ", insertedMovieId.InsertedID)

	return insertedMovieId
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
	update := bson.M{"$set": bson.M{"isWatched": true}}

	// update document
	result, err := collection.UpdateOne(context.Background(), filter, update)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("\nmodified movie result: ", result)
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

	fmt.Println("\nDelete One movie : ", deletedOne)
}

// Delete all Movies
func deleteAllMovies() *mongo.DeleteResult {

	// filter := bson.D{{}}  //* Senior Tip = insted direcly put it
	deletedAll, err := collection.DeleteMany(context.Background(), bson.D{{}}, nil) // putting "nil" , see here we can put other options , but when you have no options then you can put "nil"

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("\nDeleted All Movies : ", deletedAll)
	return deletedAll
}

// Get all Movies

func getAllMovies() []models.Netflix {

	// here mongoDB does not return all movies directly
	cursor, err := collection.Find(context.Background(), bson.D{{}})

	if err != nil {
		log.Fatal(err)
	}

	//TEST
	fmt.Printf("\nhow cursor look in getAllMovies ? : %v ", cursor)

	// var to store our movies from cursor

	var movies []models.Netflix // correct version

	//!	var movies []primitive.M //* primitive.M is basically bson.M , hitesh said bson.M crate errors
	//* Decoding DB response into Mongo document (map) insted of go struct is not recommended, it bypass JOSN tags on struct and send DB keys
	// go and read Note "JSON & BOSN tags" in Notion about it

	for cursor.Next(context.Background()) {

		//! var movie primitive.M
		var movie models.Netflix

		err := cursor.Decode(&movie)

		if err != nil {
			log.Fatal(err)
		}

		movies = append(movies, movie)
	}

	defer cursor.Close(context.Background())
	return movies
}

// --- helper functions end here ---

// Controllers - file

// Get all movies
func GetMyAllMovies(w http.ResponseWriter, r *http.Request) {

	// set headers to tell which kind a response you are sending
	w.Header().Set("Content-Type", "application/x-www-for-urlencode")

	allMovies := getAllMovies()

	// - here we can process , refine , tune , twist your data for frontend

	json.NewEncoder(w).Encode(allMovies)
	return
}

// Create a New Movie

func CreateNewMovie(w http.ResponseWriter, r *http.Request) {

	// set headers
	w.Header().Set("Content-Type", "application/x-www-form-urlencode")
	w.Header().Set("Allow-Control-Allow-Methods", "POST")

	// get new movie details from r - store it into local variable
	var newMovie models.Netflix
	err := json.NewDecoder(r.Body).Decode(&newMovie)

	if err != nil {
		log.Fatal(err)
	}

	// store in into DB
	newInsertedMovieId := insertOneMovie(newMovie)

	// send response
	json.NewEncoder(w).Encode(newInsertedMovieId)
}

// Update a Movie
func MarkAsWatched(w http.ResponseWriter, r *http.Request) {

	// set headers
	w.Header().Set("Content-Type", "application/x-www-form-urlencode")
	w.Header().Set("Allow-Control-Allow-Methods", "PUT")

	// grab id from params and store it
	params := mux.Vars(r)

	// pass id to DB helper func
	updateOneMovie(params["movieId"])

	// send response
	json.NewEncoder(w).Encode("Movie marked as watched")
}

// Delete a movie
func DeleteMyOneMovie(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/x-www-form-urlencode")
	w.Header().Set("Allow-Control-Allow-Methods", "DELETE")

	params := mux.Vars(r)

	deleteOneMovie(params["movieId"])

	json.NewEncoder(w).Encode("Successfully Delete A Movie")

}

// Delet All movies
func DeleteMyAllMovies(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-type", "application/x-www-form-urlencode")
	w.Header().Set("Allow-Control-Allow-Methods", "DELETE")

	deleteAllMovies()

	json.NewEncoder(w).Encode("All Movies are successfully deleted")
}
