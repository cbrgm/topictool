package main

import (
	"reflect"
	"testing"
)

func Test_removeDuplicateTopics(t *testing.T) {
	type args struct {
		topics []string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			name: "must remove duplicates",
			args: args{
				topics: []string{"this", "this", "is", "a", "a", "test"},
			},
			want: []string{"this", "is", "a", "test"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := removeDuplicateTopics(tt.args.topics); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("removeDuplicateTopics() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_removeFromTopics(t *testing.T) {
	type args struct {
		topics   []string
		toRemove []string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			name: "must be removed",
			args: args{
				topics:   []string{"this", "is", "a", "test"},
				toRemove: []string{"this", "a"},
			},
			want: []string{"is", "test"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := removeFromTopics(tt.args.topics, tt.args.toRemove); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("removeFromTopics() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_topicsToStr(t *testing.T) {
	type args struct {
		topics []string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "must be formatted",
			args: args{
				topics: []string{"this", "is", "a", "test"},
			},
			want: "this,is,a,test",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := topicsToStr(tt.args.topics); got != tt.want {
				t.Errorf("topicsToStr() = %v, want %v", got, tt.want)
			}
		})
	}
}
