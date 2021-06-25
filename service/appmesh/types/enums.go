// Code generated by smithy-go-codegen DO NOT EDIT.

package types

type DefaultGatewayRouteRewrite string

// Enum values for DefaultGatewayRouteRewrite
const (
	DefaultGatewayRouteRewriteEnabled  DefaultGatewayRouteRewrite = "ENABLED"
	DefaultGatewayRouteRewriteDisabled DefaultGatewayRouteRewrite = "DISABLED"
)

// Values returns all known values for DefaultGatewayRouteRewrite. Note that this
// can be expanded in the future, and so it is only as up to date as the client.
// The ordering of this slice is not guaranteed to be stable across updates.
func (DefaultGatewayRouteRewrite) Values() []DefaultGatewayRouteRewrite {
	return []DefaultGatewayRouteRewrite{
		"ENABLED",
		"DISABLED",
	}
}

type DnsResponseType string

// Enum values for DnsResponseType
const (
	DnsResponseTypeLoadbalancer DnsResponseType = "LOADBALANCER"
	DnsResponseTypeEndpoints    DnsResponseType = "ENDPOINTS"
)

// Values returns all known values for DnsResponseType. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (DnsResponseType) Values() []DnsResponseType {
	return []DnsResponseType{
		"LOADBALANCER",
		"ENDPOINTS",
	}
}

type DurationUnit string

// Enum values for DurationUnit
const (
	DurationUnitS  DurationUnit = "s"
	DurationUnitMs DurationUnit = "ms"
)

// Values returns all known values for DurationUnit. Note that this can be expanded
// in the future, and so it is only as up to date as the client. The ordering of
// this slice is not guaranteed to be stable across updates.
func (DurationUnit) Values() []DurationUnit {
	return []DurationUnit{
		"s",
		"ms",
	}
}

type EgressFilterType string

// Enum values for EgressFilterType
const (
	EgressFilterTypeAllowAll EgressFilterType = "ALLOW_ALL"
	EgressFilterTypeDropAll  EgressFilterType = "DROP_ALL"
)

// Values returns all known values for EgressFilterType. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (EgressFilterType) Values() []EgressFilterType {
	return []EgressFilterType{
		"ALLOW_ALL",
		"DROP_ALL",
	}
}

type GatewayRouteStatusCode string

// Enum values for GatewayRouteStatusCode
const (
	GatewayRouteStatusCodeActive   GatewayRouteStatusCode = "ACTIVE"
	GatewayRouteStatusCodeInactive GatewayRouteStatusCode = "INACTIVE"
	GatewayRouteStatusCodeDeleted  GatewayRouteStatusCode = "DELETED"
)

// Values returns all known values for GatewayRouteStatusCode. Note that this can
// be expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (GatewayRouteStatusCode) Values() []GatewayRouteStatusCode {
	return []GatewayRouteStatusCode{
		"ACTIVE",
		"INACTIVE",
		"DELETED",
	}
}

type GrpcRetryPolicyEvent string

// Enum values for GrpcRetryPolicyEvent
const (
	GrpcRetryPolicyEventCancelled         GrpcRetryPolicyEvent = "cancelled"
	GrpcRetryPolicyEventDeadlineExceeded  GrpcRetryPolicyEvent = "deadline-exceeded"
	GrpcRetryPolicyEventInternal          GrpcRetryPolicyEvent = "internal"
	GrpcRetryPolicyEventResourceExhausted GrpcRetryPolicyEvent = "resource-exhausted"
	GrpcRetryPolicyEventUnavailable       GrpcRetryPolicyEvent = "unavailable"
)

// Values returns all known values for GrpcRetryPolicyEvent. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (GrpcRetryPolicyEvent) Values() []GrpcRetryPolicyEvent {
	return []GrpcRetryPolicyEvent{
		"cancelled",
		"deadline-exceeded",
		"internal",
		"resource-exhausted",
		"unavailable",
	}
}

type HttpMethod string

// Enum values for HttpMethod
const (
	HttpMethodGet     HttpMethod = "GET"
	HttpMethodHead    HttpMethod = "HEAD"
	HttpMethodPost    HttpMethod = "POST"
	HttpMethodPut     HttpMethod = "PUT"
	HttpMethodDelete  HttpMethod = "DELETE"
	HttpMethodConnect HttpMethod = "CONNECT"
	HttpMethodOptions HttpMethod = "OPTIONS"
	HttpMethodTrace   HttpMethod = "TRACE"
	HttpMethodPatch   HttpMethod = "PATCH"
)

// Values returns all known values for HttpMethod. Note that this can be expanded
// in the future, and so it is only as up to date as the client. The ordering of
// this slice is not guaranteed to be stable across updates.
func (HttpMethod) Values() []HttpMethod {
	return []HttpMethod{
		"GET",
		"HEAD",
		"POST",
		"PUT",
		"DELETE",
		"CONNECT",
		"OPTIONS",
		"TRACE",
		"PATCH",
	}
}

type HttpScheme string

// Enum values for HttpScheme
const (
	HttpSchemeHttp  HttpScheme = "http"
	HttpSchemeHttps HttpScheme = "https"
)

// Values returns all known values for HttpScheme. Note that this can be expanded
// in the future, and so it is only as up to date as the client. The ordering of
// this slice is not guaranteed to be stable across updates.
func (HttpScheme) Values() []HttpScheme {
	return []HttpScheme{
		"http",
		"https",
	}
}

type ListenerTlsMode string

// Enum values for ListenerTlsMode
const (
	ListenerTlsModeStrict     ListenerTlsMode = "STRICT"
	ListenerTlsModePermissive ListenerTlsMode = "PERMISSIVE"
	ListenerTlsModeDisabled   ListenerTlsMode = "DISABLED"
)

// Values returns all known values for ListenerTlsMode. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (ListenerTlsMode) Values() []ListenerTlsMode {
	return []ListenerTlsMode{
		"STRICT",
		"PERMISSIVE",
		"DISABLED",
	}
}

type MeshStatusCode string

// Enum values for MeshStatusCode
const (
	MeshStatusCodeActive   MeshStatusCode = "ACTIVE"
	MeshStatusCodeInactive MeshStatusCode = "INACTIVE"
	MeshStatusCodeDeleted  MeshStatusCode = "DELETED"
)

// Values returns all known values for MeshStatusCode. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (MeshStatusCode) Values() []MeshStatusCode {
	return []MeshStatusCode{
		"ACTIVE",
		"INACTIVE",
		"DELETED",
	}
}

type PortProtocol string

// Enum values for PortProtocol
const (
	PortProtocolHttp  PortProtocol = "http"
	PortProtocolTcp   PortProtocol = "tcp"
	PortProtocolHttp2 PortProtocol = "http2"
	PortProtocolGrpc  PortProtocol = "grpc"
)

// Values returns all known values for PortProtocol. Note that this can be expanded
// in the future, and so it is only as up to date as the client. The ordering of
// this slice is not guaranteed to be stable across updates.
func (PortProtocol) Values() []PortProtocol {
	return []PortProtocol{
		"http",
		"tcp",
		"http2",
		"grpc",
	}
}

type RouteStatusCode string

// Enum values for RouteStatusCode
const (
	RouteStatusCodeActive   RouteStatusCode = "ACTIVE"
	RouteStatusCodeInactive RouteStatusCode = "INACTIVE"
	RouteStatusCodeDeleted  RouteStatusCode = "DELETED"
)

// Values returns all known values for RouteStatusCode. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (RouteStatusCode) Values() []RouteStatusCode {
	return []RouteStatusCode{
		"ACTIVE",
		"INACTIVE",
		"DELETED",
	}
}

type TcpRetryPolicyEvent string

// Enum values for TcpRetryPolicyEvent
const (
	TcpRetryPolicyEventConnectionError TcpRetryPolicyEvent = "connection-error"
)

// Values returns all known values for TcpRetryPolicyEvent. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (TcpRetryPolicyEvent) Values() []TcpRetryPolicyEvent {
	return []TcpRetryPolicyEvent{
		"connection-error",
	}
}

type VirtualGatewayListenerTlsMode string

// Enum values for VirtualGatewayListenerTlsMode
const (
	VirtualGatewayListenerTlsModeStrict     VirtualGatewayListenerTlsMode = "STRICT"
	VirtualGatewayListenerTlsModePermissive VirtualGatewayListenerTlsMode = "PERMISSIVE"
	VirtualGatewayListenerTlsModeDisabled   VirtualGatewayListenerTlsMode = "DISABLED"
)

// Values returns all known values for VirtualGatewayListenerTlsMode. Note that
// this can be expanded in the future, and so it is only as up to date as the
// client. The ordering of this slice is not guaranteed to be stable across
// updates.
func (VirtualGatewayListenerTlsMode) Values() []VirtualGatewayListenerTlsMode {
	return []VirtualGatewayListenerTlsMode{
		"STRICT",
		"PERMISSIVE",
		"DISABLED",
	}
}

type VirtualGatewayPortProtocol string

// Enum values for VirtualGatewayPortProtocol
const (
	VirtualGatewayPortProtocolHttp  VirtualGatewayPortProtocol = "http"
	VirtualGatewayPortProtocolHttp2 VirtualGatewayPortProtocol = "http2"
	VirtualGatewayPortProtocolGrpc  VirtualGatewayPortProtocol = "grpc"
)

// Values returns all known values for VirtualGatewayPortProtocol. Note that this
// can be expanded in the future, and so it is only as up to date as the client.
// The ordering of this slice is not guaranteed to be stable across updates.
func (VirtualGatewayPortProtocol) Values() []VirtualGatewayPortProtocol {
	return []VirtualGatewayPortProtocol{
		"http",
		"http2",
		"grpc",
	}
}

type VirtualGatewayStatusCode string

// Enum values for VirtualGatewayStatusCode
const (
	VirtualGatewayStatusCodeActive   VirtualGatewayStatusCode = "ACTIVE"
	VirtualGatewayStatusCodeInactive VirtualGatewayStatusCode = "INACTIVE"
	VirtualGatewayStatusCodeDeleted  VirtualGatewayStatusCode = "DELETED"
)

// Values returns all known values for VirtualGatewayStatusCode. Note that this can
// be expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (VirtualGatewayStatusCode) Values() []VirtualGatewayStatusCode {
	return []VirtualGatewayStatusCode{
		"ACTIVE",
		"INACTIVE",
		"DELETED",
	}
}

type VirtualNodeStatusCode string

// Enum values for VirtualNodeStatusCode
const (
	VirtualNodeStatusCodeActive   VirtualNodeStatusCode = "ACTIVE"
	VirtualNodeStatusCodeInactive VirtualNodeStatusCode = "INACTIVE"
	VirtualNodeStatusCodeDeleted  VirtualNodeStatusCode = "DELETED"
)

// Values returns all known values for VirtualNodeStatusCode. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (VirtualNodeStatusCode) Values() []VirtualNodeStatusCode {
	return []VirtualNodeStatusCode{
		"ACTIVE",
		"INACTIVE",
		"DELETED",
	}
}

type VirtualRouterStatusCode string

// Enum values for VirtualRouterStatusCode
const (
	VirtualRouterStatusCodeActive   VirtualRouterStatusCode = "ACTIVE"
	VirtualRouterStatusCodeInactive VirtualRouterStatusCode = "INACTIVE"
	VirtualRouterStatusCodeDeleted  VirtualRouterStatusCode = "DELETED"
)

// Values returns all known values for VirtualRouterStatusCode. Note that this can
// be expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (VirtualRouterStatusCode) Values() []VirtualRouterStatusCode {
	return []VirtualRouterStatusCode{
		"ACTIVE",
		"INACTIVE",
		"DELETED",
	}
}

type VirtualServiceStatusCode string

// Enum values for VirtualServiceStatusCode
const (
	VirtualServiceStatusCodeActive   VirtualServiceStatusCode = "ACTIVE"
	VirtualServiceStatusCodeInactive VirtualServiceStatusCode = "INACTIVE"
	VirtualServiceStatusCodeDeleted  VirtualServiceStatusCode = "DELETED"
)

// Values returns all known values for VirtualServiceStatusCode. Note that this can
// be expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (VirtualServiceStatusCode) Values() []VirtualServiceStatusCode {
	return []VirtualServiceStatusCode{
		"ACTIVE",
		"INACTIVE",
		"DELETED",
	}
}
