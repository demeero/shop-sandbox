package sql

import (
	"encoding/base64"
	"errors"
	"fmt"
	"strings"
	"time"
)

type pageToken struct {
	CreatedAt time.Time
	UUID      string
}

func (t pageToken) Encode() string {
	key := fmt.Sprintf("%s,%s", t.CreatedAt.Format(time.RFC3339Nano), t.UUID)
	return base64.StdEncoding.EncodeToString([]byte(key))
}

func (t *pageToken) Decode(pageToken string) error {
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
	t.UUID = decodedTokenParts[1]
	return nil
}
