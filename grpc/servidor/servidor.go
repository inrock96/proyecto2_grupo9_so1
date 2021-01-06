package main

import (
	"fmt"
	"github.com/go-redis/redis"
	//"github.com/go-redis/redis/v8"
	"net"
	"log"
	"strconv"
	"context"
	pb "../mcache"
	"google.golang.org/grpc"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	
)

type recibe struct {
	Name 		string `json:"name"`
	Location 	string `json:"location"`
	Age 		int32 `json:"age"`
	InfectedType string `json:"infected_type"`
	State 	string `json:"state"`
}


const (
	port = ":9090"
)

// SayHello implements helloworld.GreeterServer
func (s *recibe) ObtenerDatos(ctx context.Context, caso *pb.Datos) (*pb.Respuesta, error) {
	
	sendtoRedis(caso.Name, caso.Location, caso.Age, caso.InfectedType, caso.State);
	sendMongoDB(caso.Name, caso.Location, caso.Age, caso.InfectedType, caso.State);
	return &pb.Respuesta{
		Enviado: true,
	}, nil
	
}

func sendtoRedis(nombre string, depto string, edad int32, tpInfectado string, estado string){ 
	var ctx = context.Background()
	client := redis.NewClient(&redis.Options{
		Addr:     "34.123.181.161:6379",
		Password: "Sopes1Grupo9",
		DB:       0,
	})

	pivote1, err := client.Get(ctx, "CONTADOR").Result();
	if err != nil {
		fmt.Println("soy peor ahora soy peor")
	}

	pivote, err2 := strconv.Atoi(pivote1);
	if err2 != nil {
	}
	fmt.Println(pivote);
	
	client.HSet(ctx, "PACIENTES", "nombre["+pivote1+"]", nombre)
	client.HSet(ctx, "PACIENTES", "departamento["+pivote1+"]", depto)
	client.HSet(ctx, "PACIENTES", "edad["+pivote1+"]", edad)
	client.HSet(ctx, "PACIENTES", "forma de contagio["+pivote1+"]", tpInfectado)
	client.HSet(ctx, "PACIENTES", "estado["+pivote1+"]", estado)
	pivoteInt := int(pivote) + 1
	client.Set(ctx, "CONTADOR", pivoteInt, 0)
	
}


func sendMongoDB(nombre string, depto string, edad int32, tpInfectado string, estado string){
	nuevo := recibe{
		Name:	nombre,
		Location:	depto,
		Age:	edad,
		InfectedType:	tpInfectado,
		State:	estado,
	}
	 fmt.Println(nuevo);
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

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterCacheServer(s, &recibe{})

	fmt.Println("conectando...");

	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}

}

