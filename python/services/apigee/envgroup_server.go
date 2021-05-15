// Copyright 2021 Google LLC. All Rights Reserved.
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
	apigeepb "github.com/GoogleCloudPlatform/declarative-resource-client-library/python/proto/apigee/apigee_go_proto"
	emptypb "github.com/GoogleCloudPlatform/declarative-resource-client-library/python/proto/empty_go_proto"
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/apigee"
)

// Server implements the gRPC interface for Envgroup.
type EnvgroupServer struct{}

// ProtoToEnvgroupStateEnum converts a EnvgroupStateEnum enum from its proto representation.
func ProtoToApigeeEnvgroupStateEnum(e apigeepb.ApigeeEnvgroupStateEnum) *apigee.EnvgroupStateEnum {
	if e == 0 {
		return nil
	}
	if n, ok := apigeepb.ApigeeEnvgroupStateEnum_name[int32(e)]; ok {
		e := apigee.EnvgroupStateEnum(n[len("EnvgroupStateEnum"):])
		return &e
	}
	return nil
}

// ProtoToEnvgroup converts a Envgroup resource from its proto representation.
func ProtoToEnvgroup(p *apigeepb.ApigeeEnvgroup) *apigee.Envgroup {
	obj := &apigee.Envgroup{
		Name:           dcl.StringOrNil(p.Name),
		CreatedAt:      dcl.Int64OrNil(p.CreatedAt),
		LastModifiedAt: dcl.Int64OrNil(p.LastModifiedAt),
		State:          ProtoToApigeeEnvgroupStateEnum(p.GetState()),
		Organization:   dcl.StringOrNil(p.Organization),
	}
	for _, r := range p.GetHostnames() {
		obj.Hostnames = append(obj.Hostnames, r)
	}
	return obj
}

// EnvgroupStateEnumToProto converts a EnvgroupStateEnum enum to its proto representation.
func ApigeeEnvgroupStateEnumToProto(e *apigee.EnvgroupStateEnum) apigeepb.ApigeeEnvgroupStateEnum {
	if e == nil {
		return apigeepb.ApigeeEnvgroupStateEnum(0)
	}
	if v, ok := apigeepb.ApigeeEnvgroupStateEnum_value["EnvgroupStateEnum"+string(*e)]; ok {
		return apigeepb.ApigeeEnvgroupStateEnum(v)
	}
	return apigeepb.ApigeeEnvgroupStateEnum(0)
}

// EnvgroupToProto converts a Envgroup resource to its proto representation.
func EnvgroupToProto(resource *apigee.Envgroup) *apigeepb.ApigeeEnvgroup {
	p := &apigeepb.ApigeeEnvgroup{
		Name:           dcl.ValueOrEmptyString(resource.Name),
		CreatedAt:      dcl.ValueOrEmptyInt64(resource.CreatedAt),
		LastModifiedAt: dcl.ValueOrEmptyInt64(resource.LastModifiedAt),
		State:          ApigeeEnvgroupStateEnumToProto(resource.State),
		Organization:   dcl.ValueOrEmptyString(resource.Organization),
	}
	for _, r := range resource.Hostnames {
		p.Hostnames = append(p.Hostnames, r)
	}

	return p
}

// ApplyEnvgroup handles the gRPC request by passing it to the underlying Envgroup Apply() method.
func (s *EnvgroupServer) applyEnvgroup(ctx context.Context, c *apigee.Client, request *apigeepb.ApplyApigeeEnvgroupRequest) (*apigeepb.ApigeeEnvgroup, error) {
	p := ProtoToEnvgroup(request.GetResource())
	res, err := c.ApplyEnvgroup(ctx, p)
	if err != nil {
		return nil, err
	}
	r := EnvgroupToProto(res)
	return r, nil
}

// ApplyEnvgroup handles the gRPC request by passing it to the underlying Envgroup Apply() method.
func (s *EnvgroupServer) ApplyApigeeEnvgroup(ctx context.Context, request *apigeepb.ApplyApigeeEnvgroupRequest) (*apigeepb.ApigeeEnvgroup, error) {
	cl, err := createConfigEnvgroup(ctx, request.ServiceAccountFile)
	if err != nil {
		return nil, err
	}
	return s.applyEnvgroup(ctx, cl, request)
}

// DeleteEnvgroup handles the gRPC request by passing it to the underlying Envgroup Delete() method.
func (s *EnvgroupServer) DeleteApigeeEnvgroup(ctx context.Context, request *apigeepb.DeleteApigeeEnvgroupRequest) (*emptypb.Empty, error) {

	cl, err := createConfigEnvgroup(ctx, request.ServiceAccountFile)
	if err != nil {
		return nil, err
	}
	return &emptypb.Empty{}, cl.DeleteEnvgroup(ctx, ProtoToEnvgroup(request.GetResource()))

}

// ListEnvgroup handles the gRPC request by passing it to the underlying EnvgroupList() method.
func (s *EnvgroupServer) ListApigeeEnvgroup(ctx context.Context, request *apigeepb.ListApigeeEnvgroupRequest) (*apigeepb.ListApigeeEnvgroupResponse, error) {
	cl, err := createConfigEnvgroup(ctx, request.ServiceAccountFile)
	if err != nil {
		return nil, err
	}

	resources, err := cl.ListEnvgroup(ctx, request.Organization)
	if err != nil {
		return nil, err
	}
	var protos []*apigeepb.ApigeeEnvgroup
	for _, r := range resources.Items {
		rp := EnvgroupToProto(r)
		protos = append(protos, rp)
	}
	return &apigeepb.ListApigeeEnvgroupResponse{Items: protos}, nil
}

func createConfigEnvgroup(ctx context.Context, service_account_file string) (*apigee.Client, error) {

	conf := dcl.NewConfig(dcl.WithUserAgent("dcl-test"), dcl.WithCredentialsFile(service_account_file))
	return apigee.NewClient(conf), nil
}
