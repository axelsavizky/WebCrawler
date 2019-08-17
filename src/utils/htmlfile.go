package utils

import (
	"golang.org/x/net/html"
	"os"
)

const htmlRootData = "html"

func GetHtmlRootFromFile(filePath string) (*html.Node, error) {
	// Open the original file
	originalHtml, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}

	// Defer the close of the file
	defer originalHtml.Close()

	rootNode, err := html.Parse(originalHtml)
	if err != nil {
		return nil, err
	}

	return searchHtmlRoot(rootNode), nil
}

func searchHtmlRoot(rootNode *html.Node) *html.Node {
	for child := rootNode.FirstChild; child != nil; child = child.NextSibling {
		if child.Type == html.ElementNode && child.Data == htmlRootData {
			return child
		}
	}

	return nil
}
