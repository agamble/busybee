package store

const (
	foodStoresForCity = `SELECT * FROM food_store WHERE city_id = ?`
	transportForCity  = `SELECT * FROM public_transport WHERE city_id = ?`
)

const (
	Train = iota // c0 == 0
	Bus   = iota // c1 == 1
	Tram  = iota // c2 == 2
)

func cityExtras(city *City) error {
	err := db.Select(&(city.FoodStores), foodStoresForCity, city.ID)
	if err != nil {
		return err
	}

	err = db.Select(&(city.Transport), transportForCity, city.ID)
	if err != nil {
		return err
	}

	return nil
}

type FoodStore struct {
	ID         int    `db:"id"`
	Name       string `db:"name"`
	Website    string `db:"website"`
	CanDeliver bool   `db:"delivery"`
	Cost       int    `db:"cost"`
	CityID     int    `db:"city_id"`
}

type PublicTransport struct {
	ID            int    `db:"id"`
	TransportType int    `db:"transport_type"`
	CompanyName   int    `db:"company_name"`
	Cost          int    `db:"cost"`
	Description   string `db:"description"`
	CityID        int    `db:"city_id"`
}
