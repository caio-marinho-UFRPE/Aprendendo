import { CommonModule } from '@angular/common';
import { Component } from '@angular/core';
import { FormsModule } from '@angular/forms';
@Component({
  selector: 'app-calculadora',
  imports: [FormsModule, CommonModule],
  templateUrl: './calculadora.html',
  styleUrl: './calculadora.scss',
})
export class Calculadora {
  num1: number = 0;
  num2: number = 0;
  num3: number = 0;
  num4: number = 0;
  resultado: number = 0;

  calcularResultado() {
    console.log("Teste pra saber se tá pegando");
    this.resultado = this.num1 + this.num2 + this.num3 + this.num4;
  }
}
