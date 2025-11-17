package model

type ControlMethods struct {
	Organic  []string `json:"organic"`
	Chemical []string `json:"chemical"`
}

type Pest struct {
	ID             int      `json:"id"`
	CommonName     string   `json:"common_name"`
	ScientificName string   `json:"scientific_name"`
	PestType       string   `json:"pest_type"`
	AffectedParts  []string `json:"affected_parts"`
	Description    string   `json:"description"`
	Symptoms       []string `json:"symptoms"`
	ImageURL       string   `json:"image_url"`
	ControlMethods struct {
		Organic  []string `json:"organic"`
		Chemical []string `json:"chemical"`
	} `json:"control_methods"`
}

