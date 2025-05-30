package app

import (
	"time"

	epochsmoduletypes "github.com/evmos/evmos/v16/x/epochs/types"
)

const (
	ForkAddEpochsMainNetHeight = int64(6397023)
	ForkAddEpochsTestNetHeight = int64(12517742)

	ForkGasLimitMainNetHeight = int64(3760837)
	ForkGasLimitTestNetHeight = int64(10827096)
)

var (
	ForkTestNetEpochInfo = []epochsmoduletypes.EpochInfo{
		{
			Identifier:              epochsmoduletypes.WeekEpochID,
			StartTime:               time.Date(2025, 3, 6, 0, 0, 0, 0, time.UTC),
			Duration:                time.Hour * 24 * 7,
			CurrentEpoch:            0,
			CurrentEpochStartHeight: 0,
			CurrentEpochStartTime:   time.Time{},
			EpochCountingStarted:    false,
		},
		{
			Identifier:              epochsmoduletypes.DayEpochID,
			StartTime:               time.Date(2025, 3, 6, 0, 0, 0, 0, time.UTC),
			Duration:                time.Hour * 24,
			CurrentEpoch:            0,
			CurrentEpochStartHeight: 0,
			CurrentEpochStartTime:   time.Time{},
			EpochCountingStarted:    false,
		},
		{
			Identifier:              epochsmoduletypes.HourEpochID,
			StartTime:               time.Date(2025, 3, 6, 0, 0, 0, 0, time.UTC),
			Duration:                time.Hour,
			CurrentEpoch:            0,
			CurrentEpochStartHeight: 0,
			CurrentEpochStartTime:   time.Time{},
			EpochCountingStarted:    false,
		},
	}

	ForkMainNetEpochInfo = []epochsmoduletypes.EpochInfo{
		{
			Identifier:              epochsmoduletypes.WeekEpochID,
			StartTime:               time.Date(2025, 3, 6, 0, 0, 0, 0, time.UTC),
			Duration:                time.Hour * 24 * 7,
			CurrentEpoch:            0,
			CurrentEpochStartHeight: 0,
			CurrentEpochStartTime:   time.Time{},
			EpochCountingStarted:    false,
		},
		{
			Identifier:              epochsmoduletypes.DayEpochID,
			StartTime:               time.Date(2025, 3, 6, 0, 0, 0, 0, time.UTC),
			Duration:                time.Hour * 24,
			CurrentEpoch:            0,
			CurrentEpochStartHeight: 0,
			CurrentEpochStartTime:   time.Time{},
			EpochCountingStarted:    false,
		},
		{
			Identifier:              epochsmoduletypes.HourEpochID,
			StartTime:               time.Date(2025, 3, 6, 0, 0, 0, 0, time.UTC),
			Duration:                time.Hour,
			CurrentEpoch:            0,
			CurrentEpochStartHeight: 0,
			CurrentEpochStartTime:   time.Time{},
			EpochCountingStarted:    false,
		},
	}
)
