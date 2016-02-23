/*
 * generated by event_generator
 *
 * DO NOT EDIT
 */

package messages

import "github.com/joernweissenborn/eventual2go"

type MessageCompleter struct {
	*eventual2go.Completer
}

func NewMessageCompleter() *MessageCompleter {
	return &MessageCompleter{eventual2go.NewCompleter()}
}

func (c *MessageCompleter) Complete(d Message) {
	c.Completer.Complete(d)
}

func (c *MessageCompleter) Future() *MessageFuture {
	return &MessageFuture{c.Completer.Future()}
}

type MessageFuture struct {
	*eventual2go.Future
}

type MessageCompletionHandler func(Message) Message

func (ch MessageCompletionHandler) toCompletionHandler() eventual2go.CompletionHandler {
	return func(d eventual2go.Data) eventual2go.Data {
		return ch(d.(Message))
	}
}

func (f *MessageFuture) Then(ch MessageCompletionHandler) *MessageFuture {
	return &MessageFuture{f.Future.Then(ch.toCompletionHandler())}
}

func (f *MessageFuture) AsChan() chan Message {
	c := make(chan Message, 1)
	cmpl := func(d chan Message) MessageCompletionHandler {
		return func(e Message) Message {
			d <- e
			close(d)
			return e
		}
	}
	ecmpl := func(d chan Message) eventual2go.ErrorHandler {
		return func(error) (eventual2go.Data, error) {
			close(d)
			return nil, nil
		}
	}
	f.Then(cmpl(c))
	f.Err(ecmpl(c))
	return c
}

type MessageStreamController struct {
	*eventual2go.StreamController
}

func NewMessageStreamController() *MessageStreamController {
	return &MessageStreamController{eventual2go.NewStreamController()}
}

func (sc *MessageStreamController) Add(d Message) {
	sc.StreamController.Add(d)
}

func (sc *MessageStreamController) Join(s *MessageStream) {
	sc.StreamController.Join(s.Stream)
}

func (sc *MessageStreamController) JoinFuture(f *MessageFuture) {
	sc.StreamController.JoinFuture(f.Future)
}

func (sc *MessageStreamController) Stream() *MessageStream {
	return &MessageStream{sc.StreamController.Stream()}
}

type MessageStream struct {
	*eventual2go.Stream
}

type MessageSuscriber func(Message)

func (l MessageSuscriber) toSuscriber() eventual2go.Subscriber {
	return func(d eventual2go.Data) { l(d.(Message)) }
}

func (s *MessageStream) Listen(ss MessageSuscriber) *eventual2go.Subscription {
	return s.Stream.Listen(ss.toSuscriber())
}

type MessageFilter func(Message) bool

func (f MessageFilter) toFilter() eventual2go.Filter {
	return func(d eventual2go.Data) bool { return f(d.(Message)) }
}

func (s *MessageStream) Where(f MessageFilter) *MessageStream {
	return &MessageStream{s.Stream.Where(f.toFilter())}
}

func (s *MessageStream) WhereNot(f MessageFilter) *MessageStream {
	return &MessageStream{s.Stream.WhereNot(f.toFilter())}
}

func (s *MessageStream) First() *MessageFuture {
	return &MessageFuture{s.Stream.First()}
}

func (s *MessageStream) FirstWhere(f MessageFilter) *MessageFuture {
	return &MessageFuture{s.Stream.FirstWhere(f.toFilter())}
}

func (s *MessageStream) FirstWhereNot(f MessageFilter) *MessageFuture {
	return &MessageFuture{s.Stream.FirstWhereNot(f.toFilter())}
}

func (s *MessageStream) AsChan() (c chan Message) {
	c = make(chan Message)
	s.Listen(pipeToMessageChan(c)).Closed().Then(closeMessageChan(c))
	return
}

func pipeToMessageChan(c chan Message) MessageSuscriber {
	return func(d Message) {
		c <- d
	}
}

func closeMessageChan(c chan Message) eventual2go.CompletionHandler {
	return func(d eventual2go.Data) eventual2go.Data {
		close(c)
		return nil
	}
}
