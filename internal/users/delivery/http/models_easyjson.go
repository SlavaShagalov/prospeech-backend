// Code generated by easyjson for marshaling/unmarshaling. DO NOT EDIT.

package http

import (
	json "encoding/json"
	easyjson "github.com/mailru/easyjson"
	jlexer "github.com/mailru/easyjson/jlexer"
	jwriter "github.com/mailru/easyjson/jwriter"
)

// suppress unused package warning
var (
	_ *json.RawMessage
	_ *jlexer.Lexer
	_ *jwriter.Writer
	_ easyjson.Marshaler
)

func easyjsonD2b7633eDecodeGithubComSlavaShagalovProspeechBackendInternalUsersDeliveryHttp(in *jlexer.Lexer, out *partialUpdateRequest) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		if isTopLevel {
			in.Consumed()
		}
		in.Skip()
		return
	}
	in.Delim('{')
	for !in.IsDelim('}') {
		key := in.UnsafeFieldName(false)
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "username":
			if in.IsNull() {
				in.Skip()
				out.Username = nil
			} else {
				if out.Username == nil {
					out.Username = new(string)
				}
				*out.Username = string(in.String())
			}
		case "email":
			if in.IsNull() {
				in.Skip()
				out.Email = nil
			} else {
				if out.Email == nil {
					out.Email = new(string)
				}
				*out.Email = string(in.String())
			}
		case "name":
			if in.IsNull() {
				in.Skip()
				out.Name = nil
			} else {
				if out.Name == nil {
					out.Name = new(string)
				}
				*out.Name = string(in.String())
			}
		default:
			in.SkipRecursive()
		}
		in.WantComma()
	}
	in.Delim('}')
	if isTopLevel {
		in.Consumed()
	}
}
func easyjsonD2b7633eEncodeGithubComSlavaShagalovProspeechBackendInternalUsersDeliveryHttp(out *jwriter.Writer, in partialUpdateRequest) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"username\":"
		out.RawString(prefix[1:])
		if in.Username == nil {
			out.RawString("null")
		} else {
			out.String(string(*in.Username))
		}
	}
	{
		const prefix string = ",\"email\":"
		out.RawString(prefix)
		if in.Email == nil {
			out.RawString("null")
		} else {
			out.String(string(*in.Email))
		}
	}
	{
		const prefix string = ",\"name\":"
		out.RawString(prefix)
		if in.Name == nil {
			out.RawString("null")
		} else {
			out.String(string(*in.Name))
		}
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v partialUpdateRequest) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonD2b7633eEncodeGithubComSlavaShagalovProspeechBackendInternalUsersDeliveryHttp(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v partialUpdateRequest) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonD2b7633eEncodeGithubComSlavaShagalovProspeechBackendInternalUsersDeliveryHttp(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *partialUpdateRequest) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonD2b7633eDecodeGithubComSlavaShagalovProspeechBackendInternalUsersDeliveryHttp(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *partialUpdateRequest) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonD2b7633eDecodeGithubComSlavaShagalovProspeechBackendInternalUsersDeliveryHttp(l, v)
}
func easyjsonD2b7633eDecodeGithubComSlavaShagalovProspeechBackendInternalUsersDeliveryHttp1(in *jlexer.Lexer, out *getResponse) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		if isTopLevel {
			in.Consumed()
		}
		in.Skip()
		return
	}
	in.Delim('{')
	for !in.IsDelim('}') {
		key := in.UnsafeFieldName(false)
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "id":
			out.ID = int64(in.Int64())
		case "username":
			out.Username = string(in.String())
		case "email":
			out.Email = string(in.String())
		case "name":
			out.Name = string(in.String())
		case "avatar":
			if in.IsNull() {
				in.Skip()
				out.Avatar = nil
			} else {
				if out.Avatar == nil {
					out.Avatar = new(string)
				}
				*out.Avatar = string(in.String())
			}
		case "created_at":
			if data := in.Raw(); in.Ok() {
				in.AddError((out.CreatedAt).UnmarshalJSON(data))
			}
		case "updated_at":
			if data := in.Raw(); in.Ok() {
				in.AddError((out.UpdatedAt).UnmarshalJSON(data))
			}
		default:
			in.SkipRecursive()
		}
		in.WantComma()
	}
	in.Delim('}')
	if isTopLevel {
		in.Consumed()
	}
}
func easyjsonD2b7633eEncodeGithubComSlavaShagalovProspeechBackendInternalUsersDeliveryHttp1(out *jwriter.Writer, in getResponse) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"id\":"
		out.RawString(prefix[1:])
		out.Int64(int64(in.ID))
	}
	{
		const prefix string = ",\"username\":"
		out.RawString(prefix)
		out.String(string(in.Username))
	}
	{
		const prefix string = ",\"email\":"
		out.RawString(prefix)
		out.String(string(in.Email))
	}
	{
		const prefix string = ",\"name\":"
		out.RawString(prefix)
		out.String(string(in.Name))
	}
	{
		const prefix string = ",\"avatar\":"
		out.RawString(prefix)
		if in.Avatar == nil {
			out.RawString("null")
		} else {
			out.String(string(*in.Avatar))
		}
	}
	{
		const prefix string = ",\"created_at\":"
		out.RawString(prefix)
		out.Raw((in.CreatedAt).MarshalJSON())
	}
	{
		const prefix string = ",\"updated_at\":"
		out.RawString(prefix)
		out.Raw((in.UpdatedAt).MarshalJSON())
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v getResponse) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonD2b7633eEncodeGithubComSlavaShagalovProspeechBackendInternalUsersDeliveryHttp1(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v getResponse) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonD2b7633eEncodeGithubComSlavaShagalovProspeechBackendInternalUsersDeliveryHttp1(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *getResponse) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonD2b7633eDecodeGithubComSlavaShagalovProspeechBackendInternalUsersDeliveryHttp1(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *getResponse) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonD2b7633eDecodeGithubComSlavaShagalovProspeechBackendInternalUsersDeliveryHttp1(l, v)
}
