package store

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

// HTTPService ...
type HTTPService interface {
	Register(*gin.Engine)
}

type endpoint struct {
	method   string
	path     string
	function gin.HandlerFunc
}

type httpService struct {
	endpoints []*endpoint
}

// NewHTTPTransport ...
func NewHTTPTransport(s service) HTTPService {
	endpoints := makeEndpoints(s)
	return httpService{endpoints}
}

func makeEndpoints(s service) []*endpoint {
	list := []*endpoint{}

	list = append(list, &endpoint{
		method:   "GET",
		path:     "/clothes",
		function: getAll(s),
	})

	list = append(list, &endpoint{
		method:   "GET",
		path:     "/clothes/:id",
		function: getById(s),
	})

	list = append(list, &endpoint{
		method:   "POST",
		path:     "/clothes/:name",
		function: postItem(s),
	})

	list = append(list, &endpoint{
		method:   "DELETE",
		path:     "/clothes/:id",
		function: deleteItem(s),
	})

	list = append(list, &endpoint{
		method:   "PUT",
		path:     "/clothes/:id/:name",
		function: putItem(s),
	})

	return list
}

func getAll(s service) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"Clothes": s.FindAll(),
		})
	}
}

func getById(s service) gin.HandlerFunc{
	return func(c *gin.Context) {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"Error": "id inválido",
			})
		}

		result, err := s.FindById(id)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"Error": "no se encontró la prenda",
			})
		}

		c.JSON(http.StatusOK, gin.H{
			"Clothing item": *result,
		})
	}
}

func postItem(s service) gin.HandlerFunc {

	return func(c *gin.Context) {
		name := c.Param("name")

		err := s.AddItem(name)

		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"Error": "no se pudo agregar la prenda",
			})
		} else {
			c.JSON(http.StatusOK, gin.H{
				"Clothes": s.FindAll(),
			})
		}

	}
}

func deleteItem(s service) gin.HandlerFunc {
	return func(c *gin.Context) {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"Error": "id inválido",
			})
		}

		errD := s.DeleteItem(id)
		if errD != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"Error": "no se pudo eliminar la prenda",
			})
		}

		c.JSON(http.StatusOK, gin.H{
			"Prenda eliminada": "Ok",
		})
	}
}

func putItem (s service) gin.HandlerFunc {
	return func(c *gin.Context) {
		id, _ := strconv.Atoi(c.Param("id"))
		name := c.Param("name")

		result, err := s.EditItem(id, name)

		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"Error": "no se pudo editar la prenda",
			})
		} else {

			c.JSON(http.StatusOK, gin.H{
				"Clothing item": result,
			})
		}

	}
}

// Register ...
func (s httpService) Register(r *gin.Engine) {
	for _, e := range s.endpoints {
		r.Handle(e.method, e.path, e.function)
	}
}