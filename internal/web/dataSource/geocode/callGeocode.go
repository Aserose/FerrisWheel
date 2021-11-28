package geocode

func (g *geocodeApi) GetCoordinates(address string) (map[string]interface{}, string) {

	result, err := g.OpenCageData.Geocode(address)
	if err != nil {
		g.logger.Errorf("geocode: error get coordinates: %s", err.Error())
	}

	if result == nil {
		return nil, "address not found"
	}

	coordinates := map[string]interface{}{
		"lat":  result.Lat,
		"long": result.Lng,
	}

	return coordinates, ""
}
