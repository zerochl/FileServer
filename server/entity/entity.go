package entity

type FileServerConf struct {
	Port string `json:"port"`
	Route string `json:"route"`
	FilePath string `json:"file_path"`
}

type File struct {
	FileName string
	FilePath string
	FileMD5 string
	FileSize int64
}