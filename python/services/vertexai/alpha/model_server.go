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
	alphapb "github.com/GoogleCloudPlatform/declarative-resource-client-library/python/proto/vertexai/alpha/vertexai_alpha_go_proto"
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/vertexai/alpha"
)

// ModelServer implements the gRPC interface for Model.
type ModelServer struct{}

// ProtoToModelSupportedExportFormatsExportableContentsEnum converts a ModelSupportedExportFormatsExportableContentsEnum enum from its proto representation.
func ProtoToVertexaiAlphaModelSupportedExportFormatsExportableContentsEnum(e alphapb.VertexaiAlphaModelSupportedExportFormatsExportableContentsEnum) *alpha.ModelSupportedExportFormatsExportableContentsEnum {
	if e == 0 {
		return nil
	}
	if n, ok := alphapb.VertexaiAlphaModelSupportedExportFormatsExportableContentsEnum_name[int32(e)]; ok {
		e := alpha.ModelSupportedExportFormatsExportableContentsEnum(n[len("VertexaiAlphaModelSupportedExportFormatsExportableContentsEnum"):])
		return &e
	}
	return nil
}

// ProtoToModelContainerSpecAcceleratorRequirementsTypeEnum converts a ModelContainerSpecAcceleratorRequirementsTypeEnum enum from its proto representation.
func ProtoToVertexaiAlphaModelContainerSpecAcceleratorRequirementsTypeEnum(e alphapb.VertexaiAlphaModelContainerSpecAcceleratorRequirementsTypeEnum) *alpha.ModelContainerSpecAcceleratorRequirementsTypeEnum {
	if e == 0 {
		return nil
	}
	if n, ok := alphapb.VertexaiAlphaModelContainerSpecAcceleratorRequirementsTypeEnum_name[int32(e)]; ok {
		e := alpha.ModelContainerSpecAcceleratorRequirementsTypeEnum(n[len("VertexaiAlphaModelContainerSpecAcceleratorRequirementsTypeEnum"):])
		return &e
	}
	return nil
}

// ProtoToModelSupportedDeploymentResourcesTypesEnum converts a ModelSupportedDeploymentResourcesTypesEnum enum from its proto representation.
func ProtoToVertexaiAlphaModelSupportedDeploymentResourcesTypesEnum(e alphapb.VertexaiAlphaModelSupportedDeploymentResourcesTypesEnum) *alpha.ModelSupportedDeploymentResourcesTypesEnum {
	if e == 0 {
		return nil
	}
	if n, ok := alphapb.VertexaiAlphaModelSupportedDeploymentResourcesTypesEnum_name[int32(e)]; ok {
		e := alpha.ModelSupportedDeploymentResourcesTypesEnum(n[len("VertexaiAlphaModelSupportedDeploymentResourcesTypesEnum"):])
		return &e
	}
	return nil
}

// ProtoToModelSupportedExportFormats converts a ModelSupportedExportFormats object from its proto representation.
func ProtoToVertexaiAlphaModelSupportedExportFormats(p *alphapb.VertexaiAlphaModelSupportedExportFormats) *alpha.ModelSupportedExportFormats {
	if p == nil {
		return nil
	}
	obj := &alpha.ModelSupportedExportFormats{
		Id: dcl.StringOrNil(p.GetId()),
	}
	for _, r := range p.GetExportableContents() {
		obj.ExportableContents = append(obj.ExportableContents, *ProtoToVertexaiAlphaModelSupportedExportFormatsExportableContentsEnum(r))
	}
	return obj
}

// ProtoToModelOriginalModelInfo converts a ModelOriginalModelInfo object from its proto representation.
func ProtoToVertexaiAlphaModelOriginalModelInfo(p *alphapb.VertexaiAlphaModelOriginalModelInfo) *alpha.ModelOriginalModelInfo {
	if p == nil {
		return nil
	}
	obj := &alpha.ModelOriginalModelInfo{
		Model: dcl.StringOrNil(p.GetModel()),
	}
	return obj
}

// ProtoToModelContainerSpec converts a ModelContainerSpec object from its proto representation.
func ProtoToVertexaiAlphaModelContainerSpec(p *alphapb.VertexaiAlphaModelContainerSpec) *alpha.ModelContainerSpec {
	if p == nil {
		return nil
	}
	obj := &alpha.ModelContainerSpec{
		ImageUri:     dcl.StringOrNil(p.GetImageUri()),
		PredictRoute: dcl.StringOrNil(p.GetPredictRoute()),
		HealthRoute:  dcl.StringOrNil(p.GetHealthRoute()),
	}
	for _, r := range p.GetCommand() {
		obj.Command = append(obj.Command, r)
	}
	for _, r := range p.GetArgs() {
		obj.Args = append(obj.Args, r)
	}
	for _, r := range p.GetEnv() {
		obj.Env = append(obj.Env, *ProtoToVertexaiAlphaModelContainerSpecEnv(r))
	}
	for _, r := range p.GetPorts() {
		obj.Ports = append(obj.Ports, *ProtoToVertexaiAlphaModelContainerSpecPorts(r))
	}
	for _, r := range p.GetAcceleratorRequirements() {
		obj.AcceleratorRequirements = append(obj.AcceleratorRequirements, *ProtoToVertexaiAlphaModelContainerSpecAcceleratorRequirements(r))
	}
	return obj
}

// ProtoToModelContainerSpecEnv converts a ModelContainerSpecEnv object from its proto representation.
func ProtoToVertexaiAlphaModelContainerSpecEnv(p *alphapb.VertexaiAlphaModelContainerSpecEnv) *alpha.ModelContainerSpecEnv {
	if p == nil {
		return nil
	}
	obj := &alpha.ModelContainerSpecEnv{
		Name:  dcl.StringOrNil(p.GetName()),
		Value: dcl.StringOrNil(p.GetValue()),
	}
	return obj
}

// ProtoToModelContainerSpecPorts converts a ModelContainerSpecPorts object from its proto representation.
func ProtoToVertexaiAlphaModelContainerSpecPorts(p *alphapb.VertexaiAlphaModelContainerSpecPorts) *alpha.ModelContainerSpecPorts {
	if p == nil {
		return nil
	}
	obj := &alpha.ModelContainerSpecPorts{
		ContainerPort: dcl.Int64OrNil(p.GetContainerPort()),
	}
	return obj
}

// ProtoToModelContainerSpecAcceleratorRequirements converts a ModelContainerSpecAcceleratorRequirements object from its proto representation.
func ProtoToVertexaiAlphaModelContainerSpecAcceleratorRequirements(p *alphapb.VertexaiAlphaModelContainerSpecAcceleratorRequirements) *alpha.ModelContainerSpecAcceleratorRequirements {
	if p == nil {
		return nil
	}
	obj := &alpha.ModelContainerSpecAcceleratorRequirements{
		Type:  ProtoToVertexaiAlphaModelContainerSpecAcceleratorRequirementsTypeEnum(p.GetType()),
		Count: dcl.Int64OrNil(p.GetCount()),
	}
	return obj
}

// ProtoToModelDeployedModels converts a ModelDeployedModels object from its proto representation.
func ProtoToVertexaiAlphaModelDeployedModels(p *alphapb.VertexaiAlphaModelDeployedModels) *alpha.ModelDeployedModels {
	if p == nil {
		return nil
	}
	obj := &alpha.ModelDeployedModels{
		Endpoint:        dcl.StringOrNil(p.GetEndpoint()),
		DeployedModelId: dcl.StringOrNil(p.GetDeployedModelId()),
	}
	return obj
}

// ProtoToModelEncryptionSpec converts a ModelEncryptionSpec object from its proto representation.
func ProtoToVertexaiAlphaModelEncryptionSpec(p *alphapb.VertexaiAlphaModelEncryptionSpec) *alpha.ModelEncryptionSpec {
	if p == nil {
		return nil
	}
	obj := &alpha.ModelEncryptionSpec{
		KmsKeyName: dcl.StringOrNil(p.GetKmsKeyName()),
	}
	return obj
}

// ProtoToModel converts a Model resource from its proto representation.
func ProtoToModel(p *alphapb.VertexaiAlphaModel) *alpha.Model {
	obj := &alpha.Model{
		Name:               dcl.StringOrNil(p.GetName()),
		VersionId:          dcl.StringOrNil(p.GetVersionId()),
		VersionCreateTime:  dcl.StringOrNil(p.GetVersionCreateTime()),
		VersionUpdateTime:  dcl.StringOrNil(p.GetVersionUpdateTime()),
		DisplayName:        dcl.StringOrNil(p.GetDisplayName()),
		Description:        dcl.StringOrNil(p.GetDescription()),
		VersionDescription: dcl.StringOrNil(p.GetVersionDescription()),
		TrainingPipeline:   dcl.StringOrNil(p.GetTrainingPipeline()),
		OriginalModelInfo:  ProtoToVertexaiAlphaModelOriginalModelInfo(p.GetOriginalModelInfo()),
		ContainerSpec:      ProtoToVertexaiAlphaModelContainerSpec(p.GetContainerSpec()),
		ArtifactUri:        dcl.StringOrNil(p.GetArtifactUri()),
		CreateTime:         dcl.StringOrNil(p.GetCreateTime()),
		UpdateTime:         dcl.StringOrNil(p.GetUpdateTime()),
		Etag:               dcl.StringOrNil(p.GetEtag()),
		EncryptionSpec:     ProtoToVertexaiAlphaModelEncryptionSpec(p.GetEncryptionSpec()),
		Project:            dcl.StringOrNil(p.GetProject()),
		Location:           dcl.StringOrNil(p.GetLocation()),
	}
	for _, r := range p.GetVersionAliases() {
		obj.VersionAliases = append(obj.VersionAliases, r)
	}
	for _, r := range p.GetSupportedExportFormats() {
		obj.SupportedExportFormats = append(obj.SupportedExportFormats, *ProtoToVertexaiAlphaModelSupportedExportFormats(r))
	}
	for _, r := range p.GetSupportedDeploymentResourcesTypes() {
		obj.SupportedDeploymentResourcesTypes = append(obj.SupportedDeploymentResourcesTypes, *ProtoToVertexaiAlphaModelSupportedDeploymentResourcesTypesEnum(r))
	}
	for _, r := range p.GetSupportedInputStorageFormats() {
		obj.SupportedInputStorageFormats = append(obj.SupportedInputStorageFormats, r)
	}
	for _, r := range p.GetSupportedOutputStorageFormats() {
		obj.SupportedOutputStorageFormats = append(obj.SupportedOutputStorageFormats, r)
	}
	for _, r := range p.GetDeployedModels() {
		obj.DeployedModels = append(obj.DeployedModels, *ProtoToVertexaiAlphaModelDeployedModels(r))
	}
	return obj
}

// ModelSupportedExportFormatsExportableContentsEnumToProto converts a ModelSupportedExportFormatsExportableContentsEnum enum to its proto representation.
func VertexaiAlphaModelSupportedExportFormatsExportableContentsEnumToProto(e *alpha.ModelSupportedExportFormatsExportableContentsEnum) alphapb.VertexaiAlphaModelSupportedExportFormatsExportableContentsEnum {
	if e == nil {
		return alphapb.VertexaiAlphaModelSupportedExportFormatsExportableContentsEnum(0)
	}
	if v, ok := alphapb.VertexaiAlphaModelSupportedExportFormatsExportableContentsEnum_value["ModelSupportedExportFormatsExportableContentsEnum"+string(*e)]; ok {
		return alphapb.VertexaiAlphaModelSupportedExportFormatsExportableContentsEnum(v)
	}
	return alphapb.VertexaiAlphaModelSupportedExportFormatsExportableContentsEnum(0)
}

// ModelContainerSpecAcceleratorRequirementsTypeEnumToProto converts a ModelContainerSpecAcceleratorRequirementsTypeEnum enum to its proto representation.
func VertexaiAlphaModelContainerSpecAcceleratorRequirementsTypeEnumToProto(e *alpha.ModelContainerSpecAcceleratorRequirementsTypeEnum) alphapb.VertexaiAlphaModelContainerSpecAcceleratorRequirementsTypeEnum {
	if e == nil {
		return alphapb.VertexaiAlphaModelContainerSpecAcceleratorRequirementsTypeEnum(0)
	}
	if v, ok := alphapb.VertexaiAlphaModelContainerSpecAcceleratorRequirementsTypeEnum_value["ModelContainerSpecAcceleratorRequirementsTypeEnum"+string(*e)]; ok {
		return alphapb.VertexaiAlphaModelContainerSpecAcceleratorRequirementsTypeEnum(v)
	}
	return alphapb.VertexaiAlphaModelContainerSpecAcceleratorRequirementsTypeEnum(0)
}

// ModelSupportedDeploymentResourcesTypesEnumToProto converts a ModelSupportedDeploymentResourcesTypesEnum enum to its proto representation.
func VertexaiAlphaModelSupportedDeploymentResourcesTypesEnumToProto(e *alpha.ModelSupportedDeploymentResourcesTypesEnum) alphapb.VertexaiAlphaModelSupportedDeploymentResourcesTypesEnum {
	if e == nil {
		return alphapb.VertexaiAlphaModelSupportedDeploymentResourcesTypesEnum(0)
	}
	if v, ok := alphapb.VertexaiAlphaModelSupportedDeploymentResourcesTypesEnum_value["ModelSupportedDeploymentResourcesTypesEnum"+string(*e)]; ok {
		return alphapb.VertexaiAlphaModelSupportedDeploymentResourcesTypesEnum(v)
	}
	return alphapb.VertexaiAlphaModelSupportedDeploymentResourcesTypesEnum(0)
}

// ModelSupportedExportFormatsToProto converts a ModelSupportedExportFormats object to its proto representation.
func VertexaiAlphaModelSupportedExportFormatsToProto(o *alpha.ModelSupportedExportFormats) *alphapb.VertexaiAlphaModelSupportedExportFormats {
	if o == nil {
		return nil
	}
	p := &alphapb.VertexaiAlphaModelSupportedExportFormats{}
	p.SetId(dcl.ValueOrEmptyString(o.Id))
	sExportableContents := make([]alphapb.VertexaiAlphaModelSupportedExportFormatsExportableContentsEnum, len(o.ExportableContents))
	for i, r := range o.ExportableContents {
		sExportableContents[i] = alphapb.VertexaiAlphaModelSupportedExportFormatsExportableContentsEnum(alphapb.VertexaiAlphaModelSupportedExportFormatsExportableContentsEnum_value[string(r)])
	}
	p.SetExportableContents(sExportableContents)
	return p
}

// ModelOriginalModelInfoToProto converts a ModelOriginalModelInfo object to its proto representation.
func VertexaiAlphaModelOriginalModelInfoToProto(o *alpha.ModelOriginalModelInfo) *alphapb.VertexaiAlphaModelOriginalModelInfo {
	if o == nil {
		return nil
	}
	p := &alphapb.VertexaiAlphaModelOriginalModelInfo{}
	p.SetModel(dcl.ValueOrEmptyString(o.Model))
	return p
}

// ModelContainerSpecToProto converts a ModelContainerSpec object to its proto representation.
func VertexaiAlphaModelContainerSpecToProto(o *alpha.ModelContainerSpec) *alphapb.VertexaiAlphaModelContainerSpec {
	if o == nil {
		return nil
	}
	p := &alphapb.VertexaiAlphaModelContainerSpec{}
	p.SetImageUri(dcl.ValueOrEmptyString(o.ImageUri))
	p.SetPredictRoute(dcl.ValueOrEmptyString(o.PredictRoute))
	p.SetHealthRoute(dcl.ValueOrEmptyString(o.HealthRoute))
	sCommand := make([]string, len(o.Command))
	for i, r := range o.Command {
		sCommand[i] = r
	}
	p.SetCommand(sCommand)
	sArgs := make([]string, len(o.Args))
	for i, r := range o.Args {
		sArgs[i] = r
	}
	p.SetArgs(sArgs)
	sEnv := make([]*alphapb.VertexaiAlphaModelContainerSpecEnv, len(o.Env))
	for i, r := range o.Env {
		sEnv[i] = VertexaiAlphaModelContainerSpecEnvToProto(&r)
	}
	p.SetEnv(sEnv)
	sPorts := make([]*alphapb.VertexaiAlphaModelContainerSpecPorts, len(o.Ports))
	for i, r := range o.Ports {
		sPorts[i] = VertexaiAlphaModelContainerSpecPortsToProto(&r)
	}
	p.SetPorts(sPorts)
	sAcceleratorRequirements := make([]*alphapb.VertexaiAlphaModelContainerSpecAcceleratorRequirements, len(o.AcceleratorRequirements))
	for i, r := range o.AcceleratorRequirements {
		sAcceleratorRequirements[i] = VertexaiAlphaModelContainerSpecAcceleratorRequirementsToProto(&r)
	}
	p.SetAcceleratorRequirements(sAcceleratorRequirements)
	return p
}

// ModelContainerSpecEnvToProto converts a ModelContainerSpecEnv object to its proto representation.
func VertexaiAlphaModelContainerSpecEnvToProto(o *alpha.ModelContainerSpecEnv) *alphapb.VertexaiAlphaModelContainerSpecEnv {
	if o == nil {
		return nil
	}
	p := &alphapb.VertexaiAlphaModelContainerSpecEnv{}
	p.SetName(dcl.ValueOrEmptyString(o.Name))
	p.SetValue(dcl.ValueOrEmptyString(o.Value))
	return p
}

// ModelContainerSpecPortsToProto converts a ModelContainerSpecPorts object to its proto representation.
func VertexaiAlphaModelContainerSpecPortsToProto(o *alpha.ModelContainerSpecPorts) *alphapb.VertexaiAlphaModelContainerSpecPorts {
	if o == nil {
		return nil
	}
	p := &alphapb.VertexaiAlphaModelContainerSpecPorts{}
	p.SetContainerPort(dcl.ValueOrEmptyInt64(o.ContainerPort))
	return p
}

// ModelContainerSpecAcceleratorRequirementsToProto converts a ModelContainerSpecAcceleratorRequirements object to its proto representation.
func VertexaiAlphaModelContainerSpecAcceleratorRequirementsToProto(o *alpha.ModelContainerSpecAcceleratorRequirements) *alphapb.VertexaiAlphaModelContainerSpecAcceleratorRequirements {
	if o == nil {
		return nil
	}
	p := &alphapb.VertexaiAlphaModelContainerSpecAcceleratorRequirements{}
	p.SetType(VertexaiAlphaModelContainerSpecAcceleratorRequirementsTypeEnumToProto(o.Type))
	p.SetCount(dcl.ValueOrEmptyInt64(o.Count))
	return p
}

// ModelDeployedModelsToProto converts a ModelDeployedModels object to its proto representation.
func VertexaiAlphaModelDeployedModelsToProto(o *alpha.ModelDeployedModels) *alphapb.VertexaiAlphaModelDeployedModels {
	if o == nil {
		return nil
	}
	p := &alphapb.VertexaiAlphaModelDeployedModels{}
	p.SetEndpoint(dcl.ValueOrEmptyString(o.Endpoint))
	p.SetDeployedModelId(dcl.ValueOrEmptyString(o.DeployedModelId))
	return p
}

// ModelEncryptionSpecToProto converts a ModelEncryptionSpec object to its proto representation.
func VertexaiAlphaModelEncryptionSpecToProto(o *alpha.ModelEncryptionSpec) *alphapb.VertexaiAlphaModelEncryptionSpec {
	if o == nil {
		return nil
	}
	p := &alphapb.VertexaiAlphaModelEncryptionSpec{}
	p.SetKmsKeyName(dcl.ValueOrEmptyString(o.KmsKeyName))
	return p
}

// ModelToProto converts a Model resource to its proto representation.
func ModelToProto(resource *alpha.Model) *alphapb.VertexaiAlphaModel {
	p := &alphapb.VertexaiAlphaModel{}
	p.SetName(dcl.ValueOrEmptyString(resource.Name))
	p.SetVersionId(dcl.ValueOrEmptyString(resource.VersionId))
	p.SetVersionCreateTime(dcl.ValueOrEmptyString(resource.VersionCreateTime))
	p.SetVersionUpdateTime(dcl.ValueOrEmptyString(resource.VersionUpdateTime))
	p.SetDisplayName(dcl.ValueOrEmptyString(resource.DisplayName))
	p.SetDescription(dcl.ValueOrEmptyString(resource.Description))
	p.SetVersionDescription(dcl.ValueOrEmptyString(resource.VersionDescription))
	p.SetTrainingPipeline(dcl.ValueOrEmptyString(resource.TrainingPipeline))
	p.SetOriginalModelInfo(VertexaiAlphaModelOriginalModelInfoToProto(resource.OriginalModelInfo))
	p.SetContainerSpec(VertexaiAlphaModelContainerSpecToProto(resource.ContainerSpec))
	p.SetArtifactUri(dcl.ValueOrEmptyString(resource.ArtifactUri))
	p.SetCreateTime(dcl.ValueOrEmptyString(resource.CreateTime))
	p.SetUpdateTime(dcl.ValueOrEmptyString(resource.UpdateTime))
	p.SetEtag(dcl.ValueOrEmptyString(resource.Etag))
	p.SetEncryptionSpec(VertexaiAlphaModelEncryptionSpecToProto(resource.EncryptionSpec))
	p.SetProject(dcl.ValueOrEmptyString(resource.Project))
	p.SetLocation(dcl.ValueOrEmptyString(resource.Location))
	sVersionAliases := make([]string, len(resource.VersionAliases))
	for i, r := range resource.VersionAliases {
		sVersionAliases[i] = r
	}
	p.SetVersionAliases(sVersionAliases)
	sSupportedExportFormats := make([]*alphapb.VertexaiAlphaModelSupportedExportFormats, len(resource.SupportedExportFormats))
	for i, r := range resource.SupportedExportFormats {
		sSupportedExportFormats[i] = VertexaiAlphaModelSupportedExportFormatsToProto(&r)
	}
	p.SetSupportedExportFormats(sSupportedExportFormats)
	sSupportedDeploymentResourcesTypes := make([]alphapb.VertexaiAlphaModelSupportedDeploymentResourcesTypesEnum, len(resource.SupportedDeploymentResourcesTypes))
	for i, r := range resource.SupportedDeploymentResourcesTypes {
		sSupportedDeploymentResourcesTypes[i] = alphapb.VertexaiAlphaModelSupportedDeploymentResourcesTypesEnum(alphapb.VertexaiAlphaModelSupportedDeploymentResourcesTypesEnum_value[string(r)])
	}
	p.SetSupportedDeploymentResourcesTypes(sSupportedDeploymentResourcesTypes)
	sSupportedInputStorageFormats := make([]string, len(resource.SupportedInputStorageFormats))
	for i, r := range resource.SupportedInputStorageFormats {
		sSupportedInputStorageFormats[i] = r
	}
	p.SetSupportedInputStorageFormats(sSupportedInputStorageFormats)
	sSupportedOutputStorageFormats := make([]string, len(resource.SupportedOutputStorageFormats))
	for i, r := range resource.SupportedOutputStorageFormats {
		sSupportedOutputStorageFormats[i] = r
	}
	p.SetSupportedOutputStorageFormats(sSupportedOutputStorageFormats)
	sDeployedModels := make([]*alphapb.VertexaiAlphaModelDeployedModels, len(resource.DeployedModels))
	for i, r := range resource.DeployedModels {
		sDeployedModels[i] = VertexaiAlphaModelDeployedModelsToProto(&r)
	}
	p.SetDeployedModels(sDeployedModels)
	mLabels := make(map[string]string, len(resource.Labels))
	for k, r := range resource.Labels {
		mLabels[k] = r
	}
	p.SetLabels(mLabels)

	return p
}

// applyModel handles the gRPC request by passing it to the underlying Model Apply() method.
func (s *ModelServer) applyModel(ctx context.Context, c *alpha.Client, request *alphapb.ApplyVertexaiAlphaModelRequest) (*alphapb.VertexaiAlphaModel, error) {
	p := ProtoToModel(request.GetResource())
	res, err := c.ApplyModel(ctx, p)
	if err != nil {
		return nil, err
	}
	r := ModelToProto(res)
	return r, nil
}

// applyVertexaiAlphaModel handles the gRPC request by passing it to the underlying Model Apply() method.
func (s *ModelServer) ApplyVertexaiAlphaModel(ctx context.Context, request *alphapb.ApplyVertexaiAlphaModelRequest) (*alphapb.VertexaiAlphaModel, error) {
	cl, err := createConfigModel(ctx, request.GetServiceAccountFile())
	if err != nil {
		return nil, err
	}
	return s.applyModel(ctx, cl, request)
}

// DeleteModel handles the gRPC request by passing it to the underlying Model Delete() method.
func (s *ModelServer) DeleteVertexaiAlphaModel(ctx context.Context, request *alphapb.DeleteVertexaiAlphaModelRequest) (*emptypb.Empty, error) {

	cl, err := createConfigModel(ctx, request.GetServiceAccountFile())
	if err != nil {
		return nil, err
	}
	return &emptypb.Empty{}, cl.DeleteModel(ctx, ProtoToModel(request.GetResource()))

}

// ListVertexaiAlphaModel handles the gRPC request by passing it to the underlying ModelList() method.
func (s *ModelServer) ListVertexaiAlphaModel(ctx context.Context, request *alphapb.ListVertexaiAlphaModelRequest) (*alphapb.ListVertexaiAlphaModelResponse, error) {
	cl, err := createConfigModel(ctx, request.GetServiceAccountFile())
	if err != nil {
		return nil, err
	}

	resources, err := cl.ListModel(ctx, request.GetProject(), request.GetLocation())
	if err != nil {
		return nil, err
	}
	var protos []*alphapb.VertexaiAlphaModel
	for _, r := range resources.Items {
		rp := ModelToProto(r)
		protos = append(protos, rp)
	}
	p := &alphapb.ListVertexaiAlphaModelResponse{}
	p.SetItems(protos)
	return p, nil
}

func createConfigModel(ctx context.Context, service_account_file string) (*alpha.Client, error) {

	conf := dcl.NewConfig(dcl.WithUserAgent("dcl-test"), dcl.WithCredentialsFile(service_account_file))
	return alpha.NewClient(conf), nil
}
