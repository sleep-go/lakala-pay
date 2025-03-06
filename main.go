package main

import (
	"fmt"
	"github.com/sleep-go/lakala-pay/api"
)

//TIP <p>To run your code, right-click the code and select <b>Run</b>.</p> <p>Alternatively, click
// the <icon src="AllIcons.Actions.Execute"/> icon in the gutter and select the <b>Run</b> menu item from here.</p>

func main() {
	client := api.NewClient("", "", "", "", true)
	fmt.Println(client)
}
