package Model

type UserbackendResponse struct {
	Userid      string
	Username    string
	Loginid     string
	Phonenumber string
	Emailid     string
}
type Orderresponse struct {
	Orderid      string
	Billnumber   string
	Userid       string
	Loginid      string
	Totalproduct string
	Totalamount  string
	Date         string
}
type Productdata struct {
	Productid   string
	Productname string
	Productrate int
	Weight      string
}

type Orderdataresponse struct {
	Orderdetails    Orderresponse
	Productsdetails []Productdata
}

type Addproduct struct {
	Categorylist  string
	Productname   string
	Productrate   int
	Productweight string
}
type Updateproduct struct {
	Productname string
	Productrate int
}

type Gettrack struct {
	Orderid       int
	Orderplacedid string
	Loginid       string
	Status        string
}

type Ordertrackdata struct {
	Billid        string
	Phonenumber   string
	Totalproducts string
	TotalAmount   string
	Date          string
}

type Trackrespopnse struct {
	Trackdata   Gettrack
	Orderplaced Ordertrackdata
}
