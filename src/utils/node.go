package utils

import (
	"bytes"
	"github.com/agnivade/levenshtein"
	"golang.org/x/net/html"
)

func ToString(node *html.Node) string {
	var bytes bytes.Buffer
	html.Render(&bytes, node)

	return bytes.String()
}

func CompareNodes(aNode, anotherNode *html.Node) int {
	aNodeString := ToString(aNode)
	anotherNodeString := ToString(anotherNode)

	return levenshtein.ComputeDistance(aNodeString, anotherNodeString)
}