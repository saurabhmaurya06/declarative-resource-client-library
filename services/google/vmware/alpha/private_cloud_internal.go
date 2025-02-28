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
package alpha

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"strings"

	"github.com/GoogleCloudPlatform/declarative-resource-client-library/dcl"
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/dcl/operations"
)

func (r *PrivateCloud) validate() error {

	if err := dcl.Required(r, "name"); err != nil {
		return err
	}
	if err := dcl.Required(r, "networkConfig"); err != nil {
		return err
	}
	if err := dcl.Required(r, "managementCluster"); err != nil {
		return err
	}
	if err := dcl.RequiredParameter(r.Project, "Project"); err != nil {
		return err
	}
	if err := dcl.RequiredParameter(r.Location, "Location"); err != nil {
		return err
	}
	if !dcl.IsEmptyValueIndirect(r.NetworkConfig) {
		if err := r.NetworkConfig.validate(); err != nil {
			return err
		}
	}
	if !dcl.IsEmptyValueIndirect(r.ManagementCluster) {
		if err := r.ManagementCluster.validate(); err != nil {
			return err
		}
	}
	if !dcl.IsEmptyValueIndirect(r.Hcx) {
		if err := r.Hcx.validate(); err != nil {
			return err
		}
	}
	if !dcl.IsEmptyValueIndirect(r.Nsx) {
		if err := r.Nsx.validate(); err != nil {
			return err
		}
	}
	if !dcl.IsEmptyValueIndirect(r.Vcenter) {
		if err := r.Vcenter.validate(); err != nil {
			return err
		}
	}
	return nil
}
func (r *PrivateCloudNetworkConfig) validate() error {
	if err := dcl.Required(r, "network"); err != nil {
		return err
	}
	if err := dcl.Required(r, "managementCidr"); err != nil {
		return err
	}
	return nil
}
func (r *PrivateCloudManagementCluster) validate() error {
	if err := dcl.Required(r, "clusterId"); err != nil {
		return err
	}
	if err := dcl.Required(r, "nodeTypeId"); err != nil {
		return err
	}
	if err := dcl.Required(r, "nodeCount"); err != nil {
		return err
	}
	return nil
}
func (r *PrivateCloudConditions) validate() error {
	return nil
}
func (r *PrivateCloudHcx) validate() error {
	return nil
}
func (r *PrivateCloudNsx) validate() error {
	return nil
}
func (r *PrivateCloudVcenter) validate() error {
	return nil
}
func (r *PrivateCloud) basePath() string {
	params := map[string]interface{}{}
	return dcl.Nprintf("https://vmwareengine.googleapis.com/v1/", params)
}

func (r *PrivateCloud) getURL(userBasePath string) (string, error) {
	nr := r.urlNormalized()
	params := map[string]interface{}{
		"project":  dcl.ValueOrEmptyString(nr.Project),
		"location": dcl.ValueOrEmptyString(nr.Location),
		"name":     dcl.ValueOrEmptyString(nr.Name),
	}
	return dcl.URL("projects/{{project}}/locations/{{location}}/privateClouds/{{name}}", nr.basePath(), userBasePath, params), nil
}

func (r *PrivateCloud) listURL(userBasePath string) (string, error) {
	nr := r.urlNormalized()
	params := map[string]interface{}{
		"project":  dcl.ValueOrEmptyString(nr.Project),
		"location": dcl.ValueOrEmptyString(nr.Location),
	}
	return dcl.URL("projects/{{project}}/locations/{{location}}/privateClouds", nr.basePath(), userBasePath, params), nil

}

func (r *PrivateCloud) createURL(userBasePath string) (string, error) {
	nr := r.urlNormalized()
	params := map[string]interface{}{
		"project":  dcl.ValueOrEmptyString(nr.Project),
		"location": dcl.ValueOrEmptyString(nr.Location),
		"name":     dcl.ValueOrEmptyString(nr.Name),
	}
	return dcl.URL("projects/{{project}}/locations/{{location}}/privateClouds?privateCloudId={{name}}", nr.basePath(), userBasePath, params), nil

}

func (r *PrivateCloud) deleteURL(userBasePath string) (string, error) {
	nr := r.urlNormalized()
	params := map[string]interface{}{
		"project":  dcl.ValueOrEmptyString(nr.Project),
		"location": dcl.ValueOrEmptyString(nr.Location),
		"name":     dcl.ValueOrEmptyString(nr.Name),
	}
	return dcl.URL("projects/{{project}}/locations/{{location}}/privateClouds/{{name}}", nr.basePath(), userBasePath, params), nil
}

func (r *PrivateCloud) SetPolicyURL(userBasePath string) string {
	nr := r.urlNormalized()
	fields := map[string]interface{}{
		"project":  *nr.Project,
		"location": *nr.Location,
		"name":     *nr.Name,
	}
	return dcl.URL("projects/{{project}}/locations/{{location}}/privateClouds/{{name}}:setIamPolicy", nr.basePath(), userBasePath, fields)
}

func (r *PrivateCloud) SetPolicyVerb() string {
	return "POST"
}

func (r *PrivateCloud) getPolicyURL(userBasePath string) string {
	nr := r.urlNormalized()
	fields := map[string]interface{}{
		"project":  *nr.Project,
		"location": *nr.Location,
		"name":     *nr.Name,
	}
	return dcl.URL("projects/{{project}}/locations/{{location}}/privateClouds/{{name}}:getIamPolicy", nr.basePath(), userBasePath, fields)
}

func (r *PrivateCloud) IAMPolicyVersion() int {
	return 3
}

// privateCloudApiOperation represents a mutable operation in the underlying REST
// API such as Create, Update, or Delete.
type privateCloudApiOperation interface {
	do(context.Context, *PrivateCloud, *Client) error
}

// newUpdatePrivateCloudUpdateClusterRequest creates a request for an
// PrivateCloud resource's UpdateCluster update type by filling in the update
// fields based on the intended state of the resource.
func newUpdatePrivateCloudUpdateClusterRequest(ctx context.Context, f *PrivateCloud, c *Client) (map[string]interface{}, error) {
	req := map[string]interface{}{}
	res := f
	_ = res

	if v, err := expandPrivateCloudManagementCluster(c, f.ManagementCluster, res); err != nil {
		return nil, fmt.Errorf("error expanding ManagementCluster into managementCluster: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		req["managementCluster"] = v
	}
	return req, nil
}

// marshalUpdatePrivateCloudUpdateClusterRequest converts the update into
// the final JSON request body.
func marshalUpdatePrivateCloudUpdateClusterRequest(c *Client, m map[string]interface{}) ([]byte, error) {

	m = encodePrivateCloudUpdateClusterRequest(m)
	return json.Marshal(m)
}

type updatePrivateCloudUpdateClusterOperation struct {
	// If the update operation has the REQUIRES_APPLY_OPTIONS trait, this will be populated.
	// Usually it will be nil - this is to prevent us from accidentally depending on apply
	// options, which should usually be unnecessary.
	ApplyOptions []dcl.ApplyOption
	FieldDiffs   []*dcl.FieldDiff
}

// do creates a request and sends it to the appropriate URL. In most operations,
// do will transcribe a subset of the resource into a request object and send a
// PUT request to a single URL.

func (op *updatePrivateCloudUpdateClusterOperation) do(ctx context.Context, r *PrivateCloud, c *Client) error {
	_, err := c.GetPrivateCloud(ctx, r)
	if err != nil {
		return err
	}

	u, err := r.updateURL(c.Config.BasePath, "UpdateCluster")
	if err != nil {
		return err
	}
	mask := op.UpdateMask()
	u, err = dcl.AddQueryParams(u, map[string]string{"updateMask": mask})
	if err != nil {
		return err
	}

	req, err := newUpdatePrivateCloudUpdateClusterRequest(ctx, r, c)
	if err != nil {
		return err
	}

	c.Config.Logger.InfoWithContextf(ctx, "Created update: %#v", req)
	body, err := marshalUpdatePrivateCloudUpdateClusterRequest(c, req)
	if err != nil {
		return err
	}
	resp, err := dcl.SendRequest(ctx, c.Config, "PATCH", u, bytes.NewBuffer(body), c.Config.RetryProvider)
	if err != nil {
		return err
	}

	var o operations.StandardGCPOperation
	if err := dcl.ParseResponse(resp.Response, &o); err != nil {
		return err
	}
	err = o.Wait(context.WithValue(ctx, dcl.DoNotLogRequestsKey, true), c.Config, r.basePath(), "GET")

	if err != nil {
		return err
	}

	return nil
}

// newUpdatePrivateCloudUpdatePrivateCloudRequest creates a request for an
// PrivateCloud resource's UpdatePrivateCloud update type by filling in the update
// fields based on the intended state of the resource.
func newUpdatePrivateCloudUpdatePrivateCloudRequest(ctx context.Context, f *PrivateCloud, c *Client) (map[string]interface{}, error) {
	req := map[string]interface{}{}
	res := f
	_ = res

	if v := f.Description; !dcl.IsEmptyValueIndirect(v) {
		req["description"] = v
	}
	req["name"] = fmt.Sprintf("projects/%s/locations/%s/privateClouds/%s", *f.Project, *f.Location, *f.Name)

	return req, nil
}

// marshalUpdatePrivateCloudUpdatePrivateCloudRequest converts the update into
// the final JSON request body.
func marshalUpdatePrivateCloudUpdatePrivateCloudRequest(c *Client, m map[string]interface{}) ([]byte, error) {

	return json.Marshal(m)
}

type updatePrivateCloudUpdatePrivateCloudOperation struct {
	// If the update operation has the REQUIRES_APPLY_OPTIONS trait, this will be populated.
	// Usually it will be nil - this is to prevent us from accidentally depending on apply
	// options, which should usually be unnecessary.
	ApplyOptions []dcl.ApplyOption
	FieldDiffs   []*dcl.FieldDiff
}

// do creates a request and sends it to the appropriate URL. In most operations,
// do will transcribe a subset of the resource into a request object and send a
// PUT request to a single URL.

func (op *updatePrivateCloudUpdatePrivateCloudOperation) do(ctx context.Context, r *PrivateCloud, c *Client) error {
	_, err := c.GetPrivateCloud(ctx, r)
	if err != nil {
		return err
	}

	u, err := r.updateURL(c.Config.BasePath, "UpdatePrivateCloud")
	if err != nil {
		return err
	}
	mask := dcl.UpdateMask(op.FieldDiffs)
	u, err = dcl.AddQueryParams(u, map[string]string{"updateMask": mask})
	if err != nil {
		return err
	}

	req, err := newUpdatePrivateCloudUpdatePrivateCloudRequest(ctx, r, c)
	if err != nil {
		return err
	}

	c.Config.Logger.InfoWithContextf(ctx, "Created update: %#v", req)
	body, err := marshalUpdatePrivateCloudUpdatePrivateCloudRequest(c, req)
	if err != nil {
		return err
	}
	resp, err := dcl.SendRequest(ctx, c.Config, "PATCH", u, bytes.NewBuffer(body), c.Config.RetryProvider)
	if err != nil {
		return err
	}

	var o operations.StandardGCPOperation
	if err := dcl.ParseResponse(resp.Response, &o); err != nil {
		return err
	}
	err = o.Wait(context.WithValue(ctx, dcl.DoNotLogRequestsKey, true), c.Config, r.basePath(), "GET")

	if err != nil {
		return err
	}

	return nil
}

func (c *Client) listPrivateCloudRaw(ctx context.Context, r *PrivateCloud, pageToken string, pageSize int32) ([]byte, error) {
	u, err := r.urlNormalized().listURL(c.Config.BasePath)
	if err != nil {
		return nil, err
	}

	m := make(map[string]string)
	if pageToken != "" {
		m["pageToken"] = pageToken
	}

	if pageSize != PrivateCloudMaxPage {
		m["pageSize"] = fmt.Sprintf("%v", pageSize)
	}

	u, err = dcl.AddQueryParams(u, m)
	if err != nil {
		return nil, err
	}
	resp, err := dcl.SendRequest(ctx, c.Config, "GET", u, &bytes.Buffer{}, c.Config.RetryProvider)
	if err != nil {
		return nil, err
	}
	defer resp.Response.Body.Close()
	return ioutil.ReadAll(resp.Response.Body)
}

type listPrivateCloudOperation struct {
	PrivateClouds []map[string]interface{} `json:"privateClouds"`
	Token         string                   `json:"nextPageToken"`
}

func (c *Client) listPrivateCloud(ctx context.Context, r *PrivateCloud, pageToken string, pageSize int32) ([]*PrivateCloud, string, error) {
	b, err := c.listPrivateCloudRaw(ctx, r, pageToken, pageSize)
	if err != nil {
		return nil, "", err
	}

	var m listPrivateCloudOperation
	if err := json.Unmarshal(b, &m); err != nil {
		return nil, "", err
	}

	var l []*PrivateCloud
	for _, v := range m.PrivateClouds {
		res, err := unmarshalMapPrivateCloud(v, c, r)
		if err != nil {
			return nil, m.Token, err
		}
		res.Project = r.Project
		res.Location = r.Location
		l = append(l, res)
	}

	return l, m.Token, nil
}

func (c *Client) deleteAllPrivateCloud(ctx context.Context, f func(*PrivateCloud) bool, resources []*PrivateCloud) error {
	var errors []string
	for _, res := range resources {
		if f(res) {
			// We do not want deleteAll to fail on a deletion or else it will stop deleting other resources.
			err := c.DeletePrivateCloud(ctx, res)
			if err != nil {
				errors = append(errors, err.Error())
			}
		}
	}
	if len(errors) > 0 {
		return fmt.Errorf("%v", strings.Join(errors, "\n"))
	} else {
		return nil
	}
}

type deletePrivateCloudOperation struct{}

func (op *deletePrivateCloudOperation) do(ctx context.Context, r *PrivateCloud, c *Client) error {
	r, err := c.GetPrivateCloud(ctx, r)
	if err != nil {
		if dcl.IsNotFound(err) {
			c.Config.Logger.InfoWithContextf(ctx, "PrivateCloud not found, returning. Original error: %v", err)
			return nil
		}
		c.Config.Logger.WarningWithContextf(ctx, "GetPrivateCloud checking for existence. error: %v", err)
		return err
	}

	u, err := r.deleteURL(c.Config.BasePath)
	if err != nil {
		return err
	}

	// Delete should never have a body
	body := &bytes.Buffer{}
	_, err = dcl.SendRequest(ctx, c.Config, "DELETE", u, body, c.Config.RetryProvider)
	if err != nil {
		return fmt.Errorf("failed to delete PrivateCloud: %w", err)
	}
	return nil
}

// Create operations are similar to Update operations, although they do not have
// specific request objects. The Create request object is the json encoding of
// the resource, which is modified by res.marshal to form the base request body.
type createPrivateCloudOperation struct {
	response map[string]interface{}
}

func (op *createPrivateCloudOperation) FirstResponse() (map[string]interface{}, bool) {
	return op.response, len(op.response) > 0
}

func (op *createPrivateCloudOperation) do(ctx context.Context, r *PrivateCloud, c *Client) error {
	c.Config.Logger.InfoWithContextf(ctx, "Attempting to create %v", r)
	u, err := r.createURL(c.Config.BasePath)
	if err != nil {
		return err
	}

	req, err := r.marshal(c)
	if err != nil {
		return err
	}
	resp, err := dcl.SendRequest(ctx, c.Config, "POST", u, bytes.NewBuffer(req), c.Config.RetryProvider)
	if err != nil {
		return err
	}
	// wait for object to be created.
	var o operations.StandardGCPOperation
	if err := dcl.ParseResponse(resp.Response, &o); err != nil {
		return err
	}
	if err := o.Wait(context.WithValue(ctx, dcl.DoNotLogRequestsKey, true), c.Config, r.basePath(), "GET"); err != nil {
		c.Config.Logger.Warningf("Creation failed after waiting for operation: %v", err)
		return err
	}
	c.Config.Logger.InfoWithContextf(ctx, "Successfully waited for operation")
	op.response, _ = o.FirstResponse()

	if _, err := c.GetPrivateCloud(ctx, r); err != nil {
		c.Config.Logger.WarningWithContextf(ctx, "get returned error: %v", err)
		return err
	}

	return nil
}

func (c *Client) getPrivateCloudRaw(ctx context.Context, r *PrivateCloud) ([]byte, error) {

	u, err := r.getURL(c.Config.BasePath)
	if err != nil {
		return nil, err
	}
	resp, err := dcl.SendRequest(ctx, c.Config, "GET", u, &bytes.Buffer{}, c.Config.RetryProvider)
	if err != nil {
		return nil, err
	}
	defer resp.Response.Body.Close()
	b, err := ioutil.ReadAll(resp.Response.Body)
	if err != nil {
		return nil, err
	}

	return b, nil
}

func (c *Client) privateCloudDiffsForRawDesired(ctx context.Context, rawDesired *PrivateCloud, opts ...dcl.ApplyOption) (initial, desired *PrivateCloud, diffs []*dcl.FieldDiff, err error) {
	c.Config.Logger.InfoWithContext(ctx, "Fetching initial state...")
	// First, let us see if the user provided a state hint.  If they did, we will start fetching based on that.
	var fetchState *PrivateCloud
	if sh := dcl.FetchStateHint(opts); sh != nil {
		if r, ok := sh.(*PrivateCloud); !ok {
			c.Config.Logger.WarningWithContextf(ctx, "Initial state hint was of the wrong type; expected PrivateCloud, got %T", sh)
		} else {
			fetchState = r
		}
	}
	if fetchState == nil {
		fetchState = rawDesired
	}

	// 1.2: Retrieval of raw initial state from API
	rawInitial, err := c.GetPrivateCloud(ctx, fetchState)
	if rawInitial == nil {
		if !dcl.IsNotFound(err) {
			c.Config.Logger.WarningWithContextf(ctx, "Failed to retrieve whether a PrivateCloud resource already exists: %s", err)
			return nil, nil, nil, fmt.Errorf("failed to retrieve PrivateCloud resource: %v", err)
		}
		c.Config.Logger.InfoWithContext(ctx, "Found that PrivateCloud resource did not exist.")
		// Perform canonicalization to pick up defaults.
		desired, err = canonicalizePrivateCloudDesiredState(rawDesired, rawInitial)
		return nil, desired, nil, err
	}
	c.Config.Logger.InfoWithContextf(ctx, "Found initial state for PrivateCloud: %v", rawInitial)
	c.Config.Logger.InfoWithContextf(ctx, "Initial desired state for PrivateCloud: %v", rawDesired)

	// The Get call applies postReadExtract and so the result may contain fields that are not part of API version.
	if err := extractPrivateCloudFields(rawInitial); err != nil {
		return nil, nil, nil, err
	}

	// 1.3: Canonicalize raw initial state into initial state.
	initial, err = canonicalizePrivateCloudInitialState(rawInitial, rawDesired)
	if err != nil {
		return nil, nil, nil, err
	}
	c.Config.Logger.InfoWithContextf(ctx, "Canonicalized initial state for PrivateCloud: %v", initial)

	// 1.4: Canonicalize raw desired state into desired state.
	desired, err = canonicalizePrivateCloudDesiredState(rawDesired, rawInitial, opts...)
	if err != nil {
		return nil, nil, nil, err
	}
	c.Config.Logger.InfoWithContextf(ctx, "Canonicalized desired state for PrivateCloud: %v", desired)

	// 2.1: Comparison of initial and desired state.
	diffs, err = diffPrivateCloud(c, desired, initial, opts...)
	return initial, desired, diffs, err
}

func canonicalizePrivateCloudInitialState(rawInitial, rawDesired *PrivateCloud) (*PrivateCloud, error) {
	// TODO(magic-modules-eng): write canonicalizer once relevant traits are added.
	return rawInitial, nil
}

/*
* Canonicalizers
*
* These are responsible for converting either a user-specified config or a
* GCP API response to a standard format that can be used for difference checking.
* */

func canonicalizePrivateCloudDesiredState(rawDesired, rawInitial *PrivateCloud, opts ...dcl.ApplyOption) (*PrivateCloud, error) {

	if rawInitial == nil {
		// Since the initial state is empty, the desired state is all we have.
		// We canonicalize the remaining nested objects with nil to pick up defaults.
		rawDesired.NetworkConfig = canonicalizePrivateCloudNetworkConfig(rawDesired.NetworkConfig, nil, opts...)
		rawDesired.ManagementCluster = canonicalizePrivateCloudManagementCluster(rawDesired.ManagementCluster, nil, opts...)
		rawDesired.Hcx = canonicalizePrivateCloudHcx(rawDesired.Hcx, nil, opts...)
		rawDesired.Nsx = canonicalizePrivateCloudNsx(rawDesired.Nsx, nil, opts...)
		rawDesired.Vcenter = canonicalizePrivateCloudVcenter(rawDesired.Vcenter, nil, opts...)

		return rawDesired, nil
	}
	canonicalDesired := &PrivateCloud{}
	if dcl.PartialSelfLinkToSelfLink(rawDesired.Name, rawInitial.Name) {
		canonicalDesired.Name = rawInitial.Name
	} else {
		canonicalDesired.Name = rawDesired.Name
	}
	canonicalDesired.NetworkConfig = canonicalizePrivateCloudNetworkConfig(rawDesired.NetworkConfig, rawInitial.NetworkConfig, opts...)
	canonicalDesired.ManagementCluster = canonicalizePrivateCloudManagementCluster(rawDesired.ManagementCluster, rawInitial.ManagementCluster, opts...)
	if dcl.StringCanonicalize(rawDesired.Description, rawInitial.Description) {
		canonicalDesired.Description = rawInitial.Description
	} else {
		canonicalDesired.Description = rawDesired.Description
	}
	if dcl.NameToSelfLink(rawDesired.Project, rawInitial.Project) {
		canonicalDesired.Project = rawInitial.Project
	} else {
		canonicalDesired.Project = rawDesired.Project
	}
	if dcl.NameToSelfLink(rawDesired.Location, rawInitial.Location) {
		canonicalDesired.Location = rawInitial.Location
	} else {
		canonicalDesired.Location = rawDesired.Location
	}

	return canonicalDesired, nil
}

func canonicalizePrivateCloudNewState(c *Client, rawNew, rawDesired *PrivateCloud) (*PrivateCloud, error) {

	if dcl.IsNotReturnedByServer(rawNew.Name) && dcl.IsNotReturnedByServer(rawDesired.Name) {
		rawNew.Name = rawDesired.Name
	} else {
		if dcl.PartialSelfLinkToSelfLink(rawDesired.Name, rawNew.Name) {
			rawNew.Name = rawDesired.Name
		}
	}

	if dcl.IsNotReturnedByServer(rawNew.CreateTime) && dcl.IsNotReturnedByServer(rawDesired.CreateTime) {
		rawNew.CreateTime = rawDesired.CreateTime
	} else {
	}

	if dcl.IsNotReturnedByServer(rawNew.UpdateTime) && dcl.IsNotReturnedByServer(rawDesired.UpdateTime) {
		rawNew.UpdateTime = rawDesired.UpdateTime
	} else {
	}

	if dcl.IsNotReturnedByServer(rawNew.DeleteTime) && dcl.IsNotReturnedByServer(rawDesired.DeleteTime) {
		rawNew.DeleteTime = rawDesired.DeleteTime
	} else {
	}

	if dcl.IsNotReturnedByServer(rawNew.ExpireTime) && dcl.IsNotReturnedByServer(rawDesired.ExpireTime) {
		rawNew.ExpireTime = rawDesired.ExpireTime
	} else {
	}

	if dcl.IsNotReturnedByServer(rawNew.State) && dcl.IsNotReturnedByServer(rawDesired.State) {
		rawNew.State = rawDesired.State
	} else {
	}

	if dcl.IsNotReturnedByServer(rawNew.NetworkConfig) && dcl.IsNotReturnedByServer(rawDesired.NetworkConfig) {
		rawNew.NetworkConfig = rawDesired.NetworkConfig
	} else {
		rawNew.NetworkConfig = canonicalizeNewPrivateCloudNetworkConfig(c, rawDesired.NetworkConfig, rawNew.NetworkConfig)
	}

	if dcl.IsNotReturnedByServer(rawNew.ManagementCluster) && dcl.IsNotReturnedByServer(rawDesired.ManagementCluster) {
		rawNew.ManagementCluster = rawDesired.ManagementCluster
	} else {
		rawNew.ManagementCluster = rawDesired.ManagementCluster
	}

	if dcl.IsNotReturnedByServer(rawNew.Description) && dcl.IsNotReturnedByServer(rawDesired.Description) {
		rawNew.Description = rawDesired.Description
	} else {
		if dcl.StringCanonicalize(rawDesired.Description, rawNew.Description) {
			rawNew.Description = rawDesired.Description
		}
	}

	if dcl.IsNotReturnedByServer(rawNew.Conditions) && dcl.IsNotReturnedByServer(rawDesired.Conditions) {
		rawNew.Conditions = rawDesired.Conditions
	} else {
		rawNew.Conditions = canonicalizeNewPrivateCloudConditionsSlice(c, rawDesired.Conditions, rawNew.Conditions)
	}

	if dcl.IsNotReturnedByServer(rawNew.Hcx) && dcl.IsNotReturnedByServer(rawDesired.Hcx) {
		rawNew.Hcx = rawDesired.Hcx
	} else {
		rawNew.Hcx = canonicalizeNewPrivateCloudHcx(c, rawDesired.Hcx, rawNew.Hcx)
	}

	if dcl.IsNotReturnedByServer(rawNew.Nsx) && dcl.IsNotReturnedByServer(rawDesired.Nsx) {
		rawNew.Nsx = rawDesired.Nsx
	} else {
		rawNew.Nsx = canonicalizeNewPrivateCloudNsx(c, rawDesired.Nsx, rawNew.Nsx)
	}

	if dcl.IsNotReturnedByServer(rawNew.Vcenter) && dcl.IsNotReturnedByServer(rawDesired.Vcenter) {
		rawNew.Vcenter = rawDesired.Vcenter
	} else {
		rawNew.Vcenter = canonicalizeNewPrivateCloudVcenter(c, rawDesired.Vcenter, rawNew.Vcenter)
	}

	rawNew.Project = rawDesired.Project

	rawNew.Location = rawDesired.Location

	return rawNew, nil
}

func canonicalizePrivateCloudNetworkConfig(des, initial *PrivateCloudNetworkConfig, opts ...dcl.ApplyOption) *PrivateCloudNetworkConfig {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	cDes := &PrivateCloudNetworkConfig{}

	if canonicalizePrivateCloudNetwork(des.Network, initial.Network) || dcl.IsZeroValue(des.Network) {
		cDes.Network = initial.Network
	} else {
		cDes.Network = des.Network
	}
	if dcl.StringCanonicalize(des.ManagementCidr, initial.ManagementCidr) || dcl.IsZeroValue(des.ManagementCidr) {
		cDes.ManagementCidr = initial.ManagementCidr
	} else {
		cDes.ManagementCidr = des.ManagementCidr
	}

	return cDes
}

func canonicalizePrivateCloudNetworkConfigSlice(des, initial []PrivateCloudNetworkConfig, opts ...dcl.ApplyOption) []PrivateCloudNetworkConfig {
	if dcl.IsEmptyValueIndirect(des) {
		return initial
	}

	if len(des) != len(initial) {

		items := make([]PrivateCloudNetworkConfig, 0, len(des))
		for _, d := range des {
			cd := canonicalizePrivateCloudNetworkConfig(&d, nil, opts...)
			if cd != nil {
				items = append(items, *cd)
			}
		}
		return items
	}

	items := make([]PrivateCloudNetworkConfig, 0, len(des))
	for i, d := range des {
		cd := canonicalizePrivateCloudNetworkConfig(&d, &initial[i], opts...)
		if cd != nil {
			items = append(items, *cd)
		}
	}
	return items

}

func canonicalizeNewPrivateCloudNetworkConfig(c *Client, des, nw *PrivateCloudNetworkConfig) *PrivateCloudNetworkConfig {

	if des == nil {
		return nw
	}

	if nw == nil {
		if dcl.IsNotReturnedByServer(des) {
			c.Config.Logger.Info("Found explicitly empty value for PrivateCloudNetworkConfig while comparing non-nil desired to nil actual.  Returning desired object.")
			return des
		}
		return nil
	}

	if canonicalizePrivateCloudNetwork(des.Network, nw.Network) {
		nw.Network = des.Network
	}
	if dcl.StringCanonicalize(des.ServiceNetwork, nw.ServiceNetwork) {
		nw.ServiceNetwork = des.ServiceNetwork
	}
	if dcl.StringCanonicalize(des.ManagementCidr, nw.ManagementCidr) {
		nw.ManagementCidr = des.ManagementCidr
	}

	return nw
}

func canonicalizeNewPrivateCloudNetworkConfigSet(c *Client, des, nw []PrivateCloudNetworkConfig) []PrivateCloudNetworkConfig {
	if des == nil {
		return nw
	}
	var reorderedNew []PrivateCloudNetworkConfig
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if diffs, _ := comparePrivateCloudNetworkConfigNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
				matchedNew = idx
				break
			}
		}
		if matchedNew != -1 {
			reorderedNew = append(reorderedNew, nw[matchedNew])
			nw = append(nw[:matchedNew], nw[matchedNew+1:]...)
		}
	}
	reorderedNew = append(reorderedNew, nw...)

	return reorderedNew
}

func canonicalizeNewPrivateCloudNetworkConfigSlice(c *Client, des, nw []PrivateCloudNetworkConfig) []PrivateCloudNetworkConfig {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []PrivateCloudNetworkConfig
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewPrivateCloudNetworkConfig(c, &d, &n))
	}

	return items
}

func canonicalizePrivateCloudManagementCluster(des, initial *PrivateCloudManagementCluster, opts ...dcl.ApplyOption) *PrivateCloudManagementCluster {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	cDes := &PrivateCloudManagementCluster{}

	if dcl.StringCanonicalize(des.ClusterId, initial.ClusterId) || dcl.IsZeroValue(des.ClusterId) {
		cDes.ClusterId = initial.ClusterId
	} else {
		cDes.ClusterId = des.ClusterId
	}
	if dcl.StringCanonicalize(des.NodeTypeId, initial.NodeTypeId) || dcl.IsZeroValue(des.NodeTypeId) {
		cDes.NodeTypeId = initial.NodeTypeId
	} else {
		cDes.NodeTypeId = des.NodeTypeId
	}
	if dcl.IsZeroValue(des.NodeCount) || (dcl.IsEmptyValueIndirect(des.NodeCount) && dcl.IsEmptyValueIndirect(initial.NodeCount)) {
		// Desired and initial values are equivalent, so set canonical desired value to initial value.
		cDes.NodeCount = initial.NodeCount
	} else {
		cDes.NodeCount = des.NodeCount
	}

	return cDes
}

func canonicalizePrivateCloudManagementClusterSlice(des, initial []PrivateCloudManagementCluster, opts ...dcl.ApplyOption) []PrivateCloudManagementCluster {
	if dcl.IsEmptyValueIndirect(des) {
		return initial
	}

	if len(des) != len(initial) {

		items := make([]PrivateCloudManagementCluster, 0, len(des))
		for _, d := range des {
			cd := canonicalizePrivateCloudManagementCluster(&d, nil, opts...)
			if cd != nil {
				items = append(items, *cd)
			}
		}
		return items
	}

	items := make([]PrivateCloudManagementCluster, 0, len(des))
	for i, d := range des {
		cd := canonicalizePrivateCloudManagementCluster(&d, &initial[i], opts...)
		if cd != nil {
			items = append(items, *cd)
		}
	}
	return items

}

func canonicalizeNewPrivateCloudManagementCluster(c *Client, des, nw *PrivateCloudManagementCluster) *PrivateCloudManagementCluster {

	if des == nil {
		return nw
	}

	if nw == nil {
		if dcl.IsNotReturnedByServer(des) {
			c.Config.Logger.Info("Found explicitly empty value for PrivateCloudManagementCluster while comparing non-nil desired to nil actual.  Returning desired object.")
			return des
		}
		return nil
	}

	if dcl.StringCanonicalize(des.ClusterId, nw.ClusterId) {
		nw.ClusterId = des.ClusterId
	}
	if dcl.StringCanonicalize(des.NodeTypeId, nw.NodeTypeId) {
		nw.NodeTypeId = des.NodeTypeId
	}

	return nw
}

func canonicalizeNewPrivateCloudManagementClusterSet(c *Client, des, nw []PrivateCloudManagementCluster) []PrivateCloudManagementCluster {
	if des == nil {
		return nw
	}
	var reorderedNew []PrivateCloudManagementCluster
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if diffs, _ := comparePrivateCloudManagementClusterNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
				matchedNew = idx
				break
			}
		}
		if matchedNew != -1 {
			reorderedNew = append(reorderedNew, nw[matchedNew])
			nw = append(nw[:matchedNew], nw[matchedNew+1:]...)
		}
	}
	reorderedNew = append(reorderedNew, nw...)

	return reorderedNew
}

func canonicalizeNewPrivateCloudManagementClusterSlice(c *Client, des, nw []PrivateCloudManagementCluster) []PrivateCloudManagementCluster {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []PrivateCloudManagementCluster
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewPrivateCloudManagementCluster(c, &d, &n))
	}

	return items
}

func canonicalizePrivateCloudConditions(des, initial *PrivateCloudConditions, opts ...dcl.ApplyOption) *PrivateCloudConditions {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	cDes := &PrivateCloudConditions{}

	return cDes
}

func canonicalizePrivateCloudConditionsSlice(des, initial []PrivateCloudConditions, opts ...dcl.ApplyOption) []PrivateCloudConditions {
	if dcl.IsEmptyValueIndirect(des) {
		return initial
	}

	if len(des) != len(initial) {

		items := make([]PrivateCloudConditions, 0, len(des))
		for _, d := range des {
			cd := canonicalizePrivateCloudConditions(&d, nil, opts...)
			if cd != nil {
				items = append(items, *cd)
			}
		}
		return items
	}

	items := make([]PrivateCloudConditions, 0, len(des))
	for i, d := range des {
		cd := canonicalizePrivateCloudConditions(&d, &initial[i], opts...)
		if cd != nil {
			items = append(items, *cd)
		}
	}
	return items

}

func canonicalizeNewPrivateCloudConditions(c *Client, des, nw *PrivateCloudConditions) *PrivateCloudConditions {

	if des == nil {
		return nw
	}

	if nw == nil {
		if dcl.IsNotReturnedByServer(des) {
			c.Config.Logger.Info("Found explicitly empty value for PrivateCloudConditions while comparing non-nil desired to nil actual.  Returning desired object.")
			return des
		}
		return nil
	}

	if dcl.StringCanonicalize(des.Code, nw.Code) {
		nw.Code = des.Code
	}
	if dcl.StringCanonicalize(des.Message, nw.Message) {
		nw.Message = des.Message
	}

	return nw
}

func canonicalizeNewPrivateCloudConditionsSet(c *Client, des, nw []PrivateCloudConditions) []PrivateCloudConditions {
	if des == nil {
		return nw
	}
	var reorderedNew []PrivateCloudConditions
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if diffs, _ := comparePrivateCloudConditionsNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
				matchedNew = idx
				break
			}
		}
		if matchedNew != -1 {
			reorderedNew = append(reorderedNew, nw[matchedNew])
			nw = append(nw[:matchedNew], nw[matchedNew+1:]...)
		}
	}
	reorderedNew = append(reorderedNew, nw...)

	return reorderedNew
}

func canonicalizeNewPrivateCloudConditionsSlice(c *Client, des, nw []PrivateCloudConditions) []PrivateCloudConditions {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []PrivateCloudConditions
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewPrivateCloudConditions(c, &d, &n))
	}

	return items
}

func canonicalizePrivateCloudHcx(des, initial *PrivateCloudHcx, opts ...dcl.ApplyOption) *PrivateCloudHcx {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	cDes := &PrivateCloudHcx{}

	if dcl.StringCanonicalize(des.Fdqn, initial.Fdqn) || dcl.IsZeroValue(des.Fdqn) {
		cDes.Fdqn = initial.Fdqn
	} else {
		cDes.Fdqn = des.Fdqn
	}
	if dcl.StringCanonicalize(des.InternalIP, initial.InternalIP) || dcl.IsZeroValue(des.InternalIP) {
		cDes.InternalIP = initial.InternalIP
	} else {
		cDes.InternalIP = des.InternalIP
	}
	if dcl.StringCanonicalize(des.ExternalIP, initial.ExternalIP) || dcl.IsZeroValue(des.ExternalIP) {
		cDes.ExternalIP = initial.ExternalIP
	} else {
		cDes.ExternalIP = des.ExternalIP
	}
	if dcl.StringCanonicalize(des.Version, initial.Version) || dcl.IsZeroValue(des.Version) {
		cDes.Version = initial.Version
	} else {
		cDes.Version = des.Version
	}

	return cDes
}

func canonicalizePrivateCloudHcxSlice(des, initial []PrivateCloudHcx, opts ...dcl.ApplyOption) []PrivateCloudHcx {
	if dcl.IsEmptyValueIndirect(des) {
		return initial
	}

	if len(des) != len(initial) {

		items := make([]PrivateCloudHcx, 0, len(des))
		for _, d := range des {
			cd := canonicalizePrivateCloudHcx(&d, nil, opts...)
			if cd != nil {
				items = append(items, *cd)
			}
		}
		return items
	}

	items := make([]PrivateCloudHcx, 0, len(des))
	for i, d := range des {
		cd := canonicalizePrivateCloudHcx(&d, &initial[i], opts...)
		if cd != nil {
			items = append(items, *cd)
		}
	}
	return items

}

func canonicalizeNewPrivateCloudHcx(c *Client, des, nw *PrivateCloudHcx) *PrivateCloudHcx {

	if des == nil {
		return nw
	}

	if nw == nil {
		if dcl.IsNotReturnedByServer(des) {
			c.Config.Logger.Info("Found explicitly empty value for PrivateCloudHcx while comparing non-nil desired to nil actual.  Returning desired object.")
			return des
		}
		return nil
	}

	if dcl.StringCanonicalize(des.Fdqn, nw.Fdqn) {
		nw.Fdqn = des.Fdqn
	}
	if dcl.StringCanonicalize(des.InternalIP, nw.InternalIP) {
		nw.InternalIP = des.InternalIP
	}
	if dcl.StringCanonicalize(des.ExternalIP, nw.ExternalIP) {
		nw.ExternalIP = des.ExternalIP
	}
	if dcl.StringCanonicalize(des.Version, nw.Version) {
		nw.Version = des.Version
	}

	return nw
}

func canonicalizeNewPrivateCloudHcxSet(c *Client, des, nw []PrivateCloudHcx) []PrivateCloudHcx {
	if des == nil {
		return nw
	}
	var reorderedNew []PrivateCloudHcx
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if diffs, _ := comparePrivateCloudHcxNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
				matchedNew = idx
				break
			}
		}
		if matchedNew != -1 {
			reorderedNew = append(reorderedNew, nw[matchedNew])
			nw = append(nw[:matchedNew], nw[matchedNew+1:]...)
		}
	}
	reorderedNew = append(reorderedNew, nw...)

	return reorderedNew
}

func canonicalizeNewPrivateCloudHcxSlice(c *Client, des, nw []PrivateCloudHcx) []PrivateCloudHcx {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []PrivateCloudHcx
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewPrivateCloudHcx(c, &d, &n))
	}

	return items
}

func canonicalizePrivateCloudNsx(des, initial *PrivateCloudNsx, opts ...dcl.ApplyOption) *PrivateCloudNsx {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	cDes := &PrivateCloudNsx{}

	if dcl.StringCanonicalize(des.Fdqn, initial.Fdqn) || dcl.IsZeroValue(des.Fdqn) {
		cDes.Fdqn = initial.Fdqn
	} else {
		cDes.Fdqn = des.Fdqn
	}
	if dcl.StringCanonicalize(des.InternalIP, initial.InternalIP) || dcl.IsZeroValue(des.InternalIP) {
		cDes.InternalIP = initial.InternalIP
	} else {
		cDes.InternalIP = des.InternalIP
	}
	if dcl.StringCanonicalize(des.ExternalIP, initial.ExternalIP) || dcl.IsZeroValue(des.ExternalIP) {
		cDes.ExternalIP = initial.ExternalIP
	} else {
		cDes.ExternalIP = des.ExternalIP
	}
	if dcl.StringCanonicalize(des.Version, initial.Version) || dcl.IsZeroValue(des.Version) {
		cDes.Version = initial.Version
	} else {
		cDes.Version = des.Version
	}

	return cDes
}

func canonicalizePrivateCloudNsxSlice(des, initial []PrivateCloudNsx, opts ...dcl.ApplyOption) []PrivateCloudNsx {
	if dcl.IsEmptyValueIndirect(des) {
		return initial
	}

	if len(des) != len(initial) {

		items := make([]PrivateCloudNsx, 0, len(des))
		for _, d := range des {
			cd := canonicalizePrivateCloudNsx(&d, nil, opts...)
			if cd != nil {
				items = append(items, *cd)
			}
		}
		return items
	}

	items := make([]PrivateCloudNsx, 0, len(des))
	for i, d := range des {
		cd := canonicalizePrivateCloudNsx(&d, &initial[i], opts...)
		if cd != nil {
			items = append(items, *cd)
		}
	}
	return items

}

func canonicalizeNewPrivateCloudNsx(c *Client, des, nw *PrivateCloudNsx) *PrivateCloudNsx {

	if des == nil {
		return nw
	}

	if nw == nil {
		if dcl.IsNotReturnedByServer(des) {
			c.Config.Logger.Info("Found explicitly empty value for PrivateCloudNsx while comparing non-nil desired to nil actual.  Returning desired object.")
			return des
		}
		return nil
	}

	if dcl.StringCanonicalize(des.Fdqn, nw.Fdqn) {
		nw.Fdqn = des.Fdqn
	}
	if dcl.StringCanonicalize(des.InternalIP, nw.InternalIP) {
		nw.InternalIP = des.InternalIP
	}
	if dcl.StringCanonicalize(des.ExternalIP, nw.ExternalIP) {
		nw.ExternalIP = des.ExternalIP
	}
	if dcl.StringCanonicalize(des.Version, nw.Version) {
		nw.Version = des.Version
	}

	return nw
}

func canonicalizeNewPrivateCloudNsxSet(c *Client, des, nw []PrivateCloudNsx) []PrivateCloudNsx {
	if des == nil {
		return nw
	}
	var reorderedNew []PrivateCloudNsx
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if diffs, _ := comparePrivateCloudNsxNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
				matchedNew = idx
				break
			}
		}
		if matchedNew != -1 {
			reorderedNew = append(reorderedNew, nw[matchedNew])
			nw = append(nw[:matchedNew], nw[matchedNew+1:]...)
		}
	}
	reorderedNew = append(reorderedNew, nw...)

	return reorderedNew
}

func canonicalizeNewPrivateCloudNsxSlice(c *Client, des, nw []PrivateCloudNsx) []PrivateCloudNsx {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []PrivateCloudNsx
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewPrivateCloudNsx(c, &d, &n))
	}

	return items
}

func canonicalizePrivateCloudVcenter(des, initial *PrivateCloudVcenter, opts ...dcl.ApplyOption) *PrivateCloudVcenter {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	cDes := &PrivateCloudVcenter{}

	if dcl.StringCanonicalize(des.Fdqn, initial.Fdqn) || dcl.IsZeroValue(des.Fdqn) {
		cDes.Fdqn = initial.Fdqn
	} else {
		cDes.Fdqn = des.Fdqn
	}
	if dcl.StringCanonicalize(des.InternalIP, initial.InternalIP) || dcl.IsZeroValue(des.InternalIP) {
		cDes.InternalIP = initial.InternalIP
	} else {
		cDes.InternalIP = des.InternalIP
	}
	if dcl.StringCanonicalize(des.ExternalIP, initial.ExternalIP) || dcl.IsZeroValue(des.ExternalIP) {
		cDes.ExternalIP = initial.ExternalIP
	} else {
		cDes.ExternalIP = des.ExternalIP
	}
	if dcl.StringCanonicalize(des.Version, initial.Version) || dcl.IsZeroValue(des.Version) {
		cDes.Version = initial.Version
	} else {
		cDes.Version = des.Version
	}

	return cDes
}

func canonicalizePrivateCloudVcenterSlice(des, initial []PrivateCloudVcenter, opts ...dcl.ApplyOption) []PrivateCloudVcenter {
	if dcl.IsEmptyValueIndirect(des) {
		return initial
	}

	if len(des) != len(initial) {

		items := make([]PrivateCloudVcenter, 0, len(des))
		for _, d := range des {
			cd := canonicalizePrivateCloudVcenter(&d, nil, opts...)
			if cd != nil {
				items = append(items, *cd)
			}
		}
		return items
	}

	items := make([]PrivateCloudVcenter, 0, len(des))
	for i, d := range des {
		cd := canonicalizePrivateCloudVcenter(&d, &initial[i], opts...)
		if cd != nil {
			items = append(items, *cd)
		}
	}
	return items

}

func canonicalizeNewPrivateCloudVcenter(c *Client, des, nw *PrivateCloudVcenter) *PrivateCloudVcenter {

	if des == nil {
		return nw
	}

	if nw == nil {
		if dcl.IsNotReturnedByServer(des) {
			c.Config.Logger.Info("Found explicitly empty value for PrivateCloudVcenter while comparing non-nil desired to nil actual.  Returning desired object.")
			return des
		}
		return nil
	}

	if dcl.StringCanonicalize(des.Fdqn, nw.Fdqn) {
		nw.Fdqn = des.Fdqn
	}
	if dcl.StringCanonicalize(des.InternalIP, nw.InternalIP) {
		nw.InternalIP = des.InternalIP
	}
	if dcl.StringCanonicalize(des.ExternalIP, nw.ExternalIP) {
		nw.ExternalIP = des.ExternalIP
	}
	if dcl.StringCanonicalize(des.Version, nw.Version) {
		nw.Version = des.Version
	}

	return nw
}

func canonicalizeNewPrivateCloudVcenterSet(c *Client, des, nw []PrivateCloudVcenter) []PrivateCloudVcenter {
	if des == nil {
		return nw
	}
	var reorderedNew []PrivateCloudVcenter
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if diffs, _ := comparePrivateCloudVcenterNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
				matchedNew = idx
				break
			}
		}
		if matchedNew != -1 {
			reorderedNew = append(reorderedNew, nw[matchedNew])
			nw = append(nw[:matchedNew], nw[matchedNew+1:]...)
		}
	}
	reorderedNew = append(reorderedNew, nw...)

	return reorderedNew
}

func canonicalizeNewPrivateCloudVcenterSlice(c *Client, des, nw []PrivateCloudVcenter) []PrivateCloudVcenter {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []PrivateCloudVcenter
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewPrivateCloudVcenter(c, &d, &n))
	}

	return items
}

// The differ returns a list of diffs, along with a list of operations that should be taken
// to remedy them. Right now, it does not attempt to consolidate operations - if several
// fields can be fixed with a patch update, it will perform the patch several times.
// Diffs on some fields will be ignored if the `desired` state has an empty (nil)
// value. This empty value indicates that the user does not care about the state for
// the field. Empty fields on the actual object will cause diffs.
// TODO(magic-modules-eng): for efficiency in some resources, add batching.
func diffPrivateCloud(c *Client, desired, actual *PrivateCloud, opts ...dcl.ApplyOption) ([]*dcl.FieldDiff, error) {
	if desired == nil || actual == nil {
		return nil, fmt.Errorf("nil resource passed to diff - always a programming error: %#v, %#v", desired, actual)
	}

	c.Config.Logger.Infof("Diff function called with desired state: %v", desired)
	c.Config.Logger.Infof("Diff function called with actual state: %v", actual)

	var fn dcl.FieldName
	var newDiffs []*dcl.FieldDiff
	// New style diffs.
	if ds, err := dcl.Diff(desired.Name, actual.Name, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("Name")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.CreateTime, actual.CreateTime, dcl.Info{OutputOnly: true, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("CreateTime")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.UpdateTime, actual.UpdateTime, dcl.Info{OutputOnly: true, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("UpdateTime")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.DeleteTime, actual.DeleteTime, dcl.Info{OutputOnly: true, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("DeleteTime")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.ExpireTime, actual.ExpireTime, dcl.Info{OutputOnly: true, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("ExpireTime")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.State, actual.State, dcl.Info{OutputOnly: true, Type: "EnumType", OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("State")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.NetworkConfig, actual.NetworkConfig, dcl.Info{ObjectFunction: comparePrivateCloudNetworkConfigNewStyle, EmptyObject: EmptyPrivateCloudNetworkConfig, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("NetworkConfig")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.ManagementCluster, actual.ManagementCluster, dcl.Info{ObjectFunction: comparePrivateCloudManagementClusterNewStyle, EmptyObject: EmptyPrivateCloudManagementCluster, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("ManagementCluster")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Description, actual.Description, dcl.Info{OperationSelector: dcl.TriggersOperation("updatePrivateCloudUpdatePrivateCloudOperation")}, fn.AddNest("Description")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Conditions, actual.Conditions, dcl.Info{OutputOnly: true, ObjectFunction: comparePrivateCloudConditionsNewStyle, EmptyObject: EmptyPrivateCloudConditions, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("Conditions")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Hcx, actual.Hcx, dcl.Info{OutputOnly: true, ObjectFunction: comparePrivateCloudHcxNewStyle, EmptyObject: EmptyPrivateCloudHcx, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("Hcx")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Nsx, actual.Nsx, dcl.Info{OutputOnly: true, ObjectFunction: comparePrivateCloudNsxNewStyle, EmptyObject: EmptyPrivateCloudNsx, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("Nsx")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Vcenter, actual.Vcenter, dcl.Info{OutputOnly: true, ObjectFunction: comparePrivateCloudVcenterNewStyle, EmptyObject: EmptyPrivateCloudVcenter, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("Vcenter")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Project, actual.Project, dcl.Info{Type: "ReferenceType", OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("Project")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Location, actual.Location, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("Location")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	return newDiffs, nil
}
func comparePrivateCloudNetworkConfigNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*PrivateCloudNetworkConfig)
	if !ok {
		desiredNotPointer, ok := d.(PrivateCloudNetworkConfig)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a PrivateCloudNetworkConfig or *PrivateCloudNetworkConfig", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*PrivateCloudNetworkConfig)
	if !ok {
		actualNotPointer, ok := a.(PrivateCloudNetworkConfig)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a PrivateCloudNetworkConfig", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.Network, actual.Network, dcl.Info{Type: "ReferenceType", CustomDiff: canonicalizePrivateCloudNetwork, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("Network")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.ServiceNetwork, actual.ServiceNetwork, dcl.Info{OutputOnly: true, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("ServiceNetwork")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.ManagementCidr, actual.ManagementCidr, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("ManagementCidr")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func comparePrivateCloudManagementClusterNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*PrivateCloudManagementCluster)
	if !ok {
		desiredNotPointer, ok := d.(PrivateCloudManagementCluster)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a PrivateCloudManagementCluster or *PrivateCloudManagementCluster", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*PrivateCloudManagementCluster)
	if !ok {
		actualNotPointer, ok := a.(PrivateCloudManagementCluster)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a PrivateCloudManagementCluster", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.ClusterId, actual.ClusterId, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("ClusterId")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.NodeTypeId, actual.NodeTypeId, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("NodeTypeId")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.NodeCount, actual.NodeCount, dcl.Info{OperationSelector: dcl.TriggersOperation("updatePrivateCloudUpdateClusterOperation")}, fn.AddNest("NodeCount")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func comparePrivateCloudConditionsNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*PrivateCloudConditions)
	if !ok {
		desiredNotPointer, ok := d.(PrivateCloudConditions)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a PrivateCloudConditions or *PrivateCloudConditions", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*PrivateCloudConditions)
	if !ok {
		actualNotPointer, ok := a.(PrivateCloudConditions)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a PrivateCloudConditions", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.Code, actual.Code, dcl.Info{OutputOnly: true, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("Code")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Message, actual.Message, dcl.Info{OutputOnly: true, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("Message")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func comparePrivateCloudHcxNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*PrivateCloudHcx)
	if !ok {
		desiredNotPointer, ok := d.(PrivateCloudHcx)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a PrivateCloudHcx or *PrivateCloudHcx", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*PrivateCloudHcx)
	if !ok {
		actualNotPointer, ok := a.(PrivateCloudHcx)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a PrivateCloudHcx", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.Fdqn, actual.Fdqn, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("Fdqn")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.InternalIP, actual.InternalIP, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("InternalIp")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.ExternalIP, actual.ExternalIP, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("ExternalIp")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Version, actual.Version, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("Version")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func comparePrivateCloudNsxNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*PrivateCloudNsx)
	if !ok {
		desiredNotPointer, ok := d.(PrivateCloudNsx)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a PrivateCloudNsx or *PrivateCloudNsx", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*PrivateCloudNsx)
	if !ok {
		actualNotPointer, ok := a.(PrivateCloudNsx)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a PrivateCloudNsx", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.Fdqn, actual.Fdqn, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("Fdqn")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.InternalIP, actual.InternalIP, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("InternalIp")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.ExternalIP, actual.ExternalIP, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("ExternalIp")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Version, actual.Version, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("Version")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func comparePrivateCloudVcenterNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*PrivateCloudVcenter)
	if !ok {
		desiredNotPointer, ok := d.(PrivateCloudVcenter)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a PrivateCloudVcenter or *PrivateCloudVcenter", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*PrivateCloudVcenter)
	if !ok {
		actualNotPointer, ok := a.(PrivateCloudVcenter)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a PrivateCloudVcenter", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.Fdqn, actual.Fdqn, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("Fdqn")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.InternalIP, actual.InternalIP, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("InternalIp")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.ExternalIP, actual.ExternalIP, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("ExternalIp")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Version, actual.Version, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("Version")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

// urlNormalized returns a copy of the resource struct with values normalized
// for URL substitutions. For instance, it converts long-form self-links to
// short-form so they can be substituted in.
func (r *PrivateCloud) urlNormalized() *PrivateCloud {
	normalized := dcl.Copy(*r).(PrivateCloud)
	normalized.Name = dcl.SelfLinkToName(r.Name)
	normalized.Description = dcl.SelfLinkToName(r.Description)
	normalized.Project = dcl.SelfLinkToName(r.Project)
	normalized.Location = dcl.SelfLinkToName(r.Location)
	return &normalized
}

func (r *PrivateCloud) updateURL(userBasePath, updateName string) (string, error) {
	nr := r.urlNormalized()
	if updateName == "UpdateCluster" {
		fields := map[string]interface{}{
			"project":                     dcl.ValueOrEmptyString(nr.Project),
			"location":                    dcl.ValueOrEmptyString(nr.Location),
			"managementCluster.clusterId": dcl.ValueOrEmptyString(nr.ManagementCluster.ClusterId),
			"name":                        dcl.ValueOrEmptyString(nr.Name),
		}
		return dcl.URL("projects/{{project}}/locations/{{location}}/privateClouds/{{name}}/clusters/{{managementCluster.clusterId}}", nr.basePath(), userBasePath, fields), nil

	}
	if updateName == "UpdatePrivateCloud" {
		fields := map[string]interface{}{
			"project":  dcl.ValueOrEmptyString(nr.Project),
			"location": dcl.ValueOrEmptyString(nr.Location),
			"name":     dcl.ValueOrEmptyString(nr.Name),
		}
		return dcl.URL("projects/{{project}}/locations/{{location}}/privateClouds/{{name}}", nr.basePath(), userBasePath, fields), nil

	}

	return "", fmt.Errorf("unknown update name: %s", updateName)
}

// marshal encodes the PrivateCloud resource into JSON for a Create request, and
// performs transformations from the resource schema to the API schema if
// necessary.
func (r *PrivateCloud) marshal(c *Client) ([]byte, error) {
	m, err := expandPrivateCloud(c, r)
	if err != nil {
		return nil, fmt.Errorf("error marshalling PrivateCloud: %w", err)
	}

	return json.Marshal(m)
}

// unmarshalPrivateCloud decodes JSON responses into the PrivateCloud resource schema.
func unmarshalPrivateCloud(b []byte, c *Client, res *PrivateCloud) (*PrivateCloud, error) {
	var m map[string]interface{}
	if err := json.Unmarshal(b, &m); err != nil {
		return nil, err
	}
	return unmarshalMapPrivateCloud(m, c, res)
}

func unmarshalMapPrivateCloud(m map[string]interface{}, c *Client, res *PrivateCloud) (*PrivateCloud, error) {

	flattened := flattenPrivateCloud(c, m, res)
	if flattened == nil {
		return nil, fmt.Errorf("attempted to flatten empty json object")
	}
	return flattened, nil
}

// expandPrivateCloud expands PrivateCloud into a JSON request object.
func expandPrivateCloud(c *Client, f *PrivateCloud) (map[string]interface{}, error) {
	m := make(map[string]interface{})
	res := f
	_ = res
	if v, err := dcl.DeriveField("projects/%s/locations/%s/privateClouds/%s", f.Name, dcl.SelfLinkToName(f.Project), dcl.SelfLinkToName(f.Location), dcl.SelfLinkToName(f.Name)); err != nil {
		return nil, fmt.Errorf("error expanding Name into name: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["name"] = v
	}
	if v, err := expandPrivateCloudNetworkConfig(c, f.NetworkConfig, res); err != nil {
		return nil, fmt.Errorf("error expanding NetworkConfig into networkConfig: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["networkConfig"] = v
	}
	if v, err := expandPrivateCloudManagementCluster(c, f.ManagementCluster, res); err != nil {
		return nil, fmt.Errorf("error expanding ManagementCluster into managementCluster: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["managementCluster"] = v
	}
	if v := f.Description; dcl.ValueShouldBeSent(v) {
		m["description"] = v
	}
	if v, err := dcl.EmptyValue(); err != nil {
		return nil, fmt.Errorf("error expanding Project into project: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["project"] = v
	}
	if v, err := dcl.EmptyValue(); err != nil {
		return nil, fmt.Errorf("error expanding Location into location: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["location"] = v
	}

	return m, nil
}

// flattenPrivateCloud flattens PrivateCloud from a JSON request object into the
// PrivateCloud type.
func flattenPrivateCloud(c *Client, i interface{}, res *PrivateCloud) *PrivateCloud {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}
	if len(m) == 0 {
		return nil
	}

	resultRes := &PrivateCloud{}
	resultRes.Name = dcl.FlattenString(m["name"])
	resultRes.CreateTime = dcl.FlattenString(m["createTime"])
	resultRes.UpdateTime = dcl.FlattenString(m["updateTime"])
	resultRes.DeleteTime = dcl.FlattenString(m["deleteTime"])
	resultRes.ExpireTime = dcl.FlattenString(m["expireTime"])
	resultRes.State = flattenPrivateCloudStateEnum(m["state"])
	resultRes.NetworkConfig = flattenPrivateCloudNetworkConfig(c, m["networkConfig"], res)
	resultRes.ManagementCluster = flattenPrivateCloudManagementCluster(c, m["managementCluster"], res)
	resultRes.Description = dcl.FlattenString(m["description"])
	resultRes.Conditions = flattenPrivateCloudConditionsSlice(c, m["conditions"], res)
	resultRes.Hcx = flattenPrivateCloudHcx(c, m["hcx"], res)
	resultRes.Nsx = flattenPrivateCloudNsx(c, m["nsx"], res)
	resultRes.Vcenter = flattenPrivateCloudVcenter(c, m["vcenter"], res)
	resultRes.Project = dcl.FlattenString(m["project"])
	resultRes.Location = dcl.FlattenString(m["location"])

	return resultRes
}

// expandPrivateCloudNetworkConfigMap expands the contents of PrivateCloudNetworkConfig into a JSON
// request object.
func expandPrivateCloudNetworkConfigMap(c *Client, f map[string]PrivateCloudNetworkConfig, res *PrivateCloud) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandPrivateCloudNetworkConfig(c, &item, res)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandPrivateCloudNetworkConfigSlice expands the contents of PrivateCloudNetworkConfig into a JSON
// request object.
func expandPrivateCloudNetworkConfigSlice(c *Client, f []PrivateCloudNetworkConfig, res *PrivateCloud) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandPrivateCloudNetworkConfig(c, &item, res)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenPrivateCloudNetworkConfigMap flattens the contents of PrivateCloudNetworkConfig from a JSON
// response object.
func flattenPrivateCloudNetworkConfigMap(c *Client, i interface{}, res *PrivateCloud) map[string]PrivateCloudNetworkConfig {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]PrivateCloudNetworkConfig{}
	}

	if len(a) == 0 {
		return map[string]PrivateCloudNetworkConfig{}
	}

	items := make(map[string]PrivateCloudNetworkConfig)
	for k, item := range a {
		items[k] = *flattenPrivateCloudNetworkConfig(c, item.(map[string]interface{}), res)
	}

	return items
}

// flattenPrivateCloudNetworkConfigSlice flattens the contents of PrivateCloudNetworkConfig from a JSON
// response object.
func flattenPrivateCloudNetworkConfigSlice(c *Client, i interface{}, res *PrivateCloud) []PrivateCloudNetworkConfig {
	a, ok := i.([]interface{})
	if !ok {
		return []PrivateCloudNetworkConfig{}
	}

	if len(a) == 0 {
		return []PrivateCloudNetworkConfig{}
	}

	items := make([]PrivateCloudNetworkConfig, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenPrivateCloudNetworkConfig(c, item.(map[string]interface{}), res))
	}

	return items
}

// expandPrivateCloudNetworkConfig expands an instance of PrivateCloudNetworkConfig into a JSON
// request object.
func expandPrivateCloudNetworkConfig(c *Client, f *PrivateCloudNetworkConfig, res *PrivateCloud) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v := f.Network; !dcl.IsEmptyValueIndirect(v) {
		m["network"] = v
	}
	if v := f.ManagementCidr; !dcl.IsEmptyValueIndirect(v) {
		m["managementCidr"] = v
	}

	return m, nil
}

// flattenPrivateCloudNetworkConfig flattens an instance of PrivateCloudNetworkConfig from a JSON
// response object.
func flattenPrivateCloudNetworkConfig(c *Client, i interface{}, res *PrivateCloud) *PrivateCloudNetworkConfig {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &PrivateCloudNetworkConfig{}

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyPrivateCloudNetworkConfig
	}
	r.Network = dcl.FlattenString(m["network"])
	r.ServiceNetwork = dcl.FlattenString(m["serviceNetwork"])
	r.ManagementCidr = dcl.FlattenString(m["managementCidr"])

	return r
}

// expandPrivateCloudManagementClusterMap expands the contents of PrivateCloudManagementCluster into a JSON
// request object.
func expandPrivateCloudManagementClusterMap(c *Client, f map[string]PrivateCloudManagementCluster, res *PrivateCloud) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandPrivateCloudManagementCluster(c, &item, res)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandPrivateCloudManagementClusterSlice expands the contents of PrivateCloudManagementCluster into a JSON
// request object.
func expandPrivateCloudManagementClusterSlice(c *Client, f []PrivateCloudManagementCluster, res *PrivateCloud) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandPrivateCloudManagementCluster(c, &item, res)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenPrivateCloudManagementClusterMap flattens the contents of PrivateCloudManagementCluster from a JSON
// response object.
func flattenPrivateCloudManagementClusterMap(c *Client, i interface{}, res *PrivateCloud) map[string]PrivateCloudManagementCluster {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]PrivateCloudManagementCluster{}
	}

	if len(a) == 0 {
		return map[string]PrivateCloudManagementCluster{}
	}

	items := make(map[string]PrivateCloudManagementCluster)
	for k, item := range a {
		items[k] = *flattenPrivateCloudManagementCluster(c, item.(map[string]interface{}), res)
	}

	return items
}

// flattenPrivateCloudManagementClusterSlice flattens the contents of PrivateCloudManagementCluster from a JSON
// response object.
func flattenPrivateCloudManagementClusterSlice(c *Client, i interface{}, res *PrivateCloud) []PrivateCloudManagementCluster {
	a, ok := i.([]interface{})
	if !ok {
		return []PrivateCloudManagementCluster{}
	}

	if len(a) == 0 {
		return []PrivateCloudManagementCluster{}
	}

	items := make([]PrivateCloudManagementCluster, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenPrivateCloudManagementCluster(c, item.(map[string]interface{}), res))
	}

	return items
}

// expandPrivateCloudManagementCluster expands an instance of PrivateCloudManagementCluster into a JSON
// request object.
func expandPrivateCloudManagementCluster(c *Client, f *PrivateCloudManagementCluster, res *PrivateCloud) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v := f.ClusterId; !dcl.IsEmptyValueIndirect(v) {
		m["clusterId"] = v
	}
	if v := f.NodeTypeId; !dcl.IsEmptyValueIndirect(v) {
		m["nodeTypeId"] = v
	}
	if v := f.NodeCount; !dcl.IsEmptyValueIndirect(v) {
		m["nodeCount"] = v
	}

	return m, nil
}

// flattenPrivateCloudManagementCluster flattens an instance of PrivateCloudManagementCluster from a JSON
// response object.
func flattenPrivateCloudManagementCluster(c *Client, i interface{}, res *PrivateCloud) *PrivateCloudManagementCluster {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &PrivateCloudManagementCluster{}

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyPrivateCloudManagementCluster
	}
	r.ClusterId = dcl.FlattenString(m["clusterId"])
	r.NodeTypeId = dcl.FlattenString(m["nodeTypeId"])
	r.NodeCount = dcl.FlattenInteger(m["nodeCount"])

	return r
}

// expandPrivateCloudConditionsMap expands the contents of PrivateCloudConditions into a JSON
// request object.
func expandPrivateCloudConditionsMap(c *Client, f map[string]PrivateCloudConditions, res *PrivateCloud) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandPrivateCloudConditions(c, &item, res)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandPrivateCloudConditionsSlice expands the contents of PrivateCloudConditions into a JSON
// request object.
func expandPrivateCloudConditionsSlice(c *Client, f []PrivateCloudConditions, res *PrivateCloud) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandPrivateCloudConditions(c, &item, res)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenPrivateCloudConditionsMap flattens the contents of PrivateCloudConditions from a JSON
// response object.
func flattenPrivateCloudConditionsMap(c *Client, i interface{}, res *PrivateCloud) map[string]PrivateCloudConditions {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]PrivateCloudConditions{}
	}

	if len(a) == 0 {
		return map[string]PrivateCloudConditions{}
	}

	items := make(map[string]PrivateCloudConditions)
	for k, item := range a {
		items[k] = *flattenPrivateCloudConditions(c, item.(map[string]interface{}), res)
	}

	return items
}

// flattenPrivateCloudConditionsSlice flattens the contents of PrivateCloudConditions from a JSON
// response object.
func flattenPrivateCloudConditionsSlice(c *Client, i interface{}, res *PrivateCloud) []PrivateCloudConditions {
	a, ok := i.([]interface{})
	if !ok {
		return []PrivateCloudConditions{}
	}

	if len(a) == 0 {
		return []PrivateCloudConditions{}
	}

	items := make([]PrivateCloudConditions, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenPrivateCloudConditions(c, item.(map[string]interface{}), res))
	}

	return items
}

// expandPrivateCloudConditions expands an instance of PrivateCloudConditions into a JSON
// request object.
func expandPrivateCloudConditions(c *Client, f *PrivateCloudConditions, res *PrivateCloud) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})

	return m, nil
}

// flattenPrivateCloudConditions flattens an instance of PrivateCloudConditions from a JSON
// response object.
func flattenPrivateCloudConditions(c *Client, i interface{}, res *PrivateCloud) *PrivateCloudConditions {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &PrivateCloudConditions{}

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyPrivateCloudConditions
	}
	r.Code = dcl.FlattenString(m["code"])
	r.Message = dcl.FlattenString(m["message"])

	return r
}

// expandPrivateCloudHcxMap expands the contents of PrivateCloudHcx into a JSON
// request object.
func expandPrivateCloudHcxMap(c *Client, f map[string]PrivateCloudHcx, res *PrivateCloud) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandPrivateCloudHcx(c, &item, res)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandPrivateCloudHcxSlice expands the contents of PrivateCloudHcx into a JSON
// request object.
func expandPrivateCloudHcxSlice(c *Client, f []PrivateCloudHcx, res *PrivateCloud) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandPrivateCloudHcx(c, &item, res)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenPrivateCloudHcxMap flattens the contents of PrivateCloudHcx from a JSON
// response object.
func flattenPrivateCloudHcxMap(c *Client, i interface{}, res *PrivateCloud) map[string]PrivateCloudHcx {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]PrivateCloudHcx{}
	}

	if len(a) == 0 {
		return map[string]PrivateCloudHcx{}
	}

	items := make(map[string]PrivateCloudHcx)
	for k, item := range a {
		items[k] = *flattenPrivateCloudHcx(c, item.(map[string]interface{}), res)
	}

	return items
}

// flattenPrivateCloudHcxSlice flattens the contents of PrivateCloudHcx from a JSON
// response object.
func flattenPrivateCloudHcxSlice(c *Client, i interface{}, res *PrivateCloud) []PrivateCloudHcx {
	a, ok := i.([]interface{})
	if !ok {
		return []PrivateCloudHcx{}
	}

	if len(a) == 0 {
		return []PrivateCloudHcx{}
	}

	items := make([]PrivateCloudHcx, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenPrivateCloudHcx(c, item.(map[string]interface{}), res))
	}

	return items
}

// expandPrivateCloudHcx expands an instance of PrivateCloudHcx into a JSON
// request object.
func expandPrivateCloudHcx(c *Client, f *PrivateCloudHcx, res *PrivateCloud) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v := f.Fdqn; !dcl.IsEmptyValueIndirect(v) {
		m["fdqn"] = v
	}
	if v := f.InternalIP; !dcl.IsEmptyValueIndirect(v) {
		m["internalIp"] = v
	}
	if v := f.ExternalIP; !dcl.IsEmptyValueIndirect(v) {
		m["externalIp"] = v
	}
	if v := f.Version; !dcl.IsEmptyValueIndirect(v) {
		m["version"] = v
	}

	return m, nil
}

// flattenPrivateCloudHcx flattens an instance of PrivateCloudHcx from a JSON
// response object.
func flattenPrivateCloudHcx(c *Client, i interface{}, res *PrivateCloud) *PrivateCloudHcx {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &PrivateCloudHcx{}

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyPrivateCloudHcx
	}
	r.Fdqn = dcl.FlattenString(m["fdqn"])
	r.InternalIP = dcl.FlattenString(m["internalIp"])
	r.ExternalIP = dcl.FlattenString(m["externalIp"])
	r.Version = dcl.FlattenString(m["version"])

	return r
}

// expandPrivateCloudNsxMap expands the contents of PrivateCloudNsx into a JSON
// request object.
func expandPrivateCloudNsxMap(c *Client, f map[string]PrivateCloudNsx, res *PrivateCloud) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandPrivateCloudNsx(c, &item, res)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandPrivateCloudNsxSlice expands the contents of PrivateCloudNsx into a JSON
// request object.
func expandPrivateCloudNsxSlice(c *Client, f []PrivateCloudNsx, res *PrivateCloud) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandPrivateCloudNsx(c, &item, res)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenPrivateCloudNsxMap flattens the contents of PrivateCloudNsx from a JSON
// response object.
func flattenPrivateCloudNsxMap(c *Client, i interface{}, res *PrivateCloud) map[string]PrivateCloudNsx {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]PrivateCloudNsx{}
	}

	if len(a) == 0 {
		return map[string]PrivateCloudNsx{}
	}

	items := make(map[string]PrivateCloudNsx)
	for k, item := range a {
		items[k] = *flattenPrivateCloudNsx(c, item.(map[string]interface{}), res)
	}

	return items
}

// flattenPrivateCloudNsxSlice flattens the contents of PrivateCloudNsx from a JSON
// response object.
func flattenPrivateCloudNsxSlice(c *Client, i interface{}, res *PrivateCloud) []PrivateCloudNsx {
	a, ok := i.([]interface{})
	if !ok {
		return []PrivateCloudNsx{}
	}

	if len(a) == 0 {
		return []PrivateCloudNsx{}
	}

	items := make([]PrivateCloudNsx, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenPrivateCloudNsx(c, item.(map[string]interface{}), res))
	}

	return items
}

// expandPrivateCloudNsx expands an instance of PrivateCloudNsx into a JSON
// request object.
func expandPrivateCloudNsx(c *Client, f *PrivateCloudNsx, res *PrivateCloud) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v := f.Fdqn; !dcl.IsEmptyValueIndirect(v) {
		m["fdqn"] = v
	}
	if v := f.InternalIP; !dcl.IsEmptyValueIndirect(v) {
		m["internalIp"] = v
	}
	if v := f.ExternalIP; !dcl.IsEmptyValueIndirect(v) {
		m["externalIp"] = v
	}
	if v := f.Version; !dcl.IsEmptyValueIndirect(v) {
		m["version"] = v
	}

	return m, nil
}

// flattenPrivateCloudNsx flattens an instance of PrivateCloudNsx from a JSON
// response object.
func flattenPrivateCloudNsx(c *Client, i interface{}, res *PrivateCloud) *PrivateCloudNsx {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &PrivateCloudNsx{}

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyPrivateCloudNsx
	}
	r.Fdqn = dcl.FlattenString(m["fdqn"])
	r.InternalIP = dcl.FlattenString(m["internalIp"])
	r.ExternalIP = dcl.FlattenString(m["externalIp"])
	r.Version = dcl.FlattenString(m["version"])

	return r
}

// expandPrivateCloudVcenterMap expands the contents of PrivateCloudVcenter into a JSON
// request object.
func expandPrivateCloudVcenterMap(c *Client, f map[string]PrivateCloudVcenter, res *PrivateCloud) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandPrivateCloudVcenter(c, &item, res)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandPrivateCloudVcenterSlice expands the contents of PrivateCloudVcenter into a JSON
// request object.
func expandPrivateCloudVcenterSlice(c *Client, f []PrivateCloudVcenter, res *PrivateCloud) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandPrivateCloudVcenter(c, &item, res)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenPrivateCloudVcenterMap flattens the contents of PrivateCloudVcenter from a JSON
// response object.
func flattenPrivateCloudVcenterMap(c *Client, i interface{}, res *PrivateCloud) map[string]PrivateCloudVcenter {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]PrivateCloudVcenter{}
	}

	if len(a) == 0 {
		return map[string]PrivateCloudVcenter{}
	}

	items := make(map[string]PrivateCloudVcenter)
	for k, item := range a {
		items[k] = *flattenPrivateCloudVcenter(c, item.(map[string]interface{}), res)
	}

	return items
}

// flattenPrivateCloudVcenterSlice flattens the contents of PrivateCloudVcenter from a JSON
// response object.
func flattenPrivateCloudVcenterSlice(c *Client, i interface{}, res *PrivateCloud) []PrivateCloudVcenter {
	a, ok := i.([]interface{})
	if !ok {
		return []PrivateCloudVcenter{}
	}

	if len(a) == 0 {
		return []PrivateCloudVcenter{}
	}

	items := make([]PrivateCloudVcenter, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenPrivateCloudVcenter(c, item.(map[string]interface{}), res))
	}

	return items
}

// expandPrivateCloudVcenter expands an instance of PrivateCloudVcenter into a JSON
// request object.
func expandPrivateCloudVcenter(c *Client, f *PrivateCloudVcenter, res *PrivateCloud) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v := f.Fdqn; !dcl.IsEmptyValueIndirect(v) {
		m["fdqn"] = v
	}
	if v := f.InternalIP; !dcl.IsEmptyValueIndirect(v) {
		m["internalIp"] = v
	}
	if v := f.ExternalIP; !dcl.IsEmptyValueIndirect(v) {
		m["externalIp"] = v
	}
	if v := f.Version; !dcl.IsEmptyValueIndirect(v) {
		m["version"] = v
	}

	return m, nil
}

// flattenPrivateCloudVcenter flattens an instance of PrivateCloudVcenter from a JSON
// response object.
func flattenPrivateCloudVcenter(c *Client, i interface{}, res *PrivateCloud) *PrivateCloudVcenter {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &PrivateCloudVcenter{}

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyPrivateCloudVcenter
	}
	r.Fdqn = dcl.FlattenString(m["fdqn"])
	r.InternalIP = dcl.FlattenString(m["internalIp"])
	r.ExternalIP = dcl.FlattenString(m["externalIp"])
	r.Version = dcl.FlattenString(m["version"])

	return r
}

// flattenPrivateCloudStateEnumMap flattens the contents of PrivateCloudStateEnum from a JSON
// response object.
func flattenPrivateCloudStateEnumMap(c *Client, i interface{}, res *PrivateCloud) map[string]PrivateCloudStateEnum {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]PrivateCloudStateEnum{}
	}

	if len(a) == 0 {
		return map[string]PrivateCloudStateEnum{}
	}

	items := make(map[string]PrivateCloudStateEnum)
	for k, item := range a {
		items[k] = *flattenPrivateCloudStateEnum(item.(interface{}))
	}

	return items
}

// flattenPrivateCloudStateEnumSlice flattens the contents of PrivateCloudStateEnum from a JSON
// response object.
func flattenPrivateCloudStateEnumSlice(c *Client, i interface{}, res *PrivateCloud) []PrivateCloudStateEnum {
	a, ok := i.([]interface{})
	if !ok {
		return []PrivateCloudStateEnum{}
	}

	if len(a) == 0 {
		return []PrivateCloudStateEnum{}
	}

	items := make([]PrivateCloudStateEnum, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenPrivateCloudStateEnum(item.(interface{})))
	}

	return items
}

// flattenPrivateCloudStateEnum asserts that an interface is a string, and returns a
// pointer to a *PrivateCloudStateEnum with the same value as that string.
func flattenPrivateCloudStateEnum(i interface{}) *PrivateCloudStateEnum {
	s, ok := i.(string)
	if !ok {
		return nil
	}

	return PrivateCloudStateEnumRef(s)
}

// This function returns a matcher that checks whether a serialized resource matches this resource
// in its parameters (as defined by the fields in a Get, which definitionally define resource
// identity).  This is useful in extracting the element from a List call.
func (r *PrivateCloud) matcher(c *Client) func([]byte) bool {
	return func(b []byte) bool {
		cr, err := unmarshalPrivateCloud(b, c, r)
		if err != nil {
			c.Config.Logger.Warning("failed to unmarshal provided resource in matcher.")
			return false
		}
		nr := r.urlNormalized()
		ncr := cr.urlNormalized()
		c.Config.Logger.Infof("looking for %v\nin %v", nr, ncr)

		if nr.Project == nil && ncr.Project == nil {
			c.Config.Logger.Info("Both Project fields null - considering equal.")
		} else if nr.Project == nil || ncr.Project == nil {
			c.Config.Logger.Info("Only one Project field is null - considering unequal.")
			return false
		} else if *nr.Project != *ncr.Project {
			return false
		}
		if nr.Location == nil && ncr.Location == nil {
			c.Config.Logger.Info("Both Location fields null - considering equal.")
		} else if nr.Location == nil || ncr.Location == nil {
			c.Config.Logger.Info("Only one Location field is null - considering unequal.")
			return false
		} else if *nr.Location != *ncr.Location {
			return false
		}
		if nr.Name == nil && ncr.Name == nil {
			c.Config.Logger.Info("Both Name fields null - considering equal.")
		} else if nr.Name == nil || ncr.Name == nil {
			c.Config.Logger.Info("Only one Name field is null - considering unequal.")
			return false
		} else if *nr.Name != *ncr.Name {
			return false
		}
		return true
	}
}

type privateCloudDiff struct {
	// The diff should include one or the other of RequiresRecreate or UpdateOp.
	RequiresRecreate bool
	UpdateOp         privateCloudApiOperation
}

func convertFieldDiffsToPrivateCloudDiffs(config *dcl.Config, fds []*dcl.FieldDiff, opts []dcl.ApplyOption) ([]privateCloudDiff, error) {
	opNamesToFieldDiffs := make(map[string][]*dcl.FieldDiff)
	// Map each operation name to the field diffs associated with it.
	for _, fd := range fds {
		for _, ro := range fd.ResultingOperation {
			if fieldDiffs, ok := opNamesToFieldDiffs[ro]; ok {
				fieldDiffs = append(fieldDiffs, fd)
				opNamesToFieldDiffs[ro] = fieldDiffs
			} else {
				config.Logger.Infof("%s required due to diff: %v", ro, fd)
				opNamesToFieldDiffs[ro] = []*dcl.FieldDiff{fd}
			}
		}
	}
	var diffs []privateCloudDiff
	// For each operation name, create a privateCloudDiff which contains the operation.
	for opName, fieldDiffs := range opNamesToFieldDiffs {
		diff := privateCloudDiff{}
		if opName == "Recreate" {
			diff.RequiresRecreate = true
		} else {
			apiOp, err := convertOpNameToPrivateCloudApiOperation(opName, fieldDiffs, opts...)
			if err != nil {
				return diffs, err
			}
			diff.UpdateOp = apiOp
		}
		diffs = append(diffs, diff)
	}
	return diffs, nil
}

func convertOpNameToPrivateCloudApiOperation(opName string, fieldDiffs []*dcl.FieldDiff, opts ...dcl.ApplyOption) (privateCloudApiOperation, error) {
	switch opName {

	case "updatePrivateCloudUpdateClusterOperation":
		return &updatePrivateCloudUpdateClusterOperation{FieldDiffs: fieldDiffs}, nil

	case "updatePrivateCloudUpdatePrivateCloudOperation":
		return &updatePrivateCloudUpdatePrivateCloudOperation{FieldDiffs: fieldDiffs}, nil

	default:
		return nil, fmt.Errorf("no such operation with name: %v", opName)
	}
}

func extractPrivateCloudFields(r *PrivateCloud) error {
	vNetworkConfig := r.NetworkConfig
	if vNetworkConfig == nil {
		// note: explicitly not the empty object.
		vNetworkConfig = &PrivateCloudNetworkConfig{}
	}
	if err := extractPrivateCloudNetworkConfigFields(r, vNetworkConfig); err != nil {
		return err
	}
	if !dcl.IsNotReturnedByServer(vNetworkConfig) {
		r.NetworkConfig = vNetworkConfig
	}
	vManagementCluster := r.ManagementCluster
	if vManagementCluster == nil {
		// note: explicitly not the empty object.
		vManagementCluster = &PrivateCloudManagementCluster{}
	}
	if err := extractPrivateCloudManagementClusterFields(r, vManagementCluster); err != nil {
		return err
	}
	if !dcl.IsNotReturnedByServer(vManagementCluster) {
		r.ManagementCluster = vManagementCluster
	}
	vHcx := r.Hcx
	if vHcx == nil {
		// note: explicitly not the empty object.
		vHcx = &PrivateCloudHcx{}
	}
	if err := extractPrivateCloudHcxFields(r, vHcx); err != nil {
		return err
	}
	if !dcl.IsNotReturnedByServer(vHcx) {
		r.Hcx = vHcx
	}
	vNsx := r.Nsx
	if vNsx == nil {
		// note: explicitly not the empty object.
		vNsx = &PrivateCloudNsx{}
	}
	if err := extractPrivateCloudNsxFields(r, vNsx); err != nil {
		return err
	}
	if !dcl.IsNotReturnedByServer(vNsx) {
		r.Nsx = vNsx
	}
	vVcenter := r.Vcenter
	if vVcenter == nil {
		// note: explicitly not the empty object.
		vVcenter = &PrivateCloudVcenter{}
	}
	if err := extractPrivateCloudVcenterFields(r, vVcenter); err != nil {
		return err
	}
	if !dcl.IsNotReturnedByServer(vVcenter) {
		r.Vcenter = vVcenter
	}
	return nil
}
func extractPrivateCloudNetworkConfigFields(r *PrivateCloud, o *PrivateCloudNetworkConfig) error {
	return nil
}
func extractPrivateCloudManagementClusterFields(r *PrivateCloud, o *PrivateCloudManagementCluster) error {
	return nil
}
func extractPrivateCloudConditionsFields(r *PrivateCloud, o *PrivateCloudConditions) error {
	return nil
}
func extractPrivateCloudHcxFields(r *PrivateCloud, o *PrivateCloudHcx) error {
	return nil
}
func extractPrivateCloudNsxFields(r *PrivateCloud, o *PrivateCloudNsx) error {
	return nil
}
func extractPrivateCloudVcenterFields(r *PrivateCloud, o *PrivateCloudVcenter) error {
	return nil
}

func postReadExtractPrivateCloudFields(r *PrivateCloud) error {
	vNetworkConfig := r.NetworkConfig
	if vNetworkConfig == nil {
		// note: explicitly not the empty object.
		vNetworkConfig = &PrivateCloudNetworkConfig{}
	}
	if err := postReadExtractPrivateCloudNetworkConfigFields(r, vNetworkConfig); err != nil {
		return err
	}
	if !dcl.IsNotReturnedByServer(vNetworkConfig) {
		r.NetworkConfig = vNetworkConfig
	}
	vManagementCluster := r.ManagementCluster
	if vManagementCluster == nil {
		// note: explicitly not the empty object.
		vManagementCluster = &PrivateCloudManagementCluster{}
	}
	if err := postReadExtractPrivateCloudManagementClusterFields(r, vManagementCluster); err != nil {
		return err
	}
	if !dcl.IsNotReturnedByServer(vManagementCluster) {
		r.ManagementCluster = vManagementCluster
	}
	vHcx := r.Hcx
	if vHcx == nil {
		// note: explicitly not the empty object.
		vHcx = &PrivateCloudHcx{}
	}
	if err := postReadExtractPrivateCloudHcxFields(r, vHcx); err != nil {
		return err
	}
	if !dcl.IsNotReturnedByServer(vHcx) {
		r.Hcx = vHcx
	}
	vNsx := r.Nsx
	if vNsx == nil {
		// note: explicitly not the empty object.
		vNsx = &PrivateCloudNsx{}
	}
	if err := postReadExtractPrivateCloudNsxFields(r, vNsx); err != nil {
		return err
	}
	if !dcl.IsNotReturnedByServer(vNsx) {
		r.Nsx = vNsx
	}
	vVcenter := r.Vcenter
	if vVcenter == nil {
		// note: explicitly not the empty object.
		vVcenter = &PrivateCloudVcenter{}
	}
	if err := postReadExtractPrivateCloudVcenterFields(r, vVcenter); err != nil {
		return err
	}
	if !dcl.IsNotReturnedByServer(vVcenter) {
		r.Vcenter = vVcenter
	}
	return nil
}
func postReadExtractPrivateCloudNetworkConfigFields(r *PrivateCloud, o *PrivateCloudNetworkConfig) error {
	return nil
}
func postReadExtractPrivateCloudManagementClusterFields(r *PrivateCloud, o *PrivateCloudManagementCluster) error {
	return nil
}
func postReadExtractPrivateCloudConditionsFields(r *PrivateCloud, o *PrivateCloudConditions) error {
	return nil
}
func postReadExtractPrivateCloudHcxFields(r *PrivateCloud, o *PrivateCloudHcx) error {
	return nil
}
func postReadExtractPrivateCloudNsxFields(r *PrivateCloud, o *PrivateCloudNsx) error {
	return nil
}
func postReadExtractPrivateCloudVcenterFields(r *PrivateCloud, o *PrivateCloudVcenter) error {
	return nil
}
