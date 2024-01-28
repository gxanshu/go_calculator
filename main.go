package main

import (
	"bufio"
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"os"
	"strconv"
)

func Eval(exp ast.Expr) int {
	switch exp := exp.(type) {
	case *ast.BinaryExpr:
		return EvalBinaryExpr(exp)
	case *ast.BasicLit:
		switch exp.Kind {
		case token.INT:
			i, _ := strconv.Atoi(exp.Value)
			return i
		}
	}

	return 0
}

func EvalBinaryExpr(exp *ast.BinaryExpr) int {
	// fmt.Println(exp.X)
	left := Eval(exp.X)
	// fmt.Println(left)
	// fmt.Println(exp.Y)
	right := Eval(exp.Y)
	// fmt.Println(right)

	switch exp.Op {
	case token.ADD:
		return left + right
	case token.SUB:
		return left - right
	case token.MUL:
		return left * right
	case token.QUO:
		return left / right
	}

	return 0
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	const PROMPT string = "🤖 Exp >> "

	for {
		fmt.Print(PROMPT)
		scanned := scanner.Scan()

		if !scanned {
			return
		}

		line := scanner.Text()

		exp, err := parser.ParseExpr(line)

		if err != nil {
			return
		}

		fmt.Printf("👽 %d\n", Eval(exp))
	}
}
