package args

import (
	"encoding/json"
	"fmt"
	"strconv"
	"strings"
)

const TCRedirectTapName = "TC_REDIRECT_TAP_NAME"
const TCRedirectTapUID = "TC_REDIRECT_TAP_UID"
const TCRedirectTapGID = "TC_REDIRECT_TAP_GID"

type Args struct {
	TapName string `json:"tap_name"`
	TapUID  *int   `json:"uid"`
	TapGID  *int   `json:"gid"`
}

// ExtractArgs parses arguments from the JSON configuration and then overlays
// CLI args passed to us from the CNI_ARGS variable
func ExtractArgs(stdinData []byte, envArgs string) (*Args, error) {

	a := &Args{}
	if err := json.Unmarshal(stdinData, &a); err != nil {
		return nil, fmt.Errorf("Invalid cni config: %w", err)
	}

	if envArgs != "" {
		argumentsPairs := strings.Split(envArgs, ";")
		for _, pairStr := range argumentsPairs {
			pair := strings.SplitN(pairStr, "=", 2)
			if len(pair) < 2 {
				return a, fmt.Errorf("Invalid cni arguments format, %q", pairStr)
			}
			switch pair[0] {
			case TCRedirectTapName:
				a.TapName = pair[1]

			case TCRedirectTapUID:
				tapUIDVal := pair[1]
				tapUID, err := strconv.Atoi(tapUIDVal)
				if err != nil {
					return nil, fmt.Errorf("tapUID should be numeric convertible, got %q: %w", tapUIDVal, err)
				}
				a.TapUID = &tapUID

			case TCRedirectTapGID:
				tapGIDVal := pair[1]
				tapGID, err := strconv.Atoi(tapGIDVal)
				if err != nil {
					return nil, fmt.Errorf("tapGID should be numeric convertible, got %q: %w", tapGIDVal, err)
				}
				a.TapGID = &tapGID

			default:
				// support IgnoreUnknown by ignoring other args
			}
		}
	}

	return a, nil
}
