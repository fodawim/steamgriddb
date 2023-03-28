package sgdb

// Dimension  for Grids
type Dimensions []dimension

type dimension string

const (
	Dimensions460x215   dimension = "460x215"
	Dimensions920x430   dimension = "920x430"
	Dimensions600x900   dimension = "600x900"
	Dimensions342x482   dimension = "342x482"
	Dimensions660x930   dimension = "660x930"
	Dimensions512x512   dimension = "512x512"
	Dimensions1024x1024 dimension = "1024x1024"
)

// Dimension filters for Heroes
type DimensionsHero []dimensionhero

type dimensionhero string

const (
	DimensionsHero1920x620  dimension = "1920x620"
	DimensionsHero3840x1240 dimension = "3840x1240"
	DimensionsHero1600x650  dimension = "1600x650"
)

)
