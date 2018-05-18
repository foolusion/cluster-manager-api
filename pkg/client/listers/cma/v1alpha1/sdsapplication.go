/*
Copyright 2018 Samsung SDS Cloud Native Computing Team.

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

// Code generated by lister-gen. DO NOT EDIT.

package v1alpha1

import (
	v1alpha1 "github.com/samsung-cnct/cluster-manager-api/pkg/apis/cma/v1alpha1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// SDSApplicationLister helps list SDSApplications.
type SDSApplicationLister interface {
	// List lists all SDSApplications in the indexer.
	List(selector labels.Selector) (ret []*v1alpha1.SDSApplication, err error)
	// SDSApplications returns an object that can list and get SDSApplications.
	SDSApplications(namespace string) SDSApplicationNamespaceLister
	SDSApplicationListerExpansion
}

// sDSApplicationLister implements the SDSApplicationLister interface.
type sDSApplicationLister struct {
	indexer cache.Indexer
}

// NewSDSApplicationLister returns a new SDSApplicationLister.
func NewSDSApplicationLister(indexer cache.Indexer) SDSApplicationLister {
	return &sDSApplicationLister{indexer: indexer}
}

// List lists all SDSApplications in the indexer.
func (s *sDSApplicationLister) List(selector labels.Selector) (ret []*v1alpha1.SDSApplication, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.SDSApplication))
	})
	return ret, err
}

// SDSApplications returns an object that can list and get SDSApplications.
func (s *sDSApplicationLister) SDSApplications(namespace string) SDSApplicationNamespaceLister {
	return sDSApplicationNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// SDSApplicationNamespaceLister helps list and get SDSApplications.
type SDSApplicationNamespaceLister interface {
	// List lists all SDSApplications in the indexer for a given namespace.
	List(selector labels.Selector) (ret []*v1alpha1.SDSApplication, err error)
	// Get retrieves the SDSApplication from the indexer for a given namespace and name.
	Get(name string) (*v1alpha1.SDSApplication, error)
	SDSApplicationNamespaceListerExpansion
}

// sDSApplicationNamespaceLister implements the SDSApplicationNamespaceLister
// interface.
type sDSApplicationNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all SDSApplications in the indexer for a given namespace.
func (s sDSApplicationNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.SDSApplication, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.SDSApplication))
	})
	return ret, err
}

// Get retrieves the SDSApplication from the indexer for a given namespace and name.
func (s sDSApplicationNamespaceLister) Get(name string) (*v1alpha1.SDSApplication, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("sdsapplication"), name)
	}
	return obj.(*v1alpha1.SDSApplication), nil
}
