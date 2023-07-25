package middleware

//middleware package serves as the bridge between APIs and the database, handling all crud operation
import (
	"encoding/json"
	"errors"
	"fmt"
	"go-shopping-cart/models"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// Response := the data or information that is returned from server when an API request is sent
type Response struct {
	ID      int64  `json:"id,omitempty"`
	Message string `json:"message,omitempty"`
}

var Database *gorm.DB

func Database_connection() *gorm.DB {
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatal("Error in loading .env file")
	}

	host := os.Getenv("DB_HOST")
	username := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	database_name := os.Getenv("DB_NAME")
	port, err := strconv.Atoi(os.Getenv("DB_PORT"))

	if err != nil {
		log.Fatal()
	}

	psql_info := fmt.Sprintf("host=%s port=%d user=%s dbname=%s sslmode=disable password=%s", host, port, username, database_name, password)

	Database, err := gorm.Open(postgres.Open(psql_info), &gorm.Config{})
	if err != nil {
		fmt.Println(err)
		log.Fatal()
	}

	return Database
}

// Api Endpoint Handlers
func CreateProduct(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Access-Control-Allow-origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	var Product models.Products
	// decoding a json request -> process of extracting the data sent in the body of an HTTP req

	err := json.NewDecoder(r.Body).Decode(&Product) //Body bata leo Product ko data

	if err != nil {
		log.Fatalf("Unable to decode the request body. %v", err)
	}

	insertID := insertProduct(Product)

	// format a response object
	res := Response{
		ID:      insertID,
		Message: "User created Successfully",
	}

	json.NewEncoder(w).Encode(res) //writes the response by encoding

}

//get user

func GetProduct(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/x-www-form-urlencoded")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	//get the Product id from request params, key is "id"
	params := mux.Vars(r)

	id, err := strconv.Atoi(params["id"]) //convert the id type from string to int
	if err != nil {
		log.Fatalf("Unable to convert the string to int . %v", err)

	}

	//call getProduct func with user id ot retrieve a single user
	Product, err := getProduct(int64(id))
	if err != nil {
		log.Fatalf("unable to get user. %v", err)
	}

	json.NewEncoder(w).Encode(Product)
}

// "id" parameter is used when retrieving a single user in the 'GetUser' fn, to fetch a specific user based on the provided id
// fetch := action of retrieving or getting the desired data from database
func GetAllProduct(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/x-www-form-urlencoded")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	Products, err := getAllProduct()

	if err != nil {
		log.Fatalf("unable to get all the Product. %v", err)
	}

	json.NewEncoder(w).Encode(Products)

}

func UpdateProduct(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/x-www-form-urlencoded")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "PUT")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	//get the user id from req params, key "id"

	var Product models.Products
	err := json.NewDecoder(r.Body).Decode(&Product) //request lai decode

	if err != nil {
		log.Fatalf("Unable to decode request body. %v", err)
	}
	fmt.Println(Product.ID)
	updatedRows := updateProduct(Product.ID, Product)

	// message
	msg := fmt.Sprintf("Product updated successfully. Total rows affected %v", updatedRows)

	res := Response{
		ID:      Product.ID,
		Message: msg,
	}

	json.NewEncoder(w).Encode(res)

}

func DeleteProduct(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/x-www-form-urlencoded")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "PUT")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	//get the userId from the req params "id"

	params := mux.Vars(r)

	id, err := strconv.Atoi(params["id"])
	fmt.Println(id)

	if err != nil {
		log.Fatalf("Unable to convert the string into int. %v", err)
	}

	deletedRows := deleteProduct(int64(id))
	msg := fmt.Sprintf("User updated successfully. Total rows affected: %v", deletedRows)

	res := Response{
		ID:      int64(id),
		Message: msg,
	}
	json.NewEncoder(w).Encode(res)
}

// insert one user in db
func insertProduct(Product models.Products) int64 {
	db := Database_connection()
	db.AutoMigrate(&models.Products{})
	result := db.Create(&Product)
	if result.Error != nil {
		panic(fmt.Sprintf("Failed to execute the query: %v", result.Error))
	}
	fmt.Printf("Inserted a single record %v \n", Product.ID)
	return Product.ID
}

func getProduct(id int64) (models.Products, error) {
	db := Database_connection()

	var Product models.Products
	//finding Product by id
	result := db.First(&Product, id)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		fmt.Println("no rows were returned")
		return Product, nil
	} else if result.Error != nil {
		log.Fatalf("unable to query the user. %v", result.Error)
		return Product, result.Error
	}

	return Product, nil

}

func getAllProduct() ([]models.Products, error) {
	db := Database_connection()
	var Products []models.Products
	//retrieve all Products from db
	result := db.Find(&Products)
	if result.Error != nil {
		log.Fatalf("unable to find Products. %v", result.Error)
	}
	return Products, nil
}

func updateProduct(id int64, Product models.Products) int64 {
	db := Database_connection()
	result := db.Model(&models.Products{}).Where("id = ?", id).Updates(Product)
	if result.Error != nil {
		log.Fatalf("Unable to update Product: %v", result.Error)
	}
	rowsAffected := result.RowsAffected
	log.Printf("total rows affected: %d", rowsAffected)
	return rowsAffected
}

func deleteProduct(id int64) int64 {
	db := Database_connection()
	result := db.Delete(&models.Products{}, id)
	if result.Error != nil {
		log.Fatalf("the record is not deleted")
	}
	rowsAffected := result.RowsAffected
	return rowsAffected
}
