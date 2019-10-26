import { Injectable } from '@angular/core';
import { Maker } from './maker';

@Injectable({
    providedIn: 'root'
})
export class MakerCondition {

    constructor(private maker: Maker) { }

    get id(): string {
        return this.maker.id;
    }

    set id(id: string) {
        this.maker.id = id;
    }

    get name(): string {
        return this.maker.name;
    }

    set name(name: string) {
        this.maker.name = name;
    }

    toMap(): Map<string, string> {
        const result: Map<string, string> = new Map();

        if (this.id) {
            result.set('id', this.id);
        }

        if (this.name) {
            result.set('name', this.name);
        }

        return result;
    }
}
