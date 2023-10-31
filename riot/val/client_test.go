package val

import (
	"testing"

	log "github.com/sirupsen/logrus"

	"github.com/KnutZuidema/golio/api"
	"github.com/KnutZuidema/golio/internal"
	"github.com/KnutZuidema/golio/internal/mock"
)

func TestNewClient(t *testing.T) {
	c := NewClient(internal.NewClient(api.RegionBrasil, "key", mock.NewStatusMockDoer(200), log.StandardLogger()))
	if c == nil {
		t.Error("returned nil")
	}
}
