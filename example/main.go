package main

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/undeadops/awsvpc"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		return awsvpc.Vpc(ctx)
	})
}
