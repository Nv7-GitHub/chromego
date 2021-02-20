package chromego

import goselenium "github.com/bunsenapp/go-selenium"

// Navigate navigates the driver to a url
func (d *Driver) Navigate(url string) error {
	_, err := d.driver.Go(url)
	return err
}

// Element has all the element functions
type Element struct {
	el goselenium.Element
}

// FindByCSSSelector finds an element by its css selector
func (d *Driver) FindByCSSSelector(selector string) (Element, error) {
	el, err := d.driver.FindElement(goselenium.ByCSSSelector(selector))
	if err != nil {
		return Element{}, err
	}
	return Element{
		el: el,
	}, nil
}

// FindByIndex finds an element by its index on the page's elements
func (d *Driver) FindByIndex(index uint) (Element, error) {
	el, err := d.driver.FindElement(goselenium.ByIndex(index))
	if err != nil {
		return Element{}, err
	}
	return Element{
		el: el,
	}, nil
}

// FindByLinkText finds a link by link text
func (d *Driver) FindByLinkText(text string) (Element, error) {
	el, err := d.driver.FindElement(goselenium.ByLinkText(text))
	if err != nil {
		return Element{}, err
	}
	return Element{
		el: el,
	}, nil
}

// FindByPartialLinkText finds a link by part of its text
func (d *Driver) FindByPartialLinkText(partialText string) (Element, error) {
	el, err := d.driver.FindElement(goselenium.ByPartialLinkText(partialText))
	if err != nil {
		return Element{}, err
	}
	return Element{
		el: el,
	}, nil
}

// FindByXPath finds an element by its XPath
func (d *Driver) FindByXPath(xpath string) (Element, error) {
	el, err := d.driver.FindElement(goselenium.ByXPath(xpath))
	if err != nil {
		return Element{}, err
	}
	return Element{
		el: el,
	}, nil
}

// FindElemsByCSSSelector finds elements by a css selector
func (d *Driver) FindElemsByCSSSelector(selector string) ([]Element, error) {
	els, err := d.driver.FindElements(goselenium.ByCSSSelector(selector))
	if err != nil {
		return []Element{}, err
	}
	elems := make([]Element, len(els))
	for i, val := range els {
		elems[i] = Element{
			el: val,
		}
	}
	return elems, nil
}

// FindElemsByLinkText finds links by their text
func (d *Driver) FindElemsByLinkText(text string) ([]Element, error) {
	els, err := d.driver.FindElements(goselenium.ByLinkText(text))
	if err != nil {
		return []Element{}, err
	}
	elems := make([]Element, len(els))
	for i, val := range els {
		elems[i] = Element{
			el: val,
		}
	}
	return elems, nil
}

// FindElemsByPartialLinkText finds a links by part of their text
func (d *Driver) FindElemsByPartialLinkText(partialText string) ([]Element, error) {
	els, err := d.driver.FindElements(goselenium.ByPartialLinkText(partialText))
	if err != nil {
		return []Element{}, err
	}
	elems := make([]Element, len(els))
	for i, val := range els {
		elems[i] = Element{
			el: val,
		}
	}
	return elems, nil
}

// FindElemsByXPath finds elements by their XPath
func (d *Driver) FindElemsByXPath(xpath string) ([]Element, error) {
	els, err := d.driver.FindElements(goselenium.ByXPath(xpath))
	if err != nil {
		return []Element{}, err
	}
	elems := make([]Element, len(els))
	for i, val := range els {
		elems[i] = Element{
			el: val,
		}
	}
	return elems, nil
}
