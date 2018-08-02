/*
Copyright 2018 The Kubernetes Authors.

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
	v1alpha1 "github.com/christopherhein/aws-operator/pkg/apis/operator.aws/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeCloudFormationTemplates implements CloudFormationTemplateInterface
type FakeCloudFormationTemplates struct {
	Fake *FakeOperatorV1alpha1
	ns   string
}

var cloudformationtemplatesResource = schema.GroupVersionResource{Group: "operator.aws", Version: "v1alpha1", Resource: "cloudformationtemplates"}

var cloudformationtemplatesKind = schema.GroupVersionKind{Group: "operator.aws", Version: "v1alpha1", Kind: "CloudFormationTemplate"}

// Get takes name of the cloudFormationTemplate, and returns the corresponding cloudFormationTemplate object, and an error if there is any.
func (c *FakeCloudFormationTemplates) Get(name string, options v1.GetOptions) (result *v1alpha1.CloudFormationTemplate, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(cloudformationtemplatesResource, c.ns, name), &v1alpha1.CloudFormationTemplate{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.CloudFormationTemplate), err
}

// List takes label and field selectors, and returns the list of CloudFormationTemplates that match those selectors.
func (c *FakeCloudFormationTemplates) List(opts v1.ListOptions) (result *v1alpha1.CloudFormationTemplateList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(cloudformationtemplatesResource, cloudformationtemplatesKind, c.ns, opts), &v1alpha1.CloudFormationTemplateList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.CloudFormationTemplateList{}
	for _, item := range obj.(*v1alpha1.CloudFormationTemplateList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested cloudFormationTemplates.
func (c *FakeCloudFormationTemplates) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(cloudformationtemplatesResource, c.ns, opts))

}

// Create takes the representation of a cloudFormationTemplate and creates it.  Returns the server's representation of the cloudFormationTemplate, and an error, if there is any.
func (c *FakeCloudFormationTemplates) Create(cloudFormationTemplate *v1alpha1.CloudFormationTemplate) (result *v1alpha1.CloudFormationTemplate, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(cloudformationtemplatesResource, c.ns, cloudFormationTemplate), &v1alpha1.CloudFormationTemplate{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.CloudFormationTemplate), err
}

// Update takes the representation of a cloudFormationTemplate and updates it. Returns the server's representation of the cloudFormationTemplate, and an error, if there is any.
func (c *FakeCloudFormationTemplates) Update(cloudFormationTemplate *v1alpha1.CloudFormationTemplate) (result *v1alpha1.CloudFormationTemplate, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(cloudformationtemplatesResource, c.ns, cloudFormationTemplate), &v1alpha1.CloudFormationTemplate{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.CloudFormationTemplate), err
}

// Delete takes name of the cloudFormationTemplate and deletes it. Returns an error if one occurs.
func (c *FakeCloudFormationTemplates) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(cloudformationtemplatesResource, c.ns, name), &v1alpha1.CloudFormationTemplate{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeCloudFormationTemplates) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(cloudformationtemplatesResource, c.ns, listOptions)

	_, err := c.Fake.Invokes(action, &v1alpha1.CloudFormationTemplateList{})
	return err
}

// Patch applies the patch and returns the patched cloudFormationTemplate.
func (c *FakeCloudFormationTemplates) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.CloudFormationTemplate, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(cloudformationtemplatesResource, c.ns, name, data, subresources...), &v1alpha1.CloudFormationTemplate{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.CloudFormationTemplate), err
}
