package aegis

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

// DescribeYesterdayStatistics invokes the aegis.DescribeYesterdayStatistics API synchronously
// api document: https://help.aliyun.com/api/aegis/describeyesterdaystatistics.html
func (client *Client) DescribeYesterdayStatistics(request *DescribeYesterdayStatisticsRequest) (response *DescribeYesterdayStatisticsResponse, err error) {
	response = CreateDescribeYesterdayStatisticsResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeYesterdayStatisticsWithChan invokes the aegis.DescribeYesterdayStatistics API asynchronously
// api document: https://help.aliyun.com/api/aegis/describeyesterdaystatistics.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeYesterdayStatisticsWithChan(request *DescribeYesterdayStatisticsRequest) (<-chan *DescribeYesterdayStatisticsResponse, <-chan error) {
	responseChan := make(chan *DescribeYesterdayStatisticsResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeYesterdayStatistics(request)
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

// DescribeYesterdayStatisticsWithCallback invokes the aegis.DescribeYesterdayStatistics API asynchronously
// api document: https://help.aliyun.com/api/aegis/describeyesterdaystatistics.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeYesterdayStatisticsWithCallback(request *DescribeYesterdayStatisticsRequest, callback func(response *DescribeYesterdayStatisticsResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeYesterdayStatisticsResponse
		var err error
		defer close(result)
		response, err = client.DescribeYesterdayStatistics(request)
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

// DescribeYesterdayStatisticsRequest is the request struct for api DescribeYesterdayStatistics
type DescribeYesterdayStatisticsRequest struct {
	*requests.RpcRequest
	SourceIp string `position:"Query" name:"SourceIp"`
}

// DescribeYesterdayStatisticsResponse is the response struct for api DescribeYesterdayStatistics
type DescribeYesterdayStatisticsResponse struct {
	*responses.BaseResponse
	RequestId           string `json:"RequestId" xml:"RequestId"`
	TotalCount          int    `json:"TotalCount" xml:"TotalCount"`
	NewRiskCheckCount   int    `json:"NewRiskCheckCount" xml:"NewRiskCheckCount"`
	NewVulCount         int    `json:"NewVulCount" xml:"NewVulCount"`
	NewHealthCheckCount int    `json:"NewHealthCheckCount" xml:"NewHealthCheckCount"`
	NewSuspiciousCount  int    `json:"NewSuspiciousCount" xml:"NewSuspiciousCount"`
}

// CreateDescribeYesterdayStatisticsRequest creates a request to invoke DescribeYesterdayStatistics API
func CreateDescribeYesterdayStatisticsRequest() (request *DescribeYesterdayStatisticsRequest) {
	request = &DescribeYesterdayStatisticsRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("aegis", "2016-11-11", "DescribeYesterdayStatistics", "vipaegis", "openAPI")
	return
}

// CreateDescribeYesterdayStatisticsResponse creates a response to parse from DescribeYesterdayStatistics response
func CreateDescribeYesterdayStatisticsResponse() (response *DescribeYesterdayStatisticsResponse) {
	response = &DescribeYesterdayStatisticsResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
