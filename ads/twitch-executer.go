package ads

import (
	"log"
	"net/http"
	"strings"
)

func (this *TwitchRequest) RunAd() {
	if debugMode {
		this.Channel = "ipldev"
	}
	http.Post(this.Url+this.Channel+"/commercial?oauth_token="+this.OAuth, "application/json", strings.NewReader(``))
	log.Println("Fired a Twitch Ad")
}
