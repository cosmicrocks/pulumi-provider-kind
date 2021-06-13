// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package mount

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// KIND Mount
type Mount struct {
	pulumi.CustomResourceState

	ContainerPath  pulumi.StringPtrOutput `pulumi:"containerPath"`
	HostPath       pulumi.StringPtrOutput `pulumi:"hostPath"`
	Propagation    pulumi.StringPtrOutput `pulumi:"propagation"`
	ReadOnly       pulumi.BoolPtrOutput   `pulumi:"readOnly"`
	SelinuxRelabel pulumi.BoolPtrOutput   `pulumi:"selinuxRelabel"`
}

// NewMount registers a new resource with the given unique name, arguments, and options.
func NewMount(ctx *pulumi.Context,
	name string, args *MountArgs, opts ...pulumi.ResourceOption) (*Mount, error) {
	if args == nil {
		args = &MountArgs{}
	}

	var resource Mount
	err := ctx.RegisterResource("kind:Mount:Mount", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetMount gets an existing Mount resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetMount(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *MountState, opts ...pulumi.ResourceOption) (*Mount, error) {
	var resource Mount
	err := ctx.ReadResource("kind:Mount:Mount", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Mount resources.
type mountState struct {
	ContainerPath  *string `pulumi:"containerPath"`
	HostPath       *string `pulumi:"hostPath"`
	Propagation    *string `pulumi:"propagation"`
	ReadOnly       *bool   `pulumi:"readOnly"`
	SelinuxRelabel *bool   `pulumi:"selinuxRelabel"`
}

type MountState struct {
	ContainerPath  pulumi.StringPtrInput
	HostPath       pulumi.StringPtrInput
	Propagation    pulumi.StringPtrInput
	ReadOnly       pulumi.BoolPtrInput
	SelinuxRelabel pulumi.BoolPtrInput
}

func (MountState) ElementType() reflect.Type {
	return reflect.TypeOf((*mountState)(nil)).Elem()
}

type mountArgs struct {
	ContainerPath  *string `pulumi:"containerPath"`
	HostPath       *string `pulumi:"hostPath"`
	Propagation    *string `pulumi:"propagation"`
	ReadOnly       *bool   `pulumi:"readOnly"`
	SelinuxRelabel *bool   `pulumi:"selinuxRelabel"`
}

// The set of arguments for constructing a Mount resource.
type MountArgs struct {
	ContainerPath  pulumi.StringPtrInput
	HostPath       pulumi.StringPtrInput
	Propagation    pulumi.StringPtrInput
	ReadOnly       pulumi.BoolPtrInput
	SelinuxRelabel pulumi.BoolPtrInput
}

func (MountArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*mountArgs)(nil)).Elem()
}

type MountInput interface {
	pulumi.Input

	ToMountOutput() MountOutput
	ToMountOutputWithContext(ctx context.Context) MountOutput
}

func (*Mount) ElementType() reflect.Type {
	return reflect.TypeOf((*Mount)(nil))
}

func (i *Mount) ToMountOutput() MountOutput {
	return i.ToMountOutputWithContext(context.Background())
}

func (i *Mount) ToMountOutputWithContext(ctx context.Context) MountOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MountOutput)
}

func (i *Mount) ToMountPtrOutput() MountPtrOutput {
	return i.ToMountPtrOutputWithContext(context.Background())
}

func (i *Mount) ToMountPtrOutputWithContext(ctx context.Context) MountPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MountPtrOutput)
}

type MountPtrInput interface {
	pulumi.Input

	ToMountPtrOutput() MountPtrOutput
	ToMountPtrOutputWithContext(ctx context.Context) MountPtrOutput
}

type mountPtrType MountArgs

func (*mountPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**Mount)(nil))
}

func (i *mountPtrType) ToMountPtrOutput() MountPtrOutput {
	return i.ToMountPtrOutputWithContext(context.Background())
}

func (i *mountPtrType) ToMountPtrOutputWithContext(ctx context.Context) MountPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MountPtrOutput)
}

// MountArrayInput is an input type that accepts MountArray and MountArrayOutput values.
// You can construct a concrete instance of `MountArrayInput` via:
//
//          MountArray{ MountArgs{...} }
type MountArrayInput interface {
	pulumi.Input

	ToMountArrayOutput() MountArrayOutput
	ToMountArrayOutputWithContext(context.Context) MountArrayOutput
}

type MountArray []MountInput

func (MountArray) ElementType() reflect.Type {
	return reflect.TypeOf(([]*Mount)(nil))
}

func (i MountArray) ToMountArrayOutput() MountArrayOutput {
	return i.ToMountArrayOutputWithContext(context.Background())
}

func (i MountArray) ToMountArrayOutputWithContext(ctx context.Context) MountArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MountArrayOutput)
}

// MountMapInput is an input type that accepts MountMap and MountMapOutput values.
// You can construct a concrete instance of `MountMapInput` via:
//
//          MountMap{ "key": MountArgs{...} }
type MountMapInput interface {
	pulumi.Input

	ToMountMapOutput() MountMapOutput
	ToMountMapOutputWithContext(context.Context) MountMapOutput
}

type MountMap map[string]MountInput

func (MountMap) ElementType() reflect.Type {
	return reflect.TypeOf((map[string]*Mount)(nil))
}

func (i MountMap) ToMountMapOutput() MountMapOutput {
	return i.ToMountMapOutputWithContext(context.Background())
}

func (i MountMap) ToMountMapOutputWithContext(ctx context.Context) MountMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MountMapOutput)
}

type MountOutput struct {
	*pulumi.OutputState
}

func (MountOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*Mount)(nil))
}

func (o MountOutput) ToMountOutput() MountOutput {
	return o
}

func (o MountOutput) ToMountOutputWithContext(ctx context.Context) MountOutput {
	return o
}

func (o MountOutput) ToMountPtrOutput() MountPtrOutput {
	return o.ToMountPtrOutputWithContext(context.Background())
}

func (o MountOutput) ToMountPtrOutputWithContext(ctx context.Context) MountPtrOutput {
	return o.ApplyT(func(v Mount) *Mount {
		return &v
	}).(MountPtrOutput)
}

type MountPtrOutput struct {
	*pulumi.OutputState
}

func (MountPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Mount)(nil))
}

func (o MountPtrOutput) ToMountPtrOutput() MountPtrOutput {
	return o
}

func (o MountPtrOutput) ToMountPtrOutputWithContext(ctx context.Context) MountPtrOutput {
	return o
}

type MountArrayOutput struct{ *pulumi.OutputState }

func (MountArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]Mount)(nil))
}

func (o MountArrayOutput) ToMountArrayOutput() MountArrayOutput {
	return o
}

func (o MountArrayOutput) ToMountArrayOutputWithContext(ctx context.Context) MountArrayOutput {
	return o
}

func (o MountArrayOutput) Index(i pulumi.IntInput) MountOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) Mount {
		return vs[0].([]Mount)[vs[1].(int)]
	}).(MountOutput)
}

type MountMapOutput struct{ *pulumi.OutputState }

func (MountMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]Mount)(nil))
}

func (o MountMapOutput) ToMountMapOutput() MountMapOutput {
	return o
}

func (o MountMapOutput) ToMountMapOutputWithContext(ctx context.Context) MountMapOutput {
	return o
}

func (o MountMapOutput) MapIndex(k pulumi.StringInput) MountOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) Mount {
		return vs[0].(map[string]Mount)[vs[1].(string)]
	}).(MountOutput)
}

func init() {
	pulumi.RegisterOutputType(MountOutput{})
	pulumi.RegisterOutputType(MountPtrOutput{})
	pulumi.RegisterOutputType(MountArrayOutput{})
	pulumi.RegisterOutputType(MountMapOutput{})
}
