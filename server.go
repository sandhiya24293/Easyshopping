package main

import (
	Model "Easyshopping/Model"
	Service "Easyshopping/Services"
	"crypto/tls"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

const (
	contenttypeJSON = "application/json; charset=utf-8"
)

//Error provides error message in JSON format to client (utility)
func Error(w http.ResponseWriter, err int, msg string) {

	e := Model.APIError{}
	e.Error.Code = err
	e.Error.Message = msg
	w.Header().Set("Content-Type", contenttypeJSON)
	w.WriteHeader(err)
	json.NewEncoder(w).Encode(e)

}

func Serve() bool {

	router := mux.NewRouter()

	fs := http.FileServer(http.Dir("./web/"))
	router.PathPrefix("/web/").Handler(http.StripPrefix("/web/", fs))

	fs1 := http.FileServer(http.Dir("./Productimage/"))
	router.PathPrefix("/Productimage/").Handler(http.StripPrefix("/Productimage/", fs1))

	//Easy shopping Handlers
	router.HandleFunc("/EasyshoppingRegister", Service.EasyRegister)
	router.HandleFunc("/EasyshoppingLogin", Service.EasyLogin)
	router.HandleFunc("/Changepassword", Service.Changepassword)
	router.HandleFunc("/Forgotpassword", Service.Forgotpassword)
	router.HandleFunc("/Profile", Service.EasyuserProfile)
	router.HandleFunc("/InsertAddress", Service.EasyInsertProfile)
	router.HandleFunc("/UpdateProfile", Service.UpdateProfile)

	//VEGETABLES
	router.HandleFunc("/Vegtables", Service.Vegetables)
	router.HandleFunc("/NonVeg", Service.NonVeg)
	router.HandleFunc("/Seafood", Service.Seafood)
	router.HandleFunc("/Leaves", Service.Leaves)
	router.HandleFunc("/Fruits", Service.Fruits)
	router.HandleFunc("/Turkey", Service.Turkey)
	router.HandleFunc("/Chicken", Service.Chicken)
	router.HandleFunc("/Chickenbiriyani", Service.Chickenbiriyani)
	router.HandleFunc("/Muttonbiriyani", Service.Muttonbiriyani)
	router.HandleFunc("/friedrice", Service.Friedrice)
	router.HandleFunc("/noodles", Service.Noodles)
	router.HandleFunc("/Grill", Service.Grill)
	router.HandleFunc("/Legpiece", Service.Legpiece)

	router.HandleFunc("/SendVegetable", Service.SendVegetable)
	router.HandleFunc("/SendNonVeg", Service.SendNonVeg)
	router.HandleFunc("/SendFruit", Service.SendFruit)
	router.HandleFunc("/SendFood", Service.SendFood)

	//Vegetables

	router.HandleFunc("/SearchProduct", Service.SearchProduct)
	router.HandleFunc("/Orderplaced", Service.Orderplaced)
	router.HandleFunc("/Ordertracking", Service.Ordertracking)
	router.HandleFunc("/Ordertrackingadmin", Service.Ordertrackingadmin)
	router.HandleFunc("/OrderCancel", Service.Ordercancel)
	router.HandleFunc("/Trackstatus", Service.Trackstatus)
	router.HandleFunc("/updatetrackstatus", Service.Updatetrackstatus)

	router.HandleFunc("/Orderhistory", Service.Orderhistory)
	router.HandleFunc("/Instantdeliveryform", Service.Instantdeliveryform)
	router.HandleFunc("/Getinstanse", Service.Getinstanse)
	router.HandleFunc("/GetTrackinstant", Service.GetTrackinstanse)
	router.HandleFunc("/Getinstanstundelivered", Service.Getinstanstundelivered)
	router.HandleFunc("/Updateinstantdeliver", Service.Updateinstantdeliver)

	//Backenddata
	router.HandleFunc("/Getalluser", Service.Allusers)
	router.HandleFunc("/Getallorder", Service.GetAllorder)
	router.HandleFunc("/Addproduct", Service.Allproduct)
	router.HandleFunc("/NonAddproduct", Service.NonAddproduct)
	router.HandleFunc("/AddFoodproduct", Service.AddFoodproduct)

	router.HandleFunc("/UpdateRate", Service.Editproduct)
	router.HandleFunc("/UpdateNonRate", Service.UpdateNonRate)
	router.HandleFunc("/Getsingledata", Service.GetSingleproduct)
	router.HandleFunc("/Getnonsingledata", Service.Getnonsingledata)

	router.HandleFunc("/Getproductdata", Service.Getproductdata)
	//router.HandleFunc("/Getsingleorderdata", Service.Getsingleorderdata)

	router.HandleFunc("/UpdateProductstatus", Service.Updatestatus)

	router.HandleFunc("/Getordertracking", Service.Getordertracking)
	router.HandleFunc("/Getcount_service", Service.Getcount_service)

	router.HandleFunc("/Changepassword_admin", Service.Changepassword_admin)
	router.HandleFunc("/AdminLogin", Service.AdminLogin)

	//Food Instance
	router.HandleFunc("/Foodinsert", Service.Foodinsert)
	router.HandleFunc("/Getallfooddelivery", Service.Getallfooddelivery)
	//router.HandleFunc("/Getsinglefooddata", Service.Getsinglefooddata)
	router.HandleFunc("/Trackfooddeliver", Service.Trackfooddeliver)
	router.HandleFunc("/Cancelfooddelivery", Service.Cancelfooddelivery)
	router.HandleFunc("/gethistory", Service.Fooddelieryhistory)

	//backendside
	router.HandleFunc("/updatefooddelivered", Service.Updatefooddelivery)
	router.HandleFunc("/Getundelivered", Service.Getundelivered)
	router.HandleFunc("/Getsinglefood", Service.Getsinglefood)

	//Food Deliver
	//Own Food
	router.HandleFunc("/GetOwnfooddeliver", Service.Getfooddeliver)
	router.HandleFunc("/Ownfooddeliver", Service.Ownfooddeliver)
	router.HandleFunc("/Orderowndeliver", Service.Orderowndeliver)
	router.HandleFunc("/UpdatefoodRate", Service.UpdatefoodRate)
	router.HandleFunc("/Getfoodsingledata", Service.Getfoodsingledata)

	router.HandleFunc("/Getownundeliverfood", Service.Getownundeliverfood)
	router.HandleFunc("/Getallundeliverfood", Service.Getallundeliverfood)
	router.HandleFunc("/updateownfoodstatus", Service.Updateownfoodstatus)

	router.HandleFunc("/Getownsinglefood", Service.Getownsinglefood)
	router.HandleFunc("/TrackOwnfooddeliver", Service.TrackOwnfooddeliver)
	router.HandleFunc("/CancelOwnFooddeliver", Service.CancelOwnFooddeliver)

	//Search

	//For HTTPS

	// Default server - non-trusted for debugging
	serverhttp := func() {
		fmt.Println("Server should be available at http", config.Port)
		fmt.Println(http.ListenAndServe(config.Port, router))
	}

	// Setup TLS parameters for trusted server implementation
	if config.SSL && config.Key != "" && config.Cert != "" {
		// Setup TLS parameters
		tlsConfig := &tls.Config{
			ClientAuth:   tls.NoClientCert,
			MinVersion:   tls.VersionTLS12,
			Certificates: make([]tls.Certificate, 1),
		}

		var err error
		// Setup API server private key and certificate
		tlsConfig.Certificates[0], err = tls.X509KeyPair([]byte(config.Cert), []byte(config.Key))
		if err != nil {
			fmt.Println("Error during decoding service key and certificate:", err)
			return false
		}

		tlsConfig.BuildNameToCertificate()

		https := &http.Server{
			Addr:      config.Https_port,
			TLSConfig: tlsConfig,
			Handler:   router,
		}

		// Trusted server implementation
		server := func() {
			fmt.Println("Server should be available at https", config.Https_port)
			fmt.Println(https.ListenAndServeTLS("", ""))
		}
		go server()
	}

	// Schedule API server
	go serverhttp()

	return true
}
