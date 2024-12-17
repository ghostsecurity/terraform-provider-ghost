/*
Ghost API

The Ghost API

API version: 2.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"net/url"
	"strings"
)


// CampaignsAPIService CampaignsAPI service
type CampaignsAPIService service

type ApiGetCampaignRequest struct {
	ctx context.Context
	ApiService *CampaignsAPIService
	campaignId string
}

func (r ApiGetCampaignRequest) Execute() (*Campaign, *http.Response, error) {
	return r.ApiService.GetCampaignExecute(r)
}

/*
GetCampaign Get a campaign

Get a campaign

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param campaignId Campaign ID as UUID
 @return ApiGetCampaignRequest
*/
func (a *CampaignsAPIService) GetCampaign(ctx context.Context, campaignId string) ApiGetCampaignRequest {
	return ApiGetCampaignRequest{
		ApiService: a,
		ctx: ctx,
		campaignId: campaignId,
	}
}

// Execute executes the request
//  @return Campaign
func (a *CampaignsAPIService) GetCampaignExecute(r ApiGetCampaignRequest) (*Campaign, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *Campaign
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "CampaignsAPIService.GetCampaign")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/campaigns/{campaign_id}"
	localVarPath = strings.Replace(localVarPath, "{"+"campaign_id"+"}", url.PathEscape(parameterValueToString(r.campaignId, "campaignId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

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
			if apiKey, ok := auth["BearerAuth"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarHeaderParams["Authorization"] = key
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
			var v InvalidInputError
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 401 {
			var v AuthError
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
			var v AccessDeniedError
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
			var v NotFoundError
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 422 {
			var v ValidationListError
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

type ApiListCampaignIssueCategoriesRequest struct {
	ctx context.Context
	ApiService *CampaignsAPIService
	campaignId string
	size *int32
	page *int32
	orderBy *string
	vulnerabilitiesHasActive *bool
	name *string
	issueSeverity *[]string
	issueId *[]string
	id *[]string
}

// Results per page
func (r ApiListCampaignIssueCategoriesRequest) Size(size int32) ApiListCampaignIssueCategoriesRequest {
	r.size = &size
	return r
}

// Page number
func (r ApiListCampaignIssueCategoriesRequest) Page(page int32) ApiListCampaignIssueCategoriesRequest {
	r.page = &page
	return r
}

// Ordering attribute with optional &#39;-&#39; prefix for descending order. Valid attributes: name, vulnerabilities.active, vulnerabilities.resolved
func (r ApiListCampaignIssueCategoriesRequest) OrderBy(orderBy string) ApiListCampaignIssueCategoriesRequest {
	r.orderBy = &orderBy
	return r
}

// Filter for categories that have active vulnerabilities
func (r ApiListCampaignIssueCategoriesRequest) VulnerabilitiesHasActive(vulnerabilitiesHasActive bool) ApiListCampaignIssueCategoriesRequest {
	r.vulnerabilitiesHasActive = &vulnerabilitiesHasActive
	return r
}

// Filter by name. Supports partial matches.
func (r ApiListCampaignIssueCategoriesRequest) Name(name string) ApiListCampaignIssueCategoriesRequest {
	r.name = &name
	return r
}

// Filter for categories that have issues that have a severity in the provided list
func (r ApiListCampaignIssueCategoriesRequest) IssueSeverity(issueSeverity []string) ApiListCampaignIssueCategoriesRequest {
	r.issueSeverity = &issueSeverity
	return r
}

// Filter for categories that have issues that have an ID in the provided list
func (r ApiListCampaignIssueCategoriesRequest) IssueId(issueId []string) ApiListCampaignIssueCategoriesRequest {
	r.issueId = &issueId
	return r
}

// Filter for categories that have an ID in the provided list
func (r ApiListCampaignIssueCategoriesRequest) Id(id []string) ApiListCampaignIssueCategoriesRequest {
	r.id = &id
	return r
}

func (r ApiListCampaignIssueCategoriesRequest) Execute() (*PaginatedIssueCategory, *http.Response, error) {
	return r.ApiService.ListCampaignIssueCategoriesExecute(r)
}

/*
ListCampaignIssueCategories List campaign issue categories

List campaign issue categories

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param campaignId Campaign ID as UUID
 @return ApiListCampaignIssueCategoriesRequest
*/
func (a *CampaignsAPIService) ListCampaignIssueCategories(ctx context.Context, campaignId string) ApiListCampaignIssueCategoriesRequest {
	return ApiListCampaignIssueCategoriesRequest{
		ApiService: a,
		ctx: ctx,
		campaignId: campaignId,
	}
}

// Execute executes the request
//  @return PaginatedIssueCategory
func (a *CampaignsAPIService) ListCampaignIssueCategoriesExecute(r ApiListCampaignIssueCategoriesRequest) (*PaginatedIssueCategory, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *PaginatedIssueCategory
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "CampaignsAPIService.ListCampaignIssueCategories")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/campaigns/{campaign_id}/issue_categories"
	localVarPath = strings.Replace(localVarPath, "{"+"campaign_id"+"}", url.PathEscape(parameterValueToString(r.campaignId, "campaignId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.size != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "size", r.size, "form", "")
	}
	if r.page != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "page", r.page, "form", "")
	}
	if r.orderBy != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "order_by", r.orderBy, "form", "")
	}
	if r.vulnerabilitiesHasActive != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "vulnerabilities.has_active", r.vulnerabilitiesHasActive, "form", "")
	}
	if r.name != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "name", r.name, "form", "")
	}
	if r.issueSeverity != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "issue.severity", r.issueSeverity, "form", "csv")
	}
	if r.issueId != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "issue.id", r.issueId, "form", "csv")
	}
	if r.id != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "id", r.id, "form", "csv")
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
			if apiKey, ok := auth["BearerAuth"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarHeaderParams["Authorization"] = key
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
			var v InvalidInputError
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 401 {
			var v AuthError
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
			var v AccessDeniedError
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 422 {
			var v ValidationListError
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

type ApiListCampaignIssuesRequest struct {
	ctx context.Context
	ApiService *CampaignsAPIService
	campaignId string
	size *int32
	page *int32
	orderBy *string
	categoryId *[]string
	vulnerabilitiesHasActive *int32
	name *string
	severity *[]string
	id *[]string
}

// Results per page
func (r ApiListCampaignIssuesRequest) Size(size int32) ApiListCampaignIssuesRequest {
	r.size = &size
	return r
}

// Page number
func (r ApiListCampaignIssuesRequest) Page(page int32) ApiListCampaignIssuesRequest {
	r.page = &page
	return r
}

// Ordering attribute with optional &#39;-&#39; prefix for descending order. Valid attributes: name, severity, category.name, vulnerabilities.active, vulnerabilities.resolved
func (r ApiListCampaignIssuesRequest) OrderBy(orderBy string) ApiListCampaignIssuesRequest {
	r.orderBy = &orderBy
	return r
}

// Filter for issues that belong to a category in the provided list
func (r ApiListCampaignIssuesRequest) CategoryId(categoryId []string) ApiListCampaignIssuesRequest {
	r.categoryId = &categoryId
	return r
}

// Filter for issues that have active vulnerabilities
func (r ApiListCampaignIssuesRequest) VulnerabilitiesHasActive(vulnerabilitiesHasActive int32) ApiListCampaignIssuesRequest {
	r.vulnerabilitiesHasActive = &vulnerabilitiesHasActive
	return r
}

// Filter by name. Supports partially matching.
func (r ApiListCampaignIssuesRequest) Name(name string) ApiListCampaignIssuesRequest {
	r.name = &name
	return r
}

// Filter for issues that have a severity in the provided list
func (r ApiListCampaignIssuesRequest) Severity(severity []string) ApiListCampaignIssuesRequest {
	r.severity = &severity
	return r
}

// Filter for issues that have an ID in the provided list
func (r ApiListCampaignIssuesRequest) Id(id []string) ApiListCampaignIssuesRequest {
	r.id = &id
	return r
}

func (r ApiListCampaignIssuesRequest) Execute() (*PaginatedIssue, *http.Response, error) {
	return r.ApiService.ListCampaignIssuesExecute(r)
}

/*
ListCampaignIssues List campaign issues

List campaign issues

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param campaignId Campaign ID as UUID
 @return ApiListCampaignIssuesRequest
*/
func (a *CampaignsAPIService) ListCampaignIssues(ctx context.Context, campaignId string) ApiListCampaignIssuesRequest {
	return ApiListCampaignIssuesRequest{
		ApiService: a,
		ctx: ctx,
		campaignId: campaignId,
	}
}

// Execute executes the request
//  @return PaginatedIssue
func (a *CampaignsAPIService) ListCampaignIssuesExecute(r ApiListCampaignIssuesRequest) (*PaginatedIssue, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *PaginatedIssue
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "CampaignsAPIService.ListCampaignIssues")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/campaigns/{campaign_id}/issues"
	localVarPath = strings.Replace(localVarPath, "{"+"campaign_id"+"}", url.PathEscape(parameterValueToString(r.campaignId, "campaignId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.size != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "size", r.size, "form", "")
	}
	if r.page != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "page", r.page, "form", "")
	}
	if r.orderBy != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "order_by", r.orderBy, "form", "")
	}
	if r.categoryId != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "category.id", r.categoryId, "form", "csv")
	}
	if r.vulnerabilitiesHasActive != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "vulnerabilities.has_active", r.vulnerabilitiesHasActive, "form", "")
	}
	if r.name != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "name", r.name, "form", "")
	}
	if r.severity != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "severity", r.severity, "form", "csv")
	}
	if r.id != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "id", r.id, "form", "csv")
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
			if apiKey, ok := auth["BearerAuth"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarHeaderParams["Authorization"] = key
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
			var v InvalidInputError
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 401 {
			var v AuthError
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
			var v AccessDeniedError
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 422 {
			var v ValidationListError
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

type ApiListCampaignVulnerabilitiesRequest struct {
	ctx context.Context
	ApiService *CampaignsAPIService
	campaignId string
	size *int32
	page *int32
	orderBy *string
	status *[]string
	firstDetectedAt *string
	lastDetectedAt *string
	resourceKind *[]string
	issueSeverity *[]string
	issueId *[]string
	issueCategoryId *[]string
}

// Results per page
func (r ApiListCampaignVulnerabilitiesRequest) Size(size int32) ApiListCampaignVulnerabilitiesRequest {
	r.size = &size
	return r
}

// Page number
func (r ApiListCampaignVulnerabilitiesRequest) Page(page int32) ApiListCampaignVulnerabilitiesRequest {
	r.page = &page
	return r
}

// Ordering attribute with optional &#39;-&#39; prefix for descending order. Valid attributes: issue.severity, status, resource.kind, issue.category.name, issue.name, first_detected_at, and last_detected_at.
func (r ApiListCampaignVulnerabilitiesRequest) OrderBy(orderBy string) ApiListCampaignVulnerabilitiesRequest {
	r.orderBy = &orderBy
	return r
}

// Filters for vulnerabilities that have a status in the provided list.
func (r ApiListCampaignVulnerabilitiesRequest) Status(status []string) ApiListCampaignVulnerabilitiesRequest {
	r.status = &status
	return r
}

// Filters for vulnerabilities that were first seen since the provided value - can be one of day, week, month, year
func (r ApiListCampaignVulnerabilitiesRequest) FirstDetectedAt(firstDetectedAt string) ApiListCampaignVulnerabilitiesRequest {
	r.firstDetectedAt = &firstDetectedAt
	return r
}

// Filters for vulnerabilities that were last seen since the provided value - can be one of day, week, month, year
func (r ApiListCampaignVulnerabilitiesRequest) LastDetectedAt(lastDetectedAt string) ApiListCampaignVulnerabilitiesRequest {
	r.lastDetectedAt = &lastDetectedAt
	return r
}

// Filters for vulnerabilities that are associated with a resource kind in the provided list.
func (r ApiListCampaignVulnerabilitiesRequest) ResourceKind(resourceKind []string) ApiListCampaignVulnerabilitiesRequest {
	r.resourceKind = &resourceKind
	return r
}

// Filters for vulnerabilities that belongs to a rule with a severity in the provided list.
func (r ApiListCampaignVulnerabilitiesRequest) IssueSeverity(issueSeverity []string) ApiListCampaignVulnerabilitiesRequest {
	r.issueSeverity = &issueSeverity
	return r
}

// Filters for vulnerabilities that belolng to a rule in the provided list.
func (r ApiListCampaignVulnerabilitiesRequest) IssueId(issueId []string) ApiListCampaignVulnerabilitiesRequest {
	r.issueId = &issueId
	return r
}

// Filters for vulnerabilities that belong to a rule that belongs to a category in the provided list.
func (r ApiListCampaignVulnerabilitiesRequest) IssueCategoryId(issueCategoryId []string) ApiListCampaignVulnerabilitiesRequest {
	r.issueCategoryId = &issueCategoryId
	return r
}

func (r ApiListCampaignVulnerabilitiesRequest) Execute() (*PaginatedVulnerability, *http.Response, error) {
	return r.ApiService.ListCampaignVulnerabilitiesExecute(r)
}

/*
ListCampaignVulnerabilities List campaign vulnerabilities

List campaign vulnerabilities

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param campaignId Campaign ID as UUID
 @return ApiListCampaignVulnerabilitiesRequest
*/
func (a *CampaignsAPIService) ListCampaignVulnerabilities(ctx context.Context, campaignId string) ApiListCampaignVulnerabilitiesRequest {
	return ApiListCampaignVulnerabilitiesRequest{
		ApiService: a,
		ctx: ctx,
		campaignId: campaignId,
	}
}

// Execute executes the request
//  @return PaginatedVulnerability
func (a *CampaignsAPIService) ListCampaignVulnerabilitiesExecute(r ApiListCampaignVulnerabilitiesRequest) (*PaginatedVulnerability, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *PaginatedVulnerability
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "CampaignsAPIService.ListCampaignVulnerabilities")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/campaigns/{campaign_id}/vulnerabilities"
	localVarPath = strings.Replace(localVarPath, "{"+"campaign_id"+"}", url.PathEscape(parameterValueToString(r.campaignId, "campaignId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.size != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "size", r.size, "form", "")
	}
	if r.page != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "page", r.page, "form", "")
	}
	if r.orderBy != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "order_by", r.orderBy, "form", "")
	}
	if r.status != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "status", r.status, "form", "csv")
	}
	if r.firstDetectedAt != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "first_detected_at", r.firstDetectedAt, "form", "")
	}
	if r.lastDetectedAt != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "last_detected_at", r.lastDetectedAt, "form", "")
	}
	if r.resourceKind != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "resource.kind", r.resourceKind, "form", "csv")
	}
	if r.issueSeverity != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "issue.severity", r.issueSeverity, "form", "csv")
	}
	if r.issueId != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "issue.id", r.issueId, "form", "csv")
	}
	if r.issueCategoryId != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "issue.category.id", r.issueCategoryId, "form", "csv")
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
			if apiKey, ok := auth["BearerAuth"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarHeaderParams["Authorization"] = key
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
			var v InvalidInputError
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 401 {
			var v AuthError
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
			var v AccessDeniedError
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

type ApiListCampaignsRequest struct {
	ctx context.Context
	ApiService *CampaignsAPIService
	size *int32
	page *int32
	status *string
}

// Results per page
func (r ApiListCampaignsRequest) Size(size int32) ApiListCampaignsRequest {
	r.size = &size
	return r
}

// Page number
func (r ApiListCampaignsRequest) Page(page int32) ApiListCampaignsRequest {
	r.page = &page
	return r
}

// Filters for campaigns matching the given status.
func (r ApiListCampaignsRequest) Status(status string) ApiListCampaignsRequest {
	r.status = &status
	return r
}

func (r ApiListCampaignsRequest) Execute() (*PaginatedCampaign, *http.Response, error) {
	return r.ApiService.ListCampaignsExecute(r)
}

/*
ListCampaigns List campaigns

List campaigns

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiListCampaignsRequest
*/
func (a *CampaignsAPIService) ListCampaigns(ctx context.Context) ApiListCampaignsRequest {
	return ApiListCampaignsRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return PaginatedCampaign
func (a *CampaignsAPIService) ListCampaignsExecute(r ApiListCampaignsRequest) (*PaginatedCampaign, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *PaginatedCampaign
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "CampaignsAPIService.ListCampaigns")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/campaigns"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.size != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "size", r.size, "form", "")
	}
	if r.page != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "page", r.page, "form", "")
	}
	if r.status != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "status", r.status, "form", "")
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
			if apiKey, ok := auth["BearerAuth"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarHeaderParams["Authorization"] = key
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
			var v InvalidInputError
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 401 {
			var v AuthError
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
			var v AccessDeniedError
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 422 {
			var v ValidationListError
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