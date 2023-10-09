package pdf_gen

import (
	"bytes"
	"github.com/SebastiaanKlippert/go-wkhtmltopdf"
	"html/template"
	"strings"
)

type pdfGen struct {
	body string
}

func newRequestPdf(body string) *pdfGen {
	return &pdfGen{
		body: body,
	}
}

func (r *pdfGen) ParseTemplate(templateFileName string, data interface{}) error {
	t, err := template.ParseFiles(templateFileName)
	if err != nil {
		return err
	}
	buf := new(bytes.Buffer)
	if err = t.Execute(buf, data); err != nil {
		return err
	}
	r.body = buf.String()
	return nil
}

func (r *pdfGen) GeneratePDF(pdfPath string) (bool, error) {
	pdfg, err := wkhtmltopdf.NewPDFGenerator()
	if err != nil {
		return false, err
	}
	pdfg.NoPdfCompression.Set(true)
	pdfg.AddPage(wkhtmltopdf.NewPageReader(strings.NewReader(r.body)))
	pdfg.PageSize.Set(wkhtmltopdf.PageSizeA4)
	pdfg.Dpi.Set(300)

	if err = pdfg.Create(); err != nil {
		return false, err
	}
	if err = pdfg.WriteFile(pdfPath); err != nil {
		return false, err
	}
	return true, nil
}
