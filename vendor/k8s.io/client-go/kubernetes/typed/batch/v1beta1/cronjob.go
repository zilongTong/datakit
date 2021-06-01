/*
Copyright The Kubernetes Authors.

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

package v1beta1

import (
	"time"

	v1beta1 "k8s.io/api/batch/v1beta1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	scheme "k8s.io/client-go/kubernetes/scheme"
	rest "k8s.io/client-go/rest"
)

// CronJobsGetter has a method to return a CronJobInterface.
// A group's client should implement this interface.
type CronJobsGetter interface {
	CronJobs(namespace string) CronJobInterface
}

// CronJobInterface has methods to work with CronJob resources.
type CronJobInterface interface {
	Create(*v1beta1.CronJob) (*v1beta1.CronJob, error)
	Update(*v1beta1.CronJob) (*v1beta1.CronJob, error)
	UpdateStatus(*v1beta1.CronJob) (*v1beta1.CronJob, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1beta1.CronJob, error)
	List(opts v1.ListOptions) (*v1beta1.CronJobList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1beta1.CronJob, err error)
	CronJobExpansion
}

// cronJobs implements CronJobInterface
type cronJobs struct {
	client rest.Interface
	ns     string
}

// newCronJobs returns a CronJobs
func newCronJobs(c *BatchV1beta1Client, namespace string) *cronJobs {
	return &cronJobs{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the cronJob, and returns the corresponding cronJob object, and an error if there is any.
func (c *cronJobs) Get(name string, options v1.GetOptions) (result *v1beta1.CronJob, err error) {
	result = &v1beta1.CronJob{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("cronjobs").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of CronJobs that match those selectors.
func (c *cronJobs) List(opts v1.ListOptions) (result *v1beta1.CronJobList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1beta1.CronJobList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("cronjobs").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested cronJobs.
func (c *cronJobs) Watch(opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("cronjobs").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch()
}

// Create takes the representation of a cronJob and creates it.  Returns the server's representation of the cronJob, and an error, if there is any.
func (c *cronJobs) Create(cronJob *v1beta1.CronJob) (result *v1beta1.CronJob, err error) {
	result = &v1beta1.CronJob{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("cronjobs").
		Body(cronJob).
		Do().
		Into(result)
	return
}

// Update takes the representation of a cronJob and updates it. Returns the server's representation of the cronJob, and an error, if there is any.
func (c *cronJobs) Update(cronJob *v1beta1.CronJob) (result *v1beta1.CronJob, err error) {
	result = &v1beta1.CronJob{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("cronjobs").
		Name(cronJob.Name).
		Body(cronJob).
		Do().
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().

func (c *cronJobs) UpdateStatus(cronJob *v1beta1.CronJob) (result *v1beta1.CronJob, err error) {
	result = &v1beta1.CronJob{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("cronjobs").
		Name(cronJob.Name).
		SubResource("status").
		Body(cronJob).
		Do().
		Into(result)
	return
}

// Delete takes name of the cronJob and deletes it. Returns an error if one occurs.
func (c *cronJobs) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("cronjobs").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *cronJobs) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	var timeout time.Duration
	if listOptions.TimeoutSeconds != nil {
		timeout = time.Duration(*listOptions.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("cronjobs").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Timeout(timeout).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched cronJob.
func (c *cronJobs) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1beta1.CronJob, err error) {
	result = &v1beta1.CronJob{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("cronjobs").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
