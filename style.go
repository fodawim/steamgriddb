package sgdb

type Styles []style

type style string

const (
	StyleAlternate style = "alternate"
	StyleBlurred   style = "blurred"
	StyleWhiteLogo style = "white_logo"
	StyleMaterial  style = "material"
	StyleNoLogo    style = "no_logo"
)
