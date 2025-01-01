package builder

import "testing"

func Test_processWithYaml(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		f       []byte
		want    Metadata
		want2   []byte
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got2, gotErr := processWithYaml(tt.f)
			if gotErr != nil {
				if !tt.wantErr {
					t.Errorf("processWithYaml() failed: %v", gotErr)
				}
				return
			}
			if tt.wantErr {
				t.Fatal("processWithYaml() succeeded unexpectedly")
			}
			// TODO: update the condition below to compare got with tt.want.
			if true {
				t.Errorf("processWithYaml() = %v, want %v", got, tt.want)
			}
			if true {
				t.Errorf("processWithYaml() = %v, want %v", got2, tt.want2)
			}
		})
	}
}
