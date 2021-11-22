package license

import (
	"fmt"
	"os"
	"text/template"

	kbptemplate "github.com/Drumato/kubectl-plugin-builder/internal/kpbtemplate"
)

const (
	MITLicense = "MIT"
)

type MITBuilder struct {
	filePath string
	tmpl     *template.Template
	data     *MITLicenseData
}

type MITLicenseData struct {
	Year   uint
	Author string
}

func NewMITLicenseData(year uint, author string) *MITLicenseData {
	return &MITLicenseData{
		Year:   year,
		Author: author,
	}
}

var _ LicenseBuilder = &MITBuilder{}

func NewMITBuilder(filePath string, data *MITLicenseData) *MITBuilder {
	return &MITBuilder{
		filePath: filePath,
		data:     data,
	}
}

func (mb *MITBuilder) Build() error {
	tmpl, err := template.ParseFS(kbptemplate.GlobalTemplates, "templates/licenses/mit.license")
	if err != nil {
		return err
	}

	mb.tmpl = tmpl
	return nil
}

func (mb *MITBuilder) Execute() error {
	f, err := os.Create(mb.filePath)
	if err != nil {
		return err
	}

	return mb.tmpl.Execute(f, mb.data)
}

func (mb *MITBuilder) SourceFileHeader() string {
	const licenceFormat = `
/* MIT License
 * 
 * Copyright (c) %d %s
 * 
 * Permission is hereby granted, free of charge, to any person obtaining a copy
 * of this software and associated documentation files (the "Software"), to deal
 * in the Software without restriction, including without limitation the rights
 * to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
 * copies of the Software, and to permit persons to whom the Software is
 * furnished to do so, subject to the following conditions:
 * 
 * The above copyright notice and this permission notice shall be included in all
 * copies or substantial portions of the Software.
 * 
 * THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
 * IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
 * FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
 * AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
 * LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
 * OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
 * SOFTWARE.
*/`

	return fmt.Sprintf(licenceFormat, mb.data.Year, mb.data.Author)
}
