// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	maistrav1 "github.com/maistra/istio-operator/pkg/apis/maistra/v1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeServiceMeshControlPlanes implements ServiceMeshControlPlaneInterface
type FakeServiceMeshControlPlanes struct {
	Fake *FakeMaistraV1
	ns   string
}

var servicemeshcontrolplanesResource = schema.GroupVersionResource{Group: "maistra.io", Version: "v1", Resource: "servicemeshcontrolplanes"}

var servicemeshcontrolplanesKind = schema.GroupVersionKind{Group: "maistra.io", Version: "v1", Kind: "ServiceMeshControlPlane"}

// Get takes name of the serviceMeshControlPlane, and returns the corresponding serviceMeshControlPlane object, and an error if there is any.
func (c *FakeServiceMeshControlPlanes) Get(name string, options v1.GetOptions) (result *maistrav1.ServiceMeshControlPlane, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(servicemeshcontrolplanesResource, c.ns, name), &maistrav1.ServiceMeshControlPlane{})

	if obj == nil {
		return nil, err
	}
	return obj.(*maistrav1.ServiceMeshControlPlane), err
}

// List takes label and field selectors, and returns the list of ServiceMeshControlPlanes that match those selectors.
func (c *FakeServiceMeshControlPlanes) List(opts v1.ListOptions) (result *maistrav1.ServiceMeshControlPlaneList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(servicemeshcontrolplanesResource, servicemeshcontrolplanesKind, c.ns, opts), &maistrav1.ServiceMeshControlPlaneList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &maistrav1.ServiceMeshControlPlaneList{ListMeta: obj.(*maistrav1.ServiceMeshControlPlaneList).ListMeta}
	for _, item := range obj.(*maistrav1.ServiceMeshControlPlaneList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested serviceMeshControlPlanes.
func (c *FakeServiceMeshControlPlanes) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(servicemeshcontrolplanesResource, c.ns, opts))

}

// Create takes the representation of a serviceMeshControlPlane and creates it.  Returns the server's representation of the serviceMeshControlPlane, and an error, if there is any.
func (c *FakeServiceMeshControlPlanes) Create(serviceMeshControlPlane *maistrav1.ServiceMeshControlPlane) (result *maistrav1.ServiceMeshControlPlane, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(servicemeshcontrolplanesResource, c.ns, serviceMeshControlPlane), &maistrav1.ServiceMeshControlPlane{})

	if obj == nil {
		return nil, err
	}
	return obj.(*maistrav1.ServiceMeshControlPlane), err
}

// Update takes the representation of a serviceMeshControlPlane and updates it. Returns the server's representation of the serviceMeshControlPlane, and an error, if there is any.
func (c *FakeServiceMeshControlPlanes) Update(serviceMeshControlPlane *maistrav1.ServiceMeshControlPlane) (result *maistrav1.ServiceMeshControlPlane, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(servicemeshcontrolplanesResource, c.ns, serviceMeshControlPlane), &maistrav1.ServiceMeshControlPlane{})

	if obj == nil {
		return nil, err
	}
	return obj.(*maistrav1.ServiceMeshControlPlane), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeServiceMeshControlPlanes) UpdateStatus(serviceMeshControlPlane *maistrav1.ServiceMeshControlPlane) (*maistrav1.ServiceMeshControlPlane, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(servicemeshcontrolplanesResource, "status", c.ns, serviceMeshControlPlane), &maistrav1.ServiceMeshControlPlane{})

	if obj == nil {
		return nil, err
	}
	return obj.(*maistrav1.ServiceMeshControlPlane), err
}

// Delete takes name of the serviceMeshControlPlane and deletes it. Returns an error if one occurs.
func (c *FakeServiceMeshControlPlanes) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(servicemeshcontrolplanesResource, c.ns, name), &maistrav1.ServiceMeshControlPlane{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeServiceMeshControlPlanes) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(servicemeshcontrolplanesResource, c.ns, listOptions)

	_, err := c.Fake.Invokes(action, &maistrav1.ServiceMeshControlPlaneList{})
	return err
}

// Patch applies the patch and returns the patched serviceMeshControlPlane.
func (c *FakeServiceMeshControlPlanes) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *maistrav1.ServiceMeshControlPlane, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(servicemeshcontrolplanesResource, c.ns, name, pt, data, subresources...), &maistrav1.ServiceMeshControlPlane{})

	if obj == nil {
		return nil, err
	}
	return obj.(*maistrav1.ServiceMeshControlPlane), err
}
