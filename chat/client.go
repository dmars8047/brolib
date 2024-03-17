package chat

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

// BroChatClientResult is the result of a requsted operation to the BroChat API via the BroChatClient.
type BroChatClientResult struct {
	// A numeric representation the error code returned by the BroChat API.
	ResponseCode BroChatResponseCode `json:"response_code"`
	// Error details. Will be empty if the response code is a success code.
	ErrorDetails []string `json:"error_details"`
}

// makeBroChatClientResult creates a BroChatClientResult with the given code and message.
func makeBroChatClientResult(code BroChatResponseCode, details ...string) BroChatClientResult {
	return BroChatClientResult{
		ResponseCode: code,
		ErrorDetails: details,
	}
}

// Err returns an error if the response code is an error code. Will return nil if the response code is a success code.
func (c BroChatClientResult) Error() error {
	if c.ResponseCode > BROCHAT_RESPONSE_CODE_SUCCESS {
		return nil
	}

	switch c.ResponseCode {
	case BROCHAT_RESPONSE_CODE_UNHANDLED_ERROR:
		return fmt.Errorf("an unhandled/unexpected error occured")
	case BROCHAT_RESPONSE_CODE_FORBIDDEN_ERROR:
		return fmt.Errorf("forbidden operation")
	case BROCHAT_RESPONSE_CODE_VALIDATION_ERROR:
		return fmt.Errorf("validation error")
	case BROCHAT_RESPONSE_CODE_REQUEST_PARSE_ERROR:
		return fmt.Errorf("request body parsing error")
	case BROCHAT_RESPONSE_CODE_NOT_FOUND_ERROR:
		return fmt.Errorf("resource not found")
	case BROCHAT_RESPONSE_CODE_DATA_CONFLICT_ERROR:
		return fmt.Errorf("data conflict")
	case BROCHAT_RESPONSE_CODE_INVALID_OPERATION:
		return fmt.Errorf("invalid operation")
	case BROCHAT_RESPONSE_CODE_INVALID_HOST_ADDRESS:
		return fmt.Errorf("invalid host address")
	case BROCHAT_RESPONSE_CODE_CONNECTION_TIMEOUT_ERROR:
		return fmt.Errorf("connection timeout")
	case BROCHAT_RESPONSE_CODE_REQUEST_FORMATTING_ERROR:
		return fmt.Errorf("request formatting error")
	case BROCHAT_RESPONSE_CODE_UNEXEPECTED_RESPONSE_ERROR:
		return fmt.Errorf("server response parsing error")
	case BROCHAT_RESPONSE_CODE_GENERIC_REQUEST_ERROR:
		return fmt.Errorf("generic request error")
	case BROCHAT_RESPONSE_CODE_GENERIC_CONNECTION_ERROR:
		return fmt.Errorf("generic connection error")
	default:
		return fmt.Errorf("unknown error")
	}
}

// BroChatError is the response returned when an error
// is encoutered during the processing of a request to the BroChat API.
type BroChatClientContentResult[T any] struct {
	BroChatClientResult
	// The content of the response.
	Content T `json:"content"`
}

func makeBroChatClientContentResult[T any](code BroChatResponseCode, content T, details ...string) BroChatClientContentResult[T] {
	return BroChatClientContentResult[T]{
		BroChatClientResult: makeBroChatClientResult(code, details...),
		Content:             content,
	}
}

// BroChatResponseCode is a numeric representation of the error code returned by the BroChat API.
type BroChatResponseCode uint8

// Server side error codes
const (
	// Indicates an unhandled error.
	BROCHAT_RESPONSE_CODE_UNHANDLED_ERROR BroChatResponseCode = iota
	// Indicates a forbidden operation error. This means the user does not have permission to perform the operation.
	BROCHAT_RESPONSE_CODE_FORBIDDEN_ERROR
	// Indicates a validation error. This means the associated request parameters were invalid.
	BROCHAT_RESPONSE_CODE_VALIDATION_ERROR
	// Indicates a request body parsing error. This means the request body could not be parsed.
	BROCHAT_RESPONSE_CODE_REQUEST_PARSE_ERROR
	// Indicates a not found error. This means the requested resource was not found.
	BROCHAT_RESPONSE_CODE_NOT_FOUND_ERROR
	// Indicates a data conflict error. This means the request could not be completed due to a conflict with the current state of the resource.
	BROCHAT_RESPONSE_CODE_DATA_CONFLICT_ERROR
	// Indicates an invalid operation error. This means the requested operation is invalid. Example: Trying to become friends with yourself.
	BROCHAT_RESPONSE_CODE_INVALID_OPERATION
	// Indicates an unauthorized operation error. This means the user is not authorized to perform the requested operation.
	BROCHAT_RESPONSE_CODE_UNAUTHORIZED_ERROR
)

// Client side error codes
const (
	// Indicates an invalid host address error. This means the address that the client is trying to connect to is invalid.
	BROCHAT_RESPONSE_CODE_INVALID_HOST_ADDRESS = iota + 64
	// Timeout error. This indicates that the BroChat API did not respond in a timely manner.
	BROCHAT_RESPONSE_CODE_CONNECTION_TIMEOUT_ERROR
	// Indicates the request content was not formatted properly.
	BROCHAT_RESPONSE_CODE_REQUEST_FORMATTING_ERROR
	// Indicates that the response from the server was unexpected and could not be parsed.
	BROCHAT_RESPONSE_CODE_UNEXEPECTED_RESPONSE_ERROR
	// Indicates a generic request error.
	BROCHAT_RESPONSE_CODE_GENERIC_REQUEST_ERROR
	// Indicates a generic connection error.
	BROCHAT_RESPONSE_CODE_GENERIC_CONNECTION_ERROR
)

// Success codes
const (
	// Succese code 128 indicates a successful operation.
	BROCHAT_RESPONSE_CODE_SUCCESS BroChatResponseCode = iota + 128
	// Success code 256 indicates a succesful operation with no content.
	BROCHAT_RESPONSE_CODE_NO_CONTENT
)

// BroChatClient is a client for the BroChat API.
type BroChatClient struct {
	httpClient *http.Client
	baseUrl    string
}

// NewBroChatClient creates a new BroChatClient with the given http client and base url.
func NewBroChatClient(httpClient *http.Client, baseUrl string) *BroChatClient {
	return &BroChatClient{
		httpClient: httpClient,
		baseUrl:    baseUrl,
	}
}

// GetUser returns a user by their ID.
func (c *BroChatClient) GetUser(accessToken string, userId string) BroChatClientContentResult[User] {
	url, err := buildUrl(c.baseUrl, GET_USER_URL_SUFFIX)

	if err != nil {
		return makeBroChatClientContentResult(BROCHAT_RESPONSE_CODE_INVALID_HOST_ADDRESS, User{})
	}

	// Create a new request using http
	req, err := http.NewRequest("GET", url, nil)

	if err != nil {
		return makeBroChatClientContentResult(BROCHAT_RESPONSE_CODE_REQUEST_FORMATTING_ERROR, User{})
	}

	// add authorization header to the req
	req.Header.Add("Authorization", fmt.Sprintf("%s %s", defaultTokenType, accessToken))

	// Send req using http Client
	res, err := c.httpClient.Do(req)

	if err != nil {
		return handleHttpRequestErrorWithContent(err, User{})
	}

	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		return handleUnsuccessfulStatusCodeWithContent(res, User{})
	}

	var user User

	err = json.NewDecoder(res.Body).Decode(&user)

	if err != nil {
		return makeBroChatClientContentResult(BROCHAT_RESPONSE_CODE_UNEXEPECTED_RESPONSE_ERROR, User{})
	}

	return makeBroChatClientContentResult(BROCHAT_RESPONSE_CODE_SUCCESS, user)
}

// GetUsersOption is a type for the options that can be passed to the GetUsers method.
type GetUsersOption func(*option)

// An option for the GetUsers method which will exclude the user making the request from the list of users returned.
func GetUsersOption_ExcludeSelf(value bool) GetUsersOption {
	return func(o *option) {
		o.values = append(o.values, queryParam{key: "exclude-self", value: strconv.FormatBool(value)})
	}
}

// An option for the GetUsers method which will exclude the friends of the user making the request from the list of users returned.
func GetUsersOption_ExcludeFriends(value string) GetUsersOption {
	return func(o *option) {
		o.values = append(o.values, queryParam{key: "before-msg", value: value})
	}
}

// An option for the GetUsers method which will filter the list of users returned by the given username.
func GetUsersOption_UsernameFilter(value string) GetUsersOption {
	return func(o *option) {
		o.values = append(o.values, queryParam{key: "username-filter", value: value})
	}
}

// Sets the page option. This will determine which page to start the channel message query from.
func GetUsersOption_Page(page uint64) GetUsersOption {
	return func(o *option) {
		o.values = append(o.values, queryParam{key: "page", value: strconv.FormatUint(page, 10)})
	}
}

// Sets the pageSize option. This will determine the size of each page. Anything over 100 will just be set to 100.
func GetUsersOption_PageSize(pageSize uint64) GetUsersOption {
	return func(o *option) {
		o.values = append(o.values, queryParam{key: "page-size", value: strconv.FormatUint(pageSize, 10)})
	}
}

// GetUsers returns a list of users.
func (c *BroChatClient) GetUsers(accessToken string, options ...GetUsersOption) BroChatClientContentResult[[]UserInfo] {

	// Default options
	opts := option{values: make([]queryParam, 0)}

	// Apply user-defined options
	for _, opt := range options {
		opt(&opts)
	}

	url, err := buildUrl(c.baseUrl, GET_USERS_URL_SUFFIX, opts.values...)

	if err != nil {
		return makeBroChatClientContentResult(BROCHAT_RESPONSE_CODE_INVALID_HOST_ADDRESS, make([]UserInfo, 0))
	}

	// Create a new request using http
	req, err := http.NewRequest("GET", url, nil)

	if err != nil {
		return makeBroChatClientContentResult(BROCHAT_RESPONSE_CODE_REQUEST_FORMATTING_ERROR, make([]UserInfo, 0))
	}

	// add authorization header to the req
	req.Header.Add("Authorization", fmt.Sprintf("%s %s", defaultTokenType, accessToken))

	// Send req using http Client
	res, err := c.httpClient.Do(req)

	if err != nil {
		return handleHttpRequestErrorWithContent(err, make([]UserInfo, 0))
	}

	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		return handleUnsuccessfulStatusCodeWithContent(res, make([]UserInfo, 0))
	}

	var users = make([]UserInfo, 0)

	err = json.NewDecoder(res.Body).Decode(&users)

	if err != nil {
		return makeBroChatClientContentResult(BROCHAT_RESPONSE_CODE_UNEXEPECTED_RESPONSE_ERROR, make([]UserInfo, 0))
	}

	return makeBroChatClientContentResult(BROCHAT_RESPONSE_CODE_SUCCESS, users)
}

// GetChannel returns a channel by its ID.
func (c *BroChatClient) GetChannel(accessToken string, channelId string) BroChatClientContentResult[Channel] {
	url, err := buildUrl(c.baseUrl, strings.Replace(GET_CHANNEL_URL_SUFFIX, ":channelId", channelId, 1))

	if err != nil {
		return makeBroChatClientContentResult(BROCHAT_RESPONSE_CODE_INVALID_HOST_ADDRESS, Channel{})
	}

	// Create a new request using http
	req, err := http.NewRequest("GET", url, nil)

	if err != nil {
		return makeBroChatClientContentResult(BROCHAT_RESPONSE_CODE_REQUEST_FORMATTING_ERROR, Channel{})
	}

	// add authorization header to the req
	req.Header.Add("Authorization", fmt.Sprintf("%s %s", defaultTokenType, accessToken))

	// Send req using http Client
	res, err := c.httpClient.Do(req)

	if err != nil {
		return handleHttpRequestErrorWithContent(err, Channel{})
	}

	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		return handleUnsuccessfulStatusCodeWithContent(res, Channel{})
	}

	var channel Channel

	err = json.NewDecoder(res.Body).Decode(&channel)

	if err != nil {
		return makeBroChatClientContentResult(BROCHAT_RESPONSE_CODE_UNEXEPECTED_RESPONSE_ERROR, Channel{})
	}

	return makeBroChatClientContentResult(BROCHAT_RESPONSE_CODE_SUCCESS, channel)
}

// GetChannelMessagesOption is a type for the options that can be passed to the GetChannelMessages method.
// Example usage: GetChannelMessages_Page(1), GetChannelMessages_PageSize(10)... etc.
type GetChannelMessagesOption func(*option)

// An option for the GetChannelMessages method which will pull the messages before the given chat message ID.
func GetChannelMessages_BeforeMessage(value string) GetChannelMessagesOption {
	return func(o *option) {
		o.values = append(o.values, queryParam{key: "before-msg", value: value})
	}
}

// Sets the page option. This will determine which page to start the channel message query from.
func GetChannelMessages_Page(page uint64) GetChannelMessagesOption {
	return func(o *option) {
		o.values = append(o.values, queryParam{key: "page", value: strconv.FormatUint(page, 10)})
	}
}

// Sets the pageSize option. This will determine the size of each page. Anything over 100 will just be set to 100.
func GetChannelMessages_PageSize(pageSize uint64) GetChannelMessagesOption {
	return func(o *option) {
		o.values = append(o.values, queryParam{key: "page-size", value: strconv.FormatUint(pageSize, 10)})
	}
}

// GetChannelMessages returns a list of messages in a channel.
func (c *BroChatClient) GetChannelMessages(accessToken string, channelId string, options ...GetChannelMessagesOption) BroChatClientContentResult[[]ChatMessage] {
	// Default options
	opts := option{values: make([]queryParam, 0)}

	// Apply user-defined options
	for _, opt := range options {
		opt(&opts)
	}

	url, err := buildUrl(c.baseUrl, strings.Replace(GET_CHANNEL_MESSAGES_URL_SUFFIX, ":channelId", channelId, 1), opts.values...)

	if err != nil {
		return makeBroChatClientContentResult(BROCHAT_RESPONSE_CODE_INVALID_HOST_ADDRESS, make([]ChatMessage, 0))
	}

	// Create a new request using http
	req, err := http.NewRequest("GET", url, nil)

	if err != nil {
		return makeBroChatClientContentResult(BROCHAT_RESPONSE_CODE_REQUEST_FORMATTING_ERROR, make([]ChatMessage, 0))
	}

	// add authorization header to the req
	req.Header.Add("Authorization", fmt.Sprintf("%s %s", defaultTokenType, accessToken))

	// Send req using http Client
	res, err := c.httpClient.Do(req)

	if err != nil {
		return handleHttpRequestErrorWithContent(err, make([]ChatMessage, 0))
	}

	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		return handleUnsuccessfulStatusCodeWithContent(res, make([]ChatMessage, 0))
	}

	var channels = make([]ChatMessage, 0)

	err = json.NewDecoder(res.Body).Decode(&channels)

	if err != nil {
		return makeBroChatClientContentResult(BROCHAT_RESPONSE_CODE_UNEXEPECTED_RESPONSE_ERROR, make([]ChatMessage, 0))
	}

	return makeBroChatClientContentResult(BROCHAT_RESPONSE_CODE_SUCCESS, channels)
}

// SendFriendRequest sends a friend request to a user.
func (c *BroChatClient) SendFriendRequest(accessToken string, request SendFriendRequestRequest) BroChatClientResult {
	url, err := buildUrl(c.baseUrl, SEND_FRIEND_REQUEST_URL_SUFFIX)

	if err != nil {
		return makeBroChatClientResult(BROCHAT_RESPONSE_CODE_INVALID_HOST_ADDRESS)
	}

	requestBodyBytes, err := json.Marshal(request)

	if err != nil {
		return makeBroChatClientResult(BROCHAT_RESPONSE_CODE_REQUEST_FORMATTING_ERROR)
	}

	// Create a new request using http
	req, err := http.NewRequest(http.MethodPut, url, bytes.NewReader(requestBodyBytes))

	if err != nil {
		return makeBroChatClientResult(BROCHAT_RESPONSE_CODE_REQUEST_FORMATTING_ERROR)
	}

	// add authorization header to the req
	req.Header.Set("Authorization", fmt.Sprintf("%s %s", defaultTokenType, accessToken))

	// Set the content type header
	req.Header.Set("Content-Type", "application/json")

	// Send req using http Client
	res, err := c.httpClient.Do(req)

	if err != nil {
		return handleHttpRequestError(err)
	}

	defer res.Body.Close()

	if res.StatusCode != http.StatusNoContent {
		return handleUnsuccessfulStatusCode(res)
	}

	return makeBroChatClientResult(BROCHAT_RESPONSE_CODE_SUCCESS)
}

// AcceptFriendRequest accepts a friend request from a user.
func (c *BroChatClient) AcceptFriendRequest(accessToken string, request AcceptFriendRequestRequest) BroChatClientResult {
	url, err := buildUrl(c.baseUrl, ACCEPT_FRIEND_REQUEST_URL_SUFFIX)

	if err != nil {
		return makeBroChatClientResult(BROCHAT_RESPONSE_CODE_INVALID_HOST_ADDRESS)
	}

	requestBodyBytes, err := json.Marshal(request)

	if err != nil {
		return makeBroChatClientResult(BROCHAT_RESPONSE_CODE_REQUEST_FORMATTING_ERROR)
	}

	// Create a new request using http
	req, err := http.NewRequest(http.MethodPut, url, bytes.NewReader(requestBodyBytes))

	if err != nil {
		return makeBroChatClientResult(BROCHAT_RESPONSE_CODE_REQUEST_FORMATTING_ERROR)
	}

	// add authorization header to the req
	req.Header.Add("Authorization", fmt.Sprintf("%s %s", defaultTokenType, accessToken))

	// Set the content type header
	req.Header.Set("Content-Type", "application/json")

	// Send req using http Client
	res, err := c.httpClient.Do(req)

	if err != nil {
		return handleHttpRequestError(err)
	}

	defer res.Body.Close()

	if res.StatusCode != http.StatusNoContent {
		return handleUnsuccessfulStatusCode(res)
	}

	return makeBroChatClientResult(BROCHAT_RESPONSE_CODE_SUCCESS)
}

// GetRooms returns a list of rooms.
func (c *BroChatClient) GetRooms(accessToken string) BroChatClientContentResult[[]Room] {
	url, err := buildUrl(c.baseUrl, GET_ROOMS_URL_SUFFIX)

	if err != nil {
		return makeBroChatClientContentResult(BROCHAT_RESPONSE_CODE_INVALID_HOST_ADDRESS, make([]Room, 0))
	}

	// Create a new request using http
	req, err := http.NewRequest("GET", url, nil)

	if err != nil {
		return makeBroChatClientContentResult(BROCHAT_RESPONSE_CODE_REQUEST_FORMATTING_ERROR, make([]Room, 0))
	}

	// add authorization header to the req
	req.Header.Set("Authorization", fmt.Sprintf("%s %s", defaultTokenType, accessToken))

	// Send req using http Client
	res, err := c.httpClient.Do(req)

	if err != nil {
		return handleHttpRequestErrorWithContent(err, make([]Room, 0))
	}

	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		return handleUnsuccessfulStatusCodeWithContent(res, make([]Room, 0))
	}

	var rooms []Room = make([]Room, 0)

	err = json.NewDecoder(res.Body).Decode(&rooms)

	if err != nil {
		return makeBroChatClientContentResult(BROCHAT_RESPONSE_CODE_UNEXEPECTED_RESPONSE_ERROR, make([]Room, 0))
	}

	return makeBroChatClientContentResult(BROCHAT_RESPONSE_CODE_SUCCESS, rooms)
}

// CreateRoom creates a new room. Note: The user cannot create more than 20 rooms.
func (c *BroChatClient) CreateRoom(accessToken string, request CreateRoomRequest) BroChatClientContentResult[Room] {
	url, err := buildUrl(c.baseUrl, CREATE_ROOM_URL_SUFFIX)

	if err != nil {
		return makeBroChatClientContentResult(BROCHAT_RESPONSE_CODE_INVALID_HOST_ADDRESS, Room{})
	}

	requestBodyBytes, err := json.Marshal(request)

	if err != nil {
		return makeBroChatClientContentResult(BROCHAT_RESPONSE_CODE_REQUEST_FORMATTING_ERROR, Room{})
	}

	// Create a new request using http
	req, err := http.NewRequest(http.MethodPost, url, bytes.NewReader(requestBodyBytes))

	if err != nil {
		return makeBroChatClientContentResult(BROCHAT_RESPONSE_CODE_REQUEST_FORMATTING_ERROR, Room{})
	}

	// Set authorization header to the req
	req.Header.Set("Authorization", fmt.Sprintf("%s %s", defaultTokenType, accessToken))

	// Set the content type header
	req.Header.Set("Content-Type", "application/json")

	// Send req using http Client
	res, err := c.httpClient.Do(req)

	if err != nil {
		return handleHttpRequestErrorWithContent(err, Room{})
	}

	defer res.Body.Close()

	if res.StatusCode != http.StatusCreated {
		return handleUnsuccessfulStatusCodeWithContent(res, Room{})
	}

	var room Room = Room{}

	err = json.NewDecoder(res.Body).Decode(&room)

	if err != nil {
		return makeBroChatClientContentResult(BROCHAT_RESPONSE_CODE_UNEXEPECTED_RESPONSE_ERROR, Room{})
	}

	return makeBroChatClientContentResult(BROCHAT_RESPONSE_CODE_SUCCESS, room)
}

// JoinRoom joins a user to a room.
func (c *BroChatClient) JoinRoom(accessToken string, roomId string) BroChatClientResult {
	url, err := buildUrl(c.baseUrl, strings.Replace(JOIN_ROOM_URL_SUFFIX, ":roomId", roomId, 1))

	if err != nil {
		return makeBroChatClientResult(BROCHAT_RESPONSE_CODE_INVALID_HOST_ADDRESS)
	}

	// Create a new request using http
	req, err := http.NewRequest(http.MethodPut, url, nil)

	if err != nil {
		return makeBroChatClientResult(BROCHAT_RESPONSE_CODE_REQUEST_FORMATTING_ERROR)
	}

	// Set authorization header to the req
	req.Header.Set("Authorization", fmt.Sprintf("%s %s", defaultTokenType, accessToken))

	// Send req using http Client
	res, err := c.httpClient.Do(req)

	if err != nil {
		if err, ok := err.(net.Error); ok && err.Timeout() {
			// If it was a timeout error
			return makeBroChatClientResult(BROCHAT_RESPONSE_CODE_CONNECTION_TIMEOUT_ERROR)
		}

		return makeBroChatClientResult(BROCHAT_RESPONSE_CODE_GENERIC_CONNECTION_ERROR)
	}

	defer res.Body.Close()

	if res.StatusCode != http.StatusNoContent {
		return handleUnsuccessfulStatusCode(res)
	}

	return makeBroChatClientResult(BROCHAT_RESPONSE_CODE_SUCCESS)
}

// option is a type for the options that can be passed to the GetChannelMessages method.
type option struct {
	values []queryParam
}

// The default token type used for authorization.
const defaultTokenType = "Bearer"

// Struct for query parameters
type queryParam struct {
	key   string
	value string
}

// buildUrl is a helper function that builds a url from a base url, a suffix and query parameters.
func buildUrl(baseUrl, suffix string, queryParams ...queryParam) (string, error) {
	base, err := url.Parse(baseUrl)

	if err != nil {
		return "", err
	}

	suffixUrl, err := url.Parse(suffix)

	if err != nil {
		return "", err
	}

	resolvedUrl := base.ResolveReference(suffixUrl)

	if len(queryParams) > 0 {
		q := resolvedUrl.Query()

		for _, param := range queryParams {
			if param.key != "" && param.value != "" {
				q.Set(param.key, param.value)
			}
		}

		resolvedUrl.RawQuery = q.Encode()
	}

	return resolvedUrl.String(), nil
}

// handleHttpRequestErrorWithContent creates a BroChatClientContentResult generated from an error after attempting an http request.
func handleHttpRequestErrorWithContent[T any](err error, content T) BroChatClientContentResult[T] {
	if err, ok := err.(net.Error); ok && err.Timeout() {
		// If it was a timeout error
		return makeBroChatClientContentResult(BROCHAT_RESPONSE_CODE_CONNECTION_TIMEOUT_ERROR, content)
	}

	return makeBroChatClientContentResult(BROCHAT_RESPONSE_CODE_GENERIC_CONNECTION_ERROR, content)
}

// handleHttpRequestError creates a BroChatClientResult generated from an error after attempting an http request.
func handleHttpRequestError(err error) BroChatClientResult {
	if err, ok := err.(net.Error); ok && err.Timeout() {
		// If it was a timeout error
		return makeBroChatClientResult(BROCHAT_RESPONSE_CODE_CONNECTION_TIMEOUT_ERROR)
	}

	return makeBroChatClientResult(BROCHAT_RESPONSE_CODE_GENERIC_CONNECTION_ERROR)
}

// handleUnsuccessfulStatusCodeWithContent is a helper function that handles the response from the server when the response is not successful.
// It returns a BroChatClientContentResult with the given content and a BroChatClientResult with the given response code and error details.
func handleUnsuccessfulStatusCodeWithContent[T any](res *http.Response, content T) BroChatClientContentResult[T] {
	return BroChatClientContentResult[T]{Content: content,
		BroChatClientResult: handleUnsuccessfulStatusCode(res),
	}
}

// handleUnsuccessfulStatusCode is a helper function that handles the response from the server when the response is not successful.
func handleUnsuccessfulStatusCode(res *http.Response) BroChatClientResult {
	var serverSideErr BroChatError

	err := json.NewDecoder(res.Body).Decode(&serverSideErr)

	if err != nil {
		switch res.StatusCode {
		case http.StatusUnauthorized:
			return makeBroChatClientResult(BROCHAT_RESPONSE_CODE_UNAUTHORIZED_ERROR)
		case http.StatusForbidden:
			return makeBroChatClientResult(BROCHAT_RESPONSE_CODE_FORBIDDEN_ERROR)
		case http.StatusNotFound:
			return makeBroChatClientResult(BROCHAT_RESPONSE_CODE_NOT_FOUND_ERROR)
		case http.StatusBadRequest:
			return makeBroChatClientResult(BROCHAT_RESPONSE_CODE_VALIDATION_ERROR)
		default:
			return makeBroChatClientResult(BROCHAT_RESPONSE_CODE_UNHANDLED_ERROR)
		}
	}

	return makeBroChatClientResult(serverSideErr.Code, serverSideErr.ErrorDetails...)
}
