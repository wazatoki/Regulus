import { Directive, ElementRef, OnInit } from '@angular/core';
import { AutofocusDirective } from './autofocus.directive';

describe('AutofocusDirective', () => {

  it('should create an instance', () => {
    const directive = new AutofocusDirective( null );
    expect(directive).toBeTruthy();
  });
});
