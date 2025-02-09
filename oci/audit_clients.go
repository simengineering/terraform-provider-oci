// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package oci

import (
	oci_audit "github.com/oracle/oci-go-sdk/v47/audit"

	oci_common "github.com/oracle/oci-go-sdk/v47/common"
)

func init() {
	RegisterOracleClient("oci_audit.AuditClient", &OracleClient{initClientFn: initAuditAuditClient})
}

func initAuditAuditClient(configProvider oci_common.ConfigurationProvider, configureClient ConfigureClient, serviceClientOverrides ServiceClientOverrides) (interface{}, error) {
	client, err := oci_audit.NewAuditClientWithConfigurationProvider(configProvider)
	if err != nil {
		return nil, err
	}
	err = configureClient(&client.BaseClient)
	if err != nil {
		return nil, err
	}

	if serviceClientOverrides.hostUrlOverride != "" {
		client.Host = serviceClientOverrides.hostUrlOverride
	}
	return &client, nil
}

func (m *OracleClients) auditClient() *oci_audit.AuditClient {
	return m.GetClient("oci_audit.AuditClient").(*oci_audit.AuditClient)
}
