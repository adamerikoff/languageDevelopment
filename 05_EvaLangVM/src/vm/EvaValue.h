#ifndef EVAVALUE_H
#define EVAVALUE_H

#include "includes.h"

enum class EvaValueType {
    NUMBER,
    OBJECT,
};

enum class ObjectType {
    STRING,
};

struct Object {
    ObjectType type;
    Object(ObjectType type);
};

struct StringObject : public Object {
    std::string str;
    StringObject(const std::string& str);
};

struct EvaValue {
    EvaValueType type;
    union {
        double number;
        Object* object;
    };

    EvaValue();
    EvaValue(double value);
    EvaValue(const std::string& str);
    double asNumber() const;
    StringObject* asString() const;
    std::string asCPPString() const;
};

#endif // EVAVALUES_H
