// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package node

import (
	"context"
	"reflect"

	mount "github.com/frezbo/pulumi-provider-kind/sdk/v3/go/kind/mount"
	patchjson6902 "github.com/frezbo/pulumi-provider-kind/sdk/v3/go/kind/patchjson6902"
	portmapping "github.com/frezbo/pulumi-provider-kind/sdk/v3/go/kind/portmapping"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// KIND Node type
type Node struct {
	ExtraMounts                  []mount.Mount                 `pulumi:"extraMounts"`
	ExtraPortMappings            []portmapping.PortMapping     `pulumi:"extraPortMappings"`
	Image                        *string                       `pulumi:"image"`
	KubeadmConfigPatches         []string                      `pulumi:"kubeadmConfigPatches"`
	KubeadmConfigPatchesJSON6902 []patchjson6902.PatchJSON6902 `pulumi:"kubeadmConfigPatchesJSON6902"`
	Labels                       map[string]string             `pulumi:"labels"`
	// node role type
	Role *string `pulumi:"role"`
}

// NodeInput is an input type that accepts NodeArgs and NodeOutput values.
// You can construct a concrete instance of `NodeInput` via:
//
//          NodeArgs{...}
type NodeInput interface {
	pulumi.Input

	ToNodeOutput() NodeOutput
	ToNodeOutputWithContext(context.Context) NodeOutput
}

// KIND Node type
type NodeArgs struct {
	ExtraMounts                  mount.MountArrayInput                 `pulumi:"extraMounts"`
	ExtraPortMappings            portmapping.PortMappingArrayInput     `pulumi:"extraPortMappings"`
	Image                        pulumi.StringPtrInput                 `pulumi:"image"`
	KubeadmConfigPatches         pulumi.StringArrayInput               `pulumi:"kubeadmConfigPatches"`
	KubeadmConfigPatchesJSON6902 patchjson6902.PatchJSON6902ArrayInput `pulumi:"kubeadmConfigPatchesJSON6902"`
	Labels                       pulumi.StringMapInput                 `pulumi:"labels"`
	// node role type
	Role pulumi.StringPtrInput `pulumi:"role"`
}

func (NodeArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*Node)(nil)).Elem()
}

func (i NodeArgs) ToNodeOutput() NodeOutput {
	return i.ToNodeOutputWithContext(context.Background())
}

func (i NodeArgs) ToNodeOutputWithContext(ctx context.Context) NodeOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NodeOutput)
}

// NodeArrayInput is an input type that accepts NodeArray and NodeArrayOutput values.
// You can construct a concrete instance of `NodeArrayInput` via:
//
//          NodeArray{ NodeArgs{...} }
type NodeArrayInput interface {
	pulumi.Input

	ToNodeArrayOutput() NodeArrayOutput
	ToNodeArrayOutputWithContext(context.Context) NodeArrayOutput
}

type NodeArray []NodeInput

func (NodeArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]Node)(nil)).Elem()
}

func (i NodeArray) ToNodeArrayOutput() NodeArrayOutput {
	return i.ToNodeArrayOutputWithContext(context.Background())
}

func (i NodeArray) ToNodeArrayOutputWithContext(ctx context.Context) NodeArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NodeArrayOutput)
}

// KIND Node type
type NodeOutput struct{ *pulumi.OutputState }

func (NodeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*Node)(nil)).Elem()
}

func (o NodeOutput) ToNodeOutput() NodeOutput {
	return o
}

func (o NodeOutput) ToNodeOutputWithContext(ctx context.Context) NodeOutput {
	return o
}

func (o NodeOutput) ExtraMounts() mount.MountArrayOutput {
	return o.ApplyT(func(v Node) []mount.Mount { return v.ExtraMounts }).(mount.MountArrayOutput)
}

func (o NodeOutput) ExtraPortMappings() portmapping.PortMappingArrayOutput {
	return o.ApplyT(func(v Node) []portmapping.PortMapping { return v.ExtraPortMappings }).(portmapping.PortMappingArrayOutput)
}

func (o NodeOutput) Image() pulumi.StringPtrOutput {
	return o.ApplyT(func(v Node) *string { return v.Image }).(pulumi.StringPtrOutput)
}

func (o NodeOutput) KubeadmConfigPatches() pulumi.StringArrayOutput {
	return o.ApplyT(func(v Node) []string { return v.KubeadmConfigPatches }).(pulumi.StringArrayOutput)
}

func (o NodeOutput) KubeadmConfigPatchesJSON6902() patchjson6902.PatchJSON6902ArrayOutput {
	return o.ApplyT(func(v Node) []patchjson6902.PatchJSON6902 { return v.KubeadmConfigPatchesJSON6902 }).(patchjson6902.PatchJSON6902ArrayOutput)
}

func (o NodeOutput) Labels() pulumi.StringMapOutput {
	return o.ApplyT(func(v Node) map[string]string { return v.Labels }).(pulumi.StringMapOutput)
}

// node role type
func (o NodeOutput) Role() pulumi.StringPtrOutput {
	return o.ApplyT(func(v Node) *string { return v.Role }).(pulumi.StringPtrOutput)
}

type NodeArrayOutput struct{ *pulumi.OutputState }

func (NodeArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]Node)(nil)).Elem()
}

func (o NodeArrayOutput) ToNodeArrayOutput() NodeArrayOutput {
	return o
}

func (o NodeArrayOutput) ToNodeArrayOutputWithContext(ctx context.Context) NodeArrayOutput {
	return o
}

func (o NodeArrayOutput) Index(i pulumi.IntInput) NodeOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) Node {
		return vs[0].([]Node)[vs[1].(int)]
	}).(NodeOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*NodeInput)(nil)).Elem(), NodeArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*NodeArrayInput)(nil)).Elem(), NodeArray{})
	pulumi.RegisterOutputType(NodeOutput{})
	pulumi.RegisterOutputType(NodeArrayOutput{})
}
