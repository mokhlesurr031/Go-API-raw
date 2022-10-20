package querydir

import (
	"movie_review_apis/conn"
	"movie_review_apis/models"
)

// MovieGetQuery ...
func MovieGetQuery() ([]models.Movie, error) {
	db := conn.GetDB()
	movies := []models.Movie{}
	if err := db.Find(&movies).Error; err != nil {
		return nil, err
	}

	return movies, nil
}

// MovieDetailQuery ...
func MovieDetailQuery(id int) (*models.Movie, error) {
	db := conn.GetDB()
	movie := models.Movie{}
	err := db.First(&movie, "id=?", id).Error
	if err != nil {
		return nil, err
	}
	return &movie, nil
}

func MovieDeleteQuery(id int) (int, error) {
	db := conn.GetDB()
	movie := models.Movie{}
	err := db.Where("id=?", id).Delete(&movie).Error
	if err != nil {
		return 0, err
	}
	return 1, nil
}

func MovieUpdateQuery(id int, movie *models.Movie) int {
	db := conn.GetDB()
	var mv models.Movie

	db.Find(&mv, "id=?", id)
	mv.Name = movie.Name
	mv.Year = movie.Year
	db.Where("id=?", id).Save(&mv)

	//err := db.Model(&mv).Updates(models.Movie{
	//	Name: movie.Name,
	//	Year: movie.Year,
	//})
	//
	//err.Save(&mv)
	//fmt.Println("Update Err", err)
	return 1
}
