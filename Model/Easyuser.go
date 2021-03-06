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
	Productid   int
	Productname string
	Productrate int
	Weight      string
}

type Orderdataresponse struct {
	Orderdetails    Orderresponse
	Productsdetails []Productdata
}

type Addproduct struct {
	Categorylist     string
	Productname      string
	Productrate1kg   int
	Productrate500kg int
	Productrate250gm int
	Pictureurl       string
}

type NonAddproduct struct {
	Categorylist      string
	Productname       string
	Productrate1kg    int
	Productrate500kg  int
	Productrate250gm  int
	Productrate1250gm int
	Productrate1500gm int
	Productrate1750gm int
	Productrate2000gm int
	Pictureurl        string
}
type Updateproduct struct {
	Productid        int
	Productname      string
	Productrate1kg   int
	Productrate500gm int
	Productrate250gm int
}
type UpdateFoodproduct struct {
	Foodid     int
	Dishname   string
	Dishrate   int
	Platecount int
}
type Updatenonproduct struct {
	Productid         int
	Productname       string
	Productrate2kg    int
	Productrate1750gm int
	Productrate1500gm int
	Productrate1250gm int
	Productrate1kg    int
	Productrate500gm  int
	Productrate250gm  int
}
type Updatestatus struct {
	Productid     int
	Productstatus string
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

type Trackupdate struct {
	Orderid     int
	Orderstatus string
}
type Datacount struct {
	Usercount    int
	Productcount int
	Ordercount   int
	Instantcount int
}

type Changepassword struct {
	Loginid string
	Oldpass string
	Newpass string
}
type AdminLogin struct {
	User string
	Pass string
}
