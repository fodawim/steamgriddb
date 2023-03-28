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

// Dimensions for Icons
type DimensionsIcon []dimensionIcon

type dimensionIcon string

const (
	DimensionsIcon8    dimensionIcon = "8"
	DimensionsIcon10   dimensionIcon = "10"
	DimensionsIcon14   dimensionIcon = "14"
	DimensionsIcon16   dimensionIcon = "16"
	DimensionsIcon20   dimensionIcon = "20"
	DimensionsIcon24   dimensionIcon = "24"
	DimensionsIcon28   dimensionIcon = "28"
	DimensionsIcon32   dimensionIcon = "32"
	DimensionsIcon35   dimensionIcon = "35"
	DimensionsIcon40   dimensionIcon = "40"
	DimensionsIcon48   dimensionIcon = "48"
	DimensionsIcon54   dimensionIcon = "54"
	DimensionsIcon56   dimensionIcon = "56"
	DimensionsIcon57   dimensionIcon = "57"
	DimensionsIcon60   dimensionIcon = "60"
	DimensionsIcon64   dimensionIcon = "64"
	DimensionsIcon72   dimensionIcon = "72"
	DimensionsIcon76   dimensionIcon = "76"
	DimensionsIcon80   dimensionIcon = "80"
	DimensionsIcon90   dimensionIcon = "90"
	DimensionsIcon96   dimensionIcon = "96"
	DimensionsIcon100  dimensionIcon = "100"
	DimensionsIcon114  dimensionIcon = "114"
	DimensionsIcon120  dimensionIcon = "120"
	DimensionsIcon128  dimensionIcon = "128"
	DimensionsIcon144  dimensionIcon = "144"
	DimensionsIcon150  dimensionIcon = "150"
	DimensionsIcon152  dimensionIcon = "152"
	DimensionsIcon160  dimensionIcon = "160"
	DimensionsIcon180  dimensionIcon = "180"
	DimensionsIcon192  dimensionIcon = "192"
	DimensionsIcon194  dimensionIcon = "194"
	DimensionsIcon256  dimensionIcon = "256"
	DimensionsIcon310  dimensionIcon = "310"
	DimensionsIcon512  dimensionIcon = "512"
	DimensionsIcon768  dimensionIcon = "768"
	DimensionsIcon1024 dimensionIcon = "1024"
)
