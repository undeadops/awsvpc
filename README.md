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
pulumi config set --path 'vpc:tags[0]' 'costcenter:platform'
pulumi config set --path 'vpc:tags[1]' 'env:dev' # matches stack name
```


