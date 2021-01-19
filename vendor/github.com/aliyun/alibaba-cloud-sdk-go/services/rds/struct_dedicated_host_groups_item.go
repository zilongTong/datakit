package rds

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

// DedicatedHostGroupsItem is a nested struct in rds response
type DedicatedHostGroupsItem struct {
	DedicatedHostGroupId              string                 `json:"DedicatedHostGroupId" xml:"DedicatedHostGroupId"`
	DedicatedHostGroupDesc            string                 `json:"DedicatedHostGroupDesc" xml:"DedicatedHostGroupDesc"`
	CpuAllocationRatio                int                    `json:"CpuAllocationRatio" xml:"CpuAllocationRatio"`
	MemAllocationRatio                int                    `json:"MemAllocationRatio" xml:"MemAllocationRatio"`
	DiskAllocationRatio               int                    `json:"DiskAllocationRatio" xml:"DiskAllocationRatio"`
	AllocationPolicy                  string                 `json:"AllocationPolicy" xml:"AllocationPolicy"`
	HostReplacePolicy                 string                 `json:"HostReplacePolicy" xml:"HostReplacePolicy"`
	CreateTime                        string                 `json:"CreateTime" xml:"CreateTime"`
	VPCId                             string                 `json:"VPCId" xml:"VPCId"`
	HostNumber                        int                    `json:"HostNumber" xml:"HostNumber"`
	InstanceNumber                    int                    `json:"InstanceNumber" xml:"InstanceNumber"`
	Engine                            string                 `json:"Engine" xml:"Engine"`
	Text                              string                 `json:"Text" xml:"Text"`
	DedicatedHostCountGroupByHostType map[string]interface{} `json:"DedicatedHostCountGroupByHostType" xml:"DedicatedHostCountGroupByHostType"`
	BastionInstanceId                 string                 `json:"BastionInstanceId" xml:"BastionInstanceId"`
	OpenPermission                    string                 `json:"OpenPermission" xml:"OpenPermission"`
	MemUtility                        float64                `json:"MemUtility" xml:"MemUtility"`
	MemUsedAmount                     float64                `json:"MemUsedAmount" xml:"MemUsedAmount"`
	DiskUtility                       float64                `json:"DiskUtility" xml:"DiskUtility"`
	DiskUsedAmount                    float64                `json:"DiskUsedAmount" xml:"DiskUsedAmount"`
	CpuAllocateRation                 float64                `json:"CpuAllocateRation" xml:"CpuAllocateRation"`
	CpuAllocatedAmount                float64                `json:"CpuAllocatedAmount" xml:"CpuAllocatedAmount"`
	MemAllocateRation                 float64                `json:"MemAllocateRation" xml:"MemAllocateRation"`
	MemAllocatedAmount                float64                `json:"MemAllocatedAmount" xml:"MemAllocatedAmount"`
	DiskAllocateRation                float64                `json:"DiskAllocateRation" xml:"DiskAllocateRation"`
	DiskAllocatedAmount               float64                `json:"DiskAllocatedAmount" xml:"DiskAllocatedAmount"`
	ZoneIDList                        ZoneIDList             `json:"ZoneIDList" xml:"ZoneIDList"`
}
