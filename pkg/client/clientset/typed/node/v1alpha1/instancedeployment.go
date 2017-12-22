package v1alpha1

import (
	v1alpha1 "code.sapcloud.io/kubernetes/node-controller-manager/pkg/apis/node/v1alpha1"
	scheme "code.sapcloud.io/kubernetes/node-controller-manager/pkg/client/clientset/scheme"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// InstanceDeploymentsGetter has a method to return a InstanceDeploymentInterface.
// A group's client should implement this interface.
type InstanceDeploymentsGetter interface {
	InstanceDeployments() InstanceDeploymentInterface
}

// InstanceDeploymentInterface has methods to work with InstanceDeployment resources.
type InstanceDeploymentInterface interface {
	Create(*v1alpha1.InstanceDeployment) (*v1alpha1.InstanceDeployment, error)
	Update(*v1alpha1.InstanceDeployment) (*v1alpha1.InstanceDeployment, error)
	UpdateStatus(*v1alpha1.InstanceDeployment) (*v1alpha1.InstanceDeployment, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1alpha1.InstanceDeployment, error)
	List(opts v1.ListOptions) (*v1alpha1.InstanceDeploymentList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.InstanceDeployment, err error)
	GetScale(instanceDeploymentName string, options v1.GetOptions) (*v1alpha1.Scale, error)
	UpdateScale(instanceDeploymentName string, scale *v1alpha1.Scale) (*v1alpha1.Scale, error)

	InstanceDeploymentExpansion
}

// instanceDeployments implements InstanceDeploymentInterface
type instanceDeployments struct {
	client rest.Interface
}

// newInstanceDeployments returns a InstanceDeployments
func newInstanceDeployments(c *NodeV1alpha1Client) *instanceDeployments {
	return &instanceDeployments{
		client: c.RESTClient(),
	}
}

// Get takes name of the instanceDeployment, and returns the corresponding instanceDeployment object, and an error if there is any.
func (c *instanceDeployments) Get(name string, options v1.GetOptions) (result *v1alpha1.InstanceDeployment, err error) {
	result = &v1alpha1.InstanceDeployment{}
	err = c.client.Get().
		Resource("instancedeployments").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of InstanceDeployments that match those selectors.
func (c *instanceDeployments) List(opts v1.ListOptions) (result *v1alpha1.InstanceDeploymentList, err error) {
	result = &v1alpha1.InstanceDeploymentList{}
	err = c.client.Get().
		Resource("instancedeployments").
		VersionedParams(&opts, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested instanceDeployments.
func (c *instanceDeployments) Watch(opts v1.ListOptions) (watch.Interface, error) {
	opts.Watch = true
	return c.client.Get().
		Resource("instancedeployments").
		VersionedParams(&opts, scheme.ParameterCodec).
		Watch()
}

// Create takes the representation of a instanceDeployment and creates it.  Returns the server's representation of the instanceDeployment, and an error, if there is any.
func (c *instanceDeployments) Create(instanceDeployment *v1alpha1.InstanceDeployment) (result *v1alpha1.InstanceDeployment, err error) {
	result = &v1alpha1.InstanceDeployment{}
	err = c.client.Post().
		Resource("instancedeployments").
		Body(instanceDeployment).
		Do().
		Into(result)
	return
}

// Update takes the representation of a instanceDeployment and updates it. Returns the server's representation of the instanceDeployment, and an error, if there is any.
func (c *instanceDeployments) Update(instanceDeployment *v1alpha1.InstanceDeployment) (result *v1alpha1.InstanceDeployment, err error) {
	result = &v1alpha1.InstanceDeployment{}
	err = c.client.Put().
		Resource("instancedeployments").
		Name(instanceDeployment.Name).
		Body(instanceDeployment).
		Do().
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().

func (c *instanceDeployments) UpdateStatus(instanceDeployment *v1alpha1.InstanceDeployment) (result *v1alpha1.InstanceDeployment, err error) {
	result = &v1alpha1.InstanceDeployment{}
	err = c.client.Put().
		Resource("instancedeployments").
		Name(instanceDeployment.Name).
		SubResource("status").
		Body(instanceDeployment).
		Do().
		Into(result)
	return
}

// Delete takes name of the instanceDeployment and deletes it. Returns an error if one occurs.
func (c *instanceDeployments) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Resource("instancedeployments").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *instanceDeployments) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	return c.client.Delete().
		Resource("instancedeployments").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched instanceDeployment.
func (c *instanceDeployments) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.InstanceDeployment, err error) {
	result = &v1alpha1.InstanceDeployment{}
	err = c.client.Patch(pt).
		Resource("instancedeployments").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}

// GetScale takes name of the instanceDeployment, and returns the corresponding v1alpha1.Scale object, and an error if there is any.
func (c *instanceDeployments) GetScale(instanceDeploymentName string, options v1.GetOptions) (result *v1alpha1.Scale, err error) {
	result = &v1alpha1.Scale{}
	err = c.client.Get().
		Resource("instancedeployments").
		Name(instanceDeploymentName).
		SubResource("scale").
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// UpdateScale takes the top resource name and the representation of a scale and updates it. Returns the server's representation of the scale, and an error, if there is any.
func (c *instanceDeployments) UpdateScale(instanceDeploymentName string, scale *v1alpha1.Scale) (result *v1alpha1.Scale, err error) {
	result = &v1alpha1.Scale{}
	err = c.client.Put().
		Resource("instancedeployments").
		Name(instanceDeploymentName).
		SubResource("scale").
		Body(scale).
		Do().
		Into(result)
	return
}