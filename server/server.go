package server

import (
	"context"

	"github.com/gin-gonic/gin"
	"github.com/imryao/course/pg"
	"github.com/rs/zerolog/log"
)

type Server struct {
	pgClient *pg.Client
	engine   *gin.Engine
}

func NewServer(ctx context.Context) (s *Server, err error) {
	pgClient, err := pg.NewClient(ctx)
	if err != nil {
		return nil, err
	}
	s = &Server{
		pgClient: pgClient,
	}

	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()

	// OAuth2 login handlers
	r.POST("/course/oauth2/callback", s.OAuth2CallbackHandler)

	// Userinfo handlers
	r.GET("/course/userinfo", s.GetUserinfoHandler)

	// Courses handlers
	r.GET("/course/courses", s.ListCoursesHandler)
	r.GET("/course/courses/:id", s.RetrieveCourseHandler)
	r.POST("/course/courses", s.CreateCourseHandler)
	r.PUT("/course/courses/:id", s.UpdateCourseHandler)
	r.DELETE("/course/courses/:id", s.DeleteCourseHandler)

	// User courses handlers
	r.GET("/course/user/courses", s.ListUserCoursesHandler)
	r.GET("/course/user/courses/:id", s.RetrieveUserCourseHandler)
	r.POST("/course/user/courses", s.CreateUserCourseHandler)
	r.DELETE("/course/user/courses/:id", s.DeleteUserCourseHandler)

	s.engine = r
	return s, nil
}

func (s *Server) Run(addr string) (err error) {
	err = s.engine.Run(addr)
	log.Warn().Err(err).Msg("s.engine.Run error")
	return err
}
