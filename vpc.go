package awsvpc

import (
	"fmt"
	"strings"

	"github.com/pulumi/pulumi-aws/sdk/v4/go/aws/ec2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi/config"
)

// VpcOpts - Very much a work in progress
type VpcOpts struct {
	Name             string
	Cidr             string
	Azs              []string
	Publicsubnets    []string
	Privatesubnets   []string
	Singlenatgateway bool
	Tags             []string
}

func Vpc(ctx *pulumi.Context) error {
	var opts VpcOpts
	cfg := config.New(ctx, "")

	name := cfg.Require("name")
	cfg.RequireObject("vpc", &opts)

	// eh...
	opts.Name = name

	// Build VPC
	vpc, err := ec2.NewVpc(ctx, name, &ec2.VpcArgs{
		CidrBlock:          pulumi.String(opts.Cidr),
		EnableDnsHostnames: pulumi.Bool(true),
		EnableDnsSupport:   pulumi.Bool(true),
		Tags:               pulumiTags(name, opts.Tags),
	})
	if err != nil {
		return err
	}

	// Internet Gateway
	igw, err := ec2.NewInternetGateway(ctx, name+"-igw", &ec2.InternetGatewayArgs{
		VpcId: vpc.ID(),
		Tags:  pulumiTags(name+"-igw", opts.Tags),
	})
	if err != nil {
		return err
	}

	// Route Table
	_, err = ec2.NewRouteTable(ctx, name+"-rt", &ec2.RouteTableArgs{
		Routes: ec2.RouteTableRouteArray{
			ec2.RouteTableRouteArgs{
				CidrBlock: pulumi.String("0.0.0.0/0"),
				GatewayId: igw.ID(),
			},
		},
		VpcId: vpc.ID(),
		Tags:  pulumiTags(name+"-rt", opts.Tags),
	})
	if err != nil {
		return err
	}

	ctx.Export("vpc", vpc)

	_, err = opts.publicsubnets(ctx, vpc)
	if err != nil {
		return err
	}

	return nil
}

func (opts *VpcOpts) publicsubnets(ctx *pulumi.Context, vpc *ec2.Vpc) ([]*ec2.Subnet, error) {
	var subnets []*ec2.Subnet
	//var exportCidrs []string
	//var exportSubnetIds []string
	for k, s := range opts.Publicsubnets {
		name := fmt.Sprintf("%s-public-subnet-%s", opts.Name, opts.Azs[k])
		tags := pulumiTags(name, opts.Tags)
		subnet, err := ec2.NewSubnet(ctx, name, &ec2.SubnetArgs{
			VpcId:               vpc.ID(),
			CidrBlock:           pulumi.String(s),
			AvailabilityZone:    pulumi.String(opts.Azs[k]),
			MapPublicIpOnLaunch: pulumi.Bool(true),
			Tags:                tags,
		}, nil)
		if err != nil {
			return nil, err
		}

		// Append Subnets
		subnets = append(subnets, subnet)
		//exportCidrs = append(exportCidrs, s)
		//exportSubnetIds = append(exportSubnetIds, subnet.Id)
	}

	// Export Subnets
	//ctx.Export("publicSubnetCidrs", exportCidrs)
	//ctx.Export("publicSubnetIds", exportSubnetIds)
	return subnets, nil
}

func pulumiTags(name string, tags []string) pulumi.StringMap {
	retTags := pulumi.StringMap{}
	t := make(map[string]string)
	for _, l := range tags {
		x := strings.Split(l, ":")
		t[strings.TrimSpace(x[0])] = strings.TrimSpace(x[1])
	}
	for k, v := range t {
		retTags[k] = pulumi.String(v)
	}

	retTags["Name"] = pulumi.String(name)
	return retTags
}
