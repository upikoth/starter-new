package functionswithneeds

import (
	"context"
	"errors"
	"testing"
	"time"
)

func Test_validate(t *testing.T) {
	type args struct {
		functions FunctionsWithNeeds
	}

	foo := func(_ context.Context) error { return nil }
	foo1 := func(_ context.Context) error { return nil }
	foo2 := func(_ context.Context) error { return nil }
	foo3 := func(_ context.Context) error { return nil }
	bar := func(_ context.Context) error { return nil }
	bar1 := func(_ context.Context) error { return nil }
	bar2 := func(_ context.Context) error { return nil }
	bar3 := func(_ context.Context) error { return nil }

	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name:    "Empty functions",
			args:    args{},
			wantErr: false,
		},
		{
			name: "2 functions and no needs",
			args: args{
				functions: FunctionsWithNeeds{
					{
						Function: foo,
						Needs:    nil,
					},
					{
						Function: bar,
						Needs:    []func(ctx context.Context) error{},
					},
				},
			},
			wantErr: false,
		},
		{
			name: "2 function depends on each other",
			args: args{
				functions: FunctionsWithNeeds{
					{
						Function: foo,
						Needs:    []func(ctx context.Context) error{bar},
					},
					{
						Function: bar,
						Needs:    []func(ctx context.Context) error{foo},
					},
				},
			},
			wantErr: true,
		},
		{
			name: "Functions depends on other functions",
			args: args{
				functions: FunctionsWithNeeds{
					{
						Function: foo,
						Needs:    nil,
					}, {
						Function: foo1,
						Needs:    []func(ctx context.Context) error{foo},
					}, {
						Function: foo2,
						Needs:    []func(ctx context.Context) error{foo, foo1},
					}, {
						Function: foo3,
						Needs:    []func(ctx context.Context) error{foo, foo1, foo2},
					},
					{
						Function: bar,
						Needs:    nil,
					},
					{
						Function: bar1,
						Needs:    nil,
					},
					{
						Function: bar2,
						Needs:    nil,
					}, {
						Function: bar3,
						Needs:    []func(ctx context.Context) error{bar, bar1, bar2, foo3},
					},
				},
			},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := validate(tt.args.functions); (err != nil) != tt.wantErr {
				t.Errorf("validate() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_broadcastMessage(t *testing.T) {
	type args struct {
		chs     []chan string
		message string
	}

	tests := []struct {
		name string
		args args
	}{
		{
			name: "Example",
			args: args{
				chs:     make([]chan string, 3),
				message: "message example",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			go func() {
				for _, ch := range tt.args.chs {
					if message := <-ch; message != tt.args.message {
						t.Errorf("broadcastMessage() = %v, want %v", message, tt.args.message)
					}
				}
			}()

			go broadcastMessage(tt.args.chs, tt.args.message)
		})
	}
}

func Test_start(t *testing.T) {
	type args struct {
		ctx       context.Context
		functions FunctionsWithNeeds
	}

	foo := func(_ context.Context) error {
		time.Sleep(time.Millisecond * 50)
		return nil
	}
	foo1 := func(_ context.Context) error {
		time.Sleep(time.Millisecond * 50)
		return nil
	}
	foo2 := func(_ context.Context) error {
		time.Sleep(time.Millisecond * 50)
		return nil
	}
	foo3 := func(_ context.Context) error {
		time.Sleep(time.Millisecond * 50)
		return nil
	}
	bar := func(_ context.Context) error {
		time.Sleep(time.Millisecond * 50)
		return nil
	}
	bar1 := func(_ context.Context) error {
		time.Sleep(time.Millisecond * 50)
		return nil
	}
	bar2 := func(_ context.Context) error {
		time.Sleep(time.Millisecond * 50)
		return nil
	}
	bar3 := func(_ context.Context) error {
		time.Sleep(time.Millisecond * 50)
		return nil
	}
	errorFunc := func(_ context.Context) error {
		time.Sleep(time.Millisecond * 100)
		return errors.New("error")
	}

	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "Without errors",
			args: args{
				ctx: context.Background(),

				functions: FunctionsWithNeeds{
					{
						Function: foo,
						Needs:    nil,
					},
					{
						Function: foo1,
						Needs:    []func(ctx context.Context) error{foo},
					},
					{
						Function: foo2,
						Needs:    []func(ctx context.Context) error{foo, foo1},
					},
					{
						Function: foo3,
						Needs:    []func(ctx context.Context) error{foo, foo1, foo2},
					},
					{
						Function: bar,
						Needs:    nil,
					},
					{
						Function: bar1,
						Needs:    nil,
					},
					{
						Function: bar2,
						Needs:    nil,
					},
					{
						Function: bar3,
						Needs:    []func(ctx context.Context) error{bar, bar1, bar2, foo3},
					},
				},
			},
			wantErr: false,
		},
		{
			name: "With error",
			args: args{
				ctx: context.Background(),

				functions: FunctionsWithNeeds{
					{
						Function: foo,
						Needs:    nil,
					},
					{
						Function: foo1,
						Needs:    []func(ctx context.Context) error{foo},
					},
					{
						Function: errorFunc,
						Needs:    []func(ctx context.Context) error{foo, foo1},
					},
					{
						Function: foo3,
						Needs:    []func(ctx context.Context) error{foo, foo1, errorFunc},
					},
					{
						Function: bar,
						Needs:    nil,
					},
					{
						Function: bar1,
						Needs:    nil,
					},
					{
						Function: bar2,
						Needs:    nil,
					},
					{
						Function: bar3,
						Needs:    []func(ctx context.Context) error{bar, bar1, bar2, foo3},
					},
				},
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := start(tt.args.ctx, tt.args.functions); (err != nil) != tt.wantErr {
				t.Errorf("start() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
