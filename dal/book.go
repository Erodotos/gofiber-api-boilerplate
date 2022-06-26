package dal

import (
	"github.com/erodotos/gofiber-api-boilerplate/database"
	"github.com/erodotos/gofiber-api-boilerplate/models"
)

func CreateBook(book *models.Book) error {
	if result := database.DB.Create(book); result.Error != nil {
		return result.Error
	}

	return nil
}

func ReadBook(bookId string) (error, *models.Book) {
	result_book := new(models.Book)

	if result := database.DB.First(&result_book, bookId); result.Error != nil {
		return result.Error, nil
	}

	return nil, result_book
}

func ReadAllBooks() (error, []models.Book) {
	books := []models.Book{}

	if result := database.DB.Find(&books); result.Error != nil {
		return result.Error, nil
	}

	return nil, books
}

func UpdateBook(book *models.Book) error {

	// Save the changes to the database
	if result := database.DB.Save(&book); result.Error != nil {
		return result.Error
	}

	return nil
}

func DeleteBook(bookId string) error {

	if result := database.DB.Delete(&models.Book{}, bookId); result.Error != nil {
		return result.Error
	}

	return nil
}
