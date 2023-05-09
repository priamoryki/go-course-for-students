package tests

func (s *HTTPSuite) TestHTTPCreateUser() {
	client := s.Client

	userResponse, err := client.createUser("test", "user")
	s.NoError(err)
	s.Zero(userResponse.Data.ID)
	s.Equal(userResponse.Data.Nickname, "test")
	s.Equal(userResponse.Data.Email, "user")
}

func (s *HTTPSuite) TestHTTPGetUser() {
	client := s.Client

	_, err := client.getUser(0)
	s.Error(err)

	_, err = client.createUser("test", "user")
	s.NoError(err)

	userResponse, err := client.getUser(0)
	s.NoError(err)
	s.Zero(userResponse.Data.ID)
	s.Equal(userResponse.Data.Nickname, "test")
	s.Equal(userResponse.Data.Email, "user")
}

func (s *HTTPSuite) TestHTTPUpdateUser() {
	client := s.Client

	_, err := client.createUser("test", "user")
	s.NoError(err)

	userResponse, err := client.updateUser(0, "test", "user")
	s.NoError(err)
	s.Zero(userResponse.Data.ID)
	s.Equal(userResponse.Data.Nickname, "test")
	s.Equal(userResponse.Data.Email, "user")
}

func (s *HTTPSuite) TestHTTPFindUser() {
	client := s.Client

	_, err := client.findUser("test")
	s.Error(err)

	_, err = client.createUser("test", "user")
	s.NoError(err)

	userResponse, err := client.findUser("test")
	s.NoError(err)
	s.Zero(userResponse.Data.ID)
	s.Equal(userResponse.Data.Nickname, "test")
	s.Equal(userResponse.Data.Email, "user")
}

func (s *HTTPSuite) TestHTTPDeleteUser() {
	client := s.Client

	_, err := client.deleteUser(0)
	s.Error(err)

	_, err = client.createUser("test", "user")
	s.NoError(err)

	userResponse, err := client.deleteUser(0)
	s.NoError(err)
	s.Zero(userResponse.Data.ID)
	s.Equal(userResponse.Data.Nickname, "test")
	s.Equal(userResponse.Data.Email, "user")
}

func (s *HTTPSuite) TestHTTPListAds() {
	client := s.Client

	_, err := client.createUser("test", "user")
	s.NoError(err)

	response, err := client.createAd(0, "hello", "world")
	s.NoError(err)

	publishedAd, err := client.changeAdStatus(0, response.Data.ID, true)
	s.NoError(err)

	_, err = client.createAd(0, "best cat", "not for sale")
	s.NoError(err)

	ads, err := client.listAds(0)
	s.NoError(err)
	s.Len(ads.Data, 1)
	s.Equal(ads.Data[0].ID, publishedAd.Data.ID)
	s.Equal(ads.Data[0].Title, publishedAd.Data.Title)
	s.Equal(ads.Data[0].Text, publishedAd.Data.Text)
	s.Equal(ads.Data[0].AuthorID, publishedAd.Data.AuthorID)
	s.True(ads.Data[0].Published)
}

func (s *HTTPSuite) TestHTTPCreateAd() {
	client := s.Client

	_, err := client.createAd(0, "hello", "world")
	s.Error(err)

	_, err = client.createUser("test", "user")
	s.NoError(err)

	response, err := client.createAd(0, "hello", "world")
	s.NoError(err)
	s.Zero(response.Data.ID)
	s.Equal(response.Data.Title, "hello")
	s.Equal(response.Data.Text, "world")
	s.Equal(response.Data.AuthorID, int64(0))
	s.False(response.Data.Published)
}

func (s *HTTPSuite) TestHTTPChangeAdStatus() {
	client := s.Client

	_, err := client.createUser("test", "user")
	s.NoError(err)

	_, err = client.changeAdStatus(0, 0, true)
	s.Error(err)

	response, err := client.createAd(0, "hello", "world")
	s.NoError(err)

	response, err = client.changeAdStatus(0, response.Data.ID, true)
	s.NoError(err)
	s.True(response.Data.Published)

	response, err = client.changeAdStatus(0, response.Data.ID, false)
	s.NoError(err)
	s.False(response.Data.Published)

	response, err = client.changeAdStatus(0, response.Data.ID, false)
	s.NoError(err)
	s.False(response.Data.Published)
}

func (s *HTTPSuite) TestHTTPGetAd() {
	client := s.Client

	_, err := client.createUser("test", "user")
	s.NoError(err)

	_, err = client.getAd(0)
	s.Error(err)

	_, err = client.createAd(0, "hello", "world")
	s.NoError(err)

	response, err := client.getAd(0)
	s.NoError(err)
	s.Equal(response.Data.Title, "hello")
	s.Equal(response.Data.Text, "world")
}

func (s *HTTPSuite) TestHTTPUpdateAd() {
	client := s.Client

	_, err := client.createUser("test", "user")
	s.NoError(err)

	_, err = client.updateAd(0, 0, "привет", "мир")
	s.Error(err)

	response, err := client.createAd(0, "hello", "world")
	s.NoError(err)

	response, err = client.updateAd(0, response.Data.ID, "привет", "мир")
	s.NoError(err)
	s.Equal(response.Data.Title, "привет")
	s.Equal(response.Data.Text, "мир")
}

func (s *HTTPSuite) TestHTTPFindAd() {
	client := s.Client

	_, err := client.createUser("test", "user")
	s.NoError(err)

	_, err = client.findAd("hello")
	s.Error(err)

	_, err = client.createAd(0, "hello", "world")
	s.NoError(err)

	response, err := client.findAd("hello")
	s.NoError(err)
	s.Equal(response.Data.Title, "hello")
	s.Equal(response.Data.Text, "world")
}

func (s *HTTPSuite) TestHTTPDeleteAd() {
	client := s.Client

	_, err := client.createUser("test", "user")
	s.NoError(err)

	_, err = client.createAd(0, "hello", "world")
	s.NoError(err)

	response, err := client.deleteAd(0, 0)
	s.NoError(err)
	s.Equal(response.Data.Title, "hello")
	s.Equal(response.Data.Text, "world")
}
