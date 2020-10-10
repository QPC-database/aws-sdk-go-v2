// Code generated by smithy-go-codegen DO NOT EDIT.

package rds

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Deletes a DB subnet group. The specified database subnet group must not be
// associated with any DB instances.
func (c *Client) DeleteDBSubnetGroup(ctx context.Context, params *DeleteDBSubnetGroupInput, optFns ...func(*Options)) (*DeleteDBSubnetGroupOutput, error) {
	if params == nil {
		params = &DeleteDBSubnetGroupInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DeleteDBSubnetGroup", params, optFns, addOperationDeleteDBSubnetGroupMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DeleteDBSubnetGroupOutput)
	out.ResultMetadata = metadata
	return out, nil
}

//
type DeleteDBSubnetGroupInput struct {

	// The name of the database subnet group to delete. You can't delete the default
	// subnet group. Constraints: Constraints: Must match the name of an existing
	// DBSubnetGroup. Must not be default. Example: mySubnetgroup
	//
	// This member is required.
	DBSubnetGroupName *string
}

type DeleteDBSubnetGroupOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationDeleteDBSubnetGroupMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsquery_serializeOpDeleteDBSubnetGroup{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsquery_deserializeOpDeleteDBSubnetGroup{}, middleware.After)
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
	addOpDeleteDBSubnetGroupValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opDeleteDBSubnetGroup(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opDeleteDBSubnetGroup(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "rds",
		OperationName: "DeleteDBSubnetGroup",
	}
}
