package zpl

type Printable struct {
	Type PrintType `json:"type"`

	ILC            string `json:"ilc"`
	CommonILC      string `json:"common_ilc"`
	CuttingDate    string `json:"cutting_date"`
	RetouchingDate string `json:"retouching_date,omitempty"`

	Basket int     `json:"basket"`
	Loin   int     `json:"loin"`
	Weight float64 `json:"weight"`
	Grade  string  `json:"grade"`
}
