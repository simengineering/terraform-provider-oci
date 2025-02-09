// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package oci

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"

	"github.com/terraform-providers/terraform-provider-oci/httpreplay"
)

var (
	autonomousExadataInfrastructureOcpuSingularDataSourceRepresentation = map[string]interface{}{
		"autonomous_exadata_infrastructure_id": Representation{repType: Required, create: `${oci_database_autonomous_exadata_infrastructure.test_autonomous_exadata_infrastructure.id}`},
	}

	AutonomousExadataInfrastructureOcpuResourceConfig = generateResourceFromRepresentationMap("oci_database_autonomous_exadata_infrastructure", "test_autonomous_exadata_infrastructure", Required, Create, autonomousExadataInfrastructureRepresentation) +
		ExadataBaseDependencies +
		generateResourceFromRepresentationMap("oci_core_network_security_group", "test_network_security_group", Required, Create, getUpdatedRepresentationCopy("vcn_id", Representation{repType: Required, create: `${oci_core_virtual_network.t.id}`}, networkSecurityGroupRepresentation)) +
		generateResourceFromRepresentationMap("oci_core_network_security_group", "test_network_security_group2", Required, Create, getUpdatedRepresentationCopy("vcn_id", Representation{repType: Required, create: `${oci_core_virtual_network.t.id}`}, networkSecurityGroupRepresentation))
)

// issue-routing-tag: database/default
func TestDatabaseAutonomousExadataInfrastructureOcpuResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestDatabaseAutonomousExadataInfrastructureOcpuResource_basic")
	defer httpreplay.SaveScenario()

	config := testProviderConfig()

	compartmentId := getEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	singularDatasourceName := "data.oci_database_autonomous_exadata_infrastructure_ocpu.test_autonomous_exadata_infrastructure_ocpu"

	saveConfigContent("", "", "", t)

	ResourceTest(t, nil, []resource.TestStep{
		// verify singular datasource
		{
			Config: config +
				generateDataSourceFromRepresentationMap("oci_database_autonomous_exadata_infrastructure_ocpu", "test_autonomous_exadata_infrastructure_ocpu", Required, Create, autonomousExadataInfrastructureOcpuSingularDataSourceRepresentation) +
				compartmentIdVariableStr + AutonomousExadataInfrastructureOcpuResourceConfig,
			Check: ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "autonomous_exadata_infrastructure_id"),

				resource.TestCheckResourceAttr(singularDatasourceName, "by_workload_type.#", "1"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "consumed_cpu"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "total_cpu"),
			),
		},
	})
}
