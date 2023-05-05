package validators

var (
	DefaultLinkValidator   = linkValidator{minLength: 8}
	DefaultFilterValidator = filterValidator{}
	DefaultUniqueValidator = uniqueValidator{}
	DefaultUploadValidator = uploadValidator{
		MaxSize: 10 * 1024 * 1024,
		InvalidContentTypes: []string{"", "application/octet-stream", "application/x-msdos-program",
			"application/java-archive", "application/vnd.apple.installer+xml",
			"application/x-httpd-php",
		},
	}
)
