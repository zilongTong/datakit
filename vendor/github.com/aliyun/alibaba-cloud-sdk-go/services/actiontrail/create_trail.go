package actiontrail

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

// CreateTrail invokes the actiontrail.CreateTrail API synchronously
// api document: https://help.aliyun.com/api/actiontrail/createtrail.html
func (client *Client) CreateTrail(request *CreateTrailRequest) (response *CreateTrailResponse, err error) {
	response = CreateCreateTrailResponse()
	err = client.DoAction(request, response)
	return
}

// CreateTrailWithChan invokes the actiontrail.CreateTrail API asynchronously
// api document: https://help.aliyun.com/api/actiontrail/createtrail.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) CreateTrailWithChan(request *CreateTrailRequest) (<-chan *CreateTrailResponse, <-chan error) {
	responseChan := make(chan *CreateTrailResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.CreateTrail(request)
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

// CreateTrailWithCallback invokes the actiontrail.CreateTrail API asynchronously
// api document: https://help.aliyun.com/api/actiontrail/createtrail.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) CreateTrailWithCallback(request *CreateTrailRequest, callback func(response *CreateTrailResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *CreateTrailResponse
		var err error
		defer close(result)
		response, err = client.CreateTrail(request)
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

// CreateTrailRequest is the request struct for api CreateTrail
type CreateTrailRequest struct {
	*requests.RpcRequest
	SlsProjectArn   string `position:"Query" name:"SlsProjectArn"`
	SlsWriteRoleArn string `position:"Query" name:"SlsWriteRoleArn"`
	OssKeyPrefix    string `position:"Query" name:"OssKeyPrefix"`
	RoleName        string `position:"Query" name:"RoleName"`
	EventRW         string `position:"Query" name:"EventRW"`
	Name            string `position:"Query" name:"Name"`
	OssBucketName   string `position:"Query" name:"OssBucketName"`
}

// CreateTrailResponse is the response struct for api CreateTrail
type CreateTrailResponse struct {
	*responses.BaseResponse
	RequestId       string `json:"RequestId" xml:"RequestId"`
	Name            string `json:"Name" xml:"Name"`
	HomeRegion      string `json:"HomeRegion" xml:"HomeRegion"`
	OssBucketName   string `json:"OssBucketName" xml:"OssBucketName"`
	OssKeyPrefix    string `json:"OssKeyPrefix" xml:"OssKeyPrefix"`
	RoleName        string `json:"RoleName" xml:"RoleName"`
	SlsProjectArn   string `json:"SlsProjectArn" xml:"SlsProjectArn"`
	SlsWriteRoleArn string `json:"SlsWriteRoleArn" xml:"SlsWriteRoleArn"`
	EventRW         string `json:"EventRW" xml:"EventRW"`
}

// CreateCreateTrailRequest creates a request to invoke CreateTrail API
func CreateCreateTrailRequest() (request *CreateTrailRequest) {
	request = &CreateTrailRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Actiontrail", "2017-12-04", "CreateTrail", "actiontrail", "openAPI")
	return
}

// CreateCreateTrailResponse creates a response to parse from CreateTrail response
func CreateCreateTrailResponse() (response *CreateTrailResponse) {
	response = &CreateTrailResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
