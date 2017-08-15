package apis

import (
	"strconv"

	"github.com/go-ozzo/ozzo-routing"
	"gitlab.com/locatemybeer/lmb-back/app"
	"gitlab.com/locatemybeer/lmb-back/models"
)

type (
	// beerService specifies the interface for the beer service needed by beerResource.
	beerService interface {
		Get(rs app.RequestScope, id int) (*models.Beer, error)
		Query(rs app.RequestScope, offset, limit int) ([]models.Beer, error)
		Count(rs app.RequestScope) (int, error)
		Create(rs app.RequestScope, model *models.Beer) (*models.Beer, error)
		Update(rs app.RequestScope, id int, model *models.Beer) (*models.Beer, error)
		Delete(rs app.RequestScope, id int) (*models.Beer, error)
	}

	// beerResource defines the handlers for the CRUD APIs.
	beerResource struct {
		service beerService
	}
)

// ServeArtist sets up the routing of artist endpoints and the corresponding handlers.
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
		return err
	}

	response, err := r.service.Get(app.GetRequestScope(c), id)
	if err != nil {
		return err
	}

	return c.Write(response)
}

func (r *beerResource) query(c *routing.Context) error {
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

func (r *beerResource) create(c *routing.Context) error {
	var model models.Beer
	if err := c.Read(&model); err != nil {
		return err
	}
	response, err := r.service.Create(app.GetRequestScope(c), &model)
	if err != nil {
		return err
	}

	return c.Write(response)
}

func (r *beerResource) update(c *routing.Context) error {
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

func (r *beerResource) delete(c *routing.Context) error {
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
