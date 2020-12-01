package controllers

import (
	"fmt"
	"net/http"
	"pretty/models"
	"pretty/services"
	"strconv"

	"github.com/gin-gonic/gin"
)

type UserController struct {
	userService services.UserServiceInterface
}

func CreateUserController(userService services.UserServiceInterface) {
	inDB := UserController{userService}
	router := gin.New()

	router.POST("/add-user", inDB.addUser)
	router.GET("/get-users", inDB.getUser)
	router.GET("/get-user/:id", inDB.getUserbyId)
	router.PUT("/update-user/:id", inDB.updateUser)
	router.DELETE("/delete/:id", inDB.deleteUser)

	_ = router.Run(":9988")
}
func (u *UserController) addUser(c *gin.Context) {
	var user models.User

	if user.Nama == "" || user.TanggalLahir == "" || user.NoKtp == 0 || user.PekerjaanId == 0 || user.PendidikanId == 0 {
		c.JSON(http.StatusBadRequest, "Fields cannot be empty")
		return
	}
	err := c.ShouldBindJSON(&user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, "error")
		fmt.Printf("[UserController.addUser] Error when decoder data from body with error : %v\n", err)
		return
	}
	result, err := u.userService.AddUser(&user)
	if err != nil {
		c.JSON(http.StatusBadRequest, "error")
		fmt.Printf("[UserController.addUser] Error when request data to usecase with error: %v\n", err)
		return
	}

	c.JSON(http.StatusOK, result)
}
func (u *UserController) getUser(c *gin.Context) {

	result, err := u.userService.GetAllUser()
	if err != nil {
		c.JSON(http.StatusBadRequest, "opps error")
		fmt.Printf("[UserController.getUser] Error when request data to usecase with error: %v\n", err)
		return
	}

	c.JSON(http.StatusOK, result)
}

func (u *UserController) getUserbyId(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))

	result, err := u.userService.GetUserByID(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, "opps salah")
		fmt.Printf("[UserController.getUser] Error when request data to usecase with error: %v\n", err)
		return
	}

	c.JSON(http.StatusOK, result)
}

func (u *UserController) updateUser(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))

	var user models.User

	err := c.ShouldBindJSON(&user)
	if err != nil {
		c.JSON(http.StatusBadRequest, "opps salah")
		fmt.Printf("[UserController.updateUser] Error when request data to usecase with error: %v\n", err)
		return
	}

	result, err := u.userService.UpdateUser(id, &user)
	if err != nil {
		c.JSON(http.StatusBadRequest, "opps salah")
		fmt.Printf("[UserController.updateUser] Error when request data to usecase with error: %v\n", err)
		return
	}

	c.JSON(http.StatusOK, result)
}

func (u *UserController) deleteUser(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	err := u.userService.DeleteUser(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, "opps salah")
		fmt.Printf("[UserController.deleteUser] Error when request data to usecase with error: %v\n", err)
		return
	}
	c.JSON(http.StatusOK, "Delete data success")
}
