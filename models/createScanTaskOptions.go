package models

type CreateScanTaskOptions struct {
	FileUri     string  `json:"file_uri"`
	FileName    string  `json:"file_name"`
	AsyncResult bool    `json:"async_result"`
	Options     Options `json:"options"`
}

type Options struct {
	AnalysisDepth      int      `json:"analysis_depth"`
	PasswordsForUnpack []string `json:"passwords_for_unpack"`
	Sandbox            Sandbox  `json:"sandbox"`
}

type Sandbox struct {
	Enabled           bool   `json:"enabled"`
	SkipCheckMimeType bool   `json:"skip_check_mime_type"`
	ImageId           string `json:"image_id"`
	AnalysisDuration  int    `json:"analysis_duration"`
}
