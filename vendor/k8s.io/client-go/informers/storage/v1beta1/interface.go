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

package v1beta1

import (
	internalinterfaces "k8s.io/client-go/informers/internalinterfaces"
)

// Interface provides access to all the informers in this group version.
type Interface interface {
	// CSIDrivers returns a CSIDriverInformer.
	CSIDrivers() CSIDriverInformer
	// CSINodes returns a CSINodeInformer.
	CSINodes() CSINodeInformer
	// CSIStorageCapacities returns a CSIStorageCapacityInformer.
	CSIStorageCapacities() CSIStorageCapacityInformer
	// StorageClasses returns a StorageClassInformer.
	StorageClasses() StorageClassInformer
	// VolumeAttachments returns a VolumeAttachmentInformer.
	VolumeAttachments() VolumeAttachmentInformer
}

type version struct {
	factory          internalinterfaces.SharedInformerFactory
	namespace        string
	tweakListOptions internalinterfaces.TweakListOptionsFunc
}

// New returns a new Interface.
func New(f internalinterfaces.SharedInformerFactory, namespace string, tweakListOptions internalinterfaces.TweakListOptionsFunc) Interface {
	return &version{factory: f, namespace: namespace, tweakListOptions: tweakListOptions}
}

// CSIDrivers returns a CSIDriverInformer.
func (v *version) CSIDrivers() CSIDriverInformer {
	return &cSIDriverInformer{factory: v.factory, tweakListOptions: v.tweakListOptions}
}

// CSINodes returns a CSINodeInformer.
func (v *version) CSINodes() CSINodeInformer {
	return &cSINodeInformer{factory: v.factory, tweakListOptions: v.tweakListOptions}
}

// CSIStorageCapacities returns a CSIStorageCapacityInformer.
func (v *version) CSIStorageCapacities() CSIStorageCapacityInformer {
	return &cSIStorageCapacityInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// StorageClasses returns a StorageClassInformer.
func (v *version) StorageClasses() StorageClassInformer {
	return &storageClassInformer{factory: v.factory, tweakListOptions: v.tweakListOptions}
}

// VolumeAttachments returns a VolumeAttachmentInformer.
func (v *version) VolumeAttachments() VolumeAttachmentInformer {
	return &volumeAttachmentInformer{factory: v.factory, tweakListOptions: v.tweakListOptions}
}
