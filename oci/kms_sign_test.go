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
	SignRequiredOnlyResource = SignResourceDependencies +
		generateResourceFromRepresentationMap("oci_kms_sign", "test_sign", Required, Create, signRepresentation)

	signRepresentation = map[string]interface{}{
		"crypto_endpoint":   Representation{repType: Required, create: `${data.oci_kms_vault.test_vault.crypto_endpoint}`},
		"key_id":            Representation{repType: Required, create: `${lookup(data.oci_kms_keys.test_keys_dependency_RSA.keys[0], "id")}`},
		"message":           Representation{repType: Required, create: `message`},
		"signing_algorithm": Representation{repType: Required, create: `SHA_224_RSA_PKCS1_V1_5`},
		"message_type":      Representation{repType: Optional, create: `RAW`},
	}

	SignResourceDependencies = KeyResourceDependencyConfig
)

// issue-routing-tag: kms/default
func TestKmsSignResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestKmsSignResource_basic")
	defer httpreplay.SaveScenario()

	config := testProviderConfig()

	compartmentId := getEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	resourceName := "oci_kms_sign.test_sign"

	// Save TF content to create resource with optional properties. This has to be exactly the same as the config part in the "create with optionals" step in the test.
	saveConfigContent(config+compartmentIdVariableStr+SignResourceDependencies+
		generateResourceFromRepresentationMap("oci_kms_sign", "test_sign", Optional, Create, signRepresentation), "keymanagement", "sign", t)

	ResourceTest(t, nil, []resource.TestStep{
		// verify create
		{
			Config: config + compartmentIdVariableStr + SignResourceDependencies +
				generateResourceFromRepresentationMap("oci_kms_sign", "test_sign", Required, Create, signRepresentation),
			Check: ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "crypto_endpoint"),
				resource.TestCheckResourceAttrSet(resourceName, "key_id"),
				resource.TestCheckResourceAttr(resourceName, "message", "message"),
				resource.TestCheckResourceAttr(resourceName, "signing_algorithm", "SHA_224_RSA_PKCS1_V1_5"),
			),
		},

		// delete before next create
		{
			Config: config + compartmentIdVariableStr + SignResourceDependencies,
		},
		// verify create with optionals
		{
			Config: config + compartmentIdVariableStr + SignResourceDependencies +
				generateResourceFromRepresentationMap("oci_kms_sign", "test_sign", Optional, Create, signRepresentation),
			Check: ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "crypto_endpoint"),
				resource.TestCheckResourceAttrSet(resourceName, "key_id"),
				resource.TestCheckResourceAttrSet(resourceName, "key_version_id"),
				resource.TestCheckResourceAttr(resourceName, "message", "message"),
				resource.TestCheckResourceAttr(resourceName, "message_type", "RAW"),
				resource.TestCheckResourceAttrSet(resourceName, "signature"),
				resource.TestCheckResourceAttr(resourceName, "signing_algorithm", "SHA_224_RSA_PKCS1_V1_5"),
			),
		},
	})
}
