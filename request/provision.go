package request

type Provision struct {
	ServiceID         string                       `json:"service_id"`
	PlanID            string                       `json:"plan_id"`
	Parameters        *map[interface{}]interface{} `json:"parameters,omitempty"`
	AcceptsIncomplete bool                         `json:"accepts_incomplete"`
	OrganizationGUID  string                       `json:"organization_guid"`
	SpaceGUID         string                       `json:"space_guid"`
}
