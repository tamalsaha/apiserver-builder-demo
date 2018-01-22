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
package fake

import (
	insect "github.com/tamalsaha/apiserver-builder-demo/pkg/apis/insect"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeBees implements BeeInterface
type FakeBees struct {
	Fake *FakeInsect
	ns   string
}

var beesResource = schema.GroupVersionResource{Group: "insect.tamalsaha.com", Version: "", Resource: "bees"}

var beesKind = schema.GroupVersionKind{Group: "insect.tamalsaha.com", Version: "", Kind: "Bee"}

// Get takes name of the bee, and returns the corresponding bee object, and an error if there is any.
func (c *FakeBees) Get(name string, options v1.GetOptions) (result *insect.Bee, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(beesResource, c.ns, name), &insect.Bee{})

	if obj == nil {
		return nil, err
	}
	return obj.(*insect.Bee), err
}

// List takes label and field selectors, and returns the list of Bees that match those selectors.
func (c *FakeBees) List(opts v1.ListOptions) (result *insect.BeeList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(beesResource, beesKind, c.ns, opts), &insect.BeeList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &insect.BeeList{}
	for _, item := range obj.(*insect.BeeList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested bees.
func (c *FakeBees) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(beesResource, c.ns, opts))

}

// Create takes the representation of a bee and creates it.  Returns the server's representation of the bee, and an error, if there is any.
func (c *FakeBees) Create(bee *insect.Bee) (result *insect.Bee, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(beesResource, c.ns, bee), &insect.Bee{})

	if obj == nil {
		return nil, err
	}
	return obj.(*insect.Bee), err
}

// Update takes the representation of a bee and updates it. Returns the server's representation of the bee, and an error, if there is any.
func (c *FakeBees) Update(bee *insect.Bee) (result *insect.Bee, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(beesResource, c.ns, bee), &insect.Bee{})

	if obj == nil {
		return nil, err
	}
	return obj.(*insect.Bee), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeBees) UpdateStatus(bee *insect.Bee) (*insect.Bee, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(beesResource, "status", c.ns, bee), &insect.Bee{})

	if obj == nil {
		return nil, err
	}
	return obj.(*insect.Bee), err
}

// Delete takes name of the bee and deletes it. Returns an error if one occurs.
func (c *FakeBees) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(beesResource, c.ns, name), &insect.Bee{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeBees) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(beesResource, c.ns, listOptions)

	_, err := c.Fake.Invokes(action, &insect.BeeList{})
	return err
}

// Patch applies the patch and returns the patched bee.
func (c *FakeBees) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *insect.Bee, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(beesResource, c.ns, name, data, subresources...), &insect.Bee{})

	if obj == nil {
		return nil, err
	}
	return obj.(*insect.Bee), err
}
