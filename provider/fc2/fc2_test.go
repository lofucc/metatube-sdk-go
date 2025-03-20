package fc2

import (
	"testing"

	"github.com/lofucc/metatube-sdk-go/provider/internal/testkit"
)

func TestFC2_GetMovieInfoByID(t *testing.T) {
	testkit.Test(t, New, []string{
		"406996",
		"2812904",
	})
}
