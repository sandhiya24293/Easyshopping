package Services

import (
	Db "Easyshopping/Common/DB/Mysql"
	Model "Easyshopping/Model"
	"encoding/json"
	"log"

	"net/http"
)

func Foodinsert(w http.ResponseWriter, r *http.Request) {
	var Instance Model.Fooddelivery
	err := json.NewDecoder(r.Body).Decode(&Instance)
	if err != nil {
		log.Println("Error - INSTATNT DELIVERY", err)
	}

	Data := Db.Foodinsert_DB(Instance)
	Senddata, err := json.Marshal(Data)
	if err != nil {
		log.Println("Error -  INSTATNT DELIVERY", err)
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Access-Control-Allow-Orgin", "*")
	w.Write(Senddata)
}

func Getallfooddelivery(w http.ResponseWriter, r *http.Request) {

	Data := Db.Getallfooddelivery_DB()
	Senddata, err := json.Marshal(Data)
	if err != nil {
		log.Println("Error -  INSTATNT DELIVERY", err)
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Access-Control-Allow-Orgin", "*")
	w.Write(Senddata)
}

type Foodinstant struct {
	Loginid string
}

func Fooddelieryhistory(w http.ResponseWriter, r *http.Request) {
	var Trackinstant Foodinstant
	err := json.NewDecoder(r.Body).Decode(&Trackinstant)
	if err != nil {
		log.Println("Error - INSTATNT DELIVERY", err)
	}

	Data := Db.Gethistory_DB(Trackinstant.Loginid)
	Senddata, err := json.Marshal(Data)
	if err != nil {
		log.Println("Error -  INSTATNT DELIVERY", err)
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Access-Control-Allow-Orgin", "*")
	w.Write(Senddata)
}

func Trackfooddeliver(w http.ResponseWriter, r *http.Request) {
	var Trackinstant Foodinstant
	err := json.NewDecoder(r.Body).Decode(&Trackinstant)
	if err != nil {
		log.Println("Error - INSTATNT DELIVERY", err)
	}

	Data := Db.GetFooddeliver_DB(Trackinstant.Loginid)
	Senddata, err := json.Marshal(Data)
	if err != nil {
		log.Println("Error -  INSTATNT DELIVERY", err)
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Access-Control-Allow-Orgin", "*")
	w.Write(Senddata)
}

type Foodorderid struct {
	Orderid int
}

func Getundelivered(w http.ResponseWriter, r *http.Request) {

	Data := Db.Getundelivered()
	Senddata, err := json.Marshal(Data)
	if err != nil {
		log.Println("Error -  INSTATNT DELIVERY", err)
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Access-Control-Allow-Orgin", "*")
	w.Write(Senddata)
}

func Getsinglefood(w http.ResponseWriter, r *http.Request) {

	var Orderid Foodorderid
	err := json.NewDecoder(r.Body).Decode(&Orderid)
	if err != nil {
		log.Println("Error - INSTATNT DELIVERY", err)
	}

	Data := Db.Getsingle_DB(Orderid.Orderid)
	Senddata, err := json.Marshal(Data)
	if err != nil {
		log.Println("Error -  INSTATNT DELIVERY", err)
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Access-Control-Allow-Orgin", "*")
	w.Write(Senddata)
}
func Cancelfooddelivery(w http.ResponseWriter, r *http.Request) {

	var Orderid Foodorderid
	err := json.NewDecoder(r.Body).Decode(&Orderid)
	if err != nil {
		log.Println("Error - INSTATNT DELIVERY", err)
	}

	Data := Db.Updatecancel_DB(Orderid.Orderid)
	Senddata, err := json.Marshal(Data)
	if err != nil {
		log.Println("Error -  INSTATNT DELIVERY", err)
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Access-Control-Allow-Orgin", "*")
	w.Write(Senddata)
}

func Updatefooddelivery(w http.ResponseWriter, r *http.Request) {

	var Orderid Foodorderid
	err := json.NewDecoder(r.Body).Decode(&Orderid)
	if err != nil {
		log.Println("Error - INSTATNT DELIVERY", err)
	}

	Data := Db.UpdateDelivered_DB(Orderid.Orderid)
	Senddata, err := json.Marshal(Data)
	if err != nil {
		log.Println("Error -  INSTATNT DELIVERY", err)
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Access-Control-Allow-Orgin", "*")
	w.Write(Senddata)
}

func Updateownfoodstatus(w http.ResponseWriter, r *http.Request) {

	var Orderid Foodorderid
	err := json.NewDecoder(r.Body).Decode(&Orderid)
	if err != nil {
		log.Println("Error - INSTATNT DELIVERY", err)
	}

	Data := Db.UpdateOwnDelivered_DB(Orderid.Orderid)
	Senddata, err := json.Marshal(Data)
	if err != nil {
		log.Println("Error -  INSTATNT DELIVERY", err)
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Access-Control-Allow-Orgin", "*")
	w.Write(Senddata)
}

//Ownfood Services

func Ownfooddeliver(w http.ResponseWriter, r *http.Request) {
	var Ownfoo Model.Ownfood
	err := json.NewDecoder(r.Body).Decode(&Ownfoo)
	if err != nil {
		log.Println("Error - INSTATNT DELIVERY", err)
	}

	Data := Db.FoodOwn_DB(Ownfoo)
	Senddata, err := json.Marshal(Data)
	if err != nil {
		log.Println("Error -  INSTATNT DELIVERY", err)
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Access-Control-Allow-Orgin", "*")
	w.Write(Senddata)
}

func Getfooddeliver(w http.ResponseWriter, r *http.Request) {

	Data := Db.GetFood_DB()
	Senddata, err := json.Marshal(Data)
	if err != nil {
		log.Println("Error -  INSTATNT DELIVERY", err)
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Access-Control-Allow-Orgin", "*")
	w.Write(Senddata)
}

func Orderowndeliver(w http.ResponseWriter, r *http.Request) {
	var Owner Model.Ownorderplaced
	err := json.NewDecoder(r.Body).Decode(&Owner)
	if err != nil {
		log.Println("Error - INSTATNT DELIVERY", err)
	}

	Data := Db.Orderfooddeliver_DB(Owner)
	Senddata, err := json.Marshal(Data)
	if err != nil {
		log.Println("Error -  INSTATNT DELIVERY", err)
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Access-Control-Allow-Orgin", "*")
	w.Write(Senddata)
}

func Getownundeliverfood(w http.ResponseWriter, r *http.Request) {

	Data := Db.Getownundeliverfood_DB()
	Senddata, err := json.Marshal(Data)
	if err != nil {
		log.Println("Error -  INSTATNT DELIVERY", err)
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Access-Control-Allow-Orgin", "*")
	w.Write(Senddata)
}

func Getallundeliverfood(w http.ResponseWriter, r *http.Request) {

	Data := Db.Getallundeliverfood_DB()
	Senddata, err := json.Marshal(Data)
	if err != nil {
		log.Println("Error -  INSTATNT DELIVERY", err)
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Access-Control-Allow-Orgin", "*")
	w.Write(Senddata)
}
func Getownsinglefood(w http.ResponseWriter, r *http.Request) {
	var Owner Foodorderid
	err := json.NewDecoder(r.Body).Decode(&Owner)
	if err != nil {
		log.Println("Error - INSTATNT DELIVERY", err)
	}

	Data := Db.OrderOwnsinglefood_DB(Owner.Orderid)
	Senddata, err := json.Marshal(Data)
	if err != nil {
		log.Println("Error -  INSTATNT DELIVERY", err)
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Access-Control-Allow-Orgin", "*")
	w.Write(Senddata)
}

func Getfoodsingledata(w http.ResponseWriter, r *http.Request) {
	var Owner Foodorderid
	err := json.NewDecoder(r.Body).Decode(&Owner)
	if err != nil {
		log.Println("Error - INSTATNT DELIVERY", err)
	}

	Data := Db.Getownsinglefood_DB(Owner.Orderid)
	Senddata, err := json.Marshal(Data)
	if err != nil {
		log.Println("Error -  INSTATNT DELIVERY", err)
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Access-Control-Allow-Orgin", "*")
	w.Write(Senddata)
}
func TrackOwnfooddeliver(w http.ResponseWriter, r *http.Request) {
	var Trackinstant Foodinstant
	err := json.NewDecoder(r.Body).Decode(&Trackinstant)
	if err != nil {
		log.Println("Error - INSTATNT DELIVERY", err)
	}

	Data := Db.TrackOwnfood_DB(Trackinstant.Loginid)
	Senddata, err := json.Marshal(Data)
	if err != nil {
		log.Println("Error -  INSTATNT DELIVERY", err)
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Access-Control-Allow-Orgin", "*")
	w.Write(Senddata)
}

func CancelOwnFooddeliver(w http.ResponseWriter, r *http.Request) {
	var Orderid Foodorderid
	err := json.NewDecoder(r.Body).Decode(&Orderid)
	if err != nil {
		log.Println("Error - INSTATNT DELIVERY", err)
	}

	Data := Db.CancelOwnFooddeliver_DB(Orderid.Orderid)
	Senddata, err := json.Marshal(Data)
	if err != nil {
		log.Println("Error -  INSTATNT DELIVERY", err)
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Access-Control-Allow-Orgin", "*")
	w.Write(Senddata)
}
