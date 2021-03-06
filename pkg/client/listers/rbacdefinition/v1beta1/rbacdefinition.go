/*
Copyright 2019 The Openshift Evangelists

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

package v1beta1

import (
	v1beta1 "github.com/darkrat/cajadepandora/pkg/apis/RBACDefinition/v1beta1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// RBACDefinitionLister helps list RBACDefinitions.
type RBACDefinitionLister interface {
	// List lists all RBACDefinitions in the indexer.
	List(selector labels.Selector) (ret []*v1beta1.RBACDefinition, err error)
	// RBACDefinitions returns an object that can list and get RBACDefinitions.
	RBACDefinitions(namespace string) RBACDefinitionNamespaceLister
	RBACDefinitionListerExpansion
}

// rBACDefinitionLister implements the RBACDefinitionLister interface.
type rBACDefinitionLister struct {
	indexer cache.Indexer
}

// NewRBACDefinitionLister returns a new RBACDefinitionLister.
func NewRBACDefinitionLister(indexer cache.Indexer) RBACDefinitionLister {
	return &rBACDefinitionLister{indexer: indexer}
}

// List lists all RBACDefinitions in the indexer.
func (s *rBACDefinitionLister) List(selector labels.Selector) (ret []*v1beta1.RBACDefinition, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1beta1.RBACDefinition))
	})
	return ret, err
}

// RBACDefinitions returns an object that can list and get RBACDefinitions.
func (s *rBACDefinitionLister) RBACDefinitions(namespace string) RBACDefinitionNamespaceLister {
	return rBACDefinitionNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// RBACDefinitionNamespaceLister helps list and get RBACDefinitions.
type RBACDefinitionNamespaceLister interface {
	// List lists all RBACDefinitions in the indexer for a given namespace.
	List(selector labels.Selector) (ret []*v1beta1.RBACDefinition, err error)
	// Get retrieves the RBACDefinition from the indexer for a given namespace and name.
	Get(name string) (*v1beta1.RBACDefinition, error)
	RBACDefinitionNamespaceListerExpansion
}

// rBACDefinitionNamespaceLister implements the RBACDefinitionNamespaceLister
// interface.
type rBACDefinitionNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all RBACDefinitions in the indexer for a given namespace.
func (s rBACDefinitionNamespaceLister) List(selector labels.Selector) (ret []*v1beta1.RBACDefinition, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1beta1.RBACDefinition))
	})
	return ret, err
}

// Get retrieves the RBACDefinition from the indexer for a given namespace and name.
func (s rBACDefinitionNamespaceLister) Get(name string) (*v1beta1.RBACDefinition, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1beta1.Resource("rbacdefinition"), name)
	}
	return obj.(*v1beta1.RBACDefinition), nil
}
