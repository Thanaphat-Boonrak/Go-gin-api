package routes

import (
	"net/http"
	"strconv"

	"Go-gin-rest-api/models"

	"github.com/gin-gonic/gin"
)

func getEvents(c *gin.Context) {
	events, err := models.GetAllEvents()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Cannot get all events",
			"error":   err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{"events": events})
}

func getEventById(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid event ID"})
		return
	}

	event, err := models.GetEventById(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "Could not find event"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"event": event})
}

func createEvent(c *gin.Context) {
	var event models.Event

	if err := c.ShouldBindJSON(&event); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Cannot parse content json to struct",
			"error":   err.Error(),
		})
		return
	}

	event.UserId = c.GetInt64("userId")

	if err := event.Save(); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Cannot save event",
			"error":   err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "Event created successfully",
		"event":   event,
	})
}

func updateEventById(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid event ID"})
		return
	}

	userId := c.GetInt64("userId")
	event, err := models.GetEventById(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "Could not find event"})
		return
	}

	if event.UserId != userId {
		c.JSON(http.StatusUnauthorized, gin.H{"message": "User does not have permission"})
		return
	}

	var updateEvent models.Event
	if err := c.ShouldBindJSON(&updateEvent); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Cannot parse content json to struct",
			"error":   err.Error(),
		})
		return
	}

	if err := updateEvent.UpdateEventById(id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Cannot update event",
			"error":   err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Successfully updated event",
		"event":   updateEvent,
	})
}

func deleteEvent(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid event ID"})
		return
	}

	userId := c.GetInt64("userId")
	event, err := models.GetEventById(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "Could not find event"})
		return
	}

	if event.UserId != userId {
		c.JSON(http.StatusUnauthorized, gin.H{"message": "User does not have permission"})
		return
	}

	if err := event.DeleteEventById(); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Could not delete event",
			"error":   err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Event deleted successfully",
		"id":      id,
	})
}
