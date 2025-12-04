package app

import (
	"time"

	epochsmoduletypes "github.com/evmos/evmos/v18/x/epochs/types"
)

const (
	ForkAddEpochsMainNetHeight = int64(6397023)
	ForkAddEpochsTestNetHeight = int64(12517742)

	ForkGasLimitMainNetHeight = int64(3760837)
	ForkGasLimitTestNetHeight = int64(10827096)

	ForkAddressBlockedMainNetHeight         = int64(7245276)
	ForkOmissionAddressBlockedMainNetHeight = int64(7250689)
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

var (
	BlockedAddress = []string{
		"0x13Cb3734D08C79B75f34B6DDb65Ba62B8C1C46C9",
		"0x385C99dEFF3929dC0431B4EC93c9a269b276E6c9",
		"0xD40af843baeFeCa27EE0C67D5331cD3d92848cb3",
		"0x6244b7D8BcEA2F2B05a3395FCADDD33015310275",
		"0x76DE7B4Ed75c4DAB6F1991867440Df4E5c79197b",
		"0x57619a49F0862076C5E7d129Ddf7b7d750C2218B",
		"0xAd22b92f46582dC540dB3e83a91658136bb2cD18",
		"0x338220789C9eEA72e87f028f5Ea1d8C62740c157",
		"0xc3A2575E5D54a169eC832E46dCA0dcdf0562B495",
		"0x06554705BF0F89e425C5ec0691E81d632eA47C5B",
		"0xf964C0fB0867E9E3b58342286d61AC070dea12Be",
		"0xeEAc08a37DD3E3Ccbe8A4fB65643cA741eF2315c",
		"0x422fC6d98BfC720BE3B9c594a324dbdcE79C9967",
	}

	BlockedOmissionAddress = []string{
		"0xbF84701E1eF2405179fdF8D2dc531aa77B895e06",
	}
)
