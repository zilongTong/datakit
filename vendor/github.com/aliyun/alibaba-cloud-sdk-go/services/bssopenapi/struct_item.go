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

// Item is a nested struct in bssopenapi response
type Item struct {
	Tax                         float64 `json:"Tax" xml:"Tax"`
	PreviousBillingCycleBalance float64 `json:"PreviousBillingCycleBalance" xml:"PreviousBillingCycleBalance"`
	PayerAccount                string  `json:"PayerAccount" xml:"PayerAccount"`
	UsageStartTime              string  `json:"UsageStartTime" xml:"UsageStartTime"`
	SuborderID                  string  `json:"SuborderID" xml:"SuborderID"`
	SolutionCode                string  `json:"SolutionCode" xml:"SolutionCode"`
	ProductDetail               string  `json:"ProductDetail" xml:"ProductDetail"`
	Promotion                   string  `json:"Promotion" xml:"Promotion"`
	Usage                       string  `json:"Usage" xml:"Usage"`
	Seller                      string  `json:"Seller" xml:"Seller"`
	ServicePeriodUnit           string  `json:"ServicePeriodUnit" xml:"ServicePeriodUnit"`
	PretaxAmountLocal           float64 `json:"PretaxAmountLocal" xml:"PretaxAmountLocal"`
	OwnerName                   string  `json:"OwnerName" xml:"OwnerName"`
	OutstandingAmount           float64 `json:"OutstandingAmount" xml:"OutstandingAmount"`
	DeductedByResourcePackage   string  `json:"DeductedByResourcePackage" xml:"DeductedByResourcePackage"`
	ProductCode                 string  `json:"ProductCode" xml:"ProductCode"`
	ListPrice                   string  `json:"ListPrice" xml:"ListPrice"`
	Quantity                    string  `json:"Quantity" xml:"Quantity"`
	InvoiceDiscount             float64 `json:"InvoiceDiscount" xml:"InvoiceDiscount"`
	BucketOwnerId               int64   `json:"BucketOwnerId" xml:"BucketOwnerId"`
	MybankPaymentAmount         float64 `json:"MybankPaymentAmount" xml:"MybankPaymentAmount"`
	PretaxGrossAmount           float64 `json:"PretaxGrossAmount" xml:"PretaxGrossAmount"`
	InstanceID                  string  `json:"InstanceID" xml:"InstanceID"`
	SplitItemName               string  `json:"SplitItemName" xml:"SplitItemName"`
	SubscribeBucket             string  `json:"SubscribeBucket" xml:"SubscribeBucket"`
	RecordID                    string  `json:"RecordID" xml:"RecordID"`
	Config                      string  `json:"Config" xml:"Config"`
	Status                      string  `json:"Status" xml:"Status"`
	Item                        string  `json:"Item" xml:"Item"`
	ProductName                 string  `json:"ProductName" xml:"ProductName"`
	Region                      string  `json:"Region" xml:"Region"`
	PaymentAmount               float64 `json:"PaymentAmount" xml:"PaymentAmount"`
	UsageEndTime                string  `json:"UsageEndTime" xml:"UsageEndTime"`
	RoundDownDiscount           string  `json:"RoundDownDiscount" xml:"RoundDownDiscount"`
	SolutionID                  string  `json:"SolutionID" xml:"SolutionID"`
	ClearedTime                 string  `json:"ClearedTime" xml:"ClearedTime"`
	InternetIP                  string  `json:"InternetIP" xml:"InternetIP"`
	PaymentTime                 string  `json:"PaymentTime" xml:"PaymentTime"`
	CreateTime                  string  `json:"CreateTime" xml:"CreateTime"`
	LinkedCustomerOrderID       string  `json:"LinkedCustomerOrderID" xml:"LinkedCustomerOrderID"`
	CostUnit                    string  `json:"CostUnit" xml:"CostUnit"`
	AfterTaxAmount              float64 `json:"AfterTaxAmount" xml:"AfterTaxAmount"`
	SubOrderId                  string  `json:"SubOrderId" xml:"SubOrderId"`
	ResourceGroup               string  `json:"ResourceGroup" xml:"ResourceGroup"`
	SubscribeTime               string  `json:"SubscribeTime" xml:"SubscribeTime"`
	BillingType                 string  `json:"BillingType" xml:"BillingType"`
	InstanceSpec                string  `json:"InstanceSpec" xml:"InstanceSpec"`
	Tag                         string  `json:"Tag" xml:"Tag"`
	OwnerID                     string  `json:"OwnerID" xml:"OwnerID"`
	SolutionName                string  `json:"SolutionName" xml:"SolutionName"`
	NickName                    string  `json:"NickName" xml:"NickName"`
	SubscriptionType            string  `json:"SubscriptionType" xml:"SubscriptionType"`
	InstanceConfig              string  `json:"InstanceConfig" xml:"InstanceConfig"`
	DeductedByCashCoupons       float64 `json:"DeductedByCashCoupons" xml:"DeductedByCashCoupons"`
	ServicePeriod               string  `json:"ServicePeriod" xml:"ServicePeriod"`
	InvoiceNo                   string  `json:"InvoiceNo" xml:"InvoiceNo"`
	MultAccountRelSubscribe     string  `json:"MultAccountRelSubscribe" xml:"MultAccountRelSubscribe"`
	DiscountAmount              float64 `json:"DiscountAmount" xml:"DiscountAmount"`
	BillID                      string  `json:"BillID" xml:"BillID"`
	ListPriceUnit               string  `json:"ListPriceUnit" xml:"ListPriceUnit"`
	PaymentCurrency             string  `json:"PaymentCurrency" xml:"PaymentCurrency"`
	UsageUnit                   string  `json:"UsageUnit" xml:"UsageUnit"`
	SplitItemID                 string  `json:"SplitItemID" xml:"SplitItemID"`
	SubscribeType               string  `json:"SubscribeType" xml:"SubscribeType"`
	SubscribeLanguage           string  `json:"SubscribeLanguage" xml:"SubscribeLanguage"`
	ProductType                 string  `json:"ProductType" xml:"ProductType"`
	AccountDiscount             float64 `json:"AccountDiscount" xml:"AccountDiscount"`
	Currency                    string  `json:"Currency" xml:"Currency"`
	DeductedByPrepaidCard       float64 `json:"DeductedByPrepaidCard" xml:"DeductedByPrepaidCard"`
	OriginalOrderID             string  `json:"OriginalOrderID" xml:"OriginalOrderID"`
	PaymentTransactionID        string  `json:"PaymentTransactionID" xml:"PaymentTransactionID"`
	OrderID                     string  `json:"OrderID" xml:"OrderID"`
	OrderType                   string  `json:"OrderType" xml:"OrderType"`
	BillingDate                 string  `json:"BillingDate" xml:"BillingDate"`
	DeductedByCoupons           float64 `json:"DeductedByCoupons" xml:"DeductedByCoupons"`
	BillingItem                 string  `json:"BillingItem" xml:"BillingItem"`
	Zone                        string  `json:"Zone" xml:"Zone"`
	ChargeDiscount              float64 `json:"ChargeDiscount" xml:"ChargeDiscount"`
	PretaxAmount                float64 `json:"PretaxAmount" xml:"PretaxAmount"`
	IntranetIP                  string  `json:"IntranetIP" xml:"IntranetIP"`
}
