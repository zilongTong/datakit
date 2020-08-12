package elasticsearch

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

// UninstallPlugin invokes the elasticsearch.UninstallPlugin API synchronously
// api document: https://help.aliyun.com/api/elasticsearch/uninstallplugin.html
func (client *Client) UninstallPlugin(request *UninstallPluginRequest) (response *UninstallPluginResponse, err error) {
	response = CreateUninstallPluginResponse()
	err = client.DoAction(request, response)
	return
}

// UninstallPluginWithChan invokes the elasticsearch.UninstallPlugin API asynchronously
// api document: https://help.aliyun.com/api/elasticsearch/uninstallplugin.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) UninstallPluginWithChan(request *UninstallPluginRequest) (<-chan *UninstallPluginResponse, <-chan error) {
	responseChan := make(chan *UninstallPluginResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.UninstallPlugin(request)
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

// UninstallPluginWithCallback invokes the elasticsearch.UninstallPlugin API asynchronously
// api document: https://help.aliyun.com/api/elasticsearch/uninstallplugin.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) UninstallPluginWithCallback(request *UninstallPluginRequest, callback func(response *UninstallPluginResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *UninstallPluginResponse
		var err error
		defer close(result)
		response, err = client.UninstallPlugin(request)
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

// UninstallPluginRequest is the request struct for api UninstallPlugin
type UninstallPluginRequest struct {
	*requests.RoaRequest
	InstanceId  string `position:"Path" name:"InstanceId"`
	ClientToken string `position:"Query" name:"clientToken"`
}

// UninstallPluginResponse is the response struct for api UninstallPlugin
type UninstallPluginResponse struct {
	*responses.BaseResponse
	RequestId string   `json:"RequestId" xml:"RequestId"`
	Result    []string `json:"Result" xml:"Result"`
}

// CreateUninstallPluginRequest creates a request to invoke UninstallPlugin API
func CreateUninstallPluginRequest() (request *UninstallPluginRequest) {
	request = &UninstallPluginRequest{
		RoaRequest: &requests.RoaRequest{},
	}
	request.InitWithApiInfo("elasticsearch", "2017-06-13", "UninstallPlugin", "/openapi/instances/[InstanceId]/plugins/actions/uninstall", "elasticsearch", "openAPI")
	request.Method = requests.POST
	return
}

// CreateUninstallPluginResponse creates a response to parse from UninstallPlugin response
func CreateUninstallPluginResponse() (response *UninstallPluginResponse) {
	response = &UninstallPluginResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
