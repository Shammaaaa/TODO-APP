package http

import (
	"context"
	"github.com/gofiber/fiber/v3"
	"github.com/shamil/todo-app/internal/application/domain"
	"log"
	"strconv"
)

type TaskUsecase interface {
	CreateTask(ctx context.Context, task domain.Task) (domain.Task, error)
	DeleteTask(ctx context.Context, taskID int) error
	UpdateTask(ctx context.Context, task domain.Task) (domain.Task, error)
	GetAllTasks(ctx context.Context) ([]domain.Task, error)
}

type HandlerImpl struct {
	taskUsecase TaskUsecase
}

func (h *HandlerImpl) CreateTask(ctx fiber.Ctx) error {
	var task domain.Task
	if err := ctx.Bind().Body(&task); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid input",
		})
	}

	// Валидация
	if task.Title == "" {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Title is required",
		})
	}

	createdTask, err := h.taskUsecase.CreateTask(ctx.Context(), task)
	if err != nil {
		log.Printf("Error creating task: %v", err)
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return ctx.Status(fiber.StatusCreated).JSON(createdTask)
}

func (h *HandlerImpl) DeleteTask(ctx fiber.Ctx) error {
	taskID, err := strconv.Atoi(ctx.Params("id"))
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid task ID",
		})
	}

	if err := h.taskUsecase.DeleteTask(ctx.Context(), taskID); err != nil {
		log.Printf("Error deleting task: %v", err)
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return ctx.SendStatus(fiber.StatusNoContent)
}

func (h *HandlerImpl) UpdateTask(ctx fiber.Ctx) error {
	var task domain.Task
	if err := ctx.Bind().Body(&task); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid input",
		})
	}

	taskID, err := strconv.Atoi(ctx.Params("id"))
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid task ID",
		})
	}
	task.ID = taskID

	updatedTask, err := h.taskUsecase.UpdateTask(ctx.Context(), task)
	if err != nil {
		log.Printf("Error updating task: %v", err)
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return ctx.JSON(updatedTask)
}

func (h *HandlerImpl) GetAllTasks(ctx fiber.Ctx) error {
	tasks, err := h.taskUsecase.GetAllTasks(ctx.Context())
	if err != nil {
		log.Printf("Error fetching tasks: %v", err)
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return ctx.JSON(tasks)
}

func New(useCase TaskUsecase) *HandlerImpl {
	return &HandlerImpl{taskUsecase: useCase}
}

func (h *HandlerImpl) MountRoutes(app *fiber.App) {
	app.Post("/tasks", h.CreateTask)
	app.Delete("/tasks/:id", h.DeleteTask)
	app.Put("/tasks/:id", h.UpdateTask)
	app.Get("/tasks", h.GetAllTasks)
}
