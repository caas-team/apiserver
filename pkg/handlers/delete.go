package handlers

import (
	"github.com/rancher/wrangler/pkg/schemas/validation"
	"gitlab.devops.telekom.de/caas/rancher/apiserver/pkg/apierror"
	"gitlab.devops.telekom.de/caas/rancher/apiserver/pkg/types"
)

func DeleteHandler(request *types.APIRequest) (types.APIObject, error) {
	if err := request.AccessControl.CanDelete(request, types.APIObject{}, request.Schema); err != nil {
		return types.APIObject{}, err
	}

	store := request.Schema.Store
	if store == nil {
		return types.APIObject{}, apierror.NewAPIError(validation.NotFound, "no store found")
	}

	return store.Delete(request, request.Schema, request.Name)
}
