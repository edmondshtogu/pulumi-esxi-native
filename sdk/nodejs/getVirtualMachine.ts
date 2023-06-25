// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs, enums } from "./types";
import * as utilities from "./utilities";

export function getVirtualMachine(args: GetVirtualMachineArgs, opts?: pulumi.InvokeOptions): Promise<GetVirtualMachineResult> {
    if (!opts) {
        opts = {}
    }

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
    return pulumi.runtime.invoke("esxi-native:index:getVirtualMachine", {
        "name": args.name,
    }, opts);
}

export interface GetVirtualMachineArgs {
    /**
     * Virtual Machine Name to get details of
     */
    name: string;
}

export interface GetVirtualMachineResult {
    /**
     * VM boot disk size. Will expand boot disk to this size.
     */
    readonly bootDiskSize?: enums.BootFirmwareType;
    /**
     * VM boot disk type. thin, zeroedthick, eagerzeroedthick
     */
    readonly bootDiskType?: enums.BootDiskType;
    /**
     * Boot type('efi' is boot uefi mode)
     */
    readonly bootFirmware?: string;
    /**
     * esxi diskstore for boot disk.
     */
    readonly diskStore?: string;
    /**
     * pass data to VM
     */
    readonly info?: outputs.ConfigItem[];
    /**
     * The IP address reported by VMWare tools.
     */
    readonly ipAddress?: string;
    /**
     * VM memory size.
     */
    readonly memSize?: string;
    /**
     * esxi vm name.
     */
    readonly name?: string;
    /**
     * VM network interfaces.
     */
    readonly networkInterfaces?: outputs.NetworkInterface[];
    /**
     * VM memory size.
     */
    readonly notes?: string;
    /**
     * VM number of virtual cpus.
     */
    readonly numVCpus?: string;
    /**
     * VM OS type.
     */
    readonly os?: string;
    /**
     * VM power state.
     */
    readonly power?: string;
    /**
     * Resource pool name to place vm.
     */
    readonly resourcePoolName?: string;
    /**
     * The amount of vm uptime, in seconds, to wait for an available IP address on this virtual machine.
     */
    readonly shutdownTimeout?: number;
    /**
     * The amount of vm uptime, in seconds, to wait for an available IP address on this virtual machine.
     */
    readonly startupTimeout?: number;
    /**
     * VM virtual disks.
     */
    readonly virtualDisks?: outputs.VirtualDisk[];
    /**
     * VM Virtual HW version.
     */
    readonly virtualHWVer?: string;
}

export function getVirtualMachineOutput(args: GetVirtualMachineOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetVirtualMachineResult> {
    return pulumi.output(args).apply(a => getVirtualMachine(a, opts))
}

export interface GetVirtualMachineOutputArgs {
    /**
     * Virtual Machine Name to get details of
     */
    name: pulumi.Input<string>;
}
