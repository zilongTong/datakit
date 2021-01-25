/*
 * DDS
 *
 * API v3
 *
 */

package model

import (
	"encoding/json"

	"strings"
)

type CreateDatabaseUserRequestBody struct {
	// 数据库用户名称。 - 长度为1~64位，可以包含大写字母（A~Z）、小写字母（a~z）、数字（0~9）、中划线、下划线和点。
	UserName string `json:"user_name"`
	// 数据库用户密码。 - 长度为8~32位，必须是大写字母（A~Z）、小写字母（a~z）、数字（0~9）、特殊字符~!@#%^*-_=+?的组合。 - 建议您输入高强度密码，以提高安全性，防止出现密码被暴力破解等安全风险。
	UserPwd string `json:"user_pwd"`
	// 新用户所在的数据库，默认为“admin”。 - 长度为1~64位，可以包含大写字母（A~Z）、小写字母（a~z）、数字（0~9）、下划线。
	DbName *string      `json:"db_name,omitempty"`
	Roles  *RolesOption `json:"roles"`
}

func (o CreateDatabaseUserRequestBody) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "CreateDatabaseUserRequestBody struct{}"
	}

	return strings.Join([]string{"CreateDatabaseUserRequestBody", string(data)}, " ")
}
