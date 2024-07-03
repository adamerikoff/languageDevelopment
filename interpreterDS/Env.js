class Env {
    constructor(record = {}, parent = null) {
        this.record = record;
        this.parent = parent;
    }
    eAssign(name, value) {
        this.record[name] = value;
        return value;
    }
    eReassign(name, value) {
        this.resolve(name).record[name] = value;
        return value;
    }
    getVariableValue(name) {
        return this.resolve(name).record[name];
    }
    resolve(name) {
        if (this.record.hasOwnProperty(name)) {
            return this;
        }
        if (this.parent === null) {
            throw new ReferenceError(`Unable to resolve variable value for ${name}`);
        }
        return this.parent.resolve(name);
    }
}

export default  Env;