package agendagui;

import agendagui.Data.Agenda;
import agendagui.Data.ObjAgenda;
import agendagui.GUI.AgendaMain;

import java.io.File;
import java.io.IOException;
import java.net.URL;
import java.net.URLClassLoader;
import java.util.ArrayList;
import java.util.Arrays;
import java.util.HashMap;
import java.util.Iterator;
import java.util.LinkedList;
import java.util.Map;
import java.util.jar.JarFile;
import java.util.logging.Level;
import java.util.logging.Logger;

public class AgendaGUI {

    static final String PLUGIN_FOLDER = "src/Plugins";

    public static void main(String[] args) {

        Agenda miAgenda = new Agenda();

        miAgenda.setClasesCargadas(cargarPlugins());

        miAgenda.setClasesInstanciables(cargarClasesInstanciables(miAgenda.getClasesCargadas()));

        Iterator<Map.Entry<String, Class>> iter = miAgenda.getClasesInstanciables().entrySet().iterator();
        if (iter != null) {
            while (iter.hasNext()) {
                Map.Entry<String, Class> entry = iter.next();
                miAgenda.getContactos().add(entry.getKey());
            }
        }

        AgendaMain guiMain = new AgendaMain(miAgenda);
        guiMain.setVisible(true);
    }

    private static Map<String, Class> cargarPlugins() {
        Map<String, Class> classList = new HashMap<>();
        File pluginFolder = new File(PLUGIN_FOLDER);
        if (!pluginFolder.exists()) {
            if (pluginFolder.mkdirs()) {
                System.out.println("Created plugin folder");
            }
        }
        File[] files = pluginFolder.listFiles((dir, name) -> name.endsWith(".jar"));
        ArrayList<URL> urls = new ArrayList<>();
        ArrayList<String> classes = new ArrayList<>();
        if (files != null) {
            Arrays.stream(files).forEach(file -> {
                try {
                    JarFile jarFile = new JarFile(file);
                    urls.add(new URL("jar:file:" + PLUGIN_FOLDER + "/" + file.getName() + "!/"));
                    jarFile.stream().forEach(jarEntry -> {
                        if (jarEntry.getName().endsWith(".class")) {
                            classes.add(jarEntry.getName());
                        }
                    });
                } catch (IOException e) {
                    e.printStackTrace();
                }
            });
            URLClassLoader pluginLoader = new URLClassLoader(urls.toArray(new URL[urls.size()]));

            classes.forEach(s -> {
                try {
                    String className = s.replaceAll("/", ".").replace(".class", "");
                    int lastDotIndex = className.lastIndexOf(".");
                    String name = (lastDotIndex != -1) ? className.substring(lastDotIndex + 1) : className;
                    Class<?> classs = pluginLoader.loadClass(className);
                    classList.put(name, classs);
                } catch (ClassNotFoundException ex) {
                    Logger.getLogger(AgendaGUI.class.getName()).log(Level.SEVERE, null, ex);
                }
            });
        }
        return classList;
    }

    public static Map<String, Class> cargarClasesInstanciables(Map<String, Class> listaClases) {
        Map<String, Class> clasesInstanciables = new HashMap<>();
        listaClases.forEach((c, v) -> {
            Class[] interfaces = v.getInterfaces();
            for (Class anInterface : interfaces) {
                if (anInterface == ObjAgenda.class) {
                    clasesInstanciables.put(c, v);
                }
            }
        });
        return clasesInstanciables;
    }
}
