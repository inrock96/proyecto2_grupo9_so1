package main

import (
	"fmt"
	"log"
	"time"
	"context"
	adminDatos "../mcache"
	"google.golang.org/grpc"

	"io/ioutil"
	"net/http"
	"encoding/json"
	"strconv"
	str "strings"

)

type datosRecibidos struct{
	Name string `json:"name"`
	Location string `json:"location"`
	Age int32 `json:"age"`
	InfectedType string `json:"infected_type"`
	State string `json:"state"`
}

type Resultado struct{
	msg string
	msg2 string
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", datos_post)
	log.Printf("listening on port 9000")
	log.Fatal(http.ListenAndServe(":9000", mux))	
}

func datos_post(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		//var results []string
		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			fmt.Println("Error reading request body", http.StatusInternalServerError)
		}
		//results = append(results, string(body))
		data := datosRecibidos{}
		err = json.Unmarshal(body, &data)
		if err!=nil{
			fmt.Println("Error al convertir los datos: %s", err)
			w.Write([]byte(err.Error()))
		}
		fmt.Println(data.Name)

		rr, errGRPC:=sendtoGrpc(
			normalizar(data.Name),
			normalizar(data.Location), 
			data.Age,
			normalizar(data.InfectedType), 
			normalizar(data.State))

		if errGRPC!=nil{
			w.Write([]byte(errGRPC.Error()))
		}else{
			if rr.msg=="true"{
				w.Write([]byte("Enviado -> "+rr.msg2))
			}else{
				w.Write([]byte("No se pudo enviar"))
			}			
		}


		//fmt.Println("-------------------------------------------")
	} else {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
	}
}

func sendtoGrpc(nombre string, depto string, edad int32, fc string, est string) (*Resultado, error) {
	time.Sleep(2000 * time.Millisecond)
	var conn *grpc.ClientConn
	//conn, err := grpc.Dial("cnt-python-svc:50051", grpc.WithInsecure())
	conn, err := grpc.Dial("servidorgo-sopes1:9090", grpc.WithInsecure())
	if err != nil {
		fmt.Println("No se puedo conectar: %s", err)
		return nil,err;
	}
	defer conn.Close()

	c := adminDatos.NewCacheClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	response, err2 := c.ObtenerDatos(ctx, &adminDatos.Datos{
		Name:nombre,
		Location:depto,
		Age:edad,
		InfectedType:fc,
		State:est,
	})

	if err2 != nil {
		fmt.Println("Error: %s", err2)
		return nil,err2;
	}
	
	
	fmt.Println("Response from server: ", response)
	var rr = Resultado{}
	rr.msg=strconv.FormatBool(response.Enviado)
	rr.msg2=response.Msg
	return &rr,nil;

}


func normalizar(txt string) string{
	txt=str.Replace(txt, "á", "a", -1);
	txt=str.Replace(txt, "é", "e", -1);
	txt=str.Replace(txt, "í", "i", -1);
	txt=str.Replace(txt, "ó", "o", -1);
	txt=str.Replace(txt, "ú", "u", -1);
	
	txt=str.Replace(txt, "Á", "A", -1);
	txt=str.Replace(txt, "É", "E", -1);
	txt=str.Replace(txt, "Í", "I", -1);
	txt=str.Replace(txt, "Ó", "O", -1);
	txt=str.Replace(txt, "Ú", "U", -1);

	txt=str.Replace(txt, "ä", "a", -1);
	txt=str.Replace(txt, "ë", "e", -1);
	txt=str.Replace(txt, "ï", "i", -1);
	txt=str.Replace(txt, "ö", "o", -1);
	txt=str.Replace(txt, "ü", "u", -1);
	
	txt=str.Replace(txt, "Ä", "A", -1);
	txt=str.Replace(txt, "Ë", "E", -1);
	txt=str.Replace(txt, "Ï", "I", -1);
	txt=str.Replace(txt, "Ö", "O", -1);
	txt=str.Replace(txt, "Ü", "U", -1);

	txt=str.Replace(txt, "ñ", "ni", -1);
	txt=str.Replace(txt, "Ñ", "NI", -1);
	return txt;
}