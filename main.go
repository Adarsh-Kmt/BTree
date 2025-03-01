package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {

	scanner := bufio.NewReader(os.Stdin)
	fmt.Println("B-TREE")
	fmt.Println("1) enter 'exit' to exit terminal.")
	fmt.Println("2) enter 'test-btree' to test working of b tree with maximum 4 items per node.")
	for {
		logger.Print("enter input: ")
		input, err := scanner.ReadString('\n')

		if err != nil {
			return
		}

		input = strings.TrimRight(input, "\r\n")
		logger.Println("input received : " + input)

		if input == "exit" {
			logger.Println("exiting...")
			return
		}

		if input == "test-btree" {

			logger.Println("testing b-tree...")
			bTree := BTreeInit(4)
			bTree.BtreeSetup()
			ok := true
			for ok {
				logger.Print("1) enter 'insert' to insert item    \n2) enter 'search' to search for item    \n3) enter 'delete' to delete an item    \n4) enter 'traverse' to traverse tree    \n5) enter 'exit' to exit")

				option, err := scanner.ReadString('\n')
				option = strings.TrimRight(option, "\r\n")

				if option == "exit" {
					ok = false
					continue
				}
				if err != nil {
					continue
				}

				if option == "traverse" {
					BTreeTraversal(bTree.Root)
					continue
				}
				logger.Print("enter key : ")
				key, err := scanner.ReadString('\n')
				key = strings.TrimRight(key, "\r\n")
				if err != nil {
					return
				}
				if option == "insert" {

					logger.Print("enter value : ")
					value, err := scanner.ReadString('\n')
					value = strings.TrimRight(value, "\r\n")
					if err != nil {
						return
					}
					bTree.InsertItem(key, value)

				} else if option == "search" {

					value, found := bTree.SearchForItem(key)
					if found {
						logger.Printf("value for key %s = %s", key, value)
					} else {
						logger.Printf("key %s not found in btree.", key)
					}
				} else if option == "delete" {

					value, err := bTree.DeleteItem(key)

					if err != nil {
						logger.Printf(err.Error())
					} else {
						logger.Printf("key %s with value = %s was deleted", key, value)

					}
				}

			}
		}

	}

}
