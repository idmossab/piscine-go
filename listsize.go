package piscine

func ListSize(l *List) int {
  res:=0
  for i:=l.Head ;i<l.Tail;i++{
    res++
  }
  return res
}