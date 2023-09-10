package albumcontroller

import (
	

	"github.com/aidarkhanov/nanoid"
	"github.com/gin-gonic/gin"
	"github.com/ivanrafli14/OpenMusicAPI-dicoding/model"
	"gorm.io/gorm"
)

func Show(c *gin.Context){
	var album model.Album
	id := c.Param("id")
	
	
	if err := model.DB.Model(&model.Album{}).Preload("Song").Where("albums.id = ?", id).First(&album).Error; err != nil {
		switch err {
		case gorm.ErrRecordNotFound:
			c.AbortWithStatusJSON(404, gin.H{
				"status" : "fail",
				"message" : "Data tidak ditemukan",
			})
			return
		default:
			c.AbortWithStatusJSON(500, gin.H{
				"status" : "error",
				"message" : "Internal Server Error",
			})
			return
		}
	}
	var songs []model.Song

	model.DB.Where("album_id = ? ", id).Find(&songs)
	// temp := map[string] model.Album {"album": album}
	// res,_ :=  json.Marshal(temp)
	 

	// data["songs"]  {
	// 	album,
	// 	songs,
	// }

	res := gin.H{
		"album" : album,
	 }

	c.JSON(200, gin.H{
		"status": "success",
		"data":res,
	})
}

func Create(c *gin.Context){
	var album model.Album
	// var body struct {
	// 	name string
	// 	year int64
	// }

	//c.Bind(&body)
	//fmt.Println(&album)
	
	id1 := nanoid.New()


	if err := c.ShouldBindJSON(&album); err != nil {
		c.AbortWithStatusJSON(400, gin.H{
			"status" : "fail",
			"message" : err.Error(),
		})
		return
	}
	str := "album-" + id1
	data := model.Album{ID: str, Name: album.Name, Year: album.Year,}
	model.DB.Create(&data)
	
	res := gin.H{
		"albumId" : str,
	}

	c.JSON(201, gin.H {
		"status" : "success",
		"data" : res,
	})
}

func Update(c *gin.Context){
	var album model.Album

	id := c.Param("id")

	if err := c.ShouldBindJSON(&album); err != nil {
		c.AbortWithStatusJSON(400, gin.H{
			"status" : "fail",
			"message" : err.Error(),
		})
		return
	}

	if model.DB.Model(&album).Where("id = ?", id).Updates(&album).RowsAffected == 0 {
		c.AbortWithStatusJSON(404,  gin.H{
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
	var album model.Album
	id:= c.Param("id")

	if model.DB.Model(&album).Where("id = ?", id).Delete(&album).RowsAffected == 0 {
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