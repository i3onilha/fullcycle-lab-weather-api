package services

import "testing"

func TestCEPService_ValidateCEP(t *testing.T) {
	tests := []struct {
		name string
		cep  string
		want bool
	}{
		{
			name: "valid CEP with 8 digits",
			cep:  "01310100",
			want: true,
		},
		{
			name: "valid CEP with formatting",
			cep:  "01310-100",
			want: true,
		},
		{
			name: "invalid CEP with less than 8 digits",
			cep:  "01310",
			want: false,
		},
		{
			name: "invalid CEP with more than 8 digits",
			cep:  "013101001",
			want: false,
		},
		{
			name: "empty CEP",
			cep:  "",
			want: false,
		},
	}

	service := NewCEPService()
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := service.ValidateCEP(tt.cep)
			if got != tt.want {
				t.Errorf("ValidateCEP() = %v, want %v", got, tt.want)
			}
		})
	}
}
