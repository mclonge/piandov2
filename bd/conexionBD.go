package bd

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

/*MongoCN Objeto de conexión a la BBDD */
var MongoCN = ConectarBD()
var clientOptions = options.Client().ApplyURI("mongodb+srv://<user>:<passs>@cluster0.bsg7r.mongodb.net/myFirstDatabase?retryWrites=true&w=majority")

/*ConectarBD es la función que realiza la conexión con Mongo*/
func ConectarBD() *mongo.Client {

	client, err := mongo.Connect(context.TODO(), clientOptions)

	if err != nil {
		log.Fatal(err.Error())
		return client
	}

	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal(err.Error())
		return client

	}
	log.Println("Conexión exitosa con la BD")
	return client
}

/*CheckConnection un ping a BD*/

func CheckConnection() int {
	err := MongoCN.Ping(context.TODO(), nil)
	if err != nil {
		return 0
	}
	return 1
}
