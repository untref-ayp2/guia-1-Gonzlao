package ordenamientos

import (
	//"reflect"
	"reflect"
	"testing"
	//"github.com/stretchr/testify/assert"
)

func TestOrdenamientoBorbujeo(t *testing.T) {
	// Test autogenerate
	type args struct {
		f []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "Test ordenamineto por borbujeo",
			args: args{f: []int{9, 2, 6, 1, 8, 4, 7, 3, 5}},
			want: []int{1, 2, 3, 4, 5, 6, 7, 8, 9},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := OrdenamientoBorbujeo(tt.args.f); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("OrdenamientoBorbujeo() = %v, want %v", got, tt.want)
			}
		})
	}

	//Test custom
	/*arreglo := []int{9, 2, 6, 1, 8, 4, 7, 3, 5}
	arreglo_deseado := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	a_ordenado := OrdenamientoBorbujeo(arreglo)

	if !reflect.DeepEqual(a_ordenado, arreglo_deseado) {
		t.Errorf("arreglo ordenado = %v, arreglo deseado %v", a_ordenado, arreglo_deseado)
	}*/

	//Test con assert
	/*arreglo := []int{9, 2, 6, 1, 8, 4, 7, 3, 5}
	arreglo_deseado := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	a_ordenado := OrdenamientoBorbujeo(arreglo)

	if !assert.ObjectsAreEqual(a_ordenado, arreglo_deseado) {
		t.Errorf("arreglo ordenado = %v, arreglo deseado %v", a_ordenado, arreglo_deseado)
	}*/
}
