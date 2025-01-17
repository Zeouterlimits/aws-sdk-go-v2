// Code generated by smithy-go-codegen DO NOT EDIT.

package configservice

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/configservice/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Provides organization conformance pack deployment status for an organization.
// The status is not considered successful until organization conformance pack is
// successfully deployed in all the member accounts with an exception of excluded
// accounts. When you specify the limit and the next token, you receive a paginated
// response. Limit and next token are not applicable if you specify organization
// conformance pack names. They are only applicable, when you request all the
// organization conformance packs.
func (c *Client) DescribeOrganizationConformancePackStatuses(ctx context.Context, params *DescribeOrganizationConformancePackStatusesInput, optFns ...func(*Options)) (*DescribeOrganizationConformancePackStatusesOutput, error) {
	if params == nil {
		params = &DescribeOrganizationConformancePackStatusesInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeOrganizationConformancePackStatuses", params, optFns, addOperationDescribeOrganizationConformancePackStatusesMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeOrganizationConformancePackStatusesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeOrganizationConformancePackStatusesInput struct {

	// The maximum number of OrganizationConformancePackStatuses returned on each page.
	// If you do no specify a number, AWS Config uses the default. The default is 100.
	Limit int32

	// The nextToken string returned on a previous page that you use to get the next
	// page of results in a paginated response.
	NextToken *string

	// The names of organization conformance packs for which you want status details.
	// If you do not specify any names, AWS Config returns details for all your
	// organization conformance packs.
	OrganizationConformancePackNames []string
}

type DescribeOrganizationConformancePackStatusesOutput struct {

	// The nextToken string returned on a previous page that you use to get the next
	// page of results in a paginated response.
	NextToken *string

	// A list of OrganizationConformancePackStatus objects.
	OrganizationConformancePackStatuses []types.OrganizationConformancePackStatus

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationDescribeOrganizationConformancePackStatusesMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpDescribeOrganizationConformancePackStatuses{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpDescribeOrganizationConformancePackStatuses{}, middleware.After)
	if err != nil {
		return err
	}
	if err = addSetLoggerMiddleware(stack, options); err != nil {
		return err
	}
	if err = awsmiddleware.AddClientRequestIDMiddleware(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddComputeContentLengthMiddleware(stack); err != nil {
		return err
	}
	if err = addResolveEndpointMiddleware(stack, options); err != nil {
		return err
	}
	if err = v4.AddComputePayloadSHA256Middleware(stack); err != nil {
		return err
	}
	if err = addRetryMiddlewares(stack, options); err != nil {
		return err
	}
	if err = addHTTPSignerV4Middleware(stack, options); err != nil {
		return err
	}
	if err = awsmiddleware.AddRawResponseToMetadata(stack); err != nil {
		return err
	}
	if err = awsmiddleware.AddRecordResponseTiming(stack); err != nil {
		return err
	}
	if err = addClientUserAgent(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddErrorCloseResponseBodyMiddleware(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddCloseResponseBodyMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeOrganizationConformancePackStatuses(options.Region), middleware.Before); err != nil {
		return err
	}
	if err = addRequestIDRetrieverMiddleware(stack); err != nil {
		return err
	}
	if err = addResponseErrorMiddleware(stack); err != nil {
		return err
	}
	if err = addRequestResponseLogging(stack, options); err != nil {
		return err
	}
	return nil
}

func newServiceMetadataMiddleware_opDescribeOrganizationConformancePackStatuses(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "config",
		OperationName: "DescribeOrganizationConformancePackStatuses",
	}
}
