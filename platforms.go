package sgdb

type Platforms []platform

type platform string

const (
	// PlatformSteam is the Steam platform.
	PlatformSteam platform = "steam"

	// PlatformOrigin is the Origin platform.
	PlatformOrigin platform = "origin"

	// PlatformEGS is the Epic Games Store platform.
	PlatformEGS platform = "egs"

	// PlatformBattleNet is the Battle.net platform.
	PlatformBattleNet platform = "bnet"

	// PlatformUplay is the Uplay platform.
	PlatformUplay platform = "uplay"

	// PlatformFlashpoint is the Flashpoint platform.
	PlatformFlashpoint platform = "flashpoint"

	// PlatformEshop is the Nintendo eShop platform.
	PlatformEshop = "eshop"
)
