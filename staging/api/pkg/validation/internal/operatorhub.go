package internal

import (
	"fmt"
	"net/mail"
	"net/url"
	"strings"

	"github.com/operator-framework/api/pkg/manifests"
	"github.com/operator-framework/api/pkg/operators/v1alpha1"
	"github.com/operator-framework/api/pkg/validation/errors"
	interfaces "github.com/operator-framework/api/pkg/validation/interfaces"
)

var OperatorHubValidator interfaces.Validator = interfaces.ValidatorFunc(validateOperatorHub)

var validCapabilities = map[string]struct{}{
	"Basic Install":     struct{}{},
	"Seamless Upgrades": struct{}{},
	"Full Lifecycle":    struct{}{},
	"Deep Insights":     struct{}{},
	"Auto Pilot":        struct{}{},
}

var validMediatypes = map[string]struct{}{
	"image/gif":     struct{}{},
	"image/jpeg":    struct{}{},
	"image/png":     struct{}{},
	"image/svg+xml": struct{}{},
}

var validCategories = map[string]struct{}{
	"AI/Machine Learning":    struct{}{},
	"Application Runtime":    struct{}{},
	"Big Data":               struct{}{},
	"Cloud Provider":         struct{}{},
	"Developer Tools":        struct{}{},
	"Database":               struct{}{},
	"Integration & Delivery": struct{}{},
	"Logging & Tracing":      struct{}{},
	"Monitoring":             struct{}{},
	"Networking":             struct{}{},
	"OpenShift Optional":     struct{}{},
	"Security":               struct{}{},
	"Storage":                struct{}{},
	"Streaming & Messaging":  struct{}{},
}

func validateOperatorHub(objs ...interface{}) (results []errors.ManifestResult) {
	for _, obj := range objs {
		switch v := obj.(type) {
		case *manifests.Bundle:
			results = append(results, validateBundleOperatorHub(v))
		}
	}
	return results
}

func validateBundleOperatorHub(bundle *manifests.Bundle) errors.ManifestResult {
	result := errors.ManifestResult{Name: bundle.Name}

	if bundle == nil {
		result.Add(errors.ErrInvalidBundle("Bundle is nil", nil))
		return result
	}

	if bundle.CSV == nil {
		result.Add(errors.ErrInvalidBundle("Bundle csv is nil", bundle.Name))
		return result
	}

	errs := validateHubCSVSpec(*bundle.CSV)
	for _, err := range errs {
		result.Add(errors.ErrInvalidCSV(err.Error(), bundle.CSV.GetName()))
	}

	return result
}

func validateHubCSVSpec(csv v1alpha1.ClusterServiceVersion) []error {
	var errs []error

	if csv.Spec.Provider.Name == "" {
		errs = append(errs, fmt.Errorf("csv.Spec.Provider.Name not specified"))
	}

	for _, maintainer := range csv.Spec.Maintainers {
		if maintainer.Name == "" || maintainer.Email == "" {
			errs = append(errs, fmt.Errorf("csv.Spec.Maintainers elements should contain both name and email"))
		}
		if maintainer.Email != "" {
			_, err := mail.ParseAddress(maintainer.Email)
			if err != nil {
				errs = append(errs, fmt.Errorf("csv.Spec.Maintainers email %s is invalid: %v", maintainer.Email, err))
			}
		}
	}

	for _, link := range csv.Spec.Links {
		if link.Name == "" || link.URL == "" {
			errs = append(errs, fmt.Errorf("csv.Spec.Links elements should contain both name and url"))
		}
		if link.URL != "" {
			_, err := url.ParseRequestURI(link.URL)
			if err != nil {
				errs = append(errs, fmt.Errorf("csv.Spec.Links url %s is invalid: %v", link.URL, err))
			}
		}
	}

	if csv.GetAnnotations() == nil {
		csv.SetAnnotations(make(map[string]string))
	}

	if capability, ok := csv.ObjectMeta.Annotations["capabilities"]; ok {
		if _, ok := validCapabilities[capability]; !ok {
			errs = append(errs, fmt.Errorf("csv.Metadata.Annotations.Capabilities %s is not a valid capabilities level", capability))
		}
	}

	if csv.Spec.Icon != nil {
		// only one icon is allowed
		if len(csv.Spec.Icon) != 1 {
			errs = append(errs, fmt.Errorf("csv.Spec.Icon should only have one element"))
		}

		icon := csv.Spec.Icon[0]
		if icon.MediaType == "" || icon.Data == "" {
			errs = append(errs, fmt.Errorf("csv.Spec.Icon elements should contain both data and mediatype"))
		}

		if icon.MediaType != "" {
			if _, ok := validMediatypes[icon.MediaType]; !ok {
				errs = append(errs, fmt.Errorf("csv.Spec.Icon %s does not have a valid mediatype", icon.MediaType))
			}
		}
	} else {
		errs = append(errs, fmt.Errorf("csv.Spec.Icon not specified"))
	}

	if categories, ok := csv.ObjectMeta.Annotations["categories"]; ok {
		categorySlice := strings.Split(categories, ",")

		for _, category := range categorySlice {
			if _, ok := validCategories[category]; !ok {
				errs = append(errs, fmt.Errorf("csv.Metadata.Annotations.Categories %s is not a valid category", category))
			}
		}
	}

	return errs
}
