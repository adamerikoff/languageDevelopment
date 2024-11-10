class Eva {
    eval(exp) {
        if (this.isNumber(exp)) {
            return exp;
        }
        if (this.isString(exp)) {
            return exp.slice(1, -1);
        }
        throw "Unimplemented";
    }

    isNumber(exp) {
        return typeof exp === "number";
    }

    isString(exp) {
        return typeof exp === "string" && exp[0] === `"` && exp.slice(-1) === `"`;
    }
}

module.exports = { Eva };