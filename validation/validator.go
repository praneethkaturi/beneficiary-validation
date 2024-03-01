package validation

import (
	"beneficiary-validate/model"
	"reflect"

	"log"
)

func ApplyValidationRules(rules []model.ValidateRule, req *model.BeneficiaryValidateRequest) bool {
	v := reflect.ValueOf(*req)
	for _, rule := range rules {
		log.Printf("applying rule: %v", rule)
		field := v.FieldByName(rule.Field)

		if rule.Type == "Required" {
			if !field.IsValid() || field.String() == "" {
				log.Printf("field %s is required", rule.Field)
				return false
			}
		}

		if !field.IsValid() || field.Kind() != reflect.String {
			log.Printf("field not found or not a string: %s", field)
			return false
		}

		fieldValue := field.Interface().(string)
		switch rule.Type {
		case "ValidateLength":
			if !validateLength(fieldValue, rule.MinLength, rule.MaxLength) {
				log.Printf("field %s length is invalid", rule.Field)
				return false
			}
		case "ValidateRegex":
			if !validateRegex(fieldValue, rule.Regex) {
				log.Printf("field %s regex is invalid", rule.Field)
				return false
			}
		case "ValidateEnum":
			if !validateEnum(fieldValue, rule.EnumValues) {
				log.Printf("field %s enum is invalid", rule.Field)
				return false
			}
		}
	}
	return true
}

func ApplyConditionalRules(conditionalRules []model.ConditionalRule, req *model.BeneficiaryValidateRequest) bool {
	for _, cRule := range conditionalRules {
		allConditionsMet := true
		for _, condition := range cRule.Conditions {
			if !validateCondition(condition, req) {
				allConditionsMet = false
				break
			}
		}
		if allConditionsMet {
			if !ApplyValidationRules(cRule.Rules, req) {
				return false
			}
		}
	}
	return true
}

func validateCondition(condition model.Condition, req *model.BeneficiaryValidateRequest) bool {
	v := reflect.ValueOf(*req)
	field := v.FieldByName(condition.Field)
	if !field.IsValid() || field.Kind() != reflect.String {
		log.Printf("field not found or not a string: %s", field)
		return false
	}

	fieldValue := field.Interface().(string)
	return fieldValue == condition.Value
}
