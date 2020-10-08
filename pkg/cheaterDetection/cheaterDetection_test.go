package cheaterDetection

import (
	"math/big"
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
	"moritzm-mueller.de/tss/pkg/shamir"
)

func TestCalculateCheckValue(t *testing.T) {
	type args struct {
		shares []shamir.Share
	}
	tests := []struct {
		name    string
		args    args
		want    AntiCheat
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := CalculateCheckValue(tt.args.shares)
			if (err != nil) != tt.wantErr {
				t.Errorf("CalculateCheckValue() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CalculateCheckValue() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_calcT(t *testing.T) {
	type args struct {
		shares [][]byte
		prime  *big.Int
	}
	tests := []struct {
		name    string
		args    args
		want    *big.Int
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := calcT(tt.args.shares, tt.args.prime)
			if (err != nil) != tt.wantErr {
				t.Errorf("calcT() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("calcT() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_calcSum1(t *testing.T) {
	assert.NotEmpty(t, calcSum1([][]byte{{1, 2, 3, 4, 5}, {1, 2, 3, 4, 5}}, big.NewInt(13)))
}

func Test_calcSum2(t *testing.T) {
	{
		assert.Equal(t, int64(1050), calcSum2(big.NewInt(3), big.NewInt(7), 3).Int64())
	}
}

func Test_decodeSecret(t *testing.T) {
	type args struct {
		share shamir.Share
	}
	tests := []struct {
		name string
		args args
		want []byte
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := decodeSecret(tt.args.share); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("decodeSecret() = %v, want %v", got, tt.want)
			}
		})
	}
}
