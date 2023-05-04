package wkhtmltopdf

import (
	"fmt"
	"github.com/SebastiaanKlippert/go-wkhtmltopdf"
	"github.com/spf13/viper"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
	"time"
)

type RequestPdf struct {
	Body string
}

func (r *RequestPdf) GeneratePDF(pdfPath string) ([]byte, error) {
	t := time.Now().Unix()
	dir, err := os.Getwd()
	if err != nil {
		return nil, err
	}
	dirPath := fmt.Sprintf("%s/internal/application/template/cloneTemplate/", dir)

	if _, err := os.Stat(dirPath); os.IsNotExist(err) {
		errDir := os.Mkdir(dirPath, 0777)
		if errDir != nil {
			return nil, errDir
		}
	}
	clonePath := dirPath + strconv.FormatInt(int64(t), 10) + ".html"
	err1 := ioutil.WriteFile(clonePath, []byte(r.Body), 0644)
	if err1 != nil {
		return nil, err1
	}

	f, err := os.Open(clonePath)
	if f != nil {
		defer f.Close()
	}
	if err != nil {
		return nil, err
	}

	wkhtmltopdf.SetPath(fmt.Sprintf("%v", viper.Get("WKHTMLTOPDF_PATH")))

	pdfg, err := wkhtmltopdf.NewPDFGenerator()
	if err != nil {
		return nil, err
	}
	page := wkhtmltopdf.NewPageReader(strings.NewReader(r.Body))

	pdfg.AddPage(page)

	pdfg.PageSize.Set(wkhtmltopdf.PageSizeA4)

	pdfg.Dpi.Set(300)

	err = pdfg.Create()
	if err != nil {
		return nil, err
	}

	err = pdfg.WriteFile(pdfPath)
	if err != nil {
		return nil, err
	}

	defer os.RemoveAll(dirPath)
	return pdfg.Bytes(), nil
}
