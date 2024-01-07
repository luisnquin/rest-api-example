package locations

import (
	"net/http"
	"strconv"
	"strings"
)

type getCitiesParams struct {
	CountryCode string
	Count       int32
}

func parseGetCitiesParams(r *http.Request) getCitiesParams {
	p := getCitiesParams{
		Count: 1000,
	}

	query := r.URL.Query()

	if c := query.Get("count"); c != "" {
		c, err := strconv.Atoi(c)
		if err == nil && c > 0 && c < 5000 {
			p.Count = int32(c)
		}
	}

	if cc := strings.TrimSpace(query.Get("country_code")); cc != "" {
		p.CountryCode = cc
	}

	return p
}
