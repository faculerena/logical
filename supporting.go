package main

type nodeType string

const (
	_         nodeType = "unasigned"
	OPERATION          = "operation"
	LEAF               = "leaf"
) // node types

type operationType string

const (
	EMPTY   operationType = ""
	OR                    = "|"
	AND                   = "&"
	NOT                   = "!"
	IFF                   = "="
	IMPLIES               = ">"
) // operation types

type Node struct {
	kind       nodeType
	binOp      operationType
	leafName   string
	isNegatedR bool
	isNegatedL bool
	truthValue trival
	left       *Node
	right      *Node
}

type logicalExpression string

type trival int

const (
	True trival = iota
	False
	Undefined
)

var OrValues = [3][3]trival{
	{True, True, True},
	{True, False, Undefined},
	{Undefined, Undefined, Undefined}}

var AndValues = [3][3]trival{
	{True, False, Undefined},
	{False, False, False},
	{Undefined, Undefined, Undefined}}

var ImpValues = [3][3]trival{
	{True, False, Undefined},
	{True, True, True},
	{Undefined, Undefined, Undefined}}

var IffValues = [3][3]trival{
	{True, False, Undefined},
	{False, True, Undefined},
	{Undefined, Undefined, Undefined}}

var NotValues = [3]trival{False, True, Undefined}
