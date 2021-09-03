package bd

import (
	"context"
	"time"

	"github.com/mclonge/piando/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var db = MongoCN.Database("PIANDODB")
var collect = db.Collection("usuarios")

func AddUser(u models.Usuario) (string, bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second) //contexto dentro del grande el TODO

	defer cancel() //si no hago esto voy sumando contextosdentro del grande

	u.Pass, _ = EncriptarPass(u.Pass)

	result, err := collect.InsertOne(ctx, u) //InsertOne Instruccion Mongo para insertar

	if err != nil {
		return "", false, err
	}

	ObjID, _ := result.InsertedID.(primitive.ObjectID)

	return ObjID.String(), true, nil
}

func CheckUserExist(email string) (models.Usuario, bool, string) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second) //contexto dentro del grande el TODO
	defer cancel()

	condicion := bson.M{"email": email}

	var resultado models.Usuario //variable resultado de tipo models.Usuario

	err := collect.FindOne(ctx, condicion).Decode(&resultado)
	ID := resultado.ID.Hex()
	if err != nil {
		return resultado, false, ID
	}
	return resultado, true, ID
}
