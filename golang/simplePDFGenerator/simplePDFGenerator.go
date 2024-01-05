package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println("# 1764Snippets")
	fmt.Println("## Golang PDF")

	content := "1764Snippets PDF"
	err := generatePDF("simplePDFGenerator.pdf", content)
	if err != nil {
		fmt.Println("Failed to create PDF:", err)
	}
}

func generatePDF(filename, content string) error {
	f, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer f.Close()

	// A4 size in points is 595 x 842
	// The text will be positioned approximately at the top center 75 750
	pdfContent := "%PDF-1.4\n1 0 obj\n<< /Type /Catalog /Pages 2 0 R >>\nendobj\n" +
		"2 0 obj\n<< /Type /Pages /Kids [3 0 R] /Count 1 >>\nendobj\n" +
		"3 0 obj\n<< /Type /Page /Parent 2 0 R /MediaBox [0 0 595 842] /Contents 4 0 R >>\nendobj\n" +
		"4 0 obj\n<< /Length 77 >>\nstream\n" +
		"BT /F1 24 Tf 75 750 Td (" + content + ") Tj ET\n" +
		"endstream\nendobj\n" +
		"xref\n0 5\n0000000000 65535 f \n0000000010 00000 n \n0000000077 00000 n \n0000000178 00000 n \n0000000467 00000 n \n" +
		"trailer\n<< /Size 5 /Root 1 0 R >>\nstartxref\n627\n%%EOF"

	_, err = f.WriteString(pdfContent)
	if err != nil {
		return err
	}

	return nil
}
