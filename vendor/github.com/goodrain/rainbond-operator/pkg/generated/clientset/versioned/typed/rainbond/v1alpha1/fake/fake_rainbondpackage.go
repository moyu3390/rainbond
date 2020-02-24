// RAINBOND, Application Management Platform
// Copyright (C) 2014-2020 Goodrain Co., Ltd.

// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version. For any non-GPL usage of Rainbond,
// one or multiple Commercial Licenses authorized by Goodrain Co., Ltd.
// must be obtained first.

// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU General Public License for more details.

// You should have received a copy of the GNU General Public License
// along with this program. If not, see <http://www.gnu.org/licenses/>.

// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	v1alpha1 "github.com/goodrain/rainbond-operator/pkg/apis/rainbond/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeRainbondPackages implements RainbondPackageInterface
type FakeRainbondPackages struct {
	Fake *FakeRainbondV1alpha1
	ns   string
}

var rainbondpackagesResource = schema.GroupVersionResource{Group: "rainbond.io", Version: "v1alpha1", Resource: "rainbondpackages"}

var rainbondpackagesKind = schema.GroupVersionKind{Group: "rainbond.io", Version: "v1alpha1", Kind: "RainbondPackage"}

// Get takes name of the rainbondPackage, and returns the corresponding rainbondPackage object, and an error if there is any.
func (c *FakeRainbondPackages) Get(name string, options v1.GetOptions) (result *v1alpha1.RainbondPackage, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(rainbondpackagesResource, c.ns, name), &v1alpha1.RainbondPackage{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.RainbondPackage), err
}

// List takes label and field selectors, and returns the list of RainbondPackages that match those selectors.
func (c *FakeRainbondPackages) List(opts v1.ListOptions) (result *v1alpha1.RainbondPackageList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(rainbondpackagesResource, rainbondpackagesKind, c.ns, opts), &v1alpha1.RainbondPackageList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.RainbondPackageList{ListMeta: obj.(*v1alpha1.RainbondPackageList).ListMeta}
	for _, item := range obj.(*v1alpha1.RainbondPackageList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested rainbondPackages.
func (c *FakeRainbondPackages) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(rainbondpackagesResource, c.ns, opts))

}

// Create takes the representation of a rainbondPackage and creates it.  Returns the server's representation of the rainbondPackage, and an error, if there is any.
func (c *FakeRainbondPackages) Create(rainbondPackage *v1alpha1.RainbondPackage) (result *v1alpha1.RainbondPackage, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(rainbondpackagesResource, c.ns, rainbondPackage), &v1alpha1.RainbondPackage{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.RainbondPackage), err
}

// Update takes the representation of a rainbondPackage and updates it. Returns the server's representation of the rainbondPackage, and an error, if there is any.
func (c *FakeRainbondPackages) Update(rainbondPackage *v1alpha1.RainbondPackage) (result *v1alpha1.RainbondPackage, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(rainbondpackagesResource, c.ns, rainbondPackage), &v1alpha1.RainbondPackage{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.RainbondPackage), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeRainbondPackages) UpdateStatus(rainbondPackage *v1alpha1.RainbondPackage) (*v1alpha1.RainbondPackage, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(rainbondpackagesResource, "status", c.ns, rainbondPackage), &v1alpha1.RainbondPackage{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.RainbondPackage), err
}

// Delete takes name of the rainbondPackage and deletes it. Returns an error if one occurs.
func (c *FakeRainbondPackages) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(rainbondpackagesResource, c.ns, name), &v1alpha1.RainbondPackage{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeRainbondPackages) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(rainbondpackagesResource, c.ns, listOptions)

	_, err := c.Fake.Invokes(action, &v1alpha1.RainbondPackageList{})
	return err
}

// Patch applies the patch and returns the patched rainbondPackage.
func (c *FakeRainbondPackages) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.RainbondPackage, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(rainbondpackagesResource, c.ns, name, pt, data, subresources...), &v1alpha1.RainbondPackage{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.RainbondPackage), err
}