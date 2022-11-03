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

// NFSVolumeSourceApplyConfiguration represents an declarative configuration of the NFSVolumeSource type for use
// with apply.
type NFSVolumeSourceApplyConfiguration struct {
	Server   *string `json:"server,omitempty"`
	Path     *string `json:"path,omitempty"`
	ReadOnly *bool   `json:"readOnly,omitempty"`
}

// NFSVolumeSourceApplyConfiguration constructs an declarative configuration of the NFSVolumeSource type for use with
// apply.
func NFSVolumeSource() *NFSVolumeSourceApplyConfiguration {
	return &NFSVolumeSourceApplyConfiguration{}
}

// WithServer sets the Server field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Server field is set to the value of the last call.
func (b *NFSVolumeSourceApplyConfiguration) WithServer(value string) *NFSVolumeSourceApplyConfiguration {
	b.Server = &value
	return b
}

// WithPath sets the Path field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Path field is set to the value of the last call.
func (b *NFSVolumeSourceApplyConfiguration) WithPath(value string) *NFSVolumeSourceApplyConfiguration {
	b.Path = &value
	return b
}

// WithReadOnly sets the ReadOnly field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the ReadOnly field is set to the value of the last call.
func (b *NFSVolumeSourceApplyConfiguration) WithReadOnly(value bool) *NFSVolumeSourceApplyConfiguration {
	b.ReadOnly = &value
	return b
}
