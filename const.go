package httphandler

// ContentTypeHeader is the respective name for the Content-Type MIME type.
const ContentTypeHeader = "Content-Type"

// ContentType is the type of content a Handler can send.
type ContentType string

const (
	// ContentTypeTextPlain is the Content-Type MIME type value for plain text responses.
	ContentTypeTextPlain ContentType = "text/plain"
	// ContentTypeJSON is the Content-Type MIME type value for JSON responses.
	ContentTypeJSON ContentType = "application/json"
	// ContentTypeXML is the Content-Type MIME type value for XML responses.
	ContentTypeXML ContentType = "application/xml"
	// ContentTypeMsgPack is the Content-Type MIME type value for MsgPack responses.
	ContentTypeMsgPack ContentType = "application/vnd.msgpack"
	// ContentTypeXMsgPack is the Content-Type MIME type value for MsgPack responses.
	ContentTypeXMsgPack ContentType = "application/x-msgpack"
)
