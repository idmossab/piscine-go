package piscine

type NodeL struct {
	Data interface{}
	Next *NodeL
}

type List struct {
	Head *NodeL
	Tail *NodeL
}

func ListSize(l *List) int {
  size:=0
  for i:=l.Head ;i!=l.Tail;i=i.Next{
    size++
  }
  if l.Tail!=nil{
    size++
  }
  return size
}