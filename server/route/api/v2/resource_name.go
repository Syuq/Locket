package v2

import (
	"fmt"
	"strings"

	"github.com/pkg/errors"

	"github.com/Syuq/Locket/internal/util"
)

const (
	WorkspaceSettingNamePrefix = "settings/"
	UserNamePrefix             = "users/"
	LocketNamePrefix             = "lockets/"
	ResourceNamePrefix         = "resources/"
	InboxNamePrefix            = "inboxes/"
)

// GetNameParentTokens returns the tokens from a resource name.
func GetNameParentTokens(name string, tokenPrefixes ...string) ([]string, error) {
	parts := strings.Split(name, "/")
	if len(parts) != 2*len(tokenPrefixes) {
		return nil, errors.Errorf("invalid request %q", name)
	}

	var tokens []string
	for i, tokenPrefix := range tokenPrefixes {
		if fmt.Sprintf("%s/", parts[2*i]) != tokenPrefix {
			return nil, errors.Errorf("invalid prefix %q in request %q", tokenPrefix, name)
		}
		if parts[2*i+1] == "" {
			return nil, errors.Errorf("invalid request %q with empty prefix %q", name, tokenPrefix)
		}
		tokens = append(tokens, parts[2*i+1])
	}
	return tokens, nil
}

func ExtractWorkspaceSettingKeyFromName(name string) (string, error) {
	tokens, err := GetNameParentTokens(name, WorkspaceSettingNamePrefix)
	if err != nil {
		return "", err
	}
	return tokens[0], nil
}

// ExtractUserIDFromName returns the uid from a resource name.
func ExtractUserIDFromName(name string) (int32, error) {
	tokens, err := GetNameParentTokens(name, UserNamePrefix)
	if err != nil {
		return 0, err
	}
	id, err := util.ConvertStringToInt32(tokens[0])
	if err != nil {
		return 0, errors.Errorf("invalid user ID %q", tokens[0])
	}
	return id, nil
}

// ExtractLocketIDFromName returns the locket ID from a resource name.
func ExtractLocketIDFromName(name string) (int32, error) {
	tokens, err := GetNameParentTokens(name, LocketNamePrefix)
	if err != nil {
		return 0, err
	}
	id, err := util.ConvertStringToInt32(tokens[0])
	if err != nil {
		return 0, errors.Errorf("invalid locket ID %q", tokens[0])
	}
	return id, nil
}

// ExtractResourceIDFromName returns the resource ID from a resource name.
func ExtractResourceIDFromName(name string) (int32, error) {
	tokens, err := GetNameParentTokens(name, ResourceNamePrefix)
	if err != nil {
		return 0, err
	}
	id, err := util.ConvertStringToInt32(tokens[0])
	if err != nil {
		return 0, errors.Errorf("invalid resource ID %q", tokens[0])
	}
	return id, nil
}

// ExtractInboxIDFromName returns the inbox ID from a resource name.
func ExtractInboxIDFromName(name string) (int32, error) {
	tokens, err := GetNameParentTokens(name, InboxNamePrefix)
	if err != nil {
		return 0, err
	}
	id, err := util.ConvertStringToInt32(tokens[0])
	if err != nil {
		return 0, errors.Errorf("invalid inbox ID %q", tokens[0])
	}
	return id, nil
}
