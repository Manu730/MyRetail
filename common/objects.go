package common

import (
	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
)

//Hosting server info
//No tls as this is just plain http
type Server struct {
	Name         string
	Address      string
	Router       *mux.Router
	readtimeout  int
	writetimeout int
}

type Product struct {
	ID    uint64 `json:"id"`
	Name  string `json:"name"`
	Price Cost   `json:"current_price"`
}

type Cost struct {
	Value    float32 `json:"value"`
	Currency string  `json:"currency_code"`
}

//obj to db obj conversion
func (p Product) Getbson() bson.M {
	return bson.M{"ID": p.ID, "Name": p.Name, "Price": p.Price}
}

func (c Cost) Getbson() bson.M {
	return bson.M{"Value": c.Value, "Currency": c.Currency}
}
