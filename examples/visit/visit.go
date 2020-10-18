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
		return
	}

	visitor := fiql.NewSQLVisitor()
	i, err := fiql.Traverse(root, visitor)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(i)
}
