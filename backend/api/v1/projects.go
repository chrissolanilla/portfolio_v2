package v1

import (
	"os"
	"fmt"
	"context"

	data "github.com/taheralfayad/portfolio_v2/data"
	utils "github.com/taheralfayad/portfolio_v2/utils"

	"net/http"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"database/sql"
)

func AddProject(c *gin.Context, db *sql.DB, ctx context.Context, client *s3.Client) {
		var payload data.ProjectPayload
		var response data.ProjectResponse
		var err error

		if err = c.ShouldBindJSON(&payload); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}

    id := uuid.New()
		imageID := fmt.Sprintf("%s_image", id.String())

		err = utils.UploadBase64Image(
			ctx,
			client,
			"portfolio-assets",
			imageID,
			payload.Image,
			"image/png",
		)

		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}

		query := `
			INSERT INTO projects (name, description, github_link, image_link, blog_link, type)
			VALUES ($1, $2, $3, $4, $5, $6)
			RETURNING id, created_at
		`

		imageLink := fmt.Sprintf(
			"%s/%s/%s",
			os.Getenv("RUSTFS_ENDPOINT_URL"),
			"portfolio-assets",
			imageID,
		)

		err = db.QueryRow(
			query,
			payload.Name,
			payload.Description,
			payload.GithubLink,
			imageLink,
			payload.BlogLink,
			payload.Type,
		).Scan(&response.ID, &response.CreatedAt)

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": err.Error(),
			})
			return
		}

		response.Name = payload.Name
		response.Description = payload.Description
		response.GithubLink = payload.GithubLink
		response.BlogLink = payload.BlogLink
		response.Type = payload.Type

		c.JSON(http.StatusCreated, response)

}

func GetProjects(c *gin.Context, db *sql.DB) {
	rows, err := db.Query(`
		SELECT
			id,
			name,
			description,
			github_link,
			image_link,
			blog_link,
			type,
			created_at,
			updated_at
		FROM projects;
	`)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	defer rows.Close()

	var projects []data.ProjectResponse

	for rows.Next(){
		var p data.ProjectResponse

		err := rows.Scan(
			&p.ID,
			&p.Name,
			&p.Description,
			&p.GithubLink,
			&p.ImageLink,
			&p.BlogLink,
			&p.Type,
			&p.CreatedAt,
			&p.UpdatedAt,
		)

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": err.Error(),
			})
			return
		}

		projects = append(projects, p)
	}

	c.JSON(http.StatusOK, projects)
}

func EditProject(c *gin.Context, db *sql.DB, ctx context.Context, client *s3.Client) {
	var p data.ProjectPayload
	var err error

	if err = c.ShouldBindJSON(&p); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	var imageS3Key string

	if p.Image != "" {
		var previousImageLink string

		query := (`
			SELECT
				image_link 
			FROM projects
			WHERE id = $1
		`)

		err := db.QueryRow(
			query,
			p.ID,
		).Scan(&previousImageLink)

		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}

		imageS3Key, err = utils.ExtractUUIDFromImageURL(previousImageLink)

		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}

		err = utils.UploadBase64Image(
			ctx,
			client,
			"portfolio-assets",
			imageS3Key,
			p.Image,
			"image/png",
		)

		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}
	}

	query := `
		UPDATE projects
			SET name = $1,
				description = $2,
				github_link = $3,
				blog_link = $4,
				type = $5
		WHERE id = $6
	`

	_, err = db.Exec(
		query,
		p.Name,
		p.Description,
		p.GithubLink,
		p.BlogLink,
		p.Type,
		p.ID,
	)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	var response data.ProjectResponse

	c.JSON(http.StatusCreated, response)
}
