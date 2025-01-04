/*
  Copyright 2018 Amazon.com, Inc. or its affiliates. All Rights Reserved.

  Licensed under the Apache License, Version 2.0 (the "License").
  You may not use this file except in compliance with the License.
  A copy of the License is located at

      http://www.apache.org/licenses/LICENSE-2.0

  or in the "license" file accompanying this file. This file is distributed
  on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either
  express or implied. See the License for the specific language governing
  permissions and limitations under the License.
*/

package dax

import (
	"context"
	"errors"
	"io"

	"github.com/aws/aws-dax-go/dax/internal/client"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
)

func (d *Dax) PutItem(input *dynamodb.PutItemInput) (*dynamodb.PutItemOutput, error) {
	return d.PutItemWithContext(nil, input)
}

func (d *Dax) PutItemWithContext(ctx context.Context, input *dynamodb.PutItemInput, opts ...func(*request.Request)) (*dynamodb.PutItemOutput, error) {
	o, cfn, err := d.config.requestOptions(false, ctx, opts...)
	if err != nil {
		return nil, err
	}
	if cfn != nil {
		defer cfn()
	}
	return d.client.PutItemWithOptions(input, &dynamodb.PutItemOutput{}, o)
}

func (d *Dax) PutItemRequest(input *dynamodb.PutItemInput) (*request.Request, *dynamodb.PutItemOutput) {
	op := &request.Operation{Name: client.OpPutItem}
	if input == nil {
		input = &dynamodb.PutItemInput{}
	}
	output := &dynamodb.PutItemOutput{}
	opt := client.RequestOptions{Context: context.Background()}
	req := d.client.NewDaxRequest(op, input, output, opt)
	return req, output
}

func (d *Dax) DeleteItem(input *dynamodb.DeleteItemInput) (*dynamodb.DeleteItemOutput, error) {
	return d.DeleteItemWithContext(nil, input)
}

func (d *Dax) DeleteItemWithContext(ctx context.Context, input *dynamodb.DeleteItemInput, opts ...func(*request.Request)) (*dynamodb.DeleteItemOutput, error) {
	o, cfn, err := d.config.requestOptions(false, ctx, opts...)
	if err != nil {
		return nil, err
	}
	if cfn != nil {
		defer cfn()
	}
	return d.client.DeleteItemWithOptions(input, &dynamodb.DeleteItemOutput{}, o)
}

func (d *Dax) DeleteItemRequest(input *dynamodb.DeleteItemInput) (*request.Request, *dynamodb.DeleteItemOutput) {
	op := &request.Operation{Name: client.OpDeleteItem}
	if input == nil {
		input = &dynamodb.DeleteItemInput{}
	}
	output := &dynamodb.DeleteItemOutput{}
	opt := client.RequestOptions{Context: context.Background()}
	req := d.client.NewDaxRequest(op, input, output, opt)
	return req, output
}

func (d *Dax) UpdateItem(input *dynamodb.UpdateItemInput) (*dynamodb.UpdateItemOutput, error) {
	return d.UpdateItemWithContext(nil, input)
}

func (d *Dax) UpdateItemWithContext(ctx context.Context, input *dynamodb.UpdateItemInput, opts ...func(*request.Request)) (*dynamodb.UpdateItemOutput, error) {
	o, cfn, err := d.config.requestOptions(false, ctx, opts...)
	if err != nil {
		return nil, err
	}
	if cfn != nil {
		defer cfn()
	}
	return d.client.UpdateItemWithOptions(input, &dynamodb.UpdateItemOutput{}, o)
}

func (d *Dax) UpdateItemRequest(input *dynamodb.UpdateItemInput) (*request.Request, *dynamodb.UpdateItemOutput) {
	op := &request.Operation{Name: client.OpUpdateItem}
	if input == nil {
		input = &dynamodb.UpdateItemInput{}
	}
	output := &dynamodb.UpdateItemOutput{}
	opt := client.RequestOptions{Context: context.Background()}
	req := d.client.NewDaxRequest(op, input, output, opt)
	return req, output
}

func (d *Dax) GetItem(input *dynamodb.GetItemInput) (*dynamodb.GetItemOutput, error) {
	return d.GetItemWithContext(nil, input)
}

func (d *Dax) GetItemWithContext(ctx context.Context, input *dynamodb.GetItemInput, opts ...func(*request.Request)) (*dynamodb.GetItemOutput, error) {
	o, cfn, err := d.config.requestOptions(true, ctx, opts...)
	if err != nil {
		return nil, err
	}
	if cfn != nil {
		defer cfn()
	}
	return d.client.GetItemWithOptions(input, &dynamodb.GetItemOutput{}, o)
}

func (d *Dax) GetItemRequest(input *dynamodb.GetItemInput) (*request.Request, *dynamodb.GetItemOutput) {
	op := &request.Operation{Name: client.OpGetItem}
	if input == nil {
		input = &dynamodb.GetItemInput{}
	}
	output := &dynamodb.GetItemOutput{}
	opt := client.RequestOptions{Context: context.Background()}
	req := d.client.NewDaxRequest(op, input, output, opt)
	return req, output
}

func (d *Dax) Scan(input *dynamodb.ScanInput) (*dynamodb.ScanOutput, error) {
	return d.ScanWithContext(nil, input)
}

func (d *Dax) ScanWithContext(ctx context.Context, input *dynamodb.ScanInput, opts ...func(*request.Request)) (*dynamodb.ScanOutput, error) {
	o, cfn, err := d.config.requestOptions(true, ctx, opts...)
	if err != nil {
		return nil, err
	}
	if cfn != nil {
		defer cfn()
	}
	return d.client.ScanWithOptions(input, &dynamodb.ScanOutput{}, o)
}

func (d *Dax) ScanRequest(input *dynamodb.ScanInput) (*request.Request, *dynamodb.ScanOutput) {
	op := &request.Operation{
		Name: client.OpScan,
		Paginator: &request.Paginator{
			InputTokens:     []string{"ExclusiveStartKey"},
			OutputTokens:    []string{"LastEvaluatedKey"},
			LimitToken:      "Limit",
			TruncationToken: "",
		},
	}
	if input == nil {
		input = &dynamodb.ScanInput{}
	}
	output := &dynamodb.ScanOutput{}
	opt := client.RequestOptions{Context: context.Background()}
	req := d.client.NewDaxRequest(op, input, output, opt)
	return req, output
}

func (d *Dax) Query(input *dynamodb.QueryInput) (*dynamodb.QueryOutput, error) {
	return d.QueryWithContext(nil, input)
}

func (d *Dax) QueryWithContext(ctx context.Context, input *dynamodb.QueryInput, opts ...func(*request.Request)) (*dynamodb.QueryOutput, error) {
	o, cfn, err := d.config.requestOptions(true, ctx, opts...)
	if err != nil {
		return nil, err
	}
	if cfn != nil {
		defer cfn()
	}
	return d.client.QueryWithOptions(input, &dynamodb.QueryOutput{}, o)
}

func (d *Dax) QueryRequest(input *dynamodb.QueryInput) (*request.Request, *dynamodb.QueryOutput) {
	op := &request.Operation{
		Name: client.OpQuery,
		Paginator: &request.Paginator{
			InputTokens:     []string{"ExclusiveStartKey"},
			OutputTokens:    []string{"LastEvaluatedKey"},
			LimitToken:      "Limit",
			TruncationToken: "",
		},
	}
	if input == nil {
		input = &dynamodb.QueryInput{}
	}
	output := &dynamodb.QueryOutput{}
	opt := client.RequestOptions{Context: context.Background()}
	req := d.client.NewDaxRequest(op, input, output, opt)
	return req, output
}

func (d *Dax) BatchWriteItem(input *dynamodb.BatchWriteItemInput) (*dynamodb.BatchWriteItemOutput, error) {
	return d.BatchWriteItemWithContext(nil, input)
}

func (d *Dax) BatchWriteItemWithContext(ctx context.Context, input *dynamodb.BatchWriteItemInput, opts ...func(*request.Request)) (*dynamodb.BatchWriteItemOutput, error) {
	o, cfn, err := d.config.requestOptions(false, ctx, opts...)
	if err != nil {
		return nil, err
	}
	if cfn != nil {
		defer cfn()
	}
	return d.client.BatchWriteItemWithOptions(input, &dynamodb.BatchWriteItemOutput{}, o)
}

func (d *Dax) BatchWriteItemRequest(input *dynamodb.BatchWriteItemInput) (*request.Request, *dynamodb.BatchWriteItemOutput) {
	op := &request.Operation{Name: client.OpBatchWriteItem}
	if input == nil {
		input = &dynamodb.BatchWriteItemInput{}
	}
	output := &dynamodb.BatchWriteItemOutput{}
	opt := client.RequestOptions{Context: context.Background()}
	req := d.client.NewDaxRequest(op, input, output, opt)
	return req, output
}

func (d *Dax) BatchGetItem(input *dynamodb.BatchGetItemInput) (*dynamodb.BatchGetItemOutput, error) {
	return d.BatchGetItemWithContext(nil, input)
}

func (d *Dax) BatchGetItemWithContext(ctx context.Context, input *dynamodb.BatchGetItemInput, opts ...func(*request.Request)) (*dynamodb.BatchGetItemOutput, error) {
	o, cfn, err := d.config.requestOptions(true, ctx, opts...)
	if err != nil {
		return nil, err
	}
	if cfn != nil {
		defer cfn()
	}
	return d.client.BatchGetItemWithOptions(input, &dynamodb.BatchGetItemOutput{}, o)
}

func (d *Dax) BatchGetItemRequest(input *dynamodb.BatchGetItemInput) (*request.Request, *dynamodb.BatchGetItemOutput) {
	op := &request.Operation{
		Name: client.OpBatchGetItem,
		Paginator: &request.Paginator{
			InputTokens:     []string{"RequestItems"},
			OutputTokens:    []string{"UnprocessedKeys"},
			LimitToken:      "",
			TruncationToken: "",
		},
	}
	if input == nil {
		input = &dynamodb.BatchGetItemInput{}
	}
	output := &dynamodb.BatchGetItemOutput{}
	opt := client.RequestOptions{Context: context.Background()}
	req := d.client.NewDaxRequest(op, input, output, opt)
	return req, output
}

func (d *Dax) TransactWriteItems(input *dynamodb.TransactWriteItemsInput) (*dynamodb.TransactWriteItemsOutput, error) {
	return d.TransactWriteItemsWithContext(nil, input)
}

func (d *Dax) TransactWriteItemsWithContext(ctx context.Context, input *dynamodb.TransactWriteItemsInput, opts ...func(*request.Request)) (*dynamodb.TransactWriteItemsOutput, error) {
	o, cfn, err := d.config.requestOptions(false, ctx, opts...)
	if err != nil {
		return nil, err
	}
	if cfn != nil {
		defer cfn()
	}
	return d.client.TransactWriteItemsWithOptions(input, &dynamodb.TransactWriteItemsOutput{}, o)
}

func (d *Dax) TransactWriteItemsRequest(input *dynamodb.TransactWriteItemsInput) (*request.Request, *dynamodb.TransactWriteItemsOutput) {
	op := &request.Operation{Name: client.OpTransactWriteItems}
	if input == nil {
		input = &dynamodb.TransactWriteItemsInput{}
	}
	output := &dynamodb.TransactWriteItemsOutput{}
	opt := client.RequestOptions{Context: context.Background()}
	req := d.client.NewDaxRequest(op, input, output, opt)
	return req, output
}

func (d *Dax) TransactGetItems(input *dynamodb.TransactGetItemsInput) (*dynamodb.TransactGetItemsOutput, error) {
	return d.TransactGetItemsWithContext(nil, input)
}

func (d *Dax) TransactGetItemsWithContext(ctx context.Context, input *dynamodb.TransactGetItemsInput, opts ...func(*request.Request)) (*dynamodb.TransactGetItemsOutput, error) {
	o, cfn, err := d.config.requestOptions(true, ctx, opts...)
	if err != nil {
		return nil, err
	}
	if cfn != nil {
		defer cfn()
	}
	return d.client.TransactGetItemsWithOptions(input, &dynamodb.TransactGetItemsOutput{}, o)
}

func (d *Dax) TransactGetItemsRequest(input *dynamodb.TransactGetItemsInput) (*request.Request, *dynamodb.TransactGetItemsOutput) {
	op := &request.Operation{Name: client.OpTransactGetItems}
	if input == nil {
		input = &dynamodb.TransactGetItemsInput{}
	}
	output := &dynamodb.TransactGetItemsOutput{}
	opt := client.RequestOptions{Context: context.Background()}
	req := d.client.NewDaxRequest(op, input, output, opt)
	return req, output
}

func (d *Dax) BatchGetItemPages(input *dynamodb.BatchGetItemInput, fn func(*dynamodb.BatchGetItemOutput, bool) bool) error {
	return d.BatchGetItemPagesWithContext(context.Background(), input, fn)
}

func (d *Dax) BatchGetItemPagesWithContext(ctx context.Context, input *dynamodb.BatchGetItemInput, fn func(*dynamodb.BatchGetItemOutput, bool) bool, opts ...func(*request.Request)) error {
	p := request.Pagination{
		NewRequest: func() (*request.Request, error) {
			var inCpy *dynamodb.BatchGetItemInput
			if input != nil {
				tmp := *input
				inCpy = &tmp
			}
			req, _ := d.BatchGetItemRequest(inCpy)
			req.SetContext(ctx)
			req.ApplyOptions(opts...)
			return req, nil
		},
	}

	for p.Next() {
		if !fn(p.Page().(*dynamodb.BatchGetItemOutput), !p.HasNextPage()) {
			break
		}
	}

	return p.Err()
}

func (d *Dax) QueryPages(input *dynamodb.QueryInput, fn func(*dynamodb.QueryOutput, bool) bool) error {
	return d.QueryPagesWithContext(context.Background(), input, fn)
}

func (d *Dax) QueryPagesWithContext(ctx context.Context, input *dynamodb.QueryInput, fn func(*dynamodb.QueryOutput, bool) bool, opts ...func(*request.Request)) error {
	p := request.Pagination{
		NewRequest: func() (*request.Request, error) {
			var inCpy *dynamodb.QueryInput
			if input != nil {
				tmp := *input
				inCpy = &tmp
			}
			req, _ := d.QueryRequest(inCpy)
			req.SetContext(ctx)
			req.ApplyOptions(opts...)
			return req, nil
		},
	}
	for p.Next() {
		if !fn(p.Page().(*dynamodb.QueryOutput), !p.HasNextPage()) {
			break
		}
	}
	return p.Err()
}

func (d *Dax) ScanPages(input *dynamodb.ScanInput, fn func(*dynamodb.ScanOutput, bool) bool) error {
	return d.ScanPagesWithContext(context.Background(), input, fn)
}

func (d *Dax) ScanPagesWithContext(ctx context.Context, input *dynamodb.ScanInput, fn func(*dynamodb.ScanOutput, bool) bool, opts ...func(*request.Request)) error {
	p := request.Pagination{
		NewRequest: func() (*request.Request, error) {
			var inCpy *dynamodb.ScanInput
			if input != nil {
				tmp := *input
				inCpy = &tmp
			}
			req, _ := d.ScanRequest(inCpy)
			req.SetContext(ctx)
			req.ApplyOptions(opts...)
			return req, nil
		},
	}
	for p.Next() {
		if !fn(p.Page().(*dynamodb.ScanOutput), !p.HasNextPage()) {
			break
		}
	}
	return p.Err()
}

func (d *Dax) CreateBackup(*dynamodb.CreateBackupInput) (*dynamodb.CreateBackupOutput, error) {
	return nil, d.unImpl()
}

func (d *Dax) CreateBackupWithContext(context.Context, *dynamodb.CreateBackupInput, ...func(*request.Request)) (*dynamodb.CreateBackupOutput, error) {
	return nil, d.unImpl()
}

func (d *Dax) CreateBackupRequest(*dynamodb.CreateBackupInput) (*request.Request, *dynamodb.CreateBackupOutput) {
	return newRequestForUnimplementedOperation(), &dynamodb.CreateBackupOutput{}
}

func (d *Dax) CreateGlobalTable(*dynamodb.CreateGlobalTableInput) (*dynamodb.CreateGlobalTableOutput, error) {
	return nil, d.unImpl()
}

func (d *Dax) CreateGlobalTableWithContext(context.Context, *dynamodb.CreateGlobalTableInput, ...func(*request.Request)) (*dynamodb.CreateGlobalTableOutput, error) {
	return nil, d.unImpl()
}

func (d *Dax) CreateGlobalTableRequest(*dynamodb.CreateGlobalTableInput) (*request.Request, *dynamodb.CreateGlobalTableOutput) {
	return newRequestForUnimplementedOperation(), &dynamodb.CreateGlobalTableOutput{}
}

func (d *Dax) CreateTable(*dynamodb.CreateTableInput) (*dynamodb.CreateTableOutput, error) {
	return nil, d.unImpl()
}

func (d *Dax) CreateTableWithContext(context.Context, *dynamodb.CreateTableInput, ...func(*request.Request)) (*dynamodb.CreateTableOutput, error) {
	return nil, d.unImpl()
}

func (d *Dax) CreateTableRequest(*dynamodb.CreateTableInput) (*request.Request, *dynamodb.CreateTableOutput) {
	return newRequestForUnimplementedOperation(), &dynamodb.CreateTableOutput{}
}

func (d *Dax) DeleteBackup(*dynamodb.DeleteBackupInput) (*dynamodb.DeleteBackupOutput, error) {
	return nil, d.unImpl()
}

func (d *Dax) DeleteBackupWithContext(context.Context, *dynamodb.DeleteBackupInput, ...func(*request.Request)) (*dynamodb.DeleteBackupOutput, error) {
	return nil, d.unImpl()
}

func (d *Dax) DeleteBackupRequest(*dynamodb.DeleteBackupInput) (*request.Request, *dynamodb.DeleteBackupOutput) {
	return newRequestForUnimplementedOperation(), &dynamodb.DeleteBackupOutput{}
}

func (d *Dax) DeleteTable(*dynamodb.DeleteTableInput) (*dynamodb.DeleteTableOutput, error) {
	return nil, d.unImpl()
}

func (d *Dax) DeleteTableWithContext(context.Context, *dynamodb.DeleteTableInput, ...func(*request.Request)) (*dynamodb.DeleteTableOutput, error) {
	return nil, d.unImpl()
}

func (d *Dax) DeleteTableRequest(*dynamodb.DeleteTableInput) (*request.Request, *dynamodb.DeleteTableOutput) {
	return newRequestForUnimplementedOperation(), &dynamodb.DeleteTableOutput{}
}

func (d *Dax) DescribeBackup(*dynamodb.DescribeBackupInput) (*dynamodb.DescribeBackupOutput, error) {
	return nil, d.unImpl()
}

func (d *Dax) DescribeBackupWithContext(context.Context, *dynamodb.DescribeBackupInput, ...func(*request.Request)) (*dynamodb.DescribeBackupOutput, error) {
	return nil, d.unImpl()
}

func (d *Dax) DescribeBackupRequest(*dynamodb.DescribeBackupInput) (*request.Request, *dynamodb.DescribeBackupOutput) {
	return newRequestForUnimplementedOperation(), &dynamodb.DescribeBackupOutput{}
}

func (d *Dax) DescribeContinuousBackups(*dynamodb.DescribeContinuousBackupsInput) (*dynamodb.DescribeContinuousBackupsOutput, error) {
	return nil, d.unImpl()
}

func (d *Dax) DescribeContinuousBackupsWithContext(context.Context, *dynamodb.DescribeContinuousBackupsInput, ...func(*request.Request)) (*dynamodb.DescribeContinuousBackupsOutput, error) {
	return nil, d.unImpl()
}

func (d *Dax) DescribeContinuousBackupsRequest(*dynamodb.DescribeContinuousBackupsInput) (*request.Request, *dynamodb.DescribeContinuousBackupsOutput) {
	return newRequestForUnimplementedOperation(), &dynamodb.DescribeContinuousBackupsOutput{}
}

func (d *Dax) DescribeContributorInsights(*dynamodb.DescribeContributorInsightsInput) (*dynamodb.DescribeContributorInsightsOutput, error) {
	return nil, d.unImpl()
}

func (d *Dax) DescribeContributorInsightsWithContext(context.Context, *dynamodb.DescribeContributorInsightsInput, ...func(*request.Request)) (*dynamodb.DescribeContributorInsightsOutput, error) {
	return nil, d.unImpl()
}

func (d *Dax) DescribeContributorInsightsRequest(*dynamodb.DescribeContributorInsightsInput) (*request.Request, *dynamodb.DescribeContributorInsightsOutput) {
	return newRequestForUnimplementedOperation(), &dynamodb.DescribeContributorInsightsOutput{}
}

func (d *Dax) DescribeEndpoints(*dynamodb.DescribeEndpointsInput) (*dynamodb.DescribeEndpointsOutput, error) {
	return nil, d.unImpl()
}

func (d *Dax) DescribeEndpointsWithContext(context.Context, *dynamodb.DescribeEndpointsInput, ...func(*request.Request)) (*dynamodb.DescribeEndpointsOutput, error) {
	return nil, d.unImpl()
}

func (d *Dax) DescribeEndpointsRequest(*dynamodb.DescribeEndpointsInput) (*request.Request, *dynamodb.DescribeEndpointsOutput) {
	return newRequestForUnimplementedOperation(), &dynamodb.DescribeEndpointsOutput{}
}

func (d *Dax) DescribeGlobalTable(*dynamodb.DescribeGlobalTableInput) (*dynamodb.DescribeGlobalTableOutput, error) {
	return nil, d.unImpl()
}

func (d *Dax) DescribeGlobalTableWithContext(context.Context, *dynamodb.DescribeGlobalTableInput, ...func(*request.Request)) (*dynamodb.DescribeGlobalTableOutput, error) {
	return nil, d.unImpl()
}

func (d *Dax) DescribeGlobalTableRequest(*dynamodb.DescribeGlobalTableInput) (*request.Request, *dynamodb.DescribeGlobalTableOutput) {
	return newRequestForUnimplementedOperation(), &dynamodb.DescribeGlobalTableOutput{}
}

func (d *Dax) DescribeGlobalTableSettings(*dynamodb.DescribeGlobalTableSettingsInput) (*dynamodb.DescribeGlobalTableSettingsOutput, error) {
	return nil, d.unImpl()
}

func (d *Dax) DescribeGlobalTableSettingsWithContext(context.Context, *dynamodb.DescribeGlobalTableSettingsInput, ...func(*request.Request)) (*dynamodb.DescribeGlobalTableSettingsOutput, error) {
	return nil, d.unImpl()
}

func (d *Dax) DescribeGlobalTableSettingsRequest(*dynamodb.DescribeGlobalTableSettingsInput) (*request.Request, *dynamodb.DescribeGlobalTableSettingsOutput) {
	return newRequestForUnimplementedOperation(), &dynamodb.DescribeGlobalTableSettingsOutput{}
}

func (d *Dax) DescribeImport(*dynamodb.DescribeImportInput) (*dynamodb.DescribeImportOutput, error) {
	return nil, d.unImpl()
}

func (d *Dax) DescribeImportWithContext(context.Context, *dynamodb.DescribeImportInput, ...func(*request.Request)) (*dynamodb.DescribeImportOutput, error) {
	return nil, d.unImpl()
}

func (d *Dax) DescribeImportRequest(*dynamodb.DescribeImportInput) (*request.Request, *dynamodb.DescribeImportOutput) {
	return newRequestForUnimplementedOperation(), &dynamodb.DescribeImportOutput{}
}

func (d *Dax) DescribeLimits(*dynamodb.DescribeLimitsInput) (*dynamodb.DescribeLimitsOutput, error) {
	return nil, d.unImpl()
}

func (d *Dax) DescribeLimitsWithContext(context.Context, *dynamodb.DescribeLimitsInput, ...func(*request.Request)) (*dynamodb.DescribeLimitsOutput, error) {
	return nil, d.unImpl()
}

func (d *Dax) DescribeLimitsRequest(*dynamodb.DescribeLimitsInput) (*request.Request, *dynamodb.DescribeLimitsOutput) {
	return newRequestForUnimplementedOperation(), &dynamodb.DescribeLimitsOutput{}
}

func (d *Dax) DescribeTable(*dynamodb.DescribeTableInput) (*dynamodb.DescribeTableOutput, error) {
	return nil, d.unImpl()
}

func (d *Dax) DescribeTableWithContext(context.Context, *dynamodb.DescribeTableInput, ...func(*request.Request)) (*dynamodb.DescribeTableOutput, error) {
	return nil, d.unImpl()
}

func (d *Dax) DescribeTableRequest(*dynamodb.DescribeTableInput) (*request.Request, *dynamodb.DescribeTableOutput) {
	return newRequestForUnimplementedOperation(), &dynamodb.DescribeTableOutput{}
}

func (d *Dax) DescribeTableReplicaAutoScaling(*dynamodb.DescribeTableReplicaAutoScalingInput) (*dynamodb.DescribeTableReplicaAutoScalingOutput, error) {
	return nil, d.unImpl()
}

func (d *Dax) DescribeTableReplicaAutoScalingWithContext(context.Context, *dynamodb.DescribeTableReplicaAutoScalingInput, ...func(*request.Request)) (*dynamodb.DescribeTableReplicaAutoScalingOutput, error) {
	return nil, d.unImpl()
}

func (d *Dax) DescribeTableReplicaAutoScalingRequest(*dynamodb.DescribeTableReplicaAutoScalingInput) (*request.Request, *dynamodb.DescribeTableReplicaAutoScalingOutput) {
	return newRequestForUnimplementedOperation(), &dynamodb.DescribeTableReplicaAutoScalingOutput{}
}

func (d *Dax) DescribeTimeToLive(*dynamodb.DescribeTimeToLiveInput) (*dynamodb.DescribeTimeToLiveOutput, error) {
	return nil, d.unImpl()
}

func (d *Dax) DescribeTimeToLiveWithContext(context.Context, *dynamodb.DescribeTimeToLiveInput, ...func(*request.Request)) (*dynamodb.DescribeTimeToLiveOutput, error) {
	return nil, d.unImpl()
}

func (d *Dax) DescribeTimeToLiveRequest(*dynamodb.DescribeTimeToLiveInput) (*request.Request, *dynamodb.DescribeTimeToLiveOutput) {
	return newRequestForUnimplementedOperation(), &dynamodb.DescribeTimeToLiveOutput{}
}

func (d *Dax) BatchExecuteStatement(*dynamodb.BatchExecuteStatementInput) (*dynamodb.BatchExecuteStatementOutput, error) {
	return nil, d.unImpl()
}

func (d *Dax) BatchExecuteStatementRequest(*dynamodb.BatchExecuteStatementInput) (*request.Request, *dynamodb.BatchExecuteStatementOutput) {
	return newRequestForUnimplementedOperation(), &dynamodb.BatchExecuteStatementOutput{}
}

func (d *Dax) BatchExecuteStatementWithContext(context.Context, *dynamodb.BatchExecuteStatementInput, ...func(*request.Request)) (*dynamodb.BatchExecuteStatementOutput, error) {
	return nil, d.unImpl()
}

func (d *Dax) DescribeExport(*dynamodb.DescribeExportInput) (*dynamodb.DescribeExportOutput, error) {
	return nil, d.unImpl()
}

func (d *Dax) DescribeExportWithContext(context.Context, *dynamodb.DescribeExportInput, ...func(*request.Request)) (*dynamodb.DescribeExportOutput, error) {
	return nil, d.unImpl()
}

func (d *Dax) DescribeExportRequest(*dynamodb.DescribeExportInput) (*request.Request, *dynamodb.DescribeExportOutput) {
	return newRequestForUnimplementedOperation(), &dynamodb.DescribeExportOutput{}
}

func (d *Dax) DescribeKinesisStreamingDestination(*dynamodb.DescribeKinesisStreamingDestinationInput) (*dynamodb.DescribeKinesisStreamingDestinationOutput, error) {
	return nil, d.unImpl()
}

func (d *Dax) DescribeKinesisStreamingDestinationWithContext(context.Context, *dynamodb.DescribeKinesisStreamingDestinationInput, ...func(*request.Request)) (*dynamodb.DescribeKinesisStreamingDestinationOutput, error) {
	return nil, d.unImpl()
}

func (d *Dax) DescribeKinesisStreamingDestinationRequest(*dynamodb.DescribeKinesisStreamingDestinationInput) (*request.Request, *dynamodb.DescribeKinesisStreamingDestinationOutput) {
	return newRequestForUnimplementedOperation(), &dynamodb.DescribeKinesisStreamingDestinationOutput{}
}

func (d *Dax) DisableKinesisStreamingDestination(*dynamodb.DisableKinesisStreamingDestinationInput) (*dynamodb.DisableKinesisStreamingDestinationOutput, error) {
	return nil, d.unImpl()
}

func (d *Dax) DisableKinesisStreamingDestinationWithContext(context.Context, *dynamodb.DisableKinesisStreamingDestinationInput, ...func(*request.Request)) (*dynamodb.DisableKinesisStreamingDestinationOutput, error) {
	return nil, d.unImpl()
}

func (d *Dax) DisableKinesisStreamingDestinationRequest(*dynamodb.DisableKinesisStreamingDestinationInput) (*request.Request, *dynamodb.DisableKinesisStreamingDestinationOutput) {
	return newRequestForUnimplementedOperation(), &dynamodb.DisableKinesisStreamingDestinationOutput{}
}

func (d *Dax) EnableKinesisStreamingDestination(*dynamodb.EnableKinesisStreamingDestinationInput) (*dynamodb.EnableKinesisStreamingDestinationOutput, error) {
	return nil, d.unImpl()
}

func (d *Dax) EnableKinesisStreamingDestinationWithContext(context.Context, *dynamodb.EnableKinesisStreamingDestinationInput, ...func(*request.Request)) (*dynamodb.EnableKinesisStreamingDestinationOutput, error) {
	return nil, d.unImpl()
}

func (d *Dax) EnableKinesisStreamingDestinationRequest(*dynamodb.EnableKinesisStreamingDestinationInput) (*request.Request, *dynamodb.EnableKinesisStreamingDestinationOutput) {
	return newRequestForUnimplementedOperation(), &dynamodb.EnableKinesisStreamingDestinationOutput{}
}

func (d *Dax) ExecuteStatement(*dynamodb.ExecuteStatementInput) (*dynamodb.ExecuteStatementOutput, error) {
	return nil, d.unImpl()
}

func (d *Dax) ExecuteStatementWithContext(context.Context, *dynamodb.ExecuteStatementInput, ...func(*request.Request)) (*dynamodb.ExecuteStatementOutput, error) {
	return nil, d.unImpl()
}

func (d *Dax) ExecuteStatementRequest(*dynamodb.ExecuteStatementInput) (*request.Request, *dynamodb.ExecuteStatementOutput) {
	return newRequestForUnimplementedOperation(), &dynamodb.ExecuteStatementOutput{}
}

func (d *Dax) ExecuteTransaction(*dynamodb.ExecuteTransactionInput) (*dynamodb.ExecuteTransactionOutput, error) {
	return nil, d.unImpl()
}

func (d *Dax) ExecuteTransactionWithContext(context.Context, *dynamodb.ExecuteTransactionInput, ...func(*request.Request)) (*dynamodb.ExecuteTransactionOutput, error) {
	return nil, d.unImpl()
}

func (d *Dax) ExecuteTransactionRequest(*dynamodb.ExecuteTransactionInput) (*request.Request, *dynamodb.ExecuteTransactionOutput) {
	return newRequestForUnimplementedOperation(), &dynamodb.ExecuteTransactionOutput{}
}

func (d *Dax) ExportTableToPointInTime(*dynamodb.ExportTableToPointInTimeInput) (*dynamodb.ExportTableToPointInTimeOutput, error) {
	return nil, d.unImpl()
}

func (d *Dax) ExportTableToPointInTimeWithContext(context.Context, *dynamodb.ExportTableToPointInTimeInput, ...func(*request.Request)) (*dynamodb.ExportTableToPointInTimeOutput, error) {
	return nil, d.unImpl()
}

func (d *Dax) ExportTableToPointInTimeRequest(*dynamodb.ExportTableToPointInTimeInput) (*request.Request, *dynamodb.ExportTableToPointInTimeOutput) {
	return newRequestForUnimplementedOperation(), &dynamodb.ExportTableToPointInTimeOutput{}
}

func (d *Dax) ListBackups(*dynamodb.ListBackupsInput) (*dynamodb.ListBackupsOutput, error) {
	return nil, d.unImpl()
}

func (d *Dax) ListBackupsWithContext(context.Context, *dynamodb.ListBackupsInput, ...func(*request.Request)) (*dynamodb.ListBackupsOutput, error) {
	return nil, d.unImpl()
}

func (d *Dax) ListBackupsRequest(*dynamodb.ListBackupsInput) (*request.Request, *dynamodb.ListBackupsOutput) {
	return newRequestForUnimplementedOperation(), &dynamodb.ListBackupsOutput{}
}

func (d *Dax) ImportTable(*dynamodb.ImportTableInput) (*dynamodb.ImportTableOutput, error) {
	return nil, d.unImpl()
}

func (d *Dax) ImportTableWithContext(context.Context, *dynamodb.ImportTableInput, ...func(*request.Request)) (*dynamodb.ImportTableOutput, error) {
	return nil, d.unImpl()
}

func (d *Dax) ImportTableRequest(*dynamodb.ImportTableInput) (*request.Request, *dynamodb.ImportTableOutput) {
	return newRequestForUnimplementedOperation(), &dynamodb.ImportTableOutput{}
}

func (d *Dax) ListContributorInsights(*dynamodb.ListContributorInsightsInput) (*dynamodb.ListContributorInsightsOutput, error) {
	return nil, d.unImpl()
}

func (d *Dax) ListContributorInsightsWithContext(context.Context, *dynamodb.ListContributorInsightsInput, ...func(*request.Request)) (*dynamodb.ListContributorInsightsOutput, error) {
	return nil, d.unImpl()
}

func (d *Dax) ListContributorInsightsRequest(*dynamodb.ListContributorInsightsInput) (*request.Request, *dynamodb.ListContributorInsightsOutput) {
	return newRequestForUnimplementedOperation(), &dynamodb.ListContributorInsightsOutput{}
}

func (d *Dax) ListContributorInsightsPages(*dynamodb.ListContributorInsightsInput, func(*dynamodb.ListContributorInsightsOutput, bool) bool) error {
	return d.unImpl()
}

func (d *Dax) ListContributorInsightsPagesWithContext(context.Context, *dynamodb.ListContributorInsightsInput, func(*dynamodb.ListContributorInsightsOutput, bool) bool, ...func(*request.Request)) error {
	return d.unImpl()
}

func (d *Dax) ListExports(*dynamodb.ListExportsInput) (*dynamodb.ListExportsOutput, error) {
	return nil, d.unImpl()
}

func (d *Dax) ListExportsWithContext(context.Context, *dynamodb.ListExportsInput, ...func(*request.Request)) (*dynamodb.ListExportsOutput, error) {
	return nil, d.unImpl()
}

func (d *Dax) ListExportsRequest(*dynamodb.ListExportsInput) (*request.Request, *dynamodb.ListExportsOutput) {
	return newRequestForUnimplementedOperation(), &dynamodb.ListExportsOutput{}
}

func (d *Dax) ListExportsPages(*dynamodb.ListExportsInput, func(*dynamodb.ListExportsOutput, bool) bool) error {
	return d.unImpl()
}

func (d *Dax) ListExportsPagesWithContext(context.Context, *dynamodb.ListExportsInput, func(*dynamodb.ListExportsOutput, bool) bool, ...func(*request.Request)) error {
	return d.unImpl()
}

func (d *Dax) ListGlobalTables(*dynamodb.ListGlobalTablesInput) (*dynamodb.ListGlobalTablesOutput, error) {
	return nil, d.unImpl()
}

func (d *Dax) ListGlobalTablesWithContext(context.Context, *dynamodb.ListGlobalTablesInput, ...func(*request.Request)) (*dynamodb.ListGlobalTablesOutput, error) {
	return nil, d.unImpl()
}

func (d *Dax) ListGlobalTablesRequest(*dynamodb.ListGlobalTablesInput) (*request.Request, *dynamodb.ListGlobalTablesOutput) {
	return newRequestForUnimplementedOperation(), &dynamodb.ListGlobalTablesOutput{}
}

func (d *Dax) ListImports(*dynamodb.ListImportsInput) (*dynamodb.ListImportsOutput, error) {
	return nil, d.unImpl()
}

func (d *Dax) ListImportsWithContext(context.Context, *dynamodb.ListImportsInput, ...func(*request.Request)) (*dynamodb.ListImportsOutput, error) {
	return nil, d.unImpl()
}

func (d *Dax) ListImportsRequest(*dynamodb.ListImportsInput) (*request.Request, *dynamodb.ListImportsOutput) {
	return newRequestForUnimplementedOperation(), &dynamodb.ListImportsOutput{}
}

func (d *Dax) ListImportsPages(*dynamodb.ListImportsInput, func(*dynamodb.ListImportsOutput, bool) bool) error {
	return d.unImpl()
}

func (d *Dax) ListImportsPagesWithContext(context.Context, *dynamodb.ListImportsInput, func(*dynamodb.ListImportsOutput, bool) bool, ...func(*request.Request)) error {
	return d.unImpl()
}

func (d *Dax) ListTables(*dynamodb.ListTablesInput) (*dynamodb.ListTablesOutput, error) {
	return nil, d.unImpl()
}

func (d *Dax) ListTablesWithContext(context.Context, *dynamodb.ListTablesInput, ...func(*request.Request)) (*dynamodb.ListTablesOutput, error) {
	return nil, d.unImpl()
}

func (d *Dax) ListTablesRequest(*dynamodb.ListTablesInput) (*request.Request, *dynamodb.ListTablesOutput) {
	return newRequestForUnimplementedOperation(), &dynamodb.ListTablesOutput{}
}

func (d *Dax) ListTablesPages(*dynamodb.ListTablesInput, func(*dynamodb.ListTablesOutput, bool) bool) error {
	return d.unImpl()
}

func (d *Dax) ListTablesPagesWithContext(context.Context, *dynamodb.ListTablesInput, func(*dynamodb.ListTablesOutput, bool) bool, ...func(*request.Request)) error {
	return d.unImpl()
}

func (d *Dax) ListTagsOfResource(*dynamodb.ListTagsOfResourceInput) (*dynamodb.ListTagsOfResourceOutput, error) {
	return nil, d.unImpl()
}

func (d *Dax) ListTagsOfResourceWithContext(context.Context, *dynamodb.ListTagsOfResourceInput, ...func(*request.Request)) (*dynamodb.ListTagsOfResourceOutput, error) {
	return nil, d.unImpl()
}

func (d *Dax) ListTagsOfResourceRequest(*dynamodb.ListTagsOfResourceInput) (*request.Request, *dynamodb.ListTagsOfResourceOutput) {
	return newRequestForUnimplementedOperation(), &dynamodb.ListTagsOfResourceOutput{}
}

func (d *Dax) RestoreTableFromBackup(*dynamodb.RestoreTableFromBackupInput) (*dynamodb.RestoreTableFromBackupOutput, error) {
	return nil, d.unImpl()
}

func (d *Dax) RestoreTableFromBackupWithContext(context.Context, *dynamodb.RestoreTableFromBackupInput, ...func(*request.Request)) (*dynamodb.RestoreTableFromBackupOutput, error) {
	return nil, d.unImpl()
}

func (d *Dax) RestoreTableFromBackupRequest(*dynamodb.RestoreTableFromBackupInput) (*request.Request, *dynamodb.RestoreTableFromBackupOutput) {
	return newRequestForUnimplementedOperation(), &dynamodb.RestoreTableFromBackupOutput{}
}

func (d *Dax) RestoreTableToPointInTime(*dynamodb.RestoreTableToPointInTimeInput) (*dynamodb.RestoreTableToPointInTimeOutput, error) {
	return nil, d.unImpl()
}

func (d *Dax) RestoreTableToPointInTimeWithContext(context.Context, *dynamodb.RestoreTableToPointInTimeInput, ...func(*request.Request)) (*dynamodb.RestoreTableToPointInTimeOutput, error) {
	return nil, d.unImpl()
}

func (d *Dax) RestoreTableToPointInTimeRequest(*dynamodb.RestoreTableToPointInTimeInput) (*request.Request, *dynamodb.RestoreTableToPointInTimeOutput) {
	return newRequestForUnimplementedOperation(), &dynamodb.RestoreTableToPointInTimeOutput{}
}

func (d *Dax) TagResource(*dynamodb.TagResourceInput) (*dynamodb.TagResourceOutput, error) {
	return nil, d.unImpl()
}

func (d *Dax) TagResourceWithContext(context.Context, *dynamodb.TagResourceInput, ...func(*request.Request)) (*dynamodb.TagResourceOutput, error) {
	return nil, d.unImpl()
}

func (d *Dax) TagResourceRequest(*dynamodb.TagResourceInput) (*request.Request, *dynamodb.TagResourceOutput) {
	return newRequestForUnimplementedOperation(), &dynamodb.TagResourceOutput{}
}

func (d *Dax) UntagResource(*dynamodb.UntagResourceInput) (*dynamodb.UntagResourceOutput, error) {
	return nil, d.unImpl()
}

func (d *Dax) UntagResourceWithContext(context.Context, *dynamodb.UntagResourceInput, ...func(*request.Request)) (*dynamodb.UntagResourceOutput, error) {
	return nil, d.unImpl()
}

func (d *Dax) UntagResourceRequest(*dynamodb.UntagResourceInput) (*request.Request, *dynamodb.UntagResourceOutput) {
	return newRequestForUnimplementedOperation(), &dynamodb.UntagResourceOutput{}
}

func (d *Dax) UpdateContinuousBackups(*dynamodb.UpdateContinuousBackupsInput) (*dynamodb.UpdateContinuousBackupsOutput, error) {
	return nil, d.unImpl()
}

func (d *Dax) UpdateContinuousBackupsWithContext(context.Context, *dynamodb.UpdateContinuousBackupsInput, ...func(*request.Request)) (*dynamodb.UpdateContinuousBackupsOutput, error) {
	return nil, d.unImpl()
}

func (d *Dax) UpdateContinuousBackupsRequest(*dynamodb.UpdateContinuousBackupsInput) (*request.Request, *dynamodb.UpdateContinuousBackupsOutput) {
	return newRequestForUnimplementedOperation(), &dynamodb.UpdateContinuousBackupsOutput{}
}

func (d *Dax) UpdateContributorInsights(*dynamodb.UpdateContributorInsightsInput) (*dynamodb.UpdateContributorInsightsOutput, error) {
	return nil, d.unImpl()
}

func (d *Dax) UpdateContributorInsightsWithContext(context.Context, *dynamodb.UpdateContributorInsightsInput, ...func(*request.Request)) (*dynamodb.UpdateContributorInsightsOutput, error) {
	return nil, d.unImpl()
}

func (d *Dax) UpdateContributorInsightsRequest(*dynamodb.UpdateContributorInsightsInput) (*request.Request, *dynamodb.UpdateContributorInsightsOutput) {
	return newRequestForUnimplementedOperation(), &dynamodb.UpdateContributorInsightsOutput{}
}

func (d *Dax) UpdateGlobalTable(*dynamodb.UpdateGlobalTableInput) (*dynamodb.UpdateGlobalTableOutput, error) {
	return nil, d.unImpl()
}

func (d *Dax) UpdateGlobalTableWithContext(context.Context, *dynamodb.UpdateGlobalTableInput, ...func(*request.Request)) (*dynamodb.UpdateGlobalTableOutput, error) {
	return nil, d.unImpl()
}

func (d *Dax) UpdateGlobalTableRequest(*dynamodb.UpdateGlobalTableInput) (*request.Request, *dynamodb.UpdateGlobalTableOutput) {
	return newRequestForUnimplementedOperation(), &dynamodb.UpdateGlobalTableOutput{}
}

func (d *Dax) UpdateGlobalTableSettings(*dynamodb.UpdateGlobalTableSettingsInput) (*dynamodb.UpdateGlobalTableSettingsOutput, error) {
	return nil, d.unImpl()
}

func (d *Dax) UpdateGlobalTableSettingsWithContext(context.Context, *dynamodb.UpdateGlobalTableSettingsInput, ...func(*request.Request)) (*dynamodb.UpdateGlobalTableSettingsOutput, error) {
	return nil, d.unImpl()
}

func (d *Dax) UpdateGlobalTableSettingsRequest(*dynamodb.UpdateGlobalTableSettingsInput) (*request.Request, *dynamodb.UpdateGlobalTableSettingsOutput) {
	return newRequestForUnimplementedOperation(), &dynamodb.UpdateGlobalTableSettingsOutput{}
}

func (d *Dax) UpdateKinesisStreamingDestination(*dynamodb.UpdateKinesisStreamingDestinationInput) (*dynamodb.UpdateKinesisStreamingDestinationOutput, error) {
	return nil, d.unImpl()
}

func (d *Dax) UpdateKinesisStreamingDestinationWithContext(context.Context, *dynamodb.UpdateKinesisStreamingDestinationInput, ...func(*request.Request)) (*dynamodb.UpdateKinesisStreamingDestinationOutput, error) {
	return nil, d.unImpl()
}

func (d *Dax) UpdateKinesisStreamingDestinationRequest(*dynamodb.UpdateKinesisStreamingDestinationInput) (*request.Request, *dynamodb.UpdateKinesisStreamingDestinationOutput) {
	return newRequestForUnimplementedOperation(), &dynamodb.UpdateKinesisStreamingDestinationOutput{}
}

func (d *Dax) UpdateTable(*dynamodb.UpdateTableInput) (*dynamodb.UpdateTableOutput, error) {
	return nil, d.unImpl()
}

func (d *Dax) UpdateTableWithContext(context.Context, *dynamodb.UpdateTableInput, ...func(*request.Request)) (*dynamodb.UpdateTableOutput, error) {
	return nil, d.unImpl()
}

func (d *Dax) UpdateTableRequest(*dynamodb.UpdateTableInput) (*request.Request, *dynamodb.UpdateTableOutput) {
	return newRequestForUnimplementedOperation(), &dynamodb.UpdateTableOutput{}
}

func (d *Dax) UpdateTableReplicaAutoScaling(*dynamodb.UpdateTableReplicaAutoScalingInput) (*dynamodb.UpdateTableReplicaAutoScalingOutput, error) {
	return nil, d.unImpl()
}

func (d *Dax) UpdateTableReplicaAutoScalingWithContext(context.Context, *dynamodb.UpdateTableReplicaAutoScalingInput, ...func(*request.Request)) (*dynamodb.UpdateTableReplicaAutoScalingOutput, error) {
	return nil, d.unImpl()
}

func (d *Dax) UpdateTableReplicaAutoScalingRequest(*dynamodb.UpdateTableReplicaAutoScalingInput) (*request.Request, *dynamodb.UpdateTableReplicaAutoScalingOutput) {
	return newRequestForUnimplementedOperation(), &dynamodb.UpdateTableReplicaAutoScalingOutput{}
}

func (d *Dax) UpdateTimeToLive(*dynamodb.UpdateTimeToLiveInput) (*dynamodb.UpdateTimeToLiveOutput, error) {
	return nil, d.unImpl()
}

func (d *Dax) UpdateTimeToLiveWithContext(context.Context, *dynamodb.UpdateTimeToLiveInput, ...func(*request.Request)) (*dynamodb.UpdateTimeToLiveOutput, error) {
	return nil, d.unImpl()
}

func (d *Dax) UpdateTimeToLiveRequest(*dynamodb.UpdateTimeToLiveInput) (*request.Request, *dynamodb.UpdateTimeToLiveOutput) {
	return newRequestForUnimplementedOperation(), &dynamodb.UpdateTimeToLiveOutput{}
}

func (d *Dax) WaitUntilTableExists(*dynamodb.DescribeTableInput) error {
	return d.unImpl()
}

func (d *Dax) WaitUntilTableExistsWithContext(context.Context, *dynamodb.DescribeTableInput, ...request.WaiterOption) error {
	return d.unImpl()
}

func (d *Dax) WaitUntilTableNotExists(*dynamodb.DescribeTableInput) error {
	return d.unImpl()
}

func (d *Dax) WaitUntilTableNotExistsWithContext(context.Context, *dynamodb.DescribeTableInput, ...request.WaiterOption) error {
	return d.unImpl()
}

func (d *Dax) DeleteResourcePolicy(*dynamodb.DeleteResourcePolicyInput) (*dynamodb.DeleteResourcePolicyOutput, error) {
	return nil, d.unImpl()
}

func (d *Dax) DeleteResourcePolicyWithContext(context.Context, *dynamodb.DeleteResourcePolicyInput, ...func(*request.Request)) (*dynamodb.DeleteResourcePolicyOutput, error) {
	return nil, d.unImpl()
}

func (d *Dax) DeleteResourcePolicyRequest(input *dynamodb.DeleteResourcePolicyInput) (*request.Request, *dynamodb.DeleteResourcePolicyOutput) {
	return newRequestForUnimplementedOperation(), &dynamodb.DeleteResourcePolicyOutput{}
}

func (d *Dax) GetResourcePolicy(*dynamodb.GetResourcePolicyInput) (*dynamodb.GetResourcePolicyOutput, error) {
	return nil, d.unImpl()
}
func (d *Dax) GetResourcePolicyWithContext(context.Context, *dynamodb.GetResourcePolicyInput, ...func(*request.Request)) (*dynamodb.GetResourcePolicyOutput, error) {
	return nil, d.unImpl()
}

func (d *Dax) GetResourcePolicyRequest(input *dynamodb.GetResourcePolicyInput) (*request.Request, *dynamodb.GetResourcePolicyOutput) {
	return newRequestForUnimplementedOperation(), &dynamodb.GetResourcePolicyOutput{}
}

func (d *Dax) PutResourcePolicy(*dynamodb.PutResourcePolicyInput) (*dynamodb.PutResourcePolicyOutput, error) {
	return nil, d.unImpl()
}
func (d *Dax) PutResourcePolicyWithContext(context.Context, *dynamodb.PutResourcePolicyInput, ...func(*request.Request)) (*dynamodb.PutResourcePolicyOutput, error) {
	return nil, d.unImpl()
}

func (d *Dax) PutResourcePolicyRequest(input *dynamodb.PutResourcePolicyInput) (*request.Request, *dynamodb.PutResourcePolicyOutput) {
	return newRequestForUnimplementedOperation(), &dynamodb.PutResourcePolicyOutput{}
}

func (d *Dax) unImpl() error {
	return errors.New(client.ErrCodeNotImplemented)
}

func (d *Dax) Close() error {
	if c, ok := d.client.(io.Closer); ok {
		return c.Close()
	}
	return nil
}
