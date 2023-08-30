package bot

import "testing"

func TestParseTask(t *testing.T) {
	type args struct {
		msg string
	}
	tests := []struct {
		name    string
		args    args
		want    int
		wantErr bool
	}{
		{
			name: "fail case",
			args: args{
				msg: "Ответь на это: ",
			},
			want:    0,
			wantErr: true,
		},
		{
			name: "plus case",
			args: args{
				msg: "Ответь на это: 2 + 2",
			},
			want:    4,
			wantErr: false,
		},
		{
			name: "multiply case",
			args: args{
				msg: "Ответь на это: 3 * 2",
			},
			want:    6,
			wantErr: false,
		},
		{
			name: "minus case",
			args: args{
				msg: "Ответь на это: 3 - 2",
			},
			want:    1,
			wantErr: false,
		},
		{
			name: "minus negative case",
			args: args{
				msg: "Ответь на это: 2 - 3",
			},
			want:    -1,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ParseTask(tt.args.msg)
			if (err != nil) != tt.wantErr {
				t.Errorf("ParseTask() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("ParseTask() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_toFixed(t *testing.T) {
	type args struct {
		num       float64
		precision int
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		{
			name: "test",
			args: args{
				num:       0.51235,
				precision: 2,
			},
			want: 0.51,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := toFixed(tt.args.num, tt.args.precision); got != tt.want {
				t.Errorf("toFixed() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetTitle(t *testing.T) {
	type args struct {
		title int64
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "test 0",
			args: args{
				title: 0,
			},
			want: "Новичек I",
		},
		{
			name: "test 4",
			args: args{
				title: 4,
			},
			want: "Новичек V",
		},
		{
			name: "test 5",
			args: args{
				title: 5,
			},
			want: "Драчун I",
		},
		{
			name: "test 39",
			args: args{
				title: 39,
			},
			want: "Рыцарь-чемпион V",
		},
		{
			name: "test 40",
			args: args{
				title: 40,
			},
			want: "МАСТЕР УРОВНЯ(1)",
		},
		{
			name: "test 55",
			args: args{
				title: 55,
			},
			want: "МАСТЕР УРОВНЯ(16)",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetTitle(tt.args.title); got != tt.want {
				t.Errorf("GetTitle() = %v, want %v", got, tt.want)
			}
		})
	}
}
