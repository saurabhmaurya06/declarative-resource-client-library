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
package clouddeploy

import (
	"context"
	"fmt"
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/dcl"
	dclService "github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/clouddeploy"
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/unstructured"
)

type DeliveryPipeline struct{}

func DeliveryPipelineToUnstructured(r *dclService.DeliveryPipeline) *unstructured.Resource {
	u := &unstructured.Resource{
		STV: unstructured.ServiceTypeVersion{
			Service: "clouddeploy",
			Version: "ga",
			Type:    "DeliveryPipeline",
		},
		Object: make(map[string]interface{}),
	}
	if r.Annotations != nil {
		rAnnotations := make(map[string]interface{})
		for k, v := range r.Annotations {
			rAnnotations[k] = v
		}
		u.Object["annotations"] = rAnnotations
	}
	if r.Condition != nil && r.Condition != dclService.EmptyDeliveryPipelineCondition {
		rCondition := make(map[string]interface{})
		if r.Condition.PipelineReadyCondition != nil && r.Condition.PipelineReadyCondition != dclService.EmptyDeliveryPipelineConditionPipelineReadyCondition {
			rConditionPipelineReadyCondition := make(map[string]interface{})
			if r.Condition.PipelineReadyCondition.Status != nil {
				rConditionPipelineReadyCondition["status"] = *r.Condition.PipelineReadyCondition.Status
			}
			if r.Condition.PipelineReadyCondition.UpdateTime != nil {
				rConditionPipelineReadyCondition["updateTime"] = *r.Condition.PipelineReadyCondition.UpdateTime
			}
			rCondition["pipelineReadyCondition"] = rConditionPipelineReadyCondition
		}
		if r.Condition.TargetsPresentCondition != nil && r.Condition.TargetsPresentCondition != dclService.EmptyDeliveryPipelineConditionTargetsPresentCondition {
			rConditionTargetsPresentCondition := make(map[string]interface{})
			var rConditionTargetsPresentConditionMissingTargets []interface{}
			for _, rConditionTargetsPresentConditionMissingTargetsVal := range r.Condition.TargetsPresentCondition.MissingTargets {
				rConditionTargetsPresentConditionMissingTargets = append(rConditionTargetsPresentConditionMissingTargets, rConditionTargetsPresentConditionMissingTargetsVal)
			}
			rConditionTargetsPresentCondition["missingTargets"] = rConditionTargetsPresentConditionMissingTargets
			if r.Condition.TargetsPresentCondition.Status != nil {
				rConditionTargetsPresentCondition["status"] = *r.Condition.TargetsPresentCondition.Status
			}
			if r.Condition.TargetsPresentCondition.UpdateTime != nil {
				rConditionTargetsPresentCondition["updateTime"] = *r.Condition.TargetsPresentCondition.UpdateTime
			}
			rCondition["targetsPresentCondition"] = rConditionTargetsPresentCondition
		}
		u.Object["condition"] = rCondition
	}
	if r.CreateTime != nil {
		u.Object["createTime"] = *r.CreateTime
	}
	if r.Description != nil {
		u.Object["description"] = *r.Description
	}
	if r.Etag != nil {
		u.Object["etag"] = *r.Etag
	}
	if r.Labels != nil {
		rLabels := make(map[string]interface{})
		for k, v := range r.Labels {
			rLabels[k] = v
		}
		u.Object["labels"] = rLabels
	}
	if r.Location != nil {
		u.Object["location"] = *r.Location
	}
	if r.Name != nil {
		u.Object["name"] = *r.Name
	}
	if r.Project != nil {
		u.Object["project"] = *r.Project
	}
	if r.SerialPipeline != nil && r.SerialPipeline != dclService.EmptyDeliveryPipelineSerialPipeline {
		rSerialPipeline := make(map[string]interface{})
		var rSerialPipelineStages []interface{}
		for _, rSerialPipelineStagesVal := range r.SerialPipeline.Stages {
			rSerialPipelineStagesObject := make(map[string]interface{})
			var rSerialPipelineStagesValProfiles []interface{}
			for _, rSerialPipelineStagesValProfilesVal := range rSerialPipelineStagesVal.Profiles {
				rSerialPipelineStagesValProfiles = append(rSerialPipelineStagesValProfiles, rSerialPipelineStagesValProfilesVal)
			}
			rSerialPipelineStagesObject["profiles"] = rSerialPipelineStagesValProfiles
			if rSerialPipelineStagesVal.TargetId != nil {
				rSerialPipelineStagesObject["targetId"] = *rSerialPipelineStagesVal.TargetId
			}
			rSerialPipelineStages = append(rSerialPipelineStages, rSerialPipelineStagesObject)
		}
		rSerialPipeline["stages"] = rSerialPipelineStages
		u.Object["serialPipeline"] = rSerialPipeline
	}
	if r.Uid != nil {
		u.Object["uid"] = *r.Uid
	}
	if r.UpdateTime != nil {
		u.Object["updateTime"] = *r.UpdateTime
	}
	return u
}

func UnstructuredToDeliveryPipeline(u *unstructured.Resource) (*dclService.DeliveryPipeline, error) {
	r := &dclService.DeliveryPipeline{}
	if _, ok := u.Object["annotations"]; ok {
		if rAnnotations, ok := u.Object["annotations"].(map[string]interface{}); ok {
			m := make(map[string]string)
			for k, v := range rAnnotations {
				if s, ok := v.(string); ok {
					m[k] = s
				}
			}
			r.Annotations = m
		} else {
			return nil, fmt.Errorf("r.Annotations: expected map[string]interface{}")
		}
	}
	if _, ok := u.Object["condition"]; ok {
		if rCondition, ok := u.Object["condition"].(map[string]interface{}); ok {
			r.Condition = &dclService.DeliveryPipelineCondition{}
			if _, ok := rCondition["pipelineReadyCondition"]; ok {
				if rConditionPipelineReadyCondition, ok := rCondition["pipelineReadyCondition"].(map[string]interface{}); ok {
					r.Condition.PipelineReadyCondition = &dclService.DeliveryPipelineConditionPipelineReadyCondition{}
					if _, ok := rConditionPipelineReadyCondition["status"]; ok {
						if b, ok := rConditionPipelineReadyCondition["status"].(bool); ok {
							r.Condition.PipelineReadyCondition.Status = dcl.Bool(b)
						} else {
							return nil, fmt.Errorf("r.Condition.PipelineReadyCondition.Status: expected bool")
						}
					}
					if _, ok := rConditionPipelineReadyCondition["updateTime"]; ok {
						if s, ok := rConditionPipelineReadyCondition["updateTime"].(string); ok {
							r.Condition.PipelineReadyCondition.UpdateTime = dcl.String(s)
						} else {
							return nil, fmt.Errorf("r.Condition.PipelineReadyCondition.UpdateTime: expected string")
						}
					}
				} else {
					return nil, fmt.Errorf("r.Condition.PipelineReadyCondition: expected map[string]interface{}")
				}
			}
			if _, ok := rCondition["targetsPresentCondition"]; ok {
				if rConditionTargetsPresentCondition, ok := rCondition["targetsPresentCondition"].(map[string]interface{}); ok {
					r.Condition.TargetsPresentCondition = &dclService.DeliveryPipelineConditionTargetsPresentCondition{}
					if _, ok := rConditionTargetsPresentCondition["missingTargets"]; ok {
						if s, ok := rConditionTargetsPresentCondition["missingTargets"].([]interface{}); ok {
							for _, ss := range s {
								if strval, ok := ss.(string); ok {
									r.Condition.TargetsPresentCondition.MissingTargets = append(r.Condition.TargetsPresentCondition.MissingTargets, strval)
								}
							}
						} else {
							return nil, fmt.Errorf("r.Condition.TargetsPresentCondition.MissingTargets: expected []interface{}")
						}
					}
					if _, ok := rConditionTargetsPresentCondition["status"]; ok {
						if b, ok := rConditionTargetsPresentCondition["status"].(bool); ok {
							r.Condition.TargetsPresentCondition.Status = dcl.Bool(b)
						} else {
							return nil, fmt.Errorf("r.Condition.TargetsPresentCondition.Status: expected bool")
						}
					}
					if _, ok := rConditionTargetsPresentCondition["updateTime"]; ok {
						if s, ok := rConditionTargetsPresentCondition["updateTime"].(string); ok {
							r.Condition.TargetsPresentCondition.UpdateTime = dcl.String(s)
						} else {
							return nil, fmt.Errorf("r.Condition.TargetsPresentCondition.UpdateTime: expected string")
						}
					}
				} else {
					return nil, fmt.Errorf("r.Condition.TargetsPresentCondition: expected map[string]interface{}")
				}
			}
		} else {
			return nil, fmt.Errorf("r.Condition: expected map[string]interface{}")
		}
	}
	if _, ok := u.Object["createTime"]; ok {
		if s, ok := u.Object["createTime"].(string); ok {
			r.CreateTime = dcl.String(s)
		} else {
			return nil, fmt.Errorf("r.CreateTime: expected string")
		}
	}
	if _, ok := u.Object["description"]; ok {
		if s, ok := u.Object["description"].(string); ok {
			r.Description = dcl.String(s)
		} else {
			return nil, fmt.Errorf("r.Description: expected string")
		}
	}
	if _, ok := u.Object["etag"]; ok {
		if s, ok := u.Object["etag"].(string); ok {
			r.Etag = dcl.String(s)
		} else {
			return nil, fmt.Errorf("r.Etag: expected string")
		}
	}
	if _, ok := u.Object["labels"]; ok {
		if rLabels, ok := u.Object["labels"].(map[string]interface{}); ok {
			m := make(map[string]string)
			for k, v := range rLabels {
				if s, ok := v.(string); ok {
					m[k] = s
				}
			}
			r.Labels = m
		} else {
			return nil, fmt.Errorf("r.Labels: expected map[string]interface{}")
		}
	}
	if _, ok := u.Object["location"]; ok {
		if s, ok := u.Object["location"].(string); ok {
			r.Location = dcl.String(s)
		} else {
			return nil, fmt.Errorf("r.Location: expected string")
		}
	}
	if _, ok := u.Object["name"]; ok {
		if s, ok := u.Object["name"].(string); ok {
			r.Name = dcl.String(s)
		} else {
			return nil, fmt.Errorf("r.Name: expected string")
		}
	}
	if _, ok := u.Object["project"]; ok {
		if s, ok := u.Object["project"].(string); ok {
			r.Project = dcl.String(s)
		} else {
			return nil, fmt.Errorf("r.Project: expected string")
		}
	}
	if _, ok := u.Object["serialPipeline"]; ok {
		if rSerialPipeline, ok := u.Object["serialPipeline"].(map[string]interface{}); ok {
			r.SerialPipeline = &dclService.DeliveryPipelineSerialPipeline{}
			if _, ok := rSerialPipeline["stages"]; ok {
				if s, ok := rSerialPipeline["stages"].([]interface{}); ok {
					for _, o := range s {
						if objval, ok := o.(map[string]interface{}); ok {
							var rSerialPipelineStages dclService.DeliveryPipelineSerialPipelineStages
							if _, ok := objval["profiles"]; ok {
								if s, ok := objval["profiles"].([]interface{}); ok {
									for _, ss := range s {
										if strval, ok := ss.(string); ok {
											rSerialPipelineStages.Profiles = append(rSerialPipelineStages.Profiles, strval)
										}
									}
								} else {
									return nil, fmt.Errorf("rSerialPipelineStages.Profiles: expected []interface{}")
								}
							}
							if _, ok := objval["targetId"]; ok {
								if s, ok := objval["targetId"].(string); ok {
									rSerialPipelineStages.TargetId = dcl.String(s)
								} else {
									return nil, fmt.Errorf("rSerialPipelineStages.TargetId: expected string")
								}
							}
							r.SerialPipeline.Stages = append(r.SerialPipeline.Stages, rSerialPipelineStages)
						}
					}
				} else {
					return nil, fmt.Errorf("r.SerialPipeline.Stages: expected []interface{}")
				}
			}
		} else {
			return nil, fmt.Errorf("r.SerialPipeline: expected map[string]interface{}")
		}
	}
	if _, ok := u.Object["uid"]; ok {
		if s, ok := u.Object["uid"].(string); ok {
			r.Uid = dcl.String(s)
		} else {
			return nil, fmt.Errorf("r.Uid: expected string")
		}
	}
	if _, ok := u.Object["updateTime"]; ok {
		if s, ok := u.Object["updateTime"].(string); ok {
			r.UpdateTime = dcl.String(s)
		} else {
			return nil, fmt.Errorf("r.UpdateTime: expected string")
		}
	}
	return r, nil
}

func GetDeliveryPipeline(ctx context.Context, config *dcl.Config, u *unstructured.Resource) (*unstructured.Resource, error) {
	c := dclService.NewClient(config)
	r, err := UnstructuredToDeliveryPipeline(u)
	if err != nil {
		return nil, err
	}
	r, err = c.GetDeliveryPipeline(ctx, r)
	if err != nil {
		return nil, err
	}
	return DeliveryPipelineToUnstructured(r), nil
}

func ListDeliveryPipeline(ctx context.Context, config *dcl.Config, project string, location string) ([]*unstructured.Resource, error) {
	c := dclService.NewClient(config)
	l, err := c.ListDeliveryPipeline(ctx, project, location)
	if err != nil {
		return nil, err
	}
	var resources []*unstructured.Resource
	for {
		for _, r := range l.Items {
			resources = append(resources, DeliveryPipelineToUnstructured(r))
		}
		if !l.HasNext() {
			break
		}
		if err := l.Next(ctx, c); err != nil {
			return nil, err
		}
	}
	return resources, nil
}

func ApplyDeliveryPipeline(ctx context.Context, config *dcl.Config, u *unstructured.Resource, opts ...dcl.ApplyOption) (*unstructured.Resource, error) {
	c := dclService.NewClient(config)
	r, err := UnstructuredToDeliveryPipeline(u)
	if err != nil {
		return nil, err
	}
	if ush := unstructured.FetchStateHint(opts); ush != nil {
		sh, err := UnstructuredToDeliveryPipeline(ush)
		if err != nil {
			return nil, err
		}
		opts = append(opts, dcl.WithStateHint(sh))
	}
	r, err = c.ApplyDeliveryPipeline(ctx, r, opts...)
	if err != nil {
		return nil, err
	}
	return DeliveryPipelineToUnstructured(r), nil
}

func DeliveryPipelineHasDiff(ctx context.Context, config *dcl.Config, u *unstructured.Resource, opts ...dcl.ApplyOption) (bool, error) {
	c := dclService.NewClient(config)
	r, err := UnstructuredToDeliveryPipeline(u)
	if err != nil {
		return false, err
	}
	if ush := unstructured.FetchStateHint(opts); ush != nil {
		sh, err := UnstructuredToDeliveryPipeline(ush)
		if err != nil {
			return false, err
		}
		opts = append(opts, dcl.WithStateHint(sh))
	}
	opts = append(opts, dcl.WithLifecycleParam(dcl.BlockDestruction), dcl.WithLifecycleParam(dcl.BlockCreation), dcl.WithLifecycleParam(dcl.BlockModification))
	_, err = c.ApplyDeliveryPipeline(ctx, r, opts...)
	if err != nil {
		if _, ok := err.(dcl.ApplyInfeasibleError); ok {
			return true, nil
		}
		return false, err
	}
	return false, nil
}

func DeleteDeliveryPipeline(ctx context.Context, config *dcl.Config, u *unstructured.Resource) error {
	c := dclService.NewClient(config)
	r, err := UnstructuredToDeliveryPipeline(u)
	if err != nil {
		return err
	}
	return c.DeleteDeliveryPipeline(ctx, r)
}

func DeliveryPipelineID(u *unstructured.Resource) (string, error) {
	r, err := UnstructuredToDeliveryPipeline(u)
	if err != nil {
		return "", err
	}
	return r.ID()
}

func (r *DeliveryPipeline) STV() unstructured.ServiceTypeVersion {
	return unstructured.ServiceTypeVersion{
		"clouddeploy",
		"DeliveryPipeline",
		"ga",
	}
}

func (r *DeliveryPipeline) SetPolicyMember(ctx context.Context, config *dcl.Config, resource *unstructured.Resource, member *unstructured.Resource) (*unstructured.Resource, error) {
	return nil, unstructured.ErrNoSuchMethod
}

func (r *DeliveryPipeline) GetPolicyMember(ctx context.Context, config *dcl.Config, resource *unstructured.Resource, role, member string) (*unstructured.Resource, error) {
	return nil, unstructured.ErrNoSuchMethod
}

func (r *DeliveryPipeline) DeletePolicyMember(ctx context.Context, config *dcl.Config, resource *unstructured.Resource, member *unstructured.Resource) error {
	return unstructured.ErrNoSuchMethod
}

func (r *DeliveryPipeline) SetPolicy(ctx context.Context, config *dcl.Config, resource *unstructured.Resource, policy *unstructured.Resource) (*unstructured.Resource, error) {
	return nil, unstructured.ErrNoSuchMethod
}

func (r *DeliveryPipeline) SetPolicyWithEtag(ctx context.Context, config *dcl.Config, resource *unstructured.Resource, policy *unstructured.Resource) (*unstructured.Resource, error) {
	return nil, unstructured.ErrNoSuchMethod
}

func (r *DeliveryPipeline) GetPolicy(ctx context.Context, config *dcl.Config, resource *unstructured.Resource) (*unstructured.Resource, error) {
	return nil, unstructured.ErrNoSuchMethod
}

func (r *DeliveryPipeline) Get(ctx context.Context, config *dcl.Config, resource *unstructured.Resource) (*unstructured.Resource, error) {
	return GetDeliveryPipeline(ctx, config, resource)
}

func (r *DeliveryPipeline) Apply(ctx context.Context, config *dcl.Config, resource *unstructured.Resource, opts ...dcl.ApplyOption) (*unstructured.Resource, error) {
	return ApplyDeliveryPipeline(ctx, config, resource, opts...)
}

func (r *DeliveryPipeline) HasDiff(ctx context.Context, config *dcl.Config, resource *unstructured.Resource, opts ...dcl.ApplyOption) (bool, error) {
	return DeliveryPipelineHasDiff(ctx, config, resource, opts...)
}

func (r *DeliveryPipeline) Delete(ctx context.Context, config *dcl.Config, resource *unstructured.Resource) error {
	return DeleteDeliveryPipeline(ctx, config, resource)
}

func (r *DeliveryPipeline) ID(resource *unstructured.Resource) (string, error) {
	return DeliveryPipelineID(resource)
}

func init() {
	unstructured.Register(&DeliveryPipeline{})
}
