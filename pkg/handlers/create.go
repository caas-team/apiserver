package handlers

import (
	"github.com/rancher/wrangler/pkg/schemas/validation"
	"gitlab.devops.telekom.de/caas/rancher/apiserver/pkg/apierror"
	"gitlab.devops.telekom.de/caas/rancher/apiserver/pkg/parse"
	"gitlab.devops.telekom.de/caas/rancher/apiserver/pkg/types"
)

func CreateHandler(apiOp *types.APIRequest) (types.APIObject, error) {
	var err error

	if err := apiOp.AccessControl.CanCreate(apiOp, apiOp.Schema); err != nil {
		return types.APIObject{}, err
	}

	data, err := parse.Body(apiOp.Request)
	if err != nil {
		return types.APIObject{}, err
	}

	store := apiOp.Schema.Store
	if store == nil {
		return types.APIObject{}, apierror.NewAPIError(validation.NotFound, "no store found")
	}

	data, err = store.Create(apiOp, apiOp.Schema, data)
	if err != nil {
		return types.APIObject{}, err
	}

	return data, nil
}
