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

// ModifySecurityGroupEgressRule invokes the ecs.ModifySecurityGroupEgressRule API synchronously
// api document: https://help.aliyun.com/api/ecs/modifysecuritygroupegressrule.html
func (client *Client) ModifySecurityGroupEgressRule(request *ModifySecurityGroupEgressRuleRequest) (response *ModifySecurityGroupEgressRuleResponse, err error) {
	response = CreateModifySecurityGroupEgressRuleResponse()
	err = client.DoAction(request, response)
	return
}

// ModifySecurityGroupEgressRuleWithChan invokes the ecs.ModifySecurityGroupEgressRule API asynchronously
// api document: https://help.aliyun.com/api/ecs/modifysecuritygroupegressrule.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) ModifySecurityGroupEgressRuleWithChan(request *ModifySecurityGroupEgressRuleRequest) (<-chan *ModifySecurityGroupEgressRuleResponse, <-chan error) {
	responseChan := make(chan *ModifySecurityGroupEgressRuleResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ModifySecurityGroupEgressRule(request)
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

// ModifySecurityGroupEgressRuleWithCallback invokes the ecs.ModifySecurityGroupEgressRule API asynchronously
// api document: https://help.aliyun.com/api/ecs/modifysecuritygroupegressrule.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) ModifySecurityGroupEgressRuleWithCallback(request *ModifySecurityGroupEgressRuleRequest, callback func(response *ModifySecurityGroupEgressRuleResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ModifySecurityGroupEgressRuleResponse
		var err error
		defer close(result)
		response, err = client.ModifySecurityGroupEgressRule(request)
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

// ModifySecurityGroupEgressRuleRequest is the request struct for api ModifySecurityGroupEgressRule
type ModifySecurityGroupEgressRuleRequest struct {
	*requests.RpcRequest
	NicType               string           `position:"Query" name:"NicType"`
	ResourceOwnerId       requests.Integer `position:"Query" name:"ResourceOwnerId"`
	SourcePortRange       string           `position:"Query" name:"SourcePortRange"`
	ClientToken           string           `position:"Query" name:"ClientToken"`
	SecurityGroupId       string           `position:"Query" name:"SecurityGroupId"`
	Description           string           `position:"Query" name:"Description"`
	Ipv6DestCidrIp        string           `position:"Query" name:"Ipv6DestCidrIp"`
	Ipv6SourceCidrIp      string           `position:"Query" name:"Ipv6SourceCidrIp"`
	Policy                string           `position:"Query" name:"Policy"`
	PortRange             string           `position:"Query" name:"PortRange"`
	ResourceOwnerAccount  string           `position:"Query" name:"ResourceOwnerAccount"`
	IpProtocol            string           `position:"Query" name:"IpProtocol"`
	OwnerAccount          string           `position:"Query" name:"OwnerAccount"`
	SourceCidrIp          string           `position:"Query" name:"SourceCidrIp"`
	DestGroupId           string           `position:"Query" name:"DestGroupId"`
	OwnerId               requests.Integer `position:"Query" name:"OwnerId"`
	DestGroupOwnerAccount string           `position:"Query" name:"DestGroupOwnerAccount"`
	Priority              string           `position:"Query" name:"Priority"`
	DestCidrIp            string           `position:"Query" name:"DestCidrIp"`
	DestGroupOwnerId      requests.Integer `position:"Query" name:"DestGroupOwnerId"`
}

// ModifySecurityGroupEgressRuleResponse is the response struct for api ModifySecurityGroupEgressRule
type ModifySecurityGroupEgressRuleResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateModifySecurityGroupEgressRuleRequest creates a request to invoke ModifySecurityGroupEgressRule API
func CreateModifySecurityGroupEgressRuleRequest() (request *ModifySecurityGroupEgressRuleRequest) {
	request = &ModifySecurityGroupEgressRuleRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Ecs", "2014-05-26", "ModifySecurityGroupEgressRule", "ecs", "openAPI")
	request.Method = requests.POST
	return
}

// CreateModifySecurityGroupEgressRuleResponse creates a response to parse from ModifySecurityGroupEgressRule response
func CreateModifySecurityGroupEgressRuleResponse() (response *ModifySecurityGroupEgressRuleResponse) {
	response = &ModifySecurityGroupEgressRuleResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
