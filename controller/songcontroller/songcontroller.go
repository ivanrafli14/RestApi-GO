package songcontroller

import (
	"fmt"

	"github.com/aidarkhanov/nanoid"
	"github.com/gin-gonic/gin"
	"github.com/ivanrafli14/OpenMusicAPI-dicoding/model"
	"gorm.io/gorm"
)

func Index(c *gin.Context){
	var songs []model.Song

	title_qu:= c.Query("title")
	performer_qu:= c.Query("performer")

	title_qu = "%" + title_qu + "%"
	performer_qu = "%" + performer_qu + "%"

	fmt.Println(title_qu, performer_qu)

	model.DB.Where("title LIKE ? AND performer LIKE ?", title_qu, performer_qu).Find(&songs)
	var temp []map[string]interface{}

	
	for _, el := range songs {
		temp = append(temp, map[string]interface{}{
			"id":  el.ID,
			"title": el.Title,
			"performer": el.Performer,
		})
	}

	res := gin.H{
		"songs" : temp, 
	}

	c.JSON(200, gin.H{
		"status": "success",
		"data" : res,
	})

}

func Show(c *gin.Context){
	var song model.Song
	id := c.Param("id")

	if err:= model.DB.First(&song, "id =?", id).Error; err != nil {
		switch err {
		case gorm.ErrRecordNotFound :
			c.AbortWithStatusJSON(404 , gin.H {
				"status" : "fail",
				"message" : "Data tidak ditemukan",
			})
			return
		default :
			c.AbortWithStatusJSON(500, gin.H{
				"status": "fail",
				"message" : err.Error(),
			})
			return
		}
	}

	res := gin.H{
		"song" : song,
	}


	c.JSON(200, gin.H{
		"status" : "success",
		"data" : res,
	})

}

func Create(c *gin.Context){
	var song model.Song
	if err:= c.ShouldBindJSON(&song); err != nil {
		c.AbortWithStatusJSON(400, gin.H {
			"status" : "fail",
			"message" : err.Error(),
		})
		return
	}
	id:= nanoid.New()

	str_id := "song-" +id

	data := model.Song{
		ID: str_id, 
		Title: song.Title, 
		Year: song.Year,
		Performer: song.Performer,
		Genre: song.Genre,
		Duration: song.Duration,
		AlbumId: song.AlbumId,
	}

	model.DB.Create(&data)
	res := gin.H{
		"songId" : str_id,
	}

	c.JSON(201, gin.H {
		"status" : "success",
		"data" : res,
	})

}

func Update(c *gin.Context){
	var song model.Song

	id := c.Param("id")

	if err:= c.ShouldBindJSON(&song); err != nil {
		c.AbortWithStatusJSON(400, gin.H{
			"status" : "fail",
			"message" : err.Error(),
		})
		return
	}

	if model.DB.Model(&song).Where("id =?", id). Updates(&song).RowsAffected == 0 {
		c.AbortWithStatusJSON(404, gin.H{
			"status" : "fail",
			"message" : "Data tidak ditemukan",
		})
		return
	}

	c.JSON(200, gin.H{
		"status" : "success",
		"message" : "Data berhasil diperbaharui",
	})

}

func Delete(c *gin.Context){
	var song model.Song
	id:= c.Param("id")

	if model.DB.Model(&song).Where("id = ?", id).Delete(&song).RowsAffected == 0 {
		c.AbortWithStatusJSON(404,  gin.H{
			"status" : "fail",
			"message" : "Data tidak ditemukan",
		})
		return
	}

	c.JSON(200, gin.H{
		"status" : "success",
		"message" : "Data berhasil dihapus",
	})

}
