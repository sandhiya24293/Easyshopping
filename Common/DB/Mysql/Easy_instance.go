
package Mysql

import (
	Model "Easyshopping/Model"
	_ "database/sql"
     "log"

	_ "github.com/go-sql-driver/mysql"
)

    

func Instance_DB(Instancedata Model.Easyinstance , Uniqueid string) string {

	

	row, err := OpenConnection["Rentmatics"].Exec("insert into easydelivery (uniqueid,Productname,Dispactchername,Dispatcheraddess,Dispatchernumber,Deliveryaddress,Deliveryname,Deliverynumber,Message ) values (?,?,?,?,?,?,?,?,?)",Uniqueid,Instancedata.Productname,Instancedata.Dispactchername,Instancedata.Dispatchernumber,Instancedata.Dispatcheraddess,Instancedata.Deliveryaddress,Instancedata.Deliveryname,Instancedata.Deliverynumber,Instancedata.Message )
	if err != nil {
		log.Println("Error -DB: update Profile", err, row)
	}

	return "Success"

}
