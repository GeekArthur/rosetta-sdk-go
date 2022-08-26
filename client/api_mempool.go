// Copyright 2022 Coinbase, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Generated by: OpenAPI Generator (https://openapi-generator.tech)

package client

import (
	_context "context"
	"fmt"
	_ioutil "io/ioutil"
	_nethttp "net/http"

	"github.com/coinbase/rosetta-sdk-go/types"
)

// Linger please
var (
	_ _context.Context
)

// MempoolAPIService MempoolAPI service
type MempoolAPIService service

// Mempool Get all Transaction Identifiers in the mempool
func (a *MempoolAPIService) Mempool(
	ctx _context.Context,
	networkRequest *types.NetworkRequest,
) (*types.MempoolResponse, *types.Error, error) {
	var (
		localVarPostBody interface{}
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/mempool"
	localVarHeaderParams := make(map[string]string)

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
	localVarPostBody = networkRequest

	r, err := a.client.prepareRequest(ctx, localVarPath, localVarPostBody, localVarHeaderParams)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to prepare request: %w", err)
	}

	localVarHTTPResponse, err := a.client.callAPI(ctx, r)
	if err != nil || localVarHTTPResponse == nil {
		return nil, nil, fmt.Errorf("failed to call API: %w", err)
	}

	localVarBody, err := _ioutil.ReadAll(localVarHTTPResponse.Body)
	defer localVarHTTPResponse.Body.Close()
	if err != nil {
		return nil, nil, fmt.Errorf("failed to read response: %w", err)
	}

	switch localVarHTTPResponse.StatusCode {
	case _nethttp.StatusOK:
		var v types.MempoolResponse
		err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
		if err != nil {
			return nil, nil, fmt.Errorf("failed to decode when hit status code 200, response body %s: %w", string(localVarBody), err)
		}

		return &v, nil, nil
	case _nethttp.StatusInternalServerError:
		var v types.Error
		err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
		if err != nil {
			return nil, nil, fmt.Errorf("failed to decode when hit status code 500, response body %s: %w", string(localVarBody), err)
		}

		return nil, &v, fmt.Errorf("error %+v", v)
	case _nethttp.StatusBadGateway,
		_nethttp.StatusServiceUnavailable,
		_nethttp.StatusGatewayTimeout,
		_nethttp.StatusRequestTimeout:
		return nil, nil, fmt.Errorf(
			"status code %d, response body %s: %w",
			localVarHTTPResponse.StatusCode,
			string(localVarBody),
			ErrRetriable,
		)
	default:
		return nil, nil, fmt.Errorf(
			"invalid status code %d, response body %s",
			localVarHTTPResponse.StatusCode,
			string(localVarBody),
		)
	}
}

// MempoolTransaction Get a transaction in the mempool by its Transaction Identifier. This is a
// separate request than fetching a block transaction (/block/transaction) because some blockchain
// nodes need to know that a transaction query is for something in the mempool instead of a
// transaction in a block. Transactions may not be fully parsable until they are in a block (ex: may
// not be possible to determine the fee to pay before a transaction is executed). On this endpoint,
// it is ok that returned transactions are only estimates of what may actually be included in a
// block.
func (a *MempoolAPIService) MempoolTransaction(
	ctx _context.Context,
	mempoolTransactionRequest *types.MempoolTransactionRequest,
) (*types.MempoolTransactionResponse, *types.Error, error) {
	var (
		localVarPostBody interface{}
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/mempool/transaction"
	localVarHeaderParams := make(map[string]string)

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
	localVarPostBody = mempoolTransactionRequest

	r, err := a.client.prepareRequest(ctx, localVarPath, localVarPostBody, localVarHeaderParams)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to prepare request: %w", err)
	}

	localVarHTTPResponse, err := a.client.callAPI(ctx, r)
	if err != nil || localVarHTTPResponse == nil {
		return nil, nil, fmt.Errorf("failed to call API: %w", err)
	}

	localVarBody, err := _ioutil.ReadAll(localVarHTTPResponse.Body)
	defer localVarHTTPResponse.Body.Close()
	if err != nil {
		return nil, nil, fmt.Errorf("failed to read response: %w", err)
	}

	switch localVarHTTPResponse.StatusCode {
	case _nethttp.StatusOK:
		var v types.MempoolTransactionResponse
		err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
		if err != nil {
			return nil, nil, fmt.Errorf("failed to decode when hit status code 200, response body %s: %w", string(localVarBody), err)
		}

		return &v, nil, nil
	case _nethttp.StatusInternalServerError:
		var v types.Error
		err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
		if err != nil {
			return nil, nil, fmt.Errorf("failed to decode when hit status code 500, response body %s: %w", string(localVarBody), err)
		}

		return nil, &v, fmt.Errorf("error %+v", v)
	case _nethttp.StatusBadGateway,
		_nethttp.StatusServiceUnavailable,
		_nethttp.StatusGatewayTimeout,
		_nethttp.StatusRequestTimeout:
		return nil, nil, fmt.Errorf(
			"status code %d, response body %s: %w",
			localVarHTTPResponse.StatusCode,
			string(localVarBody),
			ErrRetriable,
		)
	default:
		return nil, nil, fmt.Errorf(
			"invalid status code %d, response body %s",
			localVarHTTPResponse.StatusCode,
			string(localVarBody),
		)
	}
}
