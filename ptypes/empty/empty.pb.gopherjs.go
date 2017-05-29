// Code generated by protoc-gen-gopherjs. DO NOT EDIT.
// source: empty/empty.proto

/*
Package empty is a generated protocol buffer package.

It is generated from these files:
	empty/empty.proto

It has these top-level messages:
	Empty
*/
package empty

import js "github.com/gopherjs/gopherjs/js"
import grpcweb "github.com/johanbrandhorst/gopherjs-improbable-grpc-web"

// A generic empty message that you can re-use to avoid defining duplicated
// empty messages in your APIs. A typical example is to use it as the request
// or the response type of an API method. For instance:
//
//     service Foo {
//       rpc Bar(google.protobuf.Empty) returns (google.protobuf.Empty);
//     }
//
// The JSON representation for `Empty` is empty JSON object `{}`.
type Empty struct {
	*js.Object
}

// NewEmpty creates a new Empty.
func NewEmpty() *Empty {
	m := &Empty{
		Object: js.Global.Get("proto").Get("google").Get("protobuf").Get("Empty").New([]interface{}{}),
	}

	return m
}

func (m *Empty) serialize() (rawBytes []byte, err error) {
	return grpcweb.Serialize(m)
}

func deserializeEmpty(rawBytes []byte) (*Empty, error) {
	obj, err := grpcweb.Deserialize(js.Global.Get("proto").Get("google").Get("protobuf").Get("Empty"), rawBytes)
	if err != nil {
		return nil, err
	}

	return &Empty{
		Object: obj,
	}, nil
}
