package searchads

import (
	"context"
	"fmt"
)

// AdGroupService struct to hold individual service
type AdGroupService service

type AdGroup struct {
	ID                     int64                `json:"id,omitempty"`
	CampaignID             int64                `json:"campaignId,omitempty"`
	Name                   string               `json:"name,omitempty"`
	CpaGoal                *Amount              `json:"cpaGoal"`
	StartTime              string               `json:"startTime,omitempty"`
	EndTime                string               `json:"endTime"`
	AutomatedKeywordsOptIn bool                 `json:"automatedKeywordsOptIn"`
	DefaultBidAmount       *Amount              `json:"defaultBidAmount,omitempty"`
	TargetingDimensions    *TargetingDimensions `json:"targetingDimensions,omitempty"`
	OrgID                  int64                `json:"orgId,omitempty"`
	ModificationTime       string               `json:"modificationTime,omitempty"`
	Status                 Status               `json:"status,omitempty"`
	ServingStatus          ServingStatus        `json:"servingStatus,omitempty"`
	DisplayStatus          DisplayStatus        `json:"displayStatus,omitempty"`
	ServingStateReasons    interface{}          `json:"servingStateReasons,omitempty"`
	Deleted                bool                 `json:"deleted,omitempty"`
	PricingModel           PricingModel         `json:"pricingModel,omitempty"`
}

type TargetingDimensions struct {
	Age            *Age            `json:"age"`
	Gender         *Gender         `json:"gender"`
	Country        *Country        `json:"country"`
	AdminArea      *AdminArea      `json:"adminArea"`
	Locality       *Locality       `json:"locality"`
	DeviceClass    *DeviceClass    `json:"deviceClass"`
	Daypart        *Daypart        `json:"daypart"`
	AppDownloaders *AppDownloaders `json:"appDownloaders"`
}

type Country struct {
	Included []string `json:"included,omitempty"`
}

type Age struct {
	Included []AgeObj `json:"included,omitempty"`
}

type AgeObj struct {
	MinAge int `json:"minAge,omitempty"`
	MaxAge int `json:"maxAge,omitempty"`
}

type Gender struct {
	Included []string `json:"included,omitempty"`
}

type AdminArea struct {
	Included []string `json:"included,omitempty"`
}

type Locality struct {
	Included []string `json:"included,omitempty"`
}

type Daypart struct {
	UserTime UserTime `json:"userTime,omitempty"`
}

type UserTime struct {
	Included []int `json:"included,omitempty"`
}

type DeviceClass struct {
	Included []string `json:"included,omitempty"`
}
type AppDownloaders struct {
	Included []string `json:"included,omitempty"`
	Excluded []string `json:"excluded,omitempty"`
}

// List function to get Adgroups from campaign
func (s *AdGroupService) List(ctx context.Context, campaignID int64, opt *ListOptions) ([]*AdGroup, *Response, error) {
	if campaignID == 0 {
		return nil, nil, fmt.Errorf("campaignID can not be 0")
	}
	u, err := addOptions(fmt.Sprintf("campaigns/%d/adgroups", campaignID), opt)
	if err != nil {
		return nil, nil, err
	}
	req, err := s.client.NewRequest("GET", u, nil)
	if err != nil {
		return nil, nil, err
	}

	adGroups := []*AdGroup{}

	resp, err := s.client.Do(ctx, req, &adGroups)
	if err != nil {
		return nil, resp, err
	}

	return adGroups, resp, nil
}

// Get function to get specific AdGroup by id and campaignID
func (s *AdGroupService) Get(ctx context.Context, campaignID, id int64) (*AdGroup, *Response, error) {
	if campaignID == 0 {
		return nil, nil, fmt.Errorf("campaignID can not be 0")
	}
	u := fmt.Sprintf("campaigns/%d/adgroups/%d", campaignID, id)
	req, err := s.client.NewRequest("GET", u, nil)
	if err != nil {
		return nil, nil, err
	}

	adGroup := new(AdGroup)
	resp, err := s.client.Do(ctx, req, adGroup)
	if err != nil {
		return nil, resp, err
	}

	return adGroup, resp, nil
}

// Create will create a new AdGroup on a specific campaign
func (s *AdGroupService) Create(ctx context.Context, campaignID int64, data *AdGroup) (*AdGroup, *Response, error) {
	if campaignID == 0 {
		return nil, nil, fmt.Errorf("campaignID can not be 0")
	}
	u := fmt.Sprintf("campaigns/%d/adgroups", campaignID)
	req, err := s.client.NewRequest("POST", u, data)
	if err != nil {
		return nil, nil, err
	}

	adG := new(AdGroup)
	resp, err := s.client.Do(ctx, req, adG)
	if err != nil {
		return nil, resp, err
	}

	return adG, resp, nil
}

// DoEdit will update an existing AdGroup on a Campaign
func (s *AdGroupService) Edit(ctx context.Context, campaignID int64, id int64, data *AdGroup) (*AdGroup, *Response, error) {
	if campaignID == 0 {
		return nil, nil, fmt.Errorf("campaignID can not be 0")
	}
	if id == 0 {
		return nil, nil, fmt.Errorf("id can not be 0")
	}
	u := fmt.Sprintf("campaigns/%d/adgroups/%d", campaignID, id)
	req, err := s.client.NewRequest("PUT", u, data)
	if err != nil {
		return nil, nil, err
	}

	adG := new(AdGroup)
	resp, err := s.client.Do(ctx, req, adG)
	if err != nil {
		return nil, resp, err
	}

	return adG, resp, nil
}

// Delete will remove an existing AdGroup on a Campaign
func (s *AdGroupService) Delete(ctx context.Context, campaignID int64, id int64) (*Response, error) {
	if campaignID == 0 {
		return nil, fmt.Errorf("campaignID can not be 0")
	}
	if id == 0 {
		return nil, fmt.Errorf("id can not be 0")
	}
	u := fmt.Sprintf("campaigns/%d/adgroups/%d", campaignID, id)
	req, err := s.client.NewRequest("DELETE", u, nil)
	if err != nil {
		return nil, err
	}

	return s.client.Do(ctx, req, nil)
}

// Fetches ad groups within a campaign.
func (s *AdGroupService) Find(ctx context.Context, campaignID int64, sel *Selector) ([]*AdGroup, *Response, error) {
	if campaignID == 0 {
		return nil, nil, fmt.Errorf("campaignID can not be 0")
	}
	req, err := s.client.NewRequest("POST", fmt.Sprintf("campaigns/%d/adgroups/find", campaignID), sel)
	if err != nil {
		return nil, nil, err
	}

	adGroups := []*AdGroup{}

	resp, err := s.client.Do(ctx, req, &adGroups)
	if err != nil {
		return nil, resp, err
	}

	return adGroups, resp, nil
}
