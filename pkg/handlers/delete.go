package handlers

import (
	"github.com/caas-team/apiserver/pkg/apierror"
	"github.com/caas-team/apiserver/pkg/types"
	"github.com/rancher/wrangler/v2/pkg/schemas/validation"
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
