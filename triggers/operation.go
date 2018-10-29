package triggers

import (
	"strings"
	"go/ast"
	"go/token"
	"go/parser"
	"strconv"
	"regexp"
	"fmt"
)

type Operation struct{}

func (operation Operation) ReplaceContent() string {
	// Escape if it isn't a message which contains a Country
	if !strings.Contains(Response, "${RESULT}") {
		return Response
	}

	literalOp := regexp.MustCompile("\\d(( )?[*+\\-/]( )?\\d)+")

	expr, err := parser.ParseExpr(literalOp.FindString(Entry))
	if err != nil {
		fmt.Println("Error parsing expression", err)
	}

	return strings.Replace(
		Response,
		"${RESULT}",
		strconv.FormatFloat(eval(expr), 'f', -1, 64),
		1)
}

func eval(expr ast.Expr) float64 {
	switch exp := expr.(type) {
	case *ast.ParenExpr:
		return eval(exp.X)
	case *ast.BinaryExpr:
		return evalBinary(exp)
	case *ast.BasicLit:
		switch exp.Kind {
		case token.INT:
			i, _ := strconv.Atoi(exp.Value)
			return float64(i)
		case token.FLOAT:
			i, _ := strconv.ParseFloat(exp.Value, 64)
			return i
		}
	}
	return 0
}

func evalBinary(expr *ast.BinaryExpr) float64 {
	left := eval(expr.X)
	right := eval(expr.Y)
	switch expr.Op {
	case token.ADD:
		return left + right
	case token.SUB:
		return left - right
	case token.MUL:
		return left * right
	case token.QUO:
		return left / right
	case token.REM:
		return float64(int(left) % int(right))
	}
	return 0
}
