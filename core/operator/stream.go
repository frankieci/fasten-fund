package stream

import "sort"

/**
 * Created by frankieci on 2022/3/16 16:21 pm
 */

type Element interface {
	Identify() interface{}
}

type Stream struct {
	sources <-chan Element
}

// Just converts the given Element items to a Stream.
func Just(elements ...Element) Stream {
	sources := make(chan Element, len(elements))
	for _, e := range elements {
		sources <- e
	}
	close(sources)

	return Range(sources)
}

// Range converts the given channel to a Stream.
func Range(sources <-chan Element) Stream {
	return Stream{sources: sources}
}

type (
	Key = interface{}
	// KeyFunc defines the method to generate keys for the Elements in a Stream.
	KeyFunc func(e Element) Key

	// ReduceFunc defines the method to reduce all the Elements in a Stream.
	ReduceFunc func(pipe <-chan Element) (Element, error)

	// ForEachFunc defines the method to handle each element in a Stream.
	ForEachFunc func(e Element)

	// CompareFunc defines the method to compare the Elements in a Stream.
	CompareFunc func(a, b Element) bool
)

var _ Element = Elements{}

type Elements []Element

func (es Elements) Identify() interface{} {
	return nil
}

// Group groups the Elements into different groups based on their keys.
func (s Stream) Group(fn KeyFunc) Stream {
	groups := make(map[Key]Elements)
	for e := range s.sources {
		key := fn(e)
		groups[key] = append(groups[key], e)
	}

	source := make(chan Element)
	go func() {
		for _, group := range groups {
			source <- group
		}
		close(source)
	}()

	return Range(source)
}

// Reduce is a utility method to let the caller deal with the underlying channel.
func (s Stream) Reduce(fn ReduceFunc) (Element, error) {
	return fn(s.sources)
}

// Sort sorts the Elements from the underlying source.
func (s Stream) Sort(compare CompareFunc) Stream {
	var es Elements
	for e := range s.sources {
		es = append(es, e)
	}
	sort.Slice(es, func(i, j int) bool {
		return compare(es[i], es[j])
	})

	return Just(es...)
}

// ForEach seals the Stream with the ForEachFunc on each element, no successive operations.
func (s Stream) ForEach(fn ForEachFunc) {
	for source := range s.sources {
		fn(source)
	}
}
