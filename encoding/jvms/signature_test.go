package jvms

import (
	_ "github.com/kr/pretty"
	_ "reflect"
	"strings"
	"testing"
)

func TestParseClassSignature(t *testing.T) {
	cases := []struct {
		S string
	}{
		// java.util.ArrayList
		{S: "<E:Ljava/lang/Object;>Ljava/util/AbstractList<TE;>;Ljava/util/List<TE;>;Ljava/util/RandomAccess;Ljava/lang/Cloneable;Ljava/io/Serializable;"},
		// java.util.List
		{S: "<E:Ljava/lang/Object;>Ljava/lang/Object;Ljava/util/Collection<TE;>;"},
		// java.util.LinkedList
		{S: "<E:Ljava/lang/Object;>Ljava/util/AbstractSequentialList<TE;>;Ljava/util/List<TE;>;Ljava/util/Deque<TE;>;Ljava/lang/Cloneable;Ljava/io/Serializable;"},
	}
	for _, c := range cases {
		if err := ParseClassSignature(strings.NewReader(c.S)); err != nil {
			t.Errorf("Can't parse ClassSignature %#v: %s", c.S, err)
		}
	}
}

func TestParseFieldTypeSignature(t *testing.T) {
	cases := []struct {
		S string
	}{
		// java.util.LinkedList
		{S: "Ljava/util/LinkedList$Node<TE;>;"},
		{S: "Ljava/util/LinkedList$Node<TE;>;"},
	}
	for _, c := range cases {
		if err := ParseFieldTypeSignature(strings.NewReader(c.S)); err != nil {
			t.Errorf("Can't parse FieldTypeSignature %#v: %s", c.S, err)
		}
	}
}

func TestParseMethodTypeSignature(t *testing.T) {
	cases := []struct {
		S string
	}{
		// java.util.LinkedList
		{S: "(Ljava/util/Collection<+TE;>;)V"},
		{S: "(TE;)V"},
		{S: "(TE;Ljava/util/LinkedList$Node<TE;>;)V"},
		{S: "(Ljava/util/LinkedList$Node<TE;>;)TE;"},
		{S: "()TE;"},
		{S: "(TE;)Z"},
		{S: "(Ljava/util/Collection<+TE;>;)Z"},
		{S: "(ILjava/util/Collection<+TE;>;)Z"},
		{S: "(I)TE;"},
		{S: "(ITE;)TE;"},
		{S: "(ITE;)V"},
		{S: "(I)Ljava/util/LinkedList$Node<TE;>;"},
		{S: "(I)Ljava/util/ListIterator<TE;>;"},
		{S: "()Ljava/util/Iterator<TE;>;"},
		{S: "()Ljava/util/LinkedList<TE;>;"},
		{S: "<T:Ljava/lang/Object;>([TT;)[TT;"},
		{S: "()Ljava/util/Spliterator<TE;>;"},
	}
	for _, c := range cases {
		if err := ParseMethodTypeSignature(strings.NewReader(c.S)); err != nil {
			t.Errorf("Can't parse MethodTypeSignature %#v: %s", c.S, err)
		}
	}
}
