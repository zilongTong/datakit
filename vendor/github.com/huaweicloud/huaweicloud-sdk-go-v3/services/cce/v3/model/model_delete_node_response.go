/*
 * CCE
 *
 * CCE开放API
 *
 */

package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type DeleteNodeResponse struct {
	// API版本，固定值“v3”，该值不可修改。
	ApiVersion *string `json:"apiVersion,omitempty"`
	// API类型，固定值“Node”，该值不可修改。
	Kind           *string       `json:"kind,omitempty"`
	Metadata       *NodeMetadata `json:"metadata,omitempty"`
	Spec           *V3NodeSpec   `json:"spec,omitempty"`
	Status         *V3NodeStatus `json:"status,omitempty"`
	HttpStatusCode int           `json:"-"`
}

func (o DeleteNodeResponse) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "DeleteNodeResponse struct{}"
	}

	return strings.Join([]string{"DeleteNodeResponse", string(data)}, " ")
}
