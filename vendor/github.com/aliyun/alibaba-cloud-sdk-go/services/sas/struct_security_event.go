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

// SecurityEvent is a nested struct in sas response
type SecurityEvent struct {
	SeriousCount    int      `json:"SeriousCount" xml:"SeriousCount"`
	SuspiciousCount int      `json:"SuspiciousCount" xml:"SuspiciousCount"`
	RemindCount     int      `json:"RemindCount" xml:"RemindCount"`
	TotalCount      int      `json:"TotalCount" xml:"TotalCount"`
	DateArray       []string `json:"DateArray" xml:"DateArray"`
	ValueArray      []string `json:"ValueArray" xml:"ValueArray"`
	LevelsOn        []string `json:"LevelsOn" xml:"LevelsOn"`
	SeriousList     []string `json:"SeriousList" xml:"SeriousList"`
	SuspiciousList  []string `json:"SuspiciousList" xml:"SuspiciousList"`
	RemindList      []string `json:"RemindList" xml:"RemindList"`
}
