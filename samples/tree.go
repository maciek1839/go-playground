package samples

import (
	"golang.org/x/exp/slog"
	"strconv"
	"strings"
)

type Node struct {
	data  int
	left  *Node
	right *Node
}

type BST struct {
	root *Node
}

func Trees() {
	// https://blog.devgenius.io/trees-in-go-9b6ff346dcfc
	slog.Info("")
	slog.Info("======> Tree")
	slog.Info("A tree is a data structure made up of nodes or vertices and edges without having any cycle.")

	bst := BST{}
	bst.Insert(10)
	bst.Insert(5)
	bst.Insert(15)
	bst.Insert(20)
	bst.Insert(17)
	bst.Insert(4)
	bst.Insert(6)
	var order1 strings.Builder
	bst.Inorder(bst.root, &order1)
	slog.Info("Inorder result: " + order1.String())
	var order2 strings.Builder
	bst.Levelorder(&order2)
	slog.Info("Level order result: " + order2.String())
	slog.Info("Search result: " + strconv.FormatBool(bst.Search(5)))
}

func (bst *BST) Insert(val int) {
	bst.InsertRec(bst.root, val)
}

func (bst *BST) InsertRec(node *Node, val int) *Node {
	if bst.root == nil {
		bst.root = &Node{val, nil, nil}
		return bst.root
	}
	if node == nil {
		return &Node{val, nil, nil}
	}
	if val <= node.data {
		node.left = bst.InsertRec(node.left, val)
	}
	if val > node.data {
		node.right = bst.InsertRec(node.right, val)
	}
	return node
}
func (bst *BST) Search(val int) bool {
	found := bst.SearchRec(bst.root, val)
	return found
}
func (bst *BST) SearchRec(node *Node, val int) bool {
	if node.data == val {
		return true
	}
	if node == nil {
		return false
	}
	if val < node.data {
		return bst.SearchRec(node.left, val)
	}
	if val > node.data {
		return bst.SearchRec(node.right, val)
	}
	return false
}
func (bst *BST) Inorder(node *Node, order *strings.Builder) {
	if node == nil {
		return
	} else {
		bst.Inorder(node.left, order)
		order.WriteString(strconv.Itoa(node.data) + " ")
		bst.Inorder(node.right, order)
	}
}
func (bst *BST) Levelorder(sb *strings.Builder) {
	if bst.root == nil {
		return
	}
	nodeList := make([](*Node), 0)
	nodeList = append(nodeList, bst.root)
	for !(len(nodeList) == 0) {
		current := nodeList[0]
		sb.WriteString(strconv.Itoa(current.data) + " ")
		if current.left != nil {
			nodeList = append(nodeList, current.left)
		}
		if current.right != nil {
			nodeList = append(nodeList, current.right)
		}
		nodeList = nodeList[1:]
	}
}
