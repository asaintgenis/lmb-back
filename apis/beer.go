package apis

import (
	"strconv"

	"github.com/go-ozzo/ozzo-routing"
	"github.com/go-ozzo/ozzo-validation"
	"gitlab.com/locatemybeer/lmb-back/app"
	"gitlab.com/locatemybeer/lmb-back/errors"
	"gitlab.com/locatemybeer/lmb-back/models"
)

type (
	// beerService specifies the interface for the beer service needed by beerResource.
	beerService interface {
		Get(rs app.RequestScope, id uint) (*models.Beer, error)
		Query(rs app.RequestScope, offset, limit int) ([]models.Beer, error)
		Count(rs app.RequestScope) (int, error)
		Create(rs app.RequestScope, model *models.Beer) (*models.Beer, error)
		Update(rs app.RequestScope, id uint, model *models.Beer) (*models.Beer, error)
		Delete(rs app.RequestScope, id uint) (*models.Beer, error)
	}

	// beerResource defines the handlers for the CRUD APIs.
	beerResource struct {
		service beerService
	}
)

// ServeBeerResource sets up the routing of beer endpoints and the corresponding handlers.
func ServeBeerResource(rg *routing.RouteGroup, service beerService) {
	r := &beerResource{service}
	rg.Get("/beers/<id>", r.get)
	rg.Get("/beers", r.query)
	rg.Post("/beers", r.create)
	rg.Put("/beers/<id>", r.update)
	rg.Delete("/beers/<id>", r.delete)
}

func (r *beerResource) get(c *routing.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return errors.InvalidParameter("ID must be an unsigned integer")
	}

	response, err := r.service.Get(app.GetRequestScope(c), uint(id))
	if err != nil {
		return errors.NotFound("beer")
	}

	return c.Write(response)
}

func (r *beerResource) query(c *routing.Context) error {
	rs := app.GetRequestScope(c)
	count, err := r.service.Count(rs)
	if err != nil {
		return errors.InternalServerError(err)
	}
	paginatedList := getPaginatedListFromRequest(c, count)
	items, err := r.service.Query(app.GetRequestScope(c), paginatedList.Offset(), paginatedList.Limit())
	if err != nil {
		return errors.InternalServerError(err)
	}
	paginatedList.Items = items
	return c.Write(paginatedList)
}

func (r *beerResource) create(c *routing.Context) error {
	var model models.Beer
	if err := c.Read(&model); err != nil {
		return errors.InvalidParameter("ID must be an unsigned integer")
	}
	response, err := r.service.Create(app.GetRequestScope(c), &model)
	if err != nil {
		castedErr, ok := err.(validation.Errors)
		if ok {
			return errors.InvalidData(castedErr)
		} else {
			return errors.InternalServerError(err)
		}
	}

	return c.Write(response)
}

func (r *beerResource) update(c *routing.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return errors.InvalidParameter("ID must be an unsigned integer")
	}

	rs := app.GetRequestScope(c)

	model, err := r.service.Get(rs, uint(id))
	if err != nil {
		return errors.NotFound("beer")
	}

	if err := c.Read(model); err != nil {
		return errors.InternalServerError(err)
	}

	response, err := r.service.Update(rs, uint(id), model)
	if err != nil {
		castedErr, ok := err.(validation.Errors)
		if ok {
			return errors.InvalidData(castedErr)
		} else {
			return errors.InternalServerError(err)
		}
	}

	return c.Write(response)
}

func (r *beerResource) delete(c *routing.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return errors.InvalidParameter("ID must be an unsigned integer")
	}

	response, err := r.service.Delete(app.GetRequestScope(c), uint(id))
	if err != nil {
		return errors.NotFound("beer")
	}

	return c.Write(response)
}
