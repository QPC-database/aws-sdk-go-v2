// Code generated by smithy-go-codegen DO NOT EDIT.

package codegurureviewer

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/codegurureviewer/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Removes the association between Amazon CodeGuru Reviewer and a repository.
func (c *Client) DisassociateRepository(ctx context.Context, params *DisassociateRepositoryInput, optFns ...func(*Options)) (*DisassociateRepositoryOutput, error) {
	if params == nil {
		params = &DisassociateRepositoryInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DisassociateRepository", params, optFns, addOperationDisassociateRepositoryMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DisassociateRepositoryOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DisassociateRepositoryInput struct {

	// The Amazon Resource Name (ARN) of the RepositoryAssociation
	// (https://docs.aws.amazon.com/codeguru/latest/reviewer-api/API_RepositoryAssociation.html)
	// object.
	//
	// This member is required.
	AssociationArn *string
}

type DisassociateRepositoryOutput struct {

	// Information about the disassociated repository.
	RepositoryAssociation *types.RepositoryAssociation

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationDisassociateRepositoryMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpDisassociateRepository{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpDisassociateRepository{}, middleware.After)
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
	addOpDisassociateRepositoryValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opDisassociateRepository(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opDisassociateRepository(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "codeguru-reviewer",
		OperationName: "DisassociateRepository",
	}
}
