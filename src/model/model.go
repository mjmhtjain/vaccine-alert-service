package model

type Appointments struct {
	Centers []struct {
		CenterID     int    `json:"center_id"`
		Name         string `json:"name"`
		Address      string `json:"address"`
		StateName    string `json:"state_name"`
		DistrictName string `json:"district_name"`
		BlockName    string `json:"block_name"`
		Pincode      int    `json:"pincode"`
		Lat          int    `json:"lat"`
		Long         int    `json:"long"`
		From         string `json:"from"`
		To           string `json:"to"`
		FeeType      string `json:"fee_type"`
		Sessions     []struct {
			SessionID              string   `json:"session_id"`
			Date                   string   `json:"date"`
			AvailableCapacity      int      `json:"available_capacity"`
			MinAgeLimit            int      `json:"min_age_limit"`
			Vaccine                string   `json:"vaccine"`
			Slots                  []string `json:"slots"`
			AvailableCapacityDose1 int      `json:"available_capacity_dose1"`
			AvailableCapacityDose2 int      `json:"available_capacity_dose2"`
		} `json:"sessions"`
		VaccineFees []struct {
			Vaccine string `json:"vaccine"`
			Fee     string `json:"fee"`
		} `json:"vaccine_fees"`
	} `json:"centers"`
}

type StateDistricts struct {
	Districts []StateDistrict `json:"districts"`
	TTL       int             `json:"ttl"`
}

type StateDistrict struct {
	DistrictID   int    `json:"district_id"`
	DistrictName string `json:"district_name"`
}

type States struct {
	States []State `json:"states"`
	TTL    int     `json:"ttl"`
}

type State struct {
	StateID   int    `json:"state_id"`
	StateName string `json:"state_name"`
}
