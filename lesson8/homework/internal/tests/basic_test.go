package tests

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateUser(t *testing.T) {
	client := getTestClient()

	userResponse, err := client.createUser("test", "user")
	assert.NoError(t, err)
	assert.Zero(t, userResponse.Data.ID)
	assert.Equal(t, userResponse.Data.Nickname, "test")
	assert.Equal(t, userResponse.Data.Email, "user")
}

func TestUpdateUser(t *testing.T) {
	client := getTestClient()

	_, err := client.createUser("test", "user")
	assert.NoError(t, err)

	userResponse, err := client.updateUser(0, "test", "user")
	assert.NoError(t, err)
	assert.Zero(t, userResponse.Data.ID)
	assert.Equal(t, userResponse.Data.Nickname, "test")
	assert.Equal(t, userResponse.Data.Email, "user")
}

func TestFindUser(t *testing.T) {
	client := getTestClient()

	_, err := client.createUser("test", "user")
	assert.NoError(t, err)

	userResponse, err := client.findUser("test")
	assert.NoError(t, err)
	assert.Zero(t, userResponse.Data.ID)
	assert.Equal(t, userResponse.Data.Nickname, "test")
	assert.Equal(t, userResponse.Data.Email, "user")
}

func TestListAds(t *testing.T) {
	client := getTestClient()

	_, err := client.createUser("test", "user")
	assert.NoError(t, err)

	response, err := client.createAd(0, "hello", "world")
	assert.NoError(t, err)

	publishedAd, err := client.changeAdStatus(0, response.Data.ID, true)
	assert.NoError(t, err)

	_, err = client.createAd(0, "best cat", "not for sale")
	assert.NoError(t, err)

	ads, err := client.listAds(0)
	assert.NoError(t, err)
	assert.Len(t, ads.Data, 1)
	assert.Equal(t, ads.Data[0].ID, publishedAd.Data.ID)
	assert.Equal(t, ads.Data[0].Title, publishedAd.Data.Title)
	assert.Equal(t, ads.Data[0].Text, publishedAd.Data.Text)
	assert.Equal(t, ads.Data[0].AuthorID, publishedAd.Data.AuthorID)
	assert.True(t, ads.Data[0].Published)
}

func TestCreateAd(t *testing.T) {
	client := getTestClient()

	_, err := client.createUser("test", "user")
	assert.NoError(t, err)

	response, err := client.createAd(0, "hello", "world")
	assert.NoError(t, err)
	assert.Zero(t, response.Data.ID)
	assert.Equal(t, response.Data.Title, "hello")
	assert.Equal(t, response.Data.Text, "world")
	assert.Equal(t, response.Data.AuthorID, int64(0))
	assert.False(t, response.Data.Published)
}

func TestChangeAdStatus(t *testing.T) {
	client := getTestClient()

	_, err := client.createUser("test", "user")
	assert.NoError(t, err)

	response, err := client.createAd(0, "hello", "world")
	assert.NoError(t, err)

	response, err = client.changeAdStatus(0, response.Data.ID, true)
	assert.NoError(t, err)
	assert.True(t, response.Data.Published)

	response, err = client.changeAdStatus(0, response.Data.ID, false)
	assert.NoError(t, err)
	assert.False(t, response.Data.Published)

	response, err = client.changeAdStatus(0, response.Data.ID, false)
	assert.NoError(t, err)
	assert.False(t, response.Data.Published)
}

func TestGetAd(t *testing.T) {
	client := getTestClient()

	_, err := client.createUser("test", "user")
	assert.NoError(t, err)

	_, err = client.createAd(0, "hello", "world")
	assert.NoError(t, err)

	response, err := client.getAd(0)
	assert.NoError(t, err)
	assert.Equal(t, response.Data.Title, "hello")
	assert.Equal(t, response.Data.Text, "world")
}

func TestUpdateAd(t *testing.T) {
	client := getTestClient()

	_, err := client.createUser("test", "user")
	assert.NoError(t, err)

	response, err := client.createAd(0, "hello", "world")
	assert.NoError(t, err)

	response, err = client.updateAd(0, response.Data.ID, "привет", "мир")
	assert.NoError(t, err)
	assert.Equal(t, response.Data.Title, "привет")
	assert.Equal(t, response.Data.Text, "мир")
}

func TestFindAd(t *testing.T) {
	client := getTestClient()

	_, err := client.createUser("test", "user")
	assert.NoError(t, err)

	_, err = client.createAd(0, "hello", "world")
	assert.NoError(t, err)

	response, err := client.findAd("hello")
	assert.NoError(t, err)
	assert.Equal(t, response.Data.Title, "hello")
	assert.Equal(t, response.Data.Text, "world")
}
