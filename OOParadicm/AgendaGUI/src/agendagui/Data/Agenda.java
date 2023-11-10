package agendagui.Data;

import java.util.LinkedList;
import java.util.Map;

public class Agenda {
    private LinkedList<Object> listaObjetos;
    private LinkedList<String> contactos;
    private LinkedList<String> eventos;
    private Map<String,Class> clasesCargadas;
    private Map<String,Class> clasesInstanciables;
    private Object instance;

    // Asegúrate de tener un constructor público
    public Agenda() {
        this.listaObjetos = new LinkedList<>();
        this.contactos = new LinkedList<>();
        this.eventos = new LinkedList<>();
    }

    // Método para obtener la instancia única de Agenda (Lazy Singleton)
    public  Agenda getInstance() {
        if (instance == null) {
            instance = new Agenda();
        }
        return (Agenda) instance;
    }

    public LinkedList<Object> getListaObjetos() {
        return listaObjetos;
    }

    public LinkedList<String> getContactos() {
        return contactos;
    }

    public LinkedList<String> getEventos() {
        return eventos;
    }

    public Map<String, Class> getClasesCargadas() {
        return clasesCargadas;
    }

    public Map<String, Class> getClasesInstanciables() {
        return clasesInstanciables;
    }

    public void setClasesCargadas(Map<String, Class> clasesCargadas) {
        this.clasesCargadas = clasesCargadas;
    }

    public void setClasesInstanciables(Map<String, Class> clasesInstanciables) {
        this.clasesInstanciables = clasesInstanciables;
    }
}
