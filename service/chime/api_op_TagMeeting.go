// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package chime

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
	"github.com/aws/aws-sdk-go-v2/private/protocol/restjson"
)

type TagMeetingInput struct {
	_ struct{} `type:"structure"`

	// The Amazon Chime SDK meeting ID.
	//
	// MeetingId is a required field
	MeetingId *string `location:"uri" locationName:"meetingId" type:"string" required:"true"`

	// The tag key-value pairs.
	//
	// Tags is a required field
	Tags []Tag `min:"1" type:"list" required:"true"`
}

// String returns the string representation
func (s TagMeetingInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *TagMeetingInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "TagMeetingInput"}

	if s.MeetingId == nil {
		invalidParams.Add(aws.NewErrParamRequired("MeetingId"))
	}

	if s.Tags == nil {
		invalidParams.Add(aws.NewErrParamRequired("Tags"))
	}
	if s.Tags != nil && len(s.Tags) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("Tags", 1))
	}
	if s.Tags != nil {
		for i, v := range s.Tags {
			if err := v.Validate(); err != nil {
				invalidParams.AddNested(fmt.Sprintf("%s[%v]", "Tags", i), err.(aws.ErrInvalidParams))
			}
		}
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s TagMeetingInput) MarshalFields(e protocol.FieldEncoder) error {
	e.SetValue(protocol.HeaderTarget, "Content-Type", protocol.StringValue("application/json"), protocol.Metadata{})

	if s.Tags != nil {
		v := s.Tags

		metadata := protocol.Metadata{}
		ls0 := e.List(protocol.BodyTarget, "Tags", metadata)
		ls0.Start()
		for _, v1 := range v {
			ls0.ListAddFields(v1)
		}
		ls0.End()

	}
	if s.MeetingId != nil {
		v := *s.MeetingId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "meetingId", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

type TagMeetingOutput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s TagMeetingOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s TagMeetingOutput) MarshalFields(e protocol.FieldEncoder) error {
	return nil
}

const opTagMeeting = "TagMeeting"

// TagMeetingRequest returns a request value for making API operation for
// Amazon Chime.
//
// Applies the specified tags to the specified Amazon Chime SDK meeting.
//
//    // Example sending a request using TagMeetingRequest.
//    req := client.TagMeetingRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/chime-2018-05-01/TagMeeting
func (c *Client) TagMeetingRequest(input *TagMeetingInput) TagMeetingRequest {
	op := &aws.Operation{
		Name:       opTagMeeting,
		HTTPMethod: "POST",
		HTTPPath:   "/meetings/{meetingId}/tags?operation=add",
	}

	if input == nil {
		input = &TagMeetingInput{}
	}

	req := c.newRequest(op, input, &TagMeetingOutput{})
	req.Handlers.Unmarshal.Remove(restjson.UnmarshalHandler)
	req.Handlers.Unmarshal.PushBackNamed(protocol.UnmarshalDiscardBodyHandler)
	return TagMeetingRequest{Request: req, Input: input, Copy: c.TagMeetingRequest}
}

// TagMeetingRequest is the request type for the
// TagMeeting API operation.
type TagMeetingRequest struct {
	*aws.Request
	Input *TagMeetingInput
	Copy  func(*TagMeetingInput) TagMeetingRequest
}

// Send marshals and sends the TagMeeting API request.
func (r TagMeetingRequest) Send(ctx context.Context) (*TagMeetingResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &TagMeetingResponse{
		TagMeetingOutput: r.Request.Data.(*TagMeetingOutput),
		response:         &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// TagMeetingResponse is the response type for the
// TagMeeting API operation.
type TagMeetingResponse struct {
	*TagMeetingOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// TagMeeting request.
func (r *TagMeetingResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}