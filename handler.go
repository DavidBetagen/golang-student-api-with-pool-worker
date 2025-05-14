func (h *StudentHandler) Create(c *fiber.Ctx) error {
	var student domain.Student
	if err := c.BodyParser(&student); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": err.Error()})
	}

	jobID, _ := h.Usecase.CreateAsync(&student)

	return c.JSON(fiber.Map{
		"message": "insert scheduled",
		"job_id":  jobID,
	})
}

func (h *StudentHandler) JobStatus(c *fiber.Ctx) error {
	id := c.Params("id")

	info, ok := h.Usecase.GetJobStatus(id)
	if !ok {
		return c.Status(404).JSON(fiber.Map{"error": "job not found"})
	}

	return c.JSON(info)
}
