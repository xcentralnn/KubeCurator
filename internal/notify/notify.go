package notify

import (
	"log"
	"github.com/xcentralnn/Alertless/internal/model"
)
type Alert struct {
	Labels      map[string]string
	Annotations map[string]string
}

func Send(a model.Alert) {
	log.Printf(
		"ALERT: %s | %s",
		a.Labels["alertname"],
		a.Annotations["summary"],
	)

	// TODO: add Slack / Telegram integration
}