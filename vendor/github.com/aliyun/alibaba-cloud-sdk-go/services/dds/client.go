package dds

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
	"reflect"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/auth"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/auth/credentials/provider"
)

// Client is the sdk client struct, each func corresponds to an OpenAPI
type Client struct {
	sdk.Client
}

// SetClientProperty Set Property by Reflect
func SetClientProperty(client *Client, propertyName string, propertyValue interface{}) {
	v := reflect.ValueOf(client).Elem()
	if v.FieldByName(propertyName).IsValid() && v.FieldByName(propertyName).CanSet() {
		v.FieldByName(propertyName).Set(reflect.ValueOf(propertyValue))
	}
}

// SetEndpointDataToClient Set EndpointMap and ENdpointType
func SetEndpointDataToClient(client *Client) {
	SetClientProperty(client, "EndpointMap", GetEndpointMap())
	SetClientProperty(client, "EndpointType", GetEndpointType())
}

// NewClient creates a sdk client with environment variables
func NewClient() (client *Client, err error) {
	client = &Client{}
	err = client.Init()
	SetEndpointDataToClient(client)
	return
}

// NewClientWithProvider creates a sdk client with providers
// usage: https://github.com/aliyun/alibaba-cloud-sdk-go/blob/master/docs/2-Client-EN.md
func NewClientWithProvider(regionId string, providers ...provider.Provider) (client *Client, err error) {
	client = &Client{}
	var pc provider.Provider
	if len(providers) == 0 {
		pc = provider.DefaultChain
	} else {
		pc = provider.NewProviderChain(providers)
	}
	err = client.InitWithProviderChain(regionId, pc)
	SetEndpointDataToClient(client)
	return
}

// NewClientWithOptions creates a sdk client with regionId/sdkConfig/credential
// this is the common api to create a sdk client
func NewClientWithOptions(regionId string, config *sdk.Config, credential auth.Credential) (client *Client, err error) {
	client = &Client{}
	err = client.InitWithOptions(regionId, config, credential)
	SetEndpointDataToClient(client)
	return
}

// NewClientWithAccessKey is a shortcut to create sdk client with accesskey
// usage: https://github.com/aliyun/alibaba-cloud-sdk-go/blob/master/docs/2-Client-EN.md
func NewClientWithAccessKey(regionId, accessKeyId, accessKeySecret string) (client *Client, err error) {
	client = &Client{}
	err = client.InitWithAccessKey(regionId, accessKeyId, accessKeySecret)
	SetEndpointDataToClient(client)
	return
}

// NewClientWithStsToken is a shortcut to create sdk client with sts token
// usage: https://github.com/aliyun/alibaba-cloud-sdk-go/blob/master/docs/2-Client-EN.md
func NewClientWithStsToken(regionId, stsAccessKeyId, stsAccessKeySecret, stsToken string) (client *Client, err error) {
	client = &Client{}
	err = client.InitWithStsToken(regionId, stsAccessKeyId, stsAccessKeySecret, stsToken)
	SetEndpointDataToClient(client)
	return
}

// NewClientWithRamRoleArn is a shortcut to create sdk client with ram roleArn
// usage: https://github.com/aliyun/alibaba-cloud-sdk-go/blob/master/docs/2-Client-EN.md
func NewClientWithRamRoleArn(regionId string, accessKeyId, accessKeySecret, roleArn, roleSessionName string) (client *Client, err error) {
	client = &Client{}
	err = client.InitWithRamRoleArn(regionId, accessKeyId, accessKeySecret, roleArn, roleSessionName)
	SetEndpointDataToClient(client)
	return
}

// NewClientWithRamRoleArn is a shortcut to create sdk client with ram roleArn and policy
// usage: https://github.com/aliyun/alibaba-cloud-sdk-go/blob/master/docs/2-Client-EN.md
func NewClientWithRamRoleArnAndPolicy(regionId string, accessKeyId, accessKeySecret, roleArn, roleSessionName, policy string) (client *Client, err error) {
	client = &Client{}
	err = client.InitWithRamRoleArnAndPolicy(regionId, accessKeyId, accessKeySecret, roleArn, roleSessionName, policy)
	SetEndpointDataToClient(client)
	return
}

// NewClientWithEcsRamRole is a shortcut to create sdk client with ecs ram role
// usage: https://github.com/aliyun/alibaba-cloud-sdk-go/blob/master/docs/2-Client-EN.md
func NewClientWithEcsRamRole(regionId string, roleName string) (client *Client, err error) {
	client = &Client{}
	err = client.InitWithEcsRamRole(regionId, roleName)
	SetEndpointDataToClient(client)
	return
}

// NewClientWithRsaKeyPair is a shortcut to create sdk client with rsa key pair
// usage: https://github.com/aliyun/alibaba-cloud-sdk-go/blob/master/docs/2-Client-EN.md
func NewClientWithRsaKeyPair(regionId string, publicKeyId, privateKey string, sessionExpiration int) (client *Client, err error) {
	client = &Client{}
	err = client.InitWithRsaKeyPair(regionId, publicKeyId, privateKey, sessionExpiration)
	SetEndpointDataToClient(client)
	return
}
