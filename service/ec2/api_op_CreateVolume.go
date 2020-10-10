// Code generated by smithy-go-codegen DO NOT EDIT.

package ec2

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/ec2/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
	"time"
)

// Creates an EBS volume that can be attached to an instance in the same
// Availability Zone. The volume is created in the regional endpoint that you send
// the HTTP request to. For more information see Regions and Endpoints
// (https://docs.aws.amazon.com/general/latest/gr/rande.html). You can create a new
// empty volume or restore a volume from an EBS snapshot. Any AWS Marketplace
// product codes from the snapshot are propagated to the volume. You can create
// encrypted volumes. Encrypted volumes must be attached to instances that support
// Amazon EBS encryption. Volumes that are created from encrypted snapshots are
// also automatically encrypted. For more information, see Amazon EBS Encryption
// (https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/EBSEncryption.html) in the
// Amazon Elastic Compute Cloud User Guide. You can tag your volumes during
// creation. For more information, see Tagging Your Amazon EC2 Resources
// (https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/Using_Tags.html) in the
// Amazon Elastic Compute Cloud User Guide. For more information, see Creating an
// Amazon EBS Volume
// (https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/ebs-creating-volume.html)
// in the Amazon Elastic Compute Cloud User Guide.
func (c *Client) CreateVolume(ctx context.Context, params *CreateVolumeInput, optFns ...func(*Options)) (*CreateVolumeOutput, error) {
	if params == nil {
		params = &CreateVolumeInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateVolume", params, optFns, addOperationCreateVolumeMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateVolumeOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateVolumeInput struct {

	// The Availability Zone in which to create the volume.
	//
	// This member is required.
	AvailabilityZone *string

	// Checks whether you have the required permissions for the action, without
	// actually making the request, and provides an error response. If you have the
	// required permissions, the error response is DryRunOperation. Otherwise, it is
	// UnauthorizedOperation.
	DryRun *bool

	// Specifies whether the volume should be encrypted. The effect of setting the
	// encryption state to true depends on the volume origin (new or from a snapshot),
	// starting encryption state, ownership, and whether encryption by default is
	// enabled. For more information, see Encryption by Default
	// (https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/EBSEncryption.html#encryption-by-default)
	// in the Amazon Elastic Compute Cloud User Guide. Encrypted Amazon EBS volumes
	// must be attached to instances that support Amazon EBS encryption. For more
	// information, see Supported Instance Types
	// (https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/EBSEncryption.html#EBSEncryption_supported_instances).
	Encrypted *bool

	// The number of I/O operations per second (IOPS) to provision for the volume, with
	// a maximum ratio of 50 IOPS/GiB. Range is 100 to 64,000 IOPS for volumes in most
	// Regions. Maximum IOPS of 64,000 is guaranteed only on Nitro-based instances
	// (https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/instance-types.html#ec2-nitro-instances).
	// Other instance families guarantee performance up to 32,000 IOPS. For more
	// information, see Amazon EBS Volume Types
	// (https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/EBSVolumeTypes.html) in the
	// Amazon Elastic Compute Cloud User Guide. This parameter is valid only for
	// Provisioned IOPS SSD (io1) volumes.
	Iops *int32

	// The identifier of the AWS Key Management Service (AWS KMS) customer master key
	// (CMK) to use for Amazon EBS encryption. If this parameter is not specified, your
	// AWS managed CMK for EBS is used. If KmsKeyId is specified, the encrypted state
	// must be true. You can specify the CMK using any of the following:
	//
	//     * Key ID.
	// For example, key/1234abcd-12ab-34cd-56ef-1234567890ab.
	//
	//     * Key alias. For
	// example, alias/ExampleAlias.
	//
	//     * Key ARN. For example,
	// arn:aws:kms:us-east-1:012345678910:key/abcd1234-a123-456a-a12b-a123b4cd56ef.
	//
	//
	// * Alias ARN. For example,
	// arn:aws:kms:us-east-1:012345678910:alias/ExampleAlias.
	//
	// AWS authenticates the
	// CMK asynchronously. Therefore, if you specify an ID, alias, or ARN that is not
	// valid, the action can appear to complete, but eventually fails.
	KmsKeyId *string

	// Specifies whether to enable Amazon EBS Multi-Attach. If you enable Multi-Attach,
	// you can attach the volume to up to 16 Nitro-based instances
	// (https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/instance-types.html#ec2-nitro-instances)
	// in the same Availability Zone. For more information, see  Amazon EBS
	// Multi-Attach
	// (https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/ebs-volumes-multi.html) in
	// the Amazon Elastic Compute Cloud User Guide.
	MultiAttachEnabled *bool

	// The Amazon Resource Name (ARN) of the Outpost.
	OutpostArn *string

	// The size of the volume, in GiBs. You must specify either a snapshot ID or a
	// volume size. Constraints: 1-16,384 for gp2, 4-16,384 for io1, 500-16,384 for
	// st1, 500-16,384 for sc1, and 1-1,024 for standard. If you specify a snapshot,
	// the volume size must be equal to or larger than the snapshot size. Default: If
	// you're creating the volume from a snapshot and don't specify a volume size, the
	// default is the snapshot size.
	Size *int32

	// The snapshot from which to create the volume. You must specify either a snapshot
	// ID or a volume size.
	SnapshotId *string

	// The tags to apply to the volume during creation.
	TagSpecifications []*types.TagSpecification

	// The volume type. This can be gp2 for General Purpose SSD, io1 for Provisioned
	// IOPS SSD, st1 for Throughput Optimized HDD, sc1 for Cold HDD, or standard for
	// Magnetic volumes. Default: gp2
	VolumeType types.VolumeType
}

// Describes a volume.
type CreateVolumeOutput struct {

	// Information about the volume attachments.
	Attachments []*types.VolumeAttachment

	// The Availability Zone for the volume.
	AvailabilityZone *string

	// The time stamp when volume creation was initiated.
	CreateTime *time.Time

	// Indicates whether the volume is encrypted.
	Encrypted *bool

	// Indicates whether the volume was created using fast snapshot restore.
	FastRestored *bool

	// The number of I/O operations per second (IOPS) that the volume supports. For
	// Provisioned IOPS SSD volumes, this represents the number of IOPS that are
	// provisioned for the volume. For General Purpose SSD volumes, this represents the
	// baseline performance of the volume and the rate at which the volume accumulates
	// I/O credits for bursting. For more information, see Amazon EBS Volume Types
	// (https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/EBSVolumeTypes.html) in the
	// Amazon Elastic Compute Cloud User Guide. Constraints: Range is 100-16,000 IOPS
	// for gp2 volumes and 100 to 64,000IOPS for io1 volumes, in most Regions. The
	// maximum IOPS for io1 of 64,000 is guaranteed only on Nitro-based instances
	// (https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/instance-types.html#ec2-nitro-instances).
	// Other instance families guarantee performance up to 32,000 IOPS. Condition: This
	// parameter is required for requests to create io1 volumes; it is not used in
	// requests to create gp2, st1, sc1, or standard volumes.
	Iops *int32

	// The Amazon Resource Name (ARN) of the AWS Key Management Service (AWS KMS)
	// customer master key (CMK) that was used to protect the volume encryption key for
	// the volume.
	KmsKeyId *string

	// Indicates whether Amazon EBS Multi-Attach is enabled.
	MultiAttachEnabled *bool

	// The Amazon Resource Name (ARN) of the Outpost.
	OutpostArn *string

	// The size of the volume, in GiBs.
	Size *int32

	// The snapshot from which the volume was created, if applicable.
	SnapshotId *string

	// The volume state.
	State types.VolumeState

	// Any tags assigned to the volume.
	Tags []*types.Tag

	// The ID of the volume.
	VolumeId *string

	// The volume type. This can be gp2 for General Purpose SSD, io1 for Provisioned
	// IOPS SSD, st1 for Throughput Optimized HDD, sc1 for Cold HDD, or standard for
	// Magnetic volumes.
	VolumeType types.VolumeType

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationCreateVolumeMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsEc2query_serializeOpCreateVolume{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsEc2query_deserializeOpCreateVolume{}, middleware.After)
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
	addOpCreateVolumeValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opCreateVolume(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opCreateVolume(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "ec2",
		OperationName: "CreateVolume",
	}
}
