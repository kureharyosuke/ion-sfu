package sfu

import "time"

const (
	quarterResolution = "q"
	halfResolution    = "h"
	fullResolution    = "f"
)

type SimulcastConfig struct {
	BestQualityFirst    bool `mapstructure:"bestqualityfirst"`
	EnableTemporalLayer bool `mapstructure:"enabletemporallayer"`
}

type simulcastTrackHelpers struct {
	targetSpatialLayer int
	temporalSupported  bool
	targetTempLayer    int
	currentTempLayer   int
	temporalEnabled    bool
	lTSCalc            time.Time

	// VP8Helper temporal helpers
	pRefPicID  uint16
	refPicID   uint16
	lPicID     uint16
	pRefTlZIdx uint8
	refTlZIdx  uint8
	lTlZIdx    uint8
	refSN      uint16
}
