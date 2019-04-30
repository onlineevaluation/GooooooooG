package model

type CodeInfo struct {
	FilePath   string `form:"filepath"`
	OutPath    string `form:"outpath"`
	FileName   string `form:"filename"`
	TestSet    string `form:"testSet"`
	LimiteTime int64  `form:"limittime"`
}
