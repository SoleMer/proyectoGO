package chat

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
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
		path:     "/messages",
		function: getAll(s),
	})

	list = append(list, &endpoint{
		method:   "GET",
		path:     "/messages/:id",
		function: getById(s),
	})

	list = append(list, &endpoint{
		method:   "POST",
		path:     "/messages/:text",
		function: postMessage(s),
	})

	list = append(list, &endpoint{
		method:   "DELETE",
		path:     "/messages/:id",
		function: deleteMessage(s),
	})

	list = append(list, &endpoint{
		method:   "PUT",
		path:     "/messages/:id/:text",
		function: putMessage(s),
	})

	return list
}

func getAll(s service) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"Messages": s.FindAll(),
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
				"Error": "no se encontró el mensaje",
			})
		}

		c.JSON(http.StatusOK, gin.H{
			"Message": *result,
		})
	}
}

func postMessage(s service) gin.HandlerFunc {

	return func(c *gin.Context) {
		text := c.Param("text")

		err := s.AddMessage(text)

		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"Error": "no se pudo agregar el mensaje",
			})
		} else {
			c.JSON(http.StatusOK, gin.H{
				"Messages": s.FindAll(),
			})
		}

	}
}

func deleteMessage(s service) gin.HandlerFunc {
	return func(c *gin.Context) {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"Erro": "id inválido",
			})
		}

		errD := s.DeleteMsg(id)
		if errD != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"Error": "no se pudo eliminar el mensaje",
			})
		}

		c.JSON(http.StatusOK, gin.H{
			"Mensaje eliminao": "Ok",
		})
	}
}

func putMessage (s service) gin.HandlerFunc {
	return func(c *gin.Context) {
		id, _ := strconv.Atoi(c.Param("id"))
		text := c.Param("text")

		result, err := s.EditMessage(id, text)

		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"Error": "no se pudo editar el mensaje",
			})
		} else {

			c.JSON(http.StatusOK, gin.H{
				"Message": result,
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