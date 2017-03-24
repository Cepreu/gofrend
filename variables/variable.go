package variables

import (
	"fmt"
)

const (
	ATTR_SYSTEM          uint8 = 1
	ATTR_CRM             uint8 = 2
	ATTR_EXTERNAL              = 4
	ATTR_INTERNAL              = 8
	ATTR_USER_PREDEFINED       = 16
	ATTR_TTS_ENUMERATION       = 32
	ATTR_USER_DEFINED          = ATTR_EXTERNAL | ATTR_INTERNAL | ATTR_USER_PREDEFINED
	ATTR_WRITABLE              = ATTR_CRM | ATTR_USER_DEFINED
	ATTR_ANY                   = ATTR_SYSTEM | ATTR_CRM | ATTR_USER_DEFINED
)

type Variable struct {
	value       Value
	name        string
	description string
	attributes  uint8
	isNullValue bool
}

func (v Variable) compareTo(arg Variable) (uint8, error) {
	if v.isNullValue && arg.isNullValue {
		return 0, nil
	} else if v.isNullValue || arg.isNullValue {
		return 0, fmt.Errorf("Cannot compare variable to NULL value : %v and %v", v.value, arg.value)
	}
	return v.value.compareTo(arg.value), nil
}

func (pv *Variable) assign(val Value) {
	if val.isEmpty() {
		pv.isNullValue = true
	} else {
		pv.value.assign(val)
		pv.isNullValue = false
	}
}
func (pv *Variable) assignNull() {
	pv.isNullValue = true
}
func (v Variable) isExternal() bool {
	return !((v.attributes & ATTR_EXTERNAL) == 0)
}

func (v Variable) isCrm() bool {
	return (v.attributes & ATTR_CRM) != 0
}

func (v Variable) isCav() bool {
	return false //TBD (this instanceof CavVariable);
}

func (v Variable) toString() string {
	re := "NULL"
	if !v.isNullValue {
		re = v.value
	}
	return fmt.Sprintf("{{name=\"%s\"}{description=\"%s\"} %s}", v.name, v.description, re)
}

/*
func (v Variable) hashCode() uint64 {
	prime := 31
	result := 1
	result = prime * result + attributes;
	result = prime * result + ((description == null) ? 0 : description.hashCode());
	result = prime * result + (isNull() ? 1231 : 1237);
	result = prime * result + ((name == null) ? 0 : name.hashCode());
	result = prime * result + ((getValue() == null) ? 0 : getValue().hashCode());
	return result;
}

    func equals( Object obj ) bool {
        if ( this == obj )
            return true;
        if ( obj == null )
            return false;
        if ( !(obj instanceof Variable) )
            return false;
        Variable other = (Variable) obj;
        if ( attributes != other.attributes )
            return false;
        if ( description == null ) {
            if ( other.description != null )
                return false;
        } else if ( !description.equals(other.description) )
            return false;
        if ( isNull() != other.isNull() )
            return false;
        if ( name == null ) {
            if ( other.name != null )
                return false;
        } else if ( !name.equals(other.name) )
            return false;

        //FIXME: Comparing of Value - "equals" methods should be created for each Value descendant
        if ( this.getType() != other.getType() )
            return false;
        Value thisValue = this.getValue();
        Value thatValue = other.getValue();
        if ( thisValue == null )
            return thatValue == null;
        else if ( thatValue == null )
            return false;
        else
            try {
                return thisValue.compareTo(thatValue) == 0;
            } catch ( VariableException e ) {
                return false;
            }
    }
}*/
