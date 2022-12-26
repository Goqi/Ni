package eventcreator

import (
	"Ernuclei/pkg/operators"
	"Ernuclei/pkg/output"
	"Ernuclei/pkg/protocols"
	"Ernuclei/pkg/protocols/common/utils/vardump"
	"github.com/projectdiscovery/gologger"
)

// CreateEvent wraps the outputEvent with the result of the operators defined on the request
func CreateEvent(request protocols.Request, outputEvent output.InternalEvent, isResponseDebug bool) *output.InternalWrappedEvent {
	return CreateEventWithAdditionalOptions(request, outputEvent, isResponseDebug, nil)
}

// CreateEventWithAdditionalOptions wraps the outputEvent with the result of the operators defined on the request
// and enables extending the resulting event with additional attributes or values.
func CreateEventWithAdditionalOptions(request protocols.Request, outputEvent output.InternalEvent, isResponseDebug bool,
	addAdditionalOptions func(internalWrappedEvent *output.InternalWrappedEvent)) *output.InternalWrappedEvent {
	event := &output.InternalWrappedEvent{InternalEvent: outputEvent}

	// Dump response variables if ran in debug mode
	if vardump.EnableVarDump {
		gologger.Debug().Msgf("Protocol response variables: \n%s\n", vardump.DumpVariables(outputEvent))
	}
	for _, compiledOperator := range request.GetCompiledOperators() {
		if compiledOperator != nil {
			result, ok := compiledOperator.Execute(outputEvent, request.Match, request.Extract, isResponseDebug)
			if ok && result != nil {
				event.OperatorsResult = result
				if addAdditionalOptions != nil {
					addAdditionalOptions(event)
				}
				event.Results = append(event.Results, request.MakeResultEvent(event)...)
			}
		}
	}
	return event
}

func CreateEventWithOperatorResults(request protocols.Request, internalEvent output.InternalEvent, operatorResult *operators.Result) *output.InternalWrappedEvent {
	event := &output.InternalWrappedEvent{InternalEvent: internalEvent}
	event.OperatorsResult = operatorResult
	event.Results = append(event.Results, request.MakeResultEvent(event)...)
	return event
}
