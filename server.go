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

	fs := http.FileServer(http.Dir("./rentmatics_theme/"))
	//http.Handle("/", fs)
	router.PathPrefix("/rentmatics_theme/").Handler(http.StripPrefix("/rentmatics_theme/", fs))

	fs1 := http.FileServer(http.Dir("./Houseimage/"))
	//http.Handle("/", fs)
	router.PathPrefix("/Houseimage/").Handler(http.StripPrefix("/Houseimage/", fs1))

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
	router.HandleFunc("/Leaves", Service.Leaves)
	router.HandleFunc("/Fruits", Service.Fruits)
	router.HandleFunc("/SearchProduct", Service.SearchProduct)
	router.HandleFunc("/Orderplaced", Service.Orderplaced)
	router.HandleFunc("/Ordertracking", Service.Ordertracking)
	router.HandleFunc("/Orderhistory", Service.Orderhistory)
	router.HandleFunc("/Instantdeliveryform", Service.Instantdeliveryform)

	//Backenddata
	router.HandleFunc("/Getalluser", Service.Allusers)
	router.HandleFunc("/Getallorder", Service.GetAllorder)
	router.HandleFunc("/Addproduct", Service.Allproduct)
	router.HandleFunc("/UpdateRate", Service.Editproduct)
	router.HandleFunc("/Getordertracking", Service.Getordertracking)

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
