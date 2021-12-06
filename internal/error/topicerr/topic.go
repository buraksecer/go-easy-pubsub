package topicerr

import "errors"

var TopicCanNotCreate = errors.New("topic can not create")
var TopicDoExist = errors.New("topic do exist")
var TopicDoNotExist = errors.New("topic do not exist")
var TopicCanNotExists = errors.New("topic does not exist")
var TopicCanNotDelete = errors.New("topic can not delete")
