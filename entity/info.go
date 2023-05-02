package entity

type Info struct {
	Url       string `json:"url"`
	LocalPath string `json:"local_path"`
	CommonKey string `json:"common_key"`
	Ext       string `json:"ext"`
	MimeType  string `json:"mime-type"`
	//Size indicates size of file in bytes
	Size uint64 `json:"size"`
	//DownloadedAt is time stamp format utc
	DownloadedAt uint64 `json:"downloaded_at"`
	//Mode represents a type of downloading origin : images.txt file or user-request
	Mode mode `json:"mode"`
	//Meta represents all optional information in order to forward compatibility
	Meta map[string]any `json:"meta"`
}

type mode string

const (
	DownloadedFromImageListFileMode = mode("from-list-file")
	DownloadedByUserRequestMode     = mode("by-user-request")
	UnDefinedMode                   = mode("un-defined")
)

func ResolveMode(s string) mode {
	switch s {
	case "from-list-file":
		return DownloadedFromImageListFileMode
	case "by-user-request":
		return DownloadedByUserRequestMode
	default:
		return UnDefinedMode
	}
}
