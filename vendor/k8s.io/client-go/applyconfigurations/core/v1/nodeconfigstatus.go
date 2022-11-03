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

// NodeConfigStatusApplyConfiguration represents an declarative configuration of the NodeConfigStatus type for use
// with apply.
type NodeConfigStatusApplyConfiguration struct {
	Assigned      *NodeConfigSourceApplyConfiguration `json:"assigned,omitempty"`
	Active        *NodeConfigSourceApplyConfiguration `json:"active,omitempty"`
	LastKnownGood *NodeConfigSourceApplyConfiguration `json:"lastKnownGood,omitempty"`
	Error         *string                             `json:"error,omitempty"`
}

// NodeConfigStatusApplyConfiguration constructs an declarative configuration of the NodeConfigStatus type for use with
// apply.
func NodeConfigStatus() *NodeConfigStatusApplyConfiguration {
	return &NodeConfigStatusApplyConfiguration{}
}

// WithAssigned sets the Assigned field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Assigned field is set to the value of the last call.
func (b *NodeConfigStatusApplyConfiguration) WithAssigned(value *NodeConfigSourceApplyConfiguration) *NodeConfigStatusApplyConfiguration {
	b.Assigned = value
	return b
}

// WithActive sets the Active field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Active field is set to the value of the last call.
func (b *NodeConfigStatusApplyConfiguration) WithActive(value *NodeConfigSourceApplyConfiguration) *NodeConfigStatusApplyConfiguration {
	b.Active = value
	return b
}

// WithLastKnownGood sets the LastKnownGood field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the LastKnownGood field is set to the value of the last call.
func (b *NodeConfigStatusApplyConfiguration) WithLastKnownGood(value *NodeConfigSourceApplyConfiguration) *NodeConfigStatusApplyConfiguration {
	b.LastKnownGood = value
	return b
}

// WithError sets the Error field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Error field is set to the value of the last call.
func (b *NodeConfigStatusApplyConfiguration) WithError(value string) *NodeConfigStatusApplyConfiguration {
	b.Error = &value
	return b
}
