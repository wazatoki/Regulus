import { Injectable } from '@angular/core';
import { Maker } from './maker';

@Injectable({
    providedIn: 'root'
})
export class MakerCondition {

    constructor(private maker: Maker) { }

    
}
