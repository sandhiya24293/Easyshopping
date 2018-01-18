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

func Instantdeliveryform(w http.ResponseWriter, r *http.Request) {
	var Instance Model.Easyinstance
	err := json.NewDecoder(r.Body).Decode(&Instance)
	if err != nil {
		log.Println("Error - INSTATNT DELIVERY", err)
	}
	rand.Seed(time.Now().UnixNano())
	randomno := randomInt(1234567, 7654321)
	uniqueid := "E3INSTANCE" + strconv.Itoa(randomno)

	Data := Db.Instance_DB(Instance, uniqueid)
	Senddata, err := json.Marshal(Data)
	if err != nil {
		log.Println("Error -  INSTATNT DELIVERY", err)
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Access-Control-Allow-Orgin", "*")
	w.Write(Senddata)
}

func Getinstanse(w http.ResponseWriter, r *http.Request) {

	Data := Db.EasyInstance()
	Senddata, err := json.Marshal(Data)
	if err != nil {
		log.Println("Error -  INSTATNT DELIVERY", err)
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Access-Control-Allow-Orgin", "*")
	w.Write(Senddata)
}

type Trackinstant struct {
	Loginid string
}

func GetTrackinstanse(w http.ResponseWriter, r *http.Request) {
	var Trackinstant Trackinstant
	err := json.NewDecoder(r.Body).Decode(&Trackinstant)
	if err != nil {
		log.Println("Error - INSTATNT DELIVERY", err)
	}

	Data := Db.Gettrackinginstance(Trackinstant.Loginid)
	Senddata, err := json.Marshal(Data)
	if err != nil {
		log.Println("Error -  INSTATNT DELIVERY", err)
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Access-Control-Allow-Orgin", "*")
	w.Write(Senddata)
}
func Getinstanstundelivered(w http.ResponseWriter, r *http.Request) {

	Data := Db.Instantundelivered_DB()
	Senddata, err := json.Marshal(Data)
	if err != nil {
		log.Println("Error -  INSTATNT DELIVERY", err)
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Access-Control-Allow-Orgin", "*")
	w.Write(Senddata)
}

func Updateinstantdeliver(w http.ResponseWriter, r *http.Request) {

	var Up_instance Model.Updateinstant
	err := json.NewDecoder(r.Body).Decode(&Up_instance)
	if err != nil {
		log.Println("Error - INSTATNT DELIVERY", err)
	}

	Data := Db.Updateinstant_DB(Up_instance)
	Senddata, err := json.Marshal(Data)
	if err != nil {
		log.Println("Error -  INSTATNT DELIVERY", err)
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Access-Control-Allow-Orgin", "*")
	w.Write(Senddata)
}
