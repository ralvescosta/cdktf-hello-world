package containers

import "github.com/ralvescosta/cdktf-hello-world/pkg/stack"

func NewEcsContainers(stack *stack.MyStack) {
	NewNginxContainer(stack)
}
