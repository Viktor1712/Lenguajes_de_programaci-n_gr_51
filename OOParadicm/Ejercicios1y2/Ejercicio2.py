from datetime import datetime, timedelta

class Socio:
    def __init__(self, numero, nombre, direccion):
        self.numero = numero
        self.nombre = nombre
        self.direccion = direccion
        self.libros_prestados = []

    def prestar_libro(self, libro, fecha_prestamo):
        if libro.disponible:
            libro.prestar(self.numero, fecha_prestamo)
            self.libros_prestados.append(libro.codigo)
            print(f"{self.nombre} ha tomado prestado el libro '{libro.titulo}'")

    def mostrar_libros_prestados(self):
        print(f"{self.nombre} tiene los siguientes libros prestados:")
        for codigo_libro in self.libros_prestados:
            print(f"- Código: {codigo_libro}")

class Libro:
    def __init__(self, codigo, titulo, autor, disponible=True):
        self.codigo = codigo
        self.titulo = titulo
        self.autor = autor
        self.disponible = disponible

    def prestar(self, numero_socio, fecha_prestamo):
        self.disponible = False
        print(f"Libro '{self.titulo}' prestado al socio {numero_socio} el {fecha_prestamo}")

    def devolver(self):
        self.disponible = True
        print(f"Libro '{self.titulo}' devuelto a la biblioteca")

class Prestamo:
    def __init__(self, numero_socio, codigo_libro, fecha_prestamo):
        self.numero_socio = numero_socio
        self.codigo_libro = codigo_libro
        self.fecha_prestamo = fecha_prestamo

# Crear instancias y demostrar el funcionamiento del sistema

# Crear socios
socio1 = Socio(1, "Juan Pérez", "Calle 123")
socio2 = Socio(2, "María Gómez", "Avenida XYZ")

# Crear libros
libro1 = Libro(101, "Harry Potter", "J.K. Rowling")
libro2 = Libro(102, "Cien años de soledad", "Gabriel García Márquez")
libro3 = Libro(103, "El Señor de los Anillos", "J.R.R. Tolkien")

# Prestar libros a los socios
fecha_prestamo = datetime.now()
socio1.prestar_libro(libro1, fecha_prestamo)
socio1.prestar_libro(libro2, fecha_prestamo)
socio2.prestar_libro(libro3, fecha_prestamo)

# Mostrar libros prestados por cada socio
socio1.mostrar_libros_prestados()
socio2.mostrar_libros_prestados()

# Devolver un libro a la biblioteca
libro_devuelto = libro1
libro_devuelto.devolver()

# Mostrar el estado de los libros después de la devolución
print(f"Estado actual de los libros:")
print(f"Libro '{libro1.titulo}': {'Disponible' if libro1.disponible else 'No disponible'}")
print(f"Libro '{libro2.titulo}': {'Disponible' if libro2.disponible else 'No disponible'}")
print(f"Libro '{libro3.titulo}': {'Disponible' if libro3.disponible else 'No disponible'}")

# Crear un nuevo socio con más de 3 libros prestados
socio3 = Socio(3, "Ana Rodríguez", "Avenida ABC")

libro4 = Libro(104, "Don Quijote de la Mancha", "Miguel de Cervantes")
libro5 = Libro(105, "1984", "George Orwell")
libro6 = Libro(106, "Orgullo y prejuicio", "Jane Austen")
libro7 = Libro(107, "Crónica de una muerte anunciada", "Gabriel García Márquez")

# Realizar préstamos a socio3
socio3.prestar_libro(libro4, fecha_prestamo)
socio3.prestar_libro(libro5, fecha_prestamo)
socio3.prestar_libro(libro6, fecha_prestamo)
socio3.prestar_libro(libro7, fecha_prestamo)

# Mostrar libros prestados por socio3
socio3.mostrar_libros_prestados()

# Encontrar socios con más de 3 libros prestados (actualizado)
socios_con_mas_de_3_libros = list(filter(lambda s: len(s.libros_prestados) > 3, [socio1, socio2, socio3]))
print("Socios con más de 3 libros prestados:")
for socio in socios_con_mas_de_3_libros:
    print(f"- {socio.nombre}")
