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

// KeyToPathApplyConfiguration represents an declarative configuration of the KeyToPath type for use
// with apply.
type KeyToPathApplyConfiguration struct {
	Key  *string `json:"key,omitempty"`
	Path *string `json:"path,omitempty"`
	Mode *int32  `json:"mode,omitempty"`
}

// KeyToPathApplyConfiguration constructs an declarative configuration of the KeyToPath type for use with
// apply.
func KeyToPath() *KeyToPathApplyConfiguration {
	return &KeyToPathApplyConfiguration{}
}

// WithKey sets the Key field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Key field is set to the value of the last call.
func (b *KeyToPathApplyConfiguration) WithKey(value string) *KeyToPathApplyConfiguration {
	b.Key = &value
	return b
}

// WithPath sets the Path field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Path field is set to the value of the last call.
func (b *KeyToPathApplyConfiguration) WithPath(value string) *KeyToPathApplyConfiguration {
	b.Path = &value
	return b
}

// WithMode sets the Mode field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Mode field is set to the value of the last call.
func (b *KeyToPathApplyConfiguration) WithMode(value int32) *KeyToPathApplyConfiguration {
	b.Mode = &value
	return b
}
