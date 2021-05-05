package book

import (
	"context"
	"testing"

	"github.com/graniticio/granitic/v2/test"
	"github.com/graniticio/granitic/v2/ws"
)

func Test_getBook(t *testing.T) {
	logic := GetBookLogic{}
	req := ws.Request{}
	res := ws.Response{}
	logic.Process(context.TODO(), &req, &res)
	test.ExpectString(t, res.Body.(*Book).Name, "My Book")
}
