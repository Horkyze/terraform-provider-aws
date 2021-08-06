// Code generated by "aws/internal/generators/listpages/main.go -function=DescribeIpGroups -paginator=NextToken github.com/aws/aws-sdk-go/service/workspaces"; DO NOT EDIT.

package lister

import (
	"context"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/workspaces"
)

func DescribeIpGroupsPages(conn *workspaces.WorkSpaces, input *workspaces.DescribeIpGroupsInput, fn func(*workspaces.DescribeIpGroupsOutput, bool) bool) error {
	return DescribeIpGroupsPagesWithContext(context.Background(), conn, input, fn)
}

func DescribeIpGroupsPagesWithContext(ctx context.Context, conn *workspaces.WorkSpaces, input *workspaces.DescribeIpGroupsInput, fn func(*workspaces.DescribeIpGroupsOutput, bool) bool) error {
	for {
		output, err := conn.DescribeIpGroupsWithContext(ctx, input)
		if err != nil {
			return err
		}

		lastPage := aws.StringValue(output.NextToken) == ""
		if !fn(output, lastPage) || lastPage {
			break
		}

		input.NextToken = output.NextToken
	}
	return nil
}
