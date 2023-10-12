package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/blackflame007/codeview/config"
	sitter "github.com/smacker/go-tree-sitter"
	"github.com/smacker/go-tree-sitter/golang"
)

var languageMap = map[string]*sitter.Language{
	".go": golang.GetLanguage(),
}

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Usage: treesitter_highlight <file_path>")
		os.Exit(1)
	}

	filePath := os.Args[1]
	content, err := ioutil.ReadFile(filePath)
	if err != nil {
		fmt.Println("Error reading file:", err)
		os.Exit(1)
	}

	fileExt := filepath.Ext(filePath)
	language, exists := languageMap[fileExt]
	if !exists {
		fmt.Println("Unsupported file extension:", fileExt)
		os.Exit(1)
	}

	parser := sitter.NewParser()
	parser.SetLanguage(language)
	tree := parser.Parse(nil, content)

	walkTree(tree.RootNode(), content)
}

func walkTree(node *sitter.Node, content []byte) {
	nodeText := string(content[node.StartByte():node.EndByte()])
	nodeType := node.Type()

	// Check if the node is a leaf node (has no children)
	if node.ChildCount() == 0 {
		if color, exists := config.ColorMap[nodeType]; exists {
			color.Print(nodeText)
		} else {
			fmt.Print(nodeText)
		}
		return
	}

	// Handle non-leaf nodes by checking if the child nodes cover the entire text of the parent node
	childStart := node.Child(0).StartByte()
	childEnd := node.Child(int(node.ChildCount()) - 1).EndByte()

	if childStart == node.StartByte() && childEnd == node.EndByte() {
		// If the children cover the entire text of the parent, print the children instead
		for i := 0; i < int(node.ChildCount()); i++ {
			walkTree(node.Child(i), content)
		}
	} else {
		// If the children do not cover the entire text of the parent, print the parent node text
		if color, exists := config.ColorMap[nodeType]; exists {
			color.Print(nodeText)
		} else {
			fmt.Print(nodeText)
		}
	}
}
