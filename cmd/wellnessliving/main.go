package main

import (
	"context"
	"fmt"
	"net/http"

	"github.com/tekkamanendless/wellnessliving"
)

func main() {
	ctx := context.Background()

	client := wellnessliving.Client{}

	fmt.Printf("Events:\n")
	err := client.Raw(ctx, http.MethodGet, "/Wl/Event/Book/EventList/ListModel.json")
	if err != nil {
		fmt.Printf("Error: [%T] %v\n", err, err)
	}
}
