package taraxacum

import (
	"errors"
	"github.com/FlowerLab/taraxacum/hash"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"os"
	"path"
)

type UploadInfo struct {
	No       int    `json:"no"`
	Filename string `json:"filename"`
	OK       bool   `json:"ok"`
	Path     string `json:"path"`
}

func uploadAction(c *gin.Context) {
	r := NewResult()
	defer c.JSON(http.StatusOK, r)

	form, err := c.MultipartForm()
	if err != nil {
		r.Code = -1
		r.Msg = err.Error()
		return
	}

	info := make([]*UploadInfo, len(form.File["file"]))

	for i, v := range form.File["file"] {
		info[i] = &UploadInfo{
			No:       i,
			Filename: v.Filename,
		}
		if p, err := saveFile(v); err == nil {
			info[i].OK = true
			info[i].Path = Conf.Domain + Conf.AccessPath + p
		}
	}

	r.Data = info
}

func saveFile(f *multipart.FileHeader) (string, error) {
	if f == nil {
		return "", errors.New("file is nil")
	}
	extname, ok := checkExt(f.Filename)
	if !ok {
		return "", errors.New("extname no support")
	}

	file, err := f.Open()
	if err != nil {
		return "", err
	}
	defer file.Close()

	bt, err := ioutil.ReadAll(file)
	if err != nil {
		return "", err
	}

	h := hash.New64()
	if _, err := h.Write(bt); err != nil {
		return "", err
	}
	filename := hash.Uint64Base62.Encoding(h.Sum64()) + extname
	// todo check file exists

	distPath := path.Join(Conf.FilePath, filename)
	dist, err := os.Create(distPath)
	if err != nil {
		return "", err
	}
	defer dist.Close()

	_, err = dist.Write(bt)
	return filename, err
}
