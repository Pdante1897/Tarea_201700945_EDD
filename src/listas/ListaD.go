package listas

type Nodo struct {
	dato      *Mensaje
	anterior  *Nodo
	siguiente *Nodo
}
type ListaDoble struct {
	inicio *Nodo `Mensajes`
	fin    *Nodo
}
type Mensaje struct {
	Origen  string `Origen`
	Destino string `Destino`
	Msg     []*Ms  `Msg`
}
type Ms struct {
	Fecha string `Fecha`
	Texto string `Texto`
}

func (this Ms) To_String() string {
	return this.Fecha + " " + this.Texto
}

func (this Mensaje) To_String() string {

	return "Origen: " + this.Origen + " Destino: " + this.Destino
}
func (this *ListaDoble) Vacio() bool {
	if this.inicio == nil {
		return true
	} else {
		return false
	}

}
func (this ListaDoble) Insertar(dato *Mensaje) {
	aux := Nodo{
		dato:      dato,
		siguiente: nil,
		anterior:  nil,
	}
	if this.Vacio() {
		this.inicio = &aux
		this.fin = this.inicio
	} else {
		aux.anterior = this.fin
		this.fin.siguiente = &aux
		this.fin = this.fin.siguiente

	}
}
