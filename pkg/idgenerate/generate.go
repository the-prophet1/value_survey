package idgenerate

import (
	"github.com/bwmarrin/snowflake"
)

var globalNode *snowflake.Node

func Init() error {
	node, err := snowflake.NewNode(0)
	if err != nil {
		return err
	}
	globalNode = node
	return nil
}

func Generate() int {
	return int(globalNode.Generate().Int64())
}
