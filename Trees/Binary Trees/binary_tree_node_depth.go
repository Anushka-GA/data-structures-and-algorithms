/*
  The distance between a node in a Binary Tree and the tree's root is called the
  node's depth.

  Write a function that takes in a Binary Tree and returns the sum of its nodes'
  depths. 
*/
package main

type BinaryTree struct {
	Value       int
	Left, Right *BinaryTree
}

func NodeDepths(root *BinaryTree) int {
	return nodeDepthHelper(root, 0)
	return -1
}

func nodeDepthHelper(root *BinaryTree, depth int) int {
    if root == nil {
        return 0
    }
    return depth + nodeDepthHelper(root.Left, depth + 1) + nodeDepthHelper(root.Right, depth + 1)
}