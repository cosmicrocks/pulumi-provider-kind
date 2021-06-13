// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "../types";
import * as utilities from "../utilities";

/**
 * KIND Cluster
 */
export class Cluster extends pulumi.CustomResource {
    /**
     * Get an existing Cluster resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): Cluster {
        return new Cluster(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'kind:Cluster:Cluster';

    /**
     * Returns true if the given object is an instance of Cluster.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Cluster {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Cluster.__pulumiType;
    }

    public readonly apiVersion!: pulumi.Output<string | undefined>;
    public readonly containerdConfigPatches!: pulumi.Output<string[] | undefined>;
    public readonly containerdConfigPatchesJSON6902!: pulumi.Output<string[] | undefined>;
    public readonly featureGates!: pulumi.Output<{[key: string]: string} | undefined>;
    public readonly kind!: pulumi.Output<string | undefined>;
    public readonly kubeadmConfigPatches!: pulumi.Output<string[] | undefined>;
    public readonly kubeadmConfigPatchesJSON6902!: pulumi.Output<outputs.PatchJSON6902.PatchJSON6902[] | undefined>;
    public readonly name!: pulumi.Output<string | undefined>;
    public readonly networking!: pulumi.Output<outputs.Networking.Networking | undefined>;
    public readonly nodes!: pulumi.Output<outputs.Node.Node[] | undefined>;
    public readonly runtimeConfig!: pulumi.Output<{[key: string]: string} | undefined>;

    /**
     * Create a Cluster resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args?: ClusterArgs, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        opts = opts || {};
        if (!opts.id) {
            inputs["apiVersion"] = args ? args.apiVersion : undefined;
            inputs["containerdConfigPatches"] = args ? args.containerdConfigPatches : undefined;
            inputs["containerdConfigPatchesJSON6902"] = args ? args.containerdConfigPatchesJSON6902 : undefined;
            inputs["featureGates"] = args ? args.featureGates : undefined;
            inputs["kind"] = args ? args.kind : undefined;
            inputs["kubeadmConfigPatches"] = args ? args.kubeadmConfigPatches : undefined;
            inputs["kubeadmConfigPatchesJSON6902"] = args ? args.kubeadmConfigPatchesJSON6902 : undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["networking"] = args ? args.networking : undefined;
            inputs["nodes"] = args ? args.nodes : undefined;
            inputs["runtimeConfig"] = args ? args.runtimeConfig : undefined;
        } else {
            inputs["apiVersion"] = undefined /*out*/;
            inputs["containerdConfigPatches"] = undefined /*out*/;
            inputs["containerdConfigPatchesJSON6902"] = undefined /*out*/;
            inputs["featureGates"] = undefined /*out*/;
            inputs["kind"] = undefined /*out*/;
            inputs["kubeadmConfigPatches"] = undefined /*out*/;
            inputs["kubeadmConfigPatchesJSON6902"] = undefined /*out*/;
            inputs["name"] = undefined /*out*/;
            inputs["networking"] = undefined /*out*/;
            inputs["nodes"] = undefined /*out*/;
            inputs["runtimeConfig"] = undefined /*out*/;
        }
        if (!opts.version) {
            opts = pulumi.mergeOptions(opts, { version: utilities.getVersion()});
        }
        super(Cluster.__pulumiType, name, inputs, opts);
    }
}

/**
 * The set of arguments for constructing a Cluster resource.
 */
export interface ClusterArgs {
    apiVersion?: pulumi.Input<string>;
    containerdConfigPatches?: pulumi.Input<pulumi.Input<string>[]>;
    containerdConfigPatchesJSON6902?: pulumi.Input<pulumi.Input<string>[]>;
    featureGates?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    kind?: pulumi.Input<string>;
    kubeadmConfigPatches?: pulumi.Input<pulumi.Input<string>[]>;
    kubeadmConfigPatchesJSON6902?: pulumi.Input<pulumi.Input<inputs.PatchJSON6902.PatchJSON6902Args>[]>;
    name?: pulumi.Input<string>;
    networking?: pulumi.Input<inputs.Networking.NetworkingArgs>;
    nodes?: pulumi.Input<pulumi.Input<inputs.Node.NodeArgs>[]>;
    runtimeConfig?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
}
