// Copyright (c) 2016, 2018, 2021, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Data Integration API
//
// Use the Data Integration Service APIs to perform common extract, load, and transform (ETL) tasks.
//

package dataintegration

import (
	"github.com/oracle/oci-go-sdk/v47/common"
)

// ExternalPublicationValidationSummaryCollection This is the collection of external publication validation  summaries. It may be a collection of lightweight details or full definitions.
type ExternalPublicationValidationSummaryCollection struct {

	// The array of external publication summaries.
	Items []ExternalPublicationValidationSummary `mandatory:"true" json:"items"`
}

func (m ExternalPublicationValidationSummaryCollection) String() string {
	return common.PointerString(m)
}
