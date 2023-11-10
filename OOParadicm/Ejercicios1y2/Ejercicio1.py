class ObjetoRepresentable:
    def __init__(self, nombre):
        self._nombre = nombre

    def mostrar(self):
        print(f"Objeto: {self._nombre}")

    def get_nombre(self):
        return self._nombre


class Texto(ObjetoRepresentable):
    def __init__(self, contenido):
        super().__init__("Texto")
        self._contenido = contenido

    def mostrar(self):
        super().mostrar()
        print(f"Contenido: {self._contenido}")

    def get_contenido(self):
        return self._contenido


class ObjetoGeometrico(ObjetoRepresentable):
    def __init__(self, tipo):
        super().__init__("Objeto Geométrico")
        self._tipo = tipo

    def mostrar(self):
        super().mostrar()
        print(f"Tipo: {self._tipo}")

    def get_tipo(self):
        return self._tipo


class Grupo(ObjetoRepresentable):
    def __init__(self, nombre):
        super().__init__("Grupo")
        self._nombre = nombre
        self._objetos = []

    def agregar_objeto(self, objeto):
        self._objetos.append(objeto)

    def mostrar(self):
        super().mostrar()
        print(f"Componentes del grupo {self._nombre}:")
        for objeto in self._objetos:
            objeto.mostrar()

    def get_objetos(self):
        return self._objetos


# Crear instancias para demostrar el uso exhaustivo de cada elemento

texto1 = Texto("Hola, mundo!")
texto2 = Texto("¡Hola desde el grupo!")

circulo = ObjetoGeometrico("Círculo")
rectangulo = ObjetoGeometrico("Rectángulo")
linea = ObjetoGeometrico("Línea")

grupo_interno = Grupo("Grupo Interno")
grupo_interno.agregar_objeto(texto2)
grupo_interno.agregar_objeto(rectangulo)

grupo_externo = Grupo("Grupo Externo")
grupo_externo.agregar_objeto(texto1)
grupo_externo.agregar_objeto(circulo)
grupo_externo.agregar_objeto(linea)
grupo_externo.agregar_objeto(grupo_interno)

# Mostrar la estructura del documento
grupo_externo.mostrar()