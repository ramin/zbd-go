package zebedee

import (
	"fmt"
)

type PublicUser struct {
	Name                   string            `json:"name"`
	Image                  string            `json:"image"`
	Username               string            `json:"username"`
	PublicBio              string            `json:"publicBio"`
	PublicStaticCharge     string            `json:"publicStaticCharge"`
	IsPublicPayPageEnabled bool              `json:"isPublicPayPageEnabled"`
	Social                 map[string]string `json:"social"`
}

func GetPublicGamertagData(gamertag string) (*PublicUser, error) {
	client := NewPublicClient("noop")
	return client.GetPublicUser(gamertag)
}

func (c *Client) GetPublicUser(gamertag string) (*PublicUser, error) {
	var pu PublicUser
	endpoint := fmt.Sprintf("/user/%s", gamertag)
	err := c.MakeRequest("GET", endpoint, nil, &pu)
	return &pu, err
}
