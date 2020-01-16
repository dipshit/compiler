package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

type Table struct{}
type Transducer struct{}
type TableWithTransitions struct{}
type AcceptTable struct{}
type ScannerReadaheadTable struct{}
type ReadaheadTable struct{}
type ReadbackTable struct{}
type Screener struct{}
type SemanticTable struct{}
type ReduceTable struct{}
type Scanner struct{}
type Token struct{}
type ShiftbackTable struct{}

type SampleTranslator struct {
	parser *Parser
	tree   *TreeNode
	// codeIfCompiler expressionsIfEvaluator compilationOperatorMap evaluationOperatorMap
}

type Parser struct {
	//scanner screener tokenStack tableNumberStack treeStack left right tableNumber newTree
}

type TreeNode struct {
	nodes []*TreeNode
	label string
}

func NewSampleTranslator() *SampleTranslator {
	return &SampleTranslator{
		parser: NewParser(),
	}
}

func NewParser() *Parser {
	return &Parser{}
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
func (s *SampleTranslator) evaluateExpressionFor(tree *TreeNode) (string, error) {
	return "", nil
}

func (s *SampleTranslator) eval(tree *TreeNode) {
}

func main() {
	sample := &SampleTranslator{}
	if err := sample.promptForEvaluation(); err != nil {
		log.Fatalf("could not evaluate: %w", err)
	}
}

func (p *Parser) parse(text string) *TreeNode {
	return &TreeNode{}
}
