package chat

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

type BroChatUserClient struct {
	httpClient *http.Client
	baseUrl    string
}

func NewBroChatClient(httpClient *http.Client, baseUrl string) *BroChatUserClient {
	return &BroChatUserClient{
		httpClient: httpClient,
		baseUrl:    baseUrl,
	}
}

type AuthInfo struct {
	// The JWT token.
	AccessToken string
	// The type of the token. Most likely "Bearer".
	TokenType string
}

// Get User
func (c *BroChatUserClient) GetUser(authInfo *AuthInfo, userId string) (*User, error) {

	base, err := url.Parse(c.baseUrl)

	if err != nil {
		return nil, err
	}

	suffix, err := url.Parse(GET_USER_URL_SUFFIX)

	if err != nil {
		return nil, err
	}

	resolvedUrl := base.ResolveReference(suffix)

	// Create a new request using http
	req, err := http.NewRequest("GET", resolvedUrl.String(), nil)

	if err != nil {
		return nil, err
	}

	// add authorization header to the req
	req.Header.Add("Authorization", authInfo.TokenType+" "+authInfo.AccessToken)

	// Send req using http Client
	res, err := c.httpClient.Do(req)

	if err != nil {
		return nil, err
	}

	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		if res.StatusCode == http.StatusNotFound {
			return nil, errors.New("user not found")
		}

		if res.StatusCode == http.StatusUnauthorized {
			return nil, errors.New("unauthorized")
		}

		if res.StatusCode == http.StatusForbidden {
			return nil, errors.New("forbidden")
		}

		return nil, errors.New("unexpected status code")
	}

	var user User

	err = json.NewDecoder(res.Body).Decode(&user)

	if err != nil {
		return nil, err
	}

	return &user, nil
}

// Get Users
func (c *BroChatUserClient) GetUsers(authInfo *AuthInfo, excludeFriends, excludeSelf bool, page, pageSize int, usernameFilter string) ([]UserInfo, error) {
	base, err := url.Parse(c.baseUrl)

	if err != nil {
		return nil, err
	}

	suffix, err := url.Parse(GET_USERS_URL_SUFFIX)

	if err != nil {
		return nil, err
	}

	resolvedUrl := base.ResolveReference(suffix)

	// Add query params
	q := resolvedUrl.Query()

	q.Add("exclude-friends", strconv.FormatBool(excludeFriends))
	q.Add("exclude-self", strconv.FormatBool(excludeSelf))
	q.Add("page", strconv.Itoa(page))
	q.Add("page-size", strconv.Itoa(pageSize))

	if usernameFilter != "" {
		q.Add("username-filter", usernameFilter)
	}

	resolvedUrl.RawQuery = q.Encode()

	// Create a new request using http
	req, err := http.NewRequest("GET", resolvedUrl.String(), nil)

	if err != nil {
		return nil, err
	}

	// add authorization header to the req
	req.Header.Add("Authorization", authInfo.TokenType+" "+authInfo.AccessToken)

	// Send req using http Client
	res, err := c.httpClient.Do(req)

	if err != nil {
		return nil, err
	}

	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		if res.StatusCode == http.StatusUnauthorized {
			return nil, errors.New("unauthorized")
		}

		if res.StatusCode == http.StatusForbidden {
			return nil, errors.New("forbidden")
		}

		return nil, errors.New("unexpected status code")
	}

	var users = make([]UserInfo, 0)

	err = json.NewDecoder(res.Body).Decode(&users)

	if err != nil {
		return nil, err
	}

	return users, nil
}

// Get Channel
func (c *BroChatUserClient) GetChannelManifest(authInfo *AuthInfo, channelId string) (*Channel, error) {
	base, err := url.Parse(c.baseUrl)

	if err != nil {
		return nil, err
	}

	suffix, err := url.Parse(strings.Replace(GET_CHANNEL_URL_SUFFIX, ":channelId", channelId, 1))

	if err != nil {
		return nil, err
	}

	resolvedUrl := base.ResolveReference(suffix)

	// Create a new request using http
	req, err := http.NewRequest("GET", resolvedUrl.String(), nil)

	if err != nil {
		return nil, err
	}

	// add authorization header to the req
	req.Header.Add("Authorization", authInfo.TokenType+" "+authInfo.AccessToken)

	// Send req using http Client
	res, err := c.httpClient.Do(req)

	if err != nil {
		return nil, err
	}

	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		if res.StatusCode == http.StatusNotFound {
			return nil, errors.New("channel not found")
		} else if res.StatusCode == http.StatusUnauthorized {
			return nil, errors.New("unauthorized")
		} else if res.StatusCode == http.StatusForbidden {
			return nil, errors.New("forbidden")
		}

		return nil, errors.New("unexpected status code")
	}

	var channel Channel

	err = json.NewDecoder(res.Body).Decode(&channel)

	if err != nil {
		return nil, err
	}

	return &channel, nil
}

type GetChannelMessagesOption func(*getChannelMessagesOptions)

type getChannelMessagesOptions struct {
	values map[string]string
}

// An option for the GetChannelMessages method which will pull the messages before the given chat message ID.
func BeforeMessageOption(value string) GetChannelMessagesOption {
	return func(o *getChannelMessagesOptions) {
		o.values["before-msg"] = value
	}
}

// Sets the page option. This will determine which page to start the channel message query from.
func PageOption(page uint64) GetChannelMessagesOption {
	return func(o *getChannelMessagesOptions) {
		o.values["page"] = strconv.FormatUint(page, 10)
	}
}

// Sets the pageSize option. This will determine the size of each page. Anything over 100 will just be set to 100.
func PageSizeOption(pageSize uint64) GetChannelMessagesOption {
	return func(o *getChannelMessagesOptions) {
		o.values["page-size"] = strconv.FormatUint(pageSize, 10)
	}
}

// Get Channel Messages
func (c *BroChatUserClient) GetChannelMessages(authInfo *AuthInfo, channelId string, options ...GetChannelMessagesOption) ([]ChatMessage, error) {
	base, err := url.Parse(c.baseUrl)

	if err != nil {
		return nil, err
	}

	suffix, err := url.Parse(strings.Replace(GET_CHANNEL_MESSAGES_URL_SUFFIX, ":channelId", channelId, 1))

	if err != nil {
		return nil, err
	}

	resolvedUrl := base.ResolveReference(suffix)

	// Default options
	opts := &getChannelMessagesOptions{
		values: make(map[string]string, 0),
	}

	// Apply user-defined options
	for _, opt := range options {
		opt(opts)
	}

	queryParams := resolvedUrl.Query()

	for key, val := range opts.values {
		queryParams.Set(key, val)
	}

	resolvedUrl.RawQuery = queryParams.Encode()

	// Create a new request using http
	req, err := http.NewRequest("GET", resolvedUrl.String(), nil)

	if err != nil {
		return nil, err
	}

	// add authorization header to the req
	req.Header.Add("Authorization", authInfo.TokenType+" "+authInfo.AccessToken)

	// Send req using http Client
	res, err := c.httpClient.Do(req)

	if err != nil {
		return nil, err
	}

	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		if res.StatusCode == http.StatusUnauthorized {
			return nil, errors.New("unauthorized")
		} else if res.StatusCode == http.StatusForbidden {
			return nil, errors.New("forbidden")
		}

		return nil, errors.New("unexpected status code")
	}

	var channels = make([]ChatMessage, 0)

	err = json.NewDecoder(res.Body).Decode(&channels)

	if err != nil {
		return nil, err
	}

	return channels, nil
}

// Send Friend Request
func (c *BroChatUserClient) SendFriendRequest(authInfo *AuthInfo, request *SendFriendRequestRequest) error {
	base, err := url.Parse(c.baseUrl)

	if err != nil {
		return err
	}

	suffix, err := url.Parse(SEND_FRIEND_REQUEST_URL_SUFFIX)

	if err != nil {
		return err
	}

	resolvedUrl := base.ResolveReference(suffix)

	requestBodyBytes, err := json.Marshal(request)

	if err != nil {
		return err
	}

	// Create a new request using http
	req, err := http.NewRequest(http.MethodPut, resolvedUrl.String(), bytes.NewReader(requestBodyBytes))

	if err != nil {
		return err
	}

	// add authorization header to the req
	req.Header.Add("Authorization", authInfo.TokenType+" "+authInfo.AccessToken)

	// Set the content type header
	req.Header.Set("Content-Type", "application/json")

	// Send req using http Client
	res, err := c.httpClient.Do(req)

	if err != nil {
		return err
	}

	defer res.Body.Close()

	if res.StatusCode != http.StatusNoContent {
		if res.StatusCode == http.StatusUnauthorized {
			return errors.New("unauthorized")
		} else if res.StatusCode == http.StatusForbidden {
			return errors.New("forbidden")
		} else if res.StatusCode == http.StatusNotFound {
			return errors.New("user not found")
		} else if res.StatusCode == http.StatusConflict {
			return errors.New("friend request already exists or users are already a friend")
		} else if res.StatusCode == http.StatusBadRequest {
			return errors.New("bad request")
		}

		return errors.New("unexpected status code")
	}

	return nil
}

// Accept Friend Request
func (c *BroChatUserClient) AcceptFriendRequest(authInfo *AuthInfo, request *AcceptFriendRequestRequest) error {
	base, err := url.Parse(c.baseUrl)

	if err != nil {
		return err
	}

	suffix, err := url.Parse(ACCEPT_FRIEND_REQUEST_URL_SUFFIX)

	if err != nil {
		return err
	}

	resolvedUrl := base.ResolveReference(suffix)

	requestBodyBytes, err := json.Marshal(request)

	if err != nil {
		return err
	}

	// Create a new request using http
	req, err := http.NewRequest(http.MethodPut, resolvedUrl.String(), bytes.NewReader(requestBodyBytes))

	if err != nil {
		return err
	}

	// add authorization header to the req
	req.Header.Add("Authorization", authInfo.TokenType+" "+authInfo.AccessToken)

	// Set the content type header
	req.Header.Set("Content-Type", "application/json")

	// Send req using http Client
	res, err := c.httpClient.Do(req)

	if err != nil {
		return err
	}

	defer res.Body.Close()

	if res.StatusCode != http.StatusNoContent {
		switch res.StatusCode {
		case http.StatusUnauthorized:
			return errors.New("unauthorized")
		case http.StatusForbidden:
			return errors.New("forbidden")
		case http.StatusNotFound:
			return errors.New("resource not found")
		case http.StatusBadRequest:
			return errors.New("bad request")
		default:
			return errors.New("unexpected status code")
		}
	}

	return nil
}

// Method for getting (public) rooms that the calling user does not already belong to.
func (c *BroChatUserClient) GetRooms(authInfo *AuthInfo) ([]Room, error) {
	base, err := url.Parse(c.baseUrl)

	if err != nil {
		return nil, err
	}

	suffix, err := url.Parse(GET_ROOMS_URL_SUFFIX)

	if err != nil {
		return nil, err
	}

	resolvedUrl := base.ResolveReference(suffix)

	// Create a new request using http
	req, err := http.NewRequest("GET", resolvedUrl.String(), nil)

	if err != nil {
		return nil, err
	}

	// add authorization header to the req
	req.Header.Set("Authorization", authInfo.TokenType+" "+authInfo.AccessToken)

	// Send req using http Client
	res, err := c.httpClient.Do(req)

	if err != nil {
		return nil, err
	}

	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		if res.StatusCode == http.StatusNotFound {
			return nil, errors.New("channel not found")
		} else if res.StatusCode == http.StatusUnauthorized {
			return nil, errors.New("unauthorized")
		} else if res.StatusCode == http.StatusForbidden {
			return nil, errors.New("forbidden")
		}

		return nil, errors.New("unexpected status code")
	}

	var rooms []Room = make([]Room, 0)

	err = json.NewDecoder(res.Body).Decode(&rooms)

	if err != nil {
		return nil, err
	}

	return rooms, nil
}

// Creates a room. Note: Users cannot create more than 20 rooms.
func (c *BroChatUserClient) CreateRoom(authInfo *AuthInfo, request *CreateRoomRequest) (*Room, error) {
	base, err := url.Parse(c.baseUrl)

	if err != nil {
		return nil, err
	}

	suffix, err := url.Parse(CREATE_ROOM_URL_SUFFIX)

	if err != nil {
		return nil, err
	}

	resolvedUrl := base.ResolveReference(suffix)

	requestBodyBytes, err := json.Marshal(request)

	if err != nil {
		return nil, err
	}

	// Create a new request using http
	req, err := http.NewRequest(http.MethodPost, resolvedUrl.String(), bytes.NewReader(requestBodyBytes))

	if err != nil {
		return nil, err
	}

	// Set authorization header to the req
	req.Header.Set("Authorization", authInfo.TokenType+" "+authInfo.AccessToken)

	// Set the content type header
	req.Header.Set("Content-Type", "application/json")

	// Send req using http Client
	res, err := c.httpClient.Do(req)

	if err != nil {
		return nil, err
	}

	defer res.Body.Close()

	if res.StatusCode != http.StatusCreated {
		if res.StatusCode == http.StatusUnauthorized {
			return nil, errors.New("unauthorized")
		} else if res.StatusCode == http.StatusForbidden {
			return nil, errors.New("forbidden")
		} else if res.StatusCode == http.StatusNotFound {
			return nil, errors.New("user not found")
		} else if res.StatusCode == http.StatusConflict {
			return nil, errors.New("friend request already exists or users are already a friend")
		} else if res.StatusCode == http.StatusBadRequest {
			return nil, errors.New("bad request")
		}

		return nil, errors.New("unexpected status code")
	}

	var room Room = Room{}

	err = json.NewDecoder(res.Body).Decode(&room)

	if err != nil {
		return nil, err
	}

	return &room, nil
}

func (c *BroChatUserClient) JoinRoom(authInfo *AuthInfo, roomId string) error {
	base, err := url.Parse(c.baseUrl)

	if err != nil {
		return err
	}

	suffix, err := url.Parse(strings.Replace(JOIN_ROOM_URL_SUFFIX, ":roomId", roomId, 1))

	if err != nil {
		return err
	}

	resolvedUrl := base.ResolveReference(suffix)

	// Create a new request using http
	req, err := http.NewRequest(http.MethodPut, resolvedUrl.String(), nil)

	if err != nil {
		return err
	}

	// Set authorization header to the req
	req.Header.Set("Authorization", authInfo.TokenType+" "+authInfo.AccessToken)

	// Send req using http Client
	res, err := c.httpClient.Do(req)

	if err != nil {
		return err
	}

	defer res.Body.Close()

	if res.StatusCode != http.StatusNoContent {
		switch res.StatusCode {
		case http.StatusUnauthorized:
			return errors.New("unauthorized")
		case http.StatusForbidden:
			return errors.New("forbidden")
		case http.StatusNotFound:
			return errors.New("resource not found")
		case http.StatusBadRequest:
			return errors.New("bad request")
		default:
			return errors.New("unexpected status code")
		}
	}

	return nil
}
