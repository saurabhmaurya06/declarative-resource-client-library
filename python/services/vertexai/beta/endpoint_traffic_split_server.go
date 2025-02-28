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
	"errors"
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/dcl"
	emptypb "github.com/GoogleCloudPlatform/declarative-resource-client-library/python/proto/empty_go_proto"
	betapb "github.com/GoogleCloudPlatform/declarative-resource-client-library/python/proto/vertexai/beta/vertexai_beta_go_proto"
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/vertexai/beta"
)

// EndpointTrafficSplitServer implements the gRPC interface for EndpointTrafficSplit.
type EndpointTrafficSplitServer struct{}

// ProtoToEndpointTrafficSplitTrafficSplit converts a EndpointTrafficSplitTrafficSplit object from its proto representation.
func ProtoToVertexaiBetaEndpointTrafficSplitTrafficSplit(p *betapb.VertexaiBetaEndpointTrafficSplitTrafficSplit) *beta.EndpointTrafficSplitTrafficSplit {
	if p == nil {
		return nil
	}
	obj := &beta.EndpointTrafficSplitTrafficSplit{
		DeployedModelId:   dcl.StringOrNil(p.GetDeployedModelId()),
		TrafficPercentage: dcl.Int64OrNil(p.GetTrafficPercentage()),
	}
	return obj
}

// ProtoToEndpointTrafficSplit converts a EndpointTrafficSplit resource from its proto representation.
func ProtoToEndpointTrafficSplit(p *betapb.VertexaiBetaEndpointTrafficSplit) *beta.EndpointTrafficSplit {
	obj := &beta.EndpointTrafficSplit{
		Endpoint: dcl.StringOrNil(p.GetEndpoint()),
		Project:  dcl.StringOrNil(p.GetProject()),
		Location: dcl.StringOrNil(p.GetLocation()),
		Etag:     dcl.StringOrNil(p.GetEtag()),
	}
	for _, r := range p.GetTrafficSplit() {
		obj.TrafficSplit = append(obj.TrafficSplit, *ProtoToVertexaiBetaEndpointTrafficSplitTrafficSplit(r))
	}
	return obj
}

// EndpointTrafficSplitTrafficSplitToProto converts a EndpointTrafficSplitTrafficSplit object to its proto representation.
func VertexaiBetaEndpointTrafficSplitTrafficSplitToProto(o *beta.EndpointTrafficSplitTrafficSplit) *betapb.VertexaiBetaEndpointTrafficSplitTrafficSplit {
	if o == nil {
		return nil
	}
	p := &betapb.VertexaiBetaEndpointTrafficSplitTrafficSplit{}
	p.SetDeployedModelId(dcl.ValueOrEmptyString(o.DeployedModelId))
	p.SetTrafficPercentage(dcl.ValueOrEmptyInt64(o.TrafficPercentage))
	return p
}

// EndpointTrafficSplitToProto converts a EndpointTrafficSplit resource to its proto representation.
func EndpointTrafficSplitToProto(resource *beta.EndpointTrafficSplit) *betapb.VertexaiBetaEndpointTrafficSplit {
	p := &betapb.VertexaiBetaEndpointTrafficSplit{}
	p.SetEndpoint(dcl.ValueOrEmptyString(resource.Endpoint))
	p.SetProject(dcl.ValueOrEmptyString(resource.Project))
	p.SetLocation(dcl.ValueOrEmptyString(resource.Location))
	p.SetEtag(dcl.ValueOrEmptyString(resource.Etag))
	sTrafficSplit := make([]*betapb.VertexaiBetaEndpointTrafficSplitTrafficSplit, len(resource.TrafficSplit))
	for i, r := range resource.TrafficSplit {
		sTrafficSplit[i] = VertexaiBetaEndpointTrafficSplitTrafficSplitToProto(&r)
	}
	p.SetTrafficSplit(sTrafficSplit)

	return p
}

// applyEndpointTrafficSplit handles the gRPC request by passing it to the underlying EndpointTrafficSplit Apply() method.
func (s *EndpointTrafficSplitServer) applyEndpointTrafficSplit(ctx context.Context, c *beta.Client, request *betapb.ApplyVertexaiBetaEndpointTrafficSplitRequest) (*betapb.VertexaiBetaEndpointTrafficSplit, error) {
	p := ProtoToEndpointTrafficSplit(request.GetResource())
	res, err := c.ApplyEndpointTrafficSplit(ctx, p)
	if err != nil {
		return nil, err
	}
	r := EndpointTrafficSplitToProto(res)
	return r, nil
}

// applyVertexaiBetaEndpointTrafficSplit handles the gRPC request by passing it to the underlying EndpointTrafficSplit Apply() method.
func (s *EndpointTrafficSplitServer) ApplyVertexaiBetaEndpointTrafficSplit(ctx context.Context, request *betapb.ApplyVertexaiBetaEndpointTrafficSplitRequest) (*betapb.VertexaiBetaEndpointTrafficSplit, error) {
	cl, err := createConfigEndpointTrafficSplit(ctx, request.GetServiceAccountFile())
	if err != nil {
		return nil, err
	}
	return s.applyEndpointTrafficSplit(ctx, cl, request)
}

// DeleteEndpointTrafficSplit handles the gRPC request by passing it to the underlying EndpointTrafficSplit Delete() method.
func (s *EndpointTrafficSplitServer) DeleteVertexaiBetaEndpointTrafficSplit(ctx context.Context, request *betapb.DeleteVertexaiBetaEndpointTrafficSplitRequest) (*emptypb.Empty, error) {

	return nil, errors.New("no delete endpoint for EndpointTrafficSplit")

}

// ListVertexaiBetaEndpointTrafficSplit is a no-op method because EndpointTrafficSplit has no list method.
func (s *EndpointTrafficSplitServer) ListVertexaiBetaEndpointTrafficSplit(_ context.Context, _ *betapb.ListVertexaiBetaEndpointTrafficSplitRequest) (*betapb.ListVertexaiBetaEndpointTrafficSplitResponse, error) {
	return nil, nil
}

func createConfigEndpointTrafficSplit(ctx context.Context, service_account_file string) (*beta.Client, error) {

	conf := dcl.NewConfig(dcl.WithUserAgent("dcl-test"), dcl.WithCredentialsFile(service_account_file))
	return beta.NewClient(conf), nil
}
