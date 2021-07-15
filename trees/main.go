package main

import (
	"fmt"
	"github.com/evgeniyv6/c_algorithms/lesson6/trees"
	"log"
	"strings"
)

func main() {
	var newBst = new(trees.BinarySearchTree)
	newBst.Insert(8); newBst.Insert(4); newBst.Insert(10);newBst.Insert(2)
	newBst.Insert(6); newBst.Insert(1);newBst.Insert(3);newBst.Insert(5)
	newBst.Insert(7); newBst.Insert(9); newBst.Insert(11)

	fmt.Print("Inorder: ")
	newBst.Root.InOrderPrint()
	fmt.Println()
	fmt.Print("Preorder: ")
	newBst.Root.PreOrderPrint()
	fmt.Println()
	fmt.Print("Postorder: ")
	newBst.Root.PostOrderPrint()
	fmt.Println()
	var N = 8
	fmt.Printf("Check %d in tree: ", N)
	fmt.Println(newBst.Search(N))

	fmt.Println()
	fmt.Println("Print the tree:")
	newBst.String()
}
