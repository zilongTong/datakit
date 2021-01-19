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

// CreateOrUpdateDataSource invokes the aegis.CreateOrUpdateDataSource API synchronously
// api document: https://help.aliyun.com/api/aegis/createorupdatedatasource.html
func (client *Client) CreateOrUpdateDataSource(request *CreateOrUpdateDataSourceRequest) (response *CreateOrUpdateDataSourceResponse, err error) {
	response = CreateCreateOrUpdateDataSourceResponse()
	err = client.DoAction(request, response)
	return
}

// CreateOrUpdateDataSourceWithChan invokes the aegis.CreateOrUpdateDataSource API asynchronously
// api document: https://help.aliyun.com/api/aegis/createorupdatedatasource.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) CreateOrUpdateDataSourceWithChan(request *CreateOrUpdateDataSourceRequest) (<-chan *CreateOrUpdateDataSourceResponse, <-chan error) {
	responseChan := make(chan *CreateOrUpdateDataSourceResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.CreateOrUpdateDataSource(request)
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

// CreateOrUpdateDataSourceWithCallback invokes the aegis.CreateOrUpdateDataSource API asynchronously
// api document: https://help.aliyun.com/api/aegis/createorupdatedatasource.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) CreateOrUpdateDataSourceWithCallback(request *CreateOrUpdateDataSourceRequest, callback func(response *CreateOrUpdateDataSourceResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *CreateOrUpdateDataSourceResponse
		var err error
		defer close(result)
		response, err = client.CreateOrUpdateDataSource(request)
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

// CreateOrUpdateDataSourceRequest is the request struct for api CreateOrUpdateDataSource
type CreateOrUpdateDataSourceRequest struct {
	*requests.RpcRequest
	ProjectName           string `position:"Query" name:"ProjectName"`
	SourceIp              string `position:"Query" name:"SourceIp"`
	LogStoreName          string `position:"Query" name:"LogStoreName"`
	DatasourceDescription string `position:"Query" name:"DatasourceDescription"`
	Fields                string `position:"Query" name:"Fields"`
	RegionNo              string `position:"Query" name:"RegionNo"`
}

// CreateOrUpdateDataSourceResponse is the response struct for api CreateOrUpdateDataSource
type CreateOrUpdateDataSourceResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateCreateOrUpdateDataSourceRequest creates a request to invoke CreateOrUpdateDataSource API
func CreateCreateOrUpdateDataSourceRequest() (request *CreateOrUpdateDataSourceRequest) {
	request = &CreateOrUpdateDataSourceRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("aegis", "2016-11-11", "CreateOrUpdateDataSource", "vipaegis", "openAPI")
	return
}

// CreateCreateOrUpdateDataSourceResponse creates a response to parse from CreateOrUpdateDataSource response
func CreateCreateOrUpdateDataSourceResponse() (response *CreateOrUpdateDataSourceResponse) {
	response = &CreateOrUpdateDataSourceResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
