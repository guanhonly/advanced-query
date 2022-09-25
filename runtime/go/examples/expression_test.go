package examples

import (
	"reflect"
	"testing"

	"github.com/guanhonly/advanced-query/runtime/go/entity"
	"github.com/guanhonly/advanced-query/runtime/go/handler/errlistener"
	"github.com/guanhonly/advanced-query/runtime/go/handler/kv"
)

func TestExpressionParser(t *testing.T) {
	type args struct {
		expression string
		kvHandler  kv.Handler
		el         errlistener.ErrorListener
	}
	tests := []struct {
		name    string
		args    args
		want    interface{}
		wantErr bool
	}{
		{
			name: "simple",
			args: args{
				expression: "field1:v1 AND field2:(v2 OR v3)",
				kvHandler:  kv.NewKVHandler(), el: errlistener.New(),
			},
			want: map[string][]string{
				"field1": {"v1"},
				"field2": {"v2", "v3"},
			},
		},
		{
			name: "complex",
			args: args{
				expression: "field1:v1 AnD field2:(not(v2 OR v3))",
				kvHandler:  kv.NewKVHandler(), el: errlistener.New(),
			},
			want: map[string][]string{
				"field1": {"v1"},
				"field2": {"v2", "v3"},
			},
		},
		{
			name: "special value",
			args: args{
				expression: `field1:"aNd" OR field2:nUll`,
				kvHandler:  kv.NewKVHandler(), el: errlistener.New(),
			},
			want: map[string][]*entity.ValueInfo{
				"field1": {{Value: `"aNd"`, StartIndex: 7, EndIndex: 11}},
				"field2": {{Value: "nUll", StartIndex: 23, EndIndex: 26, IsNull: true}},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res, err := kv.GetExpressionKVs(tt.args.expression, tt.args.kvHandler, tt.args.el)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetExpressionKVs got err: %v, want %v", err, tt.wantErr)
				return
			}
			var got interface{}
			switch tt.want.(type) {
			case map[string][]string:
				kvs := make(map[string][]string)
				for k, v := range res {
					var vals []string
					for _, val := range v {
						vals = append(vals, val.Value)
					}
					kvs[k] = vals
				}
				got = kvs
			case map[string][]*entity.ValueInfo:
				got = res
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetExpressionKVs got: %v, want: %v", got, tt.want)
			}
		})
	}
}
