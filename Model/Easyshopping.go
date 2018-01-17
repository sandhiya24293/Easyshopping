package Model

type EasyRegister struct {
	Username    string
	Loginid     string
	Password    string
	Phonenumber string
	Emailid     string
	Logintype   string
}
type UserResponse struct {
	Username    string
	Loginid     string
	Phonenumber string
	Status      string
}
type APIError struct {
	Error struct {
		Code    int    `json:"code"`
		Message string `json:"message"`
	} `json:"error"`
}

type EasyLogin struct {
	Loginid  string
	Password string
}

type Easyloginverify struct {
	Username    string
	Loginid     string
	Password    string
	Phonenumber string
}
type Changepass struct {
	Loginid     string
	Oldpassword string
	Newpassword string
}

type ProfileLoginid struct {
	Loginid string
}

type Profile struct {
	Userid      int
	Uniueid     string
	Username    string
	Loginid     string
	Phonenumber string
	Emailid     string
}
type Profileresponse struct {
	Profileresponse Profile
	Address         string
}
type InsertAddress struct {
	Loginid string
	Address string
}
type UpdateProfile struct {
	Loginid     string
	Username    string
	Phonenumber string
	Emailid     string
	Address     string
}
