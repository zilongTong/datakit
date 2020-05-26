package bssopenapi

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

// QueryBillOverview invokes the bssopenapi.QueryBillOverview API synchronously
// api document: https://help.aliyun.com/api/bssopenapi/querybilloverview.html
func (client *Client) QueryBillOverview(request *QueryBillOverviewRequest) (response *QueryBillOverviewResponse, err error) {
	response = CreateQueryBillOverviewResponse()
	err = client.DoAction(request, response)
	return
}

// QueryBillOverviewWithChan invokes the bssopenapi.QueryBillOverview API asynchronously
// api document: https://help.aliyun.com/api/bssopenapi/querybilloverview.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) QueryBillOverviewWithChan(request *QueryBillOverviewRequest) (<-chan *QueryBillOverviewResponse, <-chan error) {
	responseChan := make(chan *QueryBillOverviewResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.QueryBillOverview(request)
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

// QueryBillOverviewWithCallback invokes the bssopenapi.QueryBillOverview API asynchronously
// api document: https://help.aliyun.com/api/bssopenapi/querybilloverview.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) QueryBillOverviewWithCallback(request *QueryBillOverviewRequest, callback func(response *QueryBillOverviewResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *QueryBillOverviewResponse
		var err error
		defer close(result)
		response, err = client.QueryBillOverview(request)
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

// QueryBillOverviewRequest is the request struct for api QueryBillOverview
type QueryBillOverviewRequest struct {
	*requests.RpcRequest
	ProductCode      string `position:"Query" name:"ProductCode"`
	SubscriptionType string `position:"Query" name:"SubscriptionType"`
	BillingCycle     string `position:"Query" name:"BillingCycle"`
	ProductType      string `position:"Query" name:"ProductType"`
}

// QueryBillOverviewResponse is the response struct for api QueryBillOverview
type QueryBillOverviewResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	Success   bool   `json:"Success" xml:"Success"`
	Code      string `json:"Code" xml:"Code"`
	Message   string `json:"Message" xml:"Message"`
	Data      Data   `json:"Data" xml:"Data"`
}

// CreateQueryBillOverviewRequest creates a request to invoke QueryBillOverview API
func CreateQueryBillOverviewRequest() (request *QueryBillOverviewRequest) {
	request = &QueryBillOverviewRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("BssOpenApi", "2017-12-14", "QueryBillOverview", "", "")
	return
}

// CreateQueryBillOverviewResponse creates a response to parse from QueryBillOverview response
func CreateQueryBillOverviewResponse() (response *QueryBillOverviewResponse) {
	response = &QueryBillOverviewResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
