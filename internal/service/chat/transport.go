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
			"messages": s.FindAll(),
		})
	}
}

func getById(s service) gin.HandlerFunc{
	return func(c *gin.Context) {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"Message": "no se encontró el mensaje",
			})
		}

		result, err := s.FindById(id)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"Message": "no se encontró el mensaje",
			})
		}

		c.JSON(http.StatusOK, gin.H{
			"message": *result,
		})
	}
}

func postMessage(s service) gin.HandlerFunc {

	return func(c *gin.Context) {
		text, err := strconv.Atoi(c.Param("text"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"Message": "no se pudo agregar el mensaje",
			})
		}

		var msg Message

		id, err := s.AddMessage(msg)

		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"Error": "no se pudo agregar el mensaje",
			})
		} else {
			c.JSON(http.StatusOK, gin.H{
				"Message": s.FindById(id),
			})
		}

	}
}

func deleteMessage(s service) gin.HandlerFunc {
	return func(c *gin.Context) {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"Message": "no se encontró el mensaje",
			})
		}

		errD := s.DeleteMsg(id)
		if errD != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"Message": "no se encontró el mensaje",
			})
		}

		c.JSON(http.StatusOK, gin.H{
			"message": "El mensaje fue eliminado",
		})
	}
}

// Register ...
func (s httpService) Register(r *gin.Engine) {
	for _, e := range s.endpoints {
		r.Handle(e.method, e.path, e.function)
	}
}