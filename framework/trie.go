package framework

import (
	"errors"
	"strings"
)

// 判断一个 segment 是否为通配符 segment
func isWildSegment(segment string) bool {
	return strings.HasPrefix(segment, ":")
}

type Tree struct {
	root *node
}

func NewTree() *Tree {
	return &Tree{root: new(node)}
}

func (t *Tree) AddRoute(uri string, handler ControllerHandler) error {
	n := t.root
	if n.machNode(uri) != nil {
		return errors.New("route exist: " + uri)
	}

	segments := strings.Split(uri, "/")
	for index, segment := range segments {
		if !isWildSegment(segment) {
			segment = strings.ToUpper(segment)
		}
		isLast := index == len(segments)-1

		var objNode *node // 标记是否有合适的子节点

		cNodes := n.filterChildNodes(segment)
		if len(cNodes) > 0 {
			for _, cNode := range cNodes {
				if cNode.segment == segment {
					objNode = cNode
					break
				}
			}
		}

		if objNode == nil {
			cNode := new(node)
			cNode.segment = segment
			n.childs = append(n.childs, cNode)
			objNode = cNode
		}

		if isLast {
			objNode.isLast = true
			objNode.handle = handler
		}

		n = objNode
	}

	return nil
}

func (t *Tree) FindHandler(uri string) ControllerHandler {
	if machNode := t.root.machNode(uri); machNode != nil {
		return machNode.handle
	}

	return nil
}

type node struct {
	isLast  bool              // 代表该节点是否能够成为最终的路由规则
	segment string            // uri 中其中一段的字符串
	handle  ControllerHandler // 该节点包含的控制器, 当该节点为最终节点时将其返回供 core 调用
	childs  []*node           //该节点包含的子节点
}

// 过滤下一层满足 segment 规则的子节点
func (n *node) filterChildNodes(segment string) []*node {
	if len(n.childs) == 0 {
		return nil
	}

	if isWildSegment(segment) {
		return n.childs
	}

	var nodes []*node
	for _, val := range n.childs {
		if isWildSegment(val.segment) {
			nodes = append(nodes, val)
		}

		if val.segment == segment {
			nodes = append(nodes, val)
		}
	}

	return nodes
}

// 判断路由是否已在节点的所有子节点树中存在了
func (n *node) machNode(uri string) *node {
	segments := strings.SplitN(uri, "/", 2)

	segment := segments[0]
	if !isWildSegment(segment) {
		segment = strings.ToUpper(segment)
	}

	cNodes := n.filterChildNodes(segment)
	if cNodes == nil || len(cNodes) == 0 {
		return nil
	}

	if len(segments) == 1 {
		for _, val := range cNodes {
			if val.isLast {
				return val
			}
		}

		return nil
	}

	for _, val := range cNodes {
		if tnMach := val.machNode(segments[1]); tnMach != nil {
			return tnMach
		}
	}

	return nil
}
