package apis

import (
	"strconv"

	"github.com/go-ozzo/ozzo-routing"
	"gitlab.com/locatemybeer/lmb-back/app"
	"gitlab.com/locatemybeer/lmb-back/models"
)

type (
	// barService specifies the interface for the bar service needed by barResource.
	barService interface {
		Get(rs app.RequestScope, id int) (*models.Bar, error)
		Query(rs app.RequestScope, offset, limit int) ([]models.Bar, error)
		Count(rs app.RequestScope) (int, error)
		Create(rs app.RequestScope, model *models.Bar) (*models.Bar, error)
		Update(rs app.RequestScope, id int, model *models.Bar) (*models.Bar, error)
		Delete(rs app.RequestScope, id int) (*models.Bar, error)
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
		return err
	}

	response, err := r.service.Get(app.GetRequestScope(c), id)
	if err != nil {
		return err
	}

	return c.Write(response)
}

func (r *barResource) query(c *routing.Context) error {
	rs := app.GetRequestScope(c)
	count, err := r.service.Count(rs)
	if err != nil {
		return err
	}
	paginatedList := getPaginatedListFromRequest(c, count)
	items, err := r.service.Query(app.GetRequestScope(c), paginatedList.Offset(), paginatedList.Limit())
	if err != nil {
		return err
	}
	paginatedList.Items = items
	return c.Write(paginatedList)
}

func (r *barResource) create(c *routing.Context) error {
	var model models.Bar
	if err := c.Read(&model); err != nil {
		return err
	}
	response, err := r.service.Create(app.GetRequestScope(c), &model)
	if err != nil {
		return err
	}

	return c.Write(response)
}

func (r *barResource) update(c *routing.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return err
	}

	rs := app.GetRequestScope(c)

	model, err := r.service.Get(rs, id)
	if err != nil {
		return err
	}

	if err := c.Read(model); err != nil {
		return err
	}

	response, err := r.service.Update(rs, id, model)
	if err != nil {
		return err
	}

	return c.Write(response)
}

func (r *barResource) delete(c *routing.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return err
	}

	response, err := r.service.Delete(app.GetRequestScope(c), id)
	if err != nil {
		return err
	}

	return c.Write(response)
}
