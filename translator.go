package main

import (
	"bufio"
	"fmt"
	"github.com/dipshit/compiler/ast"
	"github.com/dipshit/compiler/transducer"
	"log"
	"os"
)

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

func (s *SampleTranslator) evalWith(input string) (string, error) {
	transducer := transducer.NewTransducer()
	val, err := transducer.Evaluate(input)
	if err != nil {
		return "", fmt.Errorf("could not evaluage: %w", err)
	}
	return val, nil
}

func (s *SampleTranslator) promptForEvaluation() error {
	fmt.Print("Enter input: ")
	reader := bufio.NewReader(os.Stdin)
	toeval, err := reader.ReadString('\n')
	if err != nil {
		return fmt.Errorf("could not read from Stdin: %w", err)
	}
	val, err := s.evalWith(toeval)
	log.Printf("evaluated to: %s\n", val)
	return nil
}

func main() {
	sample := NewSampleTranslator()
	if err := sample.promptForEvaluation(); err != nil {
		log.Fatalf("could not evaluate: %w", err)
	}
}
