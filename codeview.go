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

	walkTree(tree.RootNode(), content, "")
}

func walkTree(node *sitter.Node, content []byte, parentType string) {
	nodeText := string(content[node.StartByte():node.EndByte()])
	nodeType := node.Type()

	// Construct a string that represents the node type and its parent type
	var fullNodeType string
	if parentType != "" {
		fullNodeType = parentType + "." + nodeType
	} else {
		fullNodeType = nodeType
	}

	// Check if the node is a leaf node (has no children)
	if node.ChildCount() == 0 {
		// First, try to color using the full node type
		if color, exists := config.ColorMap[fullNodeType]; exists {
			color.Print(nodeText)
		} else if color, exists := config.ColorMap[nodeType]; exists {
			// If that fails, try to color using just the node type
			color.Print(nodeText)
		} else {
			// If that also fails, print without color
			fmt.Print(nodeText)
		}
		return
	}

	// Handle non-leaf nodes
	lastByte := node.StartByte()
	for i := 0; i < int(node.ChildCount()); i++ {
		child := node.Child(i)
		// Print any text in the parent node that appears before this child
		if child.StartByte() > lastByte {
			fmt.Print(string(content[lastByte:child.StartByte()]))
		}
		// Recursively print the child
		walkTree(child, content, nodeType)
		// Update lastByte to be the byte after the end of the child
		lastByte = child.EndByte()
	}
	// Print any text in the parent node that appears after the last child
	if lastByte < node.EndByte() {
		fmt.Print(string(content[lastByte:node.EndByte()]))
	}
}
