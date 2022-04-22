package main

import "fmt"

//go:generate mockgen -destination=./mocks/customClientMock.go -package=mocks -source=./main.go

type Request struct {
	queryParamInt    int
	queryParamBool   bool
	queryParamString string
}

type Response struct {
	responseInt    int
	responseBool   bool
	responseString string
}

type CustomClient interface {
	GetBoolResponse(option int) bool
	GetStringResponse(option bool) string
	GetntegerResponse(option string) int
}

type customClient struct {
}

func NewCustomClient() *customClient {
	return &customClient{}
}

func (cc *customClient) GetBoolResponse(option int) bool {
	if option < 100 {
		return true
	}
	return false
}

func (cc *customClient) GetStringResponse(option bool) string {
	if option == true {
		return "true"
	}
	return "false"
}

func (cc *customClient) GetntegerResponse(option string) int {
	switch option {
	case "uno":
		return 1
	case "dos":
		return 2
	case "tres":
		return 3
	default:
	}
	return 0
}

type CustomService struct {
	client CustomClient
}

func NewCustomService(cl CustomClient) *CustomService {
	return &CustomService{client: cl}
}

func (cs *CustomService) FetchData() string {
	req := Request{1, true, "uno"}
	resp := Response{}
	resp.responseBool = cs.client.GetBoolResponse(req.queryParamInt)
	resp.responseInt = cs.client.GetntegerResponse(req.queryParamString)
	resp.responseString = cs.client.GetStringResponse(req.queryParamBool)

	return fmt.Sprintf("Response: %t, %s, %d", resp.responseBool, resp.responseString, resp.responseInt)
}

func main() {

}
