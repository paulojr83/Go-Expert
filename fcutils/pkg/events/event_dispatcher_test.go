package events

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
	"sync"
	"testing"
	"time"
)

type TestEvent struct {
	Name    string
	Payload interface{}
}

func (t *TestEvent) GetName() string {
	return t.Name
}

func (t *TestEvent) GetPayload() interface{} {
	return t.Payload
}

func (t *TestEvent) GetDateTime() time.Time {
	return time.Now()
}

type TestEventHandler struct {
	ID int
}

func (t *TestEventHandler) Handle(event EventInterface, wg *sync.WaitGroup) {

}

type EventDispatcherTestSuite struct {
	suite.Suite
	event           TestEvent
	event2          TestEvent
	handler         TestEventHandler
	handler1        TestEventHandler
	handler2        TestEventHandler
	eventDispatcher *EventDispatcher
}

func (suite *EventDispatcherTestSuite) SetupTest() {
	suite.eventDispatcher = NewEventDispatcher()
	suite.handler = TestEventHandler{ID: 1}
	suite.handler2 = TestEventHandler{ID: 2}
	suite.handler2 = TestEventHandler{ID: 3}
	suite.event = TestEvent{
		Name:    "Test",
		Payload: "{\"name\":\"João\", \"year\":\"20\"}",
	}
	suite.event2 = TestEvent{
		Name:    "Test 2",
		Payload: "{\"name\":\"Maria\", \"year\":\"30\"}",
	}
}
func (suite *EventDispatcherTestSuite) TestEventDispatcher_Register() {
	err := suite.eventDispatcher.Register(suite.event.GetName(), &suite.handler)
	suite.Nil(err)
	suite.Equal(1, len(suite.eventDispatcher.handlers[suite.event.GetName()]))

	suite.eventDispatcher.Register(suite.event.GetName(), &suite.handler1)
	suite.Nil(err)
	suite.Equal(2, len(suite.eventDispatcher.handlers[suite.event.GetName()]))

	assert.Equal(suite.T(), &suite.handler, suite.eventDispatcher.handlers[suite.event.GetName()][0])
	assert.Equal(suite.T(), &suite.handler1, suite.eventDispatcher.handlers[suite.event.GetName()][1])
}

func (suite *EventDispatcherTestSuite) TestEventDispatcher_Register_WithSameHandler() {
	err := suite.eventDispatcher.Register(suite.event.GetName(), &suite.handler)
	suite.Nil(err)
	suite.Equal(1, len(suite.eventDispatcher.handlers[suite.event.GetName()]))

	err = suite.eventDispatcher.Register(suite.event.GetName(), &suite.handler)
	suite.Equal(ErrHandlerAlreadyRegistered, err)
	suite.Equal(1, len(suite.eventDispatcher.handlers[suite.event.GetName()]))
}

func (suite *EventDispatcherTestSuite) TestEventDispatcher_Clear() {
	// Event 1
	err := suite.eventDispatcher.Register(suite.event.GetName(), &suite.handler)
	suite.Nil(err)
	suite.Equal(1, len(suite.eventDispatcher.handlers[suite.event.GetName()]))

	suite.eventDispatcher.Register(suite.event.GetName(), &suite.handler1)
	suite.Nil(err)
	suite.Equal(2, len(suite.eventDispatcher.handlers[suite.event.GetName()]))

	// Event 2
	err = suite.eventDispatcher.Register(suite.event2.GetName(), &suite.handler2)
	suite.Nil(err)
	suite.Equal(1, len(suite.eventDispatcher.handlers[suite.event2.GetName()]))

	suite.eventDispatcher.Clear()
	suite.Equal(0, len(suite.eventDispatcher.handlers))
}

func (suite *EventDispatcherTestSuite) TestEventDispatcher_Has() {
	// Event 1
	err := suite.eventDispatcher.Register(suite.event.GetName(), &suite.handler)
	suite.Nil(err)
	suite.Equal(1, len(suite.eventDispatcher.handlers[suite.event.GetName()]))

	suite.eventDispatcher.Register(suite.event.GetName(), &suite.handler1)
	suite.Nil(err)
	suite.Equal(2, len(suite.eventDispatcher.handlers[suite.event.GetName()]))

	// Event 2
	err = suite.eventDispatcher.Register(suite.event2.GetName(), &suite.handler2)
	suite.Nil(err)
	suite.Equal(1, len(suite.eventDispatcher.handlers[suite.event2.GetName()]))

	assert.True(suite.T(), suite.eventDispatcher.Has(suite.event.GetName(), &suite.handler))
	assert.True(suite.T(), suite.eventDispatcher.Has(suite.event2.GetName(), &suite.handler2))

	assert.False(suite.T(), suite.eventDispatcher.Has(suite.event2.GetName(), &suite.handler1))
}

type MockHandler struct {
	mock.Mock
}

func (m *MockHandler) Handle(event EventInterface, wg *sync.WaitGroup) {
	m.Called(event)
	wg.Done()
}

func (suite *EventDispatcherTestSuite) TestEventDispatcher_Dispatch() {
	eh := &MockHandler{}
	eh.On("Handle", &suite.event)
	suite.eventDispatcher.Register(suite.event.GetName(), eh)
	suite.eventDispatcher.Dispatch(&suite.event)

	eh.AssertExpectations(suite.T())
	eh.AssertNumberOfCalls(suite.T(), "Handle", 1)
}

func (suite *EventDispatcherTestSuite) TestEventDispatcher_Remove() {
	// Event 1
	err := suite.eventDispatcher.Register(suite.event.GetName(), &suite.handler)
	suite.Nil(err)
	suite.Equal(1, len(suite.eventDispatcher.handlers[suite.event.GetName()]))

	suite.eventDispatcher.Register(suite.event.GetName(), &suite.handler1)
	suite.Nil(err)
	suite.Equal(2, len(suite.eventDispatcher.handlers[suite.event.GetName()]))

	// Event 2
	err = suite.eventDispatcher.Register(suite.event2.GetName(), &suite.handler2)
	suite.Nil(err)
	suite.Equal(1, len(suite.eventDispatcher.handlers[suite.event2.GetName()]))

	suite.eventDispatcher.Remove(suite.event.GetName(), &suite.handler)
	suite.Equal(1, len(suite.eventDispatcher.handlers[suite.event.GetName()]))

	suite.eventDispatcher.Remove(suite.event.GetName(), &suite.handler1)
	suite.Equal(0, len(suite.eventDispatcher.handlers[suite.event.GetName()]))

	suite.eventDispatcher.Remove(suite.event2.GetName(), &suite.handler2)
	suite.Equal(0, len(suite.eventDispatcher.handlers[suite.event2.GetName()]))

}

func TestSuite(t *testing.T) {
	suite.Run(t, new(EventDispatcherTestSuite))
}
