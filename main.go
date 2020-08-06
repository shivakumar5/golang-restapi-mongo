package main

import (
	"encoding/json"

	"fmt"
	"log"
	"net/http"

	"github.com/flexera/avihs/mongo"
	"github.com/gorilla/mux"	
)

// Employee ...
type Employee struct{
	ID string `bson:"_id", json:"id"`
	Name string `bson:"name", json :"name"`
	Age int `bson:"age", json :"age"`
}

type mongosvc struct{
	mongo mongo.Client
}

// GetConnection ...
func GetConnection() mongo.Client {

	conn, err := mongo.NewClient([]string{"localhost:27017"}, "", "", "", 123456)
	if err != nil{
		fmt.Printf("failed to connect mongodb %s", err)
	}

	return conn
}



// CreateEmployee ...
func CreateEmployee(w http.ResponseWriter, r *http.Request){

	conn := GetConnection()

	var employee mongo.Employee
	json.NewDecoder(r.Body).Decode(&employee)

	err := conn.CreateEmployee(employee)
	if err != nil{
		w.WriteHeader(http.StatusInternalServerError)
	}

	w.WriteHeader(http.StatusCreated)

}

// GetAllEmployees ...
func GetAllEmployees(w http.ResponseWriter, r *http.Request){
	conn := GetConnection()

	res, err := conn.GetEmployees()
	if err != nil{
		w.WriteHeader(http.StatusNoContent)
	}

	json.NewEncoder(w).Encode(res)
}

// GetSingleEmployee ...
func GetSingleEmployee(w http.ResponseWriter, r *http.Request){

	key := mux.Vars(r)["id"]

	conn := GetConnection()

	res, err := conn.GetEmployeeByID(key)	
	
	
	if err != nil{
		w.WriteHeader(http.StatusNoContent)
	}
	

	json.NewEncoder(w).Encode(res)
}

// UpdateEmployee ...
func UpdateEmployee(w http.ResponseWriter, r *http.Request){

	key := mux.Vars(r)["id"]

	var updateEmployee mongo.Employee
	json.NewDecoder(r.Body).Decode(&updateEmployee)
	
	conn := GetConnection()

	err := conn.Update(key, updateEmployee)
	if err != nil{
		w.WriteHeader(http.StatusNoContent)
	}

	w.WriteHeader(http.StatusOK)
	
}


// DeleteEmployee ...
func DeleteEmployee(w http.ResponseWriter, r *http.Request){
	
 	key := mux.Vars(r)["id"]
	
	conn := GetConnection()

	err := conn.Delete(key)
	if err != nil{
		w.WriteHeader(http.StatusNoContent)
	}
}



func main(){

	newRouter := mux.NewRouter()

	newRouter.HandleFunc("/CreateEmployee", CreateEmployee).Methods("POST")
	newRouter.HandleFunc("/GetAllEmployees", GetAllEmployees)
	newRouter.HandleFunc("/GetAllEmployees/{id}", GetSingleEmployee)
	newRouter.HandleFunc("/DeleteEmployee/{id}", DeleteEmployee).Methods("DELETE")
	newRouter.HandleFunc("/UpdateEmployee/{id}", UpdateEmployee).Methods("PUT")

	log.Fatal(http.ListenAndServe(":8009", newRouter))

} //end of main
