package Model

type Easyinstance struct {
	Productname     string
	Deliveryaddress string
	Deliveryname    string
	Deliverynumber  string
	Message         string
	Status          string
	Loginid         string
}
type Easyinstances struct {
	Productid       int
	Uniqueid        string
	Productname     string
	Deliveryaddress string
	Deliveryname    string
	Deliverynumber  string
	Message         string
	Status          string
	Loginid         string
}
type Updateinstant struct {
	Easyid int
	Status string
}
