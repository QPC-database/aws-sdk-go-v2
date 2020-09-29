// Code generated by smithy-go-codegen DO NOT EDIT.

package types

import (
	"fmt"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/ptr"
)

// Access denied. Check your permissions.
type AccessDeniedException struct {
	Message *string
}

func (e *AccessDeniedException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *AccessDeniedException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *AccessDeniedException) ErrorCode() string             { return "AccessDeniedException" }
func (e *AccessDeniedException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }
func (e *AccessDeniedException) GetMessage() string {
	return ptr.ToString(e.Message)
}
func (e *AccessDeniedException) HasMessage() bool {
	return e.Message != nil
}

// A BatchWrite exception has occurred.
type BatchWriteException struct {
	Message *string

	Index *int32
	Type  BatchWriteExceptionType
}

func (e *BatchWriteException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *BatchWriteException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *BatchWriteException) ErrorCode() string             { return "BatchWriteException" }
func (e *BatchWriteException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }
func (e *BatchWriteException) GetIndex() int32 {
	return ptr.ToInt32(e.Index)
}
func (e *BatchWriteException) HasIndex() bool {
	return e.Index != nil
}
func (e *BatchWriteException) GetType() BatchWriteExceptionType {
	return e.Type
}
func (e *BatchWriteException) GetMessage() string {
	return ptr.ToString(e.Message)
}
func (e *BatchWriteException) HasMessage() bool {
	return e.Message != nil
}

// Cannot list the parents of a Directory () root.
type CannotListParentOfRootException struct {
	Message *string
}

func (e *CannotListParentOfRootException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *CannotListParentOfRootException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *CannotListParentOfRootException) ErrorCode() string {
	return "CannotListParentOfRootException"
}
func (e *CannotListParentOfRootException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }
func (e *CannotListParentOfRootException) GetMessage() string {
	return ptr.ToString(e.Message)
}
func (e *CannotListParentOfRootException) HasMessage() bool {
	return e.Message != nil
}

// Indicates that a Directory () could not be created due to a naming conflict.
// Choose a different name and try again.
type DirectoryAlreadyExistsException struct {
	Message *string
}

func (e *DirectoryAlreadyExistsException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *DirectoryAlreadyExistsException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *DirectoryAlreadyExistsException) ErrorCode() string {
	return "DirectoryAlreadyExistsException"
}
func (e *DirectoryAlreadyExistsException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }
func (e *DirectoryAlreadyExistsException) GetMessage() string {
	return ptr.ToString(e.Message)
}
func (e *DirectoryAlreadyExistsException) HasMessage() bool {
	return e.Message != nil
}

// A directory that has been deleted and to which access has been attempted. Note:
// The requested resource will eventually cease to exist.
type DirectoryDeletedException struct {
	Message *string
}

func (e *DirectoryDeletedException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *DirectoryDeletedException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *DirectoryDeletedException) ErrorCode() string             { return "DirectoryDeletedException" }
func (e *DirectoryDeletedException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }
func (e *DirectoryDeletedException) GetMessage() string {
	return ptr.ToString(e.Message)
}
func (e *DirectoryDeletedException) HasMessage() bool {
	return e.Message != nil
}

// An operation can only operate on a disabled directory.
type DirectoryNotDisabledException struct {
	Message *string
}

func (e *DirectoryNotDisabledException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *DirectoryNotDisabledException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *DirectoryNotDisabledException) ErrorCode() string             { return "DirectoryNotDisabledException" }
func (e *DirectoryNotDisabledException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }
func (e *DirectoryNotDisabledException) GetMessage() string {
	return ptr.ToString(e.Message)
}
func (e *DirectoryNotDisabledException) HasMessage() bool {
	return e.Message != nil
}

// Operations are only permitted on enabled directories.
type DirectoryNotEnabledException struct {
	Message *string
}

func (e *DirectoryNotEnabledException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *DirectoryNotEnabledException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *DirectoryNotEnabledException) ErrorCode() string             { return "DirectoryNotEnabledException" }
func (e *DirectoryNotEnabledException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }
func (e *DirectoryNotEnabledException) GetMessage() string {
	return ptr.ToString(e.Message)
}
func (e *DirectoryNotEnabledException) HasMessage() bool {
	return e.Message != nil
}

// A facet with the same name already exists.
type FacetAlreadyExistsException struct {
	Message *string
}

func (e *FacetAlreadyExistsException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *FacetAlreadyExistsException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *FacetAlreadyExistsException) ErrorCode() string             { return "FacetAlreadyExistsException" }
func (e *FacetAlreadyExistsException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }
func (e *FacetAlreadyExistsException) GetMessage() string {
	return ptr.ToString(e.Message)
}
func (e *FacetAlreadyExistsException) HasMessage() bool {
	return e.Message != nil
}

// Occurs when deleting a facet that contains an attribute that is a target to an
// attribute reference in a different facet.
type FacetInUseException struct {
	Message *string
}

func (e *FacetInUseException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *FacetInUseException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *FacetInUseException) ErrorCode() string             { return "FacetInUseException" }
func (e *FacetInUseException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }
func (e *FacetInUseException) GetMessage() string {
	return ptr.ToString(e.Message)
}
func (e *FacetInUseException) HasMessage() bool {
	return e.Message != nil
}

// The specified Facet () could not be found.
type FacetNotFoundException struct {
	Message *string
}

func (e *FacetNotFoundException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *FacetNotFoundException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *FacetNotFoundException) ErrorCode() string             { return "FacetNotFoundException" }
func (e *FacetNotFoundException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }
func (e *FacetNotFoundException) GetMessage() string {
	return ptr.ToString(e.Message)
}
func (e *FacetNotFoundException) HasMessage() bool {
	return e.Message != nil
}

// The Facet () that you provided was not well formed or could not be validated
// with the schema.
type FacetValidationException struct {
	Message *string
}

func (e *FacetValidationException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *FacetValidationException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *FacetValidationException) ErrorCode() string             { return "FacetValidationException" }
func (e *FacetValidationException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }
func (e *FacetValidationException) GetMessage() string {
	return ptr.ToString(e.Message)
}
func (e *FacetValidationException) HasMessage() bool {
	return e.Message != nil
}

// Indicates a failure occurred while performing a check for backward compatibility
// between the specified schema and the schema that is currently applied to the
// directory.
type IncompatibleSchemaException struct {
	Message *string
}

func (e *IncompatibleSchemaException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *IncompatibleSchemaException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *IncompatibleSchemaException) ErrorCode() string             { return "IncompatibleSchemaException" }
func (e *IncompatibleSchemaException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }
func (e *IncompatibleSchemaException) GetMessage() string {
	return ptr.ToString(e.Message)
}
func (e *IncompatibleSchemaException) HasMessage() bool {
	return e.Message != nil
}

// An object has been attempted to be attached to an object that does not have the
// appropriate attribute value.
type IndexedAttributeMissingException struct {
	Message *string
}

func (e *IndexedAttributeMissingException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *IndexedAttributeMissingException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *IndexedAttributeMissingException) ErrorCode() string {
	return "IndexedAttributeMissingException"
}
func (e *IndexedAttributeMissingException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }
func (e *IndexedAttributeMissingException) GetMessage() string {
	return ptr.ToString(e.Message)
}
func (e *IndexedAttributeMissingException) HasMessage() bool {
	return e.Message != nil
}

// Indicates a problem that must be resolved by Amazon Web Services. This might be
// a transient error in which case you can retry your request until it succeeds.
// Otherwise, go to the AWS Service Health Dashboard
// (http://status.aws.amazon.com/) site to see if there are any operational issues
// with the service.
type InternalServiceException struct {
	Message *string
}

func (e *InternalServiceException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *InternalServiceException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *InternalServiceException) ErrorCode() string             { return "InternalServiceException" }
func (e *InternalServiceException) ErrorFault() smithy.ErrorFault { return smithy.FaultServer }
func (e *InternalServiceException) GetMessage() string {
	return ptr.ToString(e.Message)
}
func (e *InternalServiceException) HasMessage() bool {
	return e.Message != nil
}

// Indicates that the provided ARN value is not valid.
type InvalidArnException struct {
	Message *string
}

func (e *InvalidArnException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *InvalidArnException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *InvalidArnException) ErrorCode() string             { return "InvalidArnException" }
func (e *InvalidArnException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }
func (e *InvalidArnException) GetMessage() string {
	return ptr.ToString(e.Message)
}
func (e *InvalidArnException) HasMessage() bool {
	return e.Message != nil
}

// Indicates that an attempt to make an attachment was invalid. For example,
// attaching two nodes with a link type that is not applicable to the nodes or
// attempting to apply a schema to a directory a second time.
type InvalidAttachmentException struct {
	Message *string
}

func (e *InvalidAttachmentException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *InvalidAttachmentException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *InvalidAttachmentException) ErrorCode() string             { return "InvalidAttachmentException" }
func (e *InvalidAttachmentException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }
func (e *InvalidAttachmentException) GetMessage() string {
	return ptr.ToString(e.Message)
}
func (e *InvalidAttachmentException) HasMessage() bool {
	return e.Message != nil
}

// An attempt to modify a Facet () resulted in an invalid schema exception.
type InvalidFacetUpdateException struct {
	Message *string
}

func (e *InvalidFacetUpdateException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *InvalidFacetUpdateException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *InvalidFacetUpdateException) ErrorCode() string             { return "InvalidFacetUpdateException" }
func (e *InvalidFacetUpdateException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }
func (e *InvalidFacetUpdateException) GetMessage() string {
	return ptr.ToString(e.Message)
}
func (e *InvalidFacetUpdateException) HasMessage() bool {
	return e.Message != nil
}

// Indicates that the NextToken value is not valid.
type InvalidNextTokenException struct {
	Message *string
}

func (e *InvalidNextTokenException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *InvalidNextTokenException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *InvalidNextTokenException) ErrorCode() string             { return "InvalidNextTokenException" }
func (e *InvalidNextTokenException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }
func (e *InvalidNextTokenException) GetMessage() string {
	return ptr.ToString(e.Message)
}
func (e *InvalidNextTokenException) HasMessage() bool {
	return e.Message != nil
}

// Occurs when any of the rule parameter keys or values are invalid.
type InvalidRuleException struct {
	Message *string
}

func (e *InvalidRuleException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *InvalidRuleException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *InvalidRuleException) ErrorCode() string             { return "InvalidRuleException" }
func (e *InvalidRuleException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }
func (e *InvalidRuleException) GetMessage() string {
	return ptr.ToString(e.Message)
}
func (e *InvalidRuleException) HasMessage() bool {
	return e.Message != nil
}

// Indicates that the provided SchemaDoc value is not valid.
type InvalidSchemaDocException struct {
	Message *string
}

func (e *InvalidSchemaDocException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *InvalidSchemaDocException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *InvalidSchemaDocException) ErrorCode() string             { return "InvalidSchemaDocException" }
func (e *InvalidSchemaDocException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }
func (e *InvalidSchemaDocException) GetMessage() string {
	return ptr.ToString(e.Message)
}
func (e *InvalidSchemaDocException) HasMessage() bool {
	return e.Message != nil
}

// Can occur for multiple reasons such as when you tag a resource that doesn’t
// exist or if you specify a higher number of tags for a resource than the allowed
// limit. Allowed limit is 50 tags per resource.
type InvalidTaggingRequestException struct {
	Message *string
}

func (e *InvalidTaggingRequestException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *InvalidTaggingRequestException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *InvalidTaggingRequestException) ErrorCode() string             { return "InvalidTaggingRequestException" }
func (e *InvalidTaggingRequestException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }
func (e *InvalidTaggingRequestException) GetMessage() string {
	return ptr.ToString(e.Message)
}
func (e *InvalidTaggingRequestException) HasMessage() bool {
	return e.Message != nil
}

// Indicates that limits are exceeded. See Limits
// (https://docs.aws.amazon.com/clouddirectory/latest/developerguide/limits.html)
// for more information.
type LimitExceededException struct {
	Message *string
}

func (e *LimitExceededException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *LimitExceededException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *LimitExceededException) ErrorCode() string             { return "LimitExceededException" }
func (e *LimitExceededException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }
func (e *LimitExceededException) GetMessage() string {
	return ptr.ToString(e.Message)
}
func (e *LimitExceededException) HasMessage() bool {
	return e.Message != nil
}

// Indicates that a link could not be created due to a naming conflict. Choose a
// different name and then try again.
type LinkNameAlreadyInUseException struct {
	Message *string
}

func (e *LinkNameAlreadyInUseException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *LinkNameAlreadyInUseException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *LinkNameAlreadyInUseException) ErrorCode() string             { return "LinkNameAlreadyInUseException" }
func (e *LinkNameAlreadyInUseException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }
func (e *LinkNameAlreadyInUseException) GetMessage() string {
	return ptr.ToString(e.Message)
}
func (e *LinkNameAlreadyInUseException) HasMessage() bool {
	return e.Message != nil
}

// Indicates that the requested operation can only operate on index objects.
type NotIndexException struct {
	Message *string
}

func (e *NotIndexException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *NotIndexException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *NotIndexException) ErrorCode() string             { return "NotIndexException" }
func (e *NotIndexException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }
func (e *NotIndexException) GetMessage() string {
	return ptr.ToString(e.Message)
}
func (e *NotIndexException) HasMessage() bool {
	return e.Message != nil
}

// Occurs when any invalid operations are performed on an object that is not a
// node, such as calling ListObjectChildren for a leaf node object.
type NotNodeException struct {
	Message *string
}

func (e *NotNodeException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *NotNodeException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *NotNodeException) ErrorCode() string             { return "NotNodeException" }
func (e *NotNodeException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }
func (e *NotNodeException) GetMessage() string {
	return ptr.ToString(e.Message)
}
func (e *NotNodeException) HasMessage() bool {
	return e.Message != nil
}

// Indicates that the requested operation can only operate on policy objects.
type NotPolicyException struct {
	Message *string
}

func (e *NotPolicyException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *NotPolicyException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *NotPolicyException) ErrorCode() string             { return "NotPolicyException" }
func (e *NotPolicyException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }
func (e *NotPolicyException) GetMessage() string {
	return ptr.ToString(e.Message)
}
func (e *NotPolicyException) HasMessage() bool {
	return e.Message != nil
}

// Indicates that the object is not attached to the index.
type ObjectAlreadyDetachedException struct {
	Message *string
}

func (e *ObjectAlreadyDetachedException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *ObjectAlreadyDetachedException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *ObjectAlreadyDetachedException) ErrorCode() string             { return "ObjectAlreadyDetachedException" }
func (e *ObjectAlreadyDetachedException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }
func (e *ObjectAlreadyDetachedException) GetMessage() string {
	return ptr.ToString(e.Message)
}
func (e *ObjectAlreadyDetachedException) HasMessage() bool {
	return e.Message != nil
}

// Indicates that the requested operation cannot be completed because the object
// has not been detached from the tree.
type ObjectNotDetachedException struct {
	Message *string
}

func (e *ObjectNotDetachedException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *ObjectNotDetachedException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *ObjectNotDetachedException) ErrorCode() string             { return "ObjectNotDetachedException" }
func (e *ObjectNotDetachedException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }
func (e *ObjectNotDetachedException) GetMessage() string {
	return ptr.ToString(e.Message)
}
func (e *ObjectNotDetachedException) HasMessage() bool {
	return e.Message != nil
}

// The specified resource could not be found.
type ResourceNotFoundException struct {
	Message *string
}

func (e *ResourceNotFoundException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *ResourceNotFoundException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *ResourceNotFoundException) ErrorCode() string             { return "ResourceNotFoundException" }
func (e *ResourceNotFoundException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }
func (e *ResourceNotFoundException) GetMessage() string {
	return ptr.ToString(e.Message)
}
func (e *ResourceNotFoundException) HasMessage() bool {
	return e.Message != nil
}

// Occurs when a conflict with a previous successful write is detected. For
// example, if a write operation occurs on an object and then an attempt is made to
// read the object using “SERIALIZABLE” consistency, this exception may result.
// This generally occurs when the previous write did not have time to propagate to
// the host serving the current request. A retry (with appropriate backoff logic)
// is the recommended response to this exception.
type RetryableConflictException struct {
	Message *string
}

func (e *RetryableConflictException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *RetryableConflictException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *RetryableConflictException) ErrorCode() string             { return "RetryableConflictException" }
func (e *RetryableConflictException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }
func (e *RetryableConflictException) GetMessage() string {
	return ptr.ToString(e.Message)
}
func (e *RetryableConflictException) HasMessage() bool {
	return e.Message != nil
}

// Indicates that a schema could not be created due to a naming conflict. Please
// select a different name and then try again.
type SchemaAlreadyExistsException struct {
	Message *string
}

func (e *SchemaAlreadyExistsException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *SchemaAlreadyExistsException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *SchemaAlreadyExistsException) ErrorCode() string             { return "SchemaAlreadyExistsException" }
func (e *SchemaAlreadyExistsException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }
func (e *SchemaAlreadyExistsException) GetMessage() string {
	return ptr.ToString(e.Message)
}
func (e *SchemaAlreadyExistsException) HasMessage() bool {
	return e.Message != nil
}

// Indicates that a schema is already published.
type SchemaAlreadyPublishedException struct {
	Message *string
}

func (e *SchemaAlreadyPublishedException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *SchemaAlreadyPublishedException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *SchemaAlreadyPublishedException) ErrorCode() string {
	return "SchemaAlreadyPublishedException"
}
func (e *SchemaAlreadyPublishedException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }
func (e *SchemaAlreadyPublishedException) GetMessage() string {
	return ptr.ToString(e.Message)
}
func (e *SchemaAlreadyPublishedException) HasMessage() bool {
	return e.Message != nil
}

// The object could not be deleted because links still exist. Remove the links and
// then try the operation again.
type StillContainsLinksException struct {
	Message *string
}

func (e *StillContainsLinksException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *StillContainsLinksException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *StillContainsLinksException) ErrorCode() string             { return "StillContainsLinksException" }
func (e *StillContainsLinksException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }
func (e *StillContainsLinksException) GetMessage() string {
	return ptr.ToString(e.Message)
}
func (e *StillContainsLinksException) HasMessage() bool {
	return e.Message != nil
}

// Indicates that the requested index type is not supported.
type UnsupportedIndexTypeException struct {
	Message *string
}

func (e *UnsupportedIndexTypeException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *UnsupportedIndexTypeException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *UnsupportedIndexTypeException) ErrorCode() string             { return "UnsupportedIndexTypeException" }
func (e *UnsupportedIndexTypeException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }
func (e *UnsupportedIndexTypeException) GetMessage() string {
	return ptr.ToString(e.Message)
}
func (e *UnsupportedIndexTypeException) HasMessage() bool {
	return e.Message != nil
}

// Indicates that your request is malformed in some manner. See the exception
// message.
type ValidationException struct {
	Message *string
}

func (e *ValidationException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *ValidationException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *ValidationException) ErrorCode() string             { return "ValidationException" }
func (e *ValidationException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }
func (e *ValidationException) GetMessage() string {
	return ptr.ToString(e.Message)
}
func (e *ValidationException) HasMessage() bool {
	return e.Message != nil
}