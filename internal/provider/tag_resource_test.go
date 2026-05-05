package provider

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-framework/types"
)

// Regression: toApiTag once used resource.Id.String() instead of
// .ValueString(). On types.String, .String() returns the framework's debug
// representation:
//   - known value "65" → `"\"65\""` (literal quotes from %q formatting)
//   - unknown          → "<unknown>"
//   - null             → "<null>"
//
// That broke every Create/Update against the GTM API. Create requests sent
// `"<unknown>"` and got `Invalid tag_id (base 10 number expected)`. Update
// requests sent the quoted form (e.g. `"\"65\""`) while the URL path had the
// bare id (`/tags/65`), and the API rejected the mismatch with
// `tag.tag_id: Mismatched key with path or parent`.
func TestToApiTag_TagIdSerialization(t *testing.T) {
	tests := []struct {
		name      string
		id        types.String
		wantTagId string
	}{
		{
			name:      "known id (Update path) serializes to bare numeric string",
			id:        types.StringValue("65"),
			wantTagId: "65",
		},
		{
			name:      "unknown id (Create path) serializes to empty string",
			id:        types.StringUnknown(),
			wantTagId: "",
		},
		{
			name:      "null id serializes to empty string",
			id:        types.StringNull(),
			wantTagId: "",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := toApiTag(resourceTagModel{
				Name: types.StringValue("My tag"),
				Type: types.StringValue("html"),
				Id:   tt.id,
			})
			if got.TagId != tt.wantTagId {
				t.Errorf("TagId = %q, want %q", got.TagId, tt.wantTagId)
			}
		})
	}
}
