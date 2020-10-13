package main

import (
	"testing"

	"github.com/pintohutch/gotools/tree"
	"github.com/stretchr/testify/suite"
)

type InvertBinaryTreeSuite struct {
	suite.Suite
}

func (s *InvertBinaryTreeSuite) TestContainsTriplet() {
	var bt = &tree.BinaryRoot{
		Value: 'a',
	}
	bt.Left = &tree.BinaryRoot{
		Value: 'b',
		Left:  &tree.BinaryRoot{Value: 'd'},
		Right: &tree.BinaryRoot{Value: 'e'},
	}
	bt.Right = &tree.BinaryRoot{
		Value: 'c',
		Left:  &tree.BinaryRoot{Value: 'f'},
	}
	invertBinaryTree(bt)
	s.Equal('c', bt.Left.Value)
	s.Equal('b', bt.Right.Value)
	s.Nil(bt.Left.Left)
	s.Equal('f', bt.Left.Right.Value)
	s.Equal('e', bt.Right.Left.Value)
	s.Equal('d', bt.Right.Right.Value)
}

func TestInvertBinaryTreeSuite(t *testing.T) { suite.Run(t, new(InvertBinaryTreeSuite)) }
