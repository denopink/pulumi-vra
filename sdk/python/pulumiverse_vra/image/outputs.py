# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities
from . import outputs

__all__ = [
    'ProfileImageMapping',
    'ProfileImageMappingConstraint',
    'GetProfileImageMappingResult',
    'GetProfileImageMappingConstraintResult',
]

@pulumi.output_type
class ProfileImageMapping(dict):
    @staticmethod
    def __key_warning(key: str):
        suggest = None
        if key == "cloudConfig":
            suggest = "cloud_config"
        elif key == "externalId":
            suggest = "external_id"
        elif key == "externalRegionId":
            suggest = "external_region_id"
        elif key == "imageId":
            suggest = "image_id"
        elif key == "imageName":
            suggest = "image_name"
        elif key == "osFamily":
            suggest = "os_family"

        if suggest:
            pulumi.log.warn(f"Key '{key}' not found in ProfileImageMapping. Access the value via the '{suggest}' property getter instead.")

    def __getitem__(self, key: str) -> Any:
        ProfileImageMapping.__key_warning(key)
        return super().__getitem__(key)

    def get(self, key: str, default = None) -> Any:
        ProfileImageMapping.__key_warning(key)
        return super().get(key, default)

    def __init__(__self__, *,
                 name: str,
                 cloud_config: Optional[str] = None,
                 constraints: Optional[Sequence['outputs.ProfileImageMappingConstraint']] = None,
                 description: Optional[str] = None,
                 external_id: Optional[str] = None,
                 external_region_id: Optional[str] = None,
                 image_id: Optional[str] = None,
                 image_name: Optional[str] = None,
                 organization: Optional[str] = None,
                 os_family: Optional[str] = None,
                 owner: Optional[str] = None,
                 private: Optional[bool] = None):
        """
        :param str name: A human-friendly name used as an identifier in APIs that support this option.
        :param str description: A human-friendly description.
        :param str external_region_id: The external regionId of the resource.
        :param str owner: Email of the user that owns the entity.
        """
        pulumi.set(__self__, "name", name)
        if cloud_config is not None:
            pulumi.set(__self__, "cloud_config", cloud_config)
        if constraints is not None:
            pulumi.set(__self__, "constraints", constraints)
        if description is not None:
            pulumi.set(__self__, "description", description)
        if external_id is not None:
            pulumi.set(__self__, "external_id", external_id)
        if external_region_id is not None:
            pulumi.set(__self__, "external_region_id", external_region_id)
        if image_id is not None:
            pulumi.set(__self__, "image_id", image_id)
        if image_name is not None:
            pulumi.set(__self__, "image_name", image_name)
        if organization is not None:
            pulumi.set(__self__, "organization", organization)
        if os_family is not None:
            pulumi.set(__self__, "os_family", os_family)
        if owner is not None:
            pulumi.set(__self__, "owner", owner)
        if private is not None:
            pulumi.set(__self__, "private", private)

    @property
    @pulumi.getter
    def name(self) -> str:
        """
        A human-friendly name used as an identifier in APIs that support this option.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="cloudConfig")
    def cloud_config(self) -> Optional[str]:
        return pulumi.get(self, "cloud_config")

    @property
    @pulumi.getter
    def constraints(self) -> Optional[Sequence['outputs.ProfileImageMappingConstraint']]:
        return pulumi.get(self, "constraints")

    @property
    @pulumi.getter
    def description(self) -> Optional[str]:
        """
        A human-friendly description.
        """
        return pulumi.get(self, "description")

    @property
    @pulumi.getter(name="externalId")
    def external_id(self) -> Optional[str]:
        return pulumi.get(self, "external_id")

    @property
    @pulumi.getter(name="externalRegionId")
    def external_region_id(self) -> Optional[str]:
        """
        The external regionId of the resource.
        """
        return pulumi.get(self, "external_region_id")

    @property
    @pulumi.getter(name="imageId")
    def image_id(self) -> Optional[str]:
        return pulumi.get(self, "image_id")

    @property
    @pulumi.getter(name="imageName")
    def image_name(self) -> Optional[str]:
        return pulumi.get(self, "image_name")

    @property
    @pulumi.getter
    def organization(self) -> Optional[str]:
        return pulumi.get(self, "organization")

    @property
    @pulumi.getter(name="osFamily")
    def os_family(self) -> Optional[str]:
        return pulumi.get(self, "os_family")

    @property
    @pulumi.getter
    def owner(self) -> Optional[str]:
        """
        Email of the user that owns the entity.
        """
        return pulumi.get(self, "owner")

    @property
    @pulumi.getter
    def private(self) -> Optional[bool]:
        return pulumi.get(self, "private")


@pulumi.output_type
class ProfileImageMappingConstraint(dict):
    def __init__(__self__, *,
                 expression: str,
                 mandatory: bool):
        pulumi.set(__self__, "expression", expression)
        pulumi.set(__self__, "mandatory", mandatory)

    @property
    @pulumi.getter
    def expression(self) -> str:
        return pulumi.get(self, "expression")

    @property
    @pulumi.getter
    def mandatory(self) -> bool:
        return pulumi.get(self, "mandatory")


@pulumi.output_type
class GetProfileImageMappingResult(dict):
    def __init__(__self__, *,
                 description: str,
                 external_id: str,
                 external_region_id: str,
                 name: str,
                 organization: str,
                 os_family: str,
                 owner: str,
                 private: bool,
                 cloud_config: Optional[str] = None,
                 constraints: Optional[Sequence['outputs.GetProfileImageMappingConstraintResult']] = None,
                 image_id: Optional[str] = None,
                 image_name: Optional[str] = None):
        """
        :param str description: A human-friendly description.
        :param str external_region_id: The external regionId of the resource.
        :param str name: A human-friendly name used as an identifier in APIs that support this option.
        :param str owner: Email of the user that owns the entity.
        """
        pulumi.set(__self__, "description", description)
        pulumi.set(__self__, "external_id", external_id)
        pulumi.set(__self__, "external_region_id", external_region_id)
        pulumi.set(__self__, "name", name)
        pulumi.set(__self__, "organization", organization)
        pulumi.set(__self__, "os_family", os_family)
        pulumi.set(__self__, "owner", owner)
        pulumi.set(__self__, "private", private)
        if cloud_config is not None:
            pulumi.set(__self__, "cloud_config", cloud_config)
        if constraints is not None:
            pulumi.set(__self__, "constraints", constraints)
        if image_id is not None:
            pulumi.set(__self__, "image_id", image_id)
        if image_name is not None:
            pulumi.set(__self__, "image_name", image_name)

    @property
    @pulumi.getter
    def description(self) -> str:
        """
        A human-friendly description.
        """
        return pulumi.get(self, "description")

    @property
    @pulumi.getter(name="externalId")
    def external_id(self) -> str:
        return pulumi.get(self, "external_id")

    @property
    @pulumi.getter(name="externalRegionId")
    def external_region_id(self) -> str:
        """
        The external regionId of the resource.
        """
        return pulumi.get(self, "external_region_id")

    @property
    @pulumi.getter
    def name(self) -> str:
        """
        A human-friendly name used as an identifier in APIs that support this option.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def organization(self) -> str:
        return pulumi.get(self, "organization")

    @property
    @pulumi.getter(name="osFamily")
    def os_family(self) -> str:
        return pulumi.get(self, "os_family")

    @property
    @pulumi.getter
    def owner(self) -> str:
        """
        Email of the user that owns the entity.
        """
        return pulumi.get(self, "owner")

    @property
    @pulumi.getter
    def private(self) -> bool:
        return pulumi.get(self, "private")

    @property
    @pulumi.getter(name="cloudConfig")
    def cloud_config(self) -> Optional[str]:
        return pulumi.get(self, "cloud_config")

    @property
    @pulumi.getter
    def constraints(self) -> Optional[Sequence['outputs.GetProfileImageMappingConstraintResult']]:
        return pulumi.get(self, "constraints")

    @property
    @pulumi.getter(name="imageId")
    def image_id(self) -> Optional[str]:
        return pulumi.get(self, "image_id")

    @property
    @pulumi.getter(name="imageName")
    def image_name(self) -> Optional[str]:
        return pulumi.get(self, "image_name")


@pulumi.output_type
class GetProfileImageMappingConstraintResult(dict):
    def __init__(__self__, *,
                 expression: str,
                 mandatory: bool):
        pulumi.set(__self__, "expression", expression)
        pulumi.set(__self__, "mandatory", mandatory)

    @property
    @pulumi.getter
    def expression(self) -> str:
        return pulumi.get(self, "expression")

    @property
    @pulumi.getter
    def mandatory(self) -> bool:
        return pulumi.get(self, "mandatory")


