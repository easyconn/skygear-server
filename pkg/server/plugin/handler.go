package plugin

import (
	"encoding/json"
	"io/ioutil"

	"github.com/sirupsen/logrus"

	"github.com/skygeario/skygear-server/pkg/server/router"
	"github.com/skygeario/skygear-server/pkg/server/skyerr"
)

type pluginRequestPayload struct {
	Status      int                 `json:"status"`
	Method      string              `json:"method,omitempty"`
	Header      map[string][]string `json:"header"`
	Body        []byte              `json:"body"`
	Path        string              `json:"path,omitempty"`
	QueryString string              `json:"query_string,omitempty"`
}

func (payload *pluginRequestPayload) Decode(input []byte) skyerr.Error {
	if err := json.Unmarshal(input, &payload); err != nil {
		return skyerr.NewError(
			skyerr.UnexpectedError,
			"plugin resposne malformat: "+err.Error(),
		)
	}
	return nil
}

func (payload *pluginRequestPayload) Encode() ([]byte, error) {
	return json.Marshal(payload)
}

type Handler struct {
	Plugin            *Plugin
	Name              string
	AccessKeyRequired bool
	UserRequired      bool
	PreprocessorList  router.PreprocessorRegistry
	preprocessors     []router.Processor
}

func NewPluginHandler(info pluginHandlerInfo, ppreg router.PreprocessorRegistry, p *Plugin) *Handler {
	handler := &Handler{
		Plugin:            p,
		Name:              info.Name,
		AccessKeyRequired: info.KeyRequired,
		UserRequired:      info.UserRequired,
		PreprocessorList:  ppreg,
	}
	return handler
}

func (h *Handler) Setup() {
	if h.UserRequired {
		h.preprocessors = h.PreprocessorList.GetByNames(
			"authenticator",
			"dbconn",
			"inject_user",
			"require_user",
			"plugin_ready",
		)
	} else if h.AccessKeyRequired {
		h.preprocessors = h.PreprocessorList.GetByNames(
			"authenticator",
			"plugin_ready",
		)
	} else {
		h.preprocessors = h.PreprocessorList.GetByNames("plugin_ready")
	}
}

func (h *Handler) GetPreprocessors() []router.Processor {
	return h.preprocessors
}

// Handle executes lambda function implemented by the plugin.
func (h *Handler) Handle(payload *router.Payload, response *router.Response) {
	body, err := ioutil.ReadAll(payload.Req.Body)
	if err != nil {
		panic(err)
	}
	wholeRequest := &pluginRequestPayload{
		Method:      payload.Req.Method,
		Path:        payload.Req.URL.Path,
		QueryString: payload.Req.URL.RawQuery,
		Header:      payload.Req.Header,
		Body:        body,
	}
	inbytes, err := wholeRequest.Encode()
	if err != nil {
		panic(err)
	}

	outbytes, err := h.Plugin.transport.RunHandler(payload.Context, h.Name, inbytes)
	log.WithFields(logrus.Fields{
		"name": h.Name,
		"err":  err,
	}).Debugf("Executed a handler with result")

	if err != nil {
		switch e := err.(type) {
		case skyerr.Error:
			response.Err = e
		case error:
			response.Err = skyerr.MakeError(err)
		}
		return
	}
	responsePayload := &pluginRequestPayload{}
	if err := responsePayload.Decode(outbytes); err != nil {
		response.Err = err
	}

	writer := response.Writer()
	if writer == nil {
		// The response is already written.
		return
	}

	for key, values := range responsePayload.Header {
		for _, value := range values {
			writer.Header().Add(key, value)
		}
	}
	writer.WriteHeader(responsePayload.Status)
	writer.Write(responsePayload.Body)
}
