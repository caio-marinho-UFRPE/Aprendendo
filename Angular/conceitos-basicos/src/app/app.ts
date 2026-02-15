import { Component, signal } from '@angular/core';
import { Calculadora } from './calculadora/calculadora'
import { ListaDCompras } from './lista-d-compras/lista-d-compras';

@Component({
  selector: 'app-root',
  imports: [Calculadora, ListaDCompras],
  templateUrl: './app.html',
  styleUrl: './app.scss'
})
export class App {
  protected readonly title = signal('conceitos-basicos');
}
