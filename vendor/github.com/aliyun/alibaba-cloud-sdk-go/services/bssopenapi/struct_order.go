package bssopenapi

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

// Order is a nested struct in bssopenapi response
type Order struct {
	Region            string `json:"Region" xml:"Region"`
	UsageEndTime      string `json:"UsageEndTime" xml:"UsageEndTime"`
	PaymentTime       string `json:"PaymentTime" xml:"PaymentTime"`
	CreateTime        string `json:"CreateTime" xml:"CreateTime"`
	PaymentStatus     string `json:"PaymentStatus" xml:"PaymentStatus"`
	Operator          string `json:"Operator" xml:"Operator"`
	SubOrderId        string `json:"SubOrderId" xml:"SubOrderId"`
	AfterTaxAmount    string `json:"AfterTaxAmount" xml:"AfterTaxAmount"`
	OrderId           string `json:"OrderId" xml:"OrderId"`
	OriginalConfig    string `json:"OriginalConfig" xml:"OriginalConfig"`
	PretaxGrossAmount string `json:"PretaxGrossAmount" xml:"PretaxGrossAmount"`
	SubscriptionType  string `json:"SubscriptionType" xml:"SubscriptionType"`
	Tax               string `json:"Tax" xml:"Tax"`
	UsageStartTime    string `json:"UsageStartTime" xml:"UsageStartTime"`
	PretaxAmount      string `json:"PretaxAmount" xml:"PretaxAmount"`
	PaymentCurrency   string `json:"PaymentCurrency" xml:"PaymentCurrency"`
	OrderSubType      string `json:"OrderSubType" xml:"OrderSubType"`
	ProductType       string `json:"ProductType" xml:"ProductType"`
	Currency          string `json:"Currency" xml:"Currency"`
	ProductCode       string `json:"ProductCode" xml:"ProductCode"`
	InstanceIDs       string `json:"InstanceIDs" xml:"InstanceIDs"`
	OrderType         string `json:"OrderType" xml:"OrderType"`
	Quantity          string `json:"Quantity" xml:"Quantity"`
	PretaxAmountLocal string `json:"PretaxAmountLocal" xml:"PretaxAmountLocal"`
	Config            string `json:"Config" xml:"Config"`
	RelatedOrderId    string `json:"RelatedOrderId" xml:"RelatedOrderId"`
}
