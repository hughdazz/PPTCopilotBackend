package controllers

import "backend/models"

type Response struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func MakeResponse(code int, message string, data interface{}) Response {
	//如果data为project类型，将其转换为ProjectResponse类型
	if data != nil {
		switch data.(type) {
		case models.Project:
			data = models.RefactProjectTime(data.(models.Project))
		case []models.Project:
			projects := data.([]models.Project)
			temp_projects := make([]models.ProjectResponse, len(projects))
			for i, project := range projects {
				temp_projects[i] = models.RefactProjectTime(project)
			}
			data = temp_projects
		case models.User:
			data = models.RefactUserTime(data.(models.User))
		case []models.User:
			users := data.([]models.User)
			temp_users := make([]models.UserResponse, len(users))
			for i, user := range users {
				temp_users[i] = models.RefactUserTime(user)
			}
			data = temp_users
		case models.File:
			data = models.RefactFileTime(data.(models.File))
		case []models.File:
			files := data.([]models.File)
			temp_files := make([]models.FileResponse, len(files))
			for i, file := range files {
				temp_files[i] = models.RefactFileTime(file)
			}
			data = temp_files
		}
	}

	return Response{
		Code:    code,
		Message: message,
		Data:    data,
	}
}
