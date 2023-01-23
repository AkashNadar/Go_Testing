package main

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type ExampleSuite struct {
	suite.Suite
}

func TestExampleSuiteMain(t *testing.T) {
	suite.Run(t, &ExampleSuite{})
}

func (s *ExampleSuite) TestTrue() {
	s.T().Log("Test-True Executing...")
	s.True(true)
}

func (s *ExampleSuite) TestFalse() {
	s.T().Log("Test-False Executing...")
	s.False(false)
}

func (s *ExampleSuite) SetupSuite() {
	s.T().Log("Set-up Suite...")
}

func (s *ExampleSuite) TearDownSuite() {
	s.T().Log("Tear-Down Suite...")
}

func (s *ExampleSuite) SetupTest() {
	s.T().Log("Set-up Test...")
}

func (s *ExampleSuite) TearDownTest() {
	s.T().Log("Tear-down Test...")
}
func (s *ExampleSuite) BeforeTest(suiteName, testName string) {
	s.T().Log("Before Test...")
}
func (s *ExampleSuite) AfterTest(suiteName, testName string) {
	s.T().Log("After Test...")
}

// o/p
// === RUN   TestExampleSuiteMain
//     examplesuite_test.go:28: Set-up Suite...
// === RUN   TestExampleSuiteMain/TestFalse
//     examplesuite_test.go:36: Set-up Test...
//     examplesuite_test.go:43: Before Test...
//     examplesuite_test.go:23: Test-False Executing...
//     examplesuite_test.go:46: After Test...
//     examplesuite_test.go:40: Tear-down Test...
// === RUN   TestExampleSuiteMain/TestTrue
//     examplesuite_test.go:36: Set-up Test...
//     examplesuite_test.go:43: Before Test...
//     examplesuite_test.go:18: Test-True Executing...
//     examplesuite_test.go:46: After Test...
//     examplesuite_test.go:40: Tear-down Test...
// === CONT  TestExampleSuiteMain
//     examplesuite_test.go:32: Tear-Down Suite...
// --- PASS: TestExampleSuiteMain (0.00s)
//     --- PASS: TestExampleSuiteMain/TestFalse (0.00s)
//     --- PASS: TestExampleSuiteMain/TestTrue (0.00s)
// PASS
// ok      suitesTesting   0.043s
