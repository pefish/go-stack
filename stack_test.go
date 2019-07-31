package go_stack

import "testing"

func TestStackClass_GetStack(t *testing.T) {
	tests := []struct {
		name string
		this *StackClass
		want string
	}{
		{
			name: `test`,
			this: &Stack,
			want: "/usr/local/Cellar/go/1.11.4/libexec/src/testing/testing.go --- 827\n",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			this := &StackClass{}
			if got := this.GetStack(); got != tt.want {
				t.Errorf("StackClass.GetStack() = %v, want %v", got, tt.want)
			}
		})
	}
}
