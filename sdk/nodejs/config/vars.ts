// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

let __config = new pulumi.Config("kind");

/**
 * Provider to use. Default: docker
 */
export let provider: string | undefined = __config.get("provider");
