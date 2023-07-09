// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.EsxiNative.Inputs
{

    public sealed class NetworkInterfaceArgs : global::Pulumi.ResourceArgs
    {
        [Input("macAddress")]
        public Input<string>? MacAddress { get; set; }

        [Input("nicType")]
        public Input<string>? NicType { get; set; }

        [Input("virtualNetwork", required: true)]
        public Input<string> VirtualNetwork { get; set; } = null!;

        public NetworkInterfaceArgs()
        {
        }
        public static new NetworkInterfaceArgs Empty => new NetworkInterfaceArgs();
    }
}
