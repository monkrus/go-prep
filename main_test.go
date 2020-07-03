package main

import "testing"

func TestWrongType(t *testing.T) {
	type args struct {
		value string
	}

	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{ //One Word Test
			name: "Hello test",
			args: args{
				value: "Hello",
			},
			wantErr: false,
		},
		{ //Two Words Test
			name: "Viktor test",
			args: args{
				value: "Viktor Ivanov",
			},
			wantErr: false,
		},
		{ //Test with numerics and chars
			name: "R2D2 test",
			args: args{
				value: "R2D2",
			},
			wantErr: false,
		},
		{ //Test with numerics (return err)
			name: "123 test",
			args: args{
				value: "123",
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := WrongType(tt.args.value); (err != nil) != tt.wantErr {
				t.Errorf("WrongType() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}