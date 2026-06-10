package driver

import (
	"errors"
)

type (
	// ModelRegistry holds the models registered for multitenancy support, categorizing them into
	// shared and tenant-specific models. Not intended for direct use in application code.
	ModelRegistry struct {
		SharedModels []TenantTabler // SharedModels contains the models that are shared across tenants.
		TenantModels []TenantTabler // TenantModels contains the models that are specific to a tenant.
	}
)

// NewModelRegistry creates and initializes a new ModelRegistry with the provided models, categorizing them into
// shared and tenant-specific based on their characteristics. It returns an error if any model fails validation.
// Not intended for direct use in application code.
func NewModelRegistry(models ...TenantTabler) (*ModelRegistry, error) {
	var (
		registry = &ModelRegistry{
			SharedModels: make([]TenantTabler, 0, len(models)),
			TenantModels: make([]TenantTabler, 0, len(models)),
		}
		errs []error
	)

	for _, model := range models {
		if model.IsSharedModel() {
			registry.SharedModels = append(registry.SharedModels, model)
		} else {
			registry.TenantModels = append(registry.TenantModels, model)
		}
	}

	if len(errs) > 0 {
		return nil, errors.Join(errs...)
	}

	return registry, nil
}
