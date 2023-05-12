package tests

func (s *HTTPSuite) TestChangeStatusAdOfAnotherUser() {
	client := s.Client

	_, err := client.createUser("test", "user")
	s.NoError(err)

	_, err = client.createUser("test", "user")
	s.NoError(err)

	resp, err := client.createAd(0, "hello", "world")
	s.NoError(err)

	_, err = client.changeAdStatus(1, resp.Data.ID, true)
	s.ErrorIs(err, ErrForbidden)
}

func (s *HTTPSuite) TestUpdateAdOfAnotherUser() {
	client := s.Client

	_, err := client.createUser("test", "user")
	s.NoError(err)

	_, err = client.createUser("test", "user")
	s.NoError(err)

	resp, err := client.createAd(0, "hello", "world")
	s.NoError(err)

	_, err = client.updateAd(1, resp.Data.ID, "title", "text")
	s.ErrorIs(err, ErrForbidden)
}

func (s *HTTPSuite) TestCreateAd_ID() {
	client := s.Client

	_, err := client.createUser("test", "user")
	s.NoError(err)

	resp, err := client.createAd(0, "hello", "world")
	s.NoError(err)
	s.Equal(resp.Data.ID, int64(0))

	resp, err = client.createAd(0, "hello", "world")
	s.NoError(err)
	s.Equal(resp.Data.ID, int64(1))

	resp, err = client.createAd(0, "hello", "world")
	s.NoError(err)
	s.Equal(resp.Data.ID, int64(2))
}
