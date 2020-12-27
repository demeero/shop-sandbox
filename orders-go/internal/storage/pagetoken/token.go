package pagetoken

import (
	"encoding/base64"
	"errors"
	"fmt"
	"strings"
	"time"
)

type PageToken struct {
	CreatedAt time.Time
	ID        string
}

func (t PageToken) Encode() string {
	key := fmt.Sprintf("%s,%s", t.CreatedAt.Format(time.RFC3339Nano), t.ID)
	return base64.StdEncoding.EncodeToString([]byte(key))
}

func (t PageToken) Valid() bool {
	return t.ID != "" && !t.CreatedAt.IsZero()
}

func (t *PageToken) Decode(pageToken string) error {
	if pageToken == "" {
		return nil
	}
	decodedToken, err := base64.StdEncoding.DecodeString(pageToken)
	if err != nil {
		return err
	}
	decodedTokenParts := strings.Split(string(decodedToken), ",")
	if len(decodedTokenParts) != 2 {
		return errors.New("pageToken is invalid")
	}
	createdAt, err := time.Parse(time.RFC3339Nano, decodedTokenParts[0])
	if err != nil {
		return err
	}
	t.CreatedAt = createdAt
	t.ID = decodedTokenParts[1]
	return nil
}
