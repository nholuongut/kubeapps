/*
Copyright 2018 Bitnami.

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

	apprepository_v1alpha1 "github.com/kubeapps/kubeapps/cmd/apprepository-controller/pkg/apis/apprepository/v1alpha1"
	versioned "github.com/kubeapps/kubeapps/cmd/apprepository-controller/pkg/client/clientset/versioned"
	internalinterfaces "github.com/kubeapps/kubeapps/cmd/apprepository-controller/pkg/client/informers/externalversions/internalinterfaces"
	v1alpha1 "github.com/kubeapps/kubeapps/cmd/apprepository-controller/pkg/client/listers/apprepository/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	cache "k8s.io/client-go/tools/cache"
)

// AppRepositoryInformer provides access to a shared informer and lister for
// AppRepositories.
type AppRepositoryInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() v1alpha1.AppRepositoryLister
}

type appRepositoryInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
	namespace        string
}

// NewAppRepositoryInformer constructs a new informer for AppRepository type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewAppRepositoryInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredAppRepositoryInformer(client, namespace, resyncPeriod, indexers, nil)
}

// NewFilteredAppRepositoryInformer constructs a new informer for AppRepository type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredAppRepositoryInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options v1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.kubeappsV1alpha1().AppRepositories(namespace).List(options)
			},
			WatchFunc: func(options v1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.kubeappsV1alpha1().AppRepositories(namespace).Watch(options)
			},
		},
		&apprepository_v1alpha1.AppRepository{},
		resyncPeriod,
		indexers,
	)
}

func (f *appRepositoryInformer) defaultInformer(client versioned.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredAppRepositoryInformer(client, f.namespace, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *appRepositoryInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&apprepository_v1alpha1.AppRepository{}, f.defaultInformer)
}

func (f *appRepositoryInformer) Lister() v1alpha1.AppRepositoryLister {
	return v1alpha1.NewAppRepositoryLister(f.Informer().GetIndexer())
}
