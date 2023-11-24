package qso

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/jinzhu/gorm"
)

type QsoHandler struct {
	repository *QsoRepository
}

// GetAll
//	@Summary	Get all QSO records
//	@Tags		QSO
//	@Produce	json
//	@Success	200	{array}	Qso
//	@Router		/api/qso [get]
func (handler *QsoHandler) GetAll(c *fiber.Ctx) error {
	var qsos []Qso = handler.repository.FindAll()
	return c.JSON(qsos)
}

// Get
//	@Summary	Get QSO record by ID
//	@Tags		QSO
//	@Produce	json
//	@Param		id	path		int	true	"QSO ID"
//	@Success	200	{object}	Qso
//	@Failure	404	{object}	object
//	@Router		/api/qso/{id} [get]
func (handler *QsoHandler) Get(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	qso, err := handler.repository.Find(id)

	if err != nil {
		return c.Status(404).JSON(fiber.Map{
			"status": 404,
			"error":  err,
		})
	}

	return c.JSON(qso)
}

// Create
//	@Summary	Create QSO record
//	@Tags		QSO
//	@Accept		json
//	@Produce	json
//	@Param		request	body		Qso	true	"QSO data"
//	@Success	200		{object}	Qso
//	@Failure	400		{object}	object
//	@Failure	500		{object}	object
//	@Router		/api/qso [post]
func (handler *QsoHandler) Create(c *fiber.Ctx) error {
	data := new(Qso)

	if err := c.BodyParser(data); err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Review your input", "error": err})
	}

	item, err := handler.repository.Create(*data)

	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"status":  400,
			"message": "Failed creating item",
			"error":   err,
		})
	}

	return c.JSON(item)
}

// Update
//	@Summary	Update QSO record by ID
//	@Tags		QSO
//	@Accept		json
//	@Produce	json
//	@Param		id		path		int	true	"QSO ID"
//	@Param		request	body		Qso	true	"QSO data"
//	@Success	200		{object}	Qso
//	@Failure	400		{object}	object
//	@Failure	404		{object}	object
//	@Failure	500		{object}	object
//	@Router		/api/qso/{id} [put]
func (handler *QsoHandler) Update(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))

	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"status":  400,
			"message": "Item not found",
			"error":   err,
		})
	}

	qso, err := handler.repository.Find(id)

	if err != nil {
		return c.Status(404).JSON(fiber.Map{
			"message": "Item not found",
		})
	}

	qsoData := new(Qso)

	if err := c.BodyParser(qsoData); err != nil {
		return c.Status(400).JSON(fiber.Map{"status": "error", "message": "Review your input", "data": err})
	}

	qso.Dateon = qsoData.Dateon
	qso.Timeon = qsoData.Timeon
	qso.Callsign = qsoData.Callsign
	qso.Band = qsoData.Band
	qso.Mode = qsoData.Mode
	qso.City = qsoData.City
	qso.State = qsoData.State
	qso.County = qsoData.County
	qso.Country = qsoData.Country
	qso.Qslr = qsoData.Qslr
	qso.Qsls = qsoData.Qsls
	qso.Rstr = qsoData.Rstr
	qso.Rsts = qsoData.Rsts
	qso.Power = qsoData.Power
	qso.Remarks = qsoData.Remarks

	item, err := handler.repository.Save(qso)

	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"message": "Error updating qso",
			"error":   err,
		})
	}

	return c.JSON(item)
}

// Delete
//	@Summary	Delete QSO record by ID
//	@Tags		QSO
//	@Produce	json
//	@Param		id	path	int	true	"QSO ID"
//	@Success	204
//	@Failure	400	{object}	object
//	@Failure	404	{object}	object
//	@Router		/api/qso/{id} [delete]
func (handler *QsoHandler) Delete(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"status":  400,
			"message": "Failed deleting qso",
			"err":     err,
		})
	}
	RowsAffected := handler.repository.Delete(id)
	statusCode := 204
	if RowsAffected == 0 {
		statusCode = 400
	}
	return c.Status(statusCode).JSON(nil)
}

func NewQsoHandler(repository *QsoRepository) *QsoHandler {
	return &QsoHandler{
		repository: repository,
	}
}

func Register(router fiber.Router, database *gorm.DB) {
	database.AutoMigrate(&Qso{})
	qsoRepository := NewQsoRepository(database)
	qsoHandler := NewQsoHandler(qsoRepository)

	movieRouter := router.Group("/qso")
	movieRouter.Get("/", qsoHandler.GetAll)
	movieRouter.Get("/:id", qsoHandler.Get)
	movieRouter.Put("/:id", qsoHandler.Update)
	movieRouter.Post("/", qsoHandler.Create)
	movieRouter.Delete("/:id", qsoHandler.Delete)
}
