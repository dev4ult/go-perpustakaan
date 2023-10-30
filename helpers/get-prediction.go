package helpers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"perpustakaan/config"
	"perpustakaan/features/feedback/dtos"

	"github.com/labstack/gommon/log"
)

func GetPrediction(comment string) string {
	var highestLabel string = ""
	var highestScore float64

	const APIURL = "https://api-inference.huggingface.co/models/ayameRushia/bert-base-indonesian-1.5G-sentiment-analysis-smsa"

	var body = []byte(fmt.Sprintf(`{"inputs": "%s"}`, comment))

	req, err := http.NewRequest("POST", APIURL, bytes.NewBuffer(body))

	if err != nil {
		log.Error("Error Creating Request")
	} else {
		cfg := config.LoadServerConfig()
		req.Header.Set("Content-Type", "application/json")
		req.Header.Set("Authorization", "Bearer " + cfg.HGF_TOKEN)

		resp, err := http.DefaultClient.Do(req)
		if err != nil {
			log.Error("Error sending request:", err)
		}
		defer resp.Body.Close()

		buf := new(bytes.Buffer)
		_, err = buf.ReadFrom(resp.Body)
		if err != nil {
			log.Error("Error reading response:", err)
		}


		if resp.StatusCode == 200 {
			var predictions [][]dtos.Prediction
			err = json.Unmarshal(buf.Bytes(), &predictions)
			if err != nil {
				log.Error("Error parsing JSON response:", err)
			}

			for _, prediction := range predictions {
				if prediction[0].Score > highestScore {
					highestScore = prediction[0].Score
					highestLabel = prediction[0].Label
				}
			}
		}
	}

	var status = ""

	switch highestLabel {
	case "Positive":
		status = "medium"
	case "Negative":
		status = "high"
	default:
		status = "low"
	}

	return status
}