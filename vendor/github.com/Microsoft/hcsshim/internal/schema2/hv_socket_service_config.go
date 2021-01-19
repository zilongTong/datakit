/*
 * HCS API
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: 2.1
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package hcsschema

type HvSocketServiceConfig struct {

	//  SDDL string that HvSocket will check before allowing a host process to bind  to this specific service.  If not specified, defaults to the system DefaultBindSecurityDescriptor, defined in  HvSocketSystemWpConfig in V1.
	BindSecurityDescriptor string `json:"BindSecurityDescriptor,omitempty"`

	//  SDDL string that HvSocket will check before allowing a host process to connect  to this specific service.  If not specified, defaults to the system DefaultConnectSecurityDescriptor, defined in  HvSocketSystemWpConfig in V1.
	ConnectSecurityDescriptor string `json:"ConnectSecurityDescriptor,omitempty"`

	//  If true, HvSocket will process wildcard binds for this service/system combination.  Wildcard binds are secured in the registry at  SOFTWARE/Microsoft/Windows NT/CurrentVersion/Virtualization/HvSocket/WildcardDescriptors
	AllowWildcardBinds bool `json:"AllowWildcardBinds,omitempty"`
}
