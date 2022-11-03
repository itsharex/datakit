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

// LifecycleHandlerApplyConfiguration represents an declarative configuration of the LifecycleHandler type for use
// with apply.
type LifecycleHandlerApplyConfiguration struct {
	Exec      *ExecActionApplyConfiguration      `json:"exec,omitempty"`
	HTTPGet   *HTTPGetActionApplyConfiguration   `json:"httpGet,omitempty"`
	TCPSocket *TCPSocketActionApplyConfiguration `json:"tcpSocket,omitempty"`
}

// LifecycleHandlerApplyConfiguration constructs an declarative configuration of the LifecycleHandler type for use with
// apply.
func LifecycleHandler() *LifecycleHandlerApplyConfiguration {
	return &LifecycleHandlerApplyConfiguration{}
}

// WithExec sets the Exec field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Exec field is set to the value of the last call.
func (b *LifecycleHandlerApplyConfiguration) WithExec(value *ExecActionApplyConfiguration) *LifecycleHandlerApplyConfiguration {
	b.Exec = value
	return b
}

// WithHTTPGet sets the HTTPGet field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the HTTPGet field is set to the value of the last call.
func (b *LifecycleHandlerApplyConfiguration) WithHTTPGet(value *HTTPGetActionApplyConfiguration) *LifecycleHandlerApplyConfiguration {
	b.HTTPGet = value
	return b
}

// WithTCPSocket sets the TCPSocket field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the TCPSocket field is set to the value of the last call.
func (b *LifecycleHandlerApplyConfiguration) WithTCPSocket(value *TCPSocketActionApplyConfiguration) *LifecycleHandlerApplyConfiguration {
	b.TCPSocket = value
	return b
}
