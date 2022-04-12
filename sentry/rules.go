package sentry

import (
	"net/http"
	"time"

	"github.com/dghubble/sling"
)

// Rule represents an alert rule configured for this project.
// https://github.com/getsentry/sentry/blob/9.0.0/src/sentry/api/serializers/models/rule.py
type Rule struct {
	ID          string          `json:"id"`
	ActionMatch string          `json:"actionMatch"`
	FilterMatch string          `json:"filterMatch"`
	Environment *string         `json:"environment,omitempty"`
	Frequency   int             `json:"frequency"`
	Name        string          `json:"name"`
	Conditions  []ConditionType `json:"conditions"`
	Actions     []ActionType    `json:"actions"`
	Filters     []FilterType    `json:"filters"`
	Created     time.Time       `json:"dateCreated"`
}

// RuleService provides methods for accessing Sentry project
// client key API endpoints.
// https://docs.sentry.io/api/projects/
type RuleService struct {
	sling *sling.Sling
}

func newRuleService(sling *sling.Sling) *RuleService {
	return &RuleService{
		sling: sling,
	}
}

// List alert rules configured for a project.
func (s *RuleService) List(organizationSlug string, projectSlug string) ([]Rule, *http.Response, error) {
	rules := new([]Rule)
	apiError := new(APIError)
	resp, err := s.sling.New().Get("projects/"+organizationSlug+"/"+projectSlug+"/rules/").Receive(rules, apiError)
	return *rules, resp, relevantError(err, *apiError)
}

// ConditionType for defining conditions.
type ConditionType map[string]interface{}

// ActionType for defining actions.
type ActionType map[string]interface{}

// FilterType for defining actions.
type FilterType map[string]interface{}

// CreateRuleParams are the parameters for RuleService.Create.
type CreateRuleParams struct {
	ActionMatch string          `json:"actionMatch"`
	FilterMatch string          `json:"filterMatch"`
	Environment string          `json:"environment,omitempty"`
	Frequency   int             `json:"frequency"`
	Name        string          `json:"name"`
	Conditions  []ConditionType `json:"conditions"`
	Actions     []ActionType    `json:"actions"`
	Filters     []FilterType    `json:"filters"`
}

// CreateRuleActionParams models the actions when creating the action for the rule.
type CreateRuleActionParams struct {
	ID        string `json:"id"`
	Tags      string `json:"tags"`
	Channel   string `json:"channel"`
	Workspace string `json:"workspace"`

	Action    string `json:"action,omitempty"`
	Service   string `json:"service,omitempty"`
	ChannelID string `json:"channel_id,omitempty"`
}

// CreateRuleConditionParams models the conditions when creating the action for the rule.
type CreateRuleConditionParams struct {
	ID       string `json:"id"`
	Interval string `json:"interval"`
	Value    int    `json:"value"`
	Level    int    `json:"level"`
	Match    string `json:"match"`

	Attribute string `json:"attribute,omitempty"`
	Key       string `json:"key,omitempty"`
	Name      string `json:"name"`
}

// Create a new alert rule bound to a project.
func (s *RuleService) Create(organizationSlug string, projectSlug string, params *CreateRuleParams) (*Rule, *http.Response, error) {
	rule := new(Rule)
	apiError := new(APIError)
	resp, err := s.sling.New().Post("projects/"+organizationSlug+"/"+projectSlug+"/rules/").BodyJSON(params).Receive(rule, apiError)
	return rule, resp, relevantError(err, *apiError)
}

// Update a rule.
func (s *RuleService) Update(organizationSlug string, projectSlug string, ruleID string, params *Rule) (*Rule, *http.Response, error) {
	rule := new(Rule)
	apiError := new(APIError)
	resp, err := s.sling.New().Put("projects/"+organizationSlug+"/"+projectSlug+"/rules/"+ruleID+"/").BodyJSON(params).Receive(rule, apiError)
	return rule, resp, relevantError(err, *apiError)
}

// Delete a rule.
func (s *RuleService) Delete(organizationSlug string, projectSlug string, ruleID string) (*http.Response, error) {
	apiError := new(APIError)
	resp, err := s.sling.New().Delete("projects/"+organizationSlug+"/"+projectSlug+"/rules/"+ruleID+"/").Receive(nil, apiError)
	return resp, relevantError(err, *apiError)
}

// Canva

// Get a rule.
func (s *RuleService) Get(organizationSlug string, projectSlug string, ruleId string) (*Rule, *http.Response, error) {
	rule := new(Rule)
	apiError := new(APIError)
	resp, err := s.sling.New().Get("projects/"+organizationSlug+"/"+projectSlug+"/rules/"+ruleId+"/").Receive(rule, apiError)
	return rule, resp, relevantError(err, *apiError)
}
