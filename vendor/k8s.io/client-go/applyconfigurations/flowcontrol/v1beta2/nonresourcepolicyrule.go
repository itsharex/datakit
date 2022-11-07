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

package v1beta2

// NonResourcePolicyRuleApplyConfiguration represents an declarative configuration of the NonResourcePolicyRule type for use
// with apply.
type NonResourcePolicyRuleApplyConfiguration struct {
	Verbs           []string `json:"verbs,omitempty"`
	NonResourceURLs []string `json:"nonResourceURLs,omitempty"`
}

// NonResourcePolicyRuleApplyConfiguration constructs an declarative configuration of the NonResourcePolicyRule type for use with
// apply.
func NonResourcePolicyRule() *NonResourcePolicyRuleApplyConfiguration {
	return &NonResourcePolicyRuleApplyConfiguration{}
}

// WithVerbs adds the given value to the Verbs field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, values provided by each call will be appended to the Verbs field.
func (b *NonResourcePolicyRuleApplyConfiguration) WithVerbs(values ...string) *NonResourcePolicyRuleApplyConfiguration {
	for i := range values {
		b.Verbs = append(b.Verbs, values[i])
	}
	return b
}

// WithNonResourceURLs adds the given value to the NonResourceURLs field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, values provided by each call will be appended to the NonResourceURLs field.
func (b *NonResourcePolicyRuleApplyConfiguration) WithNonResourceURLs(values ...string) *NonResourcePolicyRuleApplyConfiguration {
	for i := range values {
		b.NonResourceURLs = append(b.NonResourceURLs, values[i])
	}
	return b
}
