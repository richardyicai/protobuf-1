// Code generated by protoc-gen-gopherjs. DO NOT EDIT.
// source: any/any.proto

/*
Package any is a generated protocol buffer package.

It is generated from these files:
	any/any.proto

It has these top-level messages:
	Any
*/
package any

import js "github.com/gopherjs/gopherjs/js"
import grpcweb "github.com/johanbrandhorst/gopherjs-improbable-grpc-web"

// `Any` contains an arbitrary serialized protocol buffer message along with a
// URL that describes the type of the serialized message.
//
// Protobuf library provides support to pack/unpack Any values in the form
// of utility functions or additional generated methods of the Any type.
//
// Example 1: Pack and unpack a message in C++.
//
//     Foo foo = ...;
//     Any any;
//     any.PackFrom(foo);
//     ...
//     if (any.UnpackTo(&foo)) {
//       ...
//     }
//
// Example 2: Pack and unpack a message in Java.
//
//     Foo foo = ...;
//     Any any = Any.pack(foo);
//     ...
//     if (any.is(Foo.class)) {
//       foo = any.unpack(Foo.class);
//     }
//
//  Example 3: Pack and unpack a message in Python.
//
//     foo = Foo(...)
//     any = Any()
//     any.Pack(foo)
//     ...
//     if any.Is(Foo.DESCRIPTOR):
//       any.Unpack(foo)
//       ...
//
// The pack methods provided by protobuf library will by default use
// 'type.googleapis.com/full.type.name' as the type URL and the unpack
// methods only use the fully qualified type name after the last '/'
// in the type URL, for example "foo.bar.com/x/y.z" will yield type
// name "y.z".
//
//
// JSON
// ====
// The JSON representation of an `Any` value uses the regular
// representation of the deserialized, embedded message, with an
// additional field `@type` which contains the type URL. Example:
//
//     package google.profile;
//     message Person {
//       string first_name = 1;
//       string last_name = 2;
//     }
//
//     {
//       "@type": "type.googleapis.com/google.profile.Person",
//       "firstName": <string>,
//       "lastName": <string>
//     }
//
// If the embedded message type is well-known and has a custom JSON
// representation, that representation will be embedded adding a field
// `value` which holds the custom JSON in addition to the `@type`
// field. Example (for message [google.protobuf.Duration][]):
//
//     {
//       "@type": "type.googleapis.com/google.protobuf.Duration",
//       "value": "1.212s"
//     }
//
type Any struct {
	*js.Object
}

// NewAny creates a new Any.
// A URL/resource name whose content describes the type of the
// serialized protocol buffer message.
//
// For URLs which use the scheme `http`, `https`, or no scheme, the
// following restrictions and interpretations apply:
//
// * If no scheme is provided, `https` is assumed.
// * The last segment of the URL's path must represent the fully
//   qualified name of the type (as in `path/google.protobuf.Duration`).
//   The name should be in a canonical form (e.g., leading "." is
//   not accepted).
// * An HTTP GET on the URL must yield a [google.protobuf.Type][]
//   value in binary format, or produce an error.
// * Applications are allowed to cache lookup results based on the
//   URL, or have them precompiled into a binary to avoid any
//   lookup. Therefore, binary compatibility needs to be preserved
//   on changes to types. (Use versioned type names to manage
//   breaking changes.)
//
// Schemes other than `http`, `https` (or the empty scheme) might be
// used with implementation specific semantics.
//
// Must be a valid serialized protocol buffer of the above specified type.
func NewAny(typeUrl string, value []byte) *Any {
	m := &Any{
		Object: js.Global.Get("proto").Get("google").Get("protobuf").Get("Any").New([]interface{}{
			typeUrl,
			value,
		}),
	}

	return m
}

// GetTypeUrl gets the TypeUrl of the Any.
// A URL/resource name whose content describes the type of the
// serialized protocol buffer message.
//
// For URLs which use the scheme `http`, `https`, or no scheme, the
// following restrictions and interpretations apply:
//
// * If no scheme is provided, `https` is assumed.
// * The last segment of the URL's path must represent the fully
//   qualified name of the type (as in `path/google.protobuf.Duration`).
//   The name should be in a canonical form (e.g., leading "." is
//   not accepted).
// * An HTTP GET on the URL must yield a [google.protobuf.Type][]
//   value in binary format, or produce an error.
// * Applications are allowed to cache lookup results based on the
//   URL, or have them precompiled into a binary to avoid any
//   lookup. Therefore, binary compatibility needs to be preserved
//   on changes to types. (Use versioned type names to manage
//   breaking changes.)
//
// Schemes other than `http`, `https` (or the empty scheme) might be
// used with implementation specific semantics.
//
func (m *Any) GetTypeUrl() string {
	return m.Call("getTypeUrl").String()
}

// SetTypeUrl sets the TypeUrl of the Any.
// A URL/resource name whose content describes the type of the
// serialized protocol buffer message.
//
// For URLs which use the scheme `http`, `https`, or no scheme, the
// following restrictions and interpretations apply:
//
// * If no scheme is provided, `https` is assumed.
// * The last segment of the URL's path must represent the fully
//   qualified name of the type (as in `path/google.protobuf.Duration`).
//   The name should be in a canonical form (e.g., leading "." is
//   not accepted).
// * An HTTP GET on the URL must yield a [google.protobuf.Type][]
//   value in binary format, or produce an error.
// * Applications are allowed to cache lookup results based on the
//   URL, or have them precompiled into a binary to avoid any
//   lookup. Therefore, binary compatibility needs to be preserved
//   on changes to types. (Use versioned type names to manage
//   breaking changes.)
//
// Schemes other than `http`, `https` (or the empty scheme) might be
// used with implementation specific semantics.
//
func (m *Any) SetTypeUrl(v string) {
	m.Call("setTypeUrl", v)
}

// GetValue gets the Value of the Any.
// Must be a valid serialized protocol buffer of the above specified type.
func (m *Any) GetValue() []byte {
	return m.Call("getValue").Interface().([]byte)
}

// SetValue sets the Value of the Any.
// Must be a valid serialized protocol buffer of the above specified type.
func (m *Any) SetValue(v []byte) {
	m.Call("setValue", v)
}

func (m *Any) serialize() (rawBytes []byte, err error) {
	return grpcweb.Serialize(m)
}

func deserializeAny(rawBytes []byte) (*Any, error) {
	obj, err := grpcweb.Deserialize(js.Global.Get("proto").Get("google").Get("protobuf").Get("Any"), rawBytes)
	if err != nil {
		return nil, err
	}

	return &Any{
		Object: obj,
	}, nil
}
