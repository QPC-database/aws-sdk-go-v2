// Code generated by smithy-go-codegen DO NOT EDIT.

package fms

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/fms/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Creates an AWS Firewall Manager applications list.
func (c *Client) PutAppsList(ctx context.Context, params *PutAppsListInput, optFns ...func(*Options)) (*PutAppsListOutput, error) {
	if params == nil {
		params = &PutAppsListInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "PutAppsList", params, optFns, addOperationPutAppsListMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*PutAppsListOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type PutAppsListInput struct {

	// The details of the AWS Firewall Manager applications list to be created.
	//
	// This member is required.
	AppsList *types.AppsListData

	// The tags associated with the resource.
	TagList []*types.Tag
}

type PutAppsListOutput struct {

	// The details of the AWS Firewall Manager applications list.
	AppsList *types.AppsListData

	// The Amazon Resource Name (ARN) of the applications list.
	AppsListArn *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationPutAppsListMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpPutAppsList{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpPutAppsList{}, middleware.After)
	if err != nil {
		return err
	}
	awsmiddleware.AddRequestInvocationIDMiddleware(stack)
	smithyhttp.AddContentLengthMiddleware(stack)
	addResolveEndpointMiddleware(stack, options)
	v4.AddComputePayloadSHA256Middleware(stack)
	addRetryMiddlewares(stack, options)
	addHTTPSignerV4Middleware(stack, options)
	awsmiddleware.AddAttemptClockSkewMiddleware(stack)
	addClientUserAgent(stack)
	smithyhttp.AddErrorCloseResponseBodyMiddleware(stack)
	smithyhttp.AddCloseResponseBodyMiddleware(stack)
	addOpPutAppsListValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opPutAppsList(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opPutAppsList(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "fms",
		OperationName: "PutAppsList",
	}
}
