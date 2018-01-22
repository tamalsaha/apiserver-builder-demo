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

// This file was automatically generated by informer-gen

package internalversion

import (
	insect "github.com/tamalsaha/apiserver-builder-demo/pkg/apis/insect"
	internalclientset "github.com/tamalsaha/apiserver-builder-demo/pkg/client/clientset_generated/internalclientset"
	internalinterfaces "github.com/tamalsaha/apiserver-builder-demo/pkg/client/informers_generated/internalversion/internalinterfaces"
	internalversion "github.com/tamalsaha/apiserver-builder-demo/pkg/client/listers_generated/insect/internalversion"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	cache "k8s.io/client-go/tools/cache"
	time "time"
)

// BeeInformer provides access to a shared informer and lister for
// Bees.
type BeeInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() internalversion.BeeLister
}

type beeInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
	namespace        string
}

// NewBeeInformer constructs a new informer for Bee type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewBeeInformer(client internalclientset.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredBeeInformer(client, namespace, resyncPeriod, indexers, nil)
}

// NewFilteredBeeInformer constructs a new informer for Bee type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredBeeInformer(client internalclientset.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options v1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.Insect().Bees(namespace).List(options)
			},
			WatchFunc: func(options v1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.Insect().Bees(namespace).Watch(options)
			},
		},
		&insect.Bee{},
		resyncPeriod,
		indexers,
	)
}

func (f *beeInformer) defaultInformer(client internalclientset.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredBeeInformer(client, f.namespace, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *beeInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&insect.Bee{}, f.defaultInformer)
}

func (f *beeInformer) Lister() internalversion.BeeLister {
	return internalversion.NewBeeLister(f.Informer().GetIndexer())
}