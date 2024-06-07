package tree_sitter_kdl_test

import (
	"testing"

	tree_sitter "github.com/smacker/go-tree-sitter"
	"github.com/tree-sitter/tree-sitter-kdl"
)

func TestCanLoadGrammar(t *testing.T) {
	language := tree_sitter.NewLanguage(tree_sitter_kdl.Language())
	if language == nil {
		t.Errorf("Error loading Kdl grammar")
	}
}
