package chromego

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
