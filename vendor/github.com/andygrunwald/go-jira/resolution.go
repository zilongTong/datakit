package jira

// ResolutionService handles resolutions for the JIRA instance / API.
//
// JIRA API docs: https://developer.atlassian.com/cloud/jira/platform/rest/#api-Resolution
type ResolutionService struct {
	client *Client
}

// Resolution represents a resolution of a JIRA issue.
// Typical types are "Fixed", "Suspended", "Won't Fix", ...
type Resolution struct {
	Self        string `json:"self" structs:"self"`
	ID          string `json:"id" structs:"id"`
	Description string `json:"description" structs:"description"`
	Name        string `json:"name" structs:"name"`
}

// GetList gets all resolutions from JIRA
//
// JIRA API docs: https://developer.atlassian.com/cloud/jira/platform/rest/#api-api-2-resolution-get
func (s *ResolutionService) GetList() ([]Resolution, *Response, error) {
	apiEndpoint := "rest/api/2/resolution"
	req, err := s.client.NewRequest("GET", apiEndpoint, nil)
	if err != nil {
		return nil, nil, err
	}

	resolutionList := []Resolution{}
	resp, err := s.client.Do(req, &resolutionList)
	if err != nil {
		return nil, resp, NewJiraError(resp, err)
	}
	return resolutionList, resp, nil
}
