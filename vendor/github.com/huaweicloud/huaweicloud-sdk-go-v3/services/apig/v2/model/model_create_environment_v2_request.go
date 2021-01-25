/*
 * APIG
 *
 * API网关（API Gateway）是为开发者、合作伙伴提供的高性能、高可用、高安全的API托管服务，帮助用户轻松构建、管理和发布任意规模的API。
 *
 */

package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type CreateEnvironmentV2Request struct {
	ProjectId  string  `json:"project_id"`
	InstanceId string  `json:"instance_id"`
	Body       *EnvReq `json:"body,omitempty"`
}

func (o CreateEnvironmentV2Request) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "CreateEnvironmentV2Request struct{}"
	}

	return strings.Join([]string{"CreateEnvironmentV2Request", string(data)}, " ")
}
