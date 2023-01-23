package controller

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"mocking/services"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gin-gonic/gin"
)

type mockPingService struct {
	handlePingFn func() (string, error)
}

func (mock mockPingService) HandlePing() (string, error) {
	return mock.handlePingFn()
}

type respo struct {
	Message string `json:"message"`
}

func Test_Ping_ServiceNoErr(t *testing.T) {
	pingService := mockPingService{}
	pingService.handlePingFn = func() (string, error) {
		return "pong", nil
	}
	services.PingService = pingService
	resp := httptest.NewRecorder()
	context, _ := gin.CreateTestContext(resp)

	Ping(context)

	// read response body
	body, _ := ioutil.ReadAll(resp.Body)

	var result respo

	err := json.Unmarshal(body, &result)
	if err != nil {
		fmt.Println("error while unmarshalling")
	}

	// fmt.Println(result)
	// fmt.Println(body)
	// fmt.Println(string(body))

	if resp.Code != http.StatusOK {
		t.Errorf("got %v, want %v", resp.Code, http.StatusOK)
	}

	if !strings.EqualFold(result.Message, "pong") {
		t.Errorf("message should be 'pong' but got %v", result.Message)
	}
}

func Test_Ping_ServiceHasErr(t *testing.T) {
	pingService := mockPingService{
		handlePingFn: func() (string, error) {
			return "", errors.New("error while creating ping service")
		},
	}

	services.PingService = pingService
	resp := httptest.NewRecorder()
	context, _ := gin.CreateTestContext(resp)

	Ping(context)

	if resp.Code != http.StatusBadRequest {
		t.Errorf("got %v, want %v", resp.Code, http.StatusBadRequest)
	}

}

// for testing go:- go test . (path of the test file)
// to see the coverage :- go test -cover .
// to see coverage in visual :- go test -coverprofile=coverage.out && go tool cover -html=coverage.out

// to run multiple test :-
// 1)to give a same naming convention before the test func  (Test_alpha)
// 2)with help of -run flag we can run
// eg go test -run Test_alpha
// all the test which has the test func name Test_alpha will run.
// go test -run Test_Ping -v -cover
// go test ./...
