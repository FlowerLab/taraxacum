package taraxacum

import (
	"os"
	"path"
	"strings"
)

type config struct {
	Debug           bool
	FilePath        string
	AccessPath      string
	LogPath         string
	Addr            string
	Domain          string
	MultipartMemory int64
}

var Conf *config

func LoadConf() {
	Conf = &config{
		Debug:           true,
		FilePath:        "/data/taraxacum/file",
		AccessPath:      "/static/",
		LogPath:         "/data/taraxacum/logs",
		Addr:            ":13079",
		Domain:          "https://flowerlab.github.io",
		MultipartMemory: 10 << 20,
	}
	if err := os.MkdirAll(Conf.FilePath, os.ModePerm); err != nil {
		panic(err)
	}
}

var fileExt = []string{
	".jpg", ".jpeg", ".png", ".gif", ".bmp", ".svg", ".webp",
	".jpg", ".zip", ".7z", ".mp4",
}

func checkExt(filename string) (extname string, ok bool) {
	extname = strings.ToLower(path.Ext(filename))
	for i := range fileExt {
		if extname == fileExt[i] {
			return extname, true
		}
	}
	return
}
