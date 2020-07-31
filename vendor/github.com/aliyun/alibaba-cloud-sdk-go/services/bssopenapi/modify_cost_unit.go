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

// ModifyCostUnit invokes the bssopenapi.ModifyCostUnit API synchronously
// api document: https://help.aliyun.com/api/bssopenapi/modifycostunit.html
func (client *Client) ModifyCostUnit(request *ModifyCostUnitRequest) (response *ModifyCostUnitResponse, err error) {
	response = CreateModifyCostUnitResponse()
	err = client.DoAction(request, response)
	return
}

// ModifyCostUnitWithChan invokes the bssopenapi.ModifyCostUnit API asynchronously
// api document: https://help.aliyun.com/api/bssopenapi/modifycostunit.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) ModifyCostUnitWithChan(request *ModifyCostUnitRequest) (<-chan *ModifyCostUnitResponse, <-chan error) {
	responseChan := make(chan *ModifyCostUnitResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ModifyCostUnit(request)
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

// ModifyCostUnitWithCallback invokes the bssopenapi.ModifyCostUnit API asynchronously
// api document: https://help.aliyun.com/api/bssopenapi/modifycostunit.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) ModifyCostUnitWithCallback(request *ModifyCostUnitRequest, callback func(response *ModifyCostUnitResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ModifyCostUnitResponse
		var err error
		defer close(result)
		response, err = client.ModifyCostUnit(request)
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

// ModifyCostUnitRequest is the request struct for api ModifyCostUnit
type ModifyCostUnitRequest struct {
	*requests.RpcRequest
	UnitEntityList *[]ModifyCostUnitUnitEntityList `position:"Query" name:"UnitEntityList"  type:"Repeated"`
}

// ModifyCostUnitUnitEntityList is a repeated param struct in ModifyCostUnitRequest
type ModifyCostUnitUnitEntityList struct {
	NewUnitName string `name:"NewUnitName"`
	UnitId      string `name:"UnitId"`
	OwnerUid    string `name:"OwnerUid"`
}

// ModifyCostUnitResponse is the response struct for api ModifyCostUnit
type ModifyCostUnitResponse struct {
	*responses.BaseResponse
	RequestId string     `json:"RequestId" xml:"RequestId"`
	Success   bool       `json:"Success" xml:"Success"`
	Code      string     `json:"Code" xml:"Code"`
	Message   string     `json:"Message" xml:"Message"`
	Data      []DataItem `json:"Data" xml:"Data"`
}

// CreateModifyCostUnitRequest creates a request to invoke ModifyCostUnit API
func CreateModifyCostUnitRequest() (request *ModifyCostUnitRequest) {
	request = &ModifyCostUnitRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("BssOpenApi", "2017-12-14", "ModifyCostUnit", "", "")
	request.Method = requests.POST
	return
}

// CreateModifyCostUnitResponse creates a response to parse from ModifyCostUnit response
func CreateModifyCostUnitResponse() (response *ModifyCostUnitResponse) {
	response = &ModifyCostUnitResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
