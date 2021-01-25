package v2

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/def"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/services/vpc/v2/model"
	"net/http"
)

func GenReqDefForAcceptVpcPeering() *def.HttpRequestDef {
	reqDefBuilder := def.NewHttpRequestDefBuilder().
		WithMethod(http.MethodPut).
		WithPath("/v2.0/vpc/peerings/{peering_id}/accept").
		WithResponse(new(model.AcceptVpcPeeringResponse)).
		WithContentType("application/json")

	// request
	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("PeeringId").
		WithJsonTag("peering_id").
		WithLocationType(def.Path))

	// response

	requestDef := reqDefBuilder.Build()
	return requestDef
}

func GenReqDefForCreatePort() *def.HttpRequestDef {
	reqDefBuilder := def.NewHttpRequestDefBuilder().
		WithMethod(http.MethodPost).
		WithPath("/v1/{project_id}/ports").
		WithResponse(new(model.CreatePortResponse)).
		WithContentType("application/json;charset=UTF-8")

	// request

	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("Body").
		WithLocationType(def.Body))

	// response

	requestDef := reqDefBuilder.Build()
	return requestDef
}

func GenReqDefForCreateSecurityGroup() *def.HttpRequestDef {
	reqDefBuilder := def.NewHttpRequestDefBuilder().
		WithMethod(http.MethodPost).
		WithPath("/v1/{project_id}/security-groups").
		WithResponse(new(model.CreateSecurityGroupResponse)).
		WithContentType("application/json;charset=UTF-8")

	// request

	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("Body").
		WithLocationType(def.Body))

	// response

	requestDef := reqDefBuilder.Build()
	return requestDef
}

func GenReqDefForCreateSecurityGroupRule() *def.HttpRequestDef {
	reqDefBuilder := def.NewHttpRequestDefBuilder().
		WithMethod(http.MethodPost).
		WithPath("/v1/{project_id}/security-group-rules").
		WithResponse(new(model.CreateSecurityGroupRuleResponse)).
		WithContentType("application/json;charset=UTF-8")

	// request

	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("Body").
		WithLocationType(def.Body))

	// response

	requestDef := reqDefBuilder.Build()
	return requestDef
}

func GenReqDefForCreateSubnet() *def.HttpRequestDef {
	reqDefBuilder := def.NewHttpRequestDefBuilder().
		WithMethod(http.MethodPost).
		WithPath("/v1/{project_id}/subnets").
		WithResponse(new(model.CreateSubnetResponse)).
		WithContentType("application/json;charset=UTF-8")

	// request

	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("Body").
		WithLocationType(def.Body))

	// response

	requestDef := reqDefBuilder.Build()
	return requestDef
}

func GenReqDefForCreateVpcPeering() *def.HttpRequestDef {
	reqDefBuilder := def.NewHttpRequestDefBuilder().
		WithMethod(http.MethodPost).
		WithPath("/v2.0/vpc/peerings").
		WithResponse(new(model.CreateVpcPeeringResponse)).
		WithContentType("application/json;charset=UTF-8")

	// request

	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("Body").
		WithLocationType(def.Body))

	// response

	requestDef := reqDefBuilder.Build()
	return requestDef
}

func GenReqDefForDeletePort() *def.HttpRequestDef {
	reqDefBuilder := def.NewHttpRequestDefBuilder().
		WithMethod(http.MethodDelete).
		WithPath("/v1/{project_id}/ports/{port_id}").
		WithResponse(new(model.DeletePortResponse)).
		WithContentType("application/json")

	// request
	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("PortId").
		WithJsonTag("port_id").
		WithLocationType(def.Path))

	// response

	requestDef := reqDefBuilder.Build()
	return requestDef
}

func GenReqDefForDeleteSecurityGroup() *def.HttpRequestDef {
	reqDefBuilder := def.NewHttpRequestDefBuilder().
		WithMethod(http.MethodDelete).
		WithPath("/v1/{project_id}/security-groups/{security_group_id}").
		WithResponse(new(model.DeleteSecurityGroupResponse)).
		WithContentType("application/json")

	// request
	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("SecurityGroupId").
		WithJsonTag("security_group_id").
		WithLocationType(def.Path))

	// response

	requestDef := reqDefBuilder.Build()
	return requestDef
}

func GenReqDefForDeleteSecurityGroupRule() *def.HttpRequestDef {
	reqDefBuilder := def.NewHttpRequestDefBuilder().
		WithMethod(http.MethodDelete).
		WithPath("/v1/{project_id}/security-group-rules/{security_group_rule_id}").
		WithResponse(new(model.DeleteSecurityGroupRuleResponse)).
		WithContentType("application/json")

	// request
	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("SecurityGroupRuleId").
		WithJsonTag("security_group_rule_id").
		WithLocationType(def.Path))

	// response

	requestDef := reqDefBuilder.Build()
	return requestDef
}

func GenReqDefForDeleteSubnet() *def.HttpRequestDef {
	reqDefBuilder := def.NewHttpRequestDefBuilder().
		WithMethod(http.MethodDelete).
		WithPath("/v1/{project_id}/vpcs/{vpc_id}/subnets/{subnet_id}").
		WithResponse(new(model.DeleteSubnetResponse)).
		WithContentType("application/json")

	// request
	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("VpcId").
		WithJsonTag("vpc_id").
		WithLocationType(def.Path))
	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("SubnetId").
		WithJsonTag("subnet_id").
		WithLocationType(def.Path))

	// response

	requestDef := reqDefBuilder.Build()
	return requestDef
}

func GenReqDefForDeleteVpcPeering() *def.HttpRequestDef {
	reqDefBuilder := def.NewHttpRequestDefBuilder().
		WithMethod(http.MethodDelete).
		WithPath("/v2.0/vpc/peerings/{peering_id}").
		WithResponse(new(model.DeleteVpcPeeringResponse)).
		WithContentType("application/json")

	// request
	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("PeeringId").
		WithJsonTag("peering_id").
		WithLocationType(def.Path))

	// response

	requestDef := reqDefBuilder.Build()
	return requestDef
}

func GenReqDefForListPorts() *def.HttpRequestDef {
	reqDefBuilder := def.NewHttpRequestDefBuilder().
		WithMethod(http.MethodGet).
		WithPath("/v1/{project_id}/ports").
		WithResponse(new(model.ListPortsResponse)).
		WithContentType("application/json")

	// request

	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("Name").
		WithJsonTag("name").
		WithLocationType(def.Query))
	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("Id").
		WithJsonTag("id").
		WithLocationType(def.Query))
	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("Limit").
		WithJsonTag("limit").
		WithLocationType(def.Query))
	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("AdminStateUp").
		WithJsonTag("admin_state_up").
		WithLocationType(def.Query))
	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("NetworkId").
		WithJsonTag("network_id").
		WithLocationType(def.Query))
	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("MacAddress").
		WithJsonTag("mac_address").
		WithLocationType(def.Query))
	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("DeviceId").
		WithJsonTag("device_id").
		WithLocationType(def.Query))
	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("DeviceOwner").
		WithJsonTag("device_owner").
		WithLocationType(def.Query))
	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("Status").
		WithJsonTag("status").
		WithLocationType(def.Query))
	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("Marker").
		WithJsonTag("marker").
		WithLocationType(def.Query))
	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("FixedIps").
		WithJsonTag("fixed_ips").
		WithLocationType(def.Query))
	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("EnterpriseProjectId").
		WithJsonTag("enterprise_project_id").
		WithLocationType(def.Query))
	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("ZoneId").
		WithJsonTag("zone_id").
		WithLocationType(def.Query))

	// response

	requestDef := reqDefBuilder.Build()
	return requestDef
}

func GenReqDefForListSecurityGroupRules() *def.HttpRequestDef {
	reqDefBuilder := def.NewHttpRequestDefBuilder().
		WithMethod(http.MethodGet).
		WithPath("/v1/{project_id}/security-group-rules").
		WithResponse(new(model.ListSecurityGroupRulesResponse)).
		WithContentType("application/json")

	// request

	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("Marker").
		WithJsonTag("marker").
		WithLocationType(def.Query))
	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("Limit").
		WithJsonTag("limit").
		WithLocationType(def.Query))
	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("SecurityGroupId").
		WithJsonTag("security_group_id").
		WithLocationType(def.Query))

	// response

	requestDef := reqDefBuilder.Build()
	return requestDef
}

func GenReqDefForListSecurityGroups() *def.HttpRequestDef {
	reqDefBuilder := def.NewHttpRequestDefBuilder().
		WithMethod(http.MethodGet).
		WithPath("/v1/{project_id}/security-groups").
		WithResponse(new(model.ListSecurityGroupsResponse)).
		WithContentType("application/json")

	// request

	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("Limit").
		WithJsonTag("limit").
		WithLocationType(def.Query))
	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("Marker").
		WithJsonTag("marker").
		WithLocationType(def.Query))
	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("VpcId").
		WithJsonTag("vpc_id").
		WithLocationType(def.Query))
	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("EnterpriseProjectId").
		WithJsonTag("enterprise_project_id").
		WithLocationType(def.Query))

	// response

	requestDef := reqDefBuilder.Build()
	return requestDef
}

func GenReqDefForListSubnets() *def.HttpRequestDef {
	reqDefBuilder := def.NewHttpRequestDefBuilder().
		WithMethod(http.MethodGet).
		WithPath("/v1/{project_id}/subnets").
		WithResponse(new(model.ListSubnetsResponse)).
		WithContentType("application/json")

	// request

	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("Limit").
		WithJsonTag("limit").
		WithLocationType(def.Query))
	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("Marker").
		WithJsonTag("marker").
		WithLocationType(def.Query))
	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("VpcId").
		WithJsonTag("vpc_id").
		WithLocationType(def.Query))
	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("Scope").
		WithJsonTag("scope").
		WithLocationType(def.Query))

	// response

	requestDef := reqDefBuilder.Build()
	return requestDef
}

func GenReqDefForListVpcPeerings() *def.HttpRequestDef {
	reqDefBuilder := def.NewHttpRequestDefBuilder().
		WithMethod(http.MethodGet).
		WithPath("/v2.0/vpc/peerings").
		WithResponse(new(model.ListVpcPeeringsResponse)).
		WithContentType("application/json")

	// request

	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("Limit").
		WithJsonTag("limit").
		WithLocationType(def.Query))
	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("Marker").
		WithJsonTag("marker").
		WithLocationType(def.Query))
	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("Id").
		WithJsonTag("id").
		WithLocationType(def.Query))
	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("Name").
		WithJsonTag("name").
		WithLocationType(def.Query))
	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("Status").
		WithJsonTag("status").
		WithLocationType(def.Query))
	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("TenantId").
		WithJsonTag("tenant_id").
		WithLocationType(def.Query))
	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("VpcId").
		WithJsonTag("vpc_id").
		WithLocationType(def.Query))

	// response

	requestDef := reqDefBuilder.Build()
	return requestDef
}

func GenReqDefForRejectVpcPeering() *def.HttpRequestDef {
	reqDefBuilder := def.NewHttpRequestDefBuilder().
		WithMethod(http.MethodPut).
		WithPath("/v2.0/vpc/peerings/{peering_id}/reject").
		WithResponse(new(model.RejectVpcPeeringResponse)).
		WithContentType("application/json")

	// request
	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("PeeringId").
		WithJsonTag("peering_id").
		WithLocationType(def.Path))

	// response

	requestDef := reqDefBuilder.Build()
	return requestDef
}

func GenReqDefForShowPort() *def.HttpRequestDef {
	reqDefBuilder := def.NewHttpRequestDefBuilder().
		WithMethod(http.MethodGet).
		WithPath("/v1/{project_id}/ports/{port_id}").
		WithResponse(new(model.ShowPortResponse)).
		WithContentType("application/json")

	// request
	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("PortId").
		WithJsonTag("port_id").
		WithLocationType(def.Path))

	// response

	requestDef := reqDefBuilder.Build()
	return requestDef
}

func GenReqDefForShowQuota() *def.HttpRequestDef {
	reqDefBuilder := def.NewHttpRequestDefBuilder().
		WithMethod(http.MethodGet).
		WithPath("/v1/{project_id}/quotas").
		WithResponse(new(model.ShowQuotaResponse)).
		WithContentType("application/json")

	// request

	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("Type").
		WithJsonTag("type").
		WithLocationType(def.Query))

	// response

	requestDef := reqDefBuilder.Build()
	return requestDef
}

func GenReqDefForShowSecurityGroup() *def.HttpRequestDef {
	reqDefBuilder := def.NewHttpRequestDefBuilder().
		WithMethod(http.MethodGet).
		WithPath("/v1/{project_id}/security-groups/{security_group_id}").
		WithResponse(new(model.ShowSecurityGroupResponse)).
		WithContentType("application/json")

	// request
	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("SecurityGroupId").
		WithJsonTag("security_group_id").
		WithLocationType(def.Path))

	// response

	requestDef := reqDefBuilder.Build()
	return requestDef
}

func GenReqDefForShowSecurityGroupRule() *def.HttpRequestDef {
	reqDefBuilder := def.NewHttpRequestDefBuilder().
		WithMethod(http.MethodGet).
		WithPath("/v1/{project_id}/security-group-rules/{security_group_rule_id}").
		WithResponse(new(model.ShowSecurityGroupRuleResponse)).
		WithContentType("application/json")

	// request
	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("SecurityGroupRuleId").
		WithJsonTag("security_group_rule_id").
		WithLocationType(def.Path))

	// response

	requestDef := reqDefBuilder.Build()
	return requestDef
}

func GenReqDefForShowSubnet() *def.HttpRequestDef {
	reqDefBuilder := def.NewHttpRequestDefBuilder().
		WithMethod(http.MethodGet).
		WithPath("/v1/{project_id}/subnets/{subnet_id}").
		WithResponse(new(model.ShowSubnetResponse)).
		WithContentType("application/json")

	// request
	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("SubnetId").
		WithJsonTag("subnet_id").
		WithLocationType(def.Path))

	// response

	requestDef := reqDefBuilder.Build()
	return requestDef
}

func GenReqDefForShowVpcPeering() *def.HttpRequestDef {
	reqDefBuilder := def.NewHttpRequestDefBuilder().
		WithMethod(http.MethodGet).
		WithPath("/v2.0/vpc/peerings/{peering_id}").
		WithResponse(new(model.ShowVpcPeeringResponse)).
		WithContentType("application/json")

	// request
	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("PeeringId").
		WithJsonTag("peering_id").
		WithLocationType(def.Path))

	// response

	requestDef := reqDefBuilder.Build()
	return requestDef
}

func GenReqDefForUpdatePort() *def.HttpRequestDef {
	reqDefBuilder := def.NewHttpRequestDefBuilder().
		WithMethod(http.MethodPut).
		WithPath("/v1/{project_id}/ports/{port_id}").
		WithResponse(new(model.UpdatePortResponse)).
		WithContentType("application/json;charset=UTF-8")

	// request
	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("PortId").
		WithJsonTag("port_id").
		WithLocationType(def.Path))

	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("Body").
		WithLocationType(def.Body))

	// response

	requestDef := reqDefBuilder.Build()
	return requestDef
}

func GenReqDefForUpdateSubnet() *def.HttpRequestDef {
	reqDefBuilder := def.NewHttpRequestDefBuilder().
		WithMethod(http.MethodPut).
		WithPath("/v1/{project_id}/vpcs/{vpc_id}/subnets/{subnet_id}").
		WithResponse(new(model.UpdateSubnetResponse)).
		WithContentType("application/json;charset=UTF-8")

	// request
	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("VpcId").
		WithJsonTag("vpc_id").
		WithLocationType(def.Path))
	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("SubnetId").
		WithJsonTag("subnet_id").
		WithLocationType(def.Path))

	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("Body").
		WithLocationType(def.Body))

	// response

	requestDef := reqDefBuilder.Build()
	return requestDef
}

func GenReqDefForUpdateVpcPeering() *def.HttpRequestDef {
	reqDefBuilder := def.NewHttpRequestDefBuilder().
		WithMethod(http.MethodPut).
		WithPath("/v2.0/vpc/peerings/{peering_id}").
		WithResponse(new(model.UpdateVpcPeeringResponse)).
		WithContentType("application/json;charset=UTF-8")

	// request
	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("PeeringId").
		WithJsonTag("peering_id").
		WithLocationType(def.Path))

	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("Body").
		WithLocationType(def.Body))

	// response

	requestDef := reqDefBuilder.Build()
	return requestDef
}

func GenReqDefForCreatePrivateip() *def.HttpRequestDef {
	reqDefBuilder := def.NewHttpRequestDefBuilder().
		WithMethod(http.MethodPost).
		WithPath("/v1/{project_id}/privateips").
		WithResponse(new(model.CreatePrivateipResponse)).
		WithContentType("application/json;charset=UTF-8")

	// request

	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("Body").
		WithLocationType(def.Body))

	// response

	requestDef := reqDefBuilder.Build()
	return requestDef
}

func GenReqDefForDeletePrivateip() *def.HttpRequestDef {
	reqDefBuilder := def.NewHttpRequestDefBuilder().
		WithMethod(http.MethodDelete).
		WithPath("/v1/{project_id}/privateips/{privateip_id}").
		WithResponse(new(model.DeletePrivateipResponse)).
		WithContentType("application/json")

	// request
	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("PrivateipId").
		WithJsonTag("privateip_id").
		WithLocationType(def.Path))

	// response

	requestDef := reqDefBuilder.Build()
	return requestDef
}

func GenReqDefForListPrivateips() *def.HttpRequestDef {
	reqDefBuilder := def.NewHttpRequestDefBuilder().
		WithMethod(http.MethodGet).
		WithPath("/v1/{project_id}/subnets/{subnet_id}/privateips").
		WithResponse(new(model.ListPrivateipsResponse)).
		WithContentType("application/json")

	// request
	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("SubnetId").
		WithJsonTag("subnet_id").
		WithLocationType(def.Path))

	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("Limit").
		WithJsonTag("limit").
		WithLocationType(def.Query))
	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("Marker").
		WithJsonTag("marker").
		WithLocationType(def.Query))

	// response

	requestDef := reqDefBuilder.Build()
	return requestDef
}

func GenReqDefForShowNetworkIpAvailabilities() *def.HttpRequestDef {
	reqDefBuilder := def.NewHttpRequestDefBuilder().
		WithMethod(http.MethodGet).
		WithPath("/v2.0/network-ip-availabilities/{network_id}").
		WithResponse(new(model.ShowNetworkIpAvailabilitiesResponse)).
		WithContentType("application/json")

	// request
	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("NetworkId").
		WithJsonTag("network_id").
		WithLocationType(def.Path))

	// response

	requestDef := reqDefBuilder.Build()
	return requestDef
}

func GenReqDefForShowPrivateip() *def.HttpRequestDef {
	reqDefBuilder := def.NewHttpRequestDefBuilder().
		WithMethod(http.MethodGet).
		WithPath("/v1/{project_id}/privateips/{privateip_id}").
		WithResponse(new(model.ShowPrivateipResponse)).
		WithContentType("application/json")

	// request
	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("PrivateipId").
		WithJsonTag("privateip_id").
		WithLocationType(def.Path))

	// response

	requestDef := reqDefBuilder.Build()
	return requestDef
}

func GenReqDefForNeutronAddFirewallRule() *def.HttpRequestDef {
	reqDefBuilder := def.NewHttpRequestDefBuilder().
		WithMethod(http.MethodPut).
		WithPath("/v2.0/fwaas/firewall_policies/{firewall_policy_id}/insert_rule").
		WithResponse(new(model.NeutronAddFirewallRuleResponse)).
		WithContentType("application/json;charset=UTF-8")

	// request
	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("FirewallPolicyId").
		WithJsonTag("firewall_policy_id").
		WithLocationType(def.Path))

	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("Body").
		WithLocationType(def.Body))

	// response

	requestDef := reqDefBuilder.Build()
	return requestDef
}

func GenReqDefForNeutronCreateFirewallGroup() *def.HttpRequestDef {
	reqDefBuilder := def.NewHttpRequestDefBuilder().
		WithMethod(http.MethodPost).
		WithPath("/v2.0/fwaas/firewall_groups").
		WithResponse(new(model.NeutronCreateFirewallGroupResponse)).
		WithContentType("application/json;charset=UTF-8")

	// request

	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("Body").
		WithLocationType(def.Body))

	// response

	requestDef := reqDefBuilder.Build()
	return requestDef
}

func GenReqDefForNeutronCreateFirewallPolicy() *def.HttpRequestDef {
	reqDefBuilder := def.NewHttpRequestDefBuilder().
		WithMethod(http.MethodPost).
		WithPath("/v2.0/fwaas/firewall_policies").
		WithResponse(new(model.NeutronCreateFirewallPolicyResponse)).
		WithContentType("application/json;charset=UTF-8")

	// request

	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("Body").
		WithLocationType(def.Body))

	// response

	requestDef := reqDefBuilder.Build()
	return requestDef
}

func GenReqDefForNeutronCreateFirewallRule() *def.HttpRequestDef {
	reqDefBuilder := def.NewHttpRequestDefBuilder().
		WithMethod(http.MethodPost).
		WithPath("/v2.0/fwaas/firewall_rules").
		WithResponse(new(model.NeutronCreateFirewallRuleResponse)).
		WithContentType("application/json;charset=UTF-8")

	// request

	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("Body").
		WithLocationType(def.Body))

	// response

	requestDef := reqDefBuilder.Build()
	return requestDef
}

func GenReqDefForNeutronDeleteFirewallGroup() *def.HttpRequestDef {
	reqDefBuilder := def.NewHttpRequestDefBuilder().
		WithMethod(http.MethodDelete).
		WithPath("/v2.0/fwaas/firewall_groups/{firewall_group_id}").
		WithResponse(new(model.NeutronDeleteFirewallGroupResponse)).
		WithContentType("application/json")

	// request
	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("FirewallGroupId").
		WithJsonTag("firewall_group_id").
		WithLocationType(def.Path))

	// response

	requestDef := reqDefBuilder.Build()
	return requestDef
}

func GenReqDefForNeutronDeleteFirewallPolicy() *def.HttpRequestDef {
	reqDefBuilder := def.NewHttpRequestDefBuilder().
		WithMethod(http.MethodDelete).
		WithPath("/v2.0/fwaas/firewall_policies/{firewall_policy_id}").
		WithResponse(new(model.NeutronDeleteFirewallPolicyResponse)).
		WithContentType("application/json")

	// request
	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("FirewallPolicyId").
		WithJsonTag("firewall_policy_id").
		WithLocationType(def.Path))

	// response

	requestDef := reqDefBuilder.Build()
	return requestDef
}

func GenReqDefForNeutronDeleteFirewallRule() *def.HttpRequestDef {
	reqDefBuilder := def.NewHttpRequestDefBuilder().
		WithMethod(http.MethodDelete).
		WithPath("/v2.0/fwaas/firewall_rules/{firewall_rule_id}").
		WithResponse(new(model.NeutronDeleteFirewallRuleResponse)).
		WithContentType("application/json")

	// request
	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("FirewallRuleId").
		WithJsonTag("firewall_rule_id").
		WithLocationType(def.Path))

	// response

	requestDef := reqDefBuilder.Build()
	return requestDef
}

func GenReqDefForNeutronListFirewallGroups() *def.HttpRequestDef {
	reqDefBuilder := def.NewHttpRequestDefBuilder().
		WithMethod(http.MethodGet).
		WithPath("/v2.0/fwaas/firewall_groups").
		WithResponse(new(model.NeutronListFirewallGroupsResponse)).
		WithContentType("application/json")

	// request

	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("Marker").
		WithJsonTag("marker").
		WithLocationType(def.Query))
	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("Limit").
		WithJsonTag("limit").
		WithLocationType(def.Query))
	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("Id").
		WithJsonTag("id").
		WithLocationType(def.Query))
	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("Name").
		WithJsonTag("name").
		WithLocationType(def.Query))
	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("Description").
		WithJsonTag("description").
		WithLocationType(def.Query))
	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("IngressFirewallPolicyId").
		WithJsonTag("ingress_firewall_policy_id").
		WithLocationType(def.Query))
	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("EgressFirewallPolicyId").
		WithJsonTag("egress_firewall_policy_id").
		WithLocationType(def.Query))

	// response

	requestDef := reqDefBuilder.Build()
	return requestDef
}

func GenReqDefForNeutronListFirewallPolicies() *def.HttpRequestDef {
	reqDefBuilder := def.NewHttpRequestDefBuilder().
		WithMethod(http.MethodGet).
		WithPath("/v2.0/fwaas/firewall_policies").
		WithResponse(new(model.NeutronListFirewallPoliciesResponse)).
		WithContentType("application/json")

	// request

	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("Limit").
		WithJsonTag("limit").
		WithLocationType(def.Query))
	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("Marker").
		WithJsonTag("marker").
		WithLocationType(def.Query))
	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("Id").
		WithJsonTag("id").
		WithLocationType(def.Query))
	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("Name").
		WithJsonTag("name").
		WithLocationType(def.Query))
	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("Description").
		WithJsonTag("description").
		WithLocationType(def.Query))
	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("TenantId").
		WithJsonTag("tenant_id").
		WithLocationType(def.Query))

	// response

	requestDef := reqDefBuilder.Build()
	return requestDef
}

func GenReqDefForNeutronListFirewallRules() *def.HttpRequestDef {
	reqDefBuilder := def.NewHttpRequestDefBuilder().
		WithMethod(http.MethodGet).
		WithPath("/v2.0/fwaas/firewall_rules").
		WithResponse(new(model.NeutronListFirewallRulesResponse)).
		WithContentType("application/json")

	// request

	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("Marker").
		WithJsonTag("marker").
		WithLocationType(def.Query))
	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("Limit").
		WithJsonTag("limit").
		WithLocationType(def.Query))
	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("Id").
		WithJsonTag("id").
		WithLocationType(def.Query))
	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("Name").
		WithJsonTag("name").
		WithLocationType(def.Query))
	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("Description").
		WithJsonTag("description").
		WithLocationType(def.Query))
	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("Action").
		WithJsonTag("action").
		WithLocationType(def.Query))
	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("TenantId").
		WithJsonTag("tenant_id").
		WithLocationType(def.Query))

	// response

	requestDef := reqDefBuilder.Build()
	return requestDef
}

func GenReqDefForNeutronRemoveFirewallRule() *def.HttpRequestDef {
	reqDefBuilder := def.NewHttpRequestDefBuilder().
		WithMethod(http.MethodPut).
		WithPath("/v2.0/fwaas/firewall_policies/{firewall_policy_id}/remove_rule").
		WithResponse(new(model.NeutronRemoveFirewallRuleResponse)).
		WithContentType("application/json;charset=UTF-8")

	// request
	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("FirewallPolicyId").
		WithJsonTag("firewall_policy_id").
		WithLocationType(def.Path))

	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("Body").
		WithLocationType(def.Body))

	// response

	requestDef := reqDefBuilder.Build()
	return requestDef
}

func GenReqDefForNeutronShowFirewallGroup() *def.HttpRequestDef {
	reqDefBuilder := def.NewHttpRequestDefBuilder().
		WithMethod(http.MethodGet).
		WithPath("/v2.0/fwaas/firewall_groups/{firewall_group_id}").
		WithResponse(new(model.NeutronShowFirewallGroupResponse)).
		WithContentType("application/json")

	// request
	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("FirewallGroupId").
		WithJsonTag("firewall_group_id").
		WithLocationType(def.Path))

	// response

	requestDef := reqDefBuilder.Build()
	return requestDef
}

func GenReqDefForNeutronShowFirewallPolicy() *def.HttpRequestDef {
	reqDefBuilder := def.NewHttpRequestDefBuilder().
		WithMethod(http.MethodGet).
		WithPath("/v2.0/fwaas/firewall_policies/{firewall_policy_id}").
		WithResponse(new(model.NeutronShowFirewallPolicyResponse)).
		WithContentType("application/json")

	// request
	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("FirewallPolicyId").
		WithJsonTag("firewall_policy_id").
		WithLocationType(def.Path))

	// response

	requestDef := reqDefBuilder.Build()
	return requestDef
}

func GenReqDefForNeutronShowFirewallRule() *def.HttpRequestDef {
	reqDefBuilder := def.NewHttpRequestDefBuilder().
		WithMethod(http.MethodGet).
		WithPath("/v2.0/fwaas/firewall_rules/{firewall_rule_id}").
		WithResponse(new(model.NeutronShowFirewallRuleResponse)).
		WithContentType("application/json")

	// request
	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("FirewallRuleId").
		WithJsonTag("firewall_rule_id").
		WithLocationType(def.Path))

	// response

	requestDef := reqDefBuilder.Build()
	return requestDef
}

func GenReqDefForNeutronUpdateFirewallGroup() *def.HttpRequestDef {
	reqDefBuilder := def.NewHttpRequestDefBuilder().
		WithMethod(http.MethodPut).
		WithPath("/v2.0/fwaas/firewall_groups/{firewall_group_id}").
		WithResponse(new(model.NeutronUpdateFirewallGroupResponse)).
		WithContentType("application/json;charset=UTF-8")

	// request
	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("FirewallGroupId").
		WithJsonTag("firewall_group_id").
		WithLocationType(def.Path))

	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("Body").
		WithLocationType(def.Body))

	// response

	requestDef := reqDefBuilder.Build()
	return requestDef
}

func GenReqDefForNeutronUpdateFirewallPolicy() *def.HttpRequestDef {
	reqDefBuilder := def.NewHttpRequestDefBuilder().
		WithMethod(http.MethodPut).
		WithPath("/v2.0/fwaas/firewall_policies/{firewall_policy_id}").
		WithResponse(new(model.NeutronUpdateFirewallPolicyResponse)).
		WithContentType("application/json;charset=UTF-8")

	// request
	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("FirewallPolicyId").
		WithJsonTag("firewall_policy_id").
		WithLocationType(def.Path))

	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("Body").
		WithLocationType(def.Body))

	// response

	requestDef := reqDefBuilder.Build()
	return requestDef
}

func GenReqDefForNeutronUpdateFirewallRule() *def.HttpRequestDef {
	reqDefBuilder := def.NewHttpRequestDefBuilder().
		WithMethod(http.MethodPut).
		WithPath("/v2.0/fwaas/firewall_rules/{firewall_rule_id}").
		WithResponse(new(model.NeutronUpdateFirewallRuleResponse)).
		WithContentType("application/json;charset=UTF-8")

	// request
	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("FirewallRuleId").
		WithJsonTag("firewall_rule_id").
		WithLocationType(def.Path))

	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("Body").
		WithLocationType(def.Body))

	// response

	requestDef := reqDefBuilder.Build()
	return requestDef
}

func GenReqDefForCreateVpc() *def.HttpRequestDef {
	reqDefBuilder := def.NewHttpRequestDefBuilder().
		WithMethod(http.MethodPost).
		WithPath("/v1/{project_id}/vpcs").
		WithResponse(new(model.CreateVpcResponse)).
		WithContentType("application/json;charset=UTF-8")

	// request

	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("Body").
		WithLocationType(def.Body))

	// response

	requestDef := reqDefBuilder.Build()
	return requestDef
}

func GenReqDefForCreateVpcRoute() *def.HttpRequestDef {
	reqDefBuilder := def.NewHttpRequestDefBuilder().
		WithMethod(http.MethodPost).
		WithPath("/v2.0/vpc/routes").
		WithResponse(new(model.CreateVpcRouteResponse)).
		WithContentType("application/json;charset=UTF-8")

	// request

	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("Body").
		WithLocationType(def.Body))

	// response

	requestDef := reqDefBuilder.Build()
	return requestDef
}

func GenReqDefForDeleteVpc() *def.HttpRequestDef {
	reqDefBuilder := def.NewHttpRequestDefBuilder().
		WithMethod(http.MethodDelete).
		WithPath("/v1/{project_id}/vpcs/{vpc_id}").
		WithResponse(new(model.DeleteVpcResponse)).
		WithContentType("application/json")

	// request
	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("VpcId").
		WithJsonTag("vpc_id").
		WithLocationType(def.Path))

	// response

	requestDef := reqDefBuilder.Build()
	return requestDef
}

func GenReqDefForDeleteVpcRoute() *def.HttpRequestDef {
	reqDefBuilder := def.NewHttpRequestDefBuilder().
		WithMethod(http.MethodDelete).
		WithPath("/v2.0/vpc/routes/{route_id}").
		WithResponse(new(model.DeleteVpcRouteResponse)).
		WithContentType("application/json")

	// request
	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("RouteId").
		WithJsonTag("route_id").
		WithLocationType(def.Path))

	// response

	requestDef := reqDefBuilder.Build()
	return requestDef
}

func GenReqDefForListVpcRoutes() *def.HttpRequestDef {
	reqDefBuilder := def.NewHttpRequestDefBuilder().
		WithMethod(http.MethodGet).
		WithPath("/v2.0/vpc/routes").
		WithResponse(new(model.ListVpcRoutesResponse)).
		WithContentType("application/json")

	// request

	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("Limit").
		WithJsonTag("limit").
		WithLocationType(def.Query))
	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("Marker").
		WithJsonTag("marker").
		WithLocationType(def.Query))
	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("Id").
		WithJsonTag("id").
		WithLocationType(def.Query))
	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("Type").
		WithJsonTag("type").
		WithLocationType(def.Query))
	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("VpcId").
		WithJsonTag("vpc_id").
		WithLocationType(def.Query))
	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("Destination").
		WithJsonTag("destination").
		WithLocationType(def.Query))
	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("TenantId").
		WithJsonTag("tenant_id").
		WithLocationType(def.Query))

	// response

	requestDef := reqDefBuilder.Build()
	return requestDef
}

func GenReqDefForListVpcs() *def.HttpRequestDef {
	reqDefBuilder := def.NewHttpRequestDefBuilder().
		WithMethod(http.MethodGet).
		WithPath("/v1/{project_id}/vpcs").
		WithResponse(new(model.ListVpcsResponse)).
		WithContentType("application/json")

	// request

	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("Limit").
		WithJsonTag("limit").
		WithLocationType(def.Query))
	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("Marker").
		WithJsonTag("marker").
		WithLocationType(def.Query))
	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("Id").
		WithJsonTag("id").
		WithLocationType(def.Query))
	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("EnterpriseProjectId").
		WithJsonTag("enterprise_project_id").
		WithLocationType(def.Query))

	// response

	requestDef := reqDefBuilder.Build()
	return requestDef
}

func GenReqDefForShowVpc() *def.HttpRequestDef {
	reqDefBuilder := def.NewHttpRequestDefBuilder().
		WithMethod(http.MethodGet).
		WithPath("/v1/{project_id}/vpcs/{vpc_id}").
		WithResponse(new(model.ShowVpcResponse)).
		WithContentType("application/json")

	// request
	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("VpcId").
		WithJsonTag("vpc_id").
		WithLocationType(def.Path))

	// response

	requestDef := reqDefBuilder.Build()
	return requestDef
}

func GenReqDefForShowVpcRoute() *def.HttpRequestDef {
	reqDefBuilder := def.NewHttpRequestDefBuilder().
		WithMethod(http.MethodGet).
		WithPath("/v2.0/vpc/routes/{route_id}").
		WithResponse(new(model.ShowVpcRouteResponse)).
		WithContentType("application/json")

	// request
	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("RouteId").
		WithJsonTag("route_id").
		WithLocationType(def.Path))

	// response

	requestDef := reqDefBuilder.Build()
	return requestDef
}

func GenReqDefForUpdateVpc() *def.HttpRequestDef {
	reqDefBuilder := def.NewHttpRequestDefBuilder().
		WithMethod(http.MethodPut).
		WithPath("/v1/{project_id}/vpcs/{vpc_id}").
		WithResponse(new(model.UpdateVpcResponse)).
		WithContentType("application/json;charset=UTF-8")

	// request
	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("VpcId").
		WithJsonTag("vpc_id").
		WithLocationType(def.Path))

	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("Body").
		WithLocationType(def.Body))

	// response

	requestDef := reqDefBuilder.Build()
	return requestDef
}
