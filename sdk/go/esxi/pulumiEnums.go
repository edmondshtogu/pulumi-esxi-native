// Code generated by pulumigen DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package esxi

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type BootDiskType string

const (
	BootDiskTypeThin             = BootDiskType("thin")
	BootDiskTypeZeroedThick      = BootDiskType("zeroedthick")
	BootDiskTypeEagerZeroedThick = BootDiskType("eagerzeroedthick")
)

func (BootDiskType) ElementType() reflect.Type {
	return reflect.TypeOf((*BootDiskType)(nil)).Elem()
}

func (e BootDiskType) ToBootDiskTypeOutput() BootDiskTypeOutput {
	return pulumi.ToOutput(e).(BootDiskTypeOutput)
}

func (e BootDiskType) ToBootDiskTypeOutputWithContext(ctx context.Context) BootDiskTypeOutput {
	return pulumi.ToOutputWithContext(ctx, e).(BootDiskTypeOutput)
}

func (e BootDiskType) ToBootDiskTypePtrOutput() BootDiskTypePtrOutput {
	return e.ToBootDiskTypePtrOutputWithContext(context.Background())
}

func (e BootDiskType) ToBootDiskTypePtrOutputWithContext(ctx context.Context) BootDiskTypePtrOutput {
	return BootDiskType(e).ToBootDiskTypeOutputWithContext(ctx).ToBootDiskTypePtrOutputWithContext(ctx)
}

func (e BootDiskType) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e BootDiskType) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e BootDiskType) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e BootDiskType) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type BootDiskTypeOutput struct{ *pulumi.OutputState }

func (BootDiskTypeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*BootDiskType)(nil)).Elem()
}

func (o BootDiskTypeOutput) ToBootDiskTypeOutput() BootDiskTypeOutput {
	return o
}

func (o BootDiskTypeOutput) ToBootDiskTypeOutputWithContext(ctx context.Context) BootDiskTypeOutput {
	return o
}

func (o BootDiskTypeOutput) ToBootDiskTypePtrOutput() BootDiskTypePtrOutput {
	return o.ToBootDiskTypePtrOutputWithContext(context.Background())
}

func (o BootDiskTypeOutput) ToBootDiskTypePtrOutputWithContext(ctx context.Context) BootDiskTypePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v BootDiskType) *BootDiskType {
		return &v
	}).(BootDiskTypePtrOutput)
}

func (o BootDiskTypeOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o BootDiskTypeOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e BootDiskType) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o BootDiskTypeOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o BootDiskTypeOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e BootDiskType) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type BootDiskTypePtrOutput struct{ *pulumi.OutputState }

func (BootDiskTypePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**BootDiskType)(nil)).Elem()
}

func (o BootDiskTypePtrOutput) ToBootDiskTypePtrOutput() BootDiskTypePtrOutput {
	return o
}

func (o BootDiskTypePtrOutput) ToBootDiskTypePtrOutputWithContext(ctx context.Context) BootDiskTypePtrOutput {
	return o
}

func (o BootDiskTypePtrOutput) Elem() BootDiskTypeOutput {
	return o.ApplyT(func(v *BootDiskType) BootDiskType {
		if v != nil {
			return *v
		}
		var ret BootDiskType
		return ret
	}).(BootDiskTypeOutput)
}

func (o BootDiskTypePtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o BootDiskTypePtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *BootDiskType) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}

// BootDiskTypeInput is an input type that accepts BootDiskTypeArgs and BootDiskTypeOutput values.
// You can construct a concrete instance of `BootDiskTypeInput` via:
//
//          BootDiskTypeArgs{...}
type BootDiskTypeInput interface {
	pulumi.Input

	ToBootDiskTypeOutput() BootDiskTypeOutput
	ToBootDiskTypeOutputWithContext(context.Context) BootDiskTypeOutput
}

var bootDiskTypePtrType = reflect.TypeOf((**BootDiskType)(nil)).Elem()

type BootDiskTypePtrInput interface {
	pulumi.Input

	ToBootDiskTypePtrOutput() BootDiskTypePtrOutput
	ToBootDiskTypePtrOutputWithContext(context.Context) BootDiskTypePtrOutput
}

type bootDiskTypePtr string

func BootDiskTypePtr(v string) BootDiskTypePtrInput {
	return (*bootDiskTypePtr)(&v)
}

func (*bootDiskTypePtr) ElementType() reflect.Type {
	return bootDiskTypePtrType
}

func (in *bootDiskTypePtr) ToBootDiskTypePtrOutput() BootDiskTypePtrOutput {
	return pulumi.ToOutput(in).(BootDiskTypePtrOutput)
}

func (in *bootDiskTypePtr) ToBootDiskTypePtrOutputWithContext(ctx context.Context) BootDiskTypePtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(BootDiskTypePtrOutput)
}

type BootFirmwareType string

const (
	BootFirmwareTypeBIOS = BootFirmwareType("bios")
	BootFirmwareTypeEFI  = BootFirmwareType("efi")
)

func (BootFirmwareType) ElementType() reflect.Type {
	return reflect.TypeOf((*BootFirmwareType)(nil)).Elem()
}

func (e BootFirmwareType) ToBootFirmwareTypeOutput() BootFirmwareTypeOutput {
	return pulumi.ToOutput(e).(BootFirmwareTypeOutput)
}

func (e BootFirmwareType) ToBootFirmwareTypeOutputWithContext(ctx context.Context) BootFirmwareTypeOutput {
	return pulumi.ToOutputWithContext(ctx, e).(BootFirmwareTypeOutput)
}

func (e BootFirmwareType) ToBootFirmwareTypePtrOutput() BootFirmwareTypePtrOutput {
	return e.ToBootFirmwareTypePtrOutputWithContext(context.Background())
}

func (e BootFirmwareType) ToBootFirmwareTypePtrOutputWithContext(ctx context.Context) BootFirmwareTypePtrOutput {
	return BootFirmwareType(e).ToBootFirmwareTypeOutputWithContext(ctx).ToBootFirmwareTypePtrOutputWithContext(ctx)
}

func (e BootFirmwareType) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e BootFirmwareType) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e BootFirmwareType) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e BootFirmwareType) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type BootFirmwareTypeOutput struct{ *pulumi.OutputState }

func (BootFirmwareTypeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*BootFirmwareType)(nil)).Elem()
}

func (o BootFirmwareTypeOutput) ToBootFirmwareTypeOutput() BootFirmwareTypeOutput {
	return o
}

func (o BootFirmwareTypeOutput) ToBootFirmwareTypeOutputWithContext(ctx context.Context) BootFirmwareTypeOutput {
	return o
}

func (o BootFirmwareTypeOutput) ToBootFirmwareTypePtrOutput() BootFirmwareTypePtrOutput {
	return o.ToBootFirmwareTypePtrOutputWithContext(context.Background())
}

func (o BootFirmwareTypeOutput) ToBootFirmwareTypePtrOutputWithContext(ctx context.Context) BootFirmwareTypePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v BootFirmwareType) *BootFirmwareType {
		return &v
	}).(BootFirmwareTypePtrOutput)
}

func (o BootFirmwareTypeOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o BootFirmwareTypeOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e BootFirmwareType) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o BootFirmwareTypeOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o BootFirmwareTypeOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e BootFirmwareType) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type BootFirmwareTypePtrOutput struct{ *pulumi.OutputState }

func (BootFirmwareTypePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**BootFirmwareType)(nil)).Elem()
}

func (o BootFirmwareTypePtrOutput) ToBootFirmwareTypePtrOutput() BootFirmwareTypePtrOutput {
	return o
}

func (o BootFirmwareTypePtrOutput) ToBootFirmwareTypePtrOutputWithContext(ctx context.Context) BootFirmwareTypePtrOutput {
	return o
}

func (o BootFirmwareTypePtrOutput) Elem() BootFirmwareTypeOutput {
	return o.ApplyT(func(v *BootFirmwareType) BootFirmwareType {
		if v != nil {
			return *v
		}
		var ret BootFirmwareType
		return ret
	}).(BootFirmwareTypeOutput)
}

func (o BootFirmwareTypePtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o BootFirmwareTypePtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *BootFirmwareType) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}

// BootFirmwareTypeInput is an input type that accepts BootFirmwareTypeArgs and BootFirmwareTypeOutput values.
// You can construct a concrete instance of `BootFirmwareTypeInput` via:
//
//          BootFirmwareTypeArgs{...}
type BootFirmwareTypeInput interface {
	pulumi.Input

	ToBootFirmwareTypeOutput() BootFirmwareTypeOutput
	ToBootFirmwareTypeOutputWithContext(context.Context) BootFirmwareTypeOutput
}

var bootFirmwareTypePtrType = reflect.TypeOf((**BootFirmwareType)(nil)).Elem()

type BootFirmwareTypePtrInput interface {
	pulumi.Input

	ToBootFirmwareTypePtrOutput() BootFirmwareTypePtrOutput
	ToBootFirmwareTypePtrOutputWithContext(context.Context) BootFirmwareTypePtrOutput
}

type bootFirmwareTypePtr string

func BootFirmwareTypePtr(v string) BootFirmwareTypePtrInput {
	return (*bootFirmwareTypePtr)(&v)
}

func (*bootFirmwareTypePtr) ElementType() reflect.Type {
	return bootFirmwareTypePtrType
}

func (in *bootFirmwareTypePtr) ToBootFirmwareTypePtrOutput() BootFirmwareTypePtrOutput {
	return pulumi.ToOutput(in).(BootFirmwareTypePtrOutput)
}

func (in *bootFirmwareTypePtr) ToBootFirmwareTypePtrOutputWithContext(ctx context.Context) BootFirmwareTypePtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(BootFirmwareTypePtrOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*BootDiskTypeInput)(nil)).Elem(), BootDiskType("thin"))
	pulumi.RegisterInputType(reflect.TypeOf((*BootDiskTypePtrInput)(nil)).Elem(), BootDiskType("thin"))
	pulumi.RegisterInputType(reflect.TypeOf((*BootFirmwareTypeInput)(nil)).Elem(), BootFirmwareType("bios"))
	pulumi.RegisterInputType(reflect.TypeOf((*BootFirmwareTypePtrInput)(nil)).Elem(), BootFirmwareType("bios"))
	pulumi.RegisterOutputType(BootDiskTypeOutput{})
	pulumi.RegisterOutputType(BootDiskTypePtrOutput{})
	pulumi.RegisterOutputType(BootFirmwareTypeOutput{})
	pulumi.RegisterOutputType(BootFirmwareTypePtrOutput{})
}
