package config

import (
	"errors"
	"net/http"
	"os"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/shun198/echo-crm/models"
	"gorm.io/gorm"
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

func GetUserFromCookie(c echo.Context, db *gorm.DB) *models.User {
	sc, _ := ReadSessionCookie(c)
	user, _ := GetUserFromSession(sc, c, false, db)
	return user
}

func GetUserFromSession(key string, context echo.Context, refresh bool, db *gorm.DB) (*models.User, error) {
	var session models.Session
	result := db.Preload("User").Where("token = ?", key).Take(&session)
	if result.Error != nil {
		return nil, errors.New("")
	} else {
		if time.Now().Before(session.Expiry) && time.Now().Before(session.MaxExpiry) {
			if refresh {
				db.Model(&session).Update("expiry", time.Now().Add(sessionLength))
			}
			return &session.User, nil
		} else {
			return nil, errors.New("セッションタイムアウトが発生しました\nログインしなおしてください")
		}
	}
}
