# coding=utf-8
# *** WARNING: this file was generated by pulumigen. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

import types

__config__ = pulumi.Config('esxi-native')


class _ExportableConfig(types.ModuleType):
    @property
    def host(self) -> Optional[str]:
        """
        ESXi Host Name config
        """
        return __config__.get('host')

    @property
    def ovf_tool_datastore_name(self) -> Optional[str]:
        """
        ESXi Datastore Name config were ovftool will be configured
        """
        return __config__.get('ovfToolDatastoreName')

    @property
    def password(self) -> Optional[str]:
        """
        ESXi Password config
        """
        return __config__.get('password')

    @property
    def ssh_port(self) -> str:
        """
        ESXi Host SSH Port config
        """
        return __config__.get('sshPort') or '22'

    @property
    def ssl_port(self) -> str:
        """
        ESXi Host SSL Port config
        """
        return __config__.get('sslPort') or '443'

    @property
    def username(self) -> str:
        """
        ESXi Username config
        """
        return __config__.get('username') or 'root'

