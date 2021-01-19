package rds

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

// EvaluateDedicatedHostInstanceResource invokes the rds.EvaluateDedicatedHostInstanceResource API synchronously
// api document: https://help.aliyun.com/api/rds/evaluatededicatedhostinstanceresource.html
func (client *Client) EvaluateDedicatedHostInstanceResource(request *EvaluateDedicatedHostInstanceResourceRequest) (response *EvaluateDedicatedHostInstanceResourceResponse, err error) {
	response = CreateEvaluateDedicatedHostInstanceResourceResponse()
	err = client.DoAction(request, response)
	return
}

// EvaluateDedicatedHostInstanceResourceWithChan invokes the rds.EvaluateDedicatedHostInstanceResource API asynchronously
// api document: https://help.aliyun.com/api/rds/evaluatededicatedhostinstanceresource.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) EvaluateDedicatedHostInstanceResourceWithChan(request *EvaluateDedicatedHostInstanceResourceRequest) (<-chan *EvaluateDedicatedHostInstanceResourceResponse, <-chan error) {
	responseChan := make(chan *EvaluateDedicatedHostInstanceResourceResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.EvaluateDedicatedHostInstanceResource(request)
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

// EvaluateDedicatedHostInstanceResourceWithCallback invokes the rds.EvaluateDedicatedHostInstanceResource API asynchronously
// api document: https://help.aliyun.com/api/rds/evaluatededicatedhostinstanceresource.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) EvaluateDedicatedHostInstanceResourceWithCallback(request *EvaluateDedicatedHostInstanceResourceRequest, callback func(response *EvaluateDedicatedHostInstanceResourceResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *EvaluateDedicatedHostInstanceResourceResponse
		var err error
		defer close(result)
		response, err = client.EvaluateDedicatedHostInstanceResource(request)
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

// EvaluateDedicatedHostInstanceResourceRequest is the request struct for api EvaluateDedicatedHostInstanceResource
type EvaluateDedicatedHostInstanceResourceRequest struct {
	*requests.RpcRequest
	ResourceOwnerId      requests.Integer `position:"Query" name:"ResourceOwnerId"`
	EngineVersion        string           `position:"Query" name:"EngineVersion"`
	Engine               string           `position:"Query" name:"Engine"`
	DedicatedHostGroupId string           `position:"Query" name:"DedicatedHostGroupId"`
	InstanceClassNames   string           `position:"Query" name:"InstanceClassNames"`
	DiskSize             string           `position:"Query" name:"DiskSize"`
	ResourceOwnerAccount string           `position:"Query" name:"ResourceOwnerAccount"`
	OwnerId              requests.Integer `position:"Query" name:"OwnerId"`
	DiskType             string           `position:"Query" name:"DiskType"`
}

// EvaluateDedicatedHostInstanceResourceResponse is the response struct for api EvaluateDedicatedHostInstanceResource
type EvaluateDedicatedHostInstanceResourceResponse struct {
	*responses.BaseResponse
	RequestId       string `json:"RequestId" xml:"RequestId"`
	DBInstanceClass string `json:"DBInstanceClass" xml:"DBInstanceClass"`
	Available       int    `json:"Available" xml:"Available"`
}

// CreateEvaluateDedicatedHostInstanceResourceRequest creates a request to invoke EvaluateDedicatedHostInstanceResource API
func CreateEvaluateDedicatedHostInstanceResourceRequest() (request *EvaluateDedicatedHostInstanceResourceRequest) {
	request = &EvaluateDedicatedHostInstanceResourceRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Rds", "2014-08-15", "EvaluateDedicatedHostInstanceResource", "rds", "openAPI")
	request.Method = requests.POST
	return
}

// CreateEvaluateDedicatedHostInstanceResourceResponse creates a response to parse from EvaluateDedicatedHostInstanceResource response
func CreateEvaluateDedicatedHostInstanceResourceResponse() (response *EvaluateDedicatedHostInstanceResourceResponse) {
	response = &EvaluateDedicatedHostInstanceResourceResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
