package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/nunthapongnp/time-doo-api/internal/models"
	"github.com/nunthapongnp/time-doo-api/internal/services"
)

type SubTaskController struct {
	subTaskService *services.SubTaskService
}

func NewSubTaskController(subTaskService *services.SubTaskService) *SubTaskController {
	return &SubTaskController{subTaskService: subTaskService}
}

func (tc *SubTaskController) CreateSubtask(c *gin.Context) {
	taskID := c.Param("taskId")

	var subtask models.Subtask
	if err := c.ShouldBindJSON(&subtask); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	id, err := tc.subTaskService.CreateSubtask(c, taskID, &subtask)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"id":     id,
		"taskId": taskID,
	})
}

func (tc *SubTaskController) GetSubtask(c *gin.Context) {
	taskID := c.Param("id")
	subtaskID := c.Param("subtaskId")

	subtask, err := tc.subTaskService.GetSubtask(c, taskID, subtaskID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Subtask not found"})
		return
	}

	c.JSON(http.StatusOK, subtask)
}

func (tc *SubTaskController) UpdateSubtask(c *gin.Context) {
	taskID := c.Param("id")
	subtaskID := c.Param("subtaskId")

	var subtask models.Subtask
	if err := c.ShouldBindJSON(&subtask); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := tc.subTaskService.UpdateSubtask(c, taskID, subtaskID, &subtask); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Subtask updated successfully"})
}

func (tc *SubTaskController) DeleteSubtask(c *gin.Context) {
	taskID := c.Param("id")
	subtaskID := c.Param("subtaskId")

	if err := tc.subTaskService.DeleteSubtask(c, taskID, subtaskID); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Subtask deleted successfully"})
}
