package services

import "fmt"

type pingService interface {
	HandlePing() (string, error)
}

type pingServiceImpl struct {
}

var (
	PingService pingService = pingServiceImpl{}
)

// Cannot mock once it is compiled
func (s pingServiceImpl) HandlePing() (string, error) {
	fmt.Println("some complex things......")
	return "pong", nil
}
