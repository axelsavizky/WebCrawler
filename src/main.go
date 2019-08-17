package main

import (
	"agileEngine/src/utils"
	"fmt"
	"golang.org/x/net/html"
	"math"
	"os"
	"strconv"
)

func main() {
	originalFilePath := os.Args[1]
	diffFilePath := os.Args[2]
	idToFind := os.Args[3]

	if originalFilePath == "" || diffFilePath == "" || idToFind == "" {
		fmt.Println("invalid arguments")
		return
	}

	originalHtmlRootNode, err := utils.GetHtmlRootFromFile(originalFilePath)
	handleErr(err)
	if originalHtmlRootNode == nil {
		panic("malformed html")
	}
	element := searchElementByID(originalHtmlRootNode, idToFind)

	// Open the diff file
	diffHtmlRootNode, err := utils.GetHtmlRootFromFile(diffFilePath)
	handleErr(err)
	if diffHtmlRootNode == nil {
		panic("malformed html")
	}

	currentMinDiff := math.MaxInt32
	pathToSimilarElement := findElementWithMoreSimilarity(diffHtmlRootNode, "", "", 0, element, nil, &currentMinDiff, new(string))
	fmt.Println(pathToSimilarElement)
}

// TODO: Dont panic in an error!!
func handleErr(err error) {
	if err != nil {
		panic(err)
	}
}

// TODO: reduce the amount of arguments
func findElementWithMoreSimilarity(root *html.Node, currentPath, currentElementType string, currentNumberOfElement int, originalNode, currentMostSimilarElement *html.Node, currentMinDiff *int, currentPathToMostSimilarElement *string) string {
	// This must be done because the root doesn't has data
	if root.Data != "" {
		currentPath = currentPath + root.Data
		if currentNumberOfElement > 0 {
			currentPath = currentPath + "[" + strconv.Itoa(currentNumberOfElement) + "]"
		}
		currentPath = currentPath + " -> "
	}
	currentElementType = ""
	currentNumberOfElement = 0

	for child := root.FirstChild; child != nil; child = child.NextSibling {
		// Check similarity
		diff := utils.CompareNodes(originalNode, child)
		if diff < *currentMinDiff {
			*currentMinDiff = diff
			*currentPathToMostSimilarElement = currentPath + child.Data
			currentMostSimilarElement = child
		}

		// We need to keep tracking of the index of the current element in order to print something like div[i]
		// So we ignore the comments and the text like '\n'
		// TODO: improve the way that i count the index of the element
		if child.Type != html.CommentNode && child.Type != html.TextNode {
			if currentElementType == child.Data {
				currentNumberOfElement = currentNumberOfElement + 1
			} else {
				// Reset the counter!
				currentElementType = child.Data
				currentNumberOfElement = 0
			}
		}

		// Keep searching!
		if child.Type == html.ElementNode {
			*currentPathToMostSimilarElement = findElementWithMoreSimilarity(child, currentPath, currentElementType, currentNumberOfElement, originalNode, currentMostSimilarElement, currentMinDiff, currentPathToMostSimilarElement)
		}
	}

	return *currentPathToMostSimilarElement
}

func searchElementByID(root *html.Node, idToFind string) *html.Node {
	for child := root.FirstChild; child != nil; child = child.NextSibling {
		// Check if it was found
		idAttr := findIDAttribute(child)
		if idAttr != nil && idAttr.Val == idToFind {
			return child
		}

		// Keep searching!
		if child.Type == html.ElementNode {
			foundNode := searchElementByID(child, idToFind)
			if foundNode != nil {
				return foundNode
			}
		}
	}

	return nil
}

func findIDAttribute(element *html.Node) *html.Attribute {
	for _, attr := range element.Attr {
		if attr.Key == "id" {
			return &attr
		}
	}

	return nil
}
