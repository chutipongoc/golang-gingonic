package main

import (
	"strconv"

	"github.com/maxdev/go-gingonic/dto"
	"github.com/maxdev/go-gingonic/entity"
	"github.com/maxdev/go-gingonic/repository"
	"github.com/maxdev/go-gingonic/usecase"

	"time"

	"github.com/gin-gonic/gin"
)

func main() {

	repo := repository.CreateRepository()
	uc := usecase.CreateTodoUsecase(repo)

	server := gin.Default()

	server.GET("/", func(c *gin.Context) {

		var newData dto.GetTodos

		c.ShouldBindJSON(&newData)

		res := uc.GetTodos()

		c.JSON(200, gin.H{
			"data": res,
		})
	})

	server.POST("/", func(c *gin.Context) {

		req := dto.AddTodo{}
		c.ShouldBindJSON(&req)
		data := entity.Todo{
			Content:  req.Content,
			Title:    req.Title,
			CreateAt: time.Now(),
		}
		res, err := uc.AddTodo(&data)

		if err != nil {
			c.JSON(400, gin.H{
				"error": err,
			})

			return
		}

		c.JSON(200, gin.H{
			"res": res,
		})

	})

	server.PUT("/:id", func(c *gin.Context) {
		id, _ := strconv.ParseInt(c.Param("id"), 10, 64)

		var reqUpd dto.UpdateTodo

		c.ShouldBindJSON(&reqUpd)

		data := entity.Todo{
			Id:     id,
			IsDone: reqUpd.IsDone,
		}

		res, _ := uc.UpdateTodo(id, &data)

		c.JSON(200, gin.H{
			"res": res,
		})

	})

	server.DELETE("/:id", func(c *gin.Context) {

		id, _ := strconv.ParseInt(c.Param("id"), 10, 64)

		res, _ := uc.DeleteTodo(id)

		c.JSON(200, gin.H{
			"res": res,
		})

	})

	server.Run(":3001")
}
