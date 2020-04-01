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

package v1alpha1

import (
	"time"

	v1alpha1 "github.com/operator-framework/operator-lifecycle-manager/pkg/api/apis/operators/v1alpha1"
	scheme "github.com/operator-framework/operator-lifecycle-manager/pkg/api/client/clientset/versioned/scheme"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// ClusterServiceVersionsGetter has a method to return a ClusterServiceVersionInterface.
// A group's client should implement this interface.
type ClusterServiceVersionsGetter interface {
	ClusterServiceVersions(namespace string) ClusterServiceVersionInterface
}

// ClusterServiceVersionInterface has methods to work with ClusterServiceVersion resources.
type ClusterServiceVersionInterface interface {
	Create(*v1alpha1.ClusterServiceVersion) (*v1alpha1.ClusterServiceVersion, error)
	Update(*v1alpha1.ClusterServiceVersion) (*v1alpha1.ClusterServiceVersion, error)
	UpdateStatus(*v1alpha1.ClusterServiceVersion) (*v1alpha1.ClusterServiceVersion, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1alpha1.ClusterServiceVersion, error)
	List(opts v1.ListOptions) (*v1alpha1.ClusterServiceVersionList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.ClusterServiceVersion, err error)
	ClusterServiceVersionExpansion
}

// clusterServiceVersions implements ClusterServiceVersionInterface
type clusterServiceVersions struct {
	client rest.Interface
	ns     string
}

// newClusterServiceVersions returns a ClusterServiceVersions
func newClusterServiceVersions(c *OperatorsV1alpha1Client, namespace string) *clusterServiceVersions {
	return &clusterServiceVersions{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the clusterServiceVersion, and returns the corresponding clusterServiceVersion object, and an error if there is any.
func (c *clusterServiceVersions) Get(name string, options v1.GetOptions) (result *v1alpha1.ClusterServiceVersion, err error) {
	result = &v1alpha1.ClusterServiceVersion{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("clusterserviceversions").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of ClusterServiceVersions that match those selectors.
func (c *clusterServiceVersions) List(opts v1.ListOptions) (result *v1alpha1.ClusterServiceVersionList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.ClusterServiceVersionList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("clusterserviceversions").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested clusterServiceVersions.
func (c *clusterServiceVersions) Watch(opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("clusterserviceversions").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch()
}

// Create takes the representation of a clusterServiceVersion and creates it.  Returns the server's representation of the clusterServiceVersion, and an error, if there is any.
func (c *clusterServiceVersions) Create(clusterServiceVersion *v1alpha1.ClusterServiceVersion) (result *v1alpha1.ClusterServiceVersion, err error) {
	result = &v1alpha1.ClusterServiceVersion{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("clusterserviceversions").
		Body(clusterServiceVersion).
		Do().
		Into(result)
	return
}

// Update takes the representation of a clusterServiceVersion and updates it. Returns the server's representation of the clusterServiceVersion, and an error, if there is any.
func (c *clusterServiceVersions) Update(clusterServiceVersion *v1alpha1.ClusterServiceVersion) (result *v1alpha1.ClusterServiceVersion, err error) {
	result = &v1alpha1.ClusterServiceVersion{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("clusterserviceversions").
		Name(clusterServiceVersion.Name).
		Body(clusterServiceVersion).
		Do().
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().

func (c *clusterServiceVersions) UpdateStatus(clusterServiceVersion *v1alpha1.ClusterServiceVersion) (result *v1alpha1.ClusterServiceVersion, err error) {
	result = &v1alpha1.ClusterServiceVersion{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("clusterserviceversions").
		Name(clusterServiceVersion.Name).
		SubResource("status").
		Body(clusterServiceVersion).
		Do().
		Into(result)
	return
}

// Delete takes name of the clusterServiceVersion and deletes it. Returns an error if one occurs.
func (c *clusterServiceVersions) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("clusterserviceversions").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *clusterServiceVersions) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	var timeout time.Duration
	if listOptions.TimeoutSeconds != nil {
		timeout = time.Duration(*listOptions.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("clusterserviceversions").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Timeout(timeout).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched clusterServiceVersion.
func (c *clusterServiceVersions) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.ClusterServiceVersion, err error) {
	result = &v1alpha1.ClusterServiceVersion{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("clusterserviceversions").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
