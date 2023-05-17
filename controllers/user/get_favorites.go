package user

import (
	"backend/controllers"
	"backend/models"
	"strconv"
)

func (this *Controller) GetFavorites() {
	// 获取路由参数
	id_ := this.Ctx.Input.Param(":id")
	id, err := strconv.Atoi(id_)
	if err != nil {

		this.Data["json"] = controllers.MakeResponse(controllers.Err, err.Error(), nil)
		this.ServeJSON()
		return
	}
	favorites, err := models.GetFavorites(id)
	if err != nil {

		this.Data["json"] = controllers.MakeResponse(controllers.Err, err.Error(), nil)
		this.ServeJSON()
		return
	}

	// favorites = models.RefactFavorites(favorites)
	projects := models.FavoriteToProjects(favorites)
	this.Data["json"] = controllers.MakeResponse(controllers.OK, "success", projects)
	this.ServeJSON()

}
