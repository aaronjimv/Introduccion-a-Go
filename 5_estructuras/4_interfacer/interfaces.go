package main 

import "fmt"

/*
	las interfaces en go se caracteriza por:
		1) crea metodos vacios para implementar
		2) es un tipo de dato
*/
type Usuario interface{

	//nombre_del_metodo() valor_que_retorna
	Permisos() int // 1-5
	Nombre() string
}

type Admin struct{

	nombre string
}

// no hay una palabre clave o una implementacion, 
//con solo colocar el mismo nombre del metodo se implementa la interface

func (this Admin) Permisos() int{
	return 5
}

func (this Admin) Nombre() string{
	return this.nombre
}


type Editor struct{

	nombre string
}

func (this Editor) Permisos() int{
	return 3
}

func (this Editor) Nombre() string{
	return this.nombre
}


type Lector struct{

	nombre string
}

func (this Lector) Permisos() int{
	return 2
}

func (this Lector) Nombre() string{
	return this.nombre
}



// como la interface es un tipo de dato se pude usar pasar como parametro
func autentificar(user Usuario) string{
	if user.Permisos() > 4 {
		return user.Nombre() + " tiene permisos de administrador"
	}else if user.Permisos() == 3 {
		return user.Nombre() + " tiene permisos de editor"
	}

	return user.Nombre() + " no tiene permisos, solo puede leer"
}


func main() {
	/*
	admin := Admin{"Aaron"}
	editor := Editor{"Sebastian"}

	fmt.Println(autentificar(admin)) // aunque Admin es una estructura es de tipo Usuario por la interface
	fmt.Println(autentificar(editor))
	*/

	//como es un tipo de dato tambien se pueden hacer un arreglo de ellos
	usuarios := [] Usuario{ Admin{"Aaron"}, Editor{"Sebastian"}, Lector{"Reste"} }

	for _, usuario := range usuarios{
		fmt.Println(autentificar(usuario))
	}

	// las interfaces son como las clases abstractas
}