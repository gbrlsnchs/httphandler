package httphandler

// ContentType is the type of content a Handler can send.
type ContentType string

const (
	// ContentTypeHeader is the respective name for the Content-Type MIME type.
	ContentTypeHeader = "Content-Type"
	// ContentTypeJSON is the Content-Type MIME type value for JSON responses.
	ContentTypeJSON ContentType = "application/json"
	// ContentTypeMsgPack is the Content-Type MIME type value for MsgPack responses.
	ContentTypeMsgPack = "application/vnd.msgpack"
	// ContentTypeTextPlain is the Content-Type MIME type value for plain text responses.
	ContentTypeTextPlain = "text/plain"
	// ContentTypeXMsgPack is the Content-Type MIME type value for MsgPack responses.
	ContentTypeXMsgPack = "application/x-msgpack"
	// ContentTypeXML is the Content-Type MIME type value for XML responses.
	ContentTypeXML = "application/xml"
)
