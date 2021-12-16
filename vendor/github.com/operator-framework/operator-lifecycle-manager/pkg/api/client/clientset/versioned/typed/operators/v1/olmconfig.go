/*
Copyright Red Hat, Inc.

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

package v1

import (
	"context"
	"time"

	v1 "github.com/operator-framework/api/pkg/operators/v1"
	scheme "github.com/operator-framework/operator-lifecycle-manager/pkg/api/client/clientset/versioned/scheme"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// OLMConfigsGetter has a method to return a OLMConfigInterface.
// A group's client should implement this interface.
type OLMConfigsGetter interface {
	OLMConfigs() OLMConfigInterface
}

// OLMConfigInterface has methods to work with OLMConfig resources.
type OLMConfigInterface interface {
	Create(ctx context.Context, oLMConfig *v1.OLMConfig, opts metav1.CreateOptions) (*v1.OLMConfig, error)
	Update(ctx context.Context, oLMConfig *v1.OLMConfig, opts metav1.UpdateOptions) (*v1.OLMConfig, error)
	UpdateStatus(ctx context.Context, oLMConfig *v1.OLMConfig, opts metav1.UpdateOptions) (*v1.OLMConfig, error)
	Delete(ctx context.Context, name string, opts metav1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts metav1.DeleteOptions, listOpts metav1.ListOptions) error
	Get(ctx context.Context, name string, opts metav1.GetOptions) (*v1.OLMConfig, error)
	List(ctx context.Context, opts metav1.ListOptions) (*v1.OLMConfigList, error)
	Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts metav1.PatchOptions, subresources ...string) (result *v1.OLMConfig, err error)
	OLMConfigExpansion
}

// oLMConfigs implements OLMConfigInterface
type oLMConfigs struct {
	client rest.Interface
}

// newOLMConfigs returns a OLMConfigs
func newOLMConfigs(c *OperatorsV1Client) *oLMConfigs {
	return &oLMConfigs{
		client: c.RESTClient(),
	}
}

// Get takes name of the oLMConfig, and returns the corresponding oLMConfig object, and an error if there is any.
func (c *oLMConfigs) Get(ctx context.Context, name string, options metav1.GetOptions) (result *v1.OLMConfig, err error) {
	result = &v1.OLMConfig{}
	err = c.client.Get().
		Resource("olmconfigs").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do(ctx).
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of OLMConfigs that match those selectors.
func (c *oLMConfigs) List(ctx context.Context, opts metav1.ListOptions) (result *v1.OLMConfigList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1.OLMConfigList{}
	err = c.client.Get().
		Resource("olmconfigs").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do(ctx).
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested oLMConfigs.
func (c *oLMConfigs) Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Resource("olmconfigs").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch(ctx)
}

// Create takes the representation of a oLMConfig and creates it.  Returns the server's representation of the oLMConfig, and an error, if there is any.
func (c *oLMConfigs) Create(ctx context.Context, oLMConfig *v1.OLMConfig, opts metav1.CreateOptions) (result *v1.OLMConfig, err error) {
	result = &v1.OLMConfig{}
	err = c.client.Post().
		Resource("olmconfigs").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(oLMConfig).
		Do(ctx).
		Into(result)
	return
}

// Update takes the representation of a oLMConfig and updates it. Returns the server's representation of the oLMConfig, and an error, if there is any.
func (c *oLMConfigs) Update(ctx context.Context, oLMConfig *v1.OLMConfig, opts metav1.UpdateOptions) (result *v1.OLMConfig, err error) {
	result = &v1.OLMConfig{}
	err = c.client.Put().
		Resource("olmconfigs").
		Name(oLMConfig.Name).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(oLMConfig).
		Do(ctx).
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *oLMConfigs) UpdateStatus(ctx context.Context, oLMConfig *v1.OLMConfig, opts metav1.UpdateOptions) (result *v1.OLMConfig, err error) {
	result = &v1.OLMConfig{}
	err = c.client.Put().
		Resource("olmconfigs").
		Name(oLMConfig.Name).
		SubResource("status").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(oLMConfig).
		Do(ctx).
		Into(result)
	return
}

// Delete takes name of the oLMConfig and deletes it. Returns an error if one occurs.
func (c *oLMConfigs) Delete(ctx context.Context, name string, opts metav1.DeleteOptions) error {
	return c.client.Delete().
		Resource("olmconfigs").
		Name(name).
		Body(&opts).
		Do(ctx).
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *oLMConfigs) DeleteCollection(ctx context.Context, opts metav1.DeleteOptions, listOpts metav1.ListOptions) error {
	var timeout time.Duration
	if listOpts.TimeoutSeconds != nil {
		timeout = time.Duration(*listOpts.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Resource("olmconfigs").
		VersionedParams(&listOpts, scheme.ParameterCodec).
		Timeout(timeout).
		Body(&opts).
		Do(ctx).
		Error()
}

// Patch applies the patch and returns the patched oLMConfig.
func (c *oLMConfigs) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts metav1.PatchOptions, subresources ...string) (result *v1.OLMConfig, err error) {
	result = &v1.OLMConfig{}
	err = c.client.Patch(pt).
		Resource("olmconfigs").
		Name(name).
		SubResource(subresources...).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(data).
		Do(ctx).
		Into(result)
	return
}
