package handler

import (
	"book-api-go/school"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type schoolHandler struct {
	schoolService school.Service
}

func NewSchoolHandler(schoolService school.Service) *schoolHandler {
	return &schoolHandler{
		schoolService: schoolService,
	}
}

func (h *schoolHandler) GetSchoolsHeader(c *gin.Context) {
	schools, err := h.schoolService.GetSchools()

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	var responseSchool []school.SchoolResponse

	for _, b := range schools {
		SchoolResponse := convertSchoolRequestToSchool(b)

		responseSchool = append(responseSchool, SchoolResponse)
	}

	c.JSON(http.StatusOK, gin.H{
		"data": responseSchool,
	})
}

func (h *schoolHandler) GetSchoolHandler(c *gin.Context) {
	idString := c.Param("id")

	id, _ := strconv.Atoi(idString)

	b, err := h.schoolService.GetSchool(id)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	schoolResponse := convertSchoolRequestToSchool(b)

	c.JSON(http.StatusOK, gin.H{
		"status":  "success",
		"data":    schoolResponse,
		"message": "success mengambil data",
	})
}

func (h *schoolHandler) PostSchoolHandler(c *gin.Context) {
	var schoolRequest school.SchoolRequest

	err := c.ShouldBindJSON(&schoolRequest)

	if err != nil {
		errorMessage := []string{}

		for _, err := range err.(validator.ValidationErrors) {
			errorMessage = append(errorMessage, err.Field()+": "+err.Tag())
		}

		c.JSON(http.StatusBadRequest, gin.H{
			"error": errorMessage,
		})
		return
	}

	school, err := h.schoolService.CreateSchool(schoolRequest)

	SchoolResponse := convertSchoolRequestToSchool(school)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": SchoolResponse,
	})
}

func (h *schoolHandler) DeleteSchoolHandler(c *gin.Context) {
	idString := c.Param("id")

	id, _ := strconv.Atoi(idString)

	_, err := h.schoolService.DeleteSchool(id)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Delete Success",
	})
}

func (h *schoolHandler) UpdateSchoolHandler(c *gin.Context) {
	var schoolRequest school.SchoolRequest

	err := c.ShouldBindJSON(&schoolRequest)

	if err != nil {
		errorMessage := []string{}

		for _, err := range err.(validator.ValidationErrors) {
			errorMessage = append(errorMessage, err.Field()+": "+err.Tag())
		}

		c.JSON(http.StatusBadRequest, gin.H{
			"error": errorMessage,
		})
		return
	}

	idString := c.Param("id")
	id, _ := strconv.Atoi(idString)

	b, err := h.schoolService.UpdateSchool(id, schoolRequest)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	schoolResponse := convertSchoolRequestToSchool(b)

	c.JSON(http.StatusOK, gin.H{
		"data": schoolResponse,
	})
}

func convertSchoolRequestToSchool(b school.School) school.SchoolResponse {
	return school.SchoolResponse{
		ID:      b.ID,
		Name:    b.Name,
		Address: b.Address,
		Class:   b.Class,
		Major:   b.Major,
	}
}
