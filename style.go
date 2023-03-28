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

type StylesHero []herostyle

type herostyle string

const (
	StylesHeroAlternate herostyle = "alternate"
	StylesHeroBlurred   herostyle = "blurred"
	StylesHeroMaterial  herostyle = "material"
)

type StylesLogo []logostyle

type logostyle string

const (
	StylesLogoOfficial logostyle = "official"
	StylesLogoWhite    logostyle = "white"
	StylesLogoBlack    logostyle = "black"
	StylesLogoCustom   logostyle = "custom"
)

type StylesIcon []iconstyle

type iconstyle string

const (
	StylesIconOfficial iconstyle = "official"
	StylesIconCustom   iconstyle = "custom"
)
