package envelope

import (
	senv "github.com/happyhippyhippo/slate/env"
	srest "github.com/happyhippyhippo/slate/rest"
)

const (
	// ContainerID defines the default id used to register
	// the application envelope middleware and related services.
	ContainerID = srest.ContainerID + ".envelope"
)

const (
	// EnvID defines the rest.envelope package base environment variable name.
	EnvID = srest.EnvID + "_ENVELOPE"
)

var (
	// ConfigPathServiceID defines the config path that used to store the
	// application service identifier.
	ConfigPathServiceID = senv.String(EnvID+"_CONFIG_PATH_SERVER_ID", "slate.service.id")

	// ConfigPathTransportAcceptList defines the config path that used to
	// store the application accepted mime types.
	ConfigPathTransportAcceptList = senv.String(EnvID+"_CONFIG_PATH_TRANSPORT_ACCEPT_LIST", "slate.rest.accept")

	// ConfigPathEndpointIDFormat defines the format of the configuration
	// path where the endpoint identification number can be retrieved.
	ConfigPathEndpointIDFormat = senv.String(EnvID+"_CONFIG_PATH_ENDPOINT_ID_FORMAT", "slate.endpoints.%s.id")
)
