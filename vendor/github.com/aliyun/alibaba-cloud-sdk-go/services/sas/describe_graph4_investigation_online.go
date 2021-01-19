package sas

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

// DescribeGraph4InvestigationOnline invokes the sas.DescribeGraph4InvestigationOnline API synchronously
// api document: https://help.aliyun.com/api/sas/describegraph4investigationonline.html
func (client *Client) DescribeGraph4InvestigationOnline(request *DescribeGraph4InvestigationOnlineRequest) (response *DescribeGraph4InvestigationOnlineResponse, err error) {
	response = CreateDescribeGraph4InvestigationOnlineResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeGraph4InvestigationOnlineWithChan invokes the sas.DescribeGraph4InvestigationOnline API asynchronously
// api document: https://help.aliyun.com/api/sas/describegraph4investigationonline.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeGraph4InvestigationOnlineWithChan(request *DescribeGraph4InvestigationOnlineRequest) (<-chan *DescribeGraph4InvestigationOnlineResponse, <-chan error) {
	responseChan := make(chan *DescribeGraph4InvestigationOnlineResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeGraph4InvestigationOnline(request)
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

// DescribeGraph4InvestigationOnlineWithCallback invokes the sas.DescribeGraph4InvestigationOnline API asynchronously
// api document: https://help.aliyun.com/api/sas/describegraph4investigationonline.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeGraph4InvestigationOnlineWithCallback(request *DescribeGraph4InvestigationOnlineRequest, callback func(response *DescribeGraph4InvestigationOnlineResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeGraph4InvestigationOnlineResponse
		var err error
		defer close(result)
		response, err = client.DescribeGraph4InvestigationOnline(request)
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

// DescribeGraph4InvestigationOnlineRequest is the request struct for api DescribeGraph4InvestigationOnline
type DescribeGraph4InvestigationOnlineRequest struct {
	*requests.RpcRequest
	VertexId    string           `position:"Query" name:"VertexId"`
	AnomalyId   string           `position:"Query" name:"AnomalyId"`
	SourceIp    string           `position:"Query" name:"SourceIp"`
	AnomalyUuid string           `position:"Query" name:"AnomalyUuid"`
	Lang        string           `position:"Query" name:"Lang"`
	Direction   string           `position:"Query" name:"Direction"`
	PathLength  requests.Integer `position:"Query" name:"PathLength"`
	Namespace   string           `position:"Query" name:"Namespace"`
}

// DescribeGraph4InvestigationOnlineResponse is the response struct for api DescribeGraph4InvestigationOnline
type DescribeGraph4InvestigationOnlineResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	Data      Data   `json:"Data" xml:"Data"`
}

// CreateDescribeGraph4InvestigationOnlineRequest creates a request to invoke DescribeGraph4InvestigationOnline API
func CreateDescribeGraph4InvestigationOnlineRequest() (request *DescribeGraph4InvestigationOnlineRequest) {
	request = &DescribeGraph4InvestigationOnlineRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Sas", "2018-12-03", "DescribeGraph4InvestigationOnline", "sas", "openAPI")
	return
}

// CreateDescribeGraph4InvestigationOnlineResponse creates a response to parse from DescribeGraph4InvestigationOnline response
func CreateDescribeGraph4InvestigationOnlineResponse() (response *DescribeGraph4InvestigationOnlineResponse) {
	response = &DescribeGraph4InvestigationOnlineResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
