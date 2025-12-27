package com.Marin.Aprendendo;


import org.springframework.web.bind.annotation.GetMapping;
import org.springframework.web.bind.annotation.RestController;

@RestController
public class OlaMundoController {

    @GetMapping(path = "/Ola")
    public String olaMundo() {
        return "Olá, Mundo!";
    }
}
