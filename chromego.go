package chromego

import (
	"os/exec"
	"strconv"

	goselenium "github.com/bunsenapp/go-selenium"
)

// Driver contains all the data and methods for a chrome driver
type Driver struct {
	driver goselenium.WebDriver
	cmd    exec.Cmd
}

// Close closes the session
func (d *Driver) Close() error {
	_, err := d.driver.DeleteSession()
	if err != nil {
		return err
	}
	err = d.cmd.Process.Kill()
	return err
}

// CreateDriver starts a driver and sets up the session
func CreateDriver(chromedriverpath string, port int) (Driver, error) {
	prt := strconv.Itoa(port)

	// Start driver
	cmd := exec.Command(chromedriverpath, "--port="+prt)
	if err := cmd.Start(); err != nil {
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
