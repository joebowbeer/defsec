package apigateway

import (
	"github.com/aquasecurity/defsec/internal/rules"
	"github.com/aquasecurity/defsec/pkg/providers"
	"github.com/aquasecurity/defsec/pkg/providers/aws/apigateway"
	"github.com/aquasecurity/defsec/pkg/scan"
	"github.com/aquasecurity/defsec/pkg/severity"
	"github.com/aquasecurity/defsec/pkg/state"
)

var CheckEnableCacheEncryption = rules.Register(
	scan.Rule{
		AVDID:       "AVD-AWS-0002",
		Provider:    providers.AWSProvider,
		Service:     "api-gateway",
		ShortCode:   "enable-cache-encryption",
		Summary:     "API Gateway must have cache enabled",
		Impact:      "Data stored in the cache that is unencrypted may be vulnerable to compromise",
		Resolution:  "Enable cache encryption",
		Explanation: `Method cache encryption ensures that any sensitive data in the cache is not vulnerable to compromise in the event of interception`,
		Links:       []string{},
		Terraform: &scan.EngineMetadata{
			GoodExamples:        terraformEnableCacheEncryptionGoodExamples,
			BadExamples:         terraformEnableCacheEncryptionBadExamples,
			Links:               terraformEnableCacheEncryptionLinks,
			RemediationMarkdown: terraformEnableCacheEncryptionRemediationMarkdown,
		},
		Severity: severity.Medium,
	},
	func(s *state.State) (results scan.Results) {
		for _, api := range s.AWS.APIGateway.APIs {
			if api.IsUnmanaged() || api.ProtocolType.NotEqualTo(apigateway.ProtocolTypeREST) {
				continue
			}
			for _, stage := range api.Stages {
				if stage.IsUnmanaged() || stage.Version.NotEqualTo(1) {
					continue
				}
				if stage.RESTMethodSettings.CacheEnabled.IsFalse() {
					continue
				}
				if stage.RESTMethodSettings.CacheDataEncrypted.IsFalse() {
					results.Add(
						"Cache data is not encrypted.",
						stage.RESTMethodSettings.CacheDataEncrypted,
					)
				} else {
					results.AddPassed(&stage)
				}
			}
		}
		return
	},
)
