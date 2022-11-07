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

// Code generated by applyconfiguration-gen. DO NOT EDIT.

package v1

// CephFSVolumeSourceApplyConfiguration represents an declarative configuration of the CephFSVolumeSource type for use
// with apply.
type CephFSVolumeSourceApplyConfiguration struct {
	Monitors   []string                                `json:"monitors,omitempty"`
	Path       *string                                 `json:"path,omitempty"`
	User       *string                                 `json:"user,omitempty"`
	SecretFile *string                                 `json:"secretFile,omitempty"`
	SecretRef  *LocalObjectReferenceApplyConfiguration `json:"secretRef,omitempty"`
	ReadOnly   *bool                                   `json:"readOnly,omitempty"`
}

// CephFSVolumeSourceApplyConfiguration constructs an declarative configuration of the CephFSVolumeSource type for use with
// apply.
func CephFSVolumeSource() *CephFSVolumeSourceApplyConfiguration {
	return &CephFSVolumeSourceApplyConfiguration{}
}

// WithMonitors adds the given value to the Monitors field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, values provided by each call will be appended to the Monitors field.
func (b *CephFSVolumeSourceApplyConfiguration) WithMonitors(values ...string) *CephFSVolumeSourceApplyConfiguration {
	for i := range values {
		b.Monitors = append(b.Monitors, values[i])
	}
	return b
}

// WithPath sets the Path field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Path field is set to the value of the last call.
func (b *CephFSVolumeSourceApplyConfiguration) WithPath(value string) *CephFSVolumeSourceApplyConfiguration {
	b.Path = &value
	return b
}

// WithUser sets the User field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the User field is set to the value of the last call.
func (b *CephFSVolumeSourceApplyConfiguration) WithUser(value string) *CephFSVolumeSourceApplyConfiguration {
	b.User = &value
	return b
}

// WithSecretFile sets the SecretFile field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the SecretFile field is set to the value of the last call.
func (b *CephFSVolumeSourceApplyConfiguration) WithSecretFile(value string) *CephFSVolumeSourceApplyConfiguration {
	b.SecretFile = &value
	return b
}

// WithSecretRef sets the SecretRef field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the SecretRef field is set to the value of the last call.
func (b *CephFSVolumeSourceApplyConfiguration) WithSecretRef(value *LocalObjectReferenceApplyConfiguration) *CephFSVolumeSourceApplyConfiguration {
	b.SecretRef = value
	return b
}

// WithReadOnly sets the ReadOnly field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the ReadOnly field is set to the value of the last call.
func (b *CephFSVolumeSourceApplyConfiguration) WithReadOnly(value bool) *CephFSVolumeSourceApplyConfiguration {
	b.ReadOnly = &value
	return b
}
