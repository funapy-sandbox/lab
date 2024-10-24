package presenter

import (
	"github.com/Reimei1213/lab/graphql/domain/entity"
	"github.com/Reimei1213/lab/graphql/pkg/graph/model"
)

const NodeTypeUser NodeType = "users"

type User struct {
	*entity.User
}

var _ node = (*User)(nil)

func (u *User) GetID() string {
	return u.ID
}

func (u *User) GetNodeType() NodeType {
	return NodeTypeUser
}

func (u *User) ToGraphqlModel() *model.User {
	if u == nil {
		return nil
	}
	return &model.User{
		ID:   EncodeGraphqlID(u.GetNodeType(), u.ID),
		Name: u.Name,
	}
}

func (u *User) ToGraphqlNode() model.Node {
	return model.Node(u.ToGraphqlModel())
}

func NewUser(u *entity.User) *User {
	return &User{u}
}

func NewUserNodes(us entity.Users) []node {
	res := make([]node, 0, len(us))
	for _, u := range us {
		res = append(res, NewUser(u))
	}
	return res
}