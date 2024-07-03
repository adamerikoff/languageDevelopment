class Env {
    constructor(record = {}, parent = null) {
        this.recordMap = record;
        this.parentEnv = parent;
    }

    // Affecte une variable avec un nom et une valeur donnés à l'environnement courant.
    assignVariable(varName, varValue) {
        this.recordMap[varName] = varValue;
        return varValue
    }
    // Réaffecte une variable avec un nom et une valeur donnés dans la chaîne d'environnement résolu.
    reassignVariable(varName, varValue) {
        this.determineEnv(varName).recordMap[varName] = varValue;
        return varValue;
    }
    // Récupère la valeur d'une variable avec le nom donné dans la chaîne d'environnement résolu.
    retrieveVariable(varName) {
        return this.determineEnv(varName).recordMap[varName];
    }
    // Résout l'environnement où la variable avec le nom donné est définie.
    determineEnv(varName) {
        if (varName in this.recordMap) {
            return this;
        }
        if (this.parentEnv === null) {
            throw new ReferenceError(`Variable "${varName}" is not defined.`);
        }
        return this.parentEnv.determineEnv(varName);
    }
}

export default Env;