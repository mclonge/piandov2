package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Usuario struct { //bson: entrada a Mongo json: Salida
	//omitempty -> omites si es vacio
	ID        primitive.ObjectID `bson:"_id,omitempty" json:"id`
	Nombre    string             `bson:"nombre" json:"nombre,omitempty`
	Date      time.Time          `bson:"date" json:"date,omitempty`
	Email     string             `bson:"email" json:"email`
	Pass      string             `bson:"pass" json:"pass`
	Avatar    string             `bson:"avatar" json:"avatar,omitempty`
	Banner    string             `bson:"banner" json:"banner,omitempty`
	Biografía string             `bson:"biografia" json:"biografia,omitempty`
	Ubicacion string             `bson:"ubicacion" json:"ubicacion,omitempty`
	SitioWeb  string             `bson:"sitioweb" json:"sitioWeb,omitempty`
}
