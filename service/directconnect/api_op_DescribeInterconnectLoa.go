// Code generated by smithy-go-codegen DO NOT EDIT.

package directconnect

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/directconnect/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Deprecated. Use DescribeLoa () instead. Gets the LOA-CFA for the specified
// interconnect. The Letter of Authorization - Connecting Facility Assignment
// (LOA-CFA) is a document that is used when establishing your cross connect to AWS
// at the colocation facility. For more information, see Requesting Cross Connects
// at AWS Direct Connect Locations
// (https://docs.aws.amazon.com/directconnect/latest/UserGuide/Colocation.html) in
// the AWS Direct Connect User Guide.
func (c *Client) DescribeInterconnectLoa(ctx context.Context, params *DescribeInterconnectLoaInput, optFns ...func(*Options)) (*DescribeInterconnectLoaOutput, error) {
	if params == nil {
		params = &DescribeInterconnectLoaInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeInterconnectLoa", params, optFns, addOperationDescribeInterconnectLoaMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeInterconnectLoaOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeInterconnectLoaInput struct {

	// The ID of the interconnect.
	//
	// This member is required.
	InterconnectId *string

	// The standard media type for the LOA-CFA document. The only supported value is
	// application/pdf.
	LoaContentType types.LoaContentType

	// The name of the service provider who establishes connectivity on your behalf. If
	// you supply this parameter, the LOA-CFA lists the provider name alongside your
	// company name as the requester of the cross connect.
	ProviderName *string
}

type DescribeInterconnectLoaOutput struct {

	// The Letter of Authorization - Connecting Facility Assignment (LOA-CFA).
	Loa *types.Loa

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationDescribeInterconnectLoaMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpDescribeInterconnectLoa{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpDescribeInterconnectLoa{}, middleware.After)
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
	addOpDescribeInterconnectLoaValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeInterconnectLoa(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opDescribeInterconnectLoa(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "directconnect",
		OperationName: "DescribeInterconnectLoa",
	}
}
