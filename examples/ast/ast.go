package main

import (
	"fmt"
	fiql "go-fiql/gofiql"
)

func main() {
	queries := []string{
		"(((((product==\"Apple\",product==\"Google\");(name==\"Joe\",name==\"Alan\")));label=!~=\"text\";(qty=gte=1,qty=lte=10)))",
		"(product==\"Apple\",product==\"Google\");(name==\"Joe\",name==\"Alan\");(qty=gte=1,qty=lte=10)",
		"(qty=gt=1;(qty=gte=1,qty=lte=10));(product==\"Apple\",product==\"HP\")",
		"(product==\"Apple\",qty=lt=1);name==\"Joe\"",
		"name==bar,dob=gt=1990-01-01",
		"title==foo*;(updated=lt=-P1D,title==*bar*)",
	}

	for _, query := range queries {
		fmt.Println(query)
		root, err := fiql.Parse(query)
		if err != nil {
			fmt.Println(err)
			return
		}

		fiql.PrettyPrinting(root, 0)
	}
}
