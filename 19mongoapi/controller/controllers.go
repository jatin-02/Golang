package controller

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jatin-02/mongoapi/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const connectionString = ""
const dbName = "netflix"
const colName = "watchlist"

var collection *mongo.Collection

func init() {
	// client option
	clientOption := options.Client().ApplyURI(connectionString)

	// connect to mongodb
	client, err := mongo.Connect(context.TODO(), clientOption)
	checkNilError(err)
	fmt.Println("MongoDB connection is succesful")
	
	collection = client.Database(dbName).Collection(colName)
	fmt.Println("Collection instance is ready")
}

func insertOneMovie(movie model.Netflix)  {
	inserted, err := collection.InsertOne(context.Background(), movie)
	checkNilError(err)
	fmt.Println("inserted one movie in db with id: ", inserted.InsertedID)
}

func updateOneMovie(movieId string)  {
	id, err := primitive.ObjectIDFromHex(movieId)
	checkNilError(err)

	filter := bson.M{"_id":id}
	update := bson.M{"$set": bson.M{"watched":true}}

	result, err := collection.UpdateOne(context.Background(), filter, update)
	fmt.Println("modified count", result.ModifiedCount)
}

func deleteOneMovie(movieId string)  {
	id, err := primitive.ObjectIDFromHex(movieId)
	checkNilError(err)

	filter := bson.M{"_id":id}
	deleteCount, err := collection.DeleteOne(context.Background(), filter)
	checkNilError(err)
	fmt.Println("Delete count: ", deleteCount.DeletedCount)
}

func deleteAllMovies() int64 {
	deleteResult, err := collection.DeleteMany(context.Background(), bson.D{{}}, nil)
	checkNilError(err)
	return deleteResult.DeletedCount
}

func getAllMovies() []bson.M {
	cur, err := collection.Find(context.Background(), bson.D{{}})
	checkNilError(err)

	var movies []bson.M
	for cur.Next(context.Background()) {
		var movie bson.M
		err := cur.Decode(&movie)
		checkNilError(err)
		movies = append(movies, movie)
	}

	defer cur.Close(context.Background())
	return movies
}

// Actual Controller file

func GetAllMovies(w http.ResponseWriter, r *http.Request)  {
	w.Header().Set("Content-Type", "application/x-www-form-urlencode")
	allMovies := getAllMovies()
	json.NewEncoder(w).Encode(allMovies)
}

func CreateMovie(w http.ResponseWriter, r *http.Request)  {
	w.Header().Set("Content-Type", "application/x-www-form-urlencode")
	w.Header().Set("Allow-Control-Allow-Methods", "POST")

	var movie model.Netflix
	_ = json.NewDecoder(r.Body).Decode(&movie)
	insertOneMovie(movie)
	json.NewEncoder(w).Encode(movie)
}

func MarkAsWatched(w http.ResponseWriter, r *http.Request)  {
	w.Header().Set("Content-Type", "application/x-www-form-urlencode")
	w.Header().Set("Allow-Control-Allow-Methods", "POST")

	params := mux.Vars(r)
	updateOneMovie(params["id"])
	json.NewEncoder(w).Encode(params["id"])
}

func DeleteOneMovie(w http.ResponseWriter, r *http.Request)  {
	w.Header().Set("Content-Type", "application/x-www-form-urlencode")
	w.Header().Set("Allow-Control-Allow-Methods", "DELETE")

	params := mux.Vars(r)
	deleteOneMovie(params["id"])
	json.NewEncoder(w).Encode(params["id"])
}

func DeleteAllMovies(w http.ResponseWriter, r *http.Request)  {
	w.Header().Set("Content-Type", "application/x-www-form-urlencode")
	w.Header().Set("Allow-Control-Allow-Methods", "DELETE")

	count := deleteAllMovies()
	json.NewEncoder(w).Encode(count)
}

func checkNilError(err error)  {
	if err != nil {
		log.Fatal(err)
	}
}