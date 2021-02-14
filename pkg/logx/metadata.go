package logx

import (
	"github.com/peractio/gdk/pkg/build"
	"github.com/peractio/gdk/pkg/errorx/v2"
	"github.com/peractio/gdk/pkg/tags"
)

// ErrMetadata returns a metadata constructed from error.
// If you are printing a log, this metadata should always be set.
// Useful if you want to trace log with certain information.
func ErrMetadata(err error) map[string]interface{} {
	if err == nil {
		return nil
	}

	metadata := map[string]interface{}{}
	if build.Info() != nil {
		metadata[tags.Build] = build.Info()
	}

	if e, ok := err.(*errorx.Error); ok {
		detail := map[string]interface{}{
			tags.ErrorLine: e.Line,
		}

		if len(e.OpTraces) > 0 {
			detail[tags.Ops] = e.OpTraces
		}

		if len(e.GetFields()) > 0 {
			detail[tags.Fields] = e.GetFields()
		}

		if e.Code != "" {
			detail[tags.Code] = e.Code
		}

		if e.MetricStatus != "" {
			detail[tags.MetricStatus] = e.MetricStatus
		}

		if e.Message != "" {
			detail[tags.Message] = e.Message
		}

		metadata[tags.Detail] = detail
	}

	return metadata
}

// Metadata returns a basic metadata added a more detail metadata.
func Metadata(detail interface{}) map[string]interface{} {
	metadata := map[string]interface{}{}
	if build.Info() != nil {
		metadata[tags.Build] = build.Info()
	}

	if detail != nil {
		metadata[tags.Detail] = detail
	}

	return metadata
}
