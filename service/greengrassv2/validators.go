// Code generated by smithy-go-codegen DO NOT EDIT.

package greengrassv2

import (
	"context"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/service/greengrassv2/types"
	smithy "github.com/aws/smithy-go"
	"github.com/aws/smithy-go/middleware"
)

type validateOpBatchAssociateClientDeviceWithCoreDevice struct {
}

func (*validateOpBatchAssociateClientDeviceWithCoreDevice) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpBatchAssociateClientDeviceWithCoreDevice) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*BatchAssociateClientDeviceWithCoreDeviceInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpBatchAssociateClientDeviceWithCoreDeviceInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpBatchDisassociateClientDeviceFromCoreDevice struct {
}

func (*validateOpBatchDisassociateClientDeviceFromCoreDevice) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpBatchDisassociateClientDeviceFromCoreDevice) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*BatchDisassociateClientDeviceFromCoreDeviceInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpBatchDisassociateClientDeviceFromCoreDeviceInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpCancelDeployment struct {
}

func (*validateOpCancelDeployment) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpCancelDeployment) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*CancelDeploymentInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpCancelDeploymentInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpCreateComponentVersion struct {
}

func (*validateOpCreateComponentVersion) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpCreateComponentVersion) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*CreateComponentVersionInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpCreateComponentVersionInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpCreateDeployment struct {
}

func (*validateOpCreateDeployment) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpCreateDeployment) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*CreateDeploymentInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpCreateDeploymentInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpDeleteComponent struct {
}

func (*validateOpDeleteComponent) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpDeleteComponent) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*DeleteComponentInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpDeleteComponentInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpDeleteCoreDevice struct {
}

func (*validateOpDeleteCoreDevice) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpDeleteCoreDevice) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*DeleteCoreDeviceInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpDeleteCoreDeviceInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpDescribeComponent struct {
}

func (*validateOpDescribeComponent) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpDescribeComponent) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*DescribeComponentInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpDescribeComponentInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpGetComponent struct {
}

func (*validateOpGetComponent) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpGetComponent) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*GetComponentInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpGetComponentInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpGetComponentVersionArtifact struct {
}

func (*validateOpGetComponentVersionArtifact) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpGetComponentVersionArtifact) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*GetComponentVersionArtifactInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpGetComponentVersionArtifactInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpGetCoreDevice struct {
}

func (*validateOpGetCoreDevice) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpGetCoreDevice) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*GetCoreDeviceInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpGetCoreDeviceInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpGetDeployment struct {
}

func (*validateOpGetDeployment) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpGetDeployment) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*GetDeploymentInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpGetDeploymentInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpListClientDevicesAssociatedWithCoreDevice struct {
}

func (*validateOpListClientDevicesAssociatedWithCoreDevice) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpListClientDevicesAssociatedWithCoreDevice) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*ListClientDevicesAssociatedWithCoreDeviceInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpListClientDevicesAssociatedWithCoreDeviceInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpListComponentVersions struct {
}

func (*validateOpListComponentVersions) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpListComponentVersions) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*ListComponentVersionsInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpListComponentVersionsInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpListEffectiveDeployments struct {
}

func (*validateOpListEffectiveDeployments) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpListEffectiveDeployments) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*ListEffectiveDeploymentsInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpListEffectiveDeploymentsInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpListInstalledComponents struct {
}

func (*validateOpListInstalledComponents) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpListInstalledComponents) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*ListInstalledComponentsInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpListInstalledComponentsInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpListTagsForResource struct {
}

func (*validateOpListTagsForResource) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpListTagsForResource) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*ListTagsForResourceInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpListTagsForResourceInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpResolveComponentCandidates struct {
}

func (*validateOpResolveComponentCandidates) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpResolveComponentCandidates) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*ResolveComponentCandidatesInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpResolveComponentCandidatesInput(input); err != nil {
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

func addOpBatchAssociateClientDeviceWithCoreDeviceValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpBatchAssociateClientDeviceWithCoreDevice{}, middleware.After)
}

func addOpBatchDisassociateClientDeviceFromCoreDeviceValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpBatchDisassociateClientDeviceFromCoreDevice{}, middleware.After)
}

func addOpCancelDeploymentValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpCancelDeployment{}, middleware.After)
}

func addOpCreateComponentVersionValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpCreateComponentVersion{}, middleware.After)
}

func addOpCreateDeploymentValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpCreateDeployment{}, middleware.After)
}

func addOpDeleteComponentValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpDeleteComponent{}, middleware.After)
}

func addOpDeleteCoreDeviceValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpDeleteCoreDevice{}, middleware.After)
}

func addOpDescribeComponentValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpDescribeComponent{}, middleware.After)
}

func addOpGetComponentValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpGetComponent{}, middleware.After)
}

func addOpGetComponentVersionArtifactValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpGetComponentVersionArtifact{}, middleware.After)
}

func addOpGetCoreDeviceValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpGetCoreDevice{}, middleware.After)
}

func addOpGetDeploymentValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpGetDeployment{}, middleware.After)
}

func addOpListClientDevicesAssociatedWithCoreDeviceValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpListClientDevicesAssociatedWithCoreDevice{}, middleware.After)
}

func addOpListComponentVersionsValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpListComponentVersions{}, middleware.After)
}

func addOpListEffectiveDeploymentsValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpListEffectiveDeployments{}, middleware.After)
}

func addOpListInstalledComponentsValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpListInstalledComponents{}, middleware.After)
}

func addOpListTagsForResourceValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpListTagsForResource{}, middleware.After)
}

func addOpResolveComponentCandidatesValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpResolveComponentCandidates{}, middleware.After)
}

func addOpTagResourceValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpTagResource{}, middleware.After)
}

func addOpUntagResourceValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpUntagResource{}, middleware.After)
}

func validateAssociateClientDeviceWithCoreDeviceEntry(v *types.AssociateClientDeviceWithCoreDeviceEntry) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "AssociateClientDeviceWithCoreDeviceEntry"}
	if v.ThingName == nil {
		invalidParams.Add(smithy.NewErrParamRequired("ThingName"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateAssociateClientDeviceWithCoreDeviceEntryList(v []types.AssociateClientDeviceWithCoreDeviceEntry) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "AssociateClientDeviceWithCoreDeviceEntryList"}
	for i := range v {
		if err := validateAssociateClientDeviceWithCoreDeviceEntry(&v[i]); err != nil {
			invalidParams.AddNested(fmt.Sprintf("[%d]", i), err.(smithy.InvalidParamsError))
		}
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateDeploymentIoTJobConfiguration(v *types.DeploymentIoTJobConfiguration) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "DeploymentIoTJobConfiguration"}
	if v.JobExecutionsRolloutConfig != nil {
		if err := validateIoTJobExecutionsRolloutConfig(v.JobExecutionsRolloutConfig); err != nil {
			invalidParams.AddNested("JobExecutionsRolloutConfig", err.(smithy.InvalidParamsError))
		}
	}
	if v.AbortConfig != nil {
		if err := validateIoTJobAbortConfig(v.AbortConfig); err != nil {
			invalidParams.AddNested("AbortConfig", err.(smithy.InvalidParamsError))
		}
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateDisassociateClientDeviceFromCoreDeviceEntry(v *types.DisassociateClientDeviceFromCoreDeviceEntry) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "DisassociateClientDeviceFromCoreDeviceEntry"}
	if v.ThingName == nil {
		invalidParams.Add(smithy.NewErrParamRequired("ThingName"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateDisassociateClientDeviceFromCoreDeviceEntryList(v []types.DisassociateClientDeviceFromCoreDeviceEntry) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "DisassociateClientDeviceFromCoreDeviceEntryList"}
	for i := range v {
		if err := validateDisassociateClientDeviceFromCoreDeviceEntry(&v[i]); err != nil {
			invalidParams.AddNested(fmt.Sprintf("[%d]", i), err.(smithy.InvalidParamsError))
		}
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateIoTJobAbortConfig(v *types.IoTJobAbortConfig) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "IoTJobAbortConfig"}
	if v.CriteriaList == nil {
		invalidParams.Add(smithy.NewErrParamRequired("CriteriaList"))
	} else if v.CriteriaList != nil {
		if err := validateIoTJobAbortCriteriaList(v.CriteriaList); err != nil {
			invalidParams.AddNested("CriteriaList", err.(smithy.InvalidParamsError))
		}
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateIoTJobAbortCriteria(v *types.IoTJobAbortCriteria) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "IoTJobAbortCriteria"}
	if len(v.FailureType) == 0 {
		invalidParams.Add(smithy.NewErrParamRequired("FailureType"))
	}
	if len(v.Action) == 0 {
		invalidParams.Add(smithy.NewErrParamRequired("Action"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateIoTJobAbortCriteriaList(v []types.IoTJobAbortCriteria) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "IoTJobAbortCriteriaList"}
	for i := range v {
		if err := validateIoTJobAbortCriteria(&v[i]); err != nil {
			invalidParams.AddNested(fmt.Sprintf("[%d]", i), err.(smithy.InvalidParamsError))
		}
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateIoTJobExecutionsRolloutConfig(v *types.IoTJobExecutionsRolloutConfig) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "IoTJobExecutionsRolloutConfig"}
	if v.ExponentialRate != nil {
		if err := validateIoTJobExponentialRolloutRate(v.ExponentialRate); err != nil {
			invalidParams.AddNested("ExponentialRate", err.(smithy.InvalidParamsError))
		}
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateIoTJobExponentialRolloutRate(v *types.IoTJobExponentialRolloutRate) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "IoTJobExponentialRolloutRate"}
	if v.RateIncreaseCriteria == nil {
		invalidParams.Add(smithy.NewErrParamRequired("RateIncreaseCriteria"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateLambdaContainerParams(v *types.LambdaContainerParams) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "LambdaContainerParams"}
	if v.Volumes != nil {
		if err := validateLambdaVolumeList(v.Volumes); err != nil {
			invalidParams.AddNested("Volumes", err.(smithy.InvalidParamsError))
		}
	}
	if v.Devices != nil {
		if err := validateLambdaDeviceList(v.Devices); err != nil {
			invalidParams.AddNested("Devices", err.(smithy.InvalidParamsError))
		}
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateLambdaDeviceList(v []types.LambdaDeviceMount) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "LambdaDeviceList"}
	for i := range v {
		if err := validateLambdaDeviceMount(&v[i]); err != nil {
			invalidParams.AddNested(fmt.Sprintf("[%d]", i), err.(smithy.InvalidParamsError))
		}
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateLambdaDeviceMount(v *types.LambdaDeviceMount) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "LambdaDeviceMount"}
	if v.Path == nil {
		invalidParams.Add(smithy.NewErrParamRequired("Path"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateLambdaEventSource(v *types.LambdaEventSource) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "LambdaEventSource"}
	if v.Topic == nil {
		invalidParams.Add(smithy.NewErrParamRequired("Topic"))
	}
	if len(v.Type) == 0 {
		invalidParams.Add(smithy.NewErrParamRequired("Type"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateLambdaEventSourceList(v []types.LambdaEventSource) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "LambdaEventSourceList"}
	for i := range v {
		if err := validateLambdaEventSource(&v[i]); err != nil {
			invalidParams.AddNested(fmt.Sprintf("[%d]", i), err.(smithy.InvalidParamsError))
		}
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateLambdaExecutionParameters(v *types.LambdaExecutionParameters) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "LambdaExecutionParameters"}
	if v.EventSources != nil {
		if err := validateLambdaEventSourceList(v.EventSources); err != nil {
			invalidParams.AddNested("EventSources", err.(smithy.InvalidParamsError))
		}
	}
	if v.LinuxProcessParams != nil {
		if err := validateLambdaLinuxProcessParams(v.LinuxProcessParams); err != nil {
			invalidParams.AddNested("LinuxProcessParams", err.(smithy.InvalidParamsError))
		}
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateLambdaFunctionRecipeSource(v *types.LambdaFunctionRecipeSource) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "LambdaFunctionRecipeSource"}
	if v.LambdaArn == nil {
		invalidParams.Add(smithy.NewErrParamRequired("LambdaArn"))
	}
	if v.ComponentLambdaParameters != nil {
		if err := validateLambdaExecutionParameters(v.ComponentLambdaParameters); err != nil {
			invalidParams.AddNested("ComponentLambdaParameters", err.(smithy.InvalidParamsError))
		}
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateLambdaLinuxProcessParams(v *types.LambdaLinuxProcessParams) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "LambdaLinuxProcessParams"}
	if v.ContainerParams != nil {
		if err := validateLambdaContainerParams(v.ContainerParams); err != nil {
			invalidParams.AddNested("ContainerParams", err.(smithy.InvalidParamsError))
		}
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateLambdaVolumeList(v []types.LambdaVolumeMount) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "LambdaVolumeList"}
	for i := range v {
		if err := validateLambdaVolumeMount(&v[i]); err != nil {
			invalidParams.AddNested(fmt.Sprintf("[%d]", i), err.(smithy.InvalidParamsError))
		}
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateLambdaVolumeMount(v *types.LambdaVolumeMount) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "LambdaVolumeMount"}
	if v.SourcePath == nil {
		invalidParams.Add(smithy.NewErrParamRequired("SourcePath"))
	}
	if v.DestinationPath == nil {
		invalidParams.Add(smithy.NewErrParamRequired("DestinationPath"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpBatchAssociateClientDeviceWithCoreDeviceInput(v *BatchAssociateClientDeviceWithCoreDeviceInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "BatchAssociateClientDeviceWithCoreDeviceInput"}
	if v.Entries != nil {
		if err := validateAssociateClientDeviceWithCoreDeviceEntryList(v.Entries); err != nil {
			invalidParams.AddNested("Entries", err.(smithy.InvalidParamsError))
		}
	}
	if v.CoreDeviceThingName == nil {
		invalidParams.Add(smithy.NewErrParamRequired("CoreDeviceThingName"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpBatchDisassociateClientDeviceFromCoreDeviceInput(v *BatchDisassociateClientDeviceFromCoreDeviceInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "BatchDisassociateClientDeviceFromCoreDeviceInput"}
	if v.Entries != nil {
		if err := validateDisassociateClientDeviceFromCoreDeviceEntryList(v.Entries); err != nil {
			invalidParams.AddNested("Entries", err.(smithy.InvalidParamsError))
		}
	}
	if v.CoreDeviceThingName == nil {
		invalidParams.Add(smithy.NewErrParamRequired("CoreDeviceThingName"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpCancelDeploymentInput(v *CancelDeploymentInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "CancelDeploymentInput"}
	if v.DeploymentId == nil {
		invalidParams.Add(smithy.NewErrParamRequired("DeploymentId"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpCreateComponentVersionInput(v *CreateComponentVersionInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "CreateComponentVersionInput"}
	if v.LambdaFunction != nil {
		if err := validateLambdaFunctionRecipeSource(v.LambdaFunction); err != nil {
			invalidParams.AddNested("LambdaFunction", err.(smithy.InvalidParamsError))
		}
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpCreateDeploymentInput(v *CreateDeploymentInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "CreateDeploymentInput"}
	if v.TargetArn == nil {
		invalidParams.Add(smithy.NewErrParamRequired("TargetArn"))
	}
	if v.IotJobConfiguration != nil {
		if err := validateDeploymentIoTJobConfiguration(v.IotJobConfiguration); err != nil {
			invalidParams.AddNested("IotJobConfiguration", err.(smithy.InvalidParamsError))
		}
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpDeleteComponentInput(v *DeleteComponentInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "DeleteComponentInput"}
	if v.Arn == nil {
		invalidParams.Add(smithy.NewErrParamRequired("Arn"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpDeleteCoreDeviceInput(v *DeleteCoreDeviceInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "DeleteCoreDeviceInput"}
	if v.CoreDeviceThingName == nil {
		invalidParams.Add(smithy.NewErrParamRequired("CoreDeviceThingName"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpDescribeComponentInput(v *DescribeComponentInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "DescribeComponentInput"}
	if v.Arn == nil {
		invalidParams.Add(smithy.NewErrParamRequired("Arn"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpGetComponentInput(v *GetComponentInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "GetComponentInput"}
	if v.Arn == nil {
		invalidParams.Add(smithy.NewErrParamRequired("Arn"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpGetComponentVersionArtifactInput(v *GetComponentVersionArtifactInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "GetComponentVersionArtifactInput"}
	if v.Arn == nil {
		invalidParams.Add(smithy.NewErrParamRequired("Arn"))
	}
	if v.ArtifactName == nil {
		invalidParams.Add(smithy.NewErrParamRequired("ArtifactName"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpGetCoreDeviceInput(v *GetCoreDeviceInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "GetCoreDeviceInput"}
	if v.CoreDeviceThingName == nil {
		invalidParams.Add(smithy.NewErrParamRequired("CoreDeviceThingName"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpGetDeploymentInput(v *GetDeploymentInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "GetDeploymentInput"}
	if v.DeploymentId == nil {
		invalidParams.Add(smithy.NewErrParamRequired("DeploymentId"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpListClientDevicesAssociatedWithCoreDeviceInput(v *ListClientDevicesAssociatedWithCoreDeviceInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "ListClientDevicesAssociatedWithCoreDeviceInput"}
	if v.CoreDeviceThingName == nil {
		invalidParams.Add(smithy.NewErrParamRequired("CoreDeviceThingName"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpListComponentVersionsInput(v *ListComponentVersionsInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "ListComponentVersionsInput"}
	if v.Arn == nil {
		invalidParams.Add(smithy.NewErrParamRequired("Arn"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpListEffectiveDeploymentsInput(v *ListEffectiveDeploymentsInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "ListEffectiveDeploymentsInput"}
	if v.CoreDeviceThingName == nil {
		invalidParams.Add(smithy.NewErrParamRequired("CoreDeviceThingName"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpListInstalledComponentsInput(v *ListInstalledComponentsInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "ListInstalledComponentsInput"}
	if v.CoreDeviceThingName == nil {
		invalidParams.Add(smithy.NewErrParamRequired("CoreDeviceThingName"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpListTagsForResourceInput(v *ListTagsForResourceInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "ListTagsForResourceInput"}
	if v.ResourceArn == nil {
		invalidParams.Add(smithy.NewErrParamRequired("ResourceArn"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpResolveComponentCandidatesInput(v *ResolveComponentCandidatesInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "ResolveComponentCandidatesInput"}
	if v.Platform == nil {
		invalidParams.Add(smithy.NewErrParamRequired("Platform"))
	}
	if v.ComponentCandidates == nil {
		invalidParams.Add(smithy.NewErrParamRequired("ComponentCandidates"))
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
	if v.ResourceArn == nil {
		invalidParams.Add(smithy.NewErrParamRequired("ResourceArn"))
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
	if v.ResourceArn == nil {
		invalidParams.Add(smithy.NewErrParamRequired("ResourceArn"))
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
