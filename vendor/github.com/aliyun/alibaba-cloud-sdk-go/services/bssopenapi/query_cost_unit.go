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

// QueryCostUnit invokes the bssopenapi.QueryCostUnit API synchronously
// api document: https://help.aliyun.com/api/bssopenapi/querycostunit.html
func (client *Client) QueryCostUnit(request *QueryCostUnitRequest) (response *QueryCostUnitResponse, err error) {
	response = CreateQueryCostUnitResponse()
	err = client.DoAction(request, response)
	return
}

// QueryCostUnitWithChan invokes the bssopenapi.QueryCostUnit API asynchronously
// api document: https://help.aliyun.com/api/bssopenapi/querycostunit.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) QueryCostUnitWithChan(request *QueryCostUnitRequest) (<-chan *QueryCostUnitResponse, <-chan error) {
	responseChan := make(chan *QueryCostUnitResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.QueryCostUnit(request)
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

// QueryCostUnitWithCallback invokes the bssopenapi.QueryCostUnit API asynchronously
// api document: https://help.aliyun.com/api/bssopenapi/querycostunit.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) QueryCostUnitWithCallback(request *QueryCostUnitRequest, callback func(response *QueryCostUnitResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *QueryCostUnitResponse
		var err error
		defer close(result)
		response, err = client.QueryCostUnit(request)
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

// QueryCostUnitRequest is the request struct for api QueryCostUnit
type QueryCostUnitRequest struct {
	*requests.RpcRequest
	ParentUnitId requests.Integer `position:"Query" name:"ParentUnitId"`
	PageNum      requests.Integer `position:"Query" name:"PageNum"`
	PageSize     requests.Integer `position:"Query" name:"PageSize"`
	OwnerUid     requests.Integer `position:"Query" name:"OwnerUid"`
}

// QueryCostUnitResponse is the response struct for api QueryCostUnit
type QueryCostUnitResponse struct {
	*responses.BaseResponse
	RequestId string              `json:"RequestId" xml:"RequestId"`
	Success   bool                `json:"Success" xml:"Success"`
	Code      string              `json:"Code" xml:"Code"`
	Message   string              `json:"Message" xml:"Message"`
	Data      DataInQueryCostUnit `json:"Data" xml:"Data"`
}

// CreateQueryCostUnitRequest creates a request to invoke QueryCostUnit API
func CreateQueryCostUnitRequest() (request *QueryCostUnitRequest) {
	request = &QueryCostUnitRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("BssOpenApi", "2017-12-14", "QueryCostUnit", "", "")
	request.Method = requests.POST
	return
}

// CreateQueryCostUnitResponse creates a response to parse from QueryCostUnit response
func CreateQueryCostUnitResponse() (response *QueryCostUnitResponse) {
	response = &QueryCostUnitResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
