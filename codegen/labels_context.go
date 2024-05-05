package codegen

import "fmt"

type Labels struct {
	Else           string
	NextStatement  string
	IterationStart string
}

type LabelsContext struct {
	labels  *Stack[Labels]
	labelID int
}

func NewLabelsContext() *LabelsContext {
	return &LabelsContext{
		labels:  NewStack[Labels](),
		labelID: 0,
	}
}

func (l *LabelsContext) Add() {
	l.labels.Push(Labels{
		Else:           l.NewLabel(),
		NextStatement:  l.NewLabel(),
		IterationStart: l.NewLabel(),
	})
}

func (l *LabelsContext) Rem() {
	l.labels.Pop()
}

func (l *LabelsContext) NewLabel() string {
	l.labelID++
	return fmt.Sprintf("L%d", l.labelID)
}

func (l *LabelsContext) Else() string {
	labels, _ := l.labels.Top()

	return labels.Else
}

func (l *LabelsContext) NextStatement() string {
	labels, _ := l.labels.Top()

	return labels.NextStatement
}

func (l *LabelsContext) IterationStart() string {
	labels, _ := l.labels.Top()

	return labels.IterationStart
}
