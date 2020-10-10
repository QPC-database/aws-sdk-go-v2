// Code generated by smithy-go-codegen DO NOT EDIT.

package pinpointemail

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/pinpointemail/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Verifies an email identity for use with Amazon Pinpoint. In Amazon Pinpoint, an
// identity is an email address or domain that you use when you send email. Before
// you can use an identity to send email with Amazon Pinpoint, you first have to
// verify it. By verifying an address, you demonstrate that you're the owner of the
// address, and that you've given Amazon Pinpoint permission to send email from the
// address. When you verify an email address, Amazon Pinpoint sends an email to the
// address. Your email address is verified as soon as you follow the link in the
// verification email.  </p> <p>When you verify a domain, this operation provides a
// set of DKIM tokens, which you can convert into CNAME tokens. You add these CNAME
// tokens to the DNS configuration for your domain. Your domain is verified when
// Amazon Pinpoint detects these records in the DNS configuration for your domain.
// It usually takes around 72 hours to complete the domain verification
// process.</p>
func (c *Client) CreateEmailIdentity(ctx context.Context, params *CreateEmailIdentityInput, optFns ...func(*Options)) (*CreateEmailIdentityOutput, error) {
	if params == nil {
		params = &CreateEmailIdentityInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateEmailIdentity", params, optFns, addOperationCreateEmailIdentityMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateEmailIdentityOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// A request to begin the verification process for an email identity (an email
// address or domain).
type CreateEmailIdentityInput struct {

	// The email address or domain that you want to verify.
	//
	// This member is required.
	EmailIdentity *string

	// An array of objects that define the tags (keys and values) that you want to
	// associate with the email identity.
	Tags []*types.Tag
}

// If the email identity is a domain, this object contains tokens that you can use
// to create a set of CNAME records. To sucessfully verify your domain, you have to
// add these records to the DNS configuration for your domain. If the email
// identity is an email address, this object is empty.
type CreateEmailIdentityOutput struct {

	// An object that contains information about the DKIM attributes for the identity.
	// This object includes the tokens that you use to create the CNAME records that
	// are required to complete the DKIM verification process.
	DkimAttributes *types.DkimAttributes

	// The email identity type.
	IdentityType types.IdentityType

	// Specifies whether or not the identity is verified. In Amazon Pinpoint, you can
	// only send email from verified email addresses or domains. For more information
	// about verifying identities, see the Amazon Pinpoint User Guide
	// (https://docs.aws.amazon.com/pinpoint/latest/userguide/channels-email-manage-verify.html).
	VerifiedForSendingStatus *bool

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationCreateEmailIdentityMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpCreateEmailIdentity{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpCreateEmailIdentity{}, middleware.After)
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
	addOpCreateEmailIdentityValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opCreateEmailIdentity(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opCreateEmailIdentity(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "ses",
		OperationName: "CreateEmailIdentity",
	}
}
