package r_kvstore

//Licensed under the Apache License, Version 2.0 (the "License");
//you may not use this file except in compliance with the License.
//You may obtain a copy of the License at
//
//http://www.apache.org/licenses/LICENSE-2.0
//
//Unless required by applicable law or agreed to in writing, software
//distributed under the License is distributed on an "AS IS" BASIS,
//WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//See the License for the specific language governing permissions and
//limitations under the License.
//
// Code generated by Alibaba Cloud SDK Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

// DescribeAuditRecords invokes the r_kvstore.DescribeAuditRecords API synchronously
// api document: https://help.aliyun.com/api/r-kvstore/describeauditrecords.html
func (client *Client) DescribeAuditRecords(request *DescribeAuditRecordsRequest) (response *DescribeAuditRecordsResponse, err error) {
	response = CreateDescribeAuditRecordsResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeAuditRecordsWithChan invokes the r_kvstore.DescribeAuditRecords API asynchronously
// api document: https://help.aliyun.com/api/r-kvstore/describeauditrecords.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeAuditRecordsWithChan(request *DescribeAuditRecordsRequest) (<-chan *DescribeAuditRecordsResponse, <-chan error) {
	responseChan := make(chan *DescribeAuditRecordsResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeAuditRecords(request)
		if err != nil {
			errChan <- err
		} else {
			responseChan <- response
		}
	})
	if err != nil {
		errChan <- err
		close(responseChan)
		close(errChan)
	}
	return responseChan, errChan
}

// DescribeAuditRecordsWithCallback invokes the r_kvstore.DescribeAuditRecords API asynchronously
// api document: https://help.aliyun.com/api/r-kvstore/describeauditrecords.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeAuditRecordsWithCallback(request *DescribeAuditRecordsRequest, callback func(response *DescribeAuditRecordsResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeAuditRecordsResponse
		var err error
		defer close(result)
		response, err = client.DescribeAuditRecords(request)
		callback(response, err)
		result <- 1
	})
	if err != nil {
		defer close(result)
		callback(nil, err)
		result <- 0
	}
	return result
}

// DescribeAuditRecordsRequest is the request struct for api DescribeAuditRecords
type DescribeAuditRecordsRequest struct {
	*requests.RpcRequest
	ResourceOwnerId      requests.Integer `position:"Query" name:"ResourceOwnerId"`
	StartTime            string           `position:"Query" name:"StartTime"`
	QueryKeywords        string           `position:"Query" name:"QueryKeywords"`
	PageNumber           requests.Integer `position:"Query" name:"PageNumber"`
	HostAddress          string           `position:"Query" name:"HostAddress"`
	AccountName          string           `position:"Query" name:"AccountName"`
	SecurityToken        string           `position:"Query" name:"SecurityToken"`
	PageSize             requests.Integer `position:"Query" name:"PageSize"`
	NodeId               string           `position:"Query" name:"NodeId"`
	ResourceOwnerAccount string           `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string           `position:"Query" name:"OwnerAccount"`
	EndTime              string           `position:"Query" name:"EndTime"`
	OwnerId              requests.Integer `position:"Query" name:"OwnerId"`
	InstanceId           string           `position:"Query" name:"InstanceId"`
	DatabaseName         string           `position:"Query" name:"DatabaseName"`
}

// DescribeAuditRecordsResponse is the response struct for api DescribeAuditRecords
type DescribeAuditRecordsResponse struct {
	*responses.BaseResponse
	RequestId        string                      `json:"RequestId" xml:"RequestId"`
	PageNumber       int                         `json:"PageNumber" xml:"PageNumber"`
	PageSize         int                         `json:"PageSize" xml:"PageSize"`
	TotalRecordCount int                         `json:"TotalRecordCount" xml:"TotalRecordCount"`
	InstanceName     string                      `json:"InstanceName" xml:"InstanceName"`
	StartTime        string                      `json:"StartTime" xml:"StartTime"`
	EndTime          string                      `json:"EndTime" xml:"EndTime"`
	Items            ItemsInDescribeAuditRecords `json:"Items" xml:"Items"`
}

// CreateDescribeAuditRecordsRequest creates a request to invoke DescribeAuditRecords API
func CreateDescribeAuditRecordsRequest() (request *DescribeAuditRecordsRequest) {
	request = &DescribeAuditRecordsRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("R-kvstore", "2015-01-01", "DescribeAuditRecords", "redisa", "openAPI")
	request.Method = requests.POST
	return
}

// CreateDescribeAuditRecordsResponse creates a response to parse from DescribeAuditRecords response
func CreateDescribeAuditRecordsResponse() (response *DescribeAuditRecordsResponse) {
	response = &DescribeAuditRecordsResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
