package pkg

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/ralvescosta/cdktf-hello-world/pkg/configs"
	"github.com/ralvescosta/cdktf-hello-world/pkg/ecr"
	"github.com/ralvescosta/cdktf-hello-world/pkg/ecs"
	"github.com/ralvescosta/cdktf-hello-world/pkg/ecs/containers"
	"github.com/ralvescosta/cdktf-hello-world/pkg/network"
	"github.com/ralvescosta/cdktf-hello-world/pkg/stack"
	"go.uber.org/zap"
)

func ApplyStack(logger *zap.SugaredLogger, cfgs *configs.Configs, tfStack cdktf.TerraformStack) {
	myStack := stack.MyStack{
		Cfgs:                   cfgs,
		Logger:                 logger,
		TfStack:                tfStack,
		Subnets:                &stack.Subnet{},
		NatGateways:            &stack.NatGateway{},
		RouteTables:            &stack.RouteTable{},
		PublicAppLoadBalancer:  &stack.ApplicationLoadBalancer{},
		PrivateAppLoadBalancer: &stack.ApplicationLoadBalancer{},
	}

	network.NewVpc(&myStack)
	network.NewSubnets(&myStack)
	network.NewInternetGateway(&myStack)
	network.NewNatGateway(&myStack)
	network.NewRouteTables(&myStack)
	network.NewPublicApplicationLoadBalancer(&myStack)
	ecr.NewECRRepositories(&myStack)
	ecs.NewECSFargateCluster(&myStack)
	containers.NewEcsContainers(&myStack)
}
