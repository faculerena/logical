package main

import "fmt"

func printInorder(node *Node, depth int) {
	if node != nil {

		if node.kind == LEAF {
			fmt.Printf("%*s%+v\n", depth*4, "-", node)
			return
		}
		if node != nil {
			printInorder(node.left, depth+1)
			fmt.Printf("%*s%+v\n", depth*4, "-", node)
			printInorder(node.right, depth+1)
		}
	}
}
