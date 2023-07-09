package generator

import "github.com/Dominik48N/url-shorter/url-creator/database"

const (
	linkLengthMin = 3
	linkLengthMax = 12
)

func GenerateRandomLink() (string, error) {
	var link string
	linkExists := true
	var err error

	for linkExists {
		link = GenerateRandomString(linkLengthMin, linkLengthMax)
		linkExists, err = database.CheckLinkExists(link)
		if err != nil {
			return "", err
		}
	}
	return link, nil
}
