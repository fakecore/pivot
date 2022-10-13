package domain

type HelloReq struct {
	Name string `form:"name"`
}

type HelloResp struct {
	Body string `json:"body"`
}
