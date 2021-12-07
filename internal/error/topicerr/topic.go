package topicerr

import "errors"

// ErrTopicCanNotCreate err
var ErrTopicCanNotCreate = errors.New("topic can not create")

// ErrTopicDoExist err
var ErrTopicDoExist = errors.New("topic do exist")

// ErrTopicDoNotExist err
var ErrTopicDoNotExist = errors.New("topic do not exist")

// ErrTopicCanNotExists err
var ErrTopicCanNotExists = errors.New("topic does not exist")

// ErrTopicCanNotDelete err
var ErrTopicCanNotDelete = errors.New("topic can not delete")
