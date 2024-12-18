package watch

import (
	"context"
	"fmt"
	"sync"

	"github.com/aide-family/moon/pkg/vobj"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// NewDefaultHandler 创建默认消息处理
func NewDefaultHandler(opts ...DefaultHandlerOption) Handler {
	d := &defaultHandler{
		topicHandleMap: make(map[vobj.Topic][]HandleFun),
	}
	for _, opt := range opts {
		opt(d)
	}
	return d
}

type (
	// Handler 消息处理
	Handler interface {
		// Handle 处理消息
		//
		// 	ctx 上下文
		// 	msg 消息
		Handle(ctx context.Context, msg *Message) error
	}

	// HandleFun 消息处理函数
	HandleFun func(ctx context.Context, msg *Message) error

	// defaultHandler 默认消息处理
	defaultHandler struct {
		lock           sync.Mutex
		topicHandleMap map[vobj.Topic][]HandleFun
	}

	// DefaultHandlerOption 默认消息处理配置
	DefaultHandlerOption func(d *defaultHandler)
)

func (d *defaultHandler) Handle(ctx context.Context, msg *Message) error {
	d.lock.Lock()
	defer d.lock.Unlock()
	// 获取处理器
	handles, ok := d.topicHandleMap[msg.GetTopic()]
	if !ok {
		return status.Error(codes.Unimplemented, fmt.Sprintf("%s topic not found handle", msg.GetTopic()))
	}

	// 调用处理器处理msg
	for _, handle := range handles {
		if err := handle(ctx, msg); err != nil {
			return err
		}
	}
	return nil
}

// WithDefaultHandlerTopicHandle 设置默认消息处理
func WithDefaultHandlerTopicHandle(topic vobj.Topic, handles ...HandleFun) DefaultHandlerOption {
	return func(d *defaultHandler) {
		d.lock.Lock()
		defer d.lock.Unlock()
		d.topicHandleMap[topic] = handles
	}
}
