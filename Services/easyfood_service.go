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

//func Getinstanse(w http.ResponseWriter, r *http.Request) {

//	Data := Db.EasyInstance()
//	Senddata, err := json.Marshal(Data)
//	if err != nil {
//		log.Println("Error -  INSTATNT DELIVERY", err)
//	}

//	w.WriteHeader(http.StatusOK)
//	w.Header().Set("Access-Control-Allow-Orgin", "*")
//	w.Write(Senddata)
//}

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
