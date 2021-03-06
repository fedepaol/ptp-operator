/*
Copyright The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	ptpv1 "github.com/openshift/ptp-operator/pkg/apis/ptp/v1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakePtpOperatorConfigs implements PtpOperatorConfigInterface
type FakePtpOperatorConfigs struct {
	Fake *FakePtpV1
	ns   string
}

var ptpoperatorconfigsResource = schema.GroupVersionResource{Group: "ptp.openshift.io", Version: "v1", Resource: "ptpoperatorconfigs"}

var ptpoperatorconfigsKind = schema.GroupVersionKind{Group: "ptp.openshift.io", Version: "v1", Kind: "PtpOperatorConfig"}

// Get takes name of the ptpOperatorConfig, and returns the corresponding ptpOperatorConfig object, and an error if there is any.
func (c *FakePtpOperatorConfigs) Get(name string, options v1.GetOptions) (result *ptpv1.PtpOperatorConfig, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(ptpoperatorconfigsResource, c.ns, name), &ptpv1.PtpOperatorConfig{})

	if obj == nil {
		return nil, err
	}
	return obj.(*ptpv1.PtpOperatorConfig), err
}

// List takes label and field selectors, and returns the list of PtpOperatorConfigs that match those selectors.
func (c *FakePtpOperatorConfigs) List(opts v1.ListOptions) (result *ptpv1.PtpOperatorConfigList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(ptpoperatorconfigsResource, ptpoperatorconfigsKind, c.ns, opts), &ptpv1.PtpOperatorConfigList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &ptpv1.PtpOperatorConfigList{ListMeta: obj.(*ptpv1.PtpOperatorConfigList).ListMeta}
	for _, item := range obj.(*ptpv1.PtpOperatorConfigList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested ptpOperatorConfigs.
func (c *FakePtpOperatorConfigs) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(ptpoperatorconfigsResource, c.ns, opts))

}

// Create takes the representation of a ptpOperatorConfig and creates it.  Returns the server's representation of the ptpOperatorConfig, and an error, if there is any.
func (c *FakePtpOperatorConfigs) Create(ptpOperatorConfig *ptpv1.PtpOperatorConfig) (result *ptpv1.PtpOperatorConfig, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(ptpoperatorconfigsResource, c.ns, ptpOperatorConfig), &ptpv1.PtpOperatorConfig{})

	if obj == nil {
		return nil, err
	}
	return obj.(*ptpv1.PtpOperatorConfig), err
}

// Update takes the representation of a ptpOperatorConfig and updates it. Returns the server's representation of the ptpOperatorConfig, and an error, if there is any.
func (c *FakePtpOperatorConfigs) Update(ptpOperatorConfig *ptpv1.PtpOperatorConfig) (result *ptpv1.PtpOperatorConfig, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(ptpoperatorconfigsResource, c.ns, ptpOperatorConfig), &ptpv1.PtpOperatorConfig{})

	if obj == nil {
		return nil, err
	}
	return obj.(*ptpv1.PtpOperatorConfig), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakePtpOperatorConfigs) UpdateStatus(ptpOperatorConfig *ptpv1.PtpOperatorConfig) (*ptpv1.PtpOperatorConfig, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(ptpoperatorconfigsResource, "status", c.ns, ptpOperatorConfig), &ptpv1.PtpOperatorConfig{})

	if obj == nil {
		return nil, err
	}
	return obj.(*ptpv1.PtpOperatorConfig), err
}

// Delete takes name of the ptpOperatorConfig and deletes it. Returns an error if one occurs.
func (c *FakePtpOperatorConfigs) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(ptpoperatorconfigsResource, c.ns, name), &ptpv1.PtpOperatorConfig{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakePtpOperatorConfigs) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(ptpoperatorconfigsResource, c.ns, listOptions)

	_, err := c.Fake.Invokes(action, &ptpv1.PtpOperatorConfigList{})
	return err
}

// Patch applies the patch and returns the patched ptpOperatorConfig.
func (c *FakePtpOperatorConfigs) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *ptpv1.PtpOperatorConfig, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(ptpoperatorconfigsResource, c.ns, name, pt, data, subresources...), &ptpv1.PtpOperatorConfig{})

	if obj == nil {
		return nil, err
	}
	return obj.(*ptpv1.PtpOperatorConfig), err
}
