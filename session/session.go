package session

import (
	"github.com/go-martini/martini"
	"goBlogExample/utils"
	"net/http"
	"time"
)

type Session struct {
	Id           string
	Username     string
	IsAuthorized bool
}

type Store struct {
	data map[string]*Session
}

var InMemoryStore *Store

func StartSession() {
	InMemoryStore = new(Store)

	InMemoryStore.data = make(map[string]*Session)
}

func (s *Store) Set(session *Session) {
	s.data[session.Id] = session
}

func (s *Store) Get(sessionId string) *Session {
	session := s.data[sessionId]

	if session == nil {
		return &Session{Id: sessionId}
	}

	return session
}

func EnsureCookie(r *http.Request, w http.ResponseWriter) string {
	cookie, err := r.Cookie("sessionId")
	if err == nil {
		return cookie.Value
	}

	sessionId := utils.GenerateId()

	cookie = &http.Cookie{
		Name:    "sessionId",
		Value:   sessionId,
		Expires: time.Now().Add(5 * time.Minute),
	}
	http.SetCookie(w, cookie)

	return sessionId
}

func Middleware(ctx martini.Context, r *http.Request, w http.ResponseWriter) {
	sessionId := EnsureCookie(r, w)
	session := InMemoryStore.Get(sessionId)

	ctx.Map(session)

	ctx.Next()

	InMemoryStore.Set(session)
}
