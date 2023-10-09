package pdf_gen

import (
	"github.com/HMUniversity/System/modules/config"
	"github.com/HMUniversity/System/modules/err_handler"
)

type CertificationData struct {
	Programme string
	Date      string
	Receiver  string
	Degree    string
}

func NewCertification(data CertificationData, outpath string) error {
	r := newRequestPdf("")
	templatePath := config.Get().CertTemplate
	if err := r.ParseTemplate(templatePath, data); err != nil {
		err_handler.HandleError(err)
		return err
	}
	ok, err := r.GeneratePDF(outpath)
	if ok {
		return nil
	}
	return err
}
