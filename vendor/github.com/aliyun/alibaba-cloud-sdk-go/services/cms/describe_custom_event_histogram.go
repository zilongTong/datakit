package cms

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

// DescribeCustomEventHistogram invokes the cms.DescribeCustomEventHistogram API synchronously
// api document: https://help.aliyun.com/api/cms/describecustomeventhistogram.html
func (client *Client) DescribeCustomEventHistogram(request *DescribeCustomEventHistogramRequest) (response *DescribeCustomEventHistogramResponse, err error) {
	response = CreateDescribeCustomEventHistogramResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeCustomEventHistogramWithChan invokes the cms.DescribeCustomEventHistogram API asynchronously
// api document: https://help.aliyun.com/api/cms/describecustomeventhistogram.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeCustomEventHistogramWithChan(request *DescribeCustomEventHistogramRequest) (<-chan *DescribeCustomEventHistogramResponse, <-chan error) {
	responseChan := make(chan *DescribeCustomEventHistogramResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeCustomEventHistogram(request)
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

// DescribeCustomEventHistogramWithCallback invokes the cms.DescribeCustomEventHistogram API asynchronously
// api document: https://help.aliyun.com/api/cms/describecustomeventhistogram.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeCustomEventHistogramWithCallback(request *DescribeCustomEventHistogramRequest, callback func(response *DescribeCustomEventHistogramResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeCustomEventHistogramResponse
		var err error
		defer close(result)
		response, err = client.DescribeCustomEventHistogram(request)
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

// DescribeCustomEventHistogramRequest is the request struct for api DescribeCustomEventHistogram
type DescribeCustomEventHistogramRequest struct {
	*requests.RpcRequest
	EventId        string `position:"Query" name:"EventId"`
	Level          string `position:"Query" name:"Level"`
	GroupId        string `position:"Query" name:"GroupId"`
	EndTime        string `position:"Query" name:"EndTime"`
	StartTime      string `position:"Query" name:"StartTime"`
	SearchKeywords string `position:"Query" name:"SearchKeywords"`
	Name           string `position:"Query" name:"Name"`
}

// DescribeCustomEventHistogramResponse is the response struct for api DescribeCustomEventHistogram
type DescribeCustomEventHistogramResponse struct {
	*responses.BaseResponse
	Code            string          `json:"Code" xml:"Code"`
	Message         string          `json:"Message" xml:"Message"`
	RequestId       string          `json:"RequestId" xml:"RequestId"`
	Success         string          `json:"Success" xml:"Success"`
	EventHistograms EventHistograms `json:"EventHistograms" xml:"EventHistograms"`
}

// CreateDescribeCustomEventHistogramRequest creates a request to invoke DescribeCustomEventHistogram API
func CreateDescribeCustomEventHistogramRequest() (request *DescribeCustomEventHistogramRequest) {
	request = &DescribeCustomEventHistogramRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Cms", "2019-01-01", "DescribeCustomEventHistogram", "cms", "openAPI")
	request.Method = requests.POST
	return
}

// CreateDescribeCustomEventHistogramResponse creates a response to parse from DescribeCustomEventHistogram response
func CreateDescribeCustomEventHistogramResponse() (response *DescribeCustomEventHistogramResponse) {
	response = &DescribeCustomEventHistogramResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
