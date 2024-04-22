package sourceverification

import (
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	hookdeck "github.com/hookdeck/hookdeck-go-sdk"
)

func tokenIoConfigSchema() schema.SingleNestedAttribute {
	return schema.SingleNestedAttribute{
		Optional: true,
		Attributes: map[string]schema.Attribute{
			"public_key": schema.StringAttribute{
				Required:  true,
				Sensitive: true,
			},
		},
	}
}

type tokenIoSourceVerification struct {
	PublicKey types.String `tfsdk:"public_key"`
}

func (m *tokenIoSourceVerification) toPayload() *hookdeck.VerificationConfig {
	return hookdeck.NewVerificationConfigFromVerificationTokenIo(&hookdeck.VerificationTokenIo{
		Type: hookdeck.VerificationTokenIoTypeTokenio,
		Configs: &hookdeck.VerificationTokenIoConfigs{
			PublicKey: m.PublicKey.ValueString(),
		},
	})
}
