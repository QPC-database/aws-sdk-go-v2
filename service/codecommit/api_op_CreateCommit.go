// Code generated by smithy-go-codegen DO NOT EDIT.

package codecommit

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/codecommit/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Creates a commit for a repository on the tip of a specified branch.
func (c *Client) CreateCommit(ctx context.Context, params *CreateCommitInput, optFns ...func(*Options)) (*CreateCommitOutput, error) {
	if params == nil {
		params = &CreateCommitInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateCommit", params, optFns, addOperationCreateCommitMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateCommitOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateCommitInput struct {

	// The name of the branch where you create the commit.
	//
	// This member is required.
	BranchName *string

	// The name of the repository where you create the commit.
	//
	// This member is required.
	RepositoryName *string

	// The name of the author who created the commit. This information is used as both
	// the author and committer for the commit.
	AuthorName *string

	// The commit message you want to include in the commit. Commit messages are
	// limited to 256 KB. If no message is specified, a default message is used.
	CommitMessage *string

	// The files to delete in this commit. These files still exist in earlier commits.
	DeleteFiles []*types.DeleteFileEntry

	// The email address of the person who created the commit.
	Email *string

	// If the commit contains deletions, whether to keep a folder or folder structure
	// if the changes leave the folders empty. If true, a ..gitkeep file is created for
	// empty folders. The default is false.
	KeepEmptyFolders *bool

	// The ID of the commit that is the parent of the commit you create. Not required
	// if this is an empty repository.
	ParentCommitId *string

	// The files to add or update in this commit.
	PutFiles []*types.PutFileEntry

	// The file modes to update for files in this commit.
	SetFileModes []*types.SetFileModeEntry
}

type CreateCommitOutput struct {

	// The full commit ID of the commit that contains your committed file changes.
	CommitId *string

	// The files added as part of the committed file changes.
	FilesAdded []*types.FileMetadata

	// The files deleted as part of the committed file changes.
	FilesDeleted []*types.FileMetadata

	// The files updated as part of the commited file changes.
	FilesUpdated []*types.FileMetadata

	// The full SHA-1 pointer of the tree information for the commit that contains the
	// commited file changes.
	TreeId *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationCreateCommitMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpCreateCommit{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpCreateCommit{}, middleware.After)
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
	addOpCreateCommitValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opCreateCommit(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opCreateCommit(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "codecommit",
		OperationName: "CreateCommit",
	}
}
