import { Component, signal } from '@angular/core';
import { Pagina } from './pagina/pagina'
import { HelloWorldComponent } from './helloworld/helloworld.component'

@Component({
  selector: 'app-root',
  imports: [HelloWorldComponent, Pagina],
  templateUrl: './app.html',
  styleUrl: './app.scss'
})
export class App {
  protected readonly title = signal('conceitos-basicos');
}
