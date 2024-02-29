/*
Copyright 2019 The Vitess Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package callinfo

// This file implements the CallInfo interface for gRPC contexts.

import (
	"context"
	"fmt"

	"github.com/google/safehtml"
	"github.com/google/safehtml/template"
	"google.golang.org/grpc"
	"google.golang.org/grpc/peer"
)

// GRPCCallInfo returns an augmented context with a CallInfo structure,
// only for gRPC contexts.
func GRPCCallInfo(ctx context.Context) context.Context {
	method, ok := grpc.Method(ctx)
	if !ok {
		return ctx
	}

	callinfo := &gRPCCallInfoImpl{
		method: method,
	}
	peer, ok := peer.FromContext(ctx)
	if ok {
		callinfo.remoteAddr = peer.Addr.String()
	}

	return NewContext(ctx, callinfo)
}

type gRPCCallInfoImpl struct {
	method     string
	remoteAddr string
}

func (gci *gRPCCallInfoImpl) RemoteAddr() string {
	return gci.remoteAddr
}

func (gci *gRPCCallInfoImpl) Username() string {
	return "gRPC"
}

func (gci *gRPCCallInfoImpl) Text() string {
	return fmt.Sprintf("%s:%s(gRPC)", gci.remoteAddr, gci.method)
}

var grpcTmpl = template.Must(template.New("tcs").Parse("<b>Method:</b> {{.Method}} <b>Remote Addr:</b> {{.RemoteAddr}}"))

func (gci *gRPCCallInfoImpl) HTML() safehtml.HTML {
	html, err := grpcTmpl.ExecuteToHTML(struct {
		Method     string
		RemoteAddr string
	}{
		Method:     gci.method,
		RemoteAddr: gci.remoteAddr,
	})
	if err != nil {
		panic(err)
	}
	return html
}
