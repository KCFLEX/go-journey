package main

import (
	"context"
	"fmt"
)

func enrichContext(ctx context.Context) context.Context {
	return context.WithValue(ctx, "request-id", "12345")
}

func doSomethingCool(ctx context.Context) {
	rId := ctx.Value("request-id")
	fmt.Println(rId)
}
func main() {
	fmt.Println("Go Context Tutorial")
	ctx := context.Background()
	ctx = enrichContext(ctx)
	doSomethingCool(ctx)
}
