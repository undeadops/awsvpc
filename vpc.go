package awsvpc

import (
	"strings"

	"github.com/pulumi/pulumi-aws/sdk/v4/go/aws/ec2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi/config"
)

type VpcOpts struct {
	CidrBlock    string
	DnsHostnames bool
	DnsSupport   bool
	Tags         []string
}

func Vpc(ctx *pulumi.Context) error {
	var opts VpcOpts
	cfg := config.New(ctx, "")

	name := cfg.Require("name")
	cfg.RequireObject("vpc", &opts)

	// Render Tags
	tags := pulumi.StringMap{}
	for _, l := range opts.Tags {
		t := strings.Split(l, ":")
		tags[strings.TrimSpace(t[0])] = pulumi.String(strings.TrimSpace(t[1]))
	}

	// Set/Override Name Tag
	tags["Name"] = pulumi.String(name)

	// Build VPC
	vpc, err := ec2.NewVpc(ctx, name, &ec2.VpcArgs{
		CidrBlock:          pulumi.String(opts.CidrBlock),
		EnableDnsHostnames: pulumi.Bool(opts.DnsHostnames),
		EnableDnsSupport:   pulumi.Bool(opts.DnsSupport),
		Tags:               tags,
	})
	if err != nil {
		return err
	}

	ctx.Export("vpc", vpc)
	return nil
}
