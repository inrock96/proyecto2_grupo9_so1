package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/go-redis/redis"
	"github.com/gorilla/mux"
)

var _ctx = context.Background()

var ipAddress = "34.123.181.161:6379"

type caso struct {
	Name         string `json:"name"`
	Location     string `json:"location"`
	Age          int    `json:"age"`
	InfectedType string `json:"infected_type"`
	State        string `json:"state"`
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", inicio).Methods("GET")
	r.HandleFunc("/caso", caseHandler).Methods("POST")
	err := http.ListenAndServe(os.Getenv("PORT"), r)
	if err != nil {
		log.Fatal(err)
	}
}
func inicio(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Holi")
}
func caseHandler(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		panic(err)
	}
	var nuevo caso
	err = json.Unmarshal(body, &nuevo)
	if err != nil {
		panic(err)
	}
	b, err := json.Marshal(nuevo)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	_ctx = ctx
	client := redis.NewClient(&redis.Options{
		Addr:     ipAddress,
		Password: "Sopes1Grupo9",
		DB:       0,
	})

	client.Publish("casos", b)
}
