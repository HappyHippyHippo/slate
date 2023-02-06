package envelopemw

import (
	"github.com/happyhippyhippo/slate/env"
	"github.com/happyhippyhippo/slate/rest"
)

const (
	// EnvID defines the slate.rest.envelopemw package base environment
	// variable name.
	EnvID = rest.EnvID + "_ENVELOPE"
)

var (
	// ServiceIDConfigPath defines the config path that used to store the
	// application service identifier.
	ServiceIDConfigPath = env.String(EnvID+"_SERVICE_ID_CONFIG_PATH", "slate.service.id")

	// FormatAcceptListConfigPath defines the config path that used toLogAcceptListErrorMessage
	// store the application accepted mime types formats.
	FormatAcceptListConfigPath = env.String(EnvID+"_FORMAT_ACCEPT_LIST_CONFIG_PATH", "slate.rest.accept")

	// EndpointIDConfigPathFormat defines the format of the configuration
	// path where the endpoint identification number can be retrieved.
	EndpointIDConfigPathFormat = env.String(EnvID+"_ENDPOINT_ID_CONFIG_PATH_FORMAT", "slate.rest.endpoints.%s.id")

	// LogLevel @todo doc
	LogLevel = env.String(EnvID+"_LOG_LEVEL", "error")

	// LogChannel @todo doc
	LogChannel = env.String(EnvID+"_LOG_CHANNEL", "rest")

	// LogServiceErrorMessage @todo doc
	LogServiceErrorMessage = env.String(EnvID+"_LOG_SERVICE_ERROR_MESSAGE", "Invalid service id")

	// LogAcceptListErrorMessage @todo doc
	LogAcceptListErrorMessage = env.String(EnvID+"_LOG_ACCEPT_LIST_ERROR_MESSAGE", "Invalid accept list")

	// LogEndpointErrorMessage @todo doc
	LogEndpointErrorMessage = env.String(EnvID+"_LOG_ENDPOINT_ERROR_MESSAGE", "Invalid endpoint id")
)
