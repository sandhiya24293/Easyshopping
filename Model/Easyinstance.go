package Model

type Easyinstance struct {
	Productname      string
	Dispactchername  string
	Dispatcheraddess string
	Dispatchernumber string
	Deliveryaddress  string
	Deliveryname     string
	Deliverynumber   string
	Message          string
	Status           string
}
type Easyinstances struct {
	Productid        int
	Uniqueid         string
	Productname      string
	Dispactchername  string
	Dispatcheraddess string
	Dispatchernumber string
	Deliveryaddress  string
	Deliveryname     string
	Deliverynumber   string
	Message          string
	Status           string
}
type Updateinstant struct {
	Easyid int
	Status string
}
