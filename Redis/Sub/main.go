package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"strconv"
	"time"

	"github.com/go-redis/redis"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type caso struct {
	Name         string `json:"name"`
	Location     string `json:"location"`
	Age          int    `json:"age"`
	InfectedType string `json:"infected_type"`
	State        string `json:"state"`
}

var ipAddress = "34.123.181.161:6379"
var ctx = context.Background()

func main() {
	client := redis.NewClient(&redis.Options{
		Addr:     ipAddress,
		Password: "Sopes1Grupo9",
		DB:       0,
	})
	for {
		time.Sleep(200 * time.Millisecond)
		pubsub := client.PSubscribe("casos")

		msg, err := pubsub.ReceiveMessage()
		if err != nil {
			fmt.Println(err)
			continue
		}
		fmt.Println("casos", msg.Payload)
		var nuevo caso
		datos := []byte(msg.Payload)
		err = json.Unmarshal(datos, &nuevo)
		//TODO enviar datos recibidos a redis
		pivote1, err := (client.Get("CONTADOR").Result())
		if err != nil {
			fmt.Println("No hay contador")
			continue
		}
		pivote, err2 := strconv.Atoi(pivote1)
		if err2 != nil {
		}
		client.HSet("PACIENTES", "nombre["+pivote1+"]", nuevo.Name)
		client.HSet("PACIENTES", "departamento["+pivote1+"]", nuevo.Location)
		client.HSet("PACIENTES", "edad["+pivote1+"]", nuevo.Age)
		client.HSet("PACIENTES", "forma de contagio["+pivote1+"]", nuevo.InfectedType)
		client.HSet("PACIENTES", "estado["+pivote1+"]", nuevo.State)
		pivateInt := int(pivote) + 1
		client.Set("CONTADOR", pivateInt, 0)

		//TODO enviar datos recibidos a mongo

		credential := options.Credential{
			Username: "AdminSopes1",
			Password: "Sopes1Grupo9",
		}
		clientOptions := options.Client().ApplyURI("mongodb://34.123.181.161:27017").SetAuth(credential)
		//clientOptions := options.Client().ApplyURI("mongodb+srv://AdminSopes1:Sopes1Grupo9@cluster0.p71sd.mongodb.net/CORONAVIRUS?retryWrites=true&w=majority")
		mongoClient, err := mongo.Connect(context.TODO(), clientOptions)
		if err != nil {
			panic(err)
		}
		collection := mongoClient.Database("CORONAVIRUS").Collection("PACIENTES")
		insertResult, err := collection.InsertOne(context.TODO(), nuevo)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("Inserted a Single Document: ", insertResult.InsertedID)
	}
}
