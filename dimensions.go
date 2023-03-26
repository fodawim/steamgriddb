package sgdb

type Dimensions []dimension

type dimension string

const (
	// Dimension filters for Grids
	GridDimension460x215   dimension = "460x215"
	GridDimension920x430   dimension = "920x430"
	GridDimension600x900   dimension = "600x900"
	GridDimension342x482   dimension = "342x482"
	GridDimension660x930   dimension = "660x930"
	GridDimension512x512   dimension = "512x512"
	GridDimension1024x1024 dimension = "1024x1024"

	// Dimension filters for Heroes
	HeroDimension1920x620  dimension = "1920x620"
	HeroDimension3840x1240 dimension = "3840x1240"
	HeroDimension1600x650  dimension = "1600x650"

	// Dimension filters for Logos
	
)
