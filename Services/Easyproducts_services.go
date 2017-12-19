package Services

import (
	Db "Easyshopping/Common/DB/Mysql"
	Model "Easyshopping/Model"
	"encoding/json"
	"log"
	"math/rand"
	"net/http"
	"strconv"
	"time"
)

func Vegetables(w http.ResponseWriter, r *http.Request) {
	Data := Db.Getveg()
	Senddata, err := json.Marshal(Data)
	if err != nil {
		log.Println("Error - RETRIVE PRODUCT DATA", err)
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Access-Control-Allow-Orgin", "*")
	w.Write(Senddata)
}
func Leaves(w http.ResponseWriter, r *http.Request) {
	Data := Db.GetLeaves()
	Senddata, err := json.Marshal(Data)
	if err != nil {
		log.Println("Error - RETRIVE PRODUCT DATA", err)
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Access-Control-Allow-Orgin", "*")
	w.Write(Senddata)
}
func NonVeg(w http.ResponseWriter, r *http.Request) {
	Data := Db.GetNonveg()
	Senddata, err := json.Marshal(Data)
	if err != nil {
		log.Println("Error - RETRIVE PRODUCT DATA", err)
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Access-Control-Allow-Orgin", "*")
	w.Write(Senddata)
}
func Fruits(w http.ResponseWriter, r *http.Request) {
	Data := Db.GetFruits()
	Senddata, err := json.Marshal(Data)
	if err != nil {
		log.Println("Error - RETRIVE PRODUCT DATA", err)
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Access-Control-Allow-Orgin", "*")
	w.Write(Senddata)
}
func SearchProduct(w http.ResponseWriter, r *http.Request) {
	var Searchproduct Model.Search
	err := json.NewDecoder(r.Body).Decode(&Searchproduct)
	if err != nil {
		log.Println("Error - SEARCH PRODUCT", err)
	}

	Data := Db.Search_DB(Searchproduct.Searchproduct)
	Senddata, err := json.Marshal(Data)
	if err != nil {
		log.Println("Error - SEARCH PRODUCT RESPONSE", err)
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Access-Control-Allow-Orgin", "*")
	w.Write(Senddata)
}

func Orderplaced(w http.ResponseWriter, r *http.Request) {
	var Orderplaced Model.Orderplaced
	err := json.NewDecoder(r.Body).Decode(&Orderplaced)
	if err != nil {
		log.Println("Error - ORDER PLACED", err)
	}
	rand.Seed(time.Now().UnixNano())
	randomno := randomInt(1234567, 7654321)
	uniqueid := "E3BILL" + strconv.Itoa(randomno)
	Data := Db.Orderplaced_DB(Orderplaced, uniqueid)
	Senddata, err := json.Marshal(Data)
	if err != nil {
		log.Println("Error - ORDER PLACED RESPONSE", err)
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Access-Control-Allow-Orgin", "*")
	w.Write(Senddata)
}

func Ordertracking(w http.ResponseWriter, r *http.Request) {
	var Track Model.Ordertrack
	err := json.NewDecoder(r.Body).Decode(&Track)
	if err != nil {
		log.Println("Error - ORDER TRACKING", err)
	}

	Data := Db.OrderTracking_DB(Track.Orderid)
	Senddata, err := json.Marshal(Data)
	if err != nil {
		log.Println("Error - ORDER PLACED RESPONSE", err)
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Access-Control-Allow-Orgin", "*")
	w.Write(Senddata)
}
func Orderhistory(w http.ResponseWriter, r *http.Request) {

	var Orderid Model.ProfileLoginid
	err := json.NewDecoder(r.Body).Decode(&Orderid)
	if err != nil {
		log.Println("Error - ORDER HISTORY", err)
	}

	Data := Db.OrderHistory_DB(Orderid.Loginid)
	Senddata, err := json.Marshal(Data)
	if err != nil {
		log.Println("Error - ORDER HISTORY RESPONSE", err)
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Access-Control-Allow-Orgin", "*")
	w.Write(Senddata)
}
