// Copyright 2022 Google LLC. All Rights Reserved.
// 
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
// 
//     http://www.apache.org/licenses/LICENSE-2.0
// 
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
package server

import (
	"context"
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/dcl"
	emptypb "github.com/GoogleCloudPlatform/declarative-resource-client-library/python/proto/empty_go_proto"
	betapb "github.com/GoogleCloudPlatform/declarative-resource-client-library/python/proto/vertexai/beta/vertexai_beta_go_proto"
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/vertexai/beta"
)

// ModelDeploymentServer implements the gRPC interface for ModelDeployment.
type ModelDeploymentServer struct{}

// ProtoToModelDeploymentDedicatedResources converts a ModelDeploymentDedicatedResources object from its proto representation.
func ProtoToVertexaiBetaModelDeploymentDedicatedResources(p *betapb.VertexaiBetaModelDeploymentDedicatedResources) *beta.ModelDeploymentDedicatedResources {
	if p == nil {
		return nil
	}
	obj := &beta.ModelDeploymentDedicatedResources{
		MachineSpec:     ProtoToVertexaiBetaModelDeploymentDedicatedResourcesMachineSpec(p.GetMachineSpec()),
		MinReplicaCount: dcl.Int64OrNil(p.GetMinReplicaCount()),
		MaxReplicaCount: dcl.Int64OrNil(p.GetMaxReplicaCount()),
	}
	return obj
}

// ProtoToModelDeploymentDedicatedResourcesMachineSpec converts a ModelDeploymentDedicatedResourcesMachineSpec object from its proto representation.
func ProtoToVertexaiBetaModelDeploymentDedicatedResourcesMachineSpec(p *betapb.VertexaiBetaModelDeploymentDedicatedResourcesMachineSpec) *beta.ModelDeploymentDedicatedResourcesMachineSpec {
	if p == nil {
		return nil
	}
	obj := &beta.ModelDeploymentDedicatedResourcesMachineSpec{
		MachineType: dcl.StringOrNil(p.GetMachineType()),
	}
	return obj
}

// ProtoToModelDeployment converts a ModelDeployment resource from its proto representation.
func ProtoToModelDeployment(p *betapb.VertexaiBetaModelDeployment) *beta.ModelDeployment {
	obj := &beta.ModelDeployment{
		Model:              dcl.StringOrNil(p.GetModel()),
		Id:                 dcl.StringOrNil(p.GetId()),
		DedicatedResources: ProtoToVertexaiBetaModelDeploymentDedicatedResources(p.GetDedicatedResources()),
		Endpoint:           dcl.StringOrNil(p.GetEndpoint()),
		Location:           dcl.StringOrNil(p.GetLocation()),
		Project:            dcl.StringOrNil(p.GetProject()),
	}
	return obj
}

// ModelDeploymentDedicatedResourcesToProto converts a ModelDeploymentDedicatedResources object to its proto representation.
func VertexaiBetaModelDeploymentDedicatedResourcesToProto(o *beta.ModelDeploymentDedicatedResources) *betapb.VertexaiBetaModelDeploymentDedicatedResources {
	if o == nil {
		return nil
	}
	p := &betapb.VertexaiBetaModelDeploymentDedicatedResources{}
	p.SetMachineSpec(VertexaiBetaModelDeploymentDedicatedResourcesMachineSpecToProto(o.MachineSpec))
	p.SetMinReplicaCount(dcl.ValueOrEmptyInt64(o.MinReplicaCount))
	p.SetMaxReplicaCount(dcl.ValueOrEmptyInt64(o.MaxReplicaCount))
	return p
}

// ModelDeploymentDedicatedResourcesMachineSpecToProto converts a ModelDeploymentDedicatedResourcesMachineSpec object to its proto representation.
func VertexaiBetaModelDeploymentDedicatedResourcesMachineSpecToProto(o *beta.ModelDeploymentDedicatedResourcesMachineSpec) *betapb.VertexaiBetaModelDeploymentDedicatedResourcesMachineSpec {
	if o == nil {
		return nil
	}
	p := &betapb.VertexaiBetaModelDeploymentDedicatedResourcesMachineSpec{}
	p.SetMachineType(dcl.ValueOrEmptyString(o.MachineType))
	return p
}

// ModelDeploymentToProto converts a ModelDeployment resource to its proto representation.
func ModelDeploymentToProto(resource *beta.ModelDeployment) *betapb.VertexaiBetaModelDeployment {
	p := &betapb.VertexaiBetaModelDeployment{}
	p.SetModel(dcl.ValueOrEmptyString(resource.Model))
	p.SetId(dcl.ValueOrEmptyString(resource.Id))
	p.SetDedicatedResources(VertexaiBetaModelDeploymentDedicatedResourcesToProto(resource.DedicatedResources))
	p.SetEndpoint(dcl.ValueOrEmptyString(resource.Endpoint))
	p.SetLocation(dcl.ValueOrEmptyString(resource.Location))
	p.SetProject(dcl.ValueOrEmptyString(resource.Project))

	return p
}

// applyModelDeployment handles the gRPC request by passing it to the underlying ModelDeployment Apply() method.
func (s *ModelDeploymentServer) applyModelDeployment(ctx context.Context, c *beta.Client, request *betapb.ApplyVertexaiBetaModelDeploymentRequest) (*betapb.VertexaiBetaModelDeployment, error) {
	p := ProtoToModelDeployment(request.GetResource())
	res, err := c.ApplyModelDeployment(ctx, p)
	if err != nil {
		return nil, err
	}
	r := ModelDeploymentToProto(res)
	return r, nil
}

// applyVertexaiBetaModelDeployment handles the gRPC request by passing it to the underlying ModelDeployment Apply() method.
func (s *ModelDeploymentServer) ApplyVertexaiBetaModelDeployment(ctx context.Context, request *betapb.ApplyVertexaiBetaModelDeploymentRequest) (*betapb.VertexaiBetaModelDeployment, error) {
	cl, err := createConfigModelDeployment(ctx, request.GetServiceAccountFile())
	if err != nil {
		return nil, err
	}
	return s.applyModelDeployment(ctx, cl, request)
}

// DeleteModelDeployment handles the gRPC request by passing it to the underlying ModelDeployment Delete() method.
func (s *ModelDeploymentServer) DeleteVertexaiBetaModelDeployment(ctx context.Context, request *betapb.DeleteVertexaiBetaModelDeploymentRequest) (*emptypb.Empty, error) {

	cl, err := createConfigModelDeployment(ctx, request.GetServiceAccountFile())
	if err != nil {
		return nil, err
	}
	return &emptypb.Empty{}, cl.DeleteModelDeployment(ctx, ProtoToModelDeployment(request.GetResource()))

}

// ListVertexaiBetaModelDeployment handles the gRPC request by passing it to the underlying ModelDeploymentList() method.
func (s *ModelDeploymentServer) ListVertexaiBetaModelDeployment(ctx context.Context, request *betapb.ListVertexaiBetaModelDeploymentRequest) (*betapb.ListVertexaiBetaModelDeploymentResponse, error) {
	cl, err := createConfigModelDeployment(ctx, request.GetServiceAccountFile())
	if err != nil {
		return nil, err
	}

	resources, err := cl.ListModelDeployment(ctx, request.GetProject(), request.GetLocation(), request.GetEndpoint())
	if err != nil {
		return nil, err
	}
	var protos []*betapb.VertexaiBetaModelDeployment
	for _, r := range resources.Items {
		rp := ModelDeploymentToProto(r)
		protos = append(protos, rp)
	}
	p := &betapb.ListVertexaiBetaModelDeploymentResponse{}
	p.SetItems(protos)
	return p, nil
}

func createConfigModelDeployment(ctx context.Context, service_account_file string) (*beta.Client, error) {

	conf := dcl.NewConfig(dcl.WithUserAgent("dcl-test"), dcl.WithCredentialsFile(service_account_file))
	return beta.NewClient(conf), nil
}
