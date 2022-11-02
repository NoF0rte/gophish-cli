package api

import (
	"crypto/tls"
	"fmt"
	"regexp"

	"github.com/NoF0rte/gophish-cli/pkg/api/models"
	"github.com/go-resty/resty/v2"
)

type Client struct {
	client *resty.Client
}

func NewClient(url string, apiKey string) *Client {
	return &Client{
		client: resty.New().
			SetAuthToken(apiKey).
			SetBaseURL(url).
			SetTLSClientConfig(&tls.Config{
				InsecureSkipVerify: true,
			}),
	}
}

func (c *Client) newRequest(result interface{}) *resty.Request {
	req := c.client.R()
	if result != nil {
		req = req.SetResult(result)
	}
	return req
}

func (c *Client) get(path string, result interface{}) (*resty.Response, interface{}, error) {
	resp, err := c.newRequest(result).Get(path)
	if err != nil {
		return nil, nil, err
	}

	r := resp.Result()

	return resp, r, nil
}

func (c *Client) post(path string, body interface{}, result interface{}) (*resty.Response, interface{}, error) {
	req := c.newRequest(result)
	if body != nil {
		req.SetBody(body)
	}

	resp, err := req.Post(path)
	if err != nil {
		return nil, nil, err
	}

	r := resp.Result()

	return resp, r, nil
}

func (c *Client) GetTemplates() ([]*models.Template, error) {
	var templates []*models.Template
	_, _, err := c.get("/api/templates/", &templates)
	if err != nil {
		return nil, err
	}

	return templates, nil
}

func (c *Client) GetTemplateByID(id int) (*models.Template, error) {
	t := &models.Template{}
	_, _, err := c.get(fmt.Sprintf("/api/templates/%d", id), t)
	if err != nil {
		return nil, err
	}

	if t.Id == 0 {
		return nil, nil
	}

	return t, nil
}

func (c *Client) GetTemplateByName(name string) (*models.Template, error) {
	templates, err := c.GetTemplates()
	if err != nil {
		return nil, err
	}

	for _, t := range templates {
		if t.Name == name {
			return t, nil
		}
	}

	return nil, nil
}

func (c *Client) GetTemplatesByRegex(re string) ([]*models.Template, error) {
	templates, err := c.GetTemplates()
	if err != nil {
		return nil, err
	}

	var filtered []*models.Template
	regex := regexp.MustCompile(re)
	for _, t := range templates {
		if regex.MatchString(t.Name) {
			filtered = append(filtered, t)
		}
	}

	return filtered, nil
}

func (c *Client) CreateTemplate(template models.Template) (*models.Template, error) {
	template.Id = 0 // Ensure the ID is always 0

	_, result, err := c.post("/api/templates/", template, &models.Template{})
	if err != nil {
		return nil, err
	}

	return result.(*models.Template), nil
}
