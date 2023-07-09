// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.EsxiNative
{
    [EsxiNativeResourceType("esxi-native:index:VirtualMachine")]
    public partial class VirtualMachine : global::Pulumi.CustomResource
    {
        /// <summary>
        /// VM boot disk size. Will expand boot disk to this size.
        /// </summary>
        [Output("bootDiskSize")]
        public Output<int?> BootDiskSize { get; private set; } = null!;

        /// <summary>
        /// VM boot disk type. thin, zeroedthick, eagerzeroedthick
        /// </summary>
        [Output("bootDiskType")]
        public Output<Pulumi.EsxiNative.DiskType?> BootDiskType { get; private set; } = null!;

        /// <summary>
        /// Boot type('efi' is boot uefi mode)
        /// </summary>
        [Output("bootFirmware")]
        public Output<Pulumi.EsxiNative.BootFirmwareType?> BootFirmware { get; private set; } = null!;

        /// <summary>
        /// esxi diskstore for boot disk.
        /// </summary>
        [Output("diskStore")]
        public Output<string> DiskStore { get; private set; } = null!;

        /// <summary>
        /// pass data to VM
        /// </summary>
        [Output("info")]
        public Output<ImmutableArray<Outputs.KeyValuePair>> Info { get; private set; } = null!;

        /// <summary>
        /// The IP address reported by VMWare tools.
        /// </summary>
        [Output("ipAddress")]
        public Output<string?> IpAddress { get; private set; } = null!;

        /// <summary>
        /// VM memory size.
        /// </summary>
        [Output("memSize")]
        public Output<int> MemSize { get; private set; } = null!;

        /// <summary>
        /// esxi vm name.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// VM network interfaces.
        /// </summary>
        [Output("networkInterfaces")]
        public Output<ImmutableArray<Outputs.NetworkInterface>> NetworkInterfaces { get; private set; } = null!;

        /// <summary>
        /// VM memory size.
        /// </summary>
        [Output("notes")]
        public Output<string?> Notes { get; private set; } = null!;

        /// <summary>
        /// VM number of virtual cpus.
        /// </summary>
        [Output("numVCpus")]
        public Output<int> NumVCpus { get; private set; } = null!;

        /// <summary>
        /// VM OS type.
        /// </summary>
        [Output("os")]
        public Output<string> Os { get; private set; } = null!;

        /// <summary>
        /// VM power state.
        /// </summary>
        [Output("power")]
        public Output<string?> Power { get; private set; } = null!;

        /// <summary>
        /// Resource pool name to place vm.
        /// </summary>
        [Output("resourcePoolName")]
        public Output<string> ResourcePoolName { get; private set; } = null!;

        /// <summary>
        /// The amount of vm uptime, in seconds, to wait for an available IP address on this virtual machine.
        /// </summary>
        [Output("shutdownTimeout")]
        public Output<int?> ShutdownTimeout { get; private set; } = null!;

        /// <summary>
        /// The amount of vm uptime, in seconds, to wait for an available IP address on this virtual machine.
        /// </summary>
        [Output("startupTimeout")]
        public Output<int?> StartupTimeout { get; private set; } = null!;

        /// <summary>
        /// VM virtual disks.
        /// </summary>
        [Output("virtualDisks")]
        public Output<ImmutableArray<Outputs.VMVirtualDisk>> VirtualDisks { get; private set; } = null!;

        /// <summary>
        /// VM Virtual HW version.
        /// </summary>
        [Output("virtualHWVer")]
        public Output<int?> VirtualHWVer { get; private set; } = null!;


        /// <summary>
        /// Create a VirtualMachine resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public VirtualMachine(string name, VirtualMachineArgs args, CustomResourceOptions? options = null)
            : base("esxi-native:index:VirtualMachine", name, args ?? new VirtualMachineArgs(), MakeResourceOptions(options, ""))
        {
        }

        private VirtualMachine(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("esxi-native:index:VirtualMachine", name, null, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing VirtualMachine resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static VirtualMachine Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new VirtualMachine(name, id, options);
        }
    }

    public sealed class VirtualMachineArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// VM boot disk size. Will expand boot disk to this size.
        /// </summary>
        [Input("bootDiskSize")]
        public Input<int>? BootDiskSize { get; set; }

        /// <summary>
        /// VM boot disk type. thin, zeroedthick, eagerzeroedthick
        /// </summary>
        [Input("bootDiskType")]
        public Input<Pulumi.EsxiNative.DiskType>? BootDiskType { get; set; }

        /// <summary>
        /// Boot type('efi' is boot uefi mode)
        /// </summary>
        [Input("bootFirmware")]
        public Input<Pulumi.EsxiNative.BootFirmwareType>? BootFirmware { get; set; }

        /// <summary>
        /// Source vm path on esxi host to clone.
        /// </summary>
        [Input("cloneFromVirtualMachine")]
        public Input<string>? CloneFromVirtualMachine { get; set; }

        /// <summary>
        /// esxi diskstore for boot disk.
        /// </summary>
        [Input("diskStore", required: true)]
        public Input<string> DiskStore { get; set; } = null!;

        [Input("info")]
        private InputList<Inputs.KeyValuePairArgs>? _info;

        /// <summary>
        /// pass data to VM
        /// </summary>
        public InputList<Inputs.KeyValuePairArgs> Info
        {
            get => _info ?? (_info = new InputList<Inputs.KeyValuePairArgs>());
            set => _info = value;
        }

        /// <summary>
        /// VM memory size.
        /// </summary>
        [Input("memSize")]
        public Input<int>? MemSize { get; set; }

        /// <summary>
        /// esxi vm name.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("networkInterfaces")]
        private InputList<Inputs.NetworkInterfaceArgs>? _networkInterfaces;

        /// <summary>
        /// VM network interfaces.
        /// </summary>
        public InputList<Inputs.NetworkInterfaceArgs> NetworkInterfaces
        {
            get => _networkInterfaces ?? (_networkInterfaces = new InputList<Inputs.NetworkInterfaceArgs>());
            set => _networkInterfaces = value;
        }

        /// <summary>
        /// VM memory size.
        /// </summary>
        [Input("notes")]
        public Input<string>? Notes { get; set; }

        /// <summary>
        /// VM number of virtual cpus.
        /// </summary>
        [Input("numVCpus")]
        public Input<int>? NumVCpus { get; set; }

        /// <summary>
        /// VM OS type.
        /// </summary>
        [Input("os")]
        public Input<string>? Os { get; set; }

        /// <summary>
        /// Path on esxi host of ovf files.
        /// </summary>
        [Input("ovfHostPathSource")]
        public Input<string>? OvfHostPathSource { get; set; }

        [Input("ovfProperties")]
        private InputList<Inputs.KeyValuePairArgs>? _ovfProperties;

        /// <summary>
        /// VM OVF properties.
        /// </summary>
        public InputList<Inputs.KeyValuePairArgs> OvfProperties
        {
            get => _ovfProperties ?? (_ovfProperties = new InputList<Inputs.KeyValuePairArgs>());
            set => _ovfProperties = value;
        }

        /// <summary>
        /// The amount of time, in seconds, to wait for the guest to boot and run ovfProperties. (0-6000)
        /// </summary>
        [Input("ovfPropertiesTimer")]
        public Input<int>? OvfPropertiesTimer { get; set; }

        /// <summary>
        /// Local path to source ovf files.
        /// </summary>
        [Input("ovfSourceLocalPath")]
        public Input<string>? OvfSourceLocalPath { get; set; }

        /// <summary>
        /// VM power state.
        /// </summary>
        [Input("power")]
        public Input<string>? Power { get; set; }

        /// <summary>
        /// Resource pool name to place vm.
        /// </summary>
        [Input("resourcePoolName")]
        public Input<string>? ResourcePoolName { get; set; }

        /// <summary>
        /// The amount of vm uptime, in seconds, to wait for an available IP address on this virtual machine. (0-600)
        /// </summary>
        [Input("shutdownTimeout")]
        public Input<int>? ShutdownTimeout { get; set; }

        /// <summary>
        /// The amount of vm uptime, in seconds, to wait for an available IP address on this virtual machine. (0-600)
        /// </summary>
        [Input("startupTimeout")]
        public Input<int>? StartupTimeout { get; set; }

        [Input("virtualDisks")]
        private InputList<Inputs.VMVirtualDiskArgs>? _virtualDisks;

        /// <summary>
        /// VM virtual disks.
        /// </summary>
        public InputList<Inputs.VMVirtualDiskArgs> VirtualDisks
        {
            get => _virtualDisks ?? (_virtualDisks = new InputList<Inputs.VMVirtualDiskArgs>());
            set => _virtualDisks = value;
        }

        /// <summary>
        /// VM Virtual HW version.
        /// </summary>
        [Input("virtualHWVer")]
        public Input<int>? VirtualHWVer { get; set; }

        public VirtualMachineArgs()
        {
            BootDiskSize = 16;
            BootDiskType = Pulumi.EsxiNative.DiskType.Thin;
            BootFirmware = Pulumi.EsxiNative.BootFirmwareType.BIOS;
            MemSize = 512;
            NumVCpus = 1;
            Os = "centos";
            OvfPropertiesTimer = 6000;
            ResourcePoolName = "/";
            ShutdownTimeout = 600;
            StartupTimeout = 600;
            VirtualHWVer = 13;
        }
        public static new VirtualMachineArgs Empty => new VirtualMachineArgs();
    }
}
