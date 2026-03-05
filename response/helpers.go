package response

type ResponseCode int

func (c ResponseCode) String() string {
	return MessageForCode(c)
}

func (c ResponseCode) Ok() bool {
	return c >= CodeCommandCompletedSuccessfully &&
		c <= CodeCommandCompleteSuccessfullyEndingSession
}

const (
	CodeCommandCompletedSuccessfully                  ResponseCode = 1000
	CodeCommandCompletedSuccessfullyWithActionPending ResponseCode = 1001
	CodeCommandCompleteSuccessfullyWithNoMessages     ResponseCode = 1300
	CodeCommandCompleteSuccessfullyAckToDequeue       ResponseCode = 1301
	CodeCommandCompleteSuccessfullyEndingSession      ResponseCode = 1500
	CodeUnknownCommand                                ResponseCode = 2000
	CodeCommandSyntaxError                            ResponseCode = 2001
	CodeCommandUseError                               ResponseCode = 2002
	CodeRequiredParameterMissing                      ResponseCode = 2003
	CodeParameterValueRangeError                      ResponseCode = 2004
	CodeParameterValueSyntaxError                     ResponseCode = 2005
	CodeUnimplementedProtocolVersion                  ResponseCode = 2100
	CodeUnimplementedCommand                          ResponseCode = 2101
	CodeUnimplementedOption                           ResponseCode = 2102
	CodeUnimplementedExtension                        ResponseCode = 2103
	CodeBillingFailure                                ResponseCode = 2104
	CodeObjectIsNotEligibleForRenewal                 ResponseCode = 2105
	CodeObjectIsNotEligibleForTransfer                ResponseCode = 2106
	CodeAuthenticationError                           ResponseCode = 2200
	CodeAuthorizationError                            ResponseCode = 2201
	CodeInvalidAuthorizationInformation               ResponseCode = 2202
	CodeObjectPendingTransfer                         ResponseCode = 2300
	CodeObjectNotPendingTransfer                      ResponseCode = 2301
	CodeObjectExists                                  ResponseCode = 2302
	CodeObjectDoesNotExist                            ResponseCode = 2303
	CodeObjectStatusProhibitsOperation                ResponseCode = 2304
	CodeObjectAssociationProhibitsOperation           ResponseCode = 2305
	CodeParameterValuePolicyError                     ResponseCode = 2306
	CodeUnimplementedObjectService                    ResponseCode = 2307
	CodeDataManagementPolicyViolation                 ResponseCode = 2308
	CodeCommandFailed                                 ResponseCode = 2400
	CodeCommandFailedServerClosingConnection          ResponseCode = 2500
	CodeAuthenticationErrorServerClosingConnection    ResponseCode = 2501
	CodeSessionLimitExceededServerClosingConnection   ResponseCode = 2502
)

const (
	XmlHeader = `<?xml version="1.0" encoding="UTF-8" standalone="no"?>`

	CommandCompletedSuccessfully                  = "Command completed successfully"
	CommandCompletedSuccessfullyWithActionPending = "Command completed successfully; action pending"
	CommandCompleteSuccessfullyWithNoMessages     = "Command completed successfully; no messages"
	CommandCompleteSuccessfullyAckToDequeue       = "Command completed successfully; ack to dequeue"
	CommandCompleteSuccessfullyEndingSession      = "Command completed successfully; ending session"
	UnknownCommand                                = "Unknown command"
	CommandSyntaxError                            = "Command syntax error"
	CommandUseError                               = "Command use error"
	RequiredParameterMissing                      = "Required parameter missing"
	ParameterValueRangeError                      = "Parameter value range error"
	ParameterValueSyntaxError                     = "Parameter value syntax error"
	UnimplementedProtocolVersion                  = "Unimplemented protocol version"
	UnimplementedCommand                          = "Unimplemented command"
	UnimplementedOption                           = "Unimplemented option"
	UnimplementedExtension                        = "Unimplemented extension"
	BillingFailure                                = "Billing failure"
	ObjectIsNotEligibleForRenewal                 = "Object is not eligible for renewal"
	ObjectIsNotEligibleForTransfer                = "Object is not eligible for transfer"
	AuthenticationError                           = "Authentication error"
	AuthorizationError                            = "Authorization error"
	InvalidAuthorizationInformation               = "Invalid authorization information"
	ObjectPendingTransfer                         = "Object pending transfer"
	ObjectNotPendingTransfer                      = "Object not pending transfer"
	ObjectExists                                  = "Object exists"
	ObjectDoesNotExist                            = "Object does not exist"
	ObjectStatusProhibitsOperation                = "Object status prohibits operation"
	ObjectAssociationProhibitsOperation           = "Object association prohibits operation"
	ParameterValuePolicyError                     = "Parameter value policy error"
	UnimplementedObjectService                    = "Unimplemented object service"
	DataManagementPolicyViolation                 = "Data management policy violation"
	CommandFailed                                 = "Command failed"
	CommandFailedServerClosingConnection          = "Command failed; server closing connection"
	AuthenticationErrorServerClosingConnection    = "Authentication error; server closing connection"
	SessionLimitExceededServerClosingConnection   = "Session limit exceeded; server closing connection"
)

func AnyError(code ResponseCode, msg string) *EPPResponse[struct{}, struct{}] {
	return NewResponse[struct{}, struct{}](code, msg)
}

//nolint:gocyclo,cyclop
func MessageForCode(code ResponseCode) string {
	// see https://datatracker.ietf.org/doc/html/rfc5730#section-3
	switch code {
	case CodeCommandCompletedSuccessfully:
		return CommandCompletedSuccessfully
	case CodeCommandCompletedSuccessfullyWithActionPending:
		return CommandCompletedSuccessfullyWithActionPending
	case CodeCommandCompleteSuccessfullyWithNoMessages:
		return CommandCompleteSuccessfullyWithNoMessages
	case CodeCommandCompleteSuccessfullyAckToDequeue:
		return CommandCompleteSuccessfullyAckToDequeue
	case CodeCommandCompleteSuccessfullyEndingSession:
		return CommandCompleteSuccessfullyEndingSession

	case CodeUnknownCommand:
		return UnknownCommand
	case CodeCommandSyntaxError:
		return CommandSyntaxError
	case CodeCommandUseError:
		return CommandUseError
	case CodeRequiredParameterMissing:
		return RequiredParameterMissing
	case CodeParameterValueRangeError:
		return ParameterValueRangeError
	case CodeParameterValueSyntaxError:
		return ParameterValueSyntaxError

	case CodeUnimplementedProtocolVersion:
		return UnimplementedProtocolVersion
	case CodeUnimplementedCommand:
		return UnimplementedCommand
	case CodeUnimplementedOption:
		return UnimplementedOption
	case CodeUnimplementedExtension:
		return UnimplementedExtension
	case CodeBillingFailure:
		return BillingFailure
	case CodeObjectIsNotEligibleForRenewal:
		return ObjectIsNotEligibleForRenewal
	case CodeObjectIsNotEligibleForTransfer:
		return ObjectIsNotEligibleForTransfer

	case CodeAuthenticationError:
		return AuthenticationError
	case CodeAuthorizationError:
		return AuthorizationError
	case CodeInvalidAuthorizationInformation:
		return InvalidAuthorizationInformation

	case CodeObjectPendingTransfer:
		return ObjectPendingTransfer
	case CodeObjectNotPendingTransfer:
		return ObjectNotPendingTransfer
	case CodeObjectExists:
		return ObjectExists
	case CodeObjectDoesNotExist:
		return ObjectDoesNotExist
	case CodeObjectStatusProhibitsOperation:
		return ObjectStatusProhibitsOperation
	case CodeObjectAssociationProhibitsOperation:
		return ObjectAssociationProhibitsOperation
	case CodeParameterValuePolicyError:
		return ParameterValuePolicyError
	case CodeUnimplementedObjectService:
		return UnimplementedObjectService
	case CodeDataManagementPolicyViolation:
		return DataManagementPolicyViolation

	case CodeCommandFailed:
		return CommandFailed

	case CodeCommandFailedServerClosingConnection:
		return CommandFailedServerClosingConnection
	case CodeAuthenticationErrorServerClosingConnection:
		return AuthenticationErrorServerClosingConnection
	case CodeSessionLimitExceededServerClosingConnection:
		return SessionLimitExceededServerClosingConnection

	default:
		return "Unknown code"
	}
}
