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
	// barService specifies the interface for the bar service needed by barResource.
	barService interface {
		Get(rs app.RequestScope, id uint) (*models.Bar, error)
		Query(rs app.RequestScope, offset, limit int) ([]models.Bar, error)
		Count(rs app.RequestScope) (int, error)
		Create(rs app.RequestScope, model *models.Bar) (*models.Bar, error)
		Update(rs app.RequestScope, id uint, model *models.Bar) (*models.Bar, error)
		Delete(rs app.RequestScope, id uint) (*models.Bar, error)
	}

	// barResource defines the handlers for the CRUD APIs.
	barResource struct {
		service barService
	}
)

// ServeBarResource sets up the routing of bar endpoints and the corresponding handlers.
func ServeBarResource(rg *routing.RouteGroup, service barService) {
	r := &barResource{service}
	rg.Get("/bars/<id>", r.get)
	rg.Get("/bars", r.query)
	rg.Post("/bars", r.create)
	rg.Put("/bars/<id>", r.update)
	rg.Delete("/bars/<id>", r.delete)
}

func (r *barResource) get(c *routing.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return errors.InvalidParameter("ID must be an unsigned integer")
	}

	response, err := r.service.Get(app.GetRequestScope(c), uint(id))
	if err != nil {
		return errors.NotFound("bar")
	}

	return c.Write(response)
}

func (r *barResource) query(c *routing.Context) error {
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

func (r *barResource) create(c *routing.Context) error {
	var model models.Bar
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

func (r *barResource) update(c *routing.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return errors.InvalidParameter("ID must be an unsigned integer")
	}

	rs := app.GetRequestScope(c)

	model, err := r.service.Get(rs, uint(id))
	if err != nil {
		return errors.NotFound("bar")
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

func (r *barResource) delete(c *routing.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return errors.InvalidParameter("ID must be an unsigned integer")
	}

	response, err := r.service.Delete(app.GetRequestScope(c), uint(id))
	if err != nil {
		return errors.NotFound("bar")
	}

	return c.Write(response)
}
