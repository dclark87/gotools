package main

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type ClosestPointsSuite struct {
	suite.Suite
}

func (s *ClosestPointsSuite) TestContainsTriplet() {
	var coords = []coord{
		{x: 1, y: 1},
		{x: -1, y: -1},
		{x: 3, y: 4},
		{x: 6, y: 1},
		{x: -1, y: -6},
		{x: -4, y: -3},
	}

	var closest = closestPoints(coords)

	s.Contains(closest, coord{x: -1, y: -1})
	s.Contains(closest, coord{x: 1, y: 1})
	s.Len(closest, 2)
}

func TestClosestPointsSuite(t *testing.T) { suite.Run(t, new(ClosestPointsSuite)) }
