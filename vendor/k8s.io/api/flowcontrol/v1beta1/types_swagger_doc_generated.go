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

package v1beta1

// This file contains a collection of methods that can be used from go-restful to
// generate Swagger API documentation for its models. Please read this PR for more
// information on the implementation: https://github.com/emicklei/go-restful/pull/215
//
// TODOs are ignored from the parser (e.g. TODO(andronat):... || TODO:...) if and only if
// they are on one line! For multiple line or blocks that you want to ignore use ---.
// Any context after a --- is ignored.
//
// Those methods can be generated by using hack/update-generated-swagger-docs.sh

// AUTO-GENERATED FUNCTIONS START HERE. DO NOT EDIT.
var map_FlowDistinguisherMethod = map[string]string{
	"":     "FlowDistinguisherMethod specifies the method of a flow distinguisher.",
	"type": "`type` is the type of flow distinguisher method The supported types are \"ByUser\" and \"ByNamespace\". Required.",
}

func (FlowDistinguisherMethod) SwaggerDoc() map[string]string {
	return map_FlowDistinguisherMethod
}

var map_FlowSchema = map[string]string{
	"":         "FlowSchema defines the schema of a group of flows. Note that a flow is made up of a set of inbound API requests with similar attributes and is identified by a pair of strings: the name of the FlowSchema and a \"flow distinguisher\".",
	"metadata": "`metadata` is the standard object's metadata. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata",
	"spec":     "`spec` is the specification of the desired behavior of a FlowSchema. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#spec-and-status",
	"status":   "`status` is the current status of a FlowSchema. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#spec-and-status",
}

func (FlowSchema) SwaggerDoc() map[string]string {
	return map_FlowSchema
}

var map_FlowSchemaCondition = map[string]string{
	"":                   "FlowSchemaCondition describes conditions for a FlowSchema.",
	"type":               "`type` is the type of the condition. Required.",
	"status":             "`status` is the status of the condition. Can be True, False, Unknown. Required.",
	"lastTransitionTime": "`lastTransitionTime` is the last time the condition transitioned from one status to another.",
	"reason":             "`reason` is a unique, one-word, CamelCase reason for the condition's last transition.",
	"message":            "`message` is a human-readable message indicating details about last transition.",
}

func (FlowSchemaCondition) SwaggerDoc() map[string]string {
	return map_FlowSchemaCondition
}

var map_FlowSchemaList = map[string]string{
	"":         "FlowSchemaList is a list of FlowSchema objects.",
	"metadata": "`metadata` is the standard list metadata. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata",
	"items":    "`items` is a list of FlowSchemas.",
}

func (FlowSchemaList) SwaggerDoc() map[string]string {
	return map_FlowSchemaList
}

var map_FlowSchemaSpec = map[string]string{
	"":                           "FlowSchemaSpec describes how the FlowSchema's specification looks like.",
	"priorityLevelConfiguration": "`priorityLevelConfiguration` should reference a PriorityLevelConfiguration in the cluster. If the reference cannot be resolved, the FlowSchema will be ignored and marked as invalid in its status. Required.",
	"matchingPrecedence":         "`matchingPrecedence` is used to choose among the FlowSchemas that match a given request. The chosen FlowSchema is among those with the numerically lowest (which we take to be logically highest) MatchingPrecedence.  Each MatchingPrecedence value must be ranged in [1,10000]. Note that if the precedence is not specified, it will be set to 1000 as default.",
	"distinguisherMethod":        "`distinguisherMethod` defines how to compute the flow distinguisher for requests that match this schema. `nil` specifies that the distinguisher is disabled and thus will always be the empty string.",
	"rules":                      "`rules` describes which requests will match this flow schema. This FlowSchema matches a request if and only if at least one member of rules matches the request. if it is an empty slice, there will be no requests matching the FlowSchema.",
}

func (FlowSchemaSpec) SwaggerDoc() map[string]string {
	return map_FlowSchemaSpec
}

var map_FlowSchemaStatus = map[string]string{
	"":           "FlowSchemaStatus represents the current state of a FlowSchema.",
	"conditions": "`conditions` is a list of the current states of FlowSchema.",
}

func (FlowSchemaStatus) SwaggerDoc() map[string]string {
	return map_FlowSchemaStatus
}

var map_GroupSubject = map[string]string{
	"":     "GroupSubject holds detailed information for group-kind subject.",
	"name": "name is the user group that matches, or \"*\" to match all user groups. See https://github.com/kubernetes/apiserver/blob/master/pkg/authentication/user/user.go for some well-known group names. Required.",
}

func (GroupSubject) SwaggerDoc() map[string]string {
	return map_GroupSubject
}

var map_LimitResponse = map[string]string{
	"":        "LimitResponse defines how to handle requests that can not be executed right now.",
	"type":    "`type` is \"Queue\" or \"Reject\". \"Queue\" means that requests that can not be executed upon arrival are held in a queue until they can be executed or a queuing limit is reached. \"Reject\" means that requests that can not be executed upon arrival are rejected. Required.",
	"queuing": "`queuing` holds the configuration parameters for queuing. This field may be non-empty only if `type` is `\"Queue\"`.",
}

func (LimitResponse) SwaggerDoc() map[string]string {
	return map_LimitResponse
}

var map_LimitedPriorityLevelConfiguration = map[string]string{
	"":                         "LimitedPriorityLevelConfiguration specifies how to handle requests that are subject to limits. It addresses two issues:\n  - How are requests for this priority level limited?\n  - What should be done with requests that exceed the limit?",
	"assuredConcurrencyShares": "`assuredConcurrencyShares` (ACS) configures the execution limit, which is a limit on the number of requests of this priority level that may be exeucting at a given time.  ACS must be a positive number. The server's concurrency limit (SCL) is divided among the concurrency-controlled priority levels in proportion to their assured concurrency shares. This produces the assured concurrency value (ACV) ",
	"limitResponse":            "`limitResponse` indicates what to do with requests that can not be executed right now",
}

func (LimitedPriorityLevelConfiguration) SwaggerDoc() map[string]string {
	return map_LimitedPriorityLevelConfiguration
}

var map_NonResourcePolicyRule = map[string]string{
	"":                "NonResourcePolicyRule is a predicate that matches non-resource requests according to their verb and the target non-resource URL. A NonResourcePolicyRule matches a request if and only if both (a) at least one member of verbs matches the request and (b) at least one member of nonResourceURLs matches the request.",
	"verbs":           "`verbs` is a list of matching verbs and may not be empty. \"*\" matches all verbs. If it is present, it must be the only entry. Required.",
	"nonResourceURLs": "`nonResourceURLs` is a set of url prefixes that a user should have access to and may not be empty. For example:\n  - \"/healthz\" is legal\n  - \"/hea*\" is illegal\n  - \"/hea\" is legal but matches nothing\n  - \"/hea/*\" also matches nothing\n  - \"/healthz/*\" matches all per-component health checks.\n\"*\" matches all non-resource urls. if it is present, it must be the only entry. Required.",
}

func (NonResourcePolicyRule) SwaggerDoc() map[string]string {
	return map_NonResourcePolicyRule
}

var map_PolicyRulesWithSubjects = map[string]string{
	"":                 "PolicyRulesWithSubjects prescribes a test that applies to a request to an apiserver. The test considers the subject making the request, the verb being requested, and the resource to be acted upon. This PolicyRulesWithSubjects matches a request if and only if both (a) at least one member of subjects matches the request and (b) at least one member of resourceRules or nonResourceRules matches the request.",
	"subjects":         "subjects is the list of normal user, serviceaccount, or group that this rule cares about. There must be at least one member in this slice. A slice that includes both the system:authenticated and system:unauthenticated user groups matches every request. Required.",
	"resourceRules":    "`resourceRules` is a slice of ResourcePolicyRules that identify matching requests according to their verb and the target resource. At least one of `resourceRules` and `nonResourceRules` has to be non-empty.",
	"nonResourceRules": "`nonResourceRules` is a list of NonResourcePolicyRules that identify matching requests according to their verb and the target non-resource URL.",
}

func (PolicyRulesWithSubjects) SwaggerDoc() map[string]string {
	return map_PolicyRulesWithSubjects
}

var map_PriorityLevelConfiguration = map[string]string{
	"":         "PriorityLevelConfiguration represents the configuration of a priority level.",
	"metadata": "`metadata` is the standard object's metadata. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata",
	"spec":     "`spec` is the specification of the desired behavior of a \"request-priority\". More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#spec-and-status",
	"status":   "`status` is the current status of a \"request-priority\". More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#spec-and-status",
}

func (PriorityLevelConfiguration) SwaggerDoc() map[string]string {
	return map_PriorityLevelConfiguration
}

var map_PriorityLevelConfigurationCondition = map[string]string{
	"":                   "PriorityLevelConfigurationCondition defines the condition of priority level.",
	"type":               "`type` is the type of the condition. Required.",
	"status":             "`status` is the status of the condition. Can be True, False, Unknown. Required.",
	"lastTransitionTime": "`lastTransitionTime` is the last time the condition transitioned from one status to another.",
	"reason":             "`reason` is a unique, one-word, CamelCase reason for the condition's last transition.",
	"message":            "`message` is a human-readable message indicating details about last transition.",
}

func (PriorityLevelConfigurationCondition) SwaggerDoc() map[string]string {
	return map_PriorityLevelConfigurationCondition
}

var map_PriorityLevelConfigurationList = map[string]string{
	"":         "PriorityLevelConfigurationList is a list of PriorityLevelConfiguration objects.",
	"metadata": "`metadata` is the standard object's metadata. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata",
	"items":    "`items` is a list of request-priorities.",
}

func (PriorityLevelConfigurationList) SwaggerDoc() map[string]string {
	return map_PriorityLevelConfigurationList
}

var map_PriorityLevelConfigurationReference = map[string]string{
	"":     "PriorityLevelConfigurationReference contains information that points to the \"request-priority\" being used.",
	"name": "`name` is the name of the priority level configuration being referenced Required.",
}

func (PriorityLevelConfigurationReference) SwaggerDoc() map[string]string {
	return map_PriorityLevelConfigurationReference
}

var map_PriorityLevelConfigurationSpec = map[string]string{
	"":        "PriorityLevelConfigurationSpec specifies the configuration of a priority level.",
	"type":    "`type` indicates whether this priority level is subject to limitation on request execution.  A value of `\"Exempt\"` means that requests of this priority level are not subject to a limit (and thus are never queued) and do not detract from the capacity made available to other priority levels.  A value of `\"Limited\"` means that (a) requests of this priority level _are_ subject to limits and (b) some of the server's limited capacity is made available exclusively to this priority level. Required.",
	"limited": "`limited` specifies how requests are handled for a Limited priority level. This field must be non-empty if and only if `type` is `\"Limited\"`.",
}

func (PriorityLevelConfigurationSpec) SwaggerDoc() map[string]string {
	return map_PriorityLevelConfigurationSpec
}

var map_PriorityLevelConfigurationStatus = map[string]string{
	"":           "PriorityLevelConfigurationStatus represents the current state of a \"request-priority\".",
	"conditions": "`conditions` is the current state of \"request-priority\".",
}

func (PriorityLevelConfigurationStatus) SwaggerDoc() map[string]string {
	return map_PriorityLevelConfigurationStatus
}

var map_QueuingConfiguration = map[string]string{
	"":                 "QueuingConfiguration holds the configuration parameters for queuing",
	"queues":           "`queues` is the number of queues for this priority level. The queues exist independently at each apiserver. The value must be positive.  Setting it to 1 effectively precludes shufflesharding and thus makes the distinguisher method of associated flow schemas irrelevant.  This field has a default value of 64.",
	"handSize":         "`handSize` is a small positive number that configures the shuffle sharding of requests into queues.  When enqueuing a request at this priority level the request's flow identifier (a string pair) is hashed and the hash value is used to shuffle the list of queues and deal a hand of the size specified here.  The request is put into one of the shortest queues in that hand. `handSize` must be no larger than `queues`, and should be significantly smaller (so that a few heavy flows do not saturate most of the queues).  See the user-facing documentation for more extensive guidance on setting this field.  This field has a default value of 8.",
	"queueLengthLimit": "`queueLengthLimit` is the maximum number of requests allowed to be waiting in a given queue of this priority level at a time; excess requests are rejected.  This value must be positive.  If not specified, it will be defaulted to 50.",
}

func (QueuingConfiguration) SwaggerDoc() map[string]string {
	return map_QueuingConfiguration
}

var map_ResourcePolicyRule = map[string]string{
	"":             "ResourcePolicyRule is a predicate that matches some resource requests, testing the request's verb and the target resource. A ResourcePolicyRule matches a resource request if and only if: (a) at least one member of verbs matches the request, (b) at least one member of apiGroups matches the request, (c) at least one member of resources matches the request, and (d) either (d1) the request does not specify a namespace (i.e., `Namespace==\"\"`) and clusterScope is true or (d2) the request specifies a namespace and least one member of namespaces matches the request's namespace.",
	"verbs":        "`verbs` is a list of matching verbs and may not be empty. \"*\" matches all verbs and, if present, must be the only entry. Required.",
	"apiGroups":    "`apiGroups` is a list of matching API groups and may not be empty. \"*\" matches all API groups and, if present, must be the only entry. Required.",
	"resources":    "`resources` is a list of matching resources (i.e., lowercase and plural) with, if desired, subresource.  For example, [ \"services\", \"nodes/status\" ].  This list may not be empty. \"*\" matches all resources and, if present, must be the only entry. Required.",
	"clusterScope": "`clusterScope` indicates whether to match requests that do not specify a namespace (which happens either because the resource is not namespaced or the request targets all namespaces). If this field is omitted or false then the `namespaces` field must contain a non-empty list.",
	"namespaces":   "`namespaces` is a list of target namespaces that restricts matches.  A request that specifies a target namespace matches only if either (a) this list contains that target namespace or (b) this list contains \"*\".  Note that \"*\" matches any specified namespace but does not match a request that _does not specify_ a namespace (see the `clusterScope` field for that). This list may be empty, but only if `clusterScope` is true.",
}

func (ResourcePolicyRule) SwaggerDoc() map[string]string {
	return map_ResourcePolicyRule
}

var map_ServiceAccountSubject = map[string]string{
	"":          "ServiceAccountSubject holds detailed information for service-account-kind subject.",
	"namespace": "`namespace` is the namespace of matching ServiceAccount objects. Required.",
	"name":      "`name` is the name of matching ServiceAccount objects, or \"*\" to match regardless of name. Required.",
}

func (ServiceAccountSubject) SwaggerDoc() map[string]string {
	return map_ServiceAccountSubject
}

var map_Subject = map[string]string{
	"":               "Subject matches the originator of a request, as identified by the request authentication system. There are three ways of matching an originator; by user, group, or service account.",
	"kind":           "`kind` indicates which one of the other fields is non-empty. Required",
	"user":           "`user` matches based on username.",
	"group":          "`group` matches based on user group name.",
	"serviceAccount": "`serviceAccount` matches ServiceAccounts.",
}

func (Subject) SwaggerDoc() map[string]string {
	return map_Subject
}

var map_UserSubject = map[string]string{
	"":     "UserSubject holds detailed information for user-kind subject.",
	"name": "`name` is the username that matches, or \"*\" to match all usernames. Required.",
}

func (UserSubject) SwaggerDoc() map[string]string {
	return map_UserSubject
}

// AUTO-GENERATED FUNCTIONS END HERE
