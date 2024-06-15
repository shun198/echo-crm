package config

import (
	"net/http"
	"os"
	"time"

	"github.com/labstack/echo/v4"
)

func ReadSessionCookie(c echo.Context) (string, error) {
	cookie, err := c.Cookie(cookieName)
	if err != nil {
		return "", err
	}
	return cookie.Value, nil
}

func WriteSessionCookie(c echo.Context, key string) {
	cookie := new(http.Cookie)
	cookie.Name = cookieName
	cookie.Value = key
	cookie.Expires = time.Now().Add(sessionLength)
	cookie.MaxAge = int(maxCookieSessionLength)
	setCookieValuesCommon(cookie)
	c.SetCookie(cookie)
}

func DeleteSessionCookie(c echo.Context) {
	cookie, err := c.Cookie(cookieName)
	if err != nil {
		return
	}
	cookie.MaxAge = -1
	setCookieValuesCommon(cookie)
	c.SetCookie(cookie)
}

func setCookieValuesCommon(cookie *http.Cookie) {
	cookie.HttpOnly = true
	cookie.Path = "/"
	cookie.Secure = secureCookie
	if secureCookie {
		cookie.SameSite = http.SameSiteNoneMode
	} else {
		cookie.Domain = os.Getenv("COOKIE_DOMAIN")
		cookie.SameSite = http.SameSiteStrictMode
	}
}
