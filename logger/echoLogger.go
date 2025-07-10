package logger

import (
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
)

// EchoLogger middleware para logar as requisições no Echo
func EchoLogger(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		start := time.Now()
		req := c.Request()
		res := c.Response()
		log := Get()

		// Chame o próximo middleware ou handler
		err := next(c)

		// Captura o status da resposta
		status := res.Status

		// Se houver um erro, ajuste o status adequadamente
		if err != nil {
			c.Error(err) // Envia o erro para o Echo processar a resposta

			// Verifica se é um HTTPError e ajusta o status corretamente
			if he, ok := err.(*echo.HTTPError); ok {
				status = he.Code
			} else {
				status = http.StatusInternalServerError
			}
		}

		// Loga as requisições conforme o status HTTP
		switch {
		case status >= 500:
			log.Errorf("Method: %s, URI: %s, Status: %d, Latency: %s, Error: %v",
				req.Method,
				req.RequestURI,
				status,
				time.Since(start),
				err,
			)

		case status >= 400:
			log.Warnf("Method: %s, URI: %s, Status: %d, Latency: %s, Error: %v",
				req.Method,
				req.RequestURI,
				status,
				time.Since(start),
				err,
			)

		default:
			log.Infof("Method: %s, URI: %s, Status: %d, Latency: %s",
				req.Method,
				req.RequestURI,
				status,
				time.Since(start),
			)
		}

		return err
	}
}
