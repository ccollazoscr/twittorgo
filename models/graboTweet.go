package models

import "time"

/*GraboTweet tiene los datos neecesario para guardar un twet*/
type GraboTweet struct {
	UserID  string    `bson: "userid" json:"userid,omitempty"`
	Mensaje string    `bson: "mensaje" json:"mensaje,omitempty"`
	Fecha   time.Time `bson: "fecha" json:"fecha,omitempty"`
}
