#ifndef EVAVALUE_H
#define EVAVALUE_H

#include "../includes.h"

enum class EvaValueType {
    NUMBER,
    BOOLEAN,
    OBJECT
};

enum class ObjectType {
    STRING,
    CODE,
};

struct Object {
    ObjectType type;
    Object(ObjectType type): type(type) {}
};

struct StringObject: public Object {
    std::string string;
    StringObject(const std::string& str) : Object(ObjectType::STRING), string(str) {}
};


struct EvaValue {
    EvaValueType type;
    union {
        double number;
        bool boolean;
        Object* object;
    };
};

struct CodeObject : public Object {
    std::string name;
    std::vector<EvaValue> constants;
    std::vector<uint8_t> code;
    CodeObject(const std::string& name) : Object(ObjectType::CODE), name(name) {}
};


#define NUMBER(value) ((EvaValue){EvaValueType::NUMBER, .number = value})
#define BOOLEAN(value) ((EvaValue){EvaValueType::BOOLEAN, .boolean = value})
#define ALLOC_STRING(value) ((EvaValue){EvaValueType::OBJECT, .object = new StringObject(value)})
#define ALLOC_CODE(name) ((EvaValue){EvaValueType::OBJECT, .object = new CodeObject(name)})

#define AS_NUMBER(evaValue) ((double)(evaValue).number)
#define AS_BOOLEAN(evaValue) ((bool)(evaValue).boolean)
#define AS_OBJECT(evaValue) ((Object*)(evaValue).object)
#define AS_STRING(evaValue) ((StringObject*)(evaValue).object)
#define AS_CPPSTRING(evaValue) (AS_STRING(evaValue)->string)
#define AS_CODE(evaValue) ((CodeObject*)(evaValue).object)

#define IS_NUMBER(evaValue) ((evaValue).type == EvaValueType::NUMBER)
#define IS_BOOLEAN(evaValue) ((evaValue).type == EvaValueType::BOOLEAN)
#define IS_OBJECT(evaValue) ((evaValue).type == EvaValueType::OBJECT)
#define IS_OBJECT_TYPE(evaValue, objectType) (IS_OBJECT(evaValue) && AS_OBJECT(evaValue)->type == objectType)
#define IS_STRING(evaValue) IS_OBJECT_TYPE(evaValue, ObjectType::STRING)
#define IS_CODE(evaValue) IS_OBJECT_TYPE(evaValue, ObjectType::CODE)

#endif // EVAVALUES_H
