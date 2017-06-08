package response

type Catalog struct {
	Services *[]Service `json:"services"`
}
