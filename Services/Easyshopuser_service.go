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

// Returns an int >= min, < max
func randomInt(min, max int) int {
	return min + rand.Intn(max-min)
}

func EasyRegister(w http.ResponseWriter, r *http.Request) {
	var Register Model.EasyRegister
	err := json.NewDecoder(r.Body).Decode(&Register)
	if err != nil {
		log.Println("Error - User REGISTER", err)
	}
	rand.Seed(time.Now().UnixNano())
	randomno := randomInt(1234567, 7654321)
	uniqueid := "E3" + strconv.Itoa(randomno)
	Data := Db.InsertnewUser(Register, uniqueid)

	Senddata, err := json.Marshal(Data)
	if err != nil {
		log.Println("Error -  User REGISTER", err)
	}
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Access-Control-Allow-Orgin", "*")
	w.Write(Senddata)

}

func EasyLogin(w http.ResponseWriter, r *http.Request) {
	var Easyuser Model.EasyLogin
	err := json.NewDecoder(r.Body).Decode(&Easyuser)
	if err != nil {
		log.Println("Error - User Login", err)
	}
	Data := Db.LoginUser(Easyuser)
	Senddata, err := json.Marshal(Data)

	if err != nil {
		log.Println("Error - User Login", err)
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Access-Control-Allow-Orgin", "*")

	w.Write(Senddata)

}

func Changepassword(w http.ResponseWriter, r *http.Request) {

	var User1 Model.Changepass
	err := json.NewDecoder(r.Body).Decode(&User1)
	if err != nil {
		log.Println("Error - Change password", err)
	}
	Data := Db.EasyChangepassword(User1)
	Senddata, err := json.Marshal(Data)
	if err != nil {
		log.Println("Error - Change password", err)
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Access-Control-Allow-Orgin", "*")
	w.Write(Senddata)

}

type Forgot struct {
	Email string
}

func Forgotpassword(w http.ResponseWriter, r *http.Request) {

	var Emailid Forgot
	err := json.NewDecoder(r.Body).Decode(&Emailid)
	if err != nil {
		log.Println("Error - FORGOT password", err)
	}
	Data := Db.EasyForgotpassword(Emailid.Email)
	Senddata, err := json.Marshal(Data)
	if err != nil {
		log.Println("Error - FORGOT password", err)
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Access-Control-Allow-Orgin", "*")
	w.Write(Senddata)

}
func EasyuserProfile(w http.ResponseWriter, r *http.Request) {

	var Loginidd Model.ProfileLoginid
	err := json.NewDecoder(r.Body).Decode(&Loginidd)
	if err != nil {
		log.Println("Error - USER PROFILE", err)
	}
	Data := Db.EasyRetriveprofile(Loginidd.Loginid)
	Senddata, err := json.Marshal(Data)
	if err != nil {
		log.Println("Error - USER PROFILE - RESPONSE", err)
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Access-Control-Allow-Orgin", "*")
	w.Write(Senddata)

}
func EasyInsertProfile(w http.ResponseWriter, r *http.Request) {

	var ProfileAddress Model.InsertAddress
	err := json.NewDecoder(r.Body).Decode(&ProfileAddress)
	if err != nil {
		log.Println("Error - INSERT PROFILE", err)
	}
	Data := Db.Easyinsertprofile(ProfileAddress)
	Senddata, err := json.Marshal(Data)
	if err != nil {
		log.Println("Error - INSERT PROFILE", err)
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Access-Control-Allow-Orgin", "*")
	w.Write(Senddata)

}

func UpdateProfile(w http.ResponseWriter, r *http.Request) {

	var ProfileAddress Model.UpdateProfile
	err := json.NewDecoder(r.Body).Decode(&ProfileAddress)
	if err != nil {
		log.Println("Error - UPDATE PROFILE", err)
	}
	Data := Db.Easyupdateprofile(ProfileAddress)
	Senddata, err := json.Marshal(Data)
	if err != nil {
		log.Println("Error - UPDATE PROFILE", err)
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Access-Control-Allow-Orgin", "*")
	w.Write(Senddata)

}
