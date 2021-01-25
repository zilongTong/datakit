/*
 * CBR
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 */

package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type ShowProtectableResponse struct {
	Instance       *ProtectablesResp `json:"instance,omitempty"`
	HttpStatusCode int               `json:"-"`
}

func (o ShowProtectableResponse) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ShowProtectableResponse struct{}"
	}

	return strings.Join([]string{"ShowProtectableResponse", string(data)}, " ")
}
