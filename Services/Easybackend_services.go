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

func SendVegetable(w http.ResponseWriter, r *http.Request) {
	var Response []Model.Catergorylist
	var Datasend Model.Datares

	var Data1 Model.Catergorylist

	var Category = []string{"Vegetables", "Green Leaves", "Fruits"}
	for i, v := range Category {
		Data1.CategoryName = v
		if i == 0 {
			Data1.Url = "http://176.111.105.86:8085/Vegetables"
		} else if i == 1 {
			Data1.Url = "http://176.111.105.86:8085/Leaves"
		} else {
			Data1.Url = "http://176.111.105.86:8085/Fruits"

		}
		Response = append(Response, Data1)

	}
	Datasend.Data = Response
	Senddata, err := json.Marshal(Datasend)
	if err != nil {
		log.Println("Error -  INSTATNT DELIVERY", err)
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Access-Control-Allow-Orgin", "*")
	w.Write(Senddata)
}

func SendNonVeg(w http.ResponseWriter, r *http.Request) {
	var Response []Model.Catergorylist

	var Data1 Model.Catergorylist

	var Category = []string{"Meat", "Sea food"}
	for i, v := range Category {
		Data1.CategoryName = v
		if i == 0 {
			Data1.Url = "http://176.111.105.86:8085/NonVeg"
		} else {
			Data1.Url = "http://176.111.105.86:8085/Seafood"

		}
		Response = append(Response, Data1)

	}
	Senddata, err := json.Marshal(Response)
	if err != nil {
		log.Println("Error -  INSTATNT DELIVERY", err)
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Access-Control-Allow-Orgin", "*")
	w.Write(Senddata)
}
