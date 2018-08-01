/*
 * Copyright 2018 The ThunderDB Authors.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package sqlchain

import (
	"sync"

	"gitlab.com/thunderdb/ThunderDB/proto"
	"gitlab.com/thunderdb/ThunderDB/rpc"
)

// MuxService defines multiplexing service of sql-chain.
type MuxService struct {
	ServiceName string
	serviceMap  sync.Map
}

// NewMuxService creates a new multiplexing service and registers it to rpc server.
func NewMuxService(serviceName string, server *rpc.Server) (service *MuxService) {
	service = &MuxService{
		ServiceName: serviceName,
	}

	server.RegisterService(serviceName, service)
	return service
}

func (s *MuxService) register(id proto.DatabaseID, service *ChainRPCService) {
	s.serviceMap.Store(id, service)
}

func (s *MuxService) unregister(id proto.DatabaseID) {
	s.serviceMap.Delete(id)
}

// MuxAdviseNewBlockReq defines a request of the AdviseNewBlock RPC method.
type MuxAdviseNewBlockReq struct {
	proto.Envelope
	proto.DatabaseID
	AdviseNewBlockReq
}

// MuxAdviseNewBlockResp defines a response of the AdviseNewBlock RPC method.
type MuxAdviseNewBlockResp struct {
	proto.Envelope
	proto.DatabaseID
	AdviseNewBlockResp
}

// MuxAdviseBinLogReq defines a request of the AdviseBinLog RPC method.
type MuxAdviseBinLogReq struct {
	proto.Envelope
	proto.DatabaseID
	AdviseBinLogReq
}

// MuxAdviseBinLogResp defines a response of the AdviseBinLog RPC method.
type MuxAdviseBinLogResp struct {
	proto.Envelope
	proto.DatabaseID
	AdviseBinLogResp
}

// MuxAdviseResponsedQueryReq defines a request of the AdviseAckedQuery RPC method.
type MuxAdviseResponsedQueryReq struct {
	proto.Envelope
	proto.DatabaseID
	AdviseResponsedQueryReq
}

// MuxAdviseResponsedQueryResp defines a response of the AdviseAckedQuery RPC method.
type MuxAdviseResponsedQueryResp struct {
	proto.Envelope
	proto.DatabaseID
	AdviseResponsedQueryResp
}

// MuxAdviseAckedQueryReq defines a request of the AdviseAckedQuery RPC method.
type MuxAdviseAckedQueryReq struct {
	proto.Envelope
	proto.DatabaseID
	AdviseAckedQueryReq
}

// MuxAdviseAckedQueryResp defines a response of the AdviseAckedQuery RPC method.
type MuxAdviseAckedQueryResp struct {
	proto.Envelope
	proto.DatabaseID
	AdviseAckedQueryResp
}

// MuxFetchBlockReq defines a request of the FetchBlock RPC method.
type MuxFetchBlockReq struct {
	proto.Envelope
	proto.DatabaseID
	FetchBlockReq
}

// MuxFetchBlockResp defines a response of the FetchBlock RPC method.
type MuxFetchBlockResp struct {
	proto.Envelope
	proto.DatabaseID
	FetchBlockResp
}

// MuxFetchAckedQueryReq defines a request of the FetchAckedQuery RPC method.
type MuxFetchAckedQueryReq struct {
	proto.Envelope
	proto.DatabaseID
	FetchAckedQueryReq
}

// MuxFetchAckedQueryResp defines a request of the FetchAckedQuery RPC method.
type MuxFetchAckedQueryResp struct {
	proto.Envelope
	proto.DatabaseID
	FetchAckedQueryResp
}

// MuxSignBillingReq defines a request of the SignBilling RPC method.
type MuxSignBillingReq struct {
	proto.Envelope
	proto.DatabaseID
	SignBillingReq
}

// MuxSignBillingResp defines a response of the SignBilling RPC method.
type MuxSignBillingResp struct {
	proto.Envelope
	proto.DatabaseID
	SignBillingResp
}

// AdviseNewBlock is the RPC method to advise a new produced block to the target server.
func (s *MuxService) AdviseNewBlock(req *MuxAdviseNewBlockReq, resp *MuxAdviseNewBlockResp) error {
	if v, ok := s.serviceMap.Load(req.DatabaseID); ok {
		resp.Envelope = req.Envelope
		resp.DatabaseID = req.DatabaseID
		return v.(*ChainRPCService).AdviseNewBlock(&req.AdviseNewBlockReq, &resp.AdviseNewBlockResp)
	}

	return ErrUnknownMuxRequest
}

// AdviseBinLog is the RPC method to advise a new binary log to the target server.
func (s *MuxService) AdviseBinLog(req *MuxAdviseBinLogReq, resp *MuxAdviseBinLogResp) error {
	if v, ok := s.serviceMap.Load(req.DatabaseID); ok {
		resp.Envelope = req.Envelope
		resp.DatabaseID = req.DatabaseID
		return v.(*ChainRPCService).AdviseBinLog(&req.AdviseBinLogReq, &resp.AdviseBinLogResp)
	}

	return ErrUnknownMuxRequest
}

// AdviseResponsedQuery is the RPC method to advise a new responsed query to the target server.
func (s *MuxService) AdviseResponsedQuery(
	req *MuxAdviseResponsedQueryReq, resp *MuxAdviseResponsedQueryResp) error {
	if v, ok := s.serviceMap.Load(req.DatabaseID); ok {
		resp.Envelope = req.Envelope
		resp.DatabaseID = req.DatabaseID
		return v.(*ChainRPCService).AdviseResponsedQuery(
			&req.AdviseResponsedQueryReq, &resp.AdviseResponsedQueryResp)
	}

	return ErrUnknownMuxRequest
}

// AdviseAckedQuery is the RPC method to advise a new acknowledged query to the target server.
func (s *MuxService) AdviseAckedQuery(
	req *MuxAdviseAckedQueryReq, resp *MuxAdviseAckedQueryResp) error {
	if v, ok := s.serviceMap.Load(req.DatabaseID); ok {
		resp.Envelope = req.Envelope
		resp.DatabaseID = req.DatabaseID
		return v.(*ChainRPCService).AdviseAckedQuery(
			&req.AdviseAckedQueryReq, &resp.AdviseAckedQueryResp)
	}

	return ErrUnknownMuxRequest
}

// FetchBlock is the RPC method to fetch a known block form the target server.
func (s *MuxService) FetchBlock(req *MuxFetchBlockReq, resp *MuxFetchBlockResp) (err error) {
	if v, ok := s.serviceMap.Load(req.DatabaseID); ok {
		resp.Envelope = req.Envelope
		resp.DatabaseID = req.DatabaseID
		return v.(*ChainRPCService).FetchBlock(&req.FetchBlockReq, &resp.FetchBlockResp)
	}

	return ErrUnknownMuxRequest
}

// FetchAckedQuery is the RPC method to fetch a known block form the target server.
func (s *MuxService) FetchAckedQuery(
	req *MuxFetchAckedQueryReq, resp *MuxFetchAckedQueryResp) (err error) {
	if v, ok := s.serviceMap.Load(req.DatabaseID); ok {
		resp.Envelope = req.Envelope
		resp.DatabaseID = req.DatabaseID
		return v.(*ChainRPCService).FetchAckedQuery(
			&req.FetchAckedQueryReq, &resp.FetchAckedQueryResp)
	}

	return ErrUnknownMuxRequest
}

// SignBilling is the RPC method to get signature for a billing request form the target server.
func (s *MuxService) SignBilling(req *MuxSignBillingReq, resp *MuxSignBillingResp) (err error) {
	if v, ok := s.serviceMap.Load(req.DatabaseID); ok {
		resp.Envelope = req.Envelope
		resp.DatabaseID = req.DatabaseID
		return v.(*ChainRPCService).SignBilling(&req.SignBillingReq, &resp.SignBillingResp)
	}

	return ErrUnknownMuxRequest
}