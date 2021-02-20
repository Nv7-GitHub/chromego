package chromego

import (
	"os/exec"
	"path/filepath"
	"strconv"

	goselenium "github.com/bunsenapp/go-selenium"
)

// Driver contains all the data and methods for a chrome driver
type Driver struct {
	driver goselenium.WebDriver
}

// Close closes the session
func (d *Driver) Close() {
	d.driver.DeleteSession()
}

// CreateDriver starts a driver and sets up the session
func CreateDriver(chromedriverpath string, port int) (Driver, error) {
	prt := strconv.Itoa(port)

	// Start driver
	cmd := &exec.Cmd{
		Path: filepath.Dir(chromedriverpath),
		Args: []string{chromedriverpath, "--port=" + prt, "&"},
	}
	if err := cmd.Run(); err != nil {
		return Driver{}, err
	}

	capabilities := goselenium.Capabilities{}
	capabilities.SetBrowser(goselenium.ChromeBrowser())
	driver, err := goselenium.NewSeleniumWebDriver("http://localhost:"+prt, capabilities)
	if err != nil {
		return Driver{}, err
	}
	_, err = driver.CreateSession()
	if err != nil {
		return Driver{}, err
	}

	drv := Driver{
		driver: driver,
	}
	return drv, nil
}
