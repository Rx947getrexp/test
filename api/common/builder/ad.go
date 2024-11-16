package builder

import (
	"encoding/json"
	"go-speed/api/types"
	"go-speed/global"
	"strings"
)

func BuildSlotLocations(in string) (out []types.SlotLocationItem) {
	in = strings.TrimSpace(in)
	if len(in) == 0 {
		return
	}
	err := json.Unmarshal([]byte(in), &out)
	if err != nil {
		global.Logger.Err(err).Msgf("Unmarshal failed, in: %s", in)
		return
	}
	return out
}

func BuildTargetUrls(in string) (out []types.TargetUrlItem) {
	in = strings.TrimSpace(in)
	if len(in) == 0 {
		return
	}
	err := json.Unmarshal([]byte(in), &out)
	if err != nil {
		global.Logger.Err(err).Msgf("Unmarshal failed, in: %s", in)
		return
	}
	return out
}

func BuildStringArray(in string) (out []string) {
	in = strings.TrimSpace(in)
	if len(in) == 0 {
		return
	}
	err := json.Unmarshal([]byte(in), &out)
	if err != nil {
		global.Logger.Err(err).Msgf("Unmarshal failed, in: %s", in)
		return
	}
	return out
}

func BuildIntArray(in string) (out []int) {
	in = strings.TrimSpace(in)
	if len(in) == 0 {
		return
	}
	err := json.Unmarshal([]byte(in), &out)
	if err != nil {
		global.Logger.Err(err).Msgf("Unmarshal failed, in: %s", in)
		return
	}
	return out
}
