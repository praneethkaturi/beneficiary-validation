package model

type BeneficiaryValidateRequest struct {
	AccountNumber string `json:"account_number" binding:"required"`
	Name          string `json:"name" binding:"required"`
	RoutingNumber string `json:"routing_number" binding:"required"`
	Pincode       string `json:"pincode"`
	StreetAddress string `json:"street_address"`
	Currency      string `json:"currency" binding:"required"`
	RoutingType   string `json:"routing_type" binding:"required"`
}

type ValidateRule struct {
	Field      string   `json:"Field"`
	Type       string   `json:"Type"`
	Required   string   `json:"Required"`
	MinLength  int      `json:"MinLength"`
	MaxLength  int      `json:"MaxLength"`
	Regex      string   `json:"Regex"`
	EnumValues []string `json:"EnumValues"`
	Value      string   `json:"Value"`
}

type ConditionalRule struct {
	Conditions []Condition    `json:"Conditions"`
	Rules      []ValidateRule `json:"Rules"`
}

type Condition struct {
	Field string `json:"Field"`
	Value string `json:"Value"`
}

type ConfigRule struct {
	GeneralRules     []ValidateRule    `json:"GeneralRules"`
	ConditionalRules []ConditionalRule `json:"ConditionalRules"`
}

type CurrencyConfigRule struct {
	Currency   string     `json:"Currency"`
	ConfigRule ConfigRule `json:"ConfigRule"`
}
