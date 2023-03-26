package sgdb

type Styles []style

type style string

const (
	StyleOfficial style = "official"
	StyleCustom   style = "custom"

	StyleAlternate style = "alternate"
	StyleBlurred   style = "blurred"
	StyleWhiteLogo style = "white_logo"
	StyleMaterial  style = "material"
	StyleNoLogo    style = "no_logo"
)
