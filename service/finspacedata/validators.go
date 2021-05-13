// Code generated by smithy-go-codegen DO NOT EDIT.

package finspacedata

import (
	"context"
	"fmt"
	smithy "github.com/aws/smithy-go"
	"github.com/aws/smithy-go/middleware"
)

type validateOpCreateChangeset struct {
}

func (*validateOpCreateChangeset) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpCreateChangeset) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*CreateChangesetInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpCreateChangesetInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpGetProgrammaticAccessCredentials struct {
}

func (*validateOpGetProgrammaticAccessCredentials) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpGetProgrammaticAccessCredentials) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*GetProgrammaticAccessCredentialsInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpGetProgrammaticAccessCredentialsInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

func addOpCreateChangesetValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpCreateChangeset{}, middleware.After)
}

func addOpGetProgrammaticAccessCredentialsValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpGetProgrammaticAccessCredentials{}, middleware.After)
}

func validateOpCreateChangesetInput(v *CreateChangesetInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "CreateChangesetInput"}
	if v.DatasetId == nil {
		invalidParams.Add(smithy.NewErrParamRequired("DatasetId"))
	}
	if len(v.ChangeType) == 0 {
		invalidParams.Add(smithy.NewErrParamRequired("ChangeType"))
	}
	if len(v.SourceType) == 0 {
		invalidParams.Add(smithy.NewErrParamRequired("SourceType"))
	}
	if v.SourceParams == nil {
		invalidParams.Add(smithy.NewErrParamRequired("SourceParams"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpGetProgrammaticAccessCredentialsInput(v *GetProgrammaticAccessCredentialsInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "GetProgrammaticAccessCredentialsInput"}
	if v.EnvironmentId == nil {
		invalidParams.Add(smithy.NewErrParamRequired("EnvironmentId"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}
