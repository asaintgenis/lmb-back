package apis

import (
	"strconv"

	"github.com/go-ozzo/ozzo-routing"
	"gitlab.com/locatemybeer/lmb-back/app"
	"gitlab.com/locatemybeer/lmb-back/models"
)

type (
	// beerStyleService specifies the interface for the beerstyle service needed by beerStyleResource.
	beerStyleService interface {
		Get(rs app.RequestScope, id int) (*models.BeerStyle, error)
		Query(rs app.RequestScope, offset, limit int) ([]models.BeerStyle, error)
		Count(rs app.RequestScope) (int, error)
		Create(rs app.RequestScope, model *models.BeerStyle) (*models.BeerStyle, error)
		Update(rs app.RequestScope, id int, model *models.BeerStyle) (*models.BeerStyle, error)
		Delete(rs app.RequestScope, id int) (*models.BeerStyle, error)
	}

	// beerResource defines the handlers for the CRUD APIs.
	beerStyleResource struct {
		service beerStyleService
	}
)

// ServeBeerResource sets up the routing of beer endpoints and the corresponding handlers.
func ServeBeerStyleResource(rg *routing.RouteGroup, service beerStyleService) {
	r := &beerStyleResource{service}
	rg.Get("/beers/<id>", r.get)
	rg.Get("/beers", r.query)
	rg.Post("/beers", r.create)
	rg.Put("/beers/<id>", r.update)
	rg.Delete("/beers/<id>", r.delete)
}

func (r *beerStyleResource) get(c *routing.Context) error {
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

func (r *beerStyleResource) query(c *routing.Context) error {
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

func (r *beerStyleResource) create(c *routing.Context) error {
	var model models.BeerStyle
	if err := c.Read(&model); err != nil {
		return err
	}
	response, err := r.service.Create(app.GetRequestScope(c), &model)
	if err != nil {
		return err
	}

	return c.Write(response)
}

func (r *beerStyleResource) update(c *routing.Context) error {
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

func (r *beerStyleResource) delete(c *routing.Context) error {
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
