import { ComponentFixture, TestBed } from '@angular/core/testing';

import { ListaDCompras } from './lista-d-compras';

describe('ListaDCompras', () => {
  let component: ListaDCompras;
  let fixture: ComponentFixture<ListaDCompras>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      imports: [ListaDCompras]
    })
    .compileComponents();

    fixture = TestBed.createComponent(ListaDCompras);
    component = fixture.componentInstance;
    await fixture.whenStable();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
