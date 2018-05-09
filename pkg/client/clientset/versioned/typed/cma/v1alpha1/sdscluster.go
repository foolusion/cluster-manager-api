/*
Copyright 2018 The sds cluster Authors.

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
	v1alpha1 "github.com/samsung-cnct/cluster-manager-api/pkg/apis/cma/v1alpha1"
	scheme "github.com/samsung-cnct/cluster-manager-api/pkg/client/clientset/versioned/scheme"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// SDSClustersGetter has a method to return a SDSClusterInterface.
// A group's client should implement this interface.
type SDSClustersGetter interface {
	SDSClusters(namespace string) SDSClusterInterface
}

// SDSClusterInterface has methods to work with SDSCluster resources.
type SDSClusterInterface interface {
	Create(*v1alpha1.SDSCluster) (*v1alpha1.SDSCluster, error)
	Update(*v1alpha1.SDSCluster) (*v1alpha1.SDSCluster, error)
	UpdateStatus(*v1alpha1.SDSCluster) (*v1alpha1.SDSCluster, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1alpha1.SDSCluster, error)
	List(opts v1.ListOptions) (*v1alpha1.SDSClusterList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.SDSCluster, err error)
	SDSClusterExpansion
}

// sDSClusters implements SDSClusterInterface
type sDSClusters struct {
	client rest.Interface
	ns     string
}

// newSDSClusters returns a SDSClusters
func newSDSClusters(c *CmaV1alpha1Client, namespace string) *sDSClusters {
	return &sDSClusters{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the sDSCluster, and returns the corresponding sDSCluster object, and an error if there is any.
func (c *sDSClusters) Get(name string, options v1.GetOptions) (result *v1alpha1.SDSCluster, err error) {
	result = &v1alpha1.SDSCluster{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("sdsclusters").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of SDSClusters that match those selectors.
func (c *sDSClusters) List(opts v1.ListOptions) (result *v1alpha1.SDSClusterList, err error) {
	result = &v1alpha1.SDSClusterList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("sdsclusters").
		VersionedParams(&opts, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested sDSClusters.
func (c *sDSClusters) Watch(opts v1.ListOptions) (watch.Interface, error) {
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("sdsclusters").
		VersionedParams(&opts, scheme.ParameterCodec).
		Watch()
}

// Create takes the representation of a sDSCluster and creates it.  Returns the server's representation of the sDSCluster, and an error, if there is any.
func (c *sDSClusters) Create(sDSCluster *v1alpha1.SDSCluster) (result *v1alpha1.SDSCluster, err error) {
	result = &v1alpha1.SDSCluster{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("sdsclusters").
		Body(sDSCluster).
		Do().
		Into(result)
	return
}

// Update takes the representation of a sDSCluster and updates it. Returns the server's representation of the sDSCluster, and an error, if there is any.
func (c *sDSClusters) Update(sDSCluster *v1alpha1.SDSCluster) (result *v1alpha1.SDSCluster, err error) {
	result = &v1alpha1.SDSCluster{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("sdsclusters").
		Name(sDSCluster.Name).
		Body(sDSCluster).
		Do().
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().

func (c *sDSClusters) UpdateStatus(sDSCluster *v1alpha1.SDSCluster) (result *v1alpha1.SDSCluster, err error) {
	result = &v1alpha1.SDSCluster{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("sdsclusters").
		Name(sDSCluster.Name).
		SubResource("status").
		Body(sDSCluster).
		Do().
		Into(result)
	return
}

// Delete takes name of the sDSCluster and deletes it. Returns an error if one occurs.
func (c *sDSClusters) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("sdsclusters").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *sDSClusters) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("sdsclusters").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched sDSCluster.
func (c *sDSClusters) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.SDSCluster, err error) {
	result = &v1alpha1.SDSCluster{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("sdsclusters").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
