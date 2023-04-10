package profile

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/PuerkitoBio/goquery"
)

type Client struct {
	httpClient *http.Client
}

func NewClient() *Client {
	return &Client{
		httpClient: &http.Client{Timeout: 10 * time.Second},
	}
}

func (c *Client) FetchDataByInn(inn string) (CompanyData, error) {
	req, err := http.NewRequest(http.MethodGet, "https://www.rusprofile.ru/search", nil)
	if err != nil {
		return CompanyData{}, fmt.Errorf("client: could not create request: %w", err)
	}

	q := req.URL.Query()
	q.Add("query", inn)
	req.URL.RawQuery = q.Encode()

	res, err := c.httpClient.Do(req)
	if err != nil {
		return CompanyData{}, fmt.Errorf("client: error making http request: %w", err)
	}
	defer res.Body.Close()

	if res.StatusCode != 200 {
		return CompanyData{}, fmt.Errorf("status code error: %d %s", res.StatusCode, res.Status)
	}

	data := parseResponse(res)
	return data, nil
}

func parseResponse(res *http.Response) CompanyData {
	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		log.Fatal(err)
	}

	data := CompanyData{}
	data.Kpp = doc.Find("#clip_kpp").Text()
	data.Inn = doc.Find("#clip_inn").Text()
	data.CompanyName = doc.Find(".company-name").Text()
	doc.Find(".td1 .company-row").Each(func(i int, s *goquery.Selection) {
		if s.Find(".company-info__title").Text() == "Руководитель" {
			data.OwnerFullName = s.Find(".company-info__text").Text()
		}
	})

	return data
}
