package Services

import (
	Db "Easyshopping/Common/DB/Mysql"
	Model "Easyshopping/Model"
	"encoding/json"
	"log"

	"net/http"
)

func Allusers(w http.ResponseWriter, r *http.Request) {

	Data := Db.Getuser_DB()
	Senddata, err := json.Marshal(Data)
	if err != nil {
		log.Println("Error -  INSTATNT DELIVERY", err)
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Access-Control-Allow-Orgin", "*")
	w.Write(Senddata)
}

func GetAllorder(w http.ResponseWriter, r *http.Request) {

	Data := Db.GetAllorder_DB()
	Senddata, err := json.Marshal(Data)
	if err != nil {
		log.Println("Error -  INSTATNT DELIVERY", err)
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Access-Control-Allow-Orgin", "*")
	w.Write(Senddata)
}
func Allproduct(w http.ResponseWriter, r *http.Request) {
	var Addproductdata Model.Addproduct
	err := json.NewDecoder(r.Body).Decode(&Addproductdata)
	if err != nil {
		log.Println("Error - INSTATNT DELIVERY", err)
	}

	Db.Addproduct_DB(Addproductdata)

}
func Editproduct(w http.ResponseWriter, r *http.Request) {
	var Editprod Model.Updateproduct
	err := json.NewDecoder(r.Body).Decode(&Editprod)
	if err != nil {
		log.Println("Error - INSTATNT DELIVERY", err)
	}

	Db.Updateproduct_DB(Editprod)
	//	Senddata, err := json.Marshal(Data)
	//	if err != nil {
	//		log.Println("Error -  INSTATNT DELIVERY", err)
	//	}

	//	w.WriteHeader(http.StatusOK)
	//	w.Header().Set("Access-Control-Allow-Orgin", "*")
	//	w.Write(Senddata)
}

func Getordertracking(w http.ResponseWriter, r *http.Request) {

	Data := Db.Getordertracking_DB()
	Senddata, err := json.Marshal(Data)
	if err != nil {
		log.Println("Error -  INSTATNT DELIVERY", err)
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Access-Control-Allow-Orgin", "*")
	w.Write(Senddata)
}
