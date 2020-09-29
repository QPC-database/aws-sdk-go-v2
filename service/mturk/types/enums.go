// Code generated by smithy-go-codegen DO NOT EDIT.

package types

type AssignmentStatus string

// Enum values for AssignmentStatus
const (
	AssignmentStatusSubmitted AssignmentStatus = "Submitted"
	AssignmentStatusApproved  AssignmentStatus = "Approved"
	AssignmentStatusRejected  AssignmentStatus = "Rejected"
)

type Comparator string

// Enum values for Comparator
const (
	ComparatorLessthan             Comparator = "LessThan"
	ComparatorLessthanorequalto    Comparator = "LessThanOrEqualTo"
	ComparatorGreaterthan          Comparator = "GreaterThan"
	ComparatorGreaterthanorequalto Comparator = "GreaterThanOrEqualTo"
	ComparatorEqualto              Comparator = "EqualTo"
	ComparatorNotequalto           Comparator = "NotEqualTo"
	ComparatorExists               Comparator = "Exists"
	ComparatorDoesnotexist         Comparator = "DoesNotExist"
	ComparatorIn                   Comparator = "In"
	ComparatorNotin                Comparator = "NotIn"
)

type EventType string

// Enum values for EventType
const (
	EventTypeAssignmentaccepted  EventType = "AssignmentAccepted"
	EventTypeAssignmentabandoned EventType = "AssignmentAbandoned"
	EventTypeAssignmentreturned  EventType = "AssignmentReturned"
	EventTypeAssignmentsubmitted EventType = "AssignmentSubmitted"
	EventTypeAssignmentrejected  EventType = "AssignmentRejected"
	EventTypeAssignmentapproved  EventType = "AssignmentApproved"
	EventTypeHitcreated          EventType = "HITCreated"
	EventTypeHitexpired          EventType = "HITExpired"
	EventTypeHitreviewable       EventType = "HITReviewable"
	EventTypeHitextended         EventType = "HITExtended"
	EventTypeHitdisposed         EventType = "HITDisposed"
	EventTypePing                EventType = "Ping"
)

type HITAccessActions string

// Enum values for HITAccessActions
const (
	HITAccessActionsAccept                   HITAccessActions = "Accept"
	HITAccessActionsPreviewandaccept         HITAccessActions = "PreviewAndAccept"
	HITAccessActionsDiscoverpreviewandaccept HITAccessActions = "DiscoverPreviewAndAccept"
)

type HITReviewStatus string

// Enum values for HITReviewStatus
const (
	HITReviewStatusNotreviewed           HITReviewStatus = "NotReviewed"
	HITReviewStatusMarkedforreview       HITReviewStatus = "MarkedForReview"
	HITReviewStatusReviewedappropriate   HITReviewStatus = "ReviewedAppropriate"
	HITReviewStatusReviewedinappropriate HITReviewStatus = "ReviewedInappropriate"
)

type HITStatus string

// Enum values for HITStatus
const (
	HITStatusAssignable   HITStatus = "Assignable"
	HITStatusUnassignable HITStatus = "Unassignable"
	HITStatusReviewable   HITStatus = "Reviewable"
	HITStatusReviewing    HITStatus = "Reviewing"
	HITStatusDisposed     HITStatus = "Disposed"
)

type NotificationTransport string

// Enum values for NotificationTransport
const (
	NotificationTransportEmail NotificationTransport = "Email"
	NotificationTransportSqs   NotificationTransport = "SQS"
	NotificationTransportSns   NotificationTransport = "SNS"
)

type NotifyWorkersFailureCode string

// Enum values for NotifyWorkersFailureCode
const (
	NotifyWorkersFailureCodeSoftfailure NotifyWorkersFailureCode = "SoftFailure"
	NotifyWorkersFailureCodeHardfailure NotifyWorkersFailureCode = "HardFailure"
)

type QualificationStatus string

// Enum values for QualificationStatus
const (
	QualificationStatusGranted QualificationStatus = "Granted"
	QualificationStatusRevoked QualificationStatus = "Revoked"
)

type QualificationTypeStatus string

// Enum values for QualificationTypeStatus
const (
	QualificationTypeStatusActive   QualificationTypeStatus = "Active"
	QualificationTypeStatusInactive QualificationTypeStatus = "Inactive"
)

type ReviewableHITStatus string

// Enum values for ReviewableHITStatus
const (
	ReviewableHITStatusReviewable ReviewableHITStatus = "Reviewable"
	ReviewableHITStatusReviewing  ReviewableHITStatus = "Reviewing"
)

type ReviewActionStatus string

// Enum values for ReviewActionStatus
const (
	ReviewActionStatusIntended  ReviewActionStatus = "Intended"
	ReviewActionStatusSucceeded ReviewActionStatus = "Succeeded"
	ReviewActionStatusFailed    ReviewActionStatus = "Failed"
	ReviewActionStatusCancelled ReviewActionStatus = "Cancelled"
)

type ReviewPolicyLevel string

// Enum values for ReviewPolicyLevel
const (
	ReviewPolicyLevelAssignment ReviewPolicyLevel = "Assignment"
	ReviewPolicyLevelHit        ReviewPolicyLevel = "HIT"
)