/*
Farcaster API V2

The Farcaster API allows you to interact with the Farcaster protocol. See the [Neynar docs](https://docs.neynar.com/reference) for more details.

API version: 2.43.0
Contact: team@neynar.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package neynar_sdk

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"net/url"
)

type NotificationsAPI interface {

	/*
		FetchAllNotifications For user

		Returns a list of notifications for a specific FID.

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@return ApiFetchAllNotificationsRequest
	*/
	FetchAllNotifications(ctx context.Context) ApiFetchAllNotificationsRequest

	// FetchAllNotificationsExecute executes the request
	//  @return NotificationsResponse
	FetchAllNotificationsExecute(r ApiFetchAllNotificationsRequest) (*NotificationsResponse, *http.Response, error)

	/*
		FetchChannelNotificationsForUser For user by channel

		Returns a list of notifications for a user in specific channels

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@return ApiFetchChannelNotificationsForUserRequest
	*/
	FetchChannelNotificationsForUser(ctx context.Context) ApiFetchChannelNotificationsForUserRequest

	// FetchChannelNotificationsForUserExecute executes the request
	//  @return NotificationsResponse
	FetchChannelNotificationsForUserExecute(r ApiFetchChannelNotificationsForUserRequest) (*NotificationsResponse, *http.Response, error)

	/*
		FetchNotificationsByParentUrlForUser For user by parent_urls

		Returns a list of notifications for a user in specific parent_urls

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@return ApiFetchNotificationsByParentUrlForUserRequest
	*/
	FetchNotificationsByParentUrlForUser(ctx context.Context) ApiFetchNotificationsByParentUrlForUserRequest

	// FetchNotificationsByParentUrlForUserExecute executes the request
	//  @return NotificationsResponse
	FetchNotificationsByParentUrlForUserExecute(r ApiFetchNotificationsByParentUrlForUserRequest) (*NotificationsResponse, *http.Response, error)

	/*
			MarkNotificationsAsSeen Mark as seen

			Mark notifications as seen.
		You can choose one of two authorization methods, either:
		  1. Provide a valid signer_uuid in the request body (Most common)
		  2. Provide a valid, signed "Bearer" token in the request's `Authorization` header similar to the
		     approach described [here](https://docs.farcaster.xyz/reference/warpcast/api#authentication)


			@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
			@return ApiMarkNotificationsAsSeenRequest
	*/
	MarkNotificationsAsSeen(ctx context.Context) ApiMarkNotificationsAsSeenRequest

	// MarkNotificationsAsSeenExecute executes the request
	//  @return OperationResponse
	MarkNotificationsAsSeenExecute(r ApiMarkNotificationsAsSeenRequest) (*OperationResponse, *http.Response, error)
}

// NotificationsAPIService NotificationsAPI service
type NotificationsAPIService service

type ApiFetchAllNotificationsRequest struct {
	ctx          context.Context
	ApiService   NotificationsAPI
	fid          *int32
	type_        *[]NotificationType
	priorityMode *bool
	limit        *int32
	cursor       *string
}

// FID of the user you you want to fetch notifications for. The response will respect this user&#39;s mutes and blocks.
func (r ApiFetchAllNotificationsRequest) Fid(fid int32) ApiFetchAllNotificationsRequest {
	r.fid = &fid
	return r
}

// Notification type to fetch. Comma separated values of follows, recasts, likes, mentions, replies.
func (r ApiFetchAllNotificationsRequest) Type_(type_ []NotificationType) ApiFetchAllNotificationsRequest {
	r.type_ = &type_
	return r
}

// When true, only returns notifications from power badge users and users that the user follows.
func (r ApiFetchAllNotificationsRequest) PriorityMode(priorityMode bool) ApiFetchAllNotificationsRequest {
	r.priorityMode = &priorityMode
	return r
}

// Number of results to fetch
func (r ApiFetchAllNotificationsRequest) Limit(limit int32) ApiFetchAllNotificationsRequest {
	r.limit = &limit
	return r
}

// Pagination cursor.
func (r ApiFetchAllNotificationsRequest) Cursor(cursor string) ApiFetchAllNotificationsRequest {
	r.cursor = &cursor
	return r
}

func (r ApiFetchAllNotificationsRequest) Execute() (*NotificationsResponse, *http.Response, error) {
	return r.ApiService.FetchAllNotificationsExecute(r)
}

/*
FetchAllNotifications For user

Returns a list of notifications for a specific FID.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiFetchAllNotificationsRequest
*/
func (a *NotificationsAPIService) FetchAllNotifications(ctx context.Context) ApiFetchAllNotificationsRequest {
	return ApiFetchAllNotificationsRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return NotificationsResponse
func (a *NotificationsAPIService) FetchAllNotificationsExecute(r ApiFetchAllNotificationsRequest) (*NotificationsResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *NotificationsResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "NotificationsAPIService.FetchAllNotifications")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/farcaster/notifications"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.fid == nil {
		return localVarReturnValue, nil, reportError("fid is required and must be specified")
	}

	parameterAddToHeaderOrQuery(localVarQueryParams, "fid", r.fid, "form", "")
	if r.type_ != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "type", r.type_, "form", "csv")
	}
	if r.priorityMode != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "priority_mode", r.priorityMode, "form", "")
	} else {
		var defaultValue bool = false
		r.priorityMode = &defaultValue
	}
	if r.limit != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "limit", r.limit, "form", "")
	} else {
		var defaultValue int32 = 15
		r.limit = &defaultValue
	}
	if r.cursor != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "cursor", r.cursor, "form", "")
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["ApiKeyAuth"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarHeaderParams["x-api-key"] = key
			}
		}
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 400 {
			var v ErrorRes
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
			newErr.model = v
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiFetchChannelNotificationsForUserRequest struct {
	ctx          context.Context
	ApiService   NotificationsAPI
	fid          *int32
	channelIds   *string
	priorityMode *bool
	limit        *int32
	cursor       *string
}

// FID of the user you you want to fetch notifications for. The response will respect this user&#39;s mutes and blocks.
func (r ApiFetchChannelNotificationsForUserRequest) Fid(fid int32) ApiFetchChannelNotificationsForUserRequest {
	r.fid = &fid
	return r
}

// Comma separated channel_ids (find list of all channels here - https://docs.neynar.com/reference/list-all-channels)
func (r ApiFetchChannelNotificationsForUserRequest) ChannelIds(channelIds string) ApiFetchChannelNotificationsForUserRequest {
	r.channelIds = &channelIds
	return r
}

// When true, only returns notifications from power badge users and users that the user follows.
func (r ApiFetchChannelNotificationsForUserRequest) PriorityMode(priorityMode bool) ApiFetchChannelNotificationsForUserRequest {
	r.priorityMode = &priorityMode
	return r
}

// Number of results to fetch
func (r ApiFetchChannelNotificationsForUserRequest) Limit(limit int32) ApiFetchChannelNotificationsForUserRequest {
	r.limit = &limit
	return r
}

// Pagination cursor.
func (r ApiFetchChannelNotificationsForUserRequest) Cursor(cursor string) ApiFetchChannelNotificationsForUserRequest {
	r.cursor = &cursor
	return r
}

func (r ApiFetchChannelNotificationsForUserRequest) Execute() (*NotificationsResponse, *http.Response, error) {
	return r.ApiService.FetchChannelNotificationsForUserExecute(r)
}

/*
FetchChannelNotificationsForUser For user by channel

Returns a list of notifications for a user in specific channels

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiFetchChannelNotificationsForUserRequest
*/
func (a *NotificationsAPIService) FetchChannelNotificationsForUser(ctx context.Context) ApiFetchChannelNotificationsForUserRequest {
	return ApiFetchChannelNotificationsForUserRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return NotificationsResponse
func (a *NotificationsAPIService) FetchChannelNotificationsForUserExecute(r ApiFetchChannelNotificationsForUserRequest) (*NotificationsResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *NotificationsResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "NotificationsAPIService.FetchChannelNotificationsForUser")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/farcaster/notifications/channel"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.fid == nil {
		return localVarReturnValue, nil, reportError("fid is required and must be specified")
	}
	if r.channelIds == nil {
		return localVarReturnValue, nil, reportError("channelIds is required and must be specified")
	}

	parameterAddToHeaderOrQuery(localVarQueryParams, "fid", r.fid, "form", "")
	parameterAddToHeaderOrQuery(localVarQueryParams, "channel_ids", r.channelIds, "form", "")
	if r.priorityMode != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "priority_mode", r.priorityMode, "form", "")
	} else {
		var defaultValue bool = false
		r.priorityMode = &defaultValue
	}
	if r.limit != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "limit", r.limit, "form", "")
	} else {
		var defaultValue int32 = 15
		r.limit = &defaultValue
	}
	if r.cursor != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "cursor", r.cursor, "form", "")
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["ApiKeyAuth"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarHeaderParams["x-api-key"] = key
			}
		}
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 400 {
			var v ErrorRes
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
			newErr.model = v
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiFetchNotificationsByParentUrlForUserRequest struct {
	ctx          context.Context
	ApiService   NotificationsAPI
	fid          *int32
	parentUrls   *string
	priorityMode *bool
	limit        *int32
	cursor       *string
}

// FID of the user you you want to fetch notifications for. The response will respect this user&#39;s mutes and blocks.
func (r ApiFetchNotificationsByParentUrlForUserRequest) Fid(fid int32) ApiFetchNotificationsByParentUrlForUserRequest {
	r.fid = &fid
	return r
}

// Comma separated parent_urls
func (r ApiFetchNotificationsByParentUrlForUserRequest) ParentUrls(parentUrls string) ApiFetchNotificationsByParentUrlForUserRequest {
	r.parentUrls = &parentUrls
	return r
}

// When true, only returns notifications from power badge users and users that the user follows.
func (r ApiFetchNotificationsByParentUrlForUserRequest) PriorityMode(priorityMode bool) ApiFetchNotificationsByParentUrlForUserRequest {
	r.priorityMode = &priorityMode
	return r
}

// Number of results to fetch
func (r ApiFetchNotificationsByParentUrlForUserRequest) Limit(limit int32) ApiFetchNotificationsByParentUrlForUserRequest {
	r.limit = &limit
	return r
}

// Pagination cursor.
func (r ApiFetchNotificationsByParentUrlForUserRequest) Cursor(cursor string) ApiFetchNotificationsByParentUrlForUserRequest {
	r.cursor = &cursor
	return r
}

func (r ApiFetchNotificationsByParentUrlForUserRequest) Execute() (*NotificationsResponse, *http.Response, error) {
	return r.ApiService.FetchNotificationsByParentUrlForUserExecute(r)
}

/*
FetchNotificationsByParentUrlForUser For user by parent_urls

Returns a list of notifications for a user in specific parent_urls

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiFetchNotificationsByParentUrlForUserRequest
*/
func (a *NotificationsAPIService) FetchNotificationsByParentUrlForUser(ctx context.Context) ApiFetchNotificationsByParentUrlForUserRequest {
	return ApiFetchNotificationsByParentUrlForUserRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return NotificationsResponse
func (a *NotificationsAPIService) FetchNotificationsByParentUrlForUserExecute(r ApiFetchNotificationsByParentUrlForUserRequest) (*NotificationsResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *NotificationsResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "NotificationsAPIService.FetchNotificationsByParentUrlForUser")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/farcaster/notifications/parent_url"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.fid == nil {
		return localVarReturnValue, nil, reportError("fid is required and must be specified")
	}
	if r.parentUrls == nil {
		return localVarReturnValue, nil, reportError("parentUrls is required and must be specified")
	}

	parameterAddToHeaderOrQuery(localVarQueryParams, "fid", r.fid, "form", "")
	parameterAddToHeaderOrQuery(localVarQueryParams, "parent_urls", r.parentUrls, "form", "")
	if r.priorityMode != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "priority_mode", r.priorityMode, "form", "")
	} else {
		var defaultValue bool = false
		r.priorityMode = &defaultValue
	}
	if r.limit != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "limit", r.limit, "form", "")
	} else {
		var defaultValue int32 = 15
		r.limit = &defaultValue
	}
	if r.cursor != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "cursor", r.cursor, "form", "")
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["ApiKeyAuth"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarHeaderParams["x-api-key"] = key
			}
		}
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 400 {
			var v ErrorRes
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
			newErr.model = v
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiMarkNotificationsAsSeenRequest struct {
	ctx                            context.Context
	ApiService                     NotificationsAPI
	markNotificationsAsSeenReqBody *MarkNotificationsAsSeenReqBody
	authorization                  *string
}

func (r ApiMarkNotificationsAsSeenRequest) MarkNotificationsAsSeenReqBody(markNotificationsAsSeenReqBody MarkNotificationsAsSeenReqBody) ApiMarkNotificationsAsSeenRequest {
	r.markNotificationsAsSeenReqBody = &markNotificationsAsSeenReqBody
	return r
}

// Optional Bearer token for certain endpoints. The token format is described [here](https://docs.farcaster.xyz/reference/warpcast/api#authentication).
func (r ApiMarkNotificationsAsSeenRequest) Authorization(authorization string) ApiMarkNotificationsAsSeenRequest {
	r.authorization = &authorization
	return r
}

func (r ApiMarkNotificationsAsSeenRequest) Execute() (*OperationResponse, *http.Response, error) {
	return r.ApiService.MarkNotificationsAsSeenExecute(r)
}

/*
MarkNotificationsAsSeen Mark as seen

Mark notifications as seen.
You can choose one of two authorization methods, either:

 1. Provide a valid signer_uuid in the request body (Most common)

 2. Provide a valid, signed "Bearer" token in the request's `Authorization` header similar to the
    approach described [here](https://docs.farcaster.xyz/reference/warpcast/api#authentication)

    @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
    @return ApiMarkNotificationsAsSeenRequest
*/
func (a *NotificationsAPIService) MarkNotificationsAsSeen(ctx context.Context) ApiMarkNotificationsAsSeenRequest {
	return ApiMarkNotificationsAsSeenRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return OperationResponse
func (a *NotificationsAPIService) MarkNotificationsAsSeenExecute(r ApiMarkNotificationsAsSeenRequest) (*OperationResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *OperationResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "NotificationsAPIService.MarkNotificationsAsSeen")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/farcaster/notifications/seen"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.markNotificationsAsSeenReqBody == nil {
		return localVarReturnValue, nil, reportError("markNotificationsAsSeenReqBody is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	if r.authorization != nil {
		parameterAddToHeaderOrQuery(localVarHeaderParams, "Authorization", r.authorization, "simple", "")
	}
	// body params
	localVarPostBody = r.markNotificationsAsSeenReqBody
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["ApiKeyAuth"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarHeaderParams["x-api-key"] = key
			}
		}
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 400 {
			var v ErrorRes
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 500 {
			var v ErrorRes
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
			newErr.model = v
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}
