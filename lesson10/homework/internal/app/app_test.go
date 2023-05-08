package app

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"homework10/internal/adapters/adrepo"
	"homework10/internal/adapters/userrepo"
	"testing"
)

type SuiteStruct struct {
	suite.Suite
	A App
}

func (s *SuiteStruct) SetupTest() {
	s.A = getApp()
}

func getApp() App {
	adsRepository := adrepo.New()
	userRepository := userrepo.New()
	return NewApp(adsRepository, userRepository)
}

func (s *SuiteStruct) TestCreateUser() {
	a := s.A

	res, err := a.CreateUser("Oleg", "test@gmail.com")
	s.NoError(err, "app.CreateUser")
	s.Equal(int64(0), res.ID)
	s.Equal("Oleg", res.Nickname)
	s.Equal("test@gmail.com", res.Email)
}

func (s *SuiteStruct) TestGetUser() {
	a := s.A

	_, err := a.CreateUser("Oleg", "test@gmail.com")
	s.NoError(err, "app.CreateUser")

	res, err := a.GetUser(0)
	s.NoError(err, "app.GetUser")
	s.Equal(int64(0), res.ID)
	s.Equal("Oleg", res.Nickname)
	s.Equal("test@gmail.com", res.Email)
}

func (s *SuiteStruct) TestUpdateUser() {
	a := s.A

	_, err := a.CreateUser("Oleg", "test@gmail.com")
	s.NoError(err, "app.CreateUser")

	res, err := a.UpdateUser(0, "Oleg1", "test1@gmail.com")
	s.NoError(err, "app.UpdateUser")
	s.Equal(int64(0), res.ID)
	s.Equal("Oleg1", res.Nickname)
	s.Equal("test1@gmail.com", res.Email)
}

func (s *SuiteStruct) TestFindUser() {
	a := s.A

	_, err := a.CreateUser("Oleg", "test@gmail.com")
	s.NoError(err, "app.CreateUser")

	res, err := a.FindUser("Oleg")
	s.NoError(err, "app.FindUser")
	s.Equal(int64(0), res.ID)
	s.Equal("Oleg", res.Nickname)
	s.Equal("test@gmail.com", res.Email)
}

func (s *SuiteStruct) TestDeleteUser() {
	a := s.A

	_, err := a.CreateUser("Oleg", "test@gmail.com")
	s.NoError(err, "app.CreateUser")

	res, err := a.DeleteUser(0)
	s.NoError(err, "app.DeleteUser")
	s.Equal(int64(0), res.ID)
	s.Equal("Oleg", res.Nickname)
	s.Equal("test@gmail.com", res.Email)

	_, err = a.GetUser(0)
	s.Error(err, "app.GetUser")
}

func (s *SuiteStruct) TestListAds() {
	a := s.A

	_, err := a.CreateUser("Oleg", "test@gmail.com")
	s.NoError(err, "app.CreateUser")

	_, err = a.CreateAd("title", "text", 0)
	s.NoError(err, "app.CreateAd")

	res := a.ListAds(0)
	s.NoError(err, "app.ListAds")
	s.Equal(0, len(res))

	_, err = a.ChangeAdStatus(0, 0, true)
	s.NoError(err, "app.ChangeAdStatus")

	res = a.ListAds(0)
	s.NoError(err, "app.ListAds")
	s.Equal(1, len(res))
	s.Equal(int64(0), res[0].ID)
}

func (s *SuiteStruct) TestCreateAd() {
	a := s.A

	_, err := a.CreateUser("Oleg", "test@gmail.com")
	s.NoError(err, "app.CreateUser")

	res, err := a.CreateAd("title", "text", 0)
	s.NoError(err, "app.CreateAd")
	s.Equal(int64(0), res.ID)
	s.Equal("title", res.Title)
	s.Equal("text", res.Text)
	s.Equal(false, res.Published)
}

func (s *SuiteStruct) TestGetAd() {
	a := s.A

	_, err := a.CreateUser("Oleg", "test@gmail.com")
	s.NoError(err, "app.CreateUser")

	_, err = a.CreateAd("title", "text", 0)
	s.NoError(err, "app.CreateAd")

	res, err := a.GetAd(0)
	s.NoError(err, "app.GetAd")
	s.Equal(int64(0), res.ID)
	s.Equal("title", res.Title)
	s.Equal("text", res.Text)
	s.Equal(false, res.Published)
}

func (s *SuiteStruct) TestUpdateAd() {
	a := s.A

	_, err := a.CreateUser("Oleg", "test@gmail.com")
	s.NoError(err, "app.CreateUser")

	_, err = a.CreateAd("title", "text", 0)
	s.NoError(err, "app.CreateAd")

	res, err := a.UpdateAd(0, 0, "title1", "text1")
	s.NoError(err, "app.UpdateAd")
	s.Equal(int64(0), res.ID)
	s.Equal("title1", res.Title)
	s.Equal("text1", res.Text)
	s.Equal(false, res.Published)
}

func (s *SuiteStruct) TestChangeAdStatus() {
	a := s.A

	_, err := a.CreateUser("Oleg", "test@gmail.com")
	s.NoError(err, "app.CreateUser")

	_, err = a.CreateAd("title", "text", 0)
	s.NoError(err, "app.CreateAd")

	res, err := a.ChangeAdStatus(0, 0, true)
	s.NoError(err, "app.UpdateAd")
	s.Equal(int64(0), res.ID)
	s.Equal("title", res.Title)
	s.Equal("text", res.Text)
	s.Equal(true, res.Published)
}

func (s *SuiteStruct) TestFindAd() {
	a := s.A

	_, err := a.CreateUser("Oleg", "test@gmail.com")
	s.NoError(err, "app.CreateUser")

	_, err = a.CreateAd("title", "text", 0)
	s.NoError(err, "app.CreateAd")

	res, err := a.FindAd("title")
	s.NoError(err, "app.FindAd")
	s.Equal(int64(0), res.ID)
	s.Equal("title", res.Title)
	s.Equal("text", res.Text)
	s.Equal(false, res.Published)
}

func (s *SuiteStruct) TestDeleteAd() {
	a := s.A

	_, err := a.CreateUser("Oleg", "test@gmail.com")
	s.NoError(err, "app.CreateUser")

	_, err = a.CreateAd("title", "text", 0)
	s.NoError(err, "app.CreateAd")

	res, err := a.DeleteAd(0, 0)
	s.NoError(err, "app.DeleteAd")
	s.Equal(int64(0), res.ID)
	s.Equal("title", res.Title)
	s.Equal("text", res.Text)
	s.Equal(false, res.Published)

	_, err = a.GetAd(0)
	s.Error(err, "app.GetAd")
}

func BenchmarkListAds(b *testing.B) {
	a := getApp()

	_, err := a.CreateUser("user1", "user1@gmail.com")
	assert.NoError(b, err, "can't create user")
	_, err = a.CreateUser("user2", "user2@gmail.com")
	assert.NoError(b, err, "can't create user")
	for i := int64(0); i < 100; i++ {
		name := fmt.Sprintf("ad%d", i)
		userID := i % 2
		_, err = a.CreateAd(name, name, i%2)
		if i%3 == 0 {
			_, err := a.ChangeAdStatus(i, userID, true)
			assert.NoError(b, err, "can't change ad status")
		}
		assert.NoError(b, err, "can't create ad")
	}

	assert.Equal(b, 100, len(a.ListAds(NonPublished)))

	for i, ad := range a.ListAds(ByAuthor) {
		assert.Equal(b, (3*int64(i))/50, ad.ID%2)
	}

	for i, ad := range a.ListAds(ByCreationTime) {
		assert.Equal(b, 3*int64(i), ad.ID)
	}
}

func TestSuite(t *testing.T) {
	suite.Run(t, new(SuiteStruct))
}
