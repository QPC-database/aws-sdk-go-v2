// Code generated by smithy-go-codegen DO NOT EDIT.

package polly

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Stores a pronunciation lexicon in an AWS Region. If a lexicon with the same name
// already exists in the region, it is overwritten by the new lexicon. Lexicon
// operations have eventual consistency, therefore, it might take some time before
// the lexicon is available to the SynthesizeSpeech operation. For more
// information, see Managing Lexicons
// (https://docs.aws.amazon.com/polly/latest/dg/managing-lexicons.html).
func (c *Client) PutLexicon(ctx context.Context, params *PutLexiconInput, optFns ...func(*Options)) (*PutLexiconOutput, error) {
	if params == nil {
		params = &PutLexiconInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "PutLexicon", params, optFns, addOperationPutLexiconMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*PutLexiconOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type PutLexiconInput struct {

	// Content of the PLS lexicon as string data.
	//
	// This member is required.
	Content *string

	// Name of the lexicon. The name must follow the regular express format
	// [0-9A-Za-z]{1,20}. That is, the name is a case-sensitive alphanumeric string up
	// to 20 characters long.
	//
	// This member is required.
	Name *string
}

type PutLexiconOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationPutLexiconMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpPutLexicon{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpPutLexicon{}, middleware.After)
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
	addOpPutLexiconValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opPutLexicon(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opPutLexicon(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "polly",
		OperationName: "PutLexicon",
	}
}
