package Services

import (
	Db "Easyshopping/Common/DB/Mysql"
	Model "Easyshopping/Model"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"

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

type Orderid struct {
	Getorder int
}

//func Getsingleorderdata(w http.ResponseWriter, r *http.Request) {
//	var Addproductdata Orderid
//	err := json.NewDecoder(r.Body).Decode(&Addproductdata)
//	if err != nil {
//		log.Println("Error - INSTATNT DELIVERY", err)
//	}

//	Db.SingleOrderdata_DB(Addproductdata.Getorder)

//}

//Categorylist     string
//	Productname      string
//	Productrate1kg   int
//	Productrate500kg int
//	Productrate250gm int
//	Pictureurl       string
//Executive Upload
func Allproduct(w http.ResponseWriter, r *http.Request) {

	var Userdata Model.Addproduct

	r.ParseMultipartForm(32 << 20)
	r.ParseForm()

	Userdata.Categorylist = r.Form["category"][0]
	Userdata.Productname = r.Form["namevalue"][0]
	fmt.Println(Userdata.Productname)

	Productrate1kg := r.Form["getrate1"][0]
	Productrate500kg := r.Form["getrate1"][0]
	Productrate250gm := r.Form["getrate3"][0]
	file, handler, err := r.FormFile("uploadfile")

	Pictureurl := "Productimage/" + handler.Filename
	Userdata.Productrate1kg, _ = strconv.Atoi(Productrate1kg)
	Userdata.Productrate500kg, _ = strconv.Atoi(Productrate500kg)
	Userdata.Productrate250gm, _ = strconv.Atoi(Productrate250gm)

	f, err := os.OpenFile(Pictureurl, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer f.Close()

	io.Copy(f, file)
	Userdata.Pictureurl = "http://localhost:8085/" + Pictureurl
	Db.Addproduct_DB(Userdata)

}

func Editproduct(w http.ResponseWriter, r *http.Request) {
	var Editprod Model.Updateproduct
	err := json.NewDecoder(r.Body).Decode(&Editprod)
	if err != nil {
		log.Println("Error - INSTATNT DELIVERY", err)
	}

	Db.Updateproduct_DB(Editprod)

}

func Updatestatus(w http.ResponseWriter, r *http.Request) {
	var Upprod Model.Updatestatus
	err := json.NewDecoder(r.Body).Decode(&Upprod)
	if err != nil {
		log.Println("Error - INSTATNT DELIVERY", err)
	}

	Db.Updatestatus_DB(Upprod)

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
			Data1.Url = "http://176.111.105.86:8085/Vegtables"
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
	var Datasend Model.Datares

	var Data1 Model.Catergorylist

	var Category = []string{"Meat", "Sea food", "Turkey"}
	for i, v := range Category {
		Data1.CategoryName = v
		if i == 0 {
			Data1.Url = "http://176.111.105.86:8085/NonVeg"
		} else if i == 1 {
			Data1.Url = "http://176.111.105.86:8085/Seafood"

		} else {
			Data1.Url = "http://176.111.105.86:8085/Turkey"

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

type Productid struct {
	Getproductid int
}

func GetSingleproduct(w http.ResponseWriter, r *http.Request) {
	var Getid Productid
	err := json.NewDecoder(r.Body).Decode(&Getid)
	if err != nil {
		log.Println("Error - INSTATNT DELIVERY", err)
	}

	Response := Db.GetSingleprod_DB(Getid.Getproductid)
	Senddata, err := json.Marshal(Response)
	if err != nil {
		log.Println("Error -  INSTATNT DELIVERY", err)
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Access-Control-Allow-Orgin", "*")
	w.Write(Senddata)

}
func Changepassword_admin(w http.ResponseWriter, r *http.Request) {
	var Changepass Model.Changepassword
	err := json.NewDecoder(r.Body).Decode(&Changepass)
	if err != nil {
		log.Println("Error - INSTATNT DELIVERY", err)
	}

	Response := Db.Changepassword_DB(Changepass)
	Senddata, err := json.Marshal(Response)
	if err != nil {
		log.Println("Error -  INSTATNT DELIVERY", err)
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Access-Control-Allow-Orgin", "*")
	w.Write(Senddata)

}

func AdminLogin(w http.ResponseWriter, r *http.Request) {
	var Admin Model.AdminLogin
	err := json.NewDecoder(r.Body).Decode(&Admin)
	if err != nil {
		log.Println("Error - INSTATNT DELIVERY", err)
	}

	Response := Db.Adminlogin_DB(Admin)
	Senddata, err := json.Marshal(Response)
	if err != nil {
		log.Println("Error -  INSTATNT DELIVERY", err)
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Access-Control-Allow-Orgin", "*")
	w.Write(Senddata)

}
