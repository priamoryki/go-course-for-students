package tests

import (
	"strings"
)

func (s *HTTPSuite) TestCreateAd_EmptyTitle() {
	client := s.Client

	_, err := client.createUser("test", "user")
	s.NoError(err)

	_, err = client.createAd(0, "", "world")
	s.ErrorIs(err, ErrBadRequest)
}

func (s *HTTPSuite) TestCreateAd_TooLongTitle() {
	client := s.Client

	_, err := client.createUser("test", "user")
	s.NoError(err)

	title := strings.Repeat("a", 101)

	_, err = client.createAd(0, title, "world")
	s.ErrorIs(err, ErrBadRequest)
}

func (s *HTTPSuite) TestCreateAd_EmptyText() {
	client := s.Client

	_, err := client.createUser("test", "user")
	s.NoError(err)

	_, err = client.createAd(0, "title", "")
	s.ErrorIs(err, ErrBadRequest)
}

func (s *HTTPSuite) TestCreateAd_TooLongText() {
	client := s.Client

	_, err := client.createUser("test", "user")
	s.NoError(err)

	text := strings.Repeat("a", 501)

	_, err = client.createAd(0, "title", text)
	s.ErrorIs(err, ErrBadRequest)
}

func (s *HTTPSuite) TestUpdateAd_EmptyTitle() {
	client := s.Client

	_, err := client.createUser("test", "user")
	s.NoError(err)

	resp, err := client.createAd(0, "hello", "world")
	s.NoError(err)

	_, err = client.updateAd(0, resp.Data.ID, "", "new_world")
	s.ErrorIs(err, ErrBadRequest)
}

func (s *HTTPSuite) TestUpdateAd_TooLongTitle() {
	client := s.Client

	_, err := client.createUser("test", "user")
	s.NoError(err)

	resp, err := client.createAd(0, "hello", "world")
	s.NoError(err)

	title := strings.Repeat("a", 101)

	_, err = client.updateAd(0, resp.Data.ID, title, "world")
	s.ErrorIs(err, ErrBadRequest)
}

func (s *HTTPSuite) TestUpdateAd_EmptyText() {
	client := s.Client

	_, err := client.createUser("test", "user")
	s.NoError(err)

	resp, err := client.createAd(0, "hello", "world")
	s.NoError(err)

	_, err = client.updateAd(0, resp.Data.ID, "title", "")
	s.ErrorIs(err, ErrBadRequest)
}

func (s *HTTPSuite) TestUpdateAd_TooLongText() {
	client := s.Client

	_, err := client.createUser("test", "user")
	s.NoError(err)

	text := strings.Repeat("a", 501)

	resp, err := client.createAd(0, "hello", "world")
	s.NoError(err)

	_, err = client.updateAd(0, resp.Data.ID, "title", text)
	s.ErrorIs(err, ErrBadRequest)
}
