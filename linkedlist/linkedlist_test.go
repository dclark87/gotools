package linkedlist

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

// LinkedListTestSuite for running unit tests on linked list.
type LinkedListTestSuite struct {
	suite.Suite
	ll *List
}

// SetupTest initializes a new linked list to test on.
func (suite *LinkedListTestSuite) SetupTest() {
	suite.ll = NewList()

	suite.ll.Append(1)
	suite.ll.Append(2)
	suite.ll.Append(3)
}

// All methods that begin with "Test" are run as tests within a
// suite.
func (suite *LinkedListTestSuite) TestAppendAndPop() {
	suite.Equal(suite.ll.Pop(), 3)
	suite.Equal(suite.ll.Pop(), 2)
	suite.Equal(suite.ll.Pop(), 1)
}

func (suite *LinkedListTestSuite) TestPrint() {
	suite.ll.Print()
}

// In order for 'go test' to run this suite, we need to create
// a normal test function and pass our suite to suite.Run
func TestLinkedListTestSuite(t *testing.T) {
	suite.Run(t, new(LinkedListTestSuite))
}
