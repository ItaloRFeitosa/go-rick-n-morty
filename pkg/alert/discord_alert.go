package alert

import "fmt"

const (
	purpleColor = 7811743
	redColor    = 16711680
	yellowColor = 16776960
	greenColor  = 65280
)

var makeDiscordCriticalAlert = newDiscordAlert(purpleColor)
var makeDiscordHighAlert = newDiscordAlert(redColor)
var makeDiscordMediumAlert = newDiscordAlert(yellowColor)
var makeDiscordLowAlert = newDiscordAlert(greenColor)

func newDiscordAlert(color int) func(username string, data Data) DiscordPayload {
	return func(username string, data Data) DiscordPayload {
		return DiscordPayload{
			Username: username,
			Embeds: []Embed{
				{
					Title:       data.Title,
					Description: data.Message,
					Color:       color,
				},
			},
		}
	}
}

type DiscordPayload struct {
	Username string  `json:"username"`
	Embeds   []Embed `json:"embeds"`
}

func (d DiscordPayload) String() string {
	if len(d.Embeds) == 0 {
		return ""
	}

	embed := d.Embeds[0]

	return fmt.Sprintf("%s - %s", embed.Title, embed.Description)
}

type Embed struct {
	Title       string `json:"title"`
	Color       int    `json:"color"`
	Description string `json:"description"`
}
