// Code generated by smithy-go-codegen DO NOT EDIT.

package kms

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/kms/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Decrypts ciphertext and then reencrypts it entirely within AWS KMS. You can use
// this operation to change the customer master key (CMK) under which data is
// encrypted, such as when you manually rotate
// (https://docs.aws.amazon.com/kms/latest/developerguide/rotate-keys.html#rotate-keys-manually)
// a CMK or change the CMK that protects a ciphertext. You can also use it to
// reencrypt ciphertext under the same CMK, such as to change the encryption
// context
// (https://docs.aws.amazon.com/kms/latest/developerguide/concepts.html#encrypt_context)
// of a ciphertext. The ReEncrypt operation can decrypt ciphertext that was
// encrypted by using an AWS KMS CMK in an AWS KMS operation, such as Encrypt () or
// GenerateDataKey (). It can also decrypt ciphertext that was encrypted by using
// the public key of an asymmetric CMK
// (https://docs.aws.amazon.com/kms/latest/developerguide/symm-asymm-concepts.html#asymmetric-cmks)
// outside of AWS KMS. However, it cannot decrypt ciphertext produced by other
// libraries, such as the AWS Encryption SDK
// (https://docs.aws.amazon.com/encryption-sdk/latest/developer-guide/) or Amazon
// S3 client-side encryption
// (https://docs.aws.amazon.com/AmazonS3/latest/dev/UsingClientSideEncryption.html).
// These libraries return a ciphertext format that is incompatible with AWS KMS.
// When you use the ReEncrypt operation, you need to provide information for the
// decrypt operation and the subsequent encrypt operation.
//
//     * If your
// ciphertext was encrypted under an asymmetric CMK, you must identify the source
// CMK, that is, the CMK that encrypted the ciphertext. You must also supply the
// encryption algorithm that was used. This information is required to decrypt the
// data.
//
//     * It is optional, but you can specify a source CMK even when the
// ciphertext was encrypted under a symmetric CMK. This ensures that the ciphertext
// is decrypted only by using a particular CMK. If the CMK that you specify cannot
// decrypt the ciphertext, the ReEncrypt operation fails.
//
//     * To reencrypt the
// data, you must specify the destination CMK, that is, the CMK that re-encrypts
// the data after it is decrypted. You can select a symmetric or asymmetric CMK. If
// the destination CMK is an asymmetric CMK, you must also provide the encryption
// algorithm. The algorithm that you choose must be compatible with the CMK.
// <important> <p>When you use an asymmetric CMK to encrypt or reencrypt data, be
// sure to record the CMK and encryption algorithm that you choose. You will be
// required to provide the same CMK and encryption algorithm when you decrypt the
// data. If the CMK and algorithm do not match the values used to encrypt the data,
// the decrypt operation fails.</p> <p>You are not required to supply the CMK ID
// and encryption algorithm when you decrypt with symmetric CMKs because AWS KMS
// stores this information in the ciphertext blob. AWS KMS cannot store metadata in
// ciphertext generated with asymmetric keys. The standard format for asymmetric
// key ciphertext does not include configurable fields.</p> </important> </li>
// </ul> <p>Unlike other AWS KMS API operations, <code>ReEncrypt</code> callers
// must have two permissions:</p> <ul> <li> <p> <code>kms:ReEncryptFrom</code>
// permission on the source CMK</p> </li> <li> <p> <code>kms:ReEncryptTo</code>
// permission on the destination CMK</p> </li> </ul> <p>To permit reencryption from
// or to a CMK, include the <code>"kms:ReEncrypt*"</code> permission in your <a
// href="https://docs.aws.amazon.com/kms/latest/developerguide/key-policies.html">key
// policy</a>. This permission is automatically included in the key policy when you
// use the console to create a CMK. But you must include it manually when you
// create a CMK programmatically or when you use the <a>PutKeyPolicy</a> operation
// to set a key policy.</p> <p>The CMK that you use for this operation must be in a
// compatible key state. For  details, see How Key State Affects Use of a Customer
// Master Key
// (https://docs.aws.amazon.com/kms/latest/developerguide/key-state.html) in the
// AWS Key Management Service Developer Guide.
func (c *Client) ReEncrypt(ctx context.Context, params *ReEncryptInput, optFns ...func(*Options)) (*ReEncryptOutput, error) {
	if params == nil {
		params = &ReEncryptInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ReEncrypt", params, optFns, addOperationReEncryptMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ReEncryptOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ReEncryptInput struct {

	// Ciphertext of the data to reencrypt.
	//
	// This member is required.
	CiphertextBlob []byte

	// A unique identifier for the CMK that is used to reencrypt the data. Specify a
	// symmetric or asymmetric CMK with a KeyUsage value of ENCRYPT_DECRYPT. To find
	// the KeyUsage value of a CMK, use the DescribeKey () operation. To specify a CMK,
	// use its key ID, Amazon Resource Name (ARN), alias name, or alias ARN. When using
	// an alias name, prefix it with "alias/". To specify a CMK in a different AWS
	// account, you must use the key ARN or alias ARN. For example:
	//
	//     * Key ID:
	// 1234abcd-12ab-34cd-56ef-1234567890ab
	//
	//     * Key ARN:
	// arn:aws:kms:us-east-2:111122223333:key/1234abcd-12ab-34cd-56ef-1234567890ab
	//
	//
	// * Alias name: alias/ExampleAlias
	//
	//     * Alias ARN:
	// arn:aws:kms:us-east-2:111122223333:alias/ExampleAlias
	//
	// To get the key ID and key
	// ARN for a CMK, use ListKeys () or DescribeKey (). To get the alias name and
	// alias ARN, use ListAliases ().
	//
	// This member is required.
	DestinationKeyId *string

	// Specifies the encryption algorithm that AWS KMS will use to reecrypt the data
	// after it has decrypted it. The default value, SYMMETRIC_DEFAULT, represents the
	// encryption algorithm used for symmetric CMKs. This parameter is required only
	// when the destination CMK is an asymmetric CMK.
	DestinationEncryptionAlgorithm types.EncryptionAlgorithmSpec

	// Specifies that encryption context to use when the reencrypting the data. A
	// destination encryption context is valid only when the destination CMK is a
	// symmetric CMK. The standard ciphertext format for asymmetric CMKs does not
	// include fields for metadata. An encryption context is a collection of non-secret
	// key-value pairs that represents additional authenticated data. When you use an
	// encryption context to encrypt data, you must specify the same (an exact
	// case-sensitive match) encryption context to decrypt the data. An encryption
	// context is optional when encrypting with a symmetric CMK, but it is highly
	// recommended. For more information, see Encryption Context
	// (https://docs.aws.amazon.com/kms/latest/developerguide/concepts.html#encrypt_context)
	// in the AWS Key Management Service Developer Guide.
	DestinationEncryptionContext map[string]*string

	// A list of grant tokens. For more information, see Grant Tokens
	// (https://docs.aws.amazon.com/kms/latest/developerguide/concepts.html#grant_token)
	// in the AWS Key Management Service Developer Guide.
	GrantTokens []*string

	// Specifies the encryption algorithm that AWS KMS will use to decrypt the
	// ciphertext before it is reencrypted. The default value, SYMMETRIC_DEFAULT,
	// represents the algorithm used for symmetric CMKs. Specify the same algorithm
	// that was used to encrypt the ciphertext. If you specify a different algorithm,
	// the decrypt attempt fails. This parameter is required only when the ciphertext
	// was encrypted under an asymmetric CMK.
	SourceEncryptionAlgorithm types.EncryptionAlgorithmSpec

	// Specifies the encryption context to use to decrypt the ciphertext. Enter the
	// same encryption context that was used to encrypt the ciphertext. An encryption
	// context is a collection of non-secret key-value pairs that represents additional
	// authenticated data. When you use an encryption context to encrypt data, you must
	// specify the same (an exact case-sensitive match) encryption context to decrypt
	// the data. An encryption context is optional when encrypting with a symmetric
	// CMK, but it is highly recommended. For more information, see Encryption Context
	// (https://docs.aws.amazon.com/kms/latest/developerguide/concepts.html#encrypt_context)
	// in the AWS Key Management Service Developer Guide.
	SourceEncryptionContext map[string]*string

	// A unique identifier for the CMK that is used to decrypt the ciphertext before it
	// reencrypts it using the destination CMK. This parameter is required only when
	// the ciphertext was encrypted under an asymmetric CMK. Otherwise, AWS KMS uses
	// the metadata that it adds to the ciphertext blob to determine which CMK was used
	// to encrypt the ciphertext. However, you can use this parameter to ensure that a
	// particular CMK (of any kind) is used to decrypt the ciphertext before it is
	// reencrypted. If you specify a KeyId value, the decrypt part of the ReEncrypt
	// operation succeeds only if the specified CMK was used to encrypt the ciphertext.
	// <p>To specify a CMK, use its key ID, Amazon Resource Name (ARN), alias name, or
	// alias ARN. When using an alias name, prefix it with <code>"alias/"</code>.</p>
	// <p>For example:</p> <ul> <li> <p>Key ID:
	// <code>1234abcd-12ab-34cd-56ef-1234567890ab</code> </p> </li> <li> <p>Key ARN:
	// <code>arn:aws:kms:us-east-2:111122223333:key/1234abcd-12ab-34cd-56ef-1234567890ab</code>
	// </p> </li> <li> <p>Alias name: <code>alias/ExampleAlias</code> </p> </li> <li>
	// <p>Alias ARN: <code>arn:aws:kms:us-east-2:111122223333:alias/ExampleAlias</code>
	// </p> </li> </ul> <p>To get the key ID and key ARN for a CMK, use <a>ListKeys</a>
	// or <a>DescribeKey</a>. To get the alias name and alias ARN, use
	// <a>ListAliases</a>.</p>
	SourceKeyId *string
}

type ReEncryptOutput struct {

	// The reencrypted data. When you use the HTTP API or the AWS CLI, the value is
	// Base64-encoded. Otherwise, it is not Base64-encoded.
	CiphertextBlob []byte

	// The encryption algorithm that was used to reencrypt the data.
	DestinationEncryptionAlgorithm types.EncryptionAlgorithmSpec

	// The Amazon Resource Name (key ARN
	// (https://docs.aws.amazon.com/kms/latest/developerguide/concepts.html#key-id-key-ARN))
	// of the CMK that was used to reencrypt the data.
	KeyId *string

	// The encryption algorithm that was used to decrypt the ciphertext before it was
	// reencrypted.
	SourceEncryptionAlgorithm types.EncryptionAlgorithmSpec

	// Unique identifier of the CMK used to originally encrypt the data.
	SourceKeyId *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationReEncryptMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpReEncrypt{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpReEncrypt{}, middleware.After)
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
	addOpReEncryptValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opReEncrypt(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opReEncrypt(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "kms",
		OperationName: "ReEncrypt",
	}
}
