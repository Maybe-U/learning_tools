package v3_service

type ComRespone struct {
	Code int `json:"code"`
}

type Add struct {
	A int `json:"a"`
	B int `json:"b"`
}

type AddAck struct {
	ComRespone
	Res int `json:"res"`
}

type Login struct {
	Account  string `json:"account"`
	Password string `json:"password"`
}

type LoginAck struct {
	ComRespone
	Token string `json:"token"`
}
