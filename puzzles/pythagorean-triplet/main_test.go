package main

import (
	"github.com/stretchr/testify/suite"
	"testing"
)

type PythagoreanTripletSuite struct {
	suite.Suite
}

func (s *PythagoreanTripletSuite) TestContainsTriplet() {
	var arr = []int{3, 4, 5, 6}
	s.True(containsTriplet(arr))
	arr = []int{3, 5, 6}
	s.False(containsTriplet(arr))
	arr = []int{160, 231, 19, 44, 281}
	s.True(containsTriplet(arr))
}

func TestPythagoreanTripletSuite(t *testing.T) { suite.Run(t, new(PythagoreanTripletSuite)) }
