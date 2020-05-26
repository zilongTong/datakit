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

// RouterInterfaceType is a nested struct in ecs response
type RouterInterfaceType struct {
	RouterInterfaceId               string `json:"RouterInterfaceId" xml:"RouterInterfaceId"`
	OppositeRegionId                string `json:"OppositeRegionId" xml:"OppositeRegionId"`
	Role                            string `json:"Role" xml:"Role"`
	Spec                            string `json:"Spec" xml:"Spec"`
	Name                            string `json:"Name" xml:"Name"`
	Description                     string `json:"Description" xml:"Description"`
	RouterId                        string `json:"RouterId" xml:"RouterId"`
	RouterType                      string `json:"RouterType" xml:"RouterType"`
	CreationTime                    string `json:"CreationTime" xml:"CreationTime"`
	EndTime                         string `json:"EndTime" xml:"EndTime"`
	ChargeType                      string `json:"ChargeType" xml:"ChargeType"`
	Status                          string `json:"Status" xml:"Status"`
	BusinessStatus                  string `json:"BusinessStatus" xml:"BusinessStatus"`
	ConnectedTime                   string `json:"ConnectedTime" xml:"ConnectedTime"`
	OppositeInterfaceId             string `json:"OppositeInterfaceId" xml:"OppositeInterfaceId"`
	OppositeInterfaceSpec           string `json:"OppositeInterfaceSpec" xml:"OppositeInterfaceSpec"`
	OppositeInterfaceStatus         string `json:"OppositeInterfaceStatus" xml:"OppositeInterfaceStatus"`
	OppositeInterfaceBusinessStatus string `json:"OppositeInterfaceBusinessStatus" xml:"OppositeInterfaceBusinessStatus"`
	OppositeRouterId                string `json:"OppositeRouterId" xml:"OppositeRouterId"`
	OppositeRouterType              string `json:"OppositeRouterType" xml:"OppositeRouterType"`
	OppositeInterfaceOwnerId        string `json:"OppositeInterfaceOwnerId" xml:"OppositeInterfaceOwnerId"`
	AccessPointId                   string `json:"AccessPointId" xml:"AccessPointId"`
	OppositeAccessPointId           string `json:"OppositeAccessPointId" xml:"OppositeAccessPointId"`
	HealthCheckSourceIp             string `json:"HealthCheckSourceIp" xml:"HealthCheckSourceIp"`
	HealthCheckTargetIp             string `json:"HealthCheckTargetIp" xml:"HealthCheckTargetIp"`
}
