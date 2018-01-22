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
package internalversion

import (
	insect "github.com/tamalsaha/apiserver-builder-demo/pkg/apis/insect"
	scheme "github.com/tamalsaha/apiserver-builder-demo/pkg/client/clientset_generated/internalclientset/scheme"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// BeesGetter has a method to return a BeeInterface.
// A group's client should implement this interface.
type BeesGetter interface {
	Bees(namespace string) BeeInterface
}

// BeeInterface has methods to work with Bee resources.
type BeeInterface interface {
	Create(*insect.Bee) (*insect.Bee, error)
	Update(*insect.Bee) (*insect.Bee, error)
	UpdateStatus(*insect.Bee) (*insect.Bee, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*insect.Bee, error)
	List(opts v1.ListOptions) (*insect.BeeList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *insect.Bee, err error)
	BeeExpansion
}

// bees implements BeeInterface
type bees struct {
	client rest.Interface
	ns     string
}

// newBees returns a Bees
func newBees(c *InsectClient, namespace string) *bees {
	return &bees{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the bee, and returns the corresponding bee object, and an error if there is any.
func (c *bees) Get(name string, options v1.GetOptions) (result *insect.Bee, err error) {
	result = &insect.Bee{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("bees").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of Bees that match those selectors.
func (c *bees) List(opts v1.ListOptions) (result *insect.BeeList, err error) {
	result = &insect.BeeList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("bees").
		VersionedParams(&opts, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested bees.
func (c *bees) Watch(opts v1.ListOptions) (watch.Interface, error) {
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("bees").
		VersionedParams(&opts, scheme.ParameterCodec).
		Watch()
}

// Create takes the representation of a bee and creates it.  Returns the server's representation of the bee, and an error, if there is any.
func (c *bees) Create(bee *insect.Bee) (result *insect.Bee, err error) {
	result = &insect.Bee{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("bees").
		Body(bee).
		Do().
		Into(result)
	return
}

// Update takes the representation of a bee and updates it. Returns the server's representation of the bee, and an error, if there is any.
func (c *bees) Update(bee *insect.Bee) (result *insect.Bee, err error) {
	result = &insect.Bee{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("bees").
		Name(bee.Name).
		Body(bee).
		Do().
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().

func (c *bees) UpdateStatus(bee *insect.Bee) (result *insect.Bee, err error) {
	result = &insect.Bee{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("bees").
		Name(bee.Name).
		SubResource("status").
		Body(bee).
		Do().
		Into(result)
	return
}

// Delete takes name of the bee and deletes it. Returns an error if one occurs.
func (c *bees) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("bees").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *bees) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("bees").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched bee.
func (c *bees) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *insect.Bee, err error) {
	result = &insect.Bee{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("bees").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
