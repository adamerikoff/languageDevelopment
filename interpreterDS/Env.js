class Env {
    constructor(record = {}) {
        this.record = record;
    }
    define(name, value) {
        this.record[name] = value;
        return value
    }
    getVariableValue(name) {
        if (!this.record.hasOwnProperty(name)) {
            throw new ReferenceError(`Unable to get variable value for ${name}`);
        }
        return this.record[name];
    }
}

export default  Env;