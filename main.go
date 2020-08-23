package main

import (
	"bytes"
	"fmt"
	"github.com/chermehdi/comet/parser"
)

const IndentWidth = 2

type PrintingVisitor struct {
	indent int
	buffer bytes.Buffer
}

func (p *PrintingVisitor) VisitStringLiteral(literal parser.StringLiteral) {
	p.printIndent()
	p.buffer.WriteString(fmt.Sprintf("StringLiteral(%s)\n", literal.Value))
}

func (p *PrintingVisitor) VisitIfStatement(statement parser.IfStatement) {
	p.printIndent()
	p.buffer.WriteString("IfStatement\n")
	p.indent += IndentWidth
	statement.Test.Accept(p)
	p.buffer.WriteString("(Then)")
	statement.Then.Accept(p)
	p.buffer.WriteString("(Else)")
	statement.Else.Accept(p)
	p.indent -= IndentWidth
}

func (p *PrintingVisitor) VisitBlockStatement(statement parser.BlockStatement) {
	p.printIndent()
	p.buffer.WriteString("BlockStatement\n")
	p.indent += IndentWidth
	for _, statement := range statement.Statements {
		statement.Accept(p)
	}
	p.indent -= IndentWidth
}

func (p *PrintingVisitor) printIndent() {
	for i := 0; i < p.indent; i++ {
		p.buffer.WriteRune(' ')
	}
}

func (p *PrintingVisitor) VisitExpression(parser.Expression) {
	panic("implement me")
}

func (p *PrintingVisitor) VisitStatement(statement parser.Statement) {
	panic("implement me")
}

func (p *PrintingVisitor) VisitRootNode(node parser.RootNode) {
	p.printIndent()
	p.buffer.WriteString("RootNode\n")
	p.indent += IndentWidth
	for _, st := range node.Statements {
		st.Accept(p)
	}
	p.indent -= IndentWidth
}

func (p *PrintingVisitor) VisitBinaryExpression(expression parser.BinaryExpression) {
	p.printIndent()
	p.buffer.WriteString(fmt.Sprintf("Visiting a BinaryExpression (%s) \n", expression.Op.Literal))
	p.indent += IndentWidth
	expression.Left.Accept(p)
	expression.Right.Accept(p)
	p.indent -= IndentWidth
}

func (p *PrintingVisitor) VisitPrefixExpression(expression parser.PrefixExpression) {
	p.printIndent()
	p.buffer.WriteString("PrefixExpression\n")
	p.indent += IndentWidth
	expression.Right.Accept(p)
	expression.Right.Accept(p)
	p.indent -= IndentWidth
}

func (p *PrintingVisitor) VisitNumberLiteral(expression parser.NumberLiteralExpression) {
	p.printIndent()
	p.buffer.WriteString(fmt.Sprintf("Visiting a Number (%d)\n", expression.ActualValue))
}

func (p *PrintingVisitor) VisitParenthesisedExpression(expression parser.ParenthesisedExpression) {
	p.printIndent()
	p.buffer.WriteString("ParenthesisedExpression\n")
	p.indent += IndentWidth
	expression.Expression.Accept(p)
	p.indent -= IndentWidth
}

func (p *PrintingVisitor) String() string {
	return p.buffer.String()
}

func (p *PrintingVisitor) VisitDeclarationStatement(statement parser.DeclarationStatement) {
	p.printIndent()
	p.buffer.WriteString(fmt.Sprintf("DeclarationStatement(%s)\n", statement.Identifier.Literal))
	p.indent += IndentWidth
	statement.Expression.Accept(p)
	p.indent -= IndentWidth
}

func (p *PrintingVisitor) VisitIdentifierExpression(expression parser.IdentifierExpression) {
	p.printIndent()
	p.buffer.WriteString(fmt.Sprintf("IdentifierExpression(%s)\n", expression.Name))
}

func (p *PrintingVisitor) VisitReturnStatement(statement parser.ReturnStatement) {
	p.printIndent()
	p.buffer.WriteString("ReturnStatement\n")
	p.indent += IndentWidth
	statement.Expression.Accept(p)
	p.indent -= IndentWidth
}

func (p *PrintingVisitor) VisitBooleanLiteral(literal parser.BooleanLiteral) {
	p.printIndent()
	p.buffer.WriteString(fmt.Sprintf("BooleanLiteral (%v)\n", literal.ActualValue))
}

func (p *PrintingVisitor) VisitFunctionStatement(statement parser.FunctionStatement) {
	p.printIndent()
	p.buffer.WriteString(fmt.Sprintf("FuncStatement %s\n", statement.Name))
	p.indent += IndentWidth
	p.printIndent()
	p.buffer.WriteString("Parameters: \n")
	for _, param := range statement.Parameters {
		param.Accept(p)
	}
	statement.Block.Accept(p)
	p.indent -= IndentWidth
}

func (p *PrintingVisitor) VisitCallExpression(expression parser.CallExpression) {
	p.printIndent()
	p.buffer.WriteString(fmt.Sprintf("CallExpression %s\n", expression.Name))
	p.indent += IndentWidth
	p.printIndent()
	p.buffer.WriteString("Parameters: \n")
	for _, arg := range expression.Arguments {
		arg.Accept(p)
	}
	p.indent -= IndentWidth
}

func main() {
	src := `
	var a = hello(10, 2)
	var c = goodBay("call", something(), a) 
`
	rootNode := parser.New(src).Parse()
	visitor := &PrintingVisitor{}
	rootNode.Accept(visitor)
	fmt.Println(visitor)
}
