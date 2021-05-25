package entity

type ResponseInfo struct {
	Status   string `json:"status"`
	Count    string `json:"count"`
	Info     string `json:"info"`
	InfoCode string `json:"infoCode"`
	Lives    []Live `json:"lives"`
}

type Live struct {
	Province      string `json:"province"`
	City          string `json:"city"`
	Adcode        string `json:"adcode"`
	Weather       string `json:"weather"`
	Temperature   string `json:"temperature"`
	Winddirection string `json:"winddirection"`
	Windpower     string `json:"windpower"`
	Humidity      string `json:"humidity"`
	Reporttime    string `json:"reporttime"`
}
