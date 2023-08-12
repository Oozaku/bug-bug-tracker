package handler

import (
	"github.com/Oozaku/bug-bug-tracker/backend/database"
	"github.com/Oozaku/bug-bug-tracker/backend/database/models"
	"github.com/gofiber/fiber/v2"
	"golang.org/x/exp/slices"
)

type Issue struct {
	Title       string   `json:"title"`
	Priority    string   `json:"priority"`
	Description string   `json:"description"`
	Progression string   `json:"progression"`
	Author      string   `json:"author"`
	Assignees   []string `json:"assignees"`
}

func CreateIssue(c *fiber.Ctx) error {

	issue := new(Issue)

	if err := c.BodyParser(issue); err != nil {
		c.Status(400).SendString(err.Error())
		return nil
	}

	if _, err := database.GetUser(issue.Author); err != nil {
		c.Status(400).SendString(err.Error())
		return nil
	}

	assignees, err := database.GetUsers(issue.Assignees)

	if err != nil {
		c.Status(400).SendString(err.Error())
		return nil
	}

	if !slices.Contains(models.Priorities, issue.Priority) {
		c.Status(400).SendString("Invalid priority")
	}

	if !slices.Contains(models.Progressions, issue.Progression) {
		c.Status(400).SendString("Invalid progression")
	}

	createdIssue, err := database.CreateIssue(&models.Issue{
		Title:          issue.Title,
		Priority:       issue.Priority,
		Description:    issue.Description,
		Progression:    issue.Progression,
		AuthorUsername: issue.Author,
		Assignees:      assignees,
	})

	if err != nil {
		c.Status(500).SendString("Could not persist data")
	}

	c.JSON(createdIssue)
	return nil
}
