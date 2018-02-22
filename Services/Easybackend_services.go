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
	Userdata.Pictureurl = "http://176.111.105.86:8085/" + Pictureurl
	Db.Addproduct_DB(Userdata)

}

func AddFoodproduct(w http.ResponseWriter, r *http.Request) {

	var Userdata Model.Ownfood
	var plate string
	var rate string
	r.ParseMultipartForm(32 << 20)
	r.ParseForm()

	Userdata.Dishname = r.Form["dish"][0]
	plate = r.Form["count"][0]
	rate = r.Form["rate"][0]
	Userdata.Platecount, _ = strconv.Atoi(plate)
	Userdata.Rate, _ = strconv.Atoi(rate)

	file, handler, err := r.FormFile("uploadfile2")

	Pictureurl := "Productimage/" + handler.Filename

	f, err := os.OpenFile(Pictureurl, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer f.Close()

	io.Copy(f, file)
	Userdata.Image = "http://176.111.105.86:8085/" + Pictureurl
	Db.AddFood_DB(Userdata)

}
func NonAddproduct(w http.ResponseWriter, r *http.Request) {

	var Userdata Model.NonAddproduct

	r.ParseMultipartForm(32 << 20)
	r.ParseForm()

	Userdata.Categorylist = r.Form["meat"][0]
	Userdata.Productname = r.Form["meatname"][0]
	fmt.Println(Userdata.Productname)

	Productrate1kg := r.Form["meatrate3"][0]
	Productrate500kg := r.Form["meatrate2"][0]
	Productrate250gm := r.Form["meatrate1"][0]
	Productrate1250gm := r.Form["meatrate4"][0]
	Productrate1500gm := r.Form["meatrate5"][0]
	Productrate1750gm := r.Form["meatrate6"][0]
	Productrate2000gm := r.Form["meatrate7"][0]
	file, handler, err := r.FormFile("uploadfile1")

	Pictureurl := "Productimage/" + handler.Filename
	Userdata.Productrate1kg, _ = strconv.Atoi(Productrate1kg)
	Userdata.Productrate500kg, _ = strconv.Atoi(Productrate500kg)
	Userdata.Productrate250gm, _ = strconv.Atoi(Productrate250gm)
	Userdata.Productrate1250gm, _ = strconv.Atoi(Productrate1250gm)
	Userdata.Productrate1500gm, _ = strconv.Atoi(Productrate1500gm)
	Userdata.Productrate1750gm, _ = strconv.Atoi(Productrate1750gm)
	Userdata.Productrate2000gm, _ = strconv.Atoi(Productrate2000gm)

	f, err := os.OpenFile(Pictureurl, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer f.Close()

	io.Copy(f, file)
	Userdata.Pictureurl = "http://176.111.105.86:8085/" + Pictureurl
	Db.AddNonproduct_DB(Userdata)

}
func Editproduct(w http.ResponseWriter, r *http.Request) {
	var Editprod Model.Updateproduct
	err := json.NewDecoder(r.Body).Decode(&Editprod)
	if err != nil {
		log.Println("Error - INSTATNT DELIVERY", err)
	}

	Db.Updateproduct_DB(Editprod)

}

func UpdatefoodRate(w http.ResponseWriter, r *http.Request) {
	var Editprod Model.UpdateFoodproduct
	err := json.NewDecoder(r.Body).Decode(&Editprod)
	if err != nil {
		log.Println("Error - INSTATNT DELIVERY", err)
	}

	Db.Updatefoodproduct_DB(Editprod)

}
func UpdateNonRate(w http.ResponseWriter, r *http.Request) {
	var Editprod Model.Updatenonproduct
	err := json.NewDecoder(r.Body).Decode(&Editprod)
	if err != nil {
		log.Println("Error - INSTATNT DELIVERY", err)
	}

	Db.Updatenonproduct_DB(Editprod)

}

func Updatestatus(w http.ResponseWriter, r *http.Request) {
	var Upprod Model.Updatestatus
	err := json.NewDecoder(r.Body).Decode(&Upprod)
	if err != nil {
		log.Println("Error - INSTATNT DELIVERY", err)
	}

	Db.Updatestatus_DB(Upprod)

}

func UpdatenonvegProductstatus(w http.ResponseWriter, r *http.Request) {
	var Upprod Model.Updatestatus
	err := json.NewDecoder(r.Body).Decode(&Upprod)
	if err != nil {
		log.Println("Error - INSTATNT DELIVERY", err)
	}

	Db.Updatestatusnon_DB(Upprod)

}
func Changefoodstatus(w http.ResponseWriter, r *http.Request) {
	var Upprod Model.Updatestatus
	err := json.NewDecoder(r.Body).Decode(&Upprod)
	if err != nil {
		log.Println("Error - INSTATNT DELIVERY", err)
	}

	Db.Updatefoodstatus_DB(Upprod)

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
func SendFruit(w http.ResponseWriter, r *http.Request) {
	var Response []Model.Catergorylist
	var Datasend Model.Datares

	var Data1 Model.Catergorylist

	var Category = []string{"Fruits", "Fruits1", "Fruits2", "Fruits3", "Fruits4"}
	for i, v := range Category {
		Data1.CategoryName = v
		if i == 0 {
			Data1.Url = "http://176.111.105.86:8085/Fruits"
		} else if i == 1 {
			Data1.Url = "http://176.111.105.86:8085/Fruits1"

		} else if i == 2 {
			Data1.Url = "http://176.111.105.86:8085/Fruits2"

		} else if i == 3 {
			Data1.Url = "http://176.111.105.86:8085/Fruits3"

		} else {
			Data1.Url = "http://176.111.105.86:8085/Fruits4"

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
func SendFood(w http.ResponseWriter, r *http.Request) {
	var Response []Model.Catergorylist
	var Datasend Model.Datares

	var Data1 Model.Catergorylist

	var Category = []string{"food"}
	for i, v := range Category {
		Data1.CategoryName = v
		if i == 0 {
			Data1.Url = "http://176.111.105.86:8085/GetOwnfooddeliver"
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

	var Category = []string{"Meat", "Frozen", "Sea food"}
	for i, v := range Category {
		Data1.CategoryName = v
		if i == 0 {
			Data1.Url = "http://176.111.105.86:8085/NonVeg"

		} else if i == 1 {
			Data1.Url = "http://176.111.105.86:8085/Legpiece"

		} else {
			Data1.Url = "http://176.111.105.86:8085/Seafood"

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

func Getnonsingledata(w http.ResponseWriter, r *http.Request) {
	var Getid Productid
	err := json.NewDecoder(r.Body).Decode(&Getid)
	if err != nil {
		log.Println("Error - INSTATNT DELIVERY", err)
	}

	Response := Db.GetnonSingleprod_DB(Getid.Getproductid)
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
