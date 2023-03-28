package sgdb

type Mimes []mime

type mime string

const (
	MimeJPEG mime = "image/jpeg"
	MimePNG  mime = "image/png"
	MimeWebP mime = "image/webp"
)

type MimesHero []mimehero

type mimehero string

const (
	MimeHeroJPEG mimehero = "image/jpeg"
	MimeHeroPNG  mimehero = "image/png"
	MimeHeroWebP mimehero = "image/webp"
)

type MimesLogo []mimelogo

type mimelogo string

const (
	MimeLogoPNG  mimelogo = "image/png"
	MimeLogoWebP mimelogo = "image/webp"
)

type MimesIcon []mimeicon

type mimeicon string

const (
	MimeIconPNG mimeicon = "image/png"
	MimeIconVND mimeicon = "image/vnd.microsoft.icon"
)
