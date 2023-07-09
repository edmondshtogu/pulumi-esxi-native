// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "./types/input";
import * as outputs from "./types/output";
import * as enums from "./types/enums";
import * as utilities from "./utilities";

export class VirtualMachine extends pulumi.CustomResource {
    /**
     * Get an existing VirtualMachine resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): VirtualMachine {
        return new VirtualMachine(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'esxi-native:index:VirtualMachine';

    /**
     * Returns true if the given object is an instance of VirtualMachine.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is VirtualMachine {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === VirtualMachine.__pulumiType;
    }

    /**
     * VM boot disk size. Will expand boot disk to this size.
     */
    public readonly bootDiskSize!: pulumi.Output<number | undefined>;
    /**
     * VM boot disk type. thin, zeroedthick, eagerzeroedthick
     */
    public readonly bootDiskType!: pulumi.Output<enums.DiskType | undefined>;
    /**
     * Boot type('efi' is boot uefi mode)
     */
    public readonly bootFirmware!: pulumi.Output<enums.BootFirmwareType | undefined>;
    /**
     * esxi diskstore for boot disk.
     */
    public readonly diskStore!: pulumi.Output<string>;
    /**
     * pass data to VM
     */
    public readonly info!: pulumi.Output<outputs.KeyValuePair[] | undefined>;
    /**
     * The IP address reported by VMWare tools.
     */
    public /*out*/ readonly ipAddress!: pulumi.Output<string | undefined>;
    /**
     * VM memory size.
     */
    public readonly memSize!: pulumi.Output<number>;
    /**
     * esxi vm name.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * VM network interfaces.
     */
    public readonly networkInterfaces!: pulumi.Output<outputs.NetworkInterface[] | undefined>;
    /**
     * VM memory size.
     */
    public readonly notes!: pulumi.Output<string | undefined>;
    /**
     * VM number of virtual cpus.
     */
    public readonly numVCpus!: pulumi.Output<number>;
    /**
     * VM OS type.
     */
    public readonly os!: pulumi.Output<string>;
    /**
     * VM power state.
     */
    public readonly power!: pulumi.Output<string | undefined>;
    /**
     * Resource pool name to place vm.
     */
    public readonly resourcePoolName!: pulumi.Output<string>;
    /**
     * The amount of vm uptime, in seconds, to wait for an available IP address on this virtual machine.
     */
    public readonly shutdownTimeout!: pulumi.Output<number | undefined>;
    /**
     * The amount of vm uptime, in seconds, to wait for an available IP address on this virtual machine.
     */
    public readonly startupTimeout!: pulumi.Output<number | undefined>;
    /**
     * VM virtual disks.
     */
    public readonly virtualDisks!: pulumi.Output<outputs.VMVirtualDisk[] | undefined>;
    /**
     * VM Virtual HW version.
     */
    public readonly virtualHWVer!: pulumi.Output<number | undefined>;

    /**
     * Create a VirtualMachine resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: VirtualMachineArgs, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (!opts.id) {
            if ((!args || args.diskStore === undefined) && !opts.urn) {
                throw new Error("Missing required property 'diskStore'");
            }
            resourceInputs["bootDiskSize"] = (args ? args.bootDiskSize : undefined) ?? 16;
            resourceInputs["bootDiskType"] = (args ? args.bootDiskType : undefined) ?? "thin";
            resourceInputs["bootFirmware"] = (args ? args.bootFirmware : undefined) ?? "bios";
            resourceInputs["cloneFromVirtualMachine"] = args ? args.cloneFromVirtualMachine : undefined;
            resourceInputs["diskStore"] = args ? args.diskStore : undefined;
            resourceInputs["info"] = args ? args.info : undefined;
            resourceInputs["memSize"] = (args ? args.memSize : undefined) ?? 512;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["networkInterfaces"] = args ? args.networkInterfaces : undefined;
            resourceInputs["notes"] = args ? args.notes : undefined;
            resourceInputs["numVCpus"] = (args ? args.numVCpus : undefined) ?? 1;
            resourceInputs["os"] = (args ? args.os : undefined) ?? "centos";
            resourceInputs["ovfHostPathSource"] = args ? args.ovfHostPathSource : undefined;
            resourceInputs["ovfProperties"] = args ? args.ovfProperties : undefined;
            resourceInputs["ovfPropertiesTimer"] = (args ? args.ovfPropertiesTimer : undefined) ?? 6000;
            resourceInputs["ovfSourceLocalPath"] = args ? args.ovfSourceLocalPath : undefined;
            resourceInputs["power"] = args ? args.power : undefined;
            resourceInputs["resourcePoolName"] = (args ? args.resourcePoolName : undefined) ?? "/";
            resourceInputs["shutdownTimeout"] = (args ? args.shutdownTimeout : undefined) ?? 600;
            resourceInputs["startupTimeout"] = (args ? args.startupTimeout : undefined) ?? 600;
            resourceInputs["virtualDisks"] = args ? args.virtualDisks : undefined;
            resourceInputs["virtualHWVer"] = (args ? args.virtualHWVer : undefined) ?? 13;
            resourceInputs["ipAddress"] = undefined /*out*/;
        } else {
            resourceInputs["bootDiskSize"] = undefined /*out*/;
            resourceInputs["bootDiskType"] = undefined /*out*/;
            resourceInputs["bootFirmware"] = undefined /*out*/;
            resourceInputs["diskStore"] = undefined /*out*/;
            resourceInputs["info"] = undefined /*out*/;
            resourceInputs["ipAddress"] = undefined /*out*/;
            resourceInputs["memSize"] = undefined /*out*/;
            resourceInputs["name"] = undefined /*out*/;
            resourceInputs["networkInterfaces"] = undefined /*out*/;
            resourceInputs["notes"] = undefined /*out*/;
            resourceInputs["numVCpus"] = undefined /*out*/;
            resourceInputs["os"] = undefined /*out*/;
            resourceInputs["power"] = undefined /*out*/;
            resourceInputs["resourcePoolName"] = undefined /*out*/;
            resourceInputs["shutdownTimeout"] = undefined /*out*/;
            resourceInputs["startupTimeout"] = undefined /*out*/;
            resourceInputs["virtualDisks"] = undefined /*out*/;
            resourceInputs["virtualHWVer"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(VirtualMachine.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * The set of arguments for constructing a VirtualMachine resource.
 */
export interface VirtualMachineArgs {
    /**
     * VM boot disk size. Will expand boot disk to this size.
     */
    bootDiskSize?: pulumi.Input<number>;
    /**
     * VM boot disk type. thin, zeroedthick, eagerzeroedthick
     */
    bootDiskType?: pulumi.Input<enums.DiskType>;
    /**
     * Boot type('efi' is boot uefi mode)
     */
    bootFirmware?: pulumi.Input<enums.BootFirmwareType>;
    /**
     * Source vm path on esxi host to clone.
     */
    cloneFromVirtualMachine?: pulumi.Input<string>;
    /**
     * esxi diskstore for boot disk.
     */
    diskStore: pulumi.Input<string>;
    /**
     * pass data to VM
     */
    info?: pulumi.Input<pulumi.Input<inputs.KeyValuePairArgs>[]>;
    /**
     * VM memory size.
     */
    memSize?: pulumi.Input<number>;
    /**
     * esxi vm name.
     */
    name?: pulumi.Input<string>;
    /**
     * VM network interfaces.
     */
    networkInterfaces?: pulumi.Input<pulumi.Input<inputs.NetworkInterfaceArgs>[]>;
    /**
     * VM memory size.
     */
    notes?: pulumi.Input<string>;
    /**
     * VM number of virtual cpus.
     */
    numVCpus?: pulumi.Input<number>;
    /**
     * VM OS type.
     */
    os?: pulumi.Input<string>;
    /**
     * Path on esxi host of ovf files.
     */
    ovfHostPathSource?: pulumi.Input<string>;
    /**
     * VM OVF properties.
     */
    ovfProperties?: pulumi.Input<pulumi.Input<inputs.KeyValuePairArgs>[]>;
    /**
     * The amount of time, in seconds, to wait for the guest to boot and run ovfProperties. (0-6000)
     */
    ovfPropertiesTimer?: pulumi.Input<number>;
    /**
     * Local path to source ovf files.
     */
    ovfSourceLocalPath?: pulumi.Input<string>;
    /**
     * VM power state.
     */
    power?: pulumi.Input<string>;
    /**
     * Resource pool name to place vm.
     */
    resourcePoolName?: pulumi.Input<string>;
    /**
     * The amount of vm uptime, in seconds, to wait for an available IP address on this virtual machine. (0-600)
     */
    shutdownTimeout?: pulumi.Input<number>;
    /**
     * The amount of vm uptime, in seconds, to wait for an available IP address on this virtual machine. (0-600)
     */
    startupTimeout?: pulumi.Input<number>;
    /**
     * VM virtual disks.
     */
    virtualDisks?: pulumi.Input<pulumi.Input<inputs.VMVirtualDiskArgs>[]>;
    /**
     * VM Virtual HW version.
     */
    virtualHWVer?: pulumi.Input<number>;
}
