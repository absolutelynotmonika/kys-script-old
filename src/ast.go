package kysscript

import (
	"errors"
)

// N O D E
type NodeType string
var nodeTypes = map[NodeType]NodeType{
	"Program"    : "Program",
	"NumLiteral" : "NumLiteral",
	"BinaryExpr" : "BinaryExpr",
	"Identifier" : "Identifier",
}

func isValidNodeType(nodeType NodeType) error {
	if _, exists := nodeTypes[nodeType]; exists {
		return nil
	}
	return errors.New("Invalid node type")
}

// interface for setting the default Kind for nodes
type KindDefault interface {
	SetDefaults()
}

//// implement the nodes themselves

/// statements:
type Stmt struct {
	Kind NodeType
}
// program
type Program struct {
	Stmt // embed
	Body []Stmt
}
func (p *Program) SetDefaults() {
	p.Kind = nodeTypes["Program"]
}


/// expressions:
type Expr struct {
	Kind NodeType
}

// binary expressions
type BinaryExpr struct {
	Expr
	Left, Right Expr
	Operator TokenType
}
func (be *BinaryExpr) SetDefaults() {
	be.Kind = nodeTypes["BinaryExpr"]
}

// identifiers
type Identifier struct {
	Expr
	Symbol string
}
func (i *Identifier) SetDefaults() {
	i.Kind = nodeTypes["Identifier"]
}

// numeric literals
type NumLiteral struct {
	Expr
	Value int
}
func (nl *NumLiteral) SetDefaults() {
	nl.Kind = nodeTypes["NumLiteral"]
}
