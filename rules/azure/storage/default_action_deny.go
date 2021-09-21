package storage

import (
	"github.com/aquasecurity/defsec/provider"
	"github.com/aquasecurity/defsec/rules"
	"github.com/aquasecurity/defsec/severity"
	"github.com/aquasecurity/defsec/state"
)

var CheckDefaultActionDeny = rules.Register(
	rules.Rule{
		Provider:    provider.AzureProvider,
		Service:     "storage",
		ShortCode:   "default-action-deny",
		Summary:     "The default action on Storage account network rules should be set to deny",
		Impact:      "Network rules that allow could cause data to be exposed publicly",
		Resolution:  "Set network rules to deny",
		Explanation: `The default_action for network rules should come into effect when no other rules are matched.

The default action should be set to Deny.`,
		Links: []string{ 
			"https://docs.microsoft.com/en-us/azure/firewall/rule-processing",
		},
		Severity: severity.Critical,
	},
	func(s *state.State) (results rules.Results) {
		for _, x := range s.AWS.S3.Buckets {
			if x.Encryption.Enabled.IsFalse() {
				results.Add(
					"",
					x.Encryption.Enabled,
					
				)
			}
		}
		return
	},
)
