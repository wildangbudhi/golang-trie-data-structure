package main

import (
	"fmt"

	"github.com/wildangbudhi/golang-trie-data-structure/utils"
)

func main() {

	var phoneCodeList []string = []string{"62", "628", "283"}
	var phoneCodeDict *utils.Trie = utils.NewTrie()

	for i := 0; i < len(phoneCodeList); i++ {
		phoneCodeDict.Insert(phoneCodeList[i])
	}

	fmt.Println(phoneCodeDict)
	fmt.Println(phoneCodeDict.ValidateStringPrefix("2831231"))

}
