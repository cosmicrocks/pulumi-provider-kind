// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package node

import (
	"context"
	"reflect"

	"github.com/frezbo/pulumi-provider-kind/sdk/v3/go/kind/mount"
	"github.com/frezbo/pulumi-provider-kind/sdk/v3/go/kind/patchjson6902"
	"github.com/frezbo/pulumi-provider-kind/sdk/v3/go/kind/portmapping"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// KIND Node
type NodeType struct {
	ExtraMounts                  []mount.MountType                 `pulumi:"extraMounts"`
	ExtraPortMappings            []portmapping.PortMappingType     `pulumi:"extraPortMappings"`
	Image                        *string                           `pulumi:"image"`
	KubeadmConfigPatches         []string                          `pulumi:"kubeadmConfigPatches"`
	KubeadmConfigPatchesJSON6902 []patchjson6902.PatchJSON6902Type `pulumi:"kubeadmConfigPatchesJSON6902"`
	Labels                       map[string]string                 `pulumi:"labels"`
	Role                         *string                           `pulumi:"role"`
}

// NodeTypeInput is an input type that accepts NodeTypeArgs and NodeTypeOutput values.
// You can construct a concrete instance of `NodeTypeInput` via:
//
//          NodeTypeArgs{...}
type NodeTypeInput interface {
	pulumi.Input

	ToNodeTypeOutput() NodeTypeOutput
	ToNodeTypeOutputWithContext(context.Context) NodeTypeOutput
}

// KIND Node
type NodeTypeArgs struct {
	ExtraMounts                  mount.MountTypeArrayInput                 `pulumi:"extraMounts"`
	ExtraPortMappings            portmapping.PortMappingTypeArrayInput     `pulumi:"extraPortMappings"`
	Image                        pulumi.StringPtrInput                     `pulumi:"image"`
	KubeadmConfigPatches         pulumi.StringArrayInput                   `pulumi:"kubeadmConfigPatches"`
	KubeadmConfigPatchesJSON6902 patchjson6902.PatchJSON6902TypeArrayInput `pulumi:"kubeadmConfigPatchesJSON6902"`
	Labels                       pulumi.StringMapInput                     `pulumi:"labels"`
	Role                         pulumi.StringPtrInput                     `pulumi:"role"`
}

func (NodeTypeArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*NodeType)(nil)).Elem()
}

func (i NodeTypeArgs) ToNodeTypeOutput() NodeTypeOutput {
	return i.ToNodeTypeOutputWithContext(context.Background())
}

func (i NodeTypeArgs) ToNodeTypeOutputWithContext(ctx context.Context) NodeTypeOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NodeTypeOutput)
}

// NodeTypeArrayInput is an input type that accepts NodeTypeArray and NodeTypeArrayOutput values.
// You can construct a concrete instance of `NodeTypeArrayInput` via:
//
//          NodeTypeArray{ NodeTypeArgs{...} }
type NodeTypeArrayInput interface {
	pulumi.Input

	ToNodeTypeArrayOutput() NodeTypeArrayOutput
	ToNodeTypeArrayOutputWithContext(context.Context) NodeTypeArrayOutput
}

type NodeTypeArray []NodeTypeInput

func (NodeTypeArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]NodeType)(nil)).Elem()
}

func (i NodeTypeArray) ToNodeTypeArrayOutput() NodeTypeArrayOutput {
	return i.ToNodeTypeArrayOutputWithContext(context.Background())
}

func (i NodeTypeArray) ToNodeTypeArrayOutputWithContext(ctx context.Context) NodeTypeArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NodeTypeArrayOutput)
}

// KIND Node
type NodeTypeOutput struct{ *pulumi.OutputState }

func (NodeTypeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*NodeType)(nil)).Elem()
}

func (o NodeTypeOutput) ToNodeTypeOutput() NodeTypeOutput {
	return o
}

func (o NodeTypeOutput) ToNodeTypeOutputWithContext(ctx context.Context) NodeTypeOutput {
	return o
}

func (o NodeTypeOutput) ExtraMounts() mount.MountTypeArrayOutput {
	return o.ApplyT(func(v NodeType) []mount.MountType { return v.ExtraMounts }).(mount.MountTypeArrayOutput)
}

func (o NodeTypeOutput) ExtraPortMappings() portmapping.PortMappingTypeArrayOutput {
	return o.ApplyT(func(v NodeType) []portmapping.PortMappingType { return v.ExtraPortMappings }).(portmapping.PortMappingTypeArrayOutput)
}

func (o NodeTypeOutput) Image() pulumi.StringPtrOutput {
	return o.ApplyT(func(v NodeType) *string { return v.Image }).(pulumi.StringPtrOutput)
}

func (o NodeTypeOutput) KubeadmConfigPatches() pulumi.StringArrayOutput {
	return o.ApplyT(func(v NodeType) []string { return v.KubeadmConfigPatches }).(pulumi.StringArrayOutput)
}

func (o NodeTypeOutput) KubeadmConfigPatchesJSON6902() patchjson6902.PatchJSON6902TypeArrayOutput {
	return o.ApplyT(func(v NodeType) []patchjson6902.PatchJSON6902Type { return v.KubeadmConfigPatchesJSON6902 }).(patchjson6902.PatchJSON6902TypeArrayOutput)
}

func (o NodeTypeOutput) Labels() pulumi.StringMapOutput {
	return o.ApplyT(func(v NodeType) map[string]string { return v.Labels }).(pulumi.StringMapOutput)
}

func (o NodeTypeOutput) Role() pulumi.StringPtrOutput {
	return o.ApplyT(func(v NodeType) *string { return v.Role }).(pulumi.StringPtrOutput)
}

type NodeTypeArrayOutput struct{ *pulumi.OutputState }

func (NodeTypeArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]NodeType)(nil)).Elem()
}

func (o NodeTypeArrayOutput) ToNodeTypeArrayOutput() NodeTypeArrayOutput {
	return o
}

func (o NodeTypeArrayOutput) ToNodeTypeArrayOutputWithContext(ctx context.Context) NodeTypeArrayOutput {
	return o
}

func (o NodeTypeArrayOutput) Index(i pulumi.IntInput) NodeTypeOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) NodeType {
		return vs[0].([]NodeType)[vs[1].(int)]
	}).(NodeTypeOutput)
}

func init() {
	pulumi.RegisterOutputType(NodeTypeOutput{})
	pulumi.RegisterOutputType(NodeTypeArrayOutput{})
}
