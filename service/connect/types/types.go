// Code generated by smithy-go-codegen DO NOT EDIT.

package types

import (
	"time"
)

// A toggle for an individual feature at the instance level.
type Attribute struct {

	// The type of attribute.
	AttributeType InstanceAttributeType

	// The value of the attribute.
	Value *string
}

// A chat message.
type ChatMessage struct {

	// The content of the chat message.
	//
	// This member is required.
	Content *string

	// The type of the content. Supported types are text and plain.
	//
	// This member is required.
	ContentType *string
}

// Contains information about a contact flow.
type ContactFlow struct {

	// The Amazon Resource Name (ARN) of the contact flow.
	Arn *string

	// The content of the contact flow.
	Content *string

	// The description of the contact flow.
	Description *string

	// The identifier of the contact flow.
	Id *string

	// The name of the contact flow.
	Name *string

	// One or more tags.
	Tags map[string]string

	// The type of the contact flow. For descriptions of the available types, see
	// Choose a Contact Flow Type
	// (https://docs.aws.amazon.com/connect/latest/adminguide/create-contact-flow.html#contact-flow-types)
	// in the Amazon Connect Administrator Guide.
	Type ContactFlowType
}

// Contains summary information about a contact flow. You can also create and
// update contact flows using the Amazon Connect Flow language
// (https://docs.aws.amazon.com/connect/latest/adminguide/flow-language.html).
type ContactFlowSummary struct {

	// The Amazon Resource Name (ARN) of the contact flow.
	Arn *string

	// The type of contact flow.
	ContactFlowType ContactFlowType

	// The identifier of the contact flow.
	Id *string

	// The name of the contact flow.
	Name *string
}

// Contains credentials to use for federation.
type Credentials struct {

	// An access token generated for a federated user to access Amazon Connect.
	AccessToken *string

	// A token generated with an expiration time for the session a user is logged in to
	// Amazon Connect.
	AccessTokenExpiration *time.Time

	// Renews a token generated for a user to access the Amazon Connect instance.
	RefreshToken *string

	// Renews the expiration timer for a generated token.
	RefreshTokenExpiration *time.Time
}

// Contains information about a real-time metric. For a description of each metric,
// see Real-time Metrics Definitions
// (https://docs.aws.amazon.com/connect/latest/adminguide/real-time-metrics-definitions.html)
// in the Amazon Connect Administrator Guide.
type CurrentMetric struct {

	// The name of the metric.
	Name CurrentMetricName

	// The unit for the metric.
	Unit Unit
}

// Contains the data for a real-time metric.
type CurrentMetricData struct {

	// Information about the metric.
	Metric *CurrentMetric

	// The value of the metric.
	Value *float64
}

// Contains information about a set of real-time metrics.
type CurrentMetricResult struct {

	// The set of metrics.
	Collections []CurrentMetricData

	// The dimensions for the metrics.
	Dimensions *Dimensions
}

// Contains information about the dimensions for a set of metrics.
type Dimensions struct {

	// The channel used for grouping and filters.
	Channel Channel

	// Information about the queue for which metrics are returned.
	Queue *QueueReference
}

// The encryption configuration.
type EncryptionConfig struct {

	// The type of encryption.
	//
	// This member is required.
	EncryptionType EncryptionType

	// The identifier of the encryption key.
	//
	// This member is required.
	KeyId *string
}

// Contains the filter to apply when retrieving metrics.
type Filters struct {

	// The channel to use to filter the metrics.
	Channels []Channel

	// The queues to use to filter the metrics. You can specify up to 100 queues per
	// request.
	Queues []string
}

// Contains information about a hierarchy group.
type HierarchyGroup struct {

	// The Amazon Resource Name (ARN) of the hierarchy group.
	Arn *string

	// Information about the levels in the hierarchy group.
	HierarchyPath *HierarchyPath

	// The identifier of the hierarchy group.
	Id *string

	// The identifier of the level in the hierarchy group.
	LevelId *string

	// The name of the hierarchy group.
	Name *string
}

// Contains summary information about a hierarchy group.
type HierarchyGroupSummary struct {

	// The Amazon Resource Name (ARN) of the hierarchy group.
	Arn *string

	// The identifier of the hierarchy group.
	Id *string

	// The name of the hierarchy group.
	Name *string
}

// Contains information about a hierarchy level.
type HierarchyLevel struct {

	// The Amazon Resource Name (ARN) of the hierarchy level.
	Arn *string

	// The identifier of the hierarchy level.
	Id *string

	// The name of the hierarchy level.
	Name *string
}

// Contains information about the hierarchy level to update.
type HierarchyLevelUpdate struct {

	// The name of the user hierarchy level. Must not be more than 50 characters.
	//
	// This member is required.
	Name *string
}

// Contains information about the levels of a hierarchy group.
type HierarchyPath struct {

	// Information about level five.
	LevelFive *HierarchyGroupSummary

	// Information about level four.
	LevelFour *HierarchyGroupSummary

	// Information about level one.
	LevelOne *HierarchyGroupSummary

	// Information about level three.
	LevelThree *HierarchyGroupSummary

	// Information about level two.
	LevelTwo *HierarchyGroupSummary
}

// Contains information about a hierarchy structure.
type HierarchyStructure struct {

	// Information about level five.
	LevelFive *HierarchyLevel

	// Information about level four.
	LevelFour *HierarchyLevel

	// Information about level one.
	LevelOne *HierarchyLevel

	// Information about level three.
	LevelThree *HierarchyLevel

	// Information about level two.
	LevelTwo *HierarchyLevel
}

// Contains information about the level hierarchy to update.
type HierarchyStructureUpdate struct {

	// The update for level five.
	LevelFive *HierarchyLevelUpdate

	// The update for level four.
	LevelFour *HierarchyLevelUpdate

	// The update for level one.
	LevelOne *HierarchyLevelUpdate

	// The update for level three.
	LevelThree *HierarchyLevelUpdate

	// The update for level two.
	LevelTwo *HierarchyLevelUpdate
}

// Contains information about a historical metric. For a description of each
// metric, see Historical Metrics Definitions
// (https://docs.aws.amazon.com/connect/latest/adminguide/historical-metrics-definitions.html)
// in the Amazon Connect Administrator Guide.
type HistoricalMetric struct {

	// The name of the metric.
	Name HistoricalMetricName

	// The statistic for the metric.
	Statistic Statistic

	// The threshold for the metric, used with service level metrics.
	Threshold *Threshold

	// The unit for the metric.
	Unit Unit
}

// Contains the data for a historical metric.
type HistoricalMetricData struct {

	// Information about the metric.
	Metric *HistoricalMetric

	// The value of the metric.
	Value *float64
}

// Contains information about the historical metrics retrieved.
type HistoricalMetricResult struct {

	// The set of metrics.
	Collections []HistoricalMetricData

	// The dimension for the metrics.
	Dimensions *Dimensions
}

// Information about of the hours of operation.
type HoursOfOperation struct {

	// Configuration information for the hours of operation.
	Config []HoursOfOperationConfig

	// The description for the hours of operation.
	Description *string

	// The Amazon Resource Name (ARN) for the hours of operation.
	HoursOfOperationArn *string

	// The identifier for the hours of operation.
	HoursOfOperationId *string

	// The name for the hours of operation.
	Name *string

	// One or more tags.
	Tags map[string]string

	// The time zone for the hours of operation.
	TimeZone *string
}

// Contains information about the hours of operation.
type HoursOfOperationConfig struct {

	// The day that the hours of operation applies to.
	Day HoursOfOperationDays

	// The end time that your contact center is closes.
	EndTime *HoursOfOperationTimeSlice

	// The start time that your contact center is open.
	StartTime *HoursOfOperationTimeSlice
}

// Contains summary information about hours of operation for a contact center.
type HoursOfOperationSummary struct {

	// The Amazon Resource Name (ARN) of the hours of operation.
	Arn *string

	// The identifier of the hours of operation.
	Id *string

	// The name of the hours of operation.
	Name *string
}

// The start time or end time for an hours of operation.
type HoursOfOperationTimeSlice struct {

	// The hours.
	Hours int32

	// The minutes.
	Minutes int32
}

// The Amazon Connect instance.
type Instance struct {

	// The Amazon Resource Name (ARN) of the instance.
	Arn *string

	// When the instance was created.
	CreatedTime *time.Time

	// The identifier of the Amazon Connect instance.
	Id *string

	// The identity management type.
	IdentityManagementType DirectoryType

	// Whether inbound calls are enabled.
	InboundCallsEnabled *bool

	// The alias of instance.
	InstanceAlias *string

	// The state of the instance.
	InstanceStatus InstanceStatus

	// Whether outbound calls are enabled.
	OutboundCallsEnabled *bool

	// The service role of the instance.
	ServiceRole *string

	// Relevant details why the instance was not successfully created.
	StatusReason *InstanceStatusReason
}

// Relevant details why the instance was not successfully created.
type InstanceStatusReason struct {

	// The message.
	Message *string
}

// The storage configuration for the instance.
type InstanceStorageConfig struct {

	// A valid storage type.
	//
	// This member is required.
	StorageType StorageType

	// The existing association identifier that uniquely identifies the resource type
	// and storage config for the given instance ID.
	AssociationId *string

	// The configuration of the Kinesis Firehose delivery stream.
	KinesisFirehoseConfig *KinesisFirehoseConfig

	// The configuration of the Kinesis data stream.
	KinesisStreamConfig *KinesisStreamConfig

	// The configuration of the Kinesis video stream.
	KinesisVideoStreamConfig *KinesisVideoStreamConfig

	// The S3 bucket configuration.
	S3Config *S3Config
}

// Information about the instance.
type InstanceSummary struct {

	// The Amazon Resource Name (ARN) of the instance.
	Arn *string

	// When the instance was created.
	CreatedTime *time.Time

	// The identifier of the instance.
	Id *string

	// The identity management type of the instance.
	IdentityManagementType DirectoryType

	// Whether inbound calls are enabled.
	InboundCallsEnabled *bool

	// The alias of the instance.
	InstanceAlias *string

	// The state of the instance.
	InstanceStatus InstanceStatus

	// Whether outbound calls are enabled.
	OutboundCallsEnabled *bool

	// The service role of the instance.
	ServiceRole *string
}

// Contains summary information about the associated AppIntegrations.
type IntegrationAssociationSummary struct {

	// The identifier of the Amazon Connect instance.
	InstanceId *string

	// The Amazon Resource Name (ARN) for the AppIntegration.
	IntegrationArn *string

	// The Amazon Resource Name (ARN) for the AppIntegration association.
	IntegrationAssociationArn *string

	// The identifier for the AppIntegration association.
	IntegrationAssociationId *string

	// The integration type.
	IntegrationType IntegrationType

	// The user-provided, friendly name for the external application.
	SourceApplicationName *string

	// The URL for the external application.
	SourceApplicationUrl *string

	// The name of the source.
	SourceType SourceType
}

// Configuration information of a Kinesis Data Firehose delivery stream.
type KinesisFirehoseConfig struct {

	// The Amazon Resource Name (ARN) of the delivery stream.
	//
	// This member is required.
	FirehoseArn *string
}

// Configuration information of a Kinesis data stream.
type KinesisStreamConfig struct {

	// The Amazon Resource Name (ARN) of the data stream.
	//
	// This member is required.
	StreamArn *string
}

// Configuration information of a Kinesis video stream.
type KinesisVideoStreamConfig struct {

	// The encryption configuration.
	//
	// This member is required.
	EncryptionConfig *EncryptionConfig

	// The prefix of the video stream.
	//
	// This member is required.
	Prefix *string

	// The number of hours data is retained in the stream. Kinesis Video Streams
	// retains the data in a data store that is associated with the stream. The default
	// value is 0, indicating that the stream does not persist data.
	//
	// This member is required.
	RetentionPeriodHours int32
}

// Configuration information of an Amazon Lex bot.
type LexBot struct {

	// The Region that the Amazon Lex bot was created in.
	LexRegion *string

	// The name of the Amazon Lex bot.
	Name *string
}

// Contains information about which channels are supported, and how many contacts
// an agent can have on a channel simultaneously.
type MediaConcurrency struct {

	// The channels that agents can handle in the Contact Control Panel (CCP).
	//
	// This member is required.
	Channel Channel

	// The number of contacts an agent can have on a channel simultaneously. Valid
	// Range for VOICE: Minimum value of 1. Maximum value of 1. Valid Range for CHAT:
	// Minimum value of 1. Maximum value of 5. Valid Range for TASK: Minimum value of
	// 1. Maximum value of 10.
	//
	// This member is required.
	Concurrency int32
}

// The outbound caller ID name, number, and outbound whisper flow.
type OutboundCallerConfig struct {

	// The caller ID name.
	OutboundCallerIdName *string

	// The caller ID number.
	OutboundCallerIdNumberId *string

	// The outbound whisper flow to be used during an outbound call.
	OutboundFlowId *string
}

// The customer's details.
type ParticipantDetails struct {

	// Display name of the participant.
	//
	// This member is required.
	DisplayName *string
}

// Contains information about a phone number for a quick connect.
type PhoneNumberQuickConnectConfig struct {

	// The phone number in E.164 format.
	//
	// This member is required.
	PhoneNumber *string
}

// Contains summary information about a phone number for a contact center.
type PhoneNumberSummary struct {

	// The Amazon Resource Name (ARN) of the phone number.
	Arn *string

	// The identifier of the phone number.
	Id *string

	// The phone number.
	PhoneNumber *string

	// The ISO country code.
	PhoneNumberCountryCode PhoneNumberCountryCode

	// The type of phone number.
	PhoneNumberType PhoneNumberType
}

// Information about a problem detail.
type ProblemDetail struct {

	// The problem detail's message.
	Message *string
}

// Contains information about the prompt.
type PromptSummary struct {

	// The Amazon Resource Name (ARN) of the prompt.
	Arn *string

	// The identifier of the prompt.
	Id *string

	// The name of the prompt.
	Name *string
}

// Contains information about a queue.
type Queue struct {

	// The description of the queue.
	Description *string

	// The identifier for the hours of operation.
	HoursOfOperationId *string

	// The maximum number of contacts that can be in the queue before it is considered
	// full.
	MaxContacts int32

	// The name of the queue.
	Name *string

	// The outbound caller ID name, number, and outbound whisper flow.
	OutboundCallerConfig *OutboundCallerConfig

	// The Amazon Resource Name (ARN) for the queue.
	QueueArn *string

	// The identifier for the queue.
	QueueId *string

	// The status of the queue.
	Status QueueStatus

	// One or more tags.
	Tags map[string]string
}

// Contains information about a queue for a quick connect. The contact flow must be
// of type Transfer to Queue.
type QueueQuickConnectConfig struct {

	// The identifier of the contact flow.
	//
	// This member is required.
	ContactFlowId *string

	// The identifier for the queue.
	//
	// This member is required.
	QueueId *string
}

// Contains information about a queue resource for which metrics are returned.
type QueueReference struct {

	// The Amazon Resource Name (ARN) of the queue.
	Arn *string

	// The identifier of the queue.
	Id *string
}

// Contains summary information about a queue.
type QueueSummary struct {

	// The Amazon Resource Name (ARN) of the queue.
	Arn *string

	// The identifier of the queue.
	Id *string

	// The name of the queue.
	Name *string

	// The type of queue.
	QueueType QueueType
}

// Contains information about a quick connect.
type QuickConnect struct {

	// The description.
	Description *string

	// The name of the quick connect.
	Name *string

	// The Amazon Resource Name (ARN) of the quick connect.
	QuickConnectARN *string

	// Contains information about the quick connect.
	QuickConnectConfig *QuickConnectConfig

	// The identifier for the quick connect.
	QuickConnectId *string

	// One or more tags.
	Tags map[string]string
}

// Contains configuration settings for a quick connect.
type QuickConnectConfig struct {

	// The type of quick connect. In the Amazon Connect console, when you create a
	// quick connect, you are prompted to assign one of the following types: Agent
	// (USER), External (PHONE_NUMBER), or Queue (QUEUE).
	//
	// This member is required.
	QuickConnectType QuickConnectType

	// The phone configuration. This is required only if QuickConnectType is
	// PHONE_NUMBER.
	PhoneConfig *PhoneNumberQuickConnectConfig

	// The queue configuration. This is required only if QuickConnectType is QUEUE.
	QueueConfig *QueueQuickConnectConfig

	// The user configuration. This is required only if QuickConnectType is USER.
	UserConfig *UserQuickConnectConfig
}

// Contains summary information about a quick connect.
type QuickConnectSummary struct {

	// The Amazon Resource Name (ARN) of the quick connect.
	Arn *string

	// The identifier for the quick connect.
	Id *string

	// The name of the quick connect.
	Name *string

	// The type of quick connect. In the Amazon Connect console, when you create a
	// quick connect, you are prompted to assign one of the following types: Agent
	// (USER), External (PHONE_NUMBER), or Queue (QUEUE).
	QuickConnectType QuickConnectType
}

// A link that an agent selects to complete a given task. You can have up to 4,096
// UTF-8 bytes across all references for a contact.
type Reference struct {

	// A valid URL.
	//
	// This member is required.
	Type ReferenceType

	// A formatted URL that displays to an agent in the Contact Control Panel (CCP)
	//
	// This member is required.
	Value *string
}

// Contains information about a routing profile.
type RoutingProfile struct {

	// The identifier of the default outbound queue for this routing profile.
	DefaultOutboundQueueId *string

	// The description of the routing profile.
	Description *string

	// The identifier of the Amazon Connect instance.
	InstanceId *string

	// The channels agents can handle in the Contact Control Panel (CCP) for this
	// routing profile.
	MediaConcurrencies []MediaConcurrency

	// The name of the routing profile.
	Name *string

	// The Amazon Resource Name (ARN) of the routing profile.
	RoutingProfileArn *string

	// The identifier of the routing profile.
	RoutingProfileId *string

	// One or more tags.
	Tags map[string]string
}

// Contains information about the queue and channel for which priority and delay
// can be set.
type RoutingProfileQueueConfig struct {

	// The delay, in seconds, a contact should be in the queue before they are routed
	// to an available agent. For more information, see Queues: priority and delay
	// (https://docs.aws.amazon.com/connect/latest/adminguide/concepts-routing-profiles-priority.html)
	// in the Amazon Connect Administrator Guide.
	//
	// This member is required.
	Delay int32

	// The order in which contacts are to be handled for the queue. For more
	// information, see Queues: priority and delay
	// (https://docs.aws.amazon.com/connect/latest/adminguide/concepts-routing-profiles-priority.html).
	//
	// This member is required.
	Priority int32

	// Contains information about a queue resource.
	//
	// This member is required.
	QueueReference *RoutingProfileQueueReference
}

// Contains summary information about a routing profile queue.
type RoutingProfileQueueConfigSummary struct {

	// The channels this queue supports.
	//
	// This member is required.
	Channel Channel

	// The delay, in seconds, that a contact should be in the queue before they are
	// routed to an available agent. For more information, see Queues: priority and
	// delay
	// (https://docs.aws.amazon.com/connect/latest/adminguide/concepts-routing-profiles-priority.html)
	// in the Amazon Connect Administrator Guide.
	//
	// This member is required.
	Delay int32

	// The order in which contacts are to be handled for the queue. For more
	// information, see Queues: priority and delay
	// (https://docs.aws.amazon.com/connect/latest/adminguide/concepts-routing-profiles-priority.html).
	//
	// This member is required.
	Priority int32

	// The Amazon Resource Name (ARN) of the queue.
	//
	// This member is required.
	QueueArn *string

	// The identifier for the queue.
	//
	// This member is required.
	QueueId *string

	// The name of the queue.
	//
	// This member is required.
	QueueName *string
}

// Contains the channel and queue identifier for a routing profile.
type RoutingProfileQueueReference struct {

	// The channels agents can handle in the Contact Control Panel (CCP) for this
	// routing profile.
	//
	// This member is required.
	Channel Channel

	// The identifier for the queue.
	//
	// This member is required.
	QueueId *string
}

// Contains summary information about a routing profile.
type RoutingProfileSummary struct {

	// The Amazon Resource Name (ARN) of the routing profile.
	Arn *string

	// The identifier of the routing profile.
	Id *string

	// The name of the routing profile.
	Name *string
}

// Information about the Amazon Simple Storage Service (Amazon S3) storage type.
type S3Config struct {

	// The S3 bucket name.
	//
	// This member is required.
	BucketName *string

	// The S3 bucket prefix.
	//
	// This member is required.
	BucketPrefix *string

	// The Amazon S3 encryption configuration.
	EncryptionConfig *EncryptionConfig
}

// Configuration information of the security key.
type SecurityKey struct {

	// The existing association identifier that uniquely identifies the resource type
	// and storage config for the given instance ID.
	AssociationId *string

	// When the security key was created.
	CreationTime *time.Time

	// The key of the security key.
	Key *string
}

// Contains information about a security profile.
type SecurityProfileSummary struct {

	// The Amazon Resource Name (ARN) of the security profile.
	Arn *string

	// The identifier of the security profile.
	Id *string

	// The name of the security profile.
	Name *string
}

// Contains information about the threshold for service level metrics.
type Threshold struct {

	// The type of comparison. Only "less than" (LT) comparisons are supported.
	Comparison Comparison

	// The threshold value to compare.
	ThresholdValue *float64
}

// Contains the use case.
type UseCase struct {

	// The Amazon Resource Name (ARN) for the use case.
	UseCaseArn *string

	// The identifier for the use case.
	UseCaseId *string

	// The type of use case to associate to the AppIntegration association. Each
	// AppIntegration association can have only one of each use case type.
	UseCaseType UseCaseType
}

// Contains information about a user account for a Amazon Connect instance.
type User struct {

	// The Amazon Resource Name (ARN) of the user account.
	Arn *string

	// The identifier of the user account in the directory used for identity
	// management.
	DirectoryUserId *string

	// The identifier of the hierarchy group for the user.
	HierarchyGroupId *string

	// The identifier of the user account.
	Id *string

	// Information about the user identity.
	IdentityInfo *UserIdentityInfo

	// Information about the phone configuration for the user.
	PhoneConfig *UserPhoneConfig

	// The identifier of the routing profile for the user.
	RoutingProfileId *string

	// The identifiers of the security profiles for the user.
	SecurityProfileIds []string

	// The tags.
	Tags map[string]string

	// The user name assigned to the user account.
	Username *string
}

// Contains information about the identity of a user.
type UserIdentityInfo struct {

	// The email address. If you are using SAML for identity management and include
	// this parameter, an error is returned.
	Email *string

	// The first name. This is required if you are using Amazon Connect or SAML for
	// identity management.
	FirstName *string

	// The last name. This is required if you are using Amazon Connect or SAML for
	// identity management.
	LastName *string
}

// Contains information about the phone configuration settings for a user.
type UserPhoneConfig struct {

	// The phone type.
	//
	// This member is required.
	PhoneType PhoneType

	// The After Call Work (ACW) timeout setting, in seconds.
	AfterContactWorkTimeLimit int32

	// The Auto accept setting.
	AutoAccept bool

	// The phone number for the user's desk phone.
	DeskPhoneNumber *string
}

// Contains information about the quick connect configuration settings for a user.
// The contact flow must be of type Transfer to Agent.
type UserQuickConnectConfig struct {

	// The identifier of the contact flow.
	//
	// This member is required.
	ContactFlowId *string

	// The identifier of the user.
	//
	// This member is required.
	UserId *string
}

// Contains summary information about a user.
type UserSummary struct {

	// The Amazon Resource Name (ARN) of the user account.
	Arn *string

	// The identifier of the user account.
	Id *string

	// The Amazon Connect user name of the user account.
	Username *string
}

// Contains information about the recording configuration settings.
type VoiceRecordingConfiguration struct {

	// Identifies which track is being recorded.
	VoiceRecordingTrack VoiceRecordingTrack
}
