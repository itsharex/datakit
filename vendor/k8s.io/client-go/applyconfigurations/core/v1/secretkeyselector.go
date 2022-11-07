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

// SecretKeySelectorApplyConfiguration represents an declarative configuration of the SecretKeySelector type for use
// with apply.
type SecretKeySelectorApplyConfiguration struct {
	LocalObjectReferenceApplyConfiguration `json:",inline"`
	Key                                    *string `json:"key,omitempty"`
	Optional                               *bool   `json:"optional,omitempty"`
}

// SecretKeySelectorApplyConfiguration constructs an declarative configuration of the SecretKeySelector type for use with
// apply.
func SecretKeySelector() *SecretKeySelectorApplyConfiguration {
	return &SecretKeySelectorApplyConfiguration{}
}

// WithName sets the Name field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Name field is set to the value of the last call.
func (b *SecretKeySelectorApplyConfiguration) WithName(value string) *SecretKeySelectorApplyConfiguration {
	b.Name = &value
	return b
}

// WithKey sets the Key field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Key field is set to the value of the last call.
func (b *SecretKeySelectorApplyConfiguration) WithKey(value string) *SecretKeySelectorApplyConfiguration {
	b.Key = &value
	return b
}

// WithOptional sets the Optional field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Optional field is set to the value of the last call.
func (b *SecretKeySelectorApplyConfiguration) WithOptional(value bool) *SecretKeySelectorApplyConfiguration {
	b.Optional = &value
	return b
}
