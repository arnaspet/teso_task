package http

import (
	"github.com/Sirupsen/logrus"
	"github.com/arnaspet/teso_task/server/domain"
)

func (s *server) InitRoutes(logger *logrus.Logger) {
	connectionPool := domain.NewConnectionPool(logger)
	s.router.GET("/ws/publish", NewWebsocket(logger, connectionPool).PublisherWebsocketHandler)
	s.router.GET("/ws/subscribe", NewWebsocket(logger, connectionPool).SubscriberWebsocketHandler)
}
