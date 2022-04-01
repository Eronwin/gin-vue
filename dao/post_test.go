package dao

import (
	"testing"

	"github.com/Eronwin/gin-vue/model"
)

func TestAddPost(t *testing.T) {
	post := &model.Post{}
	AddPost(post)
}
