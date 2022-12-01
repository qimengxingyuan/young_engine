package executor

import (
	"bytes"
	"encoding/xml"
	"fmt"
)

const (
	svgFormat = `<svg xmlns="http://www.w3.org/2000/svg" xmlns:xlink="http://www.w3.org/1999/xlink" version="1.1" width="%d" height="%d">%s</svg>`

	circleR = 20
	yOffset = 50

	nodeWidth  = 40
	nodeHeight = 100

	colorRed    = "red"
	colorOrange = "orange"
	colorGreen  = "lightgreen"
)

func (n *Node) layerOrder() [][]interface{} {
	var res [][]interface{}
	if n == nil {
		return res
	}

	nodes := make([][]*Node, 0)
	lastLevel := 0                    // 当前的层数
	nodes = append(nodes, []*Node{n}) // lastLevel node
	for {
		finish := true
		for _, node := range nodes[lastLevel] {
			if node != nil {
				finish = false
				break
			}
		}
		if finish {
			nodes = nodes[:lastLevel] // why
			break
		}

		node := make([]*Node, 0)
		for _, root := range nodes[lastLevel] {
			if root == nil {
				node = append(node, nil, nil)
			} else {
				node = append(node, root.leftNode, root.rightNode)
			}
		}
		nodes = append(nodes, node)
		lastLevel++
	}

	for _, innerNodes := range nodes {
		retVal := make([]interface{}, len(innerNodes))
		for i, node := range innerNodes {
			if node == nil {
				retVal[i] = nil
			} else if node.symbol == LITERAL || node.symbol == VALUE {
				retVal[i] = node.value
			} else {
				retVal[i] = node.symbol.String()
			}
		}
		res = append(res, retVal)
	}

	return res
}

func (n *Node) xmlTree() string {
	List := n.layerOrder()
	Levels := len(List)

	var nodeXml string
	for i := Levels - 1; i >= 0; i-- {
		negative := -1
		curLevelXml := ""
		for j := 0; j < pow2(i); j++ {
			t := pow2(Levels - i - 1)
			x, y := nodeWidth*(2*t*j+t), nodeHeight*i+yOffset
			if List[i][j] != nil {
				fillColor := colorOrange
				if i == Levels-1 || i > 0 && i < Levels-1 &&
					List[i+1][j*2] == nil && List[i+1][j*2+1] == nil {
					fillColor = colorGreen
				}

				//textData := fmt.Sprintf("%v", List[i][j])
				var b bytes.Buffer
				xml.Escape(&b, []byte(fmt.Sprintf("%v", List[i][j])))
				textData := b.String()
				var x1, y1, x2, y2 int
				if i > 0 {
					x1 = x
					y1 = y - circleR
					negative *= -1
					x2 = x + nodeWidth*negative*(2*t*j%2+t) - negative*circleR
					y2 = y - nodeHeight + circleR/2 // 节点的1/4处
				}
				curLevelXml += xmlNode(i, j, x, y, x1, y1, x2, y2, textData, fillColor, colorRed)
			} else {
				negative *= -1
			}
		}
		nodeXml = curLevelXml + nodeXml
	}

	width := pow2(Levels) * nodeWidth
	height := Levels * nodeHeight
	xml := fmt.Sprintf(svgFormat, width, height, nodeXml)
	return xml
}

func xmlNode(m, n, x, y, x1, y1, x2, y2 int, textData, cycleColor, textColor string) string {
	id := fmt.Sprintf("<g id=\"%d,%d\">\n", m, n)
	circle := xmlCircle(x, y, circleR, cycleColor)
	text := xmlText(x, y, textData, textColor)
	end := "</g>\n"

	var line string
	if x1 != 0 || y1 != 0 || x2 != 0 || y2 != 0 {
		line = xmlLine(x1, y1, x2, y2)
	}

	nodeStr := id + circle + text
	if line != "" {
		nodeStr += line
	}
	nodeStr += end
	return nodeStr
}

func xmlCircle(x, y, r int, color string) string {
	return fmt.Sprintf("<circle cx=\"%d\" cy=\"%d\" r=\"%d\" stroke=\"black\" stroke-width=\"2\" fill=\"%s\" />\n",
		x, y, r, color)
}

func xmlText(x, y int, text, color string) string {
	return fmt.Sprintf("<text x=\"%d\" y=\"%d\" fill=\"%s\" font-size=\"14\" text-anchor=\"middle\" dominant-baseline=\"middle\">%s</text>\n",
		x, y, color, text)
}

func xmlLine(x1, y1, x2, y2 int) string {
	return fmt.Sprintf("<line x1=\"%d\" y1=\"%d\" x2=\"%d\" y2=\"%d\" style=\"stroke:black;stroke-width:2\"/>\n", x1, y1, x2, y2)
}

func pow2(x int) int { //x>=0
	res := 1
	for i := 0; i < x; i++ {
		res *= 2
	}
	return res
}
