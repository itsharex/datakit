// Code generated by smithy-go-codegen DO NOT EDIT.

package secretsmanager

import (
	"context"
	"fmt"
	smithy "github.com/aws/smithy-go"
	"github.com/aws/smithy-go/middleware"
)

type validateOpCancelRotateSecret struct {
}

func (*validateOpCancelRotateSecret) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpCancelRotateSecret) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*CancelRotateSecretInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpCancelRotateSecretInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpCreateSecret struct {
}

func (*validateOpCreateSecret) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpCreateSecret) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*CreateSecretInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpCreateSecretInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpDeleteResourcePolicy struct {
}

func (*validateOpDeleteResourcePolicy) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpDeleteResourcePolicy) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*DeleteResourcePolicyInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpDeleteResourcePolicyInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpDeleteSecret struct {
}

func (*validateOpDeleteSecret) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpDeleteSecret) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*DeleteSecretInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpDeleteSecretInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpDescribeSecret struct {
}

func (*validateOpDescribeSecret) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpDescribeSecret) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*DescribeSecretInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpDescribeSecretInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpGetResourcePolicy struct {
}

func (*validateOpGetResourcePolicy) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpGetResourcePolicy) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*GetResourcePolicyInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpGetResourcePolicyInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpGetSecretValue struct {
}

func (*validateOpGetSecretValue) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpGetSecretValue) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*GetSecretValueInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpGetSecretValueInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpListSecretVersionIds struct {
}

func (*validateOpListSecretVersionIds) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpListSecretVersionIds) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*ListSecretVersionIdsInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpListSecretVersionIdsInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpPutResourcePolicy struct {
}

func (*validateOpPutResourcePolicy) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpPutResourcePolicy) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*PutResourcePolicyInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpPutResourcePolicyInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpPutSecretValue struct {
}

func (*validateOpPutSecretValue) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpPutSecretValue) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*PutSecretValueInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpPutSecretValueInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpRemoveRegionsFromReplication struct {
}

func (*validateOpRemoveRegionsFromReplication) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpRemoveRegionsFromReplication) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*RemoveRegionsFromReplicationInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpRemoveRegionsFromReplicationInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpReplicateSecretToRegions struct {
}

func (*validateOpReplicateSecretToRegions) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpReplicateSecretToRegions) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*ReplicateSecretToRegionsInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpReplicateSecretToRegionsInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpRestoreSecret struct {
}

func (*validateOpRestoreSecret) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpRestoreSecret) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*RestoreSecretInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpRestoreSecretInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpRotateSecret struct {
}

func (*validateOpRotateSecret) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpRotateSecret) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*RotateSecretInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpRotateSecretInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpStopReplicationToReplica struct {
}

func (*validateOpStopReplicationToReplica) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpStopReplicationToReplica) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*StopReplicationToReplicaInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpStopReplicationToReplicaInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpTagResource struct {
}

func (*validateOpTagResource) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpTagResource) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*TagResourceInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpTagResourceInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpUntagResource struct {
}

func (*validateOpUntagResource) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpUntagResource) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*UntagResourceInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpUntagResourceInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpUpdateSecret struct {
}

func (*validateOpUpdateSecret) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpUpdateSecret) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*UpdateSecretInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpUpdateSecretInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpUpdateSecretVersionStage struct {
}

func (*validateOpUpdateSecretVersionStage) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpUpdateSecretVersionStage) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*UpdateSecretVersionStageInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpUpdateSecretVersionStageInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpValidateResourcePolicy struct {
}

func (*validateOpValidateResourcePolicy) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpValidateResourcePolicy) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*ValidateResourcePolicyInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpValidateResourcePolicyInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

func addOpCancelRotateSecretValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpCancelRotateSecret{}, middleware.After)
}

func addOpCreateSecretValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpCreateSecret{}, middleware.After)
}

func addOpDeleteResourcePolicyValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpDeleteResourcePolicy{}, middleware.After)
}

func addOpDeleteSecretValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpDeleteSecret{}, middleware.After)
}

func addOpDescribeSecretValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpDescribeSecret{}, middleware.After)
}

func addOpGetResourcePolicyValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpGetResourcePolicy{}, middleware.After)
}

func addOpGetSecretValueValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpGetSecretValue{}, middleware.After)
}

func addOpListSecretVersionIdsValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpListSecretVersionIds{}, middleware.After)
}

func addOpPutResourcePolicyValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpPutResourcePolicy{}, middleware.After)
}

func addOpPutSecretValueValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpPutSecretValue{}, middleware.After)
}

func addOpRemoveRegionsFromReplicationValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpRemoveRegionsFromReplication{}, middleware.After)
}

func addOpReplicateSecretToRegionsValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpReplicateSecretToRegions{}, middleware.After)
}

func addOpRestoreSecretValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpRestoreSecret{}, middleware.After)
}

func addOpRotateSecretValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpRotateSecret{}, middleware.After)
}

func addOpStopReplicationToReplicaValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpStopReplicationToReplica{}, middleware.After)
}

func addOpTagResourceValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpTagResource{}, middleware.After)
}

func addOpUntagResourceValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpUntagResource{}, middleware.After)
}

func addOpUpdateSecretValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpUpdateSecret{}, middleware.After)
}

func addOpUpdateSecretVersionStageValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpUpdateSecretVersionStage{}, middleware.After)
}

func addOpValidateResourcePolicyValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpValidateResourcePolicy{}, middleware.After)
}

func validateOpCancelRotateSecretInput(v *CancelRotateSecretInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "CancelRotateSecretInput"}
	if v.SecretId == nil {
		invalidParams.Add(smithy.NewErrParamRequired("SecretId"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpCreateSecretInput(v *CreateSecretInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "CreateSecretInput"}
	if v.Name == nil {
		invalidParams.Add(smithy.NewErrParamRequired("Name"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpDeleteResourcePolicyInput(v *DeleteResourcePolicyInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "DeleteResourcePolicyInput"}
	if v.SecretId == nil {
		invalidParams.Add(smithy.NewErrParamRequired("SecretId"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpDeleteSecretInput(v *DeleteSecretInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "DeleteSecretInput"}
	if v.SecretId == nil {
		invalidParams.Add(smithy.NewErrParamRequired("SecretId"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpDescribeSecretInput(v *DescribeSecretInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "DescribeSecretInput"}
	if v.SecretId == nil {
		invalidParams.Add(smithy.NewErrParamRequired("SecretId"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpGetResourcePolicyInput(v *GetResourcePolicyInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "GetResourcePolicyInput"}
	if v.SecretId == nil {
		invalidParams.Add(smithy.NewErrParamRequired("SecretId"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpGetSecretValueInput(v *GetSecretValueInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "GetSecretValueInput"}
	if v.SecretId == nil {
		invalidParams.Add(smithy.NewErrParamRequired("SecretId"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpListSecretVersionIdsInput(v *ListSecretVersionIdsInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "ListSecretVersionIdsInput"}
	if v.SecretId == nil {
		invalidParams.Add(smithy.NewErrParamRequired("SecretId"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpPutResourcePolicyInput(v *PutResourcePolicyInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "PutResourcePolicyInput"}
	if v.SecretId == nil {
		invalidParams.Add(smithy.NewErrParamRequired("SecretId"))
	}
	if v.ResourcePolicy == nil {
		invalidParams.Add(smithy.NewErrParamRequired("ResourcePolicy"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpPutSecretValueInput(v *PutSecretValueInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "PutSecretValueInput"}
	if v.SecretId == nil {
		invalidParams.Add(smithy.NewErrParamRequired("SecretId"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpRemoveRegionsFromReplicationInput(v *RemoveRegionsFromReplicationInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "RemoveRegionsFromReplicationInput"}
	if v.SecretId == nil {
		invalidParams.Add(smithy.NewErrParamRequired("SecretId"))
	}
	if v.RemoveReplicaRegions == nil {
		invalidParams.Add(smithy.NewErrParamRequired("RemoveReplicaRegions"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpReplicateSecretToRegionsInput(v *ReplicateSecretToRegionsInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "ReplicateSecretToRegionsInput"}
	if v.SecretId == nil {
		invalidParams.Add(smithy.NewErrParamRequired("SecretId"))
	}
	if v.AddReplicaRegions == nil {
		invalidParams.Add(smithy.NewErrParamRequired("AddReplicaRegions"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpRestoreSecretInput(v *RestoreSecretInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "RestoreSecretInput"}
	if v.SecretId == nil {
		invalidParams.Add(smithy.NewErrParamRequired("SecretId"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpRotateSecretInput(v *RotateSecretInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "RotateSecretInput"}
	if v.SecretId == nil {
		invalidParams.Add(smithy.NewErrParamRequired("SecretId"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpStopReplicationToReplicaInput(v *StopReplicationToReplicaInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "StopReplicationToReplicaInput"}
	if v.SecretId == nil {
		invalidParams.Add(smithy.NewErrParamRequired("SecretId"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpTagResourceInput(v *TagResourceInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "TagResourceInput"}
	if v.SecretId == nil {
		invalidParams.Add(smithy.NewErrParamRequired("SecretId"))
	}
	if v.Tags == nil {
		invalidParams.Add(smithy.NewErrParamRequired("Tags"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpUntagResourceInput(v *UntagResourceInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "UntagResourceInput"}
	if v.SecretId == nil {
		invalidParams.Add(smithy.NewErrParamRequired("SecretId"))
	}
	if v.TagKeys == nil {
		invalidParams.Add(smithy.NewErrParamRequired("TagKeys"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpUpdateSecretInput(v *UpdateSecretInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "UpdateSecretInput"}
	if v.SecretId == nil {
		invalidParams.Add(smithy.NewErrParamRequired("SecretId"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpUpdateSecretVersionStageInput(v *UpdateSecretVersionStageInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "UpdateSecretVersionStageInput"}
	if v.SecretId == nil {
		invalidParams.Add(smithy.NewErrParamRequired("SecretId"))
	}
	if v.VersionStage == nil {
		invalidParams.Add(smithy.NewErrParamRequired("VersionStage"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpValidateResourcePolicyInput(v *ValidateResourcePolicyInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "ValidateResourcePolicyInput"}
	if v.ResourcePolicy == nil {
		invalidParams.Add(smithy.NewErrParamRequired("ResourcePolicy"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}
