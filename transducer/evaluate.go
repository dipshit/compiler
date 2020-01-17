package transducer

import (
	"fmt"
	"github.com/dipshit/compiler/ast"
	"github.com/dipshit/compiler/token"
	"strconv"
)

type TokenEvaluator func(*ast.TreeNode) int

var operatorMap = make(map[token.TokenType]TokenEvaluator)

func init() {
	operatorMap[token.PLUS] = evaluatePlus
	operatorMap[token.MULTIPLY] = evaluateMultiply
	operatorMap[token.INT] = evaluateInt
}

func evaluateTree(t *ast.TreeNode) int {
	return operatorMap[t.Token().Label()](t)
}

func evaluatePlus(t *ast.TreeNode) int {
	return evaluateTree(t.Children()[0]) + evaluateTree(t.Children()[1])
}

func evaluateMultiply(t *ast.TreeNode) int {
	return evaluateTree(t.Children()[0]) * evaluateTree(t.Children()[1])
}

// just convert token from string to int
func evaluateInt(t *ast.TreeNode) int {
	val, err := strconv.ParseInt(t.Token().Literal(), 10, 32)
	if err != nil {
		panic(fmt.Errorf("could evaluate token %s as int: %w", t.Token().Literal(), err))
	}
	return int(val)
}
