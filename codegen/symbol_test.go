package codegen

import "testing"

func TestSymbolGenLoadOpcode(t *testing.T) {
	type fields struct {
		Name       string
		SymbolType SymbolType
		PascalType PascalType
		Index      int
		Global     bool
	}
	type args struct {
		className string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   string
	}{
		{
			name: "global integer variable",
			fields: fields{
				Name:       "myintvar",
				SymbolType: Variable,
				PascalType: Integer,
				Index:      -1,
				Global:     true,
			},
			args: args{
				className: "MyClassInt",
			},
			want: "getstatic MyClassInt.myintvar I",
		},
		{
			name: "local integer variable",
			fields: fields{
				Name:       "myintvar",
				SymbolType: Variable,
				PascalType: Integer,
				Index:      22,
				Global:     false,
			},
			args: args{
				className: "MyClassInt",
			},
			want: "iload 22",
		},
		{
			name: "global string variable",
			fields: fields{
				Name:       "mystringvar",
				SymbolType: Variable,
				PascalType: String,
				Index:      -1,
				Global:     true,
			},
			args: args{
				className: "MyClassString",
			},
			want: "getstatic MyClassString.mystringvar java/lang/String",
		},
		{
			name: "local string variable",
			fields: fields{
				Name:       "mystringvar",
				SymbolType: Variable,
				PascalType: String,
				Index:      25,
				Global:     false,
			},
			args: args{
				className: "MyClassString",
			},
			want: "aload 25",
		},
		{
			name: "global boolean variable",
			fields: fields{
				Name:       "myboolvar",
				SymbolType: Variable,
				PascalType: Boolean,
				Index:      -1,
				Global:     true,
			},
			args: args{
				className: "MyClassBool",
			},
			want: "getstatic MyClassBool.myboolvar I",
		},
		{
			name: "local boolean variable",
			fields: fields{
				Name:       "myboolvar",
				SymbolType: Variable,
				PascalType: Boolean,
				Index:      28,
				Global:     false,
			},
			args: args{
				className: "MyClassBool",
			},
			want: "iload 28",
		},
		{
			name: "invalid type",
			fields: fields{
				Name:       "myboolvar",
				SymbolType: Function,
				PascalType: Boolean,
				Index:      -1,
				Global:     true,
			},
			args: args{
				className: "MyClassBool",
			},
			want: "invalid symbol type",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := Symbol{
				Name:       tt.fields.Name,
				SymbolType: tt.fields.SymbolType,
				PascalType: tt.fields.PascalType,
				Index:      tt.fields.Index,
				Global:     tt.fields.Global,
			}
			if got := s.GenLoadOpcode(tt.args.className); got != tt.want {
				t.Errorf("Symbol.GenLoadOpcode() = %v, want %v", got, tt.want)
			}
		})
	}
}
