#include "EvaValue.h"

// Constructor implementation
Object::Object(ObjectType type) {
    this->type = type;
}

StringObject::StringObject(const std::string& str) : Object(ObjectType::STRING) {
    this->str = str;
}

EvaValue::EvaValue() {
    
}

EvaValue::EvaValue(double value) {
    this->type = EvaValueType::NUMBER;
    this->number = value;
}

EvaValue::EvaValue(const std::string& str) {
    this->type = EvaValueType::OBJECT;
    this->object = new StringObject(str);
}

double EvaValue::asNumber() const {
    if (this->type == EvaValueType::NUMBER) {
        return this->number;
    }
    std::cerr << "Error: Attempted to access non-number value as a number." << std::endl;
    return 0.0;
}

StringObject* EvaValue::asString() const {
    if (this->type == EvaValueType::OBJECT) {
        if (this->object->type == ObjectType::STRING) {
            return static_cast<StringObject*>(this->object);
        }
    }
    std::cerr << "Error: Attempted to access non-object value as an object." << std::endl;
    return nullptr;
}

std::string EvaValue::asCPPString() const {
    if (this->type == EvaValueType::OBJECT && this->object) {
        StringObject* strObj = static_cast<StringObject*>(this->object);
        if (strObj != nullptr) {
            return strObj->str;
        }
    }

    if (this->type == EvaValueType::NUMBER) {
        return std::to_string(this->number);
    }

    std::cerr << "Error: Attempted to access non-object or non-number value as a string." << std::endl;
    return "";
}

