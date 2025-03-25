package scrapers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"api/common"
	"api/stock"

	"github.com/jackc/pgx/v5"
)

type Response struct {
	Items    []stock.Stock `json:"items"`
	NextPage string        `json:"next_page"`
}

func getDataFromEndpoint(url string, jwt string, nextPage string) (*Response, error) {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, fmt.Errorf("error creating request: %v", err)
	}

	req.Header.Add("Authorization", "Bearer "+jwt)

	q := req.URL.Query()
	q.Add("next_page", nextPage)
	req.URL.RawQuery = q.Encode()

	client := &http.Client{}
	resp, err := client.Do(req)

	if err != nil {
		return nil, fmt.Errorf("error making request: %v", err)
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("failed call to endpoint: %v", resp.StatusCode)
	}

	var response Response

	err = json.NewDecoder(resp.Body).Decode(&response)

	if err != nil {
		return nil, fmt.Errorf("error decoding response body: %v", err)
	}

	return &response, nil
}

func ScrapeAndStoreData(conn *pgx.Conn) error {
	url := common.GetEnv("INFO_URL")
	jwt := common.GetEnv("INFO_JWT")
	nextPage := GetLastPage(conn)

	log.Printf("start: %v\n", nextPage)
	for {
		response, err := getDataFromEndpoint(url, jwt, nextPage)

		if err != nil {
			log.Fatalf("end")
		}

		for _, item := range response.Items {
			stockData := stock.Stock{
				Ticker:     item.Ticker,
				Company:    item.Company,
				Brokerage:  item.Brokerage,
				Action:     item.Action,
				RatingFrom: item.RatingFrom,
				RatingTo:   item.RatingTo,
				TargetFrom: item.TargetFrom,
				TargetTo:   item.TargetTo,
				Time:       item.Time,
			}

			err = stock.InsertData(conn, stockData)

			if err != nil {
				log.Fatalf("%v", err)
			}
		}

		if response.NextPage == "" {
			break
		}

		nextPage = response.NextPage

		err = SavePageHistory(conn, nextPage)
		if err != nil {
			log.Fatalf("error saving page: %v", err)
		}
	}

	return nil
}
