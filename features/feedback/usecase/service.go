package usecase

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"perpustakaan/config"
	"perpustakaan/features/feedback"
	"perpustakaan/features/feedback/dtos"
	"perpustakaan/helpers"

	"github.com/labstack/gommon/log"
	"github.com/mashingan/smapping"
)

type service struct {
	model feedback.Repository
}

func New(model feedback.Repository) feedback.Usecase {
	return &service {
		model: model,
	}
}

func (svc *service) FindAll(page, size int) []dtos.ResFeedback {
	feedbacks := svc.model.Paginate(page, size)

	return feedbacks
}

func (svc *service) FindByID(feedbackID int) *dtos.ResFeedback {
	feedback := svc.model.SelectByID(feedbackID)

	if feedback == nil {
		log.Error("Feedback Not Found!")
		return nil
	}

	return feedback
}

func (svc *service) Create(newFeedback dtos.InputFeedback) *dtos.ResFeedback {
	feedback := feedback.Feedback{}
	
	err := smapping.FillStruct(&feedback, smapping.MapFields(newFeedback))
	if err != nil {
		log.Error(err)
		return nil
	}
	
	feedback.PriorityStatus = "low"
	
	const APIURL = "https://api-inference.huggingface.co/models/ayameRushia/bert-base-indonesian-1.5G-sentiment-analysis-smsa"

	var body = []byte(fmt.Sprintf(`{"inputs": "%s"}`, newFeedback.Comment))

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

		// Read and print the response
		buf := new(bytes.Buffer)
		_, err = buf.ReadFrom(resp.Body)
		if err != nil {
			log.Error("Error reading response:", err)
		}

		// fmt.Println(buf.String())

		if resp.StatusCode == 200 {
			var predictions [][]dtos.Prediction
			err = json.Unmarshal(buf.Bytes(), &predictions)
			if err != nil {
				log.Error("Error parsing JSON response:", err)
			}

			// Find the label with the highest score
			var highestLabel string
			var highestScore float64

			for _, prediction := range predictions {
				if prediction[0].Score > highestScore {
					highestScore = prediction[0].Score
					highestLabel = prediction[0].Label
				}
			}

			feedback.PriorityStatus = helpers.GetPrediction(highestLabel)
		} 
	}

	fb := svc.model.Insert(feedback)

	if fb == nil {
		return nil
	}

	return fb
}

func (svc *service) AddAReply(replyData dtos.InputReply, feedbackID int) *dtos.FeedbackWithReply {
	feedbackReply := feedback.FeedbackReply{}
	
	if err := smapping.FillStruct(&feedbackReply, smapping.MapFields(replyData)); err != nil {
		log.Error(err)
		return nil
	}

	feedbackReply.FeedbackID = feedbackID
	staffReply := svc.model.InsertReplyForAFeedback(feedbackReply)

	if staffReply == nil {
		log.Error("There is No Feedback Updated!")
		return nil
	}

	feedbackWithReply := dtos.FeedbackWithReply{}

	if err := smapping.FillStruct(&feedbackWithReply, smapping.MapFields(replyData)); err != nil {
		log.Error(err)
		return nil
	}

	feedbackWithReply.Reply = *staffReply
	
	return &feedbackWithReply
}

func (svc *service) Remove(feedbackID int) bool {
	rowsAffected := svc.model.DeleteByID(feedbackID)

	if rowsAffected <= 0 {
		log.Error("There is No Feedback Deleted!")
		return false
	}

	return true
}