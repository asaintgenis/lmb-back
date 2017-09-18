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
	// beerStyleService specifies the interface for the beerstyle service needed by beerStyleResource.
	beerStyleService interface {
		Get(rs app.RequestScope, id uint) (*models.BeerStyle, error)
		Query(rs app.RequestScope, offset, limit int) ([]models.BeerStyle, error)
		Count(rs app.RequestScope) (int, error)
		Create(rs app.RequestScope, model *models.BeerStyle) (*models.BeerStyle, error)
		Update(rs app.RequestScope, id uint, model *models.BeerStyle) (*models.BeerStyle, error)
		Delete(rs app.RequestScope, id uint) (*models.BeerStyle, error)
	}

	// beerResource defines the handlers for the CRUD APIs.
	beerStyleResource struct {
		service beerStyleService
	}
)

// ServeBeerResource sets up the routing of beer endpoints and the corresponding handlers.
func ServeBeerStyleResource(rg *routing.RouteGroup, service beerStyleService) {
	r := &beerStyleResource{service}
	rg.Get("/beersStyles/<id>", r.get)
	rg.Get("/beersStyles", r.query)
	rg.Post("/beersStyles", r.create)
	rg.Put("/beersStyles/<id>", r.update)
	rg.Delete("/beersStyles/<id>", r.delete)
}

func (r *beerStyleResource) get(c *routing.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return errors.InvalidParameter("ID must be an unsigned integer")
	}

	response, err := r.service.Get(app.GetRequestScope(c), uint(id))
	if err != nil {
		return errors.NotFound("beerStyle")
	}

	return c.Write(response)
}

func (r *beerStyleResource) query(c *routing.Context) error {
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

func (r *beerStyleResource) create(c *routing.Context) error {
	var model models.BeerStyle
	if err := c.Read(&model); err != nil {
		return errors.InternalServerError(err)
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

func (r *beerStyleResource) update(c *routing.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return errors.InvalidParameter("ID must be an unsigned integer")
	}

	rs := app.GetRequestScope(c)

	model, err := r.service.Get(rs, uint(id))
	if err != nil {
		return errors.NotFound("beerStyle")
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

func (r *beerStyleResource) delete(c *routing.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return errors.InvalidParameter("ID must be an unsigned integer")
	}

	response, err := r.service.Delete(app.GetRequestScope(c), uint(id))
	if err != nil {
		return errors.NotFound("beerStyle")
	}

	return c.Write(response)
}
