package pipelines

import (
	"encoding/json"
	"fmt"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ CopySource = MicrosoftAccessSource{}

type MicrosoftAccessSource struct {
	AdditionalColumns *interface{} `json:"additionalColumns,omitempty"`
	Query             *interface{} `json:"query,omitempty"`

	// Fields inherited from CopySource

	DisableMetricsCollection *bool        `json:"disableMetricsCollection,omitempty"`
	MaxConcurrentConnections *int64       `json:"maxConcurrentConnections,omitempty"`
	SourceRetryCount         *int64       `json:"sourceRetryCount,omitempty"`
	SourceRetryWait          *interface{} `json:"sourceRetryWait,omitempty"`
	Type                     string       `json:"type"`
}

func (s MicrosoftAccessSource) CopySource() BaseCopySourceImpl {
	return BaseCopySourceImpl{
		DisableMetricsCollection: s.DisableMetricsCollection,
		MaxConcurrentConnections: s.MaxConcurrentConnections,
		SourceRetryCount:         s.SourceRetryCount,
		SourceRetryWait:          s.SourceRetryWait,
		Type:                     s.Type,
	}
}

var _ json.Marshaler = MicrosoftAccessSource{}

func (s MicrosoftAccessSource) MarshalJSON() ([]byte, error) {
	type wrapper MicrosoftAccessSource
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling MicrosoftAccessSource: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling MicrosoftAccessSource: %+v", err)
	}

	decoded["type"] = "MicrosoftAccessSource"

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling MicrosoftAccessSource: %+v", err)
	}

	return encoded, nil
}
