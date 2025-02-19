package export

import (
	"time"

	"github.com/NoF0rte/gophish-client/api/models"
)

type Campaign struct {
	Name       string           `yaml:"name"`
	LaunchDate time.Time        `yaml:"launch-date"`
	SendByDate time.Time        `yaml:"send-by-date"`
	Template   string           `yaml:"template"`
	Page       string           `yaml:"page"`
	SMTP       string           `yaml:"smtp"`
	Groups     []string         `yaml:"groups,omitempty"` // Currently GoPhish's API doesn't return this'
	Results    []*models.Result `yaml:"results,omitempty"`
	Timeline   []*models.Event  `yaml:"timeline,omitempty"`
	URL        string           `yaml:"url"`
}

func NewCampaign(c *models.Campaign, includeResults bool) Campaign {
	campaign := Campaign{
		Name:       c.Name,
		LaunchDate: c.LaunchDate,
		SendByDate: c.SendByDate,
		Template:   c.Template.Name,
		Page:       c.Page.Name,
		SMTP:       c.SMTP.Name,
		Results:    c.Results,
		Timeline:   c.Timeline,
		URL:        c.URL,
	}

	var groups []string
	for _, g := range c.Groups {
		groups = append(groups, g.Name)
	}

	campaign.Groups = groups

	if !includeResults {
		campaign.Results = nil
		campaign.Timeline = nil
	}

	return campaign
}
