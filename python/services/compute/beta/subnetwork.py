# Copyright 2022 Google LLC. All Rights Reserved.
# 
# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
# 
#     http://www.apache.org/licenses/LICENSE-2.0
# 
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.
from connector import channel
from google3.cloud.graphite.mmv2.services.google.compute import subnetwork_pb2
from google3.cloud.graphite.mmv2.services.google.compute import subnetwork_pb2_grpc

from typing import List


class Subnetwork(object):
    def __init__(
        self,
        creation_timestamp: str = None,
        description: str = None,
        gateway_address: str = None,
        ip_cidr_range: str = None,
        name: str = None,
        network: str = None,
        fingerprint: str = None,
        purpose: str = None,
        role: str = None,
        secondary_ip_ranges: list = None,
        private_ip_google_access: bool = None,
        region: str = None,
        log_config: dict = None,
        project: str = None,
        self_link: str = None,
        enable_flow_logs: bool = None,
        service_account_file: str = "",
    ):

        channel.initialize()
        self.description = description
        self.ip_cidr_range = ip_cidr_range
        self.name = name
        self.network = network
        self.purpose = purpose
        self.role = role
        self.secondary_ip_ranges = secondary_ip_ranges
        self.private_ip_google_access = private_ip_google_access
        self.region = region
        self.log_config = log_config
        self.project = project
        self.enable_flow_logs = enable_flow_logs
        self.service_account_file = service_account_file

    def apply(self):
        stub = subnetwork_pb2_grpc.ComputeBetaSubnetworkServiceStub(channel.Channel())
        request = subnetwork_pb2.ApplyComputeBetaSubnetworkRequest()
        if Primitive.to_proto(self.description):
            request.resource.description = Primitive.to_proto(self.description)

        if Primitive.to_proto(self.ip_cidr_range):
            request.resource.ip_cidr_range = Primitive.to_proto(self.ip_cidr_range)

        if Primitive.to_proto(self.name):
            request.resource.name = Primitive.to_proto(self.name)

        if Primitive.to_proto(self.network):
            request.resource.network = Primitive.to_proto(self.network)

        if SubnetworkPurposeEnum.to_proto(self.purpose):
            request.resource.purpose = SubnetworkPurposeEnum.to_proto(self.purpose)

        if SubnetworkRoleEnum.to_proto(self.role):
            request.resource.role = SubnetworkRoleEnum.to_proto(self.role)

        if SubnetworkSecondaryIPRangesArray.to_proto(self.secondary_ip_ranges):
            request.resource.secondary_ip_ranges.extend(
                SubnetworkSecondaryIPRangesArray.to_proto(self.secondary_ip_ranges)
            )
        if Primitive.to_proto(self.private_ip_google_access):
            request.resource.private_ip_google_access = Primitive.to_proto(
                self.private_ip_google_access
            )

        if Primitive.to_proto(self.region):
            request.resource.region = Primitive.to_proto(self.region)

        if SubnetworkLogConfig.to_proto(self.log_config):
            request.resource.log_config.CopyFrom(
                SubnetworkLogConfig.to_proto(self.log_config)
            )
        else:
            request.resource.ClearField("log_config")
        if Primitive.to_proto(self.project):
            request.resource.project = Primitive.to_proto(self.project)

        if Primitive.to_proto(self.enable_flow_logs):
            request.resource.enable_flow_logs = Primitive.to_proto(
                self.enable_flow_logs
            )

        request.service_account_file = self.service_account_file

        response = stub.ApplyComputeBetaSubnetwork(request)
        self.creation_timestamp = Primitive.from_proto(response.creation_timestamp)
        self.description = Primitive.from_proto(response.description)
        self.gateway_address = Primitive.from_proto(response.gateway_address)
        self.ip_cidr_range = Primitive.from_proto(response.ip_cidr_range)
        self.name = Primitive.from_proto(response.name)
        self.network = Primitive.from_proto(response.network)
        self.fingerprint = Primitive.from_proto(response.fingerprint)
        self.purpose = SubnetworkPurposeEnum.from_proto(response.purpose)
        self.role = SubnetworkRoleEnum.from_proto(response.role)
        self.secondary_ip_ranges = SubnetworkSecondaryIPRangesArray.from_proto(
            response.secondary_ip_ranges
        )
        self.private_ip_google_access = Primitive.from_proto(
            response.private_ip_google_access
        )
        self.region = Primitive.from_proto(response.region)
        self.log_config = SubnetworkLogConfig.from_proto(response.log_config)
        self.project = Primitive.from_proto(response.project)
        self.self_link = Primitive.from_proto(response.self_link)
        self.enable_flow_logs = Primitive.from_proto(response.enable_flow_logs)

    def delete(self):
        stub = subnetwork_pb2_grpc.ComputeBetaSubnetworkServiceStub(channel.Channel())
        request = subnetwork_pb2.DeleteComputeBetaSubnetworkRequest()
        request.service_account_file = self.service_account_file
        if Primitive.to_proto(self.description):
            request.resource.description = Primitive.to_proto(self.description)

        if Primitive.to_proto(self.ip_cidr_range):
            request.resource.ip_cidr_range = Primitive.to_proto(self.ip_cidr_range)

        if Primitive.to_proto(self.name):
            request.resource.name = Primitive.to_proto(self.name)

        if Primitive.to_proto(self.network):
            request.resource.network = Primitive.to_proto(self.network)

        if SubnetworkPurposeEnum.to_proto(self.purpose):
            request.resource.purpose = SubnetworkPurposeEnum.to_proto(self.purpose)

        if SubnetworkRoleEnum.to_proto(self.role):
            request.resource.role = SubnetworkRoleEnum.to_proto(self.role)

        if SubnetworkSecondaryIPRangesArray.to_proto(self.secondary_ip_ranges):
            request.resource.secondary_ip_ranges.extend(
                SubnetworkSecondaryIPRangesArray.to_proto(self.secondary_ip_ranges)
            )
        if Primitive.to_proto(self.private_ip_google_access):
            request.resource.private_ip_google_access = Primitive.to_proto(
                self.private_ip_google_access
            )

        if Primitive.to_proto(self.region):
            request.resource.region = Primitive.to_proto(self.region)

        if SubnetworkLogConfig.to_proto(self.log_config):
            request.resource.log_config.CopyFrom(
                SubnetworkLogConfig.to_proto(self.log_config)
            )
        else:
            request.resource.ClearField("log_config")
        if Primitive.to_proto(self.project):
            request.resource.project = Primitive.to_proto(self.project)

        if Primitive.to_proto(self.enable_flow_logs):
            request.resource.enable_flow_logs = Primitive.to_proto(
                self.enable_flow_logs
            )

        response = stub.DeleteComputeBetaSubnetwork(request)

    @classmethod
    def list(self, project, region, service_account_file=""):
        stub = subnetwork_pb2_grpc.ComputeBetaSubnetworkServiceStub(channel.Channel())
        request = subnetwork_pb2.ListComputeBetaSubnetworkRequest()
        request.service_account_file = service_account_file
        request.Project = project

        request.Region = region

        return stub.ListComputeBetaSubnetwork(request).items

    def to_proto(self):
        resource = subnetwork_pb2.ComputeBetaSubnetwork()
        if Primitive.to_proto(self.description):
            resource.description = Primitive.to_proto(self.description)
        if Primitive.to_proto(self.ip_cidr_range):
            resource.ip_cidr_range = Primitive.to_proto(self.ip_cidr_range)
        if Primitive.to_proto(self.name):
            resource.name = Primitive.to_proto(self.name)
        if Primitive.to_proto(self.network):
            resource.network = Primitive.to_proto(self.network)
        if SubnetworkPurposeEnum.to_proto(self.purpose):
            resource.purpose = SubnetworkPurposeEnum.to_proto(self.purpose)
        if SubnetworkRoleEnum.to_proto(self.role):
            resource.role = SubnetworkRoleEnum.to_proto(self.role)
        if SubnetworkSecondaryIPRangesArray.to_proto(self.secondary_ip_ranges):
            resource.secondary_ip_ranges.extend(
                SubnetworkSecondaryIPRangesArray.to_proto(self.secondary_ip_ranges)
            )
        if Primitive.to_proto(self.private_ip_google_access):
            resource.private_ip_google_access = Primitive.to_proto(
                self.private_ip_google_access
            )
        if Primitive.to_proto(self.region):
            resource.region = Primitive.to_proto(self.region)
        if SubnetworkLogConfig.to_proto(self.log_config):
            resource.log_config.CopyFrom(SubnetworkLogConfig.to_proto(self.log_config))
        else:
            resource.ClearField("log_config")
        if Primitive.to_proto(self.project):
            resource.project = Primitive.to_proto(self.project)
        if Primitive.to_proto(self.enable_flow_logs):
            resource.enable_flow_logs = Primitive.to_proto(self.enable_flow_logs)
        return resource


class SubnetworkSecondaryIPRanges(object):
    def __init__(self, range_name: str = None, ip_cidr_range: str = None):
        self.range_name = range_name
        self.ip_cidr_range = ip_cidr_range

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = subnetwork_pb2.ComputeBetaSubnetworkSecondaryIPRanges()
        if Primitive.to_proto(resource.range_name):
            res.range_name = Primitive.to_proto(resource.range_name)
        if Primitive.to_proto(resource.ip_cidr_range):
            res.ip_cidr_range = Primitive.to_proto(resource.ip_cidr_range)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return SubnetworkSecondaryIPRanges(
            range_name=Primitive.from_proto(resource.range_name),
            ip_cidr_range=Primitive.from_proto(resource.ip_cidr_range),
        )


class SubnetworkSecondaryIPRangesArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [SubnetworkSecondaryIPRanges.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [SubnetworkSecondaryIPRanges.from_proto(i) for i in resources]


class SubnetworkLogConfig(object):
    def __init__(
        self,
        aggregation_interval: str = None,
        flow_sampling: float = None,
        metadata: str = None,
    ):
        self.aggregation_interval = aggregation_interval
        self.flow_sampling = flow_sampling
        self.metadata = metadata

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = subnetwork_pb2.ComputeBetaSubnetworkLogConfig()
        if SubnetworkLogConfigAggregationIntervalEnum.to_proto(
            resource.aggregation_interval
        ):
            res.aggregation_interval = SubnetworkLogConfigAggregationIntervalEnum.to_proto(
                resource.aggregation_interval
            )
        if Primitive.to_proto(resource.flow_sampling):
            res.flow_sampling = Primitive.to_proto(resource.flow_sampling)
        if SubnetworkLogConfigMetadataEnum.to_proto(resource.metadata):
            res.metadata = SubnetworkLogConfigMetadataEnum.to_proto(resource.metadata)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return SubnetworkLogConfig(
            aggregation_interval=SubnetworkLogConfigAggregationIntervalEnum.from_proto(
                resource.aggregation_interval
            ),
            flow_sampling=Primitive.from_proto(resource.flow_sampling),
            metadata=SubnetworkLogConfigMetadataEnum.from_proto(resource.metadata),
        )


class SubnetworkLogConfigArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [SubnetworkLogConfig.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [SubnetworkLogConfig.from_proto(i) for i in resources]


class SubnetworkPurposeEnum(object):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return subnetwork_pb2.ComputeBetaSubnetworkPurposeEnum.Value(
            "ComputeBetaSubnetworkPurposeEnum%s" % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return subnetwork_pb2.ComputeBetaSubnetworkPurposeEnum.Name(resource)[
            len("ComputeBetaSubnetworkPurposeEnum") :
        ]


class SubnetworkRoleEnum(object):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return subnetwork_pb2.ComputeBetaSubnetworkRoleEnum.Value(
            "ComputeBetaSubnetworkRoleEnum%s" % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return subnetwork_pb2.ComputeBetaSubnetworkRoleEnum.Name(resource)[
            len("ComputeBetaSubnetworkRoleEnum") :
        ]


class SubnetworkLogConfigAggregationIntervalEnum(object):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return subnetwork_pb2.ComputeBetaSubnetworkLogConfigAggregationIntervalEnum.Value(
            "ComputeBetaSubnetworkLogConfigAggregationIntervalEnum%s" % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return subnetwork_pb2.ComputeBetaSubnetworkLogConfigAggregationIntervalEnum.Name(
            resource
        )[
            len("ComputeBetaSubnetworkLogConfigAggregationIntervalEnum") :
        ]


class SubnetworkLogConfigMetadataEnum(object):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return subnetwork_pb2.ComputeBetaSubnetworkLogConfigMetadataEnum.Value(
            "ComputeBetaSubnetworkLogConfigMetadataEnum%s" % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return subnetwork_pb2.ComputeBetaSubnetworkLogConfigMetadataEnum.Name(resource)[
            len("ComputeBetaSubnetworkLogConfigMetadataEnum") :
        ]


class Primitive(object):
    @classmethod
    def to_proto(self, s):
        if not s:
            return ""
        return s

    @classmethod
    def from_proto(self, s):
        return s
