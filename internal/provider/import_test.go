package provider

import (
	"strings"
	"testing"
)

func TestParseImportID(t *testing.T) {
	tests := []struct {
		name    string
		id      string
		kind    string
		acct    string
		cont    string
		ws      string
		wantID  string
		wantErr string
	}{
		{
			name:   "tag happy path",
			id:     "accounts/123/containers/456/workspaces/7/tags/42",
			kind:   "tags",
			acct:   "123", cont: "456", ws: "7",
			wantID: "42",
		},
		{
			name:   "trigger happy path",
			id:     "accounts/123/containers/456/workspaces/7/triggers/99",
			kind:   "triggers",
			acct:   "123", cont: "456", ws: "7",
			wantID: "99",
		},
		{
			name:   "variable happy path",
			id:     "accounts/123/containers/456/workspaces/7/variables/55",
			kind:   "variables",
			acct:   "123", cont: "456", ws: "7",
			wantID: "55",
		},
		{
			name:   "workspace happy path",
			id:     "accounts/123/containers/456/workspaces/7",
			kind:   "workspaces",
			acct:   "123", cont: "456", ws: "7",
			wantID: "7",
		},
		{
			name: "tag wrong account",
			id:   "accounts/999/containers/456/workspaces/7/tags/42",
			kind: "tags",
			acct: "123", cont: "456", ws: "7",
			wantErr: "for account 999",
		},
		{
			name: "tag wrong container",
			id:   "accounts/123/containers/999/workspaces/7/tags/42",
			kind: "tags",
			acct: "123", cont: "456", ws: "7",
			wantErr: "for container 999",
		},
		{
			name: "tag wrong workspace",
			id:   "accounts/123/containers/456/workspaces/999/tags/42",
			kind: "tags",
			acct: "123", cont: "456", ws: "7",
			wantErr: "for workspace 999",
		},
		{
			name: "tag kind mismatch (asked for tags, got triggers)",
			id:   "accounts/123/containers/456/workspaces/7/triggers/42",
			kind: "tags",
			acct: "123", cont: "456", ws: "7",
			wantErr: "expected accounts/<account>",
		},
		{
			name: "tag too few segments",
			id:   "accounts/123/containers/456/workspaces/7",
			kind: "tags",
			acct: "123", cont: "456", ws: "7",
			wantErr: "expected accounts/<account>",
		},
		{
			name: "tag bare numeric id rejected",
			id:   "42",
			kind: "tags",
			acct: "123", cont: "456", ws: "7",
			wantErr: "expected accounts/<account>",
		},
		{
			name: "workspace wrong account",
			id:   "accounts/999/containers/456/workspaces/7",
			kind: "workspaces",
			acct: "123", cont: "456", ws: "7",
			wantErr: "for account 999",
		},
		{
			name: "workspace wrong container",
			id:   "accounts/123/containers/999/workspaces/7",
			kind: "workspaces",
			acct: "123", cont: "456", ws: "7",
			wantErr: "for container 999",
		},
		{
			name: "workspace wrong id",
			id:   "accounts/123/containers/456/workspaces/999",
			kind: "workspaces",
			acct: "123", cont: "456", ws: "7",
			wantErr: "provider's workspace is 7",
		},
		{
			name: "workspace too long",
			id:   "accounts/123/containers/456/workspaces/7/tags/42",
			kind: "workspaces",
			acct: "123", cont: "456", ws: "7",
			wantErr: "expected accounts/<account>/containers/<container>/workspaces/<id>",
		},
		{
			name: "empty id",
			id:   "",
			kind: "tags",
			acct: "123", cont: "456", ws: "7",
			wantErr: "expected accounts/<account>",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := parseImportID(tt.id, tt.kind, tt.acct, tt.cont, tt.ws)
			if tt.wantErr != "" {
				if err == nil {
					t.Fatalf("want error containing %q, got nil (returned id=%q)", tt.wantErr, got)
				}
				if !strings.Contains(err.Error(), tt.wantErr) {
					t.Errorf("error %q does not contain %q", err.Error(), tt.wantErr)
				}
				return
			}
			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}
			if got != tt.wantID {
				t.Errorf("id = %q, want %q", got, tt.wantID)
			}
		})
	}
}
