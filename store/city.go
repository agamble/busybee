package store

import (
	"encoding/json"
	"fmt"
	"log"
	"net/url"
	"regexp"
	"strings"
)

type City struct {
	Base

	ID           int64  `json:"-"`
	Name         string `db:"name"`
	ShortName    string `db:"short_name"`
	ThumbnailURL string `db:"thumbnail_url"`
	Latitude     string `db:"latitude"`
	Longitude    string `db:"longitude"`
	Region       string `db:"region"`
	Country      string `db:"country"`
	Enabled      bool   `db:"enabled",json:"-"`

	FoodStores []FoodStore
	Transport  []PublicTransport
}

const (
	allCitiesQuery  = `SELECT * FROM city WHERE enabled = 1 ORDER BY region, country`
	singleCityQuery = `SELECT * FROM city WHERE short_name=?`
)

func AllCities() ([]City, error) {
	var cities []City
	err := db.Select(&cities, allCitiesQuery)
	if err != nil {
		return nil, err
	}

	return cities, nil
}

func CityByShortName(sn string) (*City, error) {
	c := City{}
	err := db.Get(&c, singleCityQuery, sn)
	if err != nil {
		return nil, fmt.Errorf("Failed to get short city %v: %v", sn, err)
	}

	err = cityExtras(&c)
	if err != nil {
		return nil, err
	}

	return &c, nil
}

func (c *City) People() string {
	return c.Country
}

func (c *City) JSON() (string, error) {
	j, err := json.Marshal(c)
	return string(j), err
}

func (c *City) safeName() string {
	n := c.Name
	reg, err := regexp.Compile(`[^a-zA-Z0-9\s]+`)
	if err != nil {
		log.Fatal(err)
	}
	n = reg.ReplaceAllString(n, "")

	return strings.ToLower(strings.Replace(n, " ", "-", -1))
}

func (c *City) Link() string {
	return (&url.URL{
		Path: fmt.Sprintf("/city/%v", c.safeName()),
	}).String()
}
