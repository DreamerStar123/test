package main

import (
	"bytes"
	"html/template"
	"net/http"
	"os"

	"github.com/SebastiaanKlippert/go-wkhtmltopdf"
)

type ExportTrans struct {
	Date        string
	Time        string
	Description string
	Reference   string
	Sign        bool
	Amount      float64
	Balance     float64
}

type PageData struct {
	Currency       string
	StartDate      string
	EndDate        string
	OpeningBalance float64
	ClosingBalance float64
	CreatedAt      string
	MoneyIn        float64
	MoneyOut       float64
	Trans          []ExportTrans
}

func handler(w http.ResponseWriter, r *http.Request) {
	tmplBytes, err := os.ReadFile("C:\\Users\\Administrator\\Documents\\temp.html")
	if err != nil {
		http.Error(w, "Failed to read html", http.StatusInternalServerError)
		return
	}
	tmpl := string(tmplBytes)

	t := template.Must(template.New("hello").Parse(tmpl))
	// Render HTML to a buffer
	var buf bytes.Buffer
	t.Execute(&buf, PageData{
		Currency:       "USD",
		StartDate:      "1st May 2025",
		EndDate:        "1st Jun 2025",
		OpeningBalance: 1524.35,
		ClosingBalance: 2543.21,
		CreatedAt:      "12nd Mar 2025",
		MoneyIn:        1254.36,
		MoneyOut:       2543.21,
		Trans: []ExportTrans{
			{
				Date:        "3rd Feb 2025",
				Time:        "20:12:20",
				Description: "Deposit",
				Reference:   "d80f2351",
				Sign:        false,
				Amount:      16.24,
				Balance:     3954.53,
			},
		},
	})

	// Use wkhtmltopdf to convert HTML to PDF
	// Requires github.com/SebastiaanKlippert/go-wkhtmltopdf and wkhtmltopdf installed on your system

	pdf, err := wkhtmltopdf.NewPDFGenerator()
	if err != nil {
		http.Error(w, "Failed to initialize PDF generator", http.StatusInternalServerError)
		return
	}

	pdf.MarginLeft.Set(0)
	pdf.MarginRight.Set(0)
	pdf.MarginTop.Set(0)
	pdf.MarginBottom.Set(0)

	// page := wkhtmltopdf.NewPage("https://godoc.org/github.com/SebastiaanKlippert/go-wkhtmltopdf")
	page := wkhtmltopdf.NewPageReader(bytes.NewReader(buf.Bytes()))
	page.FooterCenter.Set("page number")
	page.FooterFontSize.Set(10)
	page.FooterSpacing.Set(5)
	pdf.AddPage(page)

	err = pdf.Create()
	if err != nil {
		http.Error(w, "Failed to generate PDF", http.StatusInternalServerError)
		return
	}

	// Set headers for PDF download
	w.Header().Set("Content-Type", "application/pdf")
	w.Header().Set("Content-Disposition", "attachment; filename=\"output.pdf\"")
	w.Write(pdf.Bytes())
}

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}
