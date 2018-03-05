package main

import (
  "encoding/json"
  "log"
  "net/http"
  "time"
  "github.com/gorilla/mux"
  "github.com/heptiolabs/healthcheck"
)

type Person struct {
  ID        string    `json:"id,omitempty"`
  FirstName string    `json:"firstname,omitempty"`
  LastName  string    `json:"lastname,omitempty"`
  Address   *Address  `json:"address,omitempty"`
}

type Address struct {
  City      string    `json:"city,omitempty"`
  State     string    `json:"state,omitempty"`
}

var people []Person

// List data for all people
func GetPeople(w http.ResponseWriter, r *http.Request) {
  json.NewEncoder(w).Encode(people)
}
// Display data for a person
func GetPerson(w http.ResponseWriter, r *http.Request) {
  params := mux.Vars(r)
  for _, item := range people {
    if item.ID == params["id"] {
      json.NewEncoder(w).Encode(item)
      return
    }
  }
  json.NewEncoder(w).Encode(&Person{})
}
// Create a person
func CreatePerson(w http.ResponseWriter, r *http.Request) {
  params := mux.Vars(r)
  var person Person
  _ = json.NewDecoder(r.Body).Decode(&person)
  person.ID = params["id"]
  people = append(people, person)
  json.NewEncoder(w).Encode(people)
}
// Delete a person from the list
func DeletePerson(w http.ResponseWriter, r *http.Request) {
  params := mux.Vars(r)
  for index, item := range people {
    if item.ID == params["id"] {
      people = append(people[:index], people[index+1:]...)
      break
    }
  }
  json.NewEncoder(w).Encode(people)
}

func main() {

  // Manually adding records
  people = append(people, Person{ID:"1", FirstName:"John", LastName:"Doe", Address: &Address{City: "Austin", State: "Texas"}})
  people = append(people, Person{ID:"2", FirstName:"Jane", LastName:"Doe", Address: &Address{City: "Houston", State: "Texas"}})
  people = append(people, Person{ID:"3", FirstName:"Francesca", LastName:"Rhodes"})

  // Add healthchecks
  health := healthcheck.NewHandler()
  // Our app is not happy if we've got more than 100 goroutines running.
  health.AddLivenessCheck("goroutine-threshold", healthcheck.GoroutineCountCheck(100))
  // Our app is not ready if we can't resolve our upstream dependency in DNS.
  health.AddReadinessCheck(
      "upstream-google",
      healthcheck.DNSResolveCheck("google.com", 50*time.Millisecond))

  router := mux.NewRouter()
  router.HandleFunc("/people", GetPeople).Methods("GET")
  router.HandleFunc("/people/{id}", GetPerson).Methods("GET")
  router.HandleFunc("/people/{id}", CreatePerson).Methods("POST")
  router.HandleFunc("/people/{id}", DeletePerson).Methods("DELETE")
  router.HandleFunc("/live", health.LiveEndpoint)
  router.HandleFunc("/ready", health.ReadyEndpoint)
  log.Fatal(http.ListenAndServe(":8080", router))
}
