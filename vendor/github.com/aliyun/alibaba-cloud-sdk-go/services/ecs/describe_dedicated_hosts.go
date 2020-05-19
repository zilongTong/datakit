package ecs

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

// DescribeDedicatedHosts invokes the ecs.DescribeDedicatedHosts API synchronously
// api document: https://help.aliyun.com/api/ecs/describededicatedhosts.html
func (client *Client) DescribeDedicatedHosts(request *DescribeDedicatedHostsRequest) (response *DescribeDedicatedHostsResponse, err error) {
	response = CreateDescribeDedicatedHostsResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeDedicatedHostsWithChan invokes the ecs.DescribeDedicatedHosts API asynchronously
// api document: https://help.aliyun.com/api/ecs/describededicatedhosts.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeDedicatedHostsWithChan(request *DescribeDedicatedHostsRequest) (<-chan *DescribeDedicatedHostsResponse, <-chan error) {
	responseChan := make(chan *DescribeDedicatedHostsResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeDedicatedHosts(request)
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

// DescribeDedicatedHostsWithCallback invokes the ecs.DescribeDedicatedHosts API asynchronously
// api document: https://help.aliyun.com/api/ecs/describededicatedhosts.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeDedicatedHostsWithCallback(request *DescribeDedicatedHostsRequest, callback func(response *DescribeDedicatedHostsResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeDedicatedHostsResponse
		var err error
		defer close(result)
		response, err = client.DescribeDedicatedHosts(request)
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

// DescribeDedicatedHostsRequest is the request struct for api DescribeDedicatedHosts
type DescribeDedicatedHostsRequest struct {
	*requests.RpcRequest
	DedicatedHostIds     string                       `position:"Query" name:"DedicatedHostIds"`
	ResourceOwnerId      requests.Integer             `position:"Query" name:"ResourceOwnerId"`
	PageNumber           requests.Integer             `position:"Query" name:"PageNumber"`
	ResourceGroupId      string                       `position:"Query" name:"ResourceGroupId"`
	LockReason           string                       `position:"Query" name:"LockReason"`
	PageSize             requests.Integer             `position:"Query" name:"PageSize"`
	DedicatedHostType    string                       `position:"Query" name:"DedicatedHostType"`
	Tag                  *[]DescribeDedicatedHostsTag `position:"Query" name:"Tag"  type:"Repeated"`
	DedicatedHostName    string                       `position:"Query" name:"DedicatedHostName"`
	ResourceOwnerAccount string                       `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string                       `position:"Query" name:"OwnerAccount"`
	OwnerId              requests.Integer             `position:"Query" name:"OwnerId"`
	ZoneId               string                       `position:"Query" name:"ZoneId"`
	Status               string                       `position:"Query" name:"Status"`
}

// DescribeDedicatedHostsTag is a repeated param struct in DescribeDedicatedHostsRequest
type DescribeDedicatedHostsTag struct {
	Value string `name:"Value"`
	Key   string `name:"Key"`
}

// DescribeDedicatedHostsResponse is the response struct for api DescribeDedicatedHosts
type DescribeDedicatedHostsResponse struct {
	*responses.BaseResponse
	RequestId      string         `json:"RequestId" xml:"RequestId"`
	TotalCount     int            `json:"TotalCount" xml:"TotalCount"`
	PageNumber     int            `json:"PageNumber" xml:"PageNumber"`
	PageSize       int            `json:"PageSize" xml:"PageSize"`
	DedicatedHosts DedicatedHosts `json:"DedicatedHosts" xml:"DedicatedHosts"`
}

// CreateDescribeDedicatedHostsRequest creates a request to invoke DescribeDedicatedHosts API
func CreateDescribeDedicatedHostsRequest() (request *DescribeDedicatedHostsRequest) {
	request = &DescribeDedicatedHostsRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Ecs", "2014-05-26", "DescribeDedicatedHosts", "ecs", "openAPI")
	return
}

// CreateDescribeDedicatedHostsResponse creates a response to parse from DescribeDedicatedHosts response
func CreateDescribeDedicatedHostsResponse() (response *DescribeDedicatedHostsResponse) {
	response = &DescribeDedicatedHostsResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
