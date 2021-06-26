package service

import (
	"fmt"
	"golang.org/x/net/context"
)

func service(ctx context.Context)  {
	fmt.Println(ctx.Value("hah"));
}