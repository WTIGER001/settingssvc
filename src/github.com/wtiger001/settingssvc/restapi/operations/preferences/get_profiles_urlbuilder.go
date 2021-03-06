// Code generated by go-swagger; DO NOT EDIT.

package preferences

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"errors"
	"net/url"
	golangswaggerpaths "path"

	"github.com/go-openapi/swag"
)

// GetProfilesURL generates an URL for the get profiles operation
type GetProfilesURL struct {
	ID      []string
	Ownerid *string

	_basePath string
	// avoid unkeyed usage
	_ struct{}
}

// WithBasePath sets the base path for this url builder, only required when it's different from the
// base path specified in the swagger spec.
// When the value of the base path is an empty string
func (o *GetProfilesURL) WithBasePath(bp string) *GetProfilesURL {
	o.SetBasePath(bp)
	return o
}

// SetBasePath sets the base path for this url builder, only required when it's different from the
// base path specified in the swagger spec.
// When the value of the base path is an empty string
func (o *GetProfilesURL) SetBasePath(bp string) {
	o._basePath = bp
}

// Build a url path and query string
func (o *GetProfilesURL) Build() (*url.URL, error) {
	var result url.URL

	var _path = "/profiles"

	_basePath := o._basePath
	if _basePath == "" {
		_basePath = "/"
	}
	result.Path = golangswaggerpaths.Join(_basePath, _path)

	qs := make(url.Values)

	var iDIR []string
	for _, iDI := range o.ID {
		iDIS := iDI
		if iDIS != "" {
			iDIR = append(iDIR, iDIS)
		}
	}

	id := swag.JoinByFormat(iDIR, "")

	if len(id) > 0 {
		qsv := id[0]
		if qsv != "" {
			qs.Set("id", qsv)
		}
	}

	var ownerid string
	if o.Ownerid != nil {
		ownerid = *o.Ownerid
	}
	if ownerid != "" {
		qs.Set("ownerid", ownerid)
	}

	result.RawQuery = qs.Encode()

	return &result, nil
}

// Must is a helper function to panic when the url builder returns an error
func (o *GetProfilesURL) Must(u *url.URL, err error) *url.URL {
	if err != nil {
		panic(err)
	}
	if u == nil {
		panic("url can't be nil")
	}
	return u
}

// String returns the string representation of the path with query string
func (o *GetProfilesURL) String() string {
	return o.Must(o.Build()).String()
}

// BuildFull builds a full url with scheme, host, path and query string
func (o *GetProfilesURL) BuildFull(scheme, host string) (*url.URL, error) {
	if scheme == "" {
		return nil, errors.New("scheme is required for a full url on GetProfilesURL")
	}
	if host == "" {
		return nil, errors.New("host is required for a full url on GetProfilesURL")
	}

	base, err := o.Build()
	if err != nil {
		return nil, err
	}

	base.Scheme = scheme
	base.Host = host
	return base, nil
}

// StringFull returns the string representation of a complete url
func (o *GetProfilesURL) StringFull(scheme, host string) string {
	return o.Must(o.BuildFull(scheme, host)).String()
}
