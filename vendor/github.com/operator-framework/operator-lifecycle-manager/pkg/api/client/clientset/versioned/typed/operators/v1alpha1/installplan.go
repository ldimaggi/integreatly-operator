/*
Copyright 2019 Red Hat, Inc.

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

package v1alpha1

import (
	v1alpha1 "github.com/operator-framework/operator-lifecycle-manager/pkg/api/apis/operators/v1alpha1"
	scheme "github.com/operator-framework/operator-lifecycle-manager/pkg/api/client/clientset/versioned/scheme"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// InstallPlansGetter has a method to return a InstallPlanInterface.
// A group's client should implement this interface.
type InstallPlansGetter interface {
	InstallPlans(namespace string) InstallPlanInterface
}

// InstallPlanInterface has methods to work with InstallPlan resources.
type InstallPlanInterface interface {
	Create(*v1alpha1.InstallPlan) (*v1alpha1.InstallPlan, error)
	Update(*v1alpha1.InstallPlan) (*v1alpha1.InstallPlan, error)
	UpdateStatus(*v1alpha1.InstallPlan) (*v1alpha1.InstallPlan, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1alpha1.InstallPlan, error)
	List(opts v1.ListOptions) (*v1alpha1.InstallPlanList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.InstallPlan, err error)
	InstallPlanExpansion
}

// installPlans implements InstallPlanInterface
type installPlans struct {
	client rest.Interface
	ns     string
}

// newInstallPlans returns a InstallPlans
func newInstallPlans(c *OperatorsV1alpha1Client, namespace string) *installPlans {
	return &installPlans{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the installPlan, and returns the corresponding installPlan object, and an error if there is any.
func (c *installPlans) Get(name string, options v1.GetOptions) (result *v1alpha1.InstallPlan, err error) {
	result = &v1alpha1.InstallPlan{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("installplans").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of InstallPlans that match those selectors.
func (c *installPlans) List(opts v1.ListOptions) (result *v1alpha1.InstallPlanList, err error) {
	result = &v1alpha1.InstallPlanList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("installplans").
		VersionedParams(&opts, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested installPlans.
func (c *installPlans) Watch(opts v1.ListOptions) (watch.Interface, error) {
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("installplans").
		VersionedParams(&opts, scheme.ParameterCodec).
		Watch()
}

// Create takes the representation of a installPlan and creates it.  Returns the server's representation of the installPlan, and an error, if there is any.
func (c *installPlans) Create(installPlan *v1alpha1.InstallPlan) (result *v1alpha1.InstallPlan, err error) {
	result = &v1alpha1.InstallPlan{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("installplans").
		Body(installPlan).
		Do().
		Into(result)
	return
}

// Update takes the representation of a installPlan and updates it. Returns the server's representation of the installPlan, and an error, if there is any.
func (c *installPlans) Update(installPlan *v1alpha1.InstallPlan) (result *v1alpha1.InstallPlan, err error) {
	result = &v1alpha1.InstallPlan{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("installplans").
		Name(installPlan.Name).
		Body(installPlan).
		Do().
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().

func (c *installPlans) UpdateStatus(installPlan *v1alpha1.InstallPlan) (result *v1alpha1.InstallPlan, err error) {
	result = &v1alpha1.InstallPlan{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("installplans").
		Name(installPlan.Name).
		SubResource("status").
		Body(installPlan).
		Do().
		Into(result)
	return
}

// Delete takes name of the installPlan and deletes it. Returns an error if one occurs.
func (c *installPlans) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("installplans").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *installPlans) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("installplans").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched installPlan.
func (c *installPlans) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.InstallPlan, err error) {
	result = &v1alpha1.InstallPlan{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("installplans").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
