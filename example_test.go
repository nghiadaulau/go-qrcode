// go-qrcode
// Copyright 2014 Tom Harwood
/*
	Amendments Thu, 2017-December-14:
	- test integration (go test -v)
	- idiomatic go code
*/
package qrcode

import (
	"fmt"
	"image/color"
	"log"
	"os"
	"testing"
)

func TestExampleEncode(t *testing.T) {
	if png, err := Encode("https://example.org", Medium, 256); err != nil {
		t.Errorf("Error: %s", err.Error())
	} else {
		fmt.Printf("PNG is %d bytes long", len(png))
	}
}

func TestExampleWriteFile(t *testing.T) {
	filename := "example.png"
	if err := WriteFile("https://example.org", Medium, 256, filename); err != nil {
		if err = os.Remove(filename); err != nil {
			t.Errorf("Error: %s", err.Error())
		}
	}
}

func TestExampleEncodeWithColourAndWithoutBorder(t *testing.T) {
	q, err := New("https://example.org", Medium)
	if err != nil {
		t.Errorf("Error: %s", err)
		return
	}

	// Optionally, disable the QR Code border.
	q.DisableBorder = true

	// Optionally, set the colours.
	q.ForegroundColor = color.RGBA{R: 0x33, G: 0x33, B: 0x66, A: 0xff}
	q.BackgroundColor = color.RGBA{R: 0xef, G: 0xef, B: 0xef, A: 0xff}

	err = q.WriteFile(256, "example2.png")
	if err != nil {
		t.Errorf("Error: %s", err)
		return
	}
}

func ExampleQRCode_QuietZoneSize() {
	// Create QR code with default quiet zone size (4)
	q1, err := New("https://example.org", Medium)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Default quiet zone size: %d\n", q1.QuietZoneSize)

	// Create QR code with custom quiet zone size
	q2, err := NewWithQuietZone("https://example.org", Medium, 8)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Custom quiet zone size: %d\n", q2.QuietZoneSize)

	// Create QR code with no quiet zone
	q3, err := NewWithQuietZone("https://example.org", Medium, 0)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("No quiet zone: %d\n", q3.QuietZoneSize)

	// Output:
	// Default quiet zone size: 4
	// Custom quiet zone size: 8
	// No quiet zone: 0
}

func TestExampleQuietZoneSizeComparison(t *testing.T) {
	content := "https://example.org"

	// Create QR codes with different quiet zone sizes
	quietZoneSizes := []int{0, 1, 2, 4, 8}

	for _, qzSize := range quietZoneSizes {
		q, err := NewWithQuietZone(content, Medium, qzSize)
		if err != nil {
			t.Errorf("Error creating QR code with quiet zone size %d: %s", qzSize, err)
			continue
		}

		filename := fmt.Sprintf("example_quietzone_%d.png", qzSize)
		err = q.WriteFile(256, filename)
		if err != nil {
			t.Errorf("Error writing file %s: %s", filename, err)
			continue
		}

		// Clean up test files

		fmt.Printf("Created QR code with quiet zone size %d: %s\n", qzSize, filename)
	}
}
