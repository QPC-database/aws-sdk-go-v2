// Code generated by smithy-go-codegen DO NOT EDIT.

package detective

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/detective/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Returns the membership details for specified member accounts for a behavior
// graph.
func (c *Client) GetMembers(ctx context.Context, params *GetMembersInput, optFns ...func(*Options)) (*GetMembersOutput, error) {
	if params == nil {
		params = &GetMembersInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetMembers", params, optFns, addOperationGetMembersMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetMembersOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetMembersInput struct {

	// The list of AWS account identifiers for the member account for which to return
	// member details. You cannot use GetMembers to retrieve information about member
	// accounts that were removed from the behavior graph.
	//
	// This member is required.
	AccountIds []*string

	// The ARN of the behavior graph for which to request the member details.
	//
	// This member is required.
	GraphArn *string
}

type GetMembersOutput struct {

	// The member account details that Detective is returning in response to the
	// request.
	MemberDetails []*types.MemberDetail

	// The requested member accounts for which Detective was unable to return member
	// details. For each account, provides the reason why the request could not be
	// processed.
	UnprocessedAccounts []*types.UnprocessedAccount

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationGetMembersMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpGetMembers{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpGetMembers{}, middleware.After)
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
	addOpGetMembersValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opGetMembers(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opGetMembers(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "detective",
		OperationName: "GetMembers",
	}
}
