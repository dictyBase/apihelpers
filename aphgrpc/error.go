// Package aphgrpc provides various interfaces, functions, types
// for building and working with gRPC services.
package aphgrpc

import (
	"encoding/json"
	"net/http"
	"strconv"
	"strings"

	context "golang.org/x/net/context"

	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"github.com/pkg/errors"
	spb "google.golang.org/genproto/googleapis/rpc/status"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

const (
	// MetaKey is the key used for storing all metadata
	MetaKey = "error"
)

var (
	//ErrDatabaseQuery represents database query related errors
	ErrDatabaseQuery = newError("Database query error")
	//ErrDatabaseInsert represents database insert related errors
	ErrDatabaseInsert = newError("Database insert error")
	//ErrDatabaseUpdate represents database update related errors
	ErrDatabaseUpdate = newError("Database update error")
	//ErrDatabaseDelete represents database update delete errors
	ErrDatabaseDelete = newError("Database delete error")
	//ErrNotFound represents the absence of an HTTP resource
	ErrNotFound = newError("Resource not found")
	//ErrExists represents the presence of an HTTP resource
	ErrExists = newError("Resource already exists")
	//ErrJSONEncoding represents any json encoding error
	ErrJSONEncoding = newError("Json encoding error")
	//ErrStructMarshal represents any error with marshalling structure
	ErrStructMarshal = newError("Structure marshalling error")
	//ErrIncludeParam represents any error with invalid include query parameter
	ErrIncludeParam = newErrorWithParam("Invalid include query parameter", "include")
	//ErrSparseFieldSets represents any error with invalid sparse fieldsets query parameter
	ErrFields = newErrorWithParam("Invalid field query parameter", "field")
	//ErrFilterParam represents any error with invalid filter query paramter
	ErrFilterParam = newErrorWithParam("Invalid filter query parameter", "filter")
	//ErrNotAcceptable represents any error with wrong or inappropriate http Accept header
	ErrNotAcceptable = newError("Accept header is not acceptable")
	//ErrUnsupportedMedia represents any error with unsupported media type in http header
	ErrUnsupportedMedia = newError("Media type is not supported")
	//ErrInValidParam represents any error with validating input parameters
	ErrInValidParam = newError("Invalid parameters")
	//ErrRetrieveMetadata represents any error to retrieve grpc metadata from the running context
	ErrRetrieveMetadata = errors.New("unable to retrieve metadata")
	//ErrXForwardedHost represents any failure or absence of x-forwarded-host HTTP header in the grpc context
	ErrXForwardedHost = errors.New("x-forwarded-host header is absent")
)

// HTTPError is used for errors
type HTTPError struct {
	err    error
	msg    string
	status int
	Errors []Error `json:"errors,omitempty"`
}

// Error can be used for all kind of application errors
// e.g. you would use it to define form errors or any
// other semantical application problems
// for more information see http://jsonapi.org/format/#errors
type Error struct {
	ID     string       `json:"id,omitempty"`
	Links  *ErrorLinks  `json:"links,omitempty"`
	Status string       `json:"status,omitempty"`
	Code   string       `json:"code,omitempty"`
	Title  string       `json:"title,omitempty"`
	Detail string       `json:"detail,omitempty"`
	Source *ErrorSource `json:"source,omitempty"`
	Meta   interface{}  `json:"meta,omitempty"`
}

// ErrorLinks is used to provide an About URL that leads to
// further details about the particular occurrence of the problem.
//
// for more information see http://jsonapi.org/format/#error-objects
type ErrorLinks struct {
	About string `json:"about,omitempty"`
}

// ErrorSource is used to provide references to the source of an error.
//
// The Pointer is a JSON Pointer to the associated entity in the request
// document.
// The Paramter is a string indicating which query parameter caused the error.
//
// for more information see http://jsonapi.org/format/#error-objects
type ErrorSource struct {
	Pointer   string `json:"pointer,omitempty"`
	Parameter string `json:"parameter,omitempty"`
}

func newErrorWithParam(msg, param string) metadata.MD {
	return metadata.Pairs(MetaKey, msg, MetaKey, param)
}

func newError(msg string) metadata.MD {
	return metadata.Pairs(MetaKey, msg)
}

// CustomHTTPError is a custom error handler for grpc-gateway to generate
// JSONAPI formatted HTTP response.
func CustomHTTPError(ctx context.Context, _ *runtime.ServeMux, marshaler runtime.Marshaler, w http.ResponseWriter, _ *http.Request, err error) {
	md, ok := runtime.ServerMetadataFromContext(ctx)
	if !ok {
		fallbackError(w, getgRPCStatus(errors.Wrap(err, "unable to retrieve metadata")))
		return
	}
	JSONAPIError(w, md.TrailerMD, getgRPCStatus(err))
}

func getgRPCStatus(err error) *status.Status {
	s, ok := status.FromError(err)
	if !ok {
		return status.New(codes.Unknown, err.Error())
	}
	return s
}

// JSONAPIError generates JSONAPI formatted error message
func JSONAPIError(w http.ResponseWriter, md metadata.MD, s *status.Status) {
	status := runtime.HTTPStatusFromCode(s.Code())
	jsnErr := Error{
		Status: strconv.Itoa(status),
		Title:  strings.Join(md["error"], "-"),
		Detail: s.Message(),
		Meta: map[string]interface{}{
			"creator": "api error helper",
		},
	}
	w.Header().Set("Content-Type", "application/vnd.api+json")
	w.WriteHeader(status)
	encErr := json.NewEncoder(w).Encode(HTTPError{Errors: []Error{jsnErr}})
	if encErr != nil {
		http.Error(w, encErr.Error(), http.StatusInternalServerError)
	}
}

func fallbackError(w http.ResponseWriter, s *status.Status) {
	status := runtime.HTTPStatusFromCode(s.Code())
	jsnErr := Error{
		Status: strconv.Itoa(status),
		Title:  "gRPC error",
		Detail: s.Message(),
		Meta: map[string]interface{}{
			"creator": "api error helper",
		},
	}
	w.Header().Set("Content-Type", "application/vnd.api+json")
	w.WriteHeader(status)
	encErr := json.NewEncoder(w).Encode(HTTPError{Errors: []Error{jsnErr}})
	if encErr != nil {
		http.Error(w, encErr.Error(), http.StatusInternalServerError)
	}
}

func CheckNoRows(err error) bool {
	if strings.Contains(err.Error(), "no rows") {
		return true
	}
	return false
}

func HandleMessagingError(ctx context.Context, st *spb.Status) error {
	err := status.ErrorProto(st)
	grpc.SetTrailer(ctx, newError(err.Error()))
	return err
}

func HandleError(ctx context.Context, err error) error {
	if CheckNoRows(err) {
		grpc.SetTrailer(ctx, ErrNotFound)
		return status.Error(codes.NotFound, err.Error())
	}
	grpc.SetTrailer(ctx, newError(err.Error()))
	return status.Error(codes.Internal, err.Error())
}

func HandleGenericError(ctx context.Context, err error) error {
	grpc.SetTrailer(ctx, newError(err.Error()))
	return status.Error(codes.Internal, err.Error())
}

func HandleDeleteError(ctx context.Context, err error) error {
	grpc.SetTrailer(ctx, ErrDatabaseDelete)
	return status.Error(codes.Internal, err.Error())
}

func HandleGetError(ctx context.Context, err error) error {
	grpc.SetTrailer(ctx, ErrDatabaseQuery)
	return status.Error(codes.Internal, err.Error())
}

func HandleInsertError(ctx context.Context, err error) error {
	grpc.SetTrailer(ctx, ErrDatabaseInsert)
	return status.Error(codes.Internal, err.Error())
}

func HandleUpdateError(ctx context.Context, err error) error {
	grpc.SetTrailer(ctx, ErrDatabaseUpdate)
	return status.Error(codes.Internal, err.Error())
}

func HandleGetArgError(ctx context.Context, err error) error {
	grpc.SetTrailer(ctx, ErrDatabaseQuery)
	return status.Error(codes.InvalidArgument, err.Error())
}

func HandleInsertArgError(ctx context.Context, err error) error {
	grpc.SetTrailer(ctx, ErrDatabaseInsert)
	return status.Error(codes.InvalidArgument, err.Error())
}

func HandleUpdateArgError(ctx context.Context, err error) error {
	grpc.SetTrailer(ctx, ErrDatabaseUpdate)
	return status.Error(codes.InvalidArgument, err.Error())
}

func HandleNotFoundError(ctx context.Context, err error) error {
	grpc.SetTrailer(ctx, ErrNotFound)
	return status.Error(codes.NotFound, err.Error())
}

func HandleExistError(ctx context.Context, err error) error {
	grpc.SetTrailer(ctx, ErrExists)
	return status.Error(codes.AlreadyExists, err.Error())
}

func HandleFilterParamError(ctx context.Context, err error) error {
	grpc.SetTrailer(ctx, ErrFilterParam)
	return status.Error(codes.InvalidArgument, err.Error())
}

func HandleInvalidParamError(ctx context.Context, err error) error {
	grpc.SetTrailer(ctx, ErrInValidParam)
	return status.Error(codes.InvalidArgument, err.Error())
}
