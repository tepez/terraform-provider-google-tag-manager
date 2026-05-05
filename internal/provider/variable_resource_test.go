package provider

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-framework/types"
)

// Regression: see the matching test in tag_resource_test.go for the full
// story. Same .String() vs .ValueString() bug existed in toApiVariable.
func TestToApiVariable_VariableIdSerialization(t *testing.T) {
	tests := []struct {
		name           string
		id             types.String
		wantVariableId string
	}{
		{
			name:           "known id (Update path) serializes to bare numeric string",
			id:             types.StringValue("12"),
			wantVariableId: "12",
		},
		{
			name:           "unknown id (Create path) serializes to empty string",
			id:             types.StringUnknown(),
			wantVariableId: "",
		},
		{
			name:           "null id serializes to empty string",
			id:             types.StringNull(),
			wantVariableId: "",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := toApiVariable(resourceVariableModel{
				Name: types.StringValue("My variable"),
				Type: types.StringValue("k"),
				Id:   tt.id,
			})
			if got.VariableId != tt.wantVariableId {
				t.Errorf("VariableId = %q, want %q", got.VariableId, tt.wantVariableId)
			}
		})
	}
}
