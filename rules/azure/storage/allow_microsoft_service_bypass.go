package storage

import (
	"github.com/aquasecurity/defsec/provider"
	"github.com/aquasecurity/defsec/rules"
	"github.com/aquasecurity/defsec/severity"
	"github.com/aquasecurity/defsec/state"
)

var CheckAllowMicrosoftServiceBypass = rules.Register(
	rules.Rule{
		Provider:    provider.AzureProvider,
		Service:     "storage",
		ShortCode:   "allow-microsoft-service-bypass",
		Summary:     "Trusted Microsoft Services should have bypass access to Storage accounts",
		Impact:      "Trusted Microsoft Services won't be able to access storage account unless rules set to allow",
		Resolution:  "Allow Trusted Microsoft Services to bypass",
		Explanation: `Some Microsoft services that interact with storage accounts operate from networks that can't be granted access through network rules. 

To help this type of service work as intended, allow the set of trusted Microsoft services to bypass the network rules`,
		Links: []string{ 
			"https://docs.microsoft.com/en-us/azure/storage/common/storage-network-security#trusted-microsoft-services",
		},
		Severity: severity.High,
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
