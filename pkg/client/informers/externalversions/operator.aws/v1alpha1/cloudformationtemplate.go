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

// Code generated by informer-gen. DO NOT EDIT.

package v1alpha1

import (
	time "time"

	operator_aws_v1alpha1 "github.com/christopherhein/aws-operator/pkg/apis/operator.aws/v1alpha1"
	versioned "github.com/christopherhein/aws-operator/pkg/client/clientset/versioned"
	internalinterfaces "github.com/christopherhein/aws-operator/pkg/client/informers/externalversions/internalinterfaces"
	v1alpha1 "github.com/christopherhein/aws-operator/pkg/client/listers/operator.aws/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	cache "k8s.io/client-go/tools/cache"
)

// CloudFormationTemplateInformer provides access to a shared informer and lister for
// CloudFormationTemplates.
type CloudFormationTemplateInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() v1alpha1.CloudFormationTemplateLister
}

type cloudFormationTemplateInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
	namespace        string
}

// NewCloudFormationTemplateInformer constructs a new informer for CloudFormationTemplate type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewCloudFormationTemplateInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredCloudFormationTemplateInformer(client, namespace, resyncPeriod, indexers, nil)
}

// NewFilteredCloudFormationTemplateInformer constructs a new informer for CloudFormationTemplate type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredCloudFormationTemplateInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options v1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.OperatorV1alpha1().CloudFormationTemplates(namespace).List(options)
			},
			WatchFunc: func(options v1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.OperatorV1alpha1().CloudFormationTemplates(namespace).Watch(options)
			},
		},
		&operator_aws_v1alpha1.CloudFormationTemplate{},
		resyncPeriod,
		indexers,
	)
}

func (f *cloudFormationTemplateInformer) defaultInformer(client versioned.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredCloudFormationTemplateInformer(client, f.namespace, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *cloudFormationTemplateInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&operator_aws_v1alpha1.CloudFormationTemplate{}, f.defaultInformer)
}

func (f *cloudFormationTemplateInformer) Lister() v1alpha1.CloudFormationTemplateLister {
	return v1alpha1.NewCloudFormationTemplateLister(f.Informer().GetIndexer())
}
