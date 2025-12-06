package hotel

import (
	"context"
	"fmt"
	"log/slog"

	"hotel/internal/config"
	"hotel/internal/repository/models"
	"hotel/internal/repository/postgres"
)

type App struct {
	Config *config.Config
}

func (app *App) MustLoad() {
	repo, err := postgres.NewRepository(app.Config)
	if err != nil {
		panic(err.Error())
	}

	desc := "super hotel description"
	testHotel := models.HotelCreate{
		Name:        "Etlon",
		OwnerID:     1,
		Description: &desc,
		Address:     "Zalupkino strict",
		Location: models.Location{
			Longitude: 12.3,
			Latitude:  3.1,
		},
	}
	desc2 := "super hotel 2 description"
	testHotel2 := models.HotelCreate{
		Name:        "Mumumu",
		OwnerID:     1,
		Description: &desc2,
		Address:     "Pupunya strict",
		Location: models.Location{
			Longitude: 2.8,
			Latitude:  9.6,
		},
	}
	ctx := context.Background()
	createdHotel, err := repo.HotelCreate(ctx, &testHotel)
	if err != nil {
		slog.Error("Error create hotel: err - %s\n", err.Error())
	}

	fmt.Println("------Hotel created success------")
	fmt.Printf("id: %s\n", createdHotel.ID.String())
	fmt.Printf("name: %s\n", createdHotel.Name)
	fmt.Printf("owner_id: %d\n", createdHotel.OwnerID)
	fmt.Printf("desc: %s\n", *createdHotel.Description)
	fmt.Printf("address: %s\n", createdHotel.Address)
	fmt.Printf("location: %v\n", createdHotel.Location)
	fmt.Printf("rating: %d\n", createdHotel.Rating)
	fmt.Printf("created: %v\n", createdHotel.CreatedAt)
	fmt.Printf("updated: %v\n", createdHotel.UpdatedAt)
	fmt.Println()

	createdHotel2, err := repo.HotelCreate(ctx, &testHotel2)
	if err != nil {
		slog.Error("Error create hotel: err - %s\n", err.Error())
	}

	fmt.Println("------Hotel 2 created success------")
	fmt.Printf("id: %s\n", createdHotel2.ID.String())
	fmt.Printf("name: %s\n", createdHotel2.Name)
	fmt.Printf("owner_id: %d\n", createdHotel2.OwnerID)
	fmt.Printf("desc: %s\n", *createdHotel2.Description)
	fmt.Printf("address: %s\n", createdHotel2.Address)
	fmt.Printf("location: %v\n", createdHotel2.Location)
	fmt.Printf("rating: %d\n", createdHotel2.Rating)
	fmt.Printf("created: %v\n", createdHotel2.CreatedAt)
	fmt.Printf("updated: %v\n", createdHotel2.UpdatedAt)
	fmt.Println()

	expHotelByID, err := repo.HotelGetByIDOrName(ctx, createdHotel.ID)
	if err != nil {
		fmt.Printf("Error get hotel: err - %s\n", err.Error())
	}

	fmt.Println("------Hotel get by id success------")
	fmt.Printf("id: %s\n", expHotelByID.ID.String())
	fmt.Printf("name: %s\n", expHotelByID.Name)
	fmt.Printf("owner_id: %d\n", expHotelByID.OwnerID)
	fmt.Printf("desc: %s\n", *expHotelByID.Description)
	fmt.Printf("address: %s\n", expHotelByID.Address)
	fmt.Printf("location: %v\n", expHotelByID.Location)
	fmt.Printf("rating: %d\n", expHotelByID.Rating)
	fmt.Printf("created: %v\n", expHotelByID.CreatedAt)
	fmt.Printf("updated: %v\n", expHotelByID.UpdatedAt)
	fmt.Println()

	expHotelByName, err := repo.HotelGetByIDOrName(ctx, createdHotel.Name)
	if err != nil {
		fmt.Printf("Error get hotel: err - %s\n", err.Error())
	}

	fmt.Println("------Hotel get by name success------")
	fmt.Printf("id: %s\n", expHotelByName.ID.String())
	fmt.Printf("name: %s\n", expHotelByName.Name)
	fmt.Printf("owner_id: %d\n", expHotelByName.OwnerID)
	fmt.Printf("desc: %s\n", *expHotelByName.Description)
	fmt.Printf("address: %s\n", expHotelByName.Address)
	fmt.Printf("location: %v\n", expHotelByName.Location)
	fmt.Printf("rating: %d\n", expHotelByName.Rating)
	fmt.Printf("created: %v\n", expHotelByName.CreatedAt)
	fmt.Printf("updated: %v\n", expHotelByName.UpdatedAt)
	fmt.Println()

	hotels, err := repo.HotelGetAll(ctx, 1, 10)
	if err != nil {
		fmt.Printf("Error get hotel: err - %s\n", err.Error())
	}

	for _, h := range hotels {
		fmt.Println("------Hotel get all success------")
		fmt.Printf("id: %s\n", h.ID.String())
		fmt.Printf("name: %s\n", h.Name)
		fmt.Printf("owner_id: %d\n", h.OwnerID)
		fmt.Printf("address: %s\n", h.Address)
		fmt.Printf("location: %v\n", h.Location)
		fmt.Printf("rating: %d\n", h.Rating)
		fmt.Println()
	}
}
