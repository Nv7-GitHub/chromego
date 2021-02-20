package chromego

// Clear clears an element
func (e *Element) Clear() error {
	_, err := e.el.Clear()
	return err
}

// Click clicks an element
func (e *Element) Click() error {
	_, err := e.el.Click()
	return err
}

// Type sends keys to an element
func (e *Element) Type(text string) error {
	_, err := e.el.SendKeys(text)
	return err
}

// Selected gets whether an element is selected
func (e *Element) Selected() (bool, error) {
	resp, err := e.el.Selected()
	if err != nil {
		return false, err
	}
	return resp.Selected, nil
}

// Enabled gets whether an element is enabled
func (e *Element) Enabled() (bool, error) {
	resp, err := e.el.Enabled()
	if err != nil {
		return false, err
	}
	return resp.Enabled, nil
}

// ID gets an element's id
func (e *Element) ID() string {
	return e.el.ID()
}

// Attr gets an element's attribute
func (e *Element) Attr(attribute string) (string, error) {
	resp, err := e.el.Attribute(attribute)
	if err != nil {
		return "", err
	}
	return resp.Value, nil
}

// Tag gets an element's tag (p, div, etc.)
func (e *Element) Tag() (string, error) {
	resp, err := e.el.TagName()
	if err != nil {
		return "", err
	}
	return resp.Tag, nil
}

// CSSAttr gets an element's css attribute
func (e *Element) CSSAttr(attribute string) (string, error) {
	resp, err := e.el.CSSValue(attribute)
	if err != nil {
		return "", err
	}
	return resp.Value, nil
}

// Rect gets an element's rectangle
func (e *Element) Rect() (Rect, error) {
	resp, err := e.el.Rectangle()
	if err != nil {
		return Rect{}, err
	}
	return Rect{
		X:      resp.Rectangle.X,
		Y:      resp.Rectangle.Y,
		Width:  int(resp.Rectangle.Width),
		Height: int(resp.Rectangle.Height),
	}, nil
}

// Rect contains a rectangle's dimensions and position
type Rect struct {
	X      int
	Y      int
	Width  int
	Height int
}
