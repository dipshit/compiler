package main

import (
	"bufio"
	"fmt"
	"github.com/dipshit/compiler/ast"
	"github.com/dipshit/compiler/transducer"
	"log"
	"os"
)

type Token struct{}

type SampleTranslator struct {
	parser *transducer.Parser
	tree   *ast.TreeNode
	// codeIfCompiler expressionsIfEvaluator compilationOperatorMap evaluationOperatorMap
}

func NewSampleTranslator() *SampleTranslator {
	return &SampleTranslator{
		parser: transducer.NewParser(),
	}
}

func (s *SampleTranslator) promptForEvaluation() error {
	fmt.Print("Enter input: ")
	reader := bufio.NewReader(os.Stdin)
	toeval, err := reader.ReadString('\n')
	if err != nil {
		return fmt.Errorf("could not read from Stdin: %w", err)
	}
	val, err := s.evaluate(toeval)
	if err != nil {
		return fmt.Errorf("could not evaluage: %w", err)
	}
	log.Printf("evaluated to: %s\n", val)
	return nil
}

// evaluate creates a Parser which builds a tree from the text
// it then calls evaluateExpressionFor to get the result from the tree
func (s *SampleTranslator) evaluate(text string) (string, error) {
	result := "hi"
	return result, nil
}

// evaluate the ast and return a result
// wrap eval and recover from its panics
func (s *SampleTranslator) evaluateExpressionFor(tree *ast.TreeNode) (string, error) {
	return "", nil
}

func (s *SampleTranslator) eval(tree *ast.TreeNode) {
}

func main() {
	sample := NewSampleTranslator()
	if err := sample.promptForEvaluation(); err != nil {
		log.Fatalf("could not evaluate: %w", err)
	}
}
