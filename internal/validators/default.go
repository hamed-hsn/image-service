package validators

var (
	DefaultLinkValidator   = linkValidator{minLength: 8}
	DefaultFilterValidator = filterValidator{}
	DefaultUniqueValidator = uniqueValidator{}
	DefaultUploadValidator = uploadValidator{}
)
