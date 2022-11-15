package api

import (
	"crypto/tls"
	"fmt"
	"html"
	"regexp"

	"github.com/NoF0rte/gophish-cli/pkg/api/models"
	"github.com/go-resty/resty/v2"
)

type Client struct {
	client *resty.Client
}

func NewClient(url string, apiKey string) *Client {
	client := resty.New().
		SetBaseURL(url).
		SetTLSClientConfig(&tls.Config{
			InsecureSkipVerify: true,
		})
	if apiKey != "" {
		client.SetAuthToken(apiKey)
	}

	return &Client{
		client: client,
	}
}

func NewClientFromCredentials(url string, username string, password string) (*Client, error) {
	c := NewClient(url, "")

	apiKey, err := c.GetAPIKey(username, password)
	if err != nil {
		return nil, err
	}

	c.client.SetAuthToken(apiKey)

	return c, nil
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

func (c *Client) delete(path string, body interface{}, result interface{}) (*resty.Response, interface{}, error) {
	req := c.newRequest(result)
	if body != nil {
		req.SetBody(body)
	}

	resp, err := req.Delete(path)
	if err != nil {
		return nil, nil, err
	}

	r := resp.Result()

	return resp, r, nil
}

func (c *Client) GetAPIKey(username string, password string) (string, error) {
	resp, err := c.client.R().Get("/login")
	if err != nil {
		return "", err
	}

	cookies := resp.Cookies()
	csrfTokenRe := regexp.MustCompile(`name="csrf_token"\s*value="([^"]+)"`)

	body := string(resp.Body())
	matches := csrfTokenRe.FindStringSubmatch(body)
	if len(matches) == 0 {
		return "", fmt.Errorf("error finding csrf_token")
	}

	csrfToken := html.UnescapeString(matches[1])

	resp, err = c.client.R().
		SetCookies(cookies).
		SetFormData(map[string]string{
			"username":   username,
			"password":   password,
			"csrf_token": csrfToken,
		}).
		Post("/login")

	if err != nil {
		return "", err
	}

	if resp.IsError() {
		return "", fmt.Errorf("error: %s", resp.Status())
	}

	resp, err = c.client.R().
		SetCookies(resp.Cookies()).
		Get("/settings")
	if err != nil {
		return "", nil
	}

	body = string(resp.Body())
	apiKeyRe := regexp.MustCompile(`api_key\s*:\s*"([^"]+)"`)

	matches = apiKeyRe.FindStringSubmatch(body)
	if len(matches) == 0 {
		return "", fmt.Errorf("error finding api key")
	}

	return matches[1], nil
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

func (c *Client) GetSendingProfiles() ([]*models.SendingProfile, error) {
	var profiles []*models.SendingProfile
	_, _, err := c.get("/api/smtp/", &profiles)
	if err != nil {
		return nil, err
	}

	return profiles, nil
}

func (c *Client) GetSendingProfileByID(id int) (*models.SendingProfile, error) {
	profile := &models.SendingProfile{}
	_, _, err := c.get(fmt.Sprintf("/api/smtp/%d", id), profile)
	if err != nil {
		return nil, err
	}

	if profile.Id == 0 {
		return nil, nil
	}

	return profile, nil
}

func (c *Client) GetSendingProfileByName(name string) (*models.SendingProfile, error) {
	profiles, err := c.GetSendingProfiles()
	if err != nil {
		return nil, err
	}

	for _, t := range profiles {
		if t.Name == name {
			return t, nil
		}
	}

	return nil, nil
}

func (c *Client) GetSendingProfileByRegex(re string) ([]*models.SendingProfile, error) {
	profiles, err := c.GetSendingProfiles()
	if err != nil {
		return nil, err
	}

	var filtered []*models.SendingProfile
	regex := regexp.MustCompile(re)
	for _, t := range profiles {
		if regex.MatchString(t.Name) {
			filtered = append(filtered, t)
		}
	}

	return filtered, nil
}

func (c *Client) DeleteTemplateByID(id int64) (*models.GenericResponse, error) {
	r := &models.GenericResponse{}
	_, _, err := c.delete(fmt.Sprintf("/api/templates/%d", id), nil, r)
	if err != nil {
		return nil, err
	}

	return r, nil
}

func (c *Client) DeleteTemplateByName(name string) (*models.GenericResponse, error) {
	templates, err := c.GetTemplates()
	if err != nil {
		return nil, err
	}

	var template *models.Template
	for _, t := range templates {
		if t.Name == name {
			template = t
			break
		}
	}

	if template == nil {
		return nil, fmt.Errorf("template %s not found", name)
	}

	return c.DeleteTemplateByID(template.Id)
}

func (c *Client) CreateTemplate(template *models.Template) (*models.Template, error) {
	template.Id = 0 // Ensure the ID is always 0

	_, result, err := c.post("/api/templates/", template, &models.Template{})
	if err != nil {
		return nil, err
	}

	return result.(*models.Template), nil
}

func (c *Client) CreateTemplateFromFile(file string, vars map[string]string) (*models.Template, error) {
	template, err := models.TemplateFromFile(file, vars)
	if err != nil {
		return nil, err
	}

	return c.CreateTemplate(template)
}

func (c *Client) CreateSendingProfile(profile *models.SendingProfile) (*models.SendingProfile, error) {
	profile.Id = 0 // Ensure the ID is always 0

	if profile.Interface == "" {
		profile.Interface = models.InterfaceSMTP
	}

	_, result, err := c.post("/api/smtp/", profile, &models.SendingProfile{})
	if err != nil {
		return nil, err
	}

	return result.(*models.SendingProfile), nil
}

func (c *Client) CreateSendingProfileFromFile(file string, vars map[string]string) (*models.SendingProfile, error) {
	profile, err := models.ProfileFromFile(file, vars)
	if err != nil {
		return nil, err
	}

	return c.CreateSendingProfile(profile)
}

func (c *Client) DeleteSendingProfileByID(id int64) (*models.GenericResponse, error) {
	r := &models.GenericResponse{}
	_, _, err := c.delete(fmt.Sprintf("/api/smtp/%d", id), nil, r)
	if err != nil {
		return nil, err
	}

	return r, nil
}

func (c *Client) DeleteSendingProfileByName(name string) (*models.GenericResponse, error) {
	profiles, err := c.GetSendingProfiles()
	if err != nil {
		return nil, err
	}

	var profile *models.SendingProfile
	for _, s := range profiles {
		if s.Name == name {
			profile = s
			break
		}
	}

	if profile == nil {
		return nil, fmt.Errorf("profile %s not found", name)
	}

	return c.DeleteSendingProfileByID(profile.Id)
}
