package chromego

import (
	"bytes"
	"encoding/base64"
	"image"
	"image/png"
)

// https://pkg.go.dev/github.com/bunsenapp/go-selenium#WebDriver - More methods need to be implemented

// URL gets the driver's current URL
func (d *Driver) URL() string {
	return d.driver.DriverURL()
}

// Refresh refreshes the driver
func (d *Driver) Refresh() error {
	_, err := d.driver.Refresh()
	return err
}

// ScreenshotBase64 gets a base64 screenshot of the window
func (d *Driver) ScreenshotBase64() (string, error) {
	resp, err := d.driver.Screenshot()
	return resp.EncodedImage, err
}

// Screenshot gets a screenshot of the window
func (d *Driver) Screenshot() (image.Image, error) {
	data, err := d.ScreenshotBase64()
	if err != nil {
		return nil, err
	}
	unbased, err := base64.StdEncoding.DecodeString(data)
	if err != nil {
		panic("Cannot decode b64")
	}
	r := bytes.NewReader(unbased)
	im, err := png.Decode(r)
	return im, err
}
