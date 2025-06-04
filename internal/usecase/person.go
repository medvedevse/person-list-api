package usecase

import (
	"net/http"

	"github.com/medvedevse/person-list-api/internal/entity"
	"github.com/medvedevse/person-list-api/internal/repository/persistent"
	"github.com/medvedevse/person-list-api/internal/repository/webapi"
	"github.com/medvedevse/person-list-api/pkg/pagination"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

// TODO: Probably it could be written better
type Handler struct {
	H persistent.DBHandler
}

// GetPersonList godoc
// @Summary GetPersonList
// @tags person
// @description get person list
// @ID person-list
// @Accept json
// @Produce json
// @Param age query string false "person search by age"
// @Param gender query string false "person search by gender"
// @Param nationality query string false "person search by nationality"
// @Param page query string false "pagination param"
// @Param limit query string false "pagination param"
// @Success 200 {array} entity.Person
// @Failure 400,404
// @Router /person [get]
func (h Handler) GetPersonList(c *gin.Context) {
	logger := c.MustGet("logger").(*zap.Logger)

	logger.Info("Executing the GetPersonList request")
	queryCount := len(c.Request.URL.Query())

	if queryCount > 0 {
		logger.Debug("GetPersonList query parameters detected", zap.Int("count", queryCount))
		query := h.H.DB.Model(&entity.Person{})

		var (
			personFilters     entity.Filters
			sortedPeople      []entity.Person
			paginationFilters pagination.PaginationFilters
		)

		if err := c.BindQuery(&personFilters); err != nil {
			logger.Error("Error binding filter parameters", zap.Error(err))
			c.AbortWithError(http.StatusBadRequest, err)
			return
		}

		logger.Debug("Filter parameters are binded",
			zap.Int("age", personFilters.Age),
			zap.String("gender", personFilters.Gender),
			zap.String("nationality", personFilters.Nationality),
		)
		if err := c.BindQuery(&paginationFilters); err != nil {
			logger.Error("Error binding pagination parameters", zap.Error(err))
			c.AbortWithError(http.StatusBadRequest, err)
			return
		}
		logger.Debug("Pagination parameters are binded",
			zap.Int("page", paginationFilters.Page),
			zap.Int("limit", paginationFilters.Limit),
		)
		if personFilters.Age > 0 {
			query = query.Where("age = ?", personFilters.Age)
			logger.Debug("Search by query parameter",
				zap.Int("age", personFilters.Age),
			)
		}

		if personFilters.Gender != "" {
			query = query.Where("gender = ?", personFilters.Gender)
			logger.Debug("Search by query parameter",
				zap.String("gender", personFilters.Gender),
			)
		}

		if personFilters.Nationality != "" {
			query = query.Where("nationality = ?", personFilters.Nationality)
			logger.Debug("Search by query parameter",
				zap.String("nationality", personFilters.Nationality),
			)
		}

		if paginationFilters.Limit > 0 && paginationFilters.Page > 0 {
			logger.Debug("Executing a GetPersonList request using pagination",
				zap.Int("limit", paginationFilters.Limit),
				zap.Int("page", paginationFilters.Page),
			)
			if result := query.Scopes(
				pagination.InitPagination(paginationFilters.Limit, paginationFilters.Page).GetPaginatedResult).
				Find(&sortedPeople); result.Error != nil {
				logger.Error("Error executing GetPersonList query with pagination", zap.Error(result.Error))
				c.AbortWithError(http.StatusNotFound, result.Error)
				return
			}
		} else if result := query.Find(&sortedPeople); result.Error != nil {
			logger.Error("Error executing GetPersonList query with filters", zap.Error(result.Error))
			c.AbortWithError(http.StatusNotFound, result.Error)
			return
		}

		logger.Info("The GetPersonList request was successfully completed")
		c.IndentedJSON(http.StatusOK, sortedPeople)
		return
	}
	var people []entity.Person

	if result := h.H.DB.Find(&people); result.Error != nil {
		logger.Error("Error executing GetPersonList request", zap.Error(result.Error))
		c.AbortWithError(http.StatusNotFound, result.Error)
		return
	}

	logger.Info("The GetPersonList request was successfully completed")
	c.IndentedJSON(http.StatusOK, people)
}

// AddPerson godoc
// @Summary AddPerson
// @tags person
// @description add new person
// @ID new-person
// @Accept json
// @Produce json
// @Param new_person_body body entity.NewPersonBody true "New person body"
// @Success 200 {object} entity.Person
// @Failure 400
// @Failure 500
// @Router /person [post]
func (h Handler) AddPerson(c *gin.Context) {
	logger := c.MustGet("logger").(*zap.Logger)

	logger.Info("Executing the AddPerson request")
	var newPerson entity.Person
	body := entity.NewPersonBody{}

	if err := c.BindJSON(&body); err != nil {
		logger.Error("Incorrect body in AddPerson request", zap.Error(err))
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	newPerson.Surname = body.Surname
	newPerson.Name = body.Name
	newPerson.Patronymic = body.Patronymic
	logger.Info("A new entity called newPerson has been created")

	webapi.AddPersonData(logger, &newPerson)

	if result := h.H.DB.Create(&newPerson); result.Error != nil {
		logger.Error("Error creating newPerson in the database", zap.Error(result.Error))
		c.AbortWithError(http.StatusInternalServerError, result.Error)
		return
	}

	logger.Info("The AddPerson request was successfully completed")
	c.IndentedJSON(http.StatusCreated, &newPerson)
}

// DeletePerson godoc
// @Summary DeletePerson
// @tags person
// @description delete person
// @ID delete-person
// @Accept json
// @Produce json
// @Param id path int true "Person ID"
// @Success 200 {integer} integer 1
// @Failure 404
// @Failure 500
// @Router /person/{id} [delete]
func (h Handler) DeletePerson(c *gin.Context) {
	logger := c.MustGet("logger").(*zap.Logger)
	logger.Info("Executing the DeletePerson request")
	id := c.Param("id")
	var deletedPerson entity.Person

	if result := h.H.DB.First(&deletedPerson, id); result.Error != nil {
		logger.Error("Error searching for deletedPerson in the database", zap.Error(result.Error))
		c.AbortWithError(http.StatusNotFound, result.Error)
		return
	}

	if result := h.H.DB.Delete(&deletedPerson); result.Error != nil {
		logger.Error("Error removing deletedPerson from the database", zap.Error(result.Error))
		c.AbortWithError(http.StatusInternalServerError, result.Error)
		return
	}
	logger.Info("The deletedPerson entity has been successfully deleted")
	c.Status(http.StatusOK)
}

// UpdatePerson godoc
// @Summary UpdatePerson
// @tags person
// @description update person
// @ID update-person
// @Accept json
// @Produce json
// @Param id path int true "Person ID"
// @Param new_person_body body entity.NewPersonBody true "New person body"
// @Success 200 {object} entity.Person
// @Failure 400,404
// @Failure 500
// @Router /person/{id} [put]
func (h Handler) UpdatePerson(c *gin.Context) {
	logger := c.MustGet("logger").(*zap.Logger)
	logger.Info("Executing the UpdatePerson request")
	id := c.Param("id")

	body := entity.NewPersonBody{}
	var updatedPerson entity.Person

	if err := c.BindJSON(&body); err != nil {
		logger.Error("Incorrect body in UpdatePerson request", zap.Error(err))
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	if result := h.H.DB.First(&updatedPerson, id); result.Error != nil {
		logger.Error("Error searching for updatedPerson in the database", zap.Error(result.Error))
		c.AbortWithError(http.StatusNotFound, result.Error)
		return
	}

	updatedPerson.Surname = body.Surname
	updatedPerson.Name = body.Name
	updatedPerson.Patronymic = body.Patronymic
	logger.Info("Data updated for the entity updatedPerson")

	webapi.AddPersonData(logger, &updatedPerson)

	if result := h.H.DB.Save(&updatedPerson); result.Error != nil {
		logger.Error("Error saving updatedPerson to database", zap.Error(result.Error))
		c.AbortWithError(http.StatusInternalServerError, result.Error)
	}
	logger.Info("The UpdatePerson request was successfully completed")
	c.IndentedJSON(http.StatusOK, updatedPerson)
}

func PreviewHandler(c *gin.Context) {
	c.String(200,
		`Examples
GET /person - get a full person list
GET /person?age=29&gender=male&nationality=RU - get a sorted person list
GET /person?page=1&limit=3 - get a paginated person list
POST /person - add a new person
PUT /person/:id - update an existing person
DELETE /person/:id - delete an existing person`,
	)
}
