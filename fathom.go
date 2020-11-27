package fathom

import (
	"fmt"
	"net/url"

	"github.com/anaskhan96/soup"
	"github.com/go-resty/resty"
)

const URL = "https://app.usefathom.com"

var Headers = map[string]string{
	"User-Agent": "Mozilla/5.0 (X11; Linux x86_64; rv:81.0) Gecko/20100101 Firefox/81.0",
}

func New() (*Config, error) {
	return &Config{}, nil
}

func (c *Config) Login(username, password string, remember bool) error {
	headers := Headers
	headers["Content-Type"] = "application/x-www-form-urlencoded"

	loginUrl := fmt.Sprintf("%s/login", URL)

	usernameFieldName := "email"
	passwordFieldName := "password"

	if len(username) == 0 {
		return fmt.Errorf("username parameter is mandatory")
	}
	if len(password) == 0 {
		return fmt.Errorf("password parameter is mandatory")
	}

	// Get login page html
	client := resty.New()

	resp, err := client.R().
		SetHeaders(headers).
		EnableTrace().
		Get(loginUrl)
	if err != nil {
		return err
	}

	// Get hidden token
        doc := soup.HTMLParse(resp.String())

        hidden := doc.Find("form").Find("input", "type", "hidden")
	hiddenName := hidden.Attrs()["name"]
	hiddenValue := hidden.Attrs()["value"]
	fmt.Println(hiddenValue)

	// Post payload
	formData := make(map[string]string)
	formData[hiddenName] = hiddenValue
	formData[usernameFieldName] = username
	formData[passwordFieldName] = password
	formData["remember"] = "1"

	resp2, err := client.R().
		SetHeaders(Headers).
		SetCookies(resp.Cookies()).
		SetFormData(formData).
		EnableTrace().
		Post(loginUrl)
	if err != nil {
		return err
	}

	c.Cookies = resp2.Cookies()

	return nil
}

func (c *Config) Data(siteID, from, to, timezone string) (*Data, error) {
	headers := map[string]string{
		"User-Agent": "Mozilla/5.0 (X11; Linux x86_64; rv:81.0) Gecko/20100101 Firefox/81.0",
	}

	dataUrl, err := url.Parse(URL)
	if err != nil {
		return nil, err
	}

	dataUrl.Path = fmt.Sprintf("sites/%s/data", siteID)

	params := url.Values{}
	params.Add("from", from)
	params.Add("to", to)
	params.Add("tz", timezone)
	// &site=36032&range=this_month

	dataUrl.RawQuery = params.Encode()

	client := resty.New()

	resp, err := client.R().
		SetHeaders(headers).
		SetCookies(c.Cookies).
		SetResult(&Data{}).
		Get(dataUrl.String())
	if err != nil {
		return nil, err
	}

	return resp.Result().(*Data), nil
}
