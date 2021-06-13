// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package patchjson6902

import (
	"fmt"

	"github.com/blang/semver"
	"github.com/frezbo/pulumi-provider-kind/sdk/v3/go/kind"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type module struct {
	version semver.Version
}

func (m *module) Version() semver.Version {
	return m.version
}

func (m *module) Construct(ctx *pulumi.Context, name, typ, urn string) (r pulumi.Resource, err error) {
	switch typ {
	case "kind:PatchJSON6902:PatchJSON6902":
		r = &PatchJSON6902{}
	default:
		return nil, fmt.Errorf("unknown resource type: %s", typ)
	}

	err = ctx.RegisterResource(typ, name, nil, r, pulumi.URN_(urn))
	return
}

func init() {
	version, err := kind.PkgVersion()
	if err != nil {
		fmt.Printf("failed to determine package version. defaulting to v1: %v\n", err)
	}
	pulumi.RegisterResourceModule(
		"kind",
		"PatchJSON6902",
		&module{version},
	)
}