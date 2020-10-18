package main

import (
	"fmt"
	fiql "go-fiql/gofiql"
)

func main() {
	f := "(qty=gt=1;(qty=gte=1,qty=lte=10));(product==\"Apple\",product==\"HP\")"
	fmt.Println(f)
	root, err := fiql.Parse(f)
	if err != nil {
		fmt.Println(err)
	} else {
		fiql.PrettyPrinting(root, 0)
	}
}
