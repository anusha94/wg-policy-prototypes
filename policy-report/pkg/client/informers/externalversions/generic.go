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

package externalversions

import (
	"fmt"

	schema "k8s.io/apimachinery/pkg/runtime/schema"
	cache "k8s.io/client-go/tools/cache"
	v1beta2 "sigs.k8s.io/wg-policy-prototypes/policy-report/api/reports.x-k8s.io/v1beta2"
	v1alpha1 "sigs.k8s.io/wg-policy-prototypes/policy-report/api/wgpolicyk8s.io/v1alpha1"
	v1alpha2 "sigs.k8s.io/wg-policy-prototypes/policy-report/api/wgpolicyk8s.io/v1alpha2"
	v1beta1 "sigs.k8s.io/wg-policy-prototypes/policy-report/api/wgpolicyk8s.io/v1beta1"
)

// GenericInformer is type of SharedIndexInformer which will locate and delegate to other
// sharedInformers based on type
type GenericInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() cache.GenericLister
}

type genericInformer struct {
	informer cache.SharedIndexInformer
	resource schema.GroupResource
}

// Informer returns the SharedIndexInformer.
func (f *genericInformer) Informer() cache.SharedIndexInformer {
	return f.informer
}

// Lister returns the GenericLister.
func (f *genericInformer) Lister() cache.GenericLister {
	return cache.NewGenericLister(f.Informer().GetIndexer(), f.resource)
}

// ForResource gives generic access to a shared informer of the matching type
// TODO extend this to unknown resources with a client pool
func (f *sharedInformerFactory) ForResource(resource schema.GroupVersionResource) (GenericInformer, error) {
	switch resource {
	// Group=reports.x-k8s.io, Version=v1beta2
	case v1beta2.SchemeGroupVersion.WithResource("clusterpolicyreports"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Reports().V1beta2().ClusterPolicyReports().Informer()}, nil
	case v1beta2.SchemeGroupVersion.WithResource("policyreports"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Reports().V1beta2().PolicyReports().Informer()}, nil

		// Group=wgpolicyk8s.io, Version=v1alpha1
	case v1alpha1.SchemeGroupVersion.WithResource("clusterpolicyreports"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Wgpolicyk8s().V1alpha1().ClusterPolicyReports().Informer()}, nil
	case v1alpha1.SchemeGroupVersion.WithResource("policyreports"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Wgpolicyk8s().V1alpha1().PolicyReports().Informer()}, nil

		// Group=wgpolicyk8s.io, Version=v1alpha2
	case v1alpha2.SchemeGroupVersion.WithResource("clusterpolicyreports"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Wgpolicyk8s().V1alpha2().ClusterPolicyReports().Informer()}, nil
	case v1alpha2.SchemeGroupVersion.WithResource("policyreports"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Wgpolicyk8s().V1alpha2().PolicyReports().Informer()}, nil

		// Group=wgpolicyk8s.io, Version=v1beta1
	case v1beta1.SchemeGroupVersion.WithResource("clusterpolicyreports"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Wgpolicyk8s().V1beta1().ClusterPolicyReports().Informer()}, nil
	case v1beta1.SchemeGroupVersion.WithResource("policyreports"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Wgpolicyk8s().V1beta1().PolicyReports().Informer()}, nil

	}

	return nil, fmt.Errorf("no informer found for %v", resource)
}
