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
// Code generated by informer-gen. DO NOT EDIT.

package v1beta2

import (
	"context"
	time "time"

	versioned "github.com/kubernetes-sigs/wg-policy-prototypes/policy-report/pkg/client/clientset/versioned"
	internalinterfaces "github.com/kubernetes-sigs/wg-policy-prototypes/policy-report/pkg/client/informers/externalversions/internalinterfaces"
	v1beta2 "github.com/kubernetes-sigs/wg-policy-prototypes/policy-report/pkg/client/listers/reports.x-k8s.io/v1beta2"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	cache "k8s.io/client-go/tools/cache"
	reportsxk8siov1beta2 "sigs.k8s.io/wg-policy-prototypes/policy-report/api/reports.x-k8s.io/v1beta2"
)

// PolicyReportInformer provides access to a shared informer and lister for
// PolicyReports.
type PolicyReportInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() v1beta2.PolicyReportLister
}

type policyReportInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
	namespace        string
}

// NewPolicyReportInformer constructs a new informer for PolicyReport type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewPolicyReportInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredPolicyReportInformer(client, namespace, resyncPeriod, indexers, nil)
}

// NewFilteredPolicyReportInformer constructs a new informer for PolicyReport type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredPolicyReportInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options v1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.ReportsV1beta2().PolicyReports(namespace).List(context.TODO(), options)
			},
			WatchFunc: func(options v1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.ReportsV1beta2().PolicyReports(namespace).Watch(context.TODO(), options)
			},
		},
		&reportsxk8siov1beta2.PolicyReport{},
		resyncPeriod,
		indexers,
	)
}

func (f *policyReportInformer) defaultInformer(client versioned.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredPolicyReportInformer(client, f.namespace, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *policyReportInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&reportsxk8siov1beta2.PolicyReport{}, f.defaultInformer)
}

func (f *policyReportInformer) Lister() v1beta2.PolicyReportLister {
	return v1beta2.NewPolicyReportLister(f.Informer().GetIndexer())
}
