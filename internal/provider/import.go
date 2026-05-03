package provider

import (
	"fmt"
	"strings"
)

// parseImportID parses a GTM API resource path and returns the bare entity ID
// (last segment). It also validates that the path's account / container /
// workspace components match what the calling provider instance is configured
// for — a mismatch means the user is trying to import a resource owned by a
// different provider alias, which would silently mis-attribute it.
//
// kindPlural must be one of: "workspaces", "tags", "triggers", "variables".
//
// Path forms accepted:
//
//	accounts/<account>/containers/<container>/workspaces/<id>                                  (workspaces)
//	accounts/<account>/containers/<container>/workspaces/<workspace>/<kindPlural>/<id>         (tags / triggers / variables)
//
// For the "workspaces" form, expectedWorkspaceId is the id the provider has
// already resolved from its workspace_name config; mismatch is an error.
func parseImportID(
	importID string,
	kindPlural string,
	expectedAccountId string,
	expectedContainerId string,
	expectedWorkspaceId string,
) (string, error) {
	parts := strings.Split(importID, "/")

	if kindPlural == "workspaces" {
		if len(parts) != 6 || parts[0] != "accounts" || parts[2] != "containers" || parts[4] != "workspaces" {
			return "", fmt.Errorf("expected accounts/<account>/containers/<container>/workspaces/<id>, got: %s", importID)
		}
		if parts[1] != expectedAccountId {
			return "", fmt.Errorf("import ID is for account %s but provider is configured for %s", parts[1], expectedAccountId)
		}
		if parts[3] != expectedContainerId {
			return "", fmt.Errorf("import ID is for container %s but provider is configured for %s", parts[3], expectedContainerId)
		}
		if parts[5] != expectedWorkspaceId {
			return "", fmt.Errorf("import ID is for workspace %s but provider's workspace is %s", parts[5], expectedWorkspaceId)
		}
		return parts[5], nil
	}

	if len(parts) != 8 || parts[0] != "accounts" || parts[2] != "containers" || parts[4] != "workspaces" || parts[6] != kindPlural {
		return "", fmt.Errorf("expected accounts/<account>/containers/<container>/workspaces/<workspace>/%s/<id>, got: %s", kindPlural, importID)
	}
	if parts[1] != expectedAccountId {
		return "", fmt.Errorf("import ID is for account %s but provider is configured for %s", parts[1], expectedAccountId)
	}
	if parts[3] != expectedContainerId {
		return "", fmt.Errorf("import ID is for container %s but provider is configured for %s", parts[3], expectedContainerId)
	}
	if parts[5] != expectedWorkspaceId {
		return "", fmt.Errorf("import ID is for workspace %s but provider's workspace is %s", parts[5], expectedWorkspaceId)
	}
	return parts[7], nil
}
