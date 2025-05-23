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

type ReactionAPI interface {

	/*
			DeleteReaction Delete reaction

			Delete a reaction (like or recast) to a cast \
		(In order to delete a reaction `signer_uuid` must be approved)


			@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
			@return ApiDeleteReactionRequest
	*/
	DeleteReaction(ctx context.Context) ApiDeleteReactionRequest

	// DeleteReactionExecute executes the request
	//  @return OperationResponse
	DeleteReactionExecute(r ApiDeleteReactionRequest) (*OperationResponse, *http.Response, error)

	/*
		FetchCastReactions Reactions for cast

		Fetches reactions for a given cast

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@return ApiFetchCastReactionsRequest
	*/
	FetchCastReactions(ctx context.Context) ApiFetchCastReactionsRequest

	// FetchCastReactionsExecute executes the request
	//  @return ReactionsCastResponse
	FetchCastReactionsExecute(r ApiFetchCastReactionsRequest) (*ReactionsCastResponse, *http.Response, error)

	/*
		FetchUserReactions Reactions for user

		Fetches reactions for a given user

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@return ApiFetchUserReactionsRequest
	*/
	FetchUserReactions(ctx context.Context) ApiFetchUserReactionsRequest

	// FetchUserReactionsExecute executes the request
	//  @return ReactionsResponse
	FetchUserReactionsExecute(r ApiFetchUserReactionsRequest) (*ReactionsResponse, *http.Response, error)

	/*
			PublishReaction Post a reaction

			Post a reaction (like or recast) to a given cast \
		(In order to post a reaction `signer_uuid` must be approved)


			@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
			@return ApiPublishReactionRequest
	*/
	PublishReaction(ctx context.Context) ApiPublishReactionRequest

	// PublishReactionExecute executes the request
	//  @return OperationResponse
	PublishReactionExecute(r ApiPublishReactionRequest) (*OperationResponse, *http.Response, error)
}

// ReactionAPIService ReactionAPI service
type ReactionAPIService service

type ApiDeleteReactionRequest struct {
	ctx             context.Context
	ApiService      ReactionAPI
	reactionReqBody *ReactionReqBody
}

func (r ApiDeleteReactionRequest) ReactionReqBody(reactionReqBody ReactionReqBody) ApiDeleteReactionRequest {
	r.reactionReqBody = &reactionReqBody
	return r
}

func (r ApiDeleteReactionRequest) Execute() (*OperationResponse, *http.Response, error) {
	return r.ApiService.DeleteReactionExecute(r)
}

/*
DeleteReaction Delete reaction

Delete a reaction (like or recast) to a cast \
(In order to delete a reaction `signer_uuid` must be approved)

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiDeleteReactionRequest
*/
func (a *ReactionAPIService) DeleteReaction(ctx context.Context) ApiDeleteReactionRequest {
	return ApiDeleteReactionRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return OperationResponse
func (a *ReactionAPIService) DeleteReactionExecute(r ApiDeleteReactionRequest) (*OperationResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodDelete
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *OperationResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ReactionAPIService.DeleteReaction")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/farcaster/reaction"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.reactionReqBody == nil {
		return localVarReturnValue, nil, reportError("reactionReqBody is required and must be specified")
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
	// body params
	localVarPostBody = r.reactionReqBody
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
		if localVarHTTPResponse.StatusCode == 403 {
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
		if localVarHTTPResponse.StatusCode == 404 {
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

type ApiFetchCastReactionsRequest struct {
	ctx        context.Context
	ApiService ReactionAPI
	hash       *string
	types      *[]ReactionsType
	viewerFid  *int32
	limit      *int32
	cursor     *string
}

func (r ApiFetchCastReactionsRequest) Hash(hash string) ApiFetchCastReactionsRequest {
	r.hash = &hash
	return r
}

// Customize which reaction types the request should search for. This is a comma-separated string that can include the following values: &#39;likes&#39; and &#39;recasts&#39;. By default api returns both. To select multiple types, use a comma-separated list of these values.
func (r ApiFetchCastReactionsRequest) Types(types []ReactionsType) ApiFetchCastReactionsRequest {
	r.types = &types
	return r
}

// Providing this will return a list of reactions that respects this user&#39;s mutes and blocks and includes &#x60;viewer_context&#x60;.
func (r ApiFetchCastReactionsRequest) ViewerFid(viewerFid int32) ApiFetchCastReactionsRequest {
	r.viewerFid = &viewerFid
	return r
}

// Number of results to fetch
func (r ApiFetchCastReactionsRequest) Limit(limit int32) ApiFetchCastReactionsRequest {
	r.limit = &limit
	return r
}

// Pagination cursor.
func (r ApiFetchCastReactionsRequest) Cursor(cursor string) ApiFetchCastReactionsRequest {
	r.cursor = &cursor
	return r
}

func (r ApiFetchCastReactionsRequest) Execute() (*ReactionsCastResponse, *http.Response, error) {
	return r.ApiService.FetchCastReactionsExecute(r)
}

/*
FetchCastReactions Reactions for cast

Fetches reactions for a given cast

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiFetchCastReactionsRequest
*/
func (a *ReactionAPIService) FetchCastReactions(ctx context.Context) ApiFetchCastReactionsRequest {
	return ApiFetchCastReactionsRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return ReactionsCastResponse
func (a *ReactionAPIService) FetchCastReactionsExecute(r ApiFetchCastReactionsRequest) (*ReactionsCastResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *ReactionsCastResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ReactionAPIService.FetchCastReactions")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/farcaster/reactions/cast"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.hash == nil {
		return localVarReturnValue, nil, reportError("hash is required and must be specified")
	}
	if r.types == nil {
		return localVarReturnValue, nil, reportError("types is required and must be specified")
	}

	parameterAddToHeaderOrQuery(localVarQueryParams, "hash", r.hash, "form", "")
	parameterAddToHeaderOrQuery(localVarQueryParams, "types", r.types, "form", "csv")
	if r.viewerFid != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "viewer_fid", r.viewerFid, "form", "")
	}
	if r.limit != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "limit", r.limit, "form", "")
	} else {
		var defaultValue int32 = 25
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

type ApiFetchUserReactionsRequest struct {
	ctx        context.Context
	ApiService ReactionAPI
	fid        *int32
	type_      *ReactionsType
	viewerFid  *int32
	limit      *int32
	cursor     *string
}

func (r ApiFetchUserReactionsRequest) Fid(fid int32) ApiFetchUserReactionsRequest {
	r.fid = &fid
	return r
}

// Type of reaction to fetch (likes or recasts or all)
func (r ApiFetchUserReactionsRequest) Type_(type_ ReactionsType) ApiFetchUserReactionsRequest {
	r.type_ = &type_
	return r
}

// Providing this will return a list of reactions that respects this user&#39;s mutes and blocks and includes &#x60;viewer_context&#x60;.
func (r ApiFetchUserReactionsRequest) ViewerFid(viewerFid int32) ApiFetchUserReactionsRequest {
	r.viewerFid = &viewerFid
	return r
}

// Number of results to fetch
func (r ApiFetchUserReactionsRequest) Limit(limit int32) ApiFetchUserReactionsRequest {
	r.limit = &limit
	return r
}

// Pagination cursor.
func (r ApiFetchUserReactionsRequest) Cursor(cursor string) ApiFetchUserReactionsRequest {
	r.cursor = &cursor
	return r
}

func (r ApiFetchUserReactionsRequest) Execute() (*ReactionsResponse, *http.Response, error) {
	return r.ApiService.FetchUserReactionsExecute(r)
}

/*
FetchUserReactions Reactions for user

Fetches reactions for a given user

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiFetchUserReactionsRequest
*/
func (a *ReactionAPIService) FetchUserReactions(ctx context.Context) ApiFetchUserReactionsRequest {
	return ApiFetchUserReactionsRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return ReactionsResponse
func (a *ReactionAPIService) FetchUserReactionsExecute(r ApiFetchUserReactionsRequest) (*ReactionsResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *ReactionsResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ReactionAPIService.FetchUserReactions")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/farcaster/reactions/user"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.fid == nil {
		return localVarReturnValue, nil, reportError("fid is required and must be specified")
	}
	if r.type_ == nil {
		return localVarReturnValue, nil, reportError("type_ is required and must be specified")
	}

	parameterAddToHeaderOrQuery(localVarQueryParams, "fid", r.fid, "form", "")
	if r.viewerFid != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "viewer_fid", r.viewerFid, "form", "")
	}
	parameterAddToHeaderOrQuery(localVarQueryParams, "type", r.type_, "form", "")
	if r.limit != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "limit", r.limit, "form", "")
	} else {
		var defaultValue int32 = 25
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

type ApiPublishReactionRequest struct {
	ctx             context.Context
	ApiService      ReactionAPI
	reactionReqBody *ReactionReqBody
}

func (r ApiPublishReactionRequest) ReactionReqBody(reactionReqBody ReactionReqBody) ApiPublishReactionRequest {
	r.reactionReqBody = &reactionReqBody
	return r
}

func (r ApiPublishReactionRequest) Execute() (*OperationResponse, *http.Response, error) {
	return r.ApiService.PublishReactionExecute(r)
}

/*
PublishReaction Post a reaction

Post a reaction (like or recast) to a given cast \
(In order to post a reaction `signer_uuid` must be approved)

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiPublishReactionRequest
*/
func (a *ReactionAPIService) PublishReaction(ctx context.Context) ApiPublishReactionRequest {
	return ApiPublishReactionRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return OperationResponse
func (a *ReactionAPIService) PublishReactionExecute(r ApiPublishReactionRequest) (*OperationResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *OperationResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ReactionAPIService.PublishReaction")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/farcaster/reaction"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.reactionReqBody == nil {
		return localVarReturnValue, nil, reportError("reactionReqBody is required and must be specified")
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
	// body params
	localVarPostBody = r.reactionReqBody
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
		if localVarHTTPResponse.StatusCode == 403 {
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
		if localVarHTTPResponse.StatusCode == 404 {
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
