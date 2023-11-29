package utils

import (
	"bytes"
	"encoding/base64"
	"fmt"
	"github.com/NicholasLiem/IF4031_M1_Ticket_App/internal/dto"
	"io/ioutil"
	"os"

	"github.com/go-pdf/fpdf"
)

func GeneratePDFWithQRCode(pdfData dto.EventPDF) (*bytes.Buffer, error) {
	base64QRCode, err := GenerateQRCode(pdfData.QRContent)
	if err != nil {
		return nil, fmt.Errorf("error generating QR Code: %w", err)
	}

	qrCodeBytes, err := base64.StdEncoding.DecodeString(base64QRCode)
	if err != nil {
		return nil, fmt.Errorf("error decoding base64 QR Code: %w", err)
	}

	tmpFile, err := ioutil.TempFile("", "qr-*.png")
	if err != nil {
		return nil, fmt.Errorf("error creating temporary file: %w", err)
	}
	defer os.Remove(tmpFile.Name())

	_, err = tmpFile.Write(qrCodeBytes)
	if err != nil {
		return nil, fmt.Errorf("error writing to temporary file: %w", err)
	}
	err = tmpFile.Close()
	if err != nil {
		return nil, fmt.Errorf("error closing temporary file: %w", err)
	}

	pdf := fpdf.New("P", "mm", "A4", "")
	pdf.AddPage()

	// Set Font
	pdf.SetFont("Arial", "B", 16)

	// Title - Event Name
	pdf.CellFormat(0, 10, pdfData.EventName, "0", 1, "C", false, 0, "")

	// User Details
	pdf.SetFont("Arial", "", 12)
	pdf.Ln(10)
	pdf.CellFormat(0, 10, fmt.Sprintf("Event Date: %s", pdfData.EventDate), "0", 1, "C", false, 0, "")
	pdf.CellFormat(0, 10, fmt.Sprintf("Email: %s", pdfData.UserEmail), "0", 1, "C", false, 0, "")
	pdf.CellFormat(0, 10, fmt.Sprintf("Seat ID: %s", pdfData.SeatID), "0", 1, "C", false, 0, "")

	// QR Code
	qrCodeSize := 100.0
	x := (210 - qrCodeSize) / 2
	y := pdf.GetY() + 10
	pdf.Image(tmpFile.Name(), x, y, qrCodeSize, qrCodeSize, false, "", 0, "")

	pdf.SetY(y + qrCodeSize + 10)
	pdf.CellFormat(0, 10, "Scan QR", "0", 1, "C", false, 0, "")

	var buf bytes.Buffer
	err = pdf.Output(&buf)
	if err != nil {
		return nil, fmt.Errorf("error generating PDF: %w", err)
	}

	return &buf, nil
}
