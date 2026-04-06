package main

import (
	"fmt"
	"os"
)

func main() {
	if cat, err := newCat(os.Args); err != nil {
		fmt.Println(err);
		return;	
	} else if err = cat.run(); err != nil {
		fmt.Println(err);
		return;	
	}
}
