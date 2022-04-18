# awsvpc
Pulumi Module to create an AWS VPC

# Example Usage

Example usage will be in the `example/` directory.

Start your pulumi project with:
`pulumi new aws-go`

I used the project name of `myInfra`.  Next you will need to setup some basic configuration values in your pulumi stack.

```bash
pulumi config set name dev # Matches stack name
pulumi config set --path 'vpc:cidrblock' 10.10.0.0/16
pulumi config set --path 'vpc:dnshostnames' true
pulumi config set --path 'vpc:dnssupport' true
pulumi config set --path 'vpc.azs[0]' "us-east-2a"
pulumi config set --path 'vpc.azs[1]' "us-east-2b"
pulumi config set --path 'vpc.azs[2]' "us-east-2c"
pulumi config set --path 'vpc.publicsubnets[0]' '10.20.0.0/24'
pulumi config set --path 'vpc.publicsubnets[1]' '10.20.1.0/24'
pulumi config set --path 'vpc.publicsubnets[2]' '10.20.2.0/24'
pulumi config set --path 'vpc.privatesubnets[0]' '10.20.32.0/19'
pulumi config set --path 'vpc.privatesubnets[1]' '10.20.64.0/19'
pulumi config set --path 'vpc.privatesubnets[2]' '10.20.96.0/19'
pulumi config set --path 'vpc.dbsubnets[0]' '10.20.20.0/24'
pulumi config set --path 'vpc.dbsubnets[1]' '10.20.22.0/24'
pulumi config set --path 'vpc.dbsubnets[2]' '10.20.24.0/24'
pulumi config set --path 'vpc:tags[0]' 'costcenter:platform'
pulumi config set --path 'vpc:singlenatgateway' true # false creates one nat gateway per private subnet
```

As long as your AWS credentials are setup and working.

`pulumi preview --diff` should return something that looks like:

```bash
~/code/myInfra ᐅ pulumi preview --diff
Previewing update (dev)

View Live: https://app.pulumi.com/undeadops/myInfra/dev/previews/

+ pulumi:pulumi:Stack: (create)
    [urn=urn:pulumi:dev::myInfra::pulumi:pulumi:Stack::myInfra-dev]
    + aws:ec2/vpc:Vpc: (create)
        [urn=urn:pulumi:dev::myInfra::aws:ec2/vpc:Vpc::dev]
        cidrBlock         : "10.20.0.0/16"
        enableDnsHostnames: true
        enableDnsSupport  : true
        instanceTenancy   : "default"
        tags              : {
            Name      : "dev"
            costcenter: "platform"
            env       : "dev"
        }
Resources:
    + 2 to create
```
