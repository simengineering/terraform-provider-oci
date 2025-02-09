// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package oci

import (
	"fmt"
	"strconv"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"

	"github.com/terraform-providers/terraform-provider-oci/httpreplay"
)

var (
	PingProbeRequiredOnlyResource = PingProbeResourceDependencies +
		generateResourceFromRepresentationMap("oci_health_checks_ping_probe", "test_ping_probe", Required, Create, pingProbeRepresentation)

	pingProbeRepresentation = map[string]interface{}{
		"compartment_id":      Representation{repType: Required, create: `${var.compartment_id}`},
		"protocol":            Representation{repType: Required, create: `TCP`},
		"targets":             Representation{repType: Required, create: []string{`www.oracle.com`}},
		"port":                Representation{repType: Optional, create: `80`},
		"timeout_in_seconds":  Representation{repType: Optional, create: `10`},
		"vantage_point_names": Representation{repType: Optional, create: []string{`goo-chs`}},
	}

	PingProbeResourceDependencies = ""
)

// issue-routing-tag: health_checks/default
func TestHealthChecksPingProbeResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestHealthChecksPingProbeResource_basic")
	defer httpreplay.SaveScenario()

	config := testProviderConfig()

	compartmentId := getEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	resourceName := "oci_health_checks_ping_probe.test_ping_probe"

	var resId string
	// Save TF content to create resource with optional properties. This has to be exactly the same as the config part in the "create with optionals" step in the test.
	saveConfigContent(config+compartmentIdVariableStr+PingProbeResourceDependencies+
		generateResourceFromRepresentationMap("oci_health_checks_ping_probe", "test_ping_probe", Optional, Create, pingProbeRepresentation), "healthchecks", "pingProbe", t)

	ResourceTest(t, nil, []resource.TestStep{
		// verify create
		{
			Config: config + compartmentIdVariableStr + PingProbeResourceDependencies +
				generateResourceFromRepresentationMap("oci_health_checks_ping_probe", "test_ping_probe", Required, Create, pingProbeRepresentation),
			Check: ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "protocol", "TCP"),
				resource.TestCheckResourceAttr(resourceName, "targets.#", "1"),
			),
		},

		// delete before next create
		{
			Config: config + compartmentIdVariableStr + PingProbeResourceDependencies,
		},
		// verify create with optionals
		{
			Config: config + compartmentIdVariableStr + PingProbeResourceDependencies +
				generateResourceFromRepresentationMap("oci_health_checks_ping_probe", "test_ping_probe", Optional, Create, pingProbeRepresentation),
			Check: ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "port", "80"),
				resource.TestCheckResourceAttr(resourceName, "protocol", "TCP"),
				resource.TestCheckResourceAttr(resourceName, "targets.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "timeout_in_seconds", "10"),
				resource.TestCheckResourceAttr(resourceName, "vantage_point_names.#", "1"),

				func(s *terraform.State) (err error) {
					resId, err = fromInstanceState(s, resourceName, "id")
					if isEnableExportCompartment, _ := strconv.ParseBool(getEnvSettingWithDefault("enable_export_compartment", "true")); isEnableExportCompartment {
						if errExport := testExportCompartmentWithResourceName(&resId, &compartmentId, resourceName); errExport != nil {
							return errExport
						}
					}
					return err
				},
			),
		},
	})
}
