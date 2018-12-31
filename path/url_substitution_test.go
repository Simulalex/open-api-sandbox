package path

import "testing"

func Test_urlSubstitutor_ToRegex(t *testing.T) {
	type args struct {
		swaggerPath string
	}
	tests := []struct {
		name      string
		subs      urlSubstitutor
		args      args
		wantRegex string
	}{
		{
			name: "No params",
			subs: urlSubstitutor{},
			args: args{
				swaggerPath: "/some/path",
			},
			wantRegex: "/some/path",
		},
		{
			name: "Single param",
			subs: urlSubstitutor{},
			args: args{
				swaggerPath: "/some/{param}/path",
			},
			wantRegex: `/some/[^/\s]+/path`,
		},
		{
			name: "Double param",
			subs: urlSubstitutor{},
			args: args{
				swaggerPath: "/some/{param}/path/{otherparam}",
			},
			wantRegex: `/some/[^/\s]+/path/[^/\s]+`,
		},
		{
			name: "Triple param",
			subs: urlSubstitutor{},
			args: args{
				swaggerPath: "/some/{param}/path/{otherparam}/sure/{one_more_param}",
			},
			wantRegex: `/some/[^/\s]+/path/[^/\s]+/sure/[^/\s]+`,
		},
		{
			name: "Missing closing bracked",
			subs: urlSubstitutor{},
			args: args{
				swaggerPath: "/some/{param/path",
			},
			wantRegex: "/some/{param/path",
		},
		{
			name: "Missing opening bracket",
			subs: urlSubstitutor{},
			args: args{
				swaggerPath: "/some/param}/path",
			},
			wantRegex: "/some/param}/path",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			subs := urlSubstitutor{}
			gotRegex := subs.ToRegex(tt.args.swaggerPath)

			if gotRegex != tt.wantRegex {
				t.Errorf("urlSubstitutor.ToRegex() = %v, want %v", gotRegex, tt.wantRegex)
			}
		})
	}
}
