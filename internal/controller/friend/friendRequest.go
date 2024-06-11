package friend

import (
	"log"
	"net/http"

	"github.com/MeiramSh/dostap/internal/models"
	"github.com/gin-gonic/gin"
)

type FriendController struct {
	FriendRepository models.FriendRepository
}

func (fc *FriendController) SendFriendRequest(c *gin.Context) {
	var request models.FriendRequest

	if err := c.ShouldBind(&request); err != nil {
		c.JSON(http.StatusBadRequest, models.ErrorResponse{
			Result: []models.ErrorDetail{
				{
					Code:    "ERROR_BIND_JSON",
					Message: "Datas dont match with struct of friend request",
				},
			},
		})
		return
	}
	request.RequestorId = c.GetInt("userID")
	log.Println(request)
	_,err := fc.FriendRepository.CreateFriendRequest(c, &request)
	if err != nil {
		log.Println("--------------------------------",err.Error())
		c.JSON(http.StatusBadRequest, models.ErrorResponse{
			Result: []models.ErrorDetail{
				{
					Code:    "ERROR_CREATE_FRIEND_REQUEST",
					Message: "Couldn't create friend request",
				},
			},
		})
		return
	}
	c.JSON(http.StatusCreated, models.SuccessResponse{Result:"created friend request"})
}
