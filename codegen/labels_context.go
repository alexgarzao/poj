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
		labelID: 100, // TODO: must be zero :-)
	}
}

func (l *LabelsContext) Add() {
	labels := Labels{}

	l.labelID++
	labels.Else = fmt.Sprintf("L%d", l.labelID)

	l.labelID++
	labels.NextStatement = fmt.Sprintf("L%d", l.labelID)

	l.labelID++
	labels.IterationStart = fmt.Sprintf("L%d", l.labelID)

	l.labels.Push(labels)
}

func (l *LabelsContext) Rem() {
	l.labels.Pop()
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
