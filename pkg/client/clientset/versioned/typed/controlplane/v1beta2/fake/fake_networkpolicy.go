// Copyright 2024 Antrea Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	"context"

	v1beta2 "antrea.io/antrea/pkg/apis/controlplane/v1beta2"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeNetworkPolicies implements NetworkPolicyInterface
type FakeNetworkPolicies struct {
	Fake *FakeControlplaneV1beta2
}

var networkpoliciesResource = v1beta2.SchemeGroupVersion.WithResource("networkpolicies")

var networkpoliciesKind = v1beta2.SchemeGroupVersion.WithKind("NetworkPolicy")

// Get takes name of the networkPolicy, and returns the corresponding networkPolicy object, and an error if there is any.
func (c *FakeNetworkPolicies) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1beta2.NetworkPolicy, err error) {
	emptyResult := &v1beta2.NetworkPolicy{}
	obj, err := c.Fake.
		Invokes(testing.NewRootGetActionWithOptions(networkpoliciesResource, name, options), emptyResult)
	if obj == nil {
		return emptyResult, err
	}
	return obj.(*v1beta2.NetworkPolicy), err
}

// List takes label and field selectors, and returns the list of NetworkPolicies that match those selectors.
func (c *FakeNetworkPolicies) List(ctx context.Context, opts v1.ListOptions) (result *v1beta2.NetworkPolicyList, err error) {
	emptyResult := &v1beta2.NetworkPolicyList{}
	obj, err := c.Fake.
		Invokes(testing.NewRootListActionWithOptions(networkpoliciesResource, networkpoliciesKind, opts), emptyResult)
	if obj == nil {
		return emptyResult, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1beta2.NetworkPolicyList{ListMeta: obj.(*v1beta2.NetworkPolicyList).ListMeta}
	for _, item := range obj.(*v1beta2.NetworkPolicyList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested networkPolicies.
func (c *FakeNetworkPolicies) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewRootWatchActionWithOptions(networkpoliciesResource, opts))
}
