package main

import (
	"fmt"
	"strings"
)

var truthValues map[rune]trival

var exp logicalExpression

func start(n logicalExpression) (exp logicalExpression) {

	if n != "" {
		exp = n
		askValues(exp)
		return
	}
	fmt.Println("Enter a logical expression: ") // "(a&b)>c"  // "(a|b)>c"
	fmt.Scanln(&exp)
	askValues(exp)
	return
}

func main() {

	exp = start("(a&b)|c")
	a := parse(exp)

	printInorder(a, 0)
	a.setTruthValues()
	fmt.Printf("Given that values, the expression %v, is %v\n", exp, a.eval().String())

}

func parse(input logicalExpression) *Node {
	var depth int
	var a *Node
	var opSeenOutside int
	if input.isOnlyVar() {
		a = &Node{kind: LEAF, binOp: EMPTY, left: nil, right: nil, leafName: string(input)}
		return a
	}

	for i, c := range input {
		isOp, op := isOperation(c)
		if isOp && depth == 0 {
			opSeenOutside++
		}

		switch c {
		case ' ':
			continue
		case '(':
			depth++
		case ')':
			depth--

			if i == len(input)-1 && depth == 0 && opSeenOutside == 0 {
				s1 := input[1 : len(input)-1]
				return parse(s1)

			}

		default:
			if isOp && depth == 0 {

				lhs := input[:i]
				rhs := input[i+1:]
				a = &Node{binOp: op, kind: OPERATION, left: parse(lhs), right: parse(rhs)}
			} else {
				continue
			}
		}
	}
	return a
}

func isOperation(c rune) (bool, operationType) {
	a := map[rune]operationType{
		'&': AND,
		'|': OR,
		'=': IFF,
		'!': NOT,
		'>': IMPLIES,
	}
	return c == '&' || c == '|' || c == '=' || c == '>' || c == '!', a[c]
}

func (exp logicalExpression) isOnlyVar() bool {

	r := []rune(exp)
	if len(exp) == 1 && r[0] >= 'a' && r[0] <= 'z' {
		return true
	}
	return false
}

func (n *Node) eval() trival {
	switch n.kind {
	case LEAF:
		return n.truthValue
	case OPERATION:
		return n.doOp()
	default:
		return Undefined
	}
}

func (t trival) String() string {
	switch t {
	case True:
		return "True"
	case False:
		return "False"
	case Undefined:
		return "Undefined"
	}
	return "Undefined"
}

func (n *Node) doOp() (ans trival) {
	lhs, op, rhs := n.left.truthValue, n.binOp, n.right.truthValue

	switch op {
	case OR:
		ans = OrValues[lhs][rhs]
	case AND:
		ans = AndValues[lhs][rhs]
	case IFF:
		ans = IffValues[lhs][rhs]
	case IMPLIES:
		ans = ImpValues[lhs][rhs]

	}
	return
}

func askValues(input logicalExpression) {
	truthValues = make(map[rune]trival)
	for _, elem := range input {
		if elem >= 'a' && elem <= 'z' {
			truthValues[elem] = Undefined
		}
	}
	fmt.Println("Los valores validos son (T/F/U), insertar el valor de verdad de: ")
	for key := range truthValues {
		var in string
		var val trival
		fmt.Printf("variable %v\n", string(key))
		fmt.Scanln(&in)

		in = strings.ToLower(in)
		if in == "true" || in == "t" || in == "v" {
			val = True
		} else if in == "false" || in == "f" {
			val = False
		} else {
			val = Undefined
		}

		truthValues[key] = val

	}

}

func (n *Node) setTruthValues() {

	if n.kind == LEAF {
		n.truthValue = truthValues[[]rune(n.leafName)[0]]
	} else {
		n.left.setTruthValues()
		n.right.setTruthValues()
	}
}
