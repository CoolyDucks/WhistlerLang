package main

type NodeKind int

const (
	NodeUnknown NodeKind = iota
	NodeSay
	NodeAssign
	NodeMath
	NodeMathBlock
	NodeRun
	NodeBuild
	NodeExec
	NodeTimePrint
	NodeTimeSet
	NodeIf
)

type Node struct {
	Kind     NodeKind
	Raw      string
	Expr     string
	Lines    []string
	Args     []string
	Left     string
	Right    string
	Children []*Node
}
