package infra

import (
	"net/http"
	"time"

	"github.com/go-playground/validator"
	"github.com/labstack/echo/v4"
)

type (
	// AppCfg application configuration
	// @envconfig (prefix:"APP").
	AppCfg struct {
		Name         string        `default:"Pari" envconfig:"NAME"          required:"true"`
		RESTPort     string        `default:":8089"             envconfig:"ADDRESS"       required:"true"`
		ReadTimeout  time.Duration `default:"5s"                envconfig:"READ_TIMEOUT"`
		WriteTimeout time.Duration `default:"10s"               envconfig:"WRITE_TIMEOUT"`
		Debug        bool          `default:"true"              envconfig:"DEBUG"`
		Timezone     string        `default:"Asia/Jakarta"      envconfig:"TIMEZONE"      required:"true"`
	}
)

// NewEcho return new instance of server
// @ctor.
func NewEcho(cfg *AppCfg) *echo.Echo {
	e := echo.New()
	e.Validator = &CustomValidator{validator: validator.New()}
	e.HideBanner = true
	e.Debug = cfg.Debug

	return e
}

type CustomValidator struct {
	validator *validator.Validate
}

func (cv *CustomValidator) Validate(i interface{}) error {
	if err := cv.validator.Struct(i); err != nil {
		// Optionally, you could return the error to give each route more control over the status code
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return nil
}
