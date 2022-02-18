package gke

import (
	"testing"

	"github.com/aquasecurity/defsec/parsers/types"
	"github.com/aquasecurity/defsec/providers/google/gke"
	"github.com/aquasecurity/defsec/rules"
	"github.com/aquasecurity/defsec/state"
	"github.com/stretchr/testify/assert"
)

func TestCheckEnableAutoRepair(t *testing.T) {
	tests := []struct {
		name     string
		input    gke.GKE
		expected bool
	}{
		{
			name: "Node pool auto repair disabled",
			input: gke.GKE{
				Metadata: types.NewTestMetadata(),
				Clusters: []gke.Cluster{
					{
						Metadata: types.NewTestMetadata(),
						NodePools: []gke.NodePool{
							{
								Metadata: types.NewTestMetadata(),
								Management: gke.Management{
									Metadata:         types.NewTestMetadata(),
									EnableAutoRepair: types.Bool(false, types.NewTestMetadata()),
								},
							},
						},
					},
				},
			},
			expected: true,
		},
		{
			name: "Node pool auto repair enabled",
			input: gke.GKE{
				Metadata: types.NewTestMetadata(),
				Clusters: []gke.Cluster{
					{
						Metadata: types.NewTestMetadata(),
						NodePools: []gke.NodePool{
							{
								Metadata: types.NewTestMetadata(),
								Management: gke.Management{
									Metadata:         types.NewTestMetadata(),
									EnableAutoRepair: types.Bool(true, types.NewTestMetadata()),
								},
							},
						},
					},
				},
			},
			expected: false,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			var testState state.State
			testState.Google.GKE = test.input
			results := CheckEnableAutoRepair.Evaluate(&testState)
			var found bool
			for _, result := range results {
				if result.Status() != rules.StatusPassed && result.Rule().LongID() == CheckEnableAutoRepair.Rule().LongID() {
					found = true
				}
			}
			if test.expected {
				assert.True(t, found, "Rule should have been found")
			} else {
				assert.False(t, found, "Rule should not have been found")
			}
		})
	}
}
