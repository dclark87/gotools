package main

import (
	"github.com/stretchr/testify/suite"
	"testing"
)

type CoinHeadsSuite struct {
	suite.Suite
}

func (s *CoinHeadsSuite) TestNumRounds() {
	s.Equal(4, numRounds(10))
}

func (s *CoinHeadsSuite) TestNumRounds2() {
	s.Equal(3, numRounds2(10))
}

func TestCoinHeadsSuite(t *testing.T) { suite.Run(t, new(CoinHeadsSuite)) }
