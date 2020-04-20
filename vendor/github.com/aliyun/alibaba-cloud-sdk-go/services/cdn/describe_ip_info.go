package cdn

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

// DescribeIpInfo invokes the cdn.DescribeIpInfo API synchronously
// api document: https://help.aliyun.com/api/cdn/describeipinfo.html
func (client *Client) DescribeIpInfo(request *DescribeIpInfoRequest) (response *DescribeIpInfoResponse, err error) {
	response = CreateDescribeIpInfoResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeIpInfoWithChan invokes the cdn.DescribeIpInfo API asynchronously
// api document: https://help.aliyun.com/api/cdn/describeipinfo.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeIpInfoWithChan(request *DescribeIpInfoRequest) (<-chan *DescribeIpInfoResponse, <-chan error) {
	responseChan := make(chan *DescribeIpInfoResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeIpInfo(request)
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

// DescribeIpInfoWithCallback invokes the cdn.DescribeIpInfo API asynchronously
// api document: https://help.aliyun.com/api/cdn/describeipinfo.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeIpInfoWithCallback(request *DescribeIpInfoRequest, callback func(response *DescribeIpInfoResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeIpInfoResponse
		var err error
		defer close(result)
		response, err = client.DescribeIpInfo(request)
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

// DescribeIpInfoRequest is the request struct for api DescribeIpInfo
type DescribeIpInfoRequest struct {
	*requests.RpcRequest
	IP            string           `position:"Query" name:"IP"`
	OwnerId       requests.Integer `position:"Query" name:"OwnerId"`
	SecurityToken string           `position:"Query" name:"SecurityToken"`
}

// DescribeIpInfoResponse is the response struct for api DescribeIpInfo
type DescribeIpInfoResponse struct {
	*responses.BaseResponse
	RequestId   string `json:"RequestId" xml:"RequestId"`
	CdnIp       string `json:"CdnIp" xml:"CdnIp"`
	ISP         string `json:"ISP" xml:"ISP"`
	IspEname    string `json:"IspEname" xml:"IspEname"`
	Region      string `json:"Region" xml:"Region"`
	RegionEname string `json:"RegionEname" xml:"RegionEname"`
}

// CreateDescribeIpInfoRequest creates a request to invoke DescribeIpInfo API
func CreateDescribeIpInfoRequest() (request *DescribeIpInfoRequest) {
	request = &DescribeIpInfoRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Cdn", "2018-05-10", "DescribeIpInfo", "", "")
	return
}

// CreateDescribeIpInfoResponse creates a response to parse from DescribeIpInfo response
func CreateDescribeIpInfoResponse() (response *DescribeIpInfoResponse) {
	response = &DescribeIpInfoResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
