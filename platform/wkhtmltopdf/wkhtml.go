package wkhtmltopdf

import (
	"attempt4/platform/zap"
	"fmt"
	"github.com/SebastiaanKlippert/go-wkhtmltopdf"
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
	err1 := ioutil.WriteFile(dirPath+strconv.FormatInt(int64(t), 10)+".html", []byte(r.Body), 0644)
	if err1 != nil {
		return nil, err1
	}

	f, err := os.Open(dirPath + strconv.FormatInt(int64(t), 10) + ".html")
	if f != nil {
		defer f.Close()
	}
	if err != nil {
		return nil, err
	}

	zap.Logger.Info(wkhtmltopdf.GetPath())

	wkhtmltopdf.SetPath("C:/Program Files/wkhtmltopdf/bin/wkhtmltopdf.exe")
	pdfg, err := wkhtmltopdf.NewPDFGenerator()
	if err != nil {
		return nil, err
	}

	pdfg.AddPage(wkhtmltopdf.NewPageReader(strings.NewReader(r.Body)))

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
