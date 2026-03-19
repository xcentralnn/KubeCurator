package handler

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/xcentralnn/Alertless/internal/dedup"
	"github.com/xcentralnn/Alertless/internal/notify"
	"github.com/xcentralnn/Alertless/internal/model"
)

type Alert struct {
	Labels      map[string]string `json:"labels"`
	Annotations map[string]string `json:"annotations"`
}

type AlertRequest struct {
	Alerts []model.Alert `json:"alerts"`
}

func HandleWebhook(w http.ResponseWriter, r *http.Request) {
	var req AlertRequest

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	for _, alert := range req.Alerts {
		key := buildKey(alert)

		if dedup.ShouldSend(key) {
			log.Printf("Sending alert: %s", key)
			notify.Send(alert)
		} else {
			log.Printf("Duplicate skipped: %s", key)
		}
	}

	w.WriteHeader(http.StatusOK)
}

func buildKey(a model.Alert) string {
	return a.Labels["alertname"] + "|" +
		a.Labels["instance"] + "|" +
		a.Labels["severity"]
}