package logx

import (
	"context"
	"testing"

	"github.com/jackc/pgx/v4"
)

func TestPGX_Log(t *testing.T) {
	type args struct {
		ctx   context.Context
		level pgx.LogLevel
		msg   string
		data  map[string]interface{}
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "Success",
			args: args{
				ctx:   context.Background(),
				level: pgx.LogLevelInfo,
				msg:   "testing",
				data: map[string]interface{}{
					"sql": " INSERT INTO cronx_histories (\tid,\tcreated_at,\tname,\tstatus,\tstatus_code,\tstarted_at,\tfinished_at,\tlatency,\tmetadata   )   VALUES (\t  $1,\t  $2,\t  $3,\t  $4,\t  $5,\t  $6,\t  $7,\t  $8,\t  $9   )\n;  ",
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &PGX{}
			p.Log(tt.args.ctx, tt.args.level, tt.args.msg, tt.args.data)
		})
	}
}
