package gpm

type GPM struct {
	nodeManager *Manager
	hostManager *Manager
}

type Host struct {
	IDName
	Host string
	Port string
}

type Node struct {
	IDName
	Host Host
}

func (g *GPM) Init() {
	g.nodeManager = NewManager()
	g.hostManager = NewManager()
}

func (g *GPM) AddHost(h *Host) error {
	return g.hostManager.Add(h.IDName, h)
}

func (g *GPM)RemoveHost(h *Host) {
	g.hostManager.Remove(h.IDName)
}

func (g *GPM) GetHost(in IDName) *Host {
	h := g.hostManager.Get(in)
	if h == nil {
		return nil
	}
	return h.(*Host)
}

func (g *GPM) GetHostByName(host string) *Host {
	return g.GetHost(IDName{Name: host})
}

func (g *GPM) GetHostByID(id int64) *Host {
	return g.GetHost(IDName{ID: id})
}

func (g *GPM) AddNode(n *Node) error {
	return g.nodeManager.Add(n.IDName, n)
}

func (g *GPM)RemoveNode(n *Node) {
	g.nodeManager.Remove(n.IDName)
}

func (g *GPM) GetNode(in IDName) *Node {
	n := g.nodeManager.Get(in)
	if n == nil {
		return nil
	}
	return n.(*Node)
}

func (g *GPM) GetNodeByName(node string) *Node {
	return g.GetNode(IDName{Name: node})
}

func (g *GPM) GetNodeByID(id int64) *Node {
	return g.GetNode(IDName{ID: id})
}
