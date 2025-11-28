package com.example.demo;

import org.apache.commons.lang3.StringUtils;
import com.google.gson.Gson;
import java.util.HashMap;
import java.util.Map;

public class App {
    public static void main(String[] args) {
        System.out.println("=== Maven Demo Application ===\n");

        // Using Apache Commons Lang
        String message = "hello maven";
        String capitalized = StringUtils.capitalize(message);
        System.out.println("Original: " + message);
        System.out.println("Capitalized: " + capitalized);

        // Using Gson
        Map<String, Object> data = new HashMap<>();
        data.put("tool", "Maven");
        data.put("language", "Java");
        data.put("version", "1.0.0");
        data.put("dependencies", 2);

        Gson gson = new Gson();
        String json = gson.toJson(data);
        
        System.out.println("\nProject Info (JSON):");
        System.out.println(json);

        System.out.println("\n✅ Maven demo completed successfully!");
    }
}
