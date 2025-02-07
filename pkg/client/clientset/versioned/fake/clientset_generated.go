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
	clientset "antrea.io/antrea/pkg/client/clientset/versioned"
	controlplanev1beta2 "antrea.io/antrea/pkg/client/clientset/versioned/typed/controlplane/v1beta2"
	fakecontrolplanev1beta2 "antrea.io/antrea/pkg/client/clientset/versioned/typed/controlplane/v1beta2/fake"
	crdv1alpha1 "antrea.io/antrea/pkg/client/clientset/versioned/typed/crd/v1alpha1"
	fakecrdv1alpha1 "antrea.io/antrea/pkg/client/clientset/versioned/typed/crd/v1alpha1/fake"
	crdv1alpha2 "antrea.io/antrea/pkg/client/clientset/versioned/typed/crd/v1alpha2"
	fakecrdv1alpha2 "antrea.io/antrea/pkg/client/clientset/versioned/typed/crd/v1alpha2/fake"
	crdv1beta1 "antrea.io/antrea/pkg/client/clientset/versioned/typed/crd/v1beta1"
	fakecrdv1beta1 "antrea.io/antrea/pkg/client/clientset/versioned/typed/crd/v1beta1/fake"
	statsv1alpha1 "antrea.io/antrea/pkg/client/clientset/versioned/typed/stats/v1alpha1"
	fakestatsv1alpha1 "antrea.io/antrea/pkg/client/clientset/versioned/typed/stats/v1alpha1/fake"
	systemv1beta1 "antrea.io/antrea/pkg/client/clientset/versioned/typed/system/v1beta1"
	fakesystemv1beta1 "antrea.io/antrea/pkg/client/clientset/versioned/typed/system/v1beta1/fake"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/watch"
	"k8s.io/client-go/discovery"
	fakediscovery "k8s.io/client-go/discovery/fake"
	"k8s.io/client-go/testing"
)

// NewSimpleClientset returns a clientset that will respond with the provided objects.
// It's backed by a very simple object tracker that processes creates, updates and deletions as-is,
// without applying any field management, validations and/or defaults. It shouldn't be considered a replacement
// for a real clientset and is mostly useful in simple unit tests.
//
// DEPRECATED: NewClientset replaces this with support for field management, which significantly improves
// server side apply testing. NewClientset is only available when apply configurations are generated (e.g.
// via --with-applyconfig).
func NewSimpleClientset(objects ...runtime.Object) *Clientset {
	o := testing.NewObjectTracker(scheme, codecs.UniversalDecoder())
	for _, obj := range objects {
		if err := o.Add(obj); err != nil {
			panic(err)
		}
	}

	cs := &Clientset{tracker: o}
	cs.discovery = &fakediscovery.FakeDiscovery{Fake: &cs.Fake}
	cs.AddReactor("*", "*", testing.ObjectReaction(o))
	cs.AddWatchReactor("*", func(action testing.Action) (handled bool, ret watch.Interface, err error) {
		gvr := action.GetResource()
		ns := action.GetNamespace()
		watch, err := o.Watch(gvr, ns)
		if err != nil {
			return false, nil, err
		}
		return true, watch, nil
	})

	return cs
}

// Clientset implements clientset.Interface. Meant to be embedded into a
// struct to get a default implementation. This makes faking out just the method
// you want to test easier.
type Clientset struct {
	testing.Fake
	discovery *fakediscovery.FakeDiscovery
	tracker   testing.ObjectTracker
}

func (c *Clientset) Discovery() discovery.DiscoveryInterface {
	return c.discovery
}

func (c *Clientset) Tracker() testing.ObjectTracker {
	return c.tracker
}

var (
	_ clientset.Interface = &Clientset{}
	_ testing.FakeClient  = &Clientset{}
)

// ControlplaneV1beta2 retrieves the ControlplaneV1beta2Client
func (c *Clientset) ControlplaneV1beta2() controlplanev1beta2.ControlplaneV1beta2Interface {
	return &fakecontrolplanev1beta2.FakeControlplaneV1beta2{Fake: &c.Fake}
}

// CrdV1alpha1 retrieves the CrdV1alpha1Client
func (c *Clientset) CrdV1alpha1() crdv1alpha1.CrdV1alpha1Interface {
	return &fakecrdv1alpha1.FakeCrdV1alpha1{Fake: &c.Fake}
}

// CrdV1alpha2 retrieves the CrdV1alpha2Client
func (c *Clientset) CrdV1alpha2() crdv1alpha2.CrdV1alpha2Interface {
	return &fakecrdv1alpha2.FakeCrdV1alpha2{Fake: &c.Fake}
}

// CrdV1beta1 retrieves the CrdV1beta1Client
func (c *Clientset) CrdV1beta1() crdv1beta1.CrdV1beta1Interface {
	return &fakecrdv1beta1.FakeCrdV1beta1{Fake: &c.Fake}
}

// StatsV1alpha1 retrieves the StatsV1alpha1Client
func (c *Clientset) StatsV1alpha1() statsv1alpha1.StatsV1alpha1Interface {
	return &fakestatsv1alpha1.FakeStatsV1alpha1{Fake: &c.Fake}
}

// SystemV1beta1 retrieves the SystemV1beta1Client
func (c *Clientset) SystemV1beta1() systemv1beta1.SystemV1beta1Interface {
	return &fakesystemv1beta1.FakeSystemV1beta1{Fake: &c.Fake}
}
