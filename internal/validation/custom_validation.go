package validation

import (
	"path/filepath"
	"regexp"
	"strconv"
	"strings"
	"user-management-api/internal/utils"

	"github.com/go-playground/validator/v10"
)

func RegisterCustomValidation(v *validator.Validate) error {
	var blockedDomains = map[string]bool{
		"gmail.com": true,
		"yahoo.com": true,
		"hotmail.com": true,
		"outlook.com": true,
		"icloud.com": true,
		"live.com": true,
		"msn.com": true,
	}

	v.RegisterValidation("email_address", func(fl validator.FieldLevel) bool {
		email := fl.Field().String()
		domain := strings.Split(email, "@")[1]
		return !blockedDomains[utils.NormalizeString(domain)]
	})

	v.RegisterValidation("password_strong", func(fl validator.FieldLevel) bool {
		password := fl.Field().String()

		if len(password) < 8 {
			return false
		}

		
		if !regexp.MustCompile(`[A-Z]`).MatchString(password) {
			return false
		}

		if !regexp.MustCompile(`[a-z]`).MatchString(password) {
			return false
		}
		
		if !regexp.MustCompile(`[0-9]`).MatchString(password) {
			return false
		}

		return true
	})

	var slugRegex = regexp.MustCompile(`^[a-z0-9]+(?:[-.][a-z0-9]+)*$`)
	v.RegisterValidation("slug", func(fl validator.FieldLevel) bool {
		return slugRegex.MatchString(fl.Field().String())
	})

	var searchRegex = regexp.MustCompile(`^[a-zA-Z0-9\s]+$`)
	v.RegisterValidation("search", func(fl validator.FieldLevel) bool {
		return searchRegex.MatchString(fl.Field().String())
	})

	v.RegisterValidation("min_int", func(fl validator.FieldLevel) bool {
		minStr := fl.Param()
		minVal, err := strconv.ParseInt(minStr, 10, 64)
		if err != nil {
			return false
		}

		return fl.Field().Int() >= minVal
	})

	v.RegisterValidation("max_int", func(fl validator.FieldLevel) bool {
		maxStr := fl.Param()
		maxVal, err := strconv.ParseInt(maxStr, 10, 64)
		if err != nil {
			return false
		}

		return fl.Field().Int() <= maxVal
	})

	v.RegisterValidation("file_ext", func(fl validator.FieldLevel) bool {
		filename := fl.Field().String()

		allowedStr := fl.Param()
		if allowedStr == "" {
			return false
		}

		allowedExt := strings.Fields(allowedStr)
		ext := strings.TrimPrefix(strings.ToLower(filepath.Ext(filename)), ".")

		for _, allowed := range allowedExt {
			if ext == strings.ToLower(allowed) {
				return true
			}
		}

		return false
	})

	return nil
}