package seeders

import (
	"log"

	"github.com/adamwahyuh/coloroutfit/src/database"
	"github.com/adamwahyuh/coloroutfit/src/models"
)

func CategorySeeder() {
	categories := []models.Category{
		{Name: "Minimalist Streetwear"},
		{Name: "Earthy Vintage"},
		{Name: "Pastel Dream"},
		{Name: "Y2K Aesthetic"},
		{Name: "Monochrome Chic"},
		{Name: "Cottagecore"},
		{Name: "Cyberpunk Neon"},
		{Name: "Old Money Luxe"},
		{Name: "Dark Academia"},
		{Name: "Cozy Casual"},
	}

	for _, c := range categories {
		if err := database.DB.FirstOrCreate(&c, models.Category{Name: c.Name}).Error; err != nil {
			log.Print(err)
		}
	}
}
