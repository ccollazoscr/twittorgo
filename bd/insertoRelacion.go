package bd

import (
	"context"
	"time"

	"github.com/ccollazoscr/twittorgo/models"
)

/*InsertoRelacion inserta una relación entre dos usuarios*/
func InsertoRelacion(t models.Relacion) (bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	db := MongoCN.Database("twittor")
	col := db.Collection("relacion")

	_, err := col.InsertOne(ctx, t)
	if err != nil {
		return false, err
	}
	return true, nil
}
