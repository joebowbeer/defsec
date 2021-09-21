package functionapp

import (
	"github.com/aquasecurity/defsec/provider"
	"github.com/aquasecurity/defsec/rules"
	"github.com/aquasecurity/defsec/severity"
	"github.com/aquasecurity/defsec/state"
)

var CheckAuthenticationEnabled = rules.Register(
	rules.Rule{
		Provider:    provider.AzureProvider,
		Service:     "functionapp",
		ShortCode:   "authentication-enabled",
		Summary:     "Function App authentication is activated",
		Impact:      "Anonymous HTTP requests will be accepted",
		Resolution:  "Enable authentication to prevent anonymous request being accepted",
		Explanation: `Enabling authentication ensures that all communications in the application are authenticated. The auth_settings block needs to be filled out with the appropriate auth backend settings`,
		Links: []string{ 
		},
		Severity: severity.Low,
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
