package jwt

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"strings"

	"github.com/gin-gonic/gin"
)

func GetSubFromTokenFromContext(c *gin.Context) (string, error) {
	tokens, ok := c.Request.Header["Authorization"]
	if !ok {
		return "", fmt.Errorf("no authorization header provided")
	}

	token := strings.Split(tokens[0], ".")
	if len(token) != 3 {
		return "", fmt.Errorf("invalid token")
	}

	payload, err := base64.StdEncoding.DecodeString(token[1])
	if err != nil {
		return "", fmt.Errorf("invalid token")
	}

	var payloadMap map[string]interface{}
	err = json.Unmarshal(payload, &payloadMap)
	if err != nil {
		return "", fmt.Errorf("invalid token")
	}

	sub, ok := payloadMap["sub"].(string)
	if !ok {
		return "", fmt.Errorf("invalid token")
	}

	return sub, nil
}
