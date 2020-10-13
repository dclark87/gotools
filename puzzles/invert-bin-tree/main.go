package main

import (
	"fmt"

	"github.com/pintohutch/gotools/tree"
)

/**Invert a binary tree.

For example, given the following tree:

    a
   / \
  b   c
 / \  /
d   e f
should become:

  a
 / \
 c  b
 \  / \
  f e  d
  **/

func invertBinaryTree(root *tree.BinaryRoot) {
	// base-case: if leaf node or nil, return
	if root == nil || (root.Left == nil && root.Right == nil) {
		return
	}
	invertBinaryTree(root.Left)
	invertBinaryTree(root.Right)
	var tmp = root.Left
	root.Left = root.Right
	root.Right = tmp
}

func main() {
	fmt.Println("hello!")
}
