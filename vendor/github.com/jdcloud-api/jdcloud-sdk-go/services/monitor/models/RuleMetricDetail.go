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

package models


type RuleMetricDetail struct {

    /* 指标的计算单位，比如bit/s、%、k等 (Optional) */
    CalculateUnit string `json:"calculateUnit"`

    /* 维度标识 (Optional) */
    Dimension string `json:"dimension"`

    /* 监控项唯一标识，可根据DescribeMetricsForCreateAlarm接口查询各产品线可用的监控项（创建规则时使用Metric字段）。格式：metric:downsample (Optional) */
    Metric string `json:"metric"`

    /* 监控项名称 (Optional) */
    MetricName string `json:"metricName"`

    /* 产品标识 (Optional) */
    Product string `json:"product"`

    /* 产品线标识 (Optional) */
    ServiceCode string `json:"serviceCode"`
}
