package models

import (
	. "github.com/SpectoLabs/hoverfly/core/util"
	. "github.com/onsi/gomega"
	"testing"
)

func TestGetDelayWithAMatchingDestination(t *testing.T) {
	RegisterTestingT(t)

	delay := ResponseDelay{
		Destination: StringToPointer("example.com"),
		Delay:       100,
	}
	delays := ResponseDelayList{delay}

	request1 := RequestDetails{
		Destination: "example.com",
	}

	Expect(*delays.GetDelay(request1, false)).To(Equal(delay))

	request2 := RequestDetails{
		Destination: "nodelay.com",
	}

	Expect(delays.GetDelay(request2, false)).To(BeNil())
}

func TestGetDelayWithAMatchingDestinationWithWildcard(t *testing.T) {
	RegisterTestingT(t)

	delay := ResponseDelay{
		Destination: StringToPointer("example.*"),
		Delay:       100,
	}
	delays := ResponseDelayList{delay}

	request1 := RequestDetails{
		Destination: "example.com",
	}

	Expect(*delays.GetDelay(request1, false)).To(Equal(delay))

	request2 := RequestDetails{
		Destination: "example.net",
	}

	Expect(*delays.GetDelay(request2, false)).To(Equal(delay))

	request3 := RequestDetails{
		Destination: "notexample.com",
	}

	Expect(delays.GetDelay(request3, false)).To(BeNil())
}

func TestGetDelayWithAMatchingPath(t *testing.T) {
	RegisterTestingT(t)

	delay := ResponseDelay{
		Path:  StringToPointer("/api/example"),
		Delay: 100,
	}
	delays := ResponseDelayList{delay}

	request1 := RequestDetails{
		Path: "/api/example",
	}

	Expect(*delays.GetDelay(request1, false)).To(Equal(delay))

	request2 := RequestDetails{
		Path: "/api/nodelay",
	}

	Expect(delays.GetDelay(request2, false)).To(BeNil())
}

func TestGetDelayWithAMatchingPathWithWildcard(t *testing.T) {
	RegisterTestingT(t)

	delay := ResponseDelay{
		Path:  StringToPointer("/*/example"),
		Delay: 100,
	}
	delays := ResponseDelayList{delay}

	request1 := RequestDetails{
		Path: "/api/example",
	}

	Expect(*delays.GetDelay(request1, false)).To(Equal(delay))

	request2 := RequestDetails{
		Path: "/something/example",
	}

	Expect(*delays.GetDelay(request2, false)).To(Equal(delay))

	request3 := RequestDetails{
		Path: "notexample.com",
	}

	Expect(delays.GetDelay(request3, false)).To(BeNil())
}

func TestGetDelayWithAMatchingQuery(t *testing.T) {
	RegisterTestingT(t)

	delay := ResponseDelay{
		Query: StringToPointer("?q=query"),
		Delay: 100,
	}
	delays := ResponseDelayList{delay}

	request1 := RequestDetails{
		Query: "?q=query",
	}

	Expect(*delays.GetDelay(request1, false)).To(Equal(delay))

	request2 := RequestDetails{
		Query: "?q=different",
	}

	Expect(delays.GetDelay(request2, false)).To(BeNil())
}

func TestGetDelayWithAMatchingQueryWithWildcard(t *testing.T) {
	RegisterTestingT(t)

	delay := ResponseDelay{
		Query: StringToPointer("?q=*"),
		Delay: 100,
	}
	delays := ResponseDelayList{delay}

	request1 := RequestDetails{
		Query: "?q=query",
	}

	Expect(*delays.GetDelay(request1, false)).To(Equal(delay))

	request2 := RequestDetails{
		Query: "?q=different",
	}

	Expect(*delays.GetDelay(request2, false)).To(Equal(delay))

	request3 := RequestDetails{
		Query: "?p=different-key",
	}

	Expect(delays.GetDelay(request3, false)).To(BeNil())

}

func TestGetDelayWithAMatchingMethod(t *testing.T) {
	RegisterTestingT(t)

	delay := ResponseDelay{
		Method: StringToPointer("GET"),
		Delay:  100,
	}
	delays := ResponseDelayList{delay}

	request1 := RequestDetails{
		Method: "GET",
	}

	Expect(*delays.GetDelay(request1, false)).To(Equal(delay))

	request2 := RequestDetails{
		Path: "POST",
	}

	Expect(delays.GetDelay(request2, false)).To(BeNil())
}

func TestGetDelayWithAMatchingMethodWithWildcard(t *testing.T) {
	RegisterTestingT(t)

	delay := ResponseDelay{
		Method: StringToPointer("*T"),
		Delay:  100,
	}
	delays := ResponseDelayList{delay}

	request1 := RequestDetails{
		Method: "GET",
	}

	Expect(*delays.GetDelay(request1, false)).To(Equal(delay))

	request2 := RequestDetails{
		Method: "POST",
	}

	Expect(*delays.GetDelay(request2, false)).To(Equal(delay))

	request3 := RequestDetails{
		Method: "DELETE",
	}

	Expect(delays.GetDelay(request3, false)).To(BeNil())
}

func TestGetDelayWithAMatchingScheme(t *testing.T) {
	RegisterTestingT(t)

	delay := ResponseDelay{
		Scheme: StringToPointer("https"),
		Delay:  100,
	}
	delays := ResponseDelayList{delay}

	request1 := RequestDetails{
		Scheme: "https",
	}

	Expect(*delays.GetDelay(request1, false)).To(Equal(delay))

	request2 := RequestDetails{
		Scheme: "trhrhr",
	}

	Expect(delays.GetDelay(request2, false)).To(BeNil())
}

func TestGetDelayWithAMatchingSchemeWithWildcard(t *testing.T) {
	RegisterTestingT(t)

	delay := ResponseDelay{
		Scheme: StringToPointer("h*ps"),
		Delay:  100,
	}
	delays := ResponseDelayList{delay}

	request1 := RequestDetails{
		Scheme: "https",
	}

	Expect(*delays.GetDelay(request1, false)).To(Equal(delay))

	request2 := RequestDetails{
		Scheme: "trhrhr",
	}

	Expect(delays.GetDelay(request2, false)).To(BeNil())
}

func TestGetDelayWithAMatchingBody(t *testing.T) {
	RegisterTestingT(t)

	delay := ResponseDelay{
		Body:  StringToPointer("this is a body"),
		Delay: 100,
	}
	delays := ResponseDelayList{delay}

	request1 := RequestDetails{
		Body: "this is a body",
	}

	Expect(*delays.GetDelay(request1, false)).To(Equal(delay))

	request2 := RequestDetails{
		Body: "this is another body",
	}

	Expect(delays.GetDelay(request2, false)).To(BeNil())
}

func TestGetDelayWithAMatchingBodyWithWildcard(t *testing.T) {
	RegisterTestingT(t)

	delay := ResponseDelay{
		Body:  StringToPointer("this is a *"),
		Delay: 100,
	}
	delays := ResponseDelayList{delay}

	request1 := RequestDetails{
		Body: "this is a cat",
	}

	Expect(*delays.GetDelay(request1, false)).To(Equal(delay))

	request2 := RequestDetails{
		Body: "this is not a cat",
	}

	Expect(delays.GetDelay(request2, false)).To(BeNil())
}

func TestGetDelayWithAMatchingDestinationAndPath(t *testing.T) {
	RegisterTestingT(t)

	delay := ResponseDelay{
		Destination: StringToPointer("example.com"),
		Path:        StringToPointer("/api/delay"),
		Delay:       100,
	}
	delays := ResponseDelayList{delay}

	request1 := RequestDetails{
		Destination: "example.com",
		Path:        "/api/delay",
	}

	Expect(*delays.GetDelay(request1, false)).To(Equal(delay))

	request2 := RequestDetails{
		Destination: "example.com",
		Path:        "/api/nodelay",
	}

	Expect(delays.GetDelay(request2, false)).To(BeNil())

	request3 := RequestDetails{
		Destination: "nodelay.com",
		Path:        "/api/delay",
	}

	Expect(delays.GetDelay(request3, false)).To(BeNil())
}

func TestGetDelayShouldIgnoreDestinationIfWebserverBooleanIsTrue(t *testing.T) {
	RegisterTestingT(t)

	delay := ResponseDelay{
		Destination: StringToPointer("example.com"),
		Path:        StringToPointer("/api/delay"),
		Delay:       100,
	}
	delays := ResponseDelayList{delay}

	request1 := RequestDetails{
		Destination: "example.com",
		Path:        "/api/delay",
	}

	Expect(*delays.GetDelay(request1, true)).To(Equal(delay))

	request2 := RequestDetails{
		Destination: "ignored.com",
		Path:        "/api/delay",
	}

	Expect(*delays.GetDelay(request2, true)).To(Equal(delay))

	request3 := RequestDetails{
		Destination: "nodelay.com",
		Path:        "/api/nodelay",
	}

	Expect(delays.GetDelay(request3, true)).To(BeNil())
}

func TestGetDelayWillIterateUntilItFindsAMatch(t *testing.T) {
	RegisterTestingT(t)

	delay1 := ResponseDelay{
		Destination: StringToPointer("wrongdely.org"),
		Delay:       100,
	}
	delay2 := ResponseDelay{
		Destination: StringToPointer("example.com"),
		Body:        StringToPointer("h*llo"),
		Delay:       100,
	}
	delays := ResponseDelayList{delay1, delay2}

	request1 := RequestDetails{
		Destination: "example.com",
		Body:        "hallo",
	}

	Expect(*delays.GetDelay(request1, false)).To(Equal(delay2))
}
