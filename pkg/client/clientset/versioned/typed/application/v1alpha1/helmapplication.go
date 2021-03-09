/*
Copyright 2020 The KubeSphere Authors.

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
	"context"
	"time"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
	v1alpha1 "kubesphere.io/kubesphere/pkg/apis/application/v1alpha1"
	scheme "kubesphere.io/kubesphere/pkg/client/clientset/versioned/scheme"
)

// HelmApplicationsGetter has a method to return a HelmApplicationInterface.
// A group's client should implement this interface.
type HelmApplicationsGetter interface {
	HelmApplications() HelmApplicationInterface
}

// HelmApplicationInterface has methods to work with HelmApplication resources.
type HelmApplicationInterface interface {
	Create(ctx context.Context, helmApplication *v1alpha1.HelmApplication, opts v1.CreateOptions) (*v1alpha1.HelmApplication, error)
	Update(ctx context.Context, helmApplication *v1alpha1.HelmApplication, opts v1.UpdateOptions) (*v1alpha1.HelmApplication, error)
	UpdateStatus(ctx context.Context, helmApplication *v1alpha1.HelmApplication, opts v1.UpdateOptions) (*v1alpha1.HelmApplication, error)
	Delete(ctx context.Context, name string, opts v1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error
	Get(ctx context.Context, name string, opts v1.GetOptions) (*v1alpha1.HelmApplication, error)
	List(ctx context.Context, opts v1.ListOptions) (*v1alpha1.HelmApplicationList, error)
	Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.HelmApplication, err error)
	HelmApplicationExpansion
}

// helmApplications implements HelmApplicationInterface
type helmApplications struct {
	client rest.Interface
}

// newHelmApplications returns a HelmApplications
func newHelmApplications(c *ApplicationV1alpha1Client) *helmApplications {
	return &helmApplications{
		client: c.RESTClient(),
	}
}

// Get takes name of the helmApplication, and returns the corresponding helmApplication object, and an error if there is any.
func (c *helmApplications) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.HelmApplication, err error) {
	result = &v1alpha1.HelmApplication{}
	err = c.client.Get().
		Resource("helmapplications").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do(ctx).
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of HelmApplications that match those selectors.
func (c *helmApplications) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.HelmApplicationList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.HelmApplicationList{}
	err = c.client.Get().
		Resource("helmapplications").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do(ctx).
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested helmApplications.
func (c *helmApplications) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Resource("helmapplications").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch(ctx)
}

// Create takes the representation of a helmApplication and creates it.  Returns the server's representation of the helmApplication, and an error, if there is any.
func (c *helmApplications) Create(ctx context.Context, helmApplication *v1alpha1.HelmApplication, opts v1.CreateOptions) (result *v1alpha1.HelmApplication, err error) {
	result = &v1alpha1.HelmApplication{}
	err = c.client.Post().
		Resource("helmapplications").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(helmApplication).
		Do(ctx).
		Into(result)
	return
}

// Update takes the representation of a helmApplication and updates it. Returns the server's representation of the helmApplication, and an error, if there is any.
func (c *helmApplications) Update(ctx context.Context, helmApplication *v1alpha1.HelmApplication, opts v1.UpdateOptions) (result *v1alpha1.HelmApplication, err error) {
	result = &v1alpha1.HelmApplication{}
	err = c.client.Put().
		Resource("helmapplications").
		Name(helmApplication.Name).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(helmApplication).
		Do(ctx).
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *helmApplications) UpdateStatus(ctx context.Context, helmApplication *v1alpha1.HelmApplication, opts v1.UpdateOptions) (result *v1alpha1.HelmApplication, err error) {
	result = &v1alpha1.HelmApplication{}
	err = c.client.Put().
		Resource("helmapplications").
		Name(helmApplication.Name).
		SubResource("status").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(helmApplication).
		Do(ctx).
		Into(result)
	return
}

// Delete takes name of the helmApplication and deletes it. Returns an error if one occurs.
func (c *helmApplications) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	return c.client.Delete().
		Resource("helmapplications").
		Name(name).
		Body(&opts).
		Do(ctx).
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *helmApplications) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	var timeout time.Duration
	if listOpts.TimeoutSeconds != nil {
		timeout = time.Duration(*listOpts.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Resource("helmapplications").
		VersionedParams(&listOpts, scheme.ParameterCodec).
		Timeout(timeout).
		Body(&opts).
		Do(ctx).
		Error()
}

// Patch applies the patch and returns the patched helmApplication.
func (c *helmApplications) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.HelmApplication, err error) {
	result = &v1alpha1.HelmApplication{}
	err = c.client.Patch(pt).
		Resource("helmapplications").
		Name(name).
		SubResource(subresources...).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(data).
		Do(ctx).
		Into(result)
	return
}
