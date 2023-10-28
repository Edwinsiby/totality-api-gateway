package delivery

import (
	"gateway/pkg/client"
	"gateway/pkg/common"
	"gateway/pkg/pb"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Handlers struct {
	UserClient client.UserClient
}

func NewUserHandlers(client client.UserClient) Handlers {
	return Handlers{
		UserClient: client,
	}
}

func (h Handlers) HealthCheck(c *gin.Context) {
	resp, err := h.UserClient.HealthCheck(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": resp.Result})
	return
}

func (h Handlers) UserDetailsByID(c *gin.Context) {
	var input common.UserID
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"error parsing json": err.Error()})
		return
	}

	resp, err := h.UserClient.UserDetails(c, &pb.UserDetailsRequest{UserID: input.UserID})
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusFound, gin.H{"result": resp})
	return
}

func (h Handlers) UserDetailsByList(c *gin.Context) {
	var input common.UserIDList
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"error parsing json": err.Error()})
		return
	}

	resp, err := h.UserClient.UserListDetails(c, &pb.UserListDetailsRequest{UserIDList: input.UserID})
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusFound, gin.H{"User Details": resp.Result, "Not found user ids": resp.NotFound})
	return
}
