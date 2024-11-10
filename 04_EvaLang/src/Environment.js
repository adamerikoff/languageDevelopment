class Environment {
    constructor(record = {}) {
        this.record = record;
    }

    define(name, value) {
        this.record[name] = value;
        return value;
    }
}

module.exports = { Environment };