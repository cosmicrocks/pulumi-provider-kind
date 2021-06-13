// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package patchjson6902

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// KIND PatchJSON6902
type PatchJSON6902Type struct {
	Group   *string `pulumi:"group"`
	Kind    *string `pulumi:"kind"`
	Patch   *string `pulumi:"patch"`
	Version *string `pulumi:"version"`
}

// PatchJSON6902TypeInput is an input type that accepts PatchJSON6902TypeArgs and PatchJSON6902TypeOutput values.
// You can construct a concrete instance of `PatchJSON6902TypeInput` via:
//
//          PatchJSON6902TypeArgs{...}
type PatchJSON6902TypeInput interface {
	pulumi.Input

	ToPatchJSON6902TypeOutput() PatchJSON6902TypeOutput
	ToPatchJSON6902TypeOutputWithContext(context.Context) PatchJSON6902TypeOutput
}

// KIND PatchJSON6902
type PatchJSON6902TypeArgs struct {
	Group   pulumi.StringPtrInput `pulumi:"group"`
	Kind    pulumi.StringPtrInput `pulumi:"kind"`
	Patch   pulumi.StringPtrInput `pulumi:"patch"`
	Version pulumi.StringPtrInput `pulumi:"version"`
}

func (PatchJSON6902TypeArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*PatchJSON6902Type)(nil)).Elem()
}

func (i PatchJSON6902TypeArgs) ToPatchJSON6902TypeOutput() PatchJSON6902TypeOutput {
	return i.ToPatchJSON6902TypeOutputWithContext(context.Background())
}

func (i PatchJSON6902TypeArgs) ToPatchJSON6902TypeOutputWithContext(ctx context.Context) PatchJSON6902TypeOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PatchJSON6902TypeOutput)
}

// PatchJSON6902TypeArrayInput is an input type that accepts PatchJSON6902TypeArray and PatchJSON6902TypeArrayOutput values.
// You can construct a concrete instance of `PatchJSON6902TypeArrayInput` via:
//
//          PatchJSON6902TypeArray{ PatchJSON6902TypeArgs{...} }
type PatchJSON6902TypeArrayInput interface {
	pulumi.Input

	ToPatchJSON6902TypeArrayOutput() PatchJSON6902TypeArrayOutput
	ToPatchJSON6902TypeArrayOutputWithContext(context.Context) PatchJSON6902TypeArrayOutput
}

type PatchJSON6902TypeArray []PatchJSON6902TypeInput

func (PatchJSON6902TypeArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]PatchJSON6902Type)(nil)).Elem()
}

func (i PatchJSON6902TypeArray) ToPatchJSON6902TypeArrayOutput() PatchJSON6902TypeArrayOutput {
	return i.ToPatchJSON6902TypeArrayOutputWithContext(context.Background())
}

func (i PatchJSON6902TypeArray) ToPatchJSON6902TypeArrayOutputWithContext(ctx context.Context) PatchJSON6902TypeArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PatchJSON6902TypeArrayOutput)
}

// KIND PatchJSON6902
type PatchJSON6902TypeOutput struct{ *pulumi.OutputState }

func (PatchJSON6902TypeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*PatchJSON6902Type)(nil)).Elem()
}

func (o PatchJSON6902TypeOutput) ToPatchJSON6902TypeOutput() PatchJSON6902TypeOutput {
	return o
}

func (o PatchJSON6902TypeOutput) ToPatchJSON6902TypeOutputWithContext(ctx context.Context) PatchJSON6902TypeOutput {
	return o
}

func (o PatchJSON6902TypeOutput) Group() pulumi.StringPtrOutput {
	return o.ApplyT(func(v PatchJSON6902Type) *string { return v.Group }).(pulumi.StringPtrOutput)
}

func (o PatchJSON6902TypeOutput) Kind() pulumi.StringPtrOutput {
	return o.ApplyT(func(v PatchJSON6902Type) *string { return v.Kind }).(pulumi.StringPtrOutput)
}

func (o PatchJSON6902TypeOutput) Patch() pulumi.StringPtrOutput {
	return o.ApplyT(func(v PatchJSON6902Type) *string { return v.Patch }).(pulumi.StringPtrOutput)
}

func (o PatchJSON6902TypeOutput) Version() pulumi.StringPtrOutput {
	return o.ApplyT(func(v PatchJSON6902Type) *string { return v.Version }).(pulumi.StringPtrOutput)
}

type PatchJSON6902TypeArrayOutput struct{ *pulumi.OutputState }

func (PatchJSON6902TypeArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]PatchJSON6902Type)(nil)).Elem()
}

func (o PatchJSON6902TypeArrayOutput) ToPatchJSON6902TypeArrayOutput() PatchJSON6902TypeArrayOutput {
	return o
}

func (o PatchJSON6902TypeArrayOutput) ToPatchJSON6902TypeArrayOutputWithContext(ctx context.Context) PatchJSON6902TypeArrayOutput {
	return o
}

func (o PatchJSON6902TypeArrayOutput) Index(i pulumi.IntInput) PatchJSON6902TypeOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) PatchJSON6902Type {
		return vs[0].([]PatchJSON6902Type)[vs[1].(int)]
	}).(PatchJSON6902TypeOutput)
}

func init() {
	pulumi.RegisterOutputType(PatchJSON6902TypeOutput{})
	pulumi.RegisterOutputType(PatchJSON6902TypeArrayOutput{})
}
