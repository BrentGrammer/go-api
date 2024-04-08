package middleware

import (
	"errors"
	"net/http"

	"github.com/BrentGrammer/goapi/api"
	"github.com/BrentGrammer/goapi/internal/tools"
	log "github.com/sirupsen/logrus"
)

var UnAuthorizedError = errors.New("Invalid username or token.")

1:01:50