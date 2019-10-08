package graph

import (
	"reflect"
	"testing"
)

func TestEnqueue(t *testing.T) {
	tests := []struct {
		input []string
		want  struct {
			q   queue
			len int
		}
	}{
		{
			input: []string{"a", "b", "c"},
			want: struct {
				q   queue
				len int
			}{
				q: queue{
					front: &queueItem{
						value: "a",
						next: &queueItem{
							value: "b",
							next: &queueItem{
								value: "c",
							},
						},
					},
					back: &queueItem{
						value: "c",
					},
					len: 3,
				},
			},
		},
	}

	for _, test := range tests {
		var got queue
		for _, x := range test.input {
			got.enqueue(x)
		}

		if !reflect.DeepEqual(got, test.want.q) {
			t.Errorf("%+v != %+v", got, test.want.q)
		}
	}
}

func TestDequeue(t *testing.T) {
	tests := []struct {
		input queue
		want  []string
	}{
		{
			input: queue{
				front: &queueItem{
					value: "a",
					next: &queueItem{
						value: "b",
						next: &queueItem{
							value: "c",
						},
					},
				},
				back: &queueItem{
					value: "c",
				},
				len: 3,
			},
			want: []string{"a", "b", "c"},
		},
	}

	for _, test := range tests {
		got := make([]string, 0)
		for next := test.input.dequeue(); next != nil; next = test.input.dequeue() {
			got = append(got, next.(string))
		}

		if !reflect.DeepEqual(got, test.want) {
			t.Errorf("%+v != %+v", got, test.want)
		}
	}
}
