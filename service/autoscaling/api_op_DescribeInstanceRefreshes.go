// Code generated by smithy-go-codegen DO NOT EDIT.

package autoscaling

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/autoscaling/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Describes one or more instance refreshes. You can determine the status of a
// request by looking at the Status parameter. The following are the possible
// statuses:
//
//     * Pending - The request was created, but the operation has not
// started.
//
//     * InProgress - The operation is in progress.
//
//     * Successful -
// The operation completed successfully.
//
//     * Failed - The operation failed to
// complete. You can troubleshoot using the status reason and the scaling
// activities.
//
//     * Cancelling - An ongoing operation is being cancelled.
// Cancellation does not roll back any replacements that have already been
// completed, but it prevents new replacements from being started.
//
//     * Cancelled
// - The operation is cancelled.
//
// For more information, see Replacing Auto Scaling
// Instances Based on an Instance Refresh
// (https://docs.aws.amazon.com/autoscaling/ec2/userguide/asg-instance-refresh.html).
func (c *Client) DescribeInstanceRefreshes(ctx context.Context, params *DescribeInstanceRefreshesInput, optFns ...func(*Options)) (*DescribeInstanceRefreshesOutput, error) {
	stack := middleware.NewStack("DescribeInstanceRefreshes", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsquery_serdeOpDescribeInstanceRefreshesMiddlewares(stack)
	awsmiddleware.AddRequestInvocationIDMiddleware(stack)
	smithyhttp.AddContentLengthMiddleware(stack)
	AddResolveEndpointMiddleware(stack, options)
	v4.AddComputePayloadSHA256Middleware(stack)
	retry.AddRetryMiddlewares(stack, options)
	addHTTPSignerV4Middleware(stack, options)
	awsmiddleware.AddAttemptClockSkewMiddleware(stack)
	addClientUserAgent(stack)
	smithyhttp.AddErrorCloseResponseBodyMiddleware(stack)
	smithyhttp.AddCloseResponseBodyMiddleware(stack)
	addOpDescribeInstanceRefreshesValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeInstanceRefreshes(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)

	for _, fn := range options.APIOptions {
		if err := fn(stack); err != nil {
			return nil, err
		}
	}
	handler := middleware.DecorateHandler(smithyhttp.NewClientHandler(options.HTTPClient), stack)
	result, metadata, err := handler.Handle(ctx, params)
	if err != nil {
		return nil, &smithy.OperationError{
			ServiceID:     ServiceID,
			OperationName: "DescribeInstanceRefreshes",
			Err:           err,
		}
	}
	out := result.(*DescribeInstanceRefreshesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeInstanceRefreshesInput struct {
	// The name of the Auto Scaling group.
	AutoScalingGroupName *string
	// The maximum number of items to return with this call. The default value is 50
	// and the maximum value is 100.
	MaxRecords *int32
	// The token for the next set of items to return. (You received this token from a
	// previous call.)
	NextToken *string
	// One or more instance refresh IDs.
	InstanceRefreshIds []*string
}

type DescribeInstanceRefreshesOutput struct {
	// A string that indicates that the response contains more items than can be
	// returned in a single response. To receive additional items, specify this string
	// for the NextToken value when requesting the next set of items. This value is
	// null when there are no more items to return.
	NextToken *string
	// The instance refreshes for the specified group.
	InstanceRefreshes []*types.InstanceRefresh

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsquery_serdeOpDescribeInstanceRefreshesMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsquery_serializeOpDescribeInstanceRefreshes{}, middleware.After)
	stack.Deserialize.Add(&awsAwsquery_deserializeOpDescribeInstanceRefreshes{}, middleware.After)
}

func newServiceMetadataMiddleware_opDescribeInstanceRefreshes(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "autoscaling",
		OperationName: "DescribeInstanceRefreshes",
	}
}