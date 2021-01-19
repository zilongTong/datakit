// Copyright 2018 JDCLOUD.COM
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//
// NOTE: This class is auto generated by the jdcloud code generator program.

package apis

import (
    "github.com/jdcloud-api/jdcloud-sdk-go/core"
    monitor "github.com/jdcloud-api/jdcloud-sdk-go/services/monitor/models"
)

type PutMetricDataRequest struct {

    core.JDCloudRequest

    /* 数据参数 (Optional) */
    MetricDataList []monitor.MetricDataCm `json:"metricDataList"`
}

/*
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewPutMetricDataRequest(
) *PutMetricDataRequest {

	return &PutMetricDataRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/customMetrics",
			Method:  "POST",
			Header:  nil,
			Version: "v2",
		},
	}
}

/*
 * param metricDataList: 数据参数 (Optional)
 */
func NewPutMetricDataRequestWithAllParams(
    metricDataList []monitor.MetricDataCm,
) *PutMetricDataRequest {

    return &PutMetricDataRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/customMetrics",
            Method:  "POST",
            Header:  nil,
            Version: "v2",
        },
        MetricDataList: metricDataList,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewPutMetricDataRequestWithoutParam() *PutMetricDataRequest {

    return &PutMetricDataRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/customMetrics",
            Method:  "POST",
            Header:  nil,
            Version: "v2",
        },
    }
}

/* param metricDataList: 数据参数(Optional) */
func (r *PutMetricDataRequest) SetMetricDataList(metricDataList []monitor.MetricDataCm) {
    r.MetricDataList = metricDataList
}

// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r PutMetricDataRequest) GetRegionId() string {
    return ""
}

type PutMetricDataResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result PutMetricDataResult `json:"result"`
}

type PutMetricDataResult struct {
    Success bool `json:"success"`
    ErrMetricDataList []monitor.MetricDataList `json:"errMetricDataList"`
}