export class Exercise {
    private _name: string = "";
    private _rep: number = 8;
    private _weight: number = 0;
    private _sets: number = 3;

    constructor() {

    }

    get sets(): number {
        return this._sets;
    }

    set sets(sets: number) {
        this.sets = sets;
    }

    get weight(): number {
        return this._weight;
    }

    set weight(weight: number) {
        this.weight = weight;
    }

    get name(): string {
        return this._name;
    }

    set name(name: string) {
        this._name = name;
    }

    get rep(): number {
        return this._rep;
    }

    set rep(rep: number) {
        this._rep = rep;
    }
}
