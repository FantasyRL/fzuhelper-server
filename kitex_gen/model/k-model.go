/*
Copyright 2024 The west2-online Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// Code generated by Kitex v0.11.3. DO NOT EDIT.

package model

import (
	"bytes"
	"fmt"
	"reflect"
	"strings"

	"github.com/cloudwego/gopkg/protocol/thrift"
)

// unused protection
var (
	_ = fmt.Formatter(nil)
	_ = (*bytes.Buffer)(nil)
	_ = (*strings.Builder)(nil)
	_ = reflect.Type(nil)
	_ = thrift.STOP
)

func (p *BaseResp) FastRead(buf []byte) (int, error) {
	var err error
	var offset int
	var l int
	var fieldTypeId thrift.TType
	var fieldId int16
	for {
		fieldTypeId, fieldId, l, err = thrift.Binary.ReadFieldBegin(buf[offset:])
		offset += l
		if err != nil {
			goto ReadFieldBeginError
		}
		if fieldTypeId == thrift.STOP {
			break
		}
		switch fieldId {
		case 1:
			if fieldTypeId == thrift.I64 {
				l, err = p.FastReadField1(buf[offset:])
				offset += l
				if err != nil {
					goto ReadFieldError
				}
			} else {
				l, err = thrift.Binary.Skip(buf[offset:], fieldTypeId)
				offset += l
				if err != nil {
					goto SkipFieldError
				}
			}
		case 2:
			if fieldTypeId == thrift.STRING {
				l, err = p.FastReadField2(buf[offset:])
				offset += l
				if err != nil {
					goto ReadFieldError
				}
			} else {
				l, err = thrift.Binary.Skip(buf[offset:], fieldTypeId)
				offset += l
				if err != nil {
					goto SkipFieldError
				}
			}
		default:
			l, err = thrift.Binary.Skip(buf[offset:], fieldTypeId)
			offset += l
			if err != nil {
				goto SkipFieldError
			}
		}
	}

	return offset, nil
ReadFieldBeginError:
	return offset, thrift.PrependError(fmt.Sprintf("%T read field %d begin error: ", p, fieldId), err)
ReadFieldError:
	return offset, thrift.PrependError(fmt.Sprintf("%T read field %d '%s' error: ", p, fieldId, fieldIDToName_BaseResp[fieldId]), err)
SkipFieldError:
	return offset, thrift.PrependError(fmt.Sprintf("%T field %d skip type %d error: ", p, fieldId, fieldTypeId), err)
}

func (p *BaseResp) FastReadField1(buf []byte) (int, error) {
	offset := 0

	var _field int64
	if v, l, err := thrift.Binary.ReadI64(buf[offset:]); err != nil {
		return offset, err
	} else {
		offset += l
		_field = v
	}
	p.Code = _field
	return offset, nil
}

func (p *BaseResp) FastReadField2(buf []byte) (int, error) {
	offset := 0

	var _field string
	if v, l, err := thrift.Binary.ReadString(buf[offset:]); err != nil {
		return offset, err
	} else {
		offset += l
		_field = v
	}
	p.Msg = _field
	return offset, nil
}

// for compatibility
func (p *BaseResp) FastWrite(buf []byte) int {
	return 0
}

func (p *BaseResp) FastWriteNocopy(buf []byte, w thrift.NocopyWriter) int {
	offset := 0
	if p != nil {
		offset += p.fastWriteField1(buf[offset:], w)
		offset += p.fastWriteField2(buf[offset:], w)
	}
	offset += thrift.Binary.WriteFieldStop(buf[offset:])
	return offset
}

func (p *BaseResp) BLength() int {
	l := 0
	if p != nil {
		l += p.field1Length()
		l += p.field2Length()
	}
	l += thrift.Binary.FieldStopLength()
	return l
}

func (p *BaseResp) fastWriteField1(buf []byte, w thrift.NocopyWriter) int {
	offset := 0
	offset += thrift.Binary.WriteFieldBegin(buf[offset:], thrift.I64, 1)
	offset += thrift.Binary.WriteI64(buf[offset:], p.Code)
	return offset
}

func (p *BaseResp) fastWriteField2(buf []byte, w thrift.NocopyWriter) int {
	offset := 0
	offset += thrift.Binary.WriteFieldBegin(buf[offset:], thrift.STRING, 2)
	offset += thrift.Binary.WriteStringNocopy(buf[offset:], w, p.Msg)
	return offset
}

func (p *BaseResp) field1Length() int {
	l := 0
	l += thrift.Binary.FieldBeginLength()
	l += thrift.Binary.I64Length()
	return l
}

func (p *BaseResp) field2Length() int {
	l := 0
	l += thrift.Binary.FieldBeginLength()
	l += thrift.Binary.StringLengthNocopy(p.Msg)
	return l
}

func (p *LoginData) FastRead(buf []byte) (int, error) {
	var err error
	var offset int
	var l int
	var fieldTypeId thrift.TType
	var fieldId int16
	var issetId bool = false
	var issetCookies bool = false
	for {
		fieldTypeId, fieldId, l, err = thrift.Binary.ReadFieldBegin(buf[offset:])
		offset += l
		if err != nil {
			goto ReadFieldBeginError
		}
		if fieldTypeId == thrift.STOP {
			break
		}
		switch fieldId {
		case 1:
			if fieldTypeId == thrift.STRING {
				l, err = p.FastReadField1(buf[offset:])
				offset += l
				if err != nil {
					goto ReadFieldError
				}
				issetId = true
			} else {
				l, err = thrift.Binary.Skip(buf[offset:], fieldTypeId)
				offset += l
				if err != nil {
					goto SkipFieldError
				}
			}
		case 2:
			if fieldTypeId == thrift.LIST {
				l, err = p.FastReadField2(buf[offset:])
				offset += l
				if err != nil {
					goto ReadFieldError
				}
				issetCookies = true
			} else {
				l, err = thrift.Binary.Skip(buf[offset:], fieldTypeId)
				offset += l
				if err != nil {
					goto SkipFieldError
				}
			}
		default:
			l, err = thrift.Binary.Skip(buf[offset:], fieldTypeId)
			offset += l
			if err != nil {
				goto SkipFieldError
			}
		}
	}

	if !issetId {
		fieldId = 1
		goto RequiredFieldNotSetError
	}

	if !issetCookies {
		fieldId = 2
		goto RequiredFieldNotSetError
	}
	return offset, nil
ReadFieldBeginError:
	return offset, thrift.PrependError(fmt.Sprintf("%T read field %d begin error: ", p, fieldId), err)
ReadFieldError:
	return offset, thrift.PrependError(fmt.Sprintf("%T read field %d '%s' error: ", p, fieldId, fieldIDToName_LoginData[fieldId]), err)
SkipFieldError:
	return offset, thrift.PrependError(fmt.Sprintf("%T field %d skip type %d error: ", p, fieldId, fieldTypeId), err)
RequiredFieldNotSetError:
	return offset, thrift.NewProtocolException(thrift.INVALID_DATA, fmt.Sprintf("required field %s is not set", fieldIDToName_LoginData[fieldId]))
}

func (p *LoginData) FastReadField1(buf []byte) (int, error) {
	offset := 0

	var _field string
	if v, l, err := thrift.Binary.ReadString(buf[offset:]); err != nil {
		return offset, err
	} else {
		offset += l
		_field = v
	}
	p.Id = _field
	return offset, nil
}

func (p *LoginData) FastReadField2(buf []byte) (int, error) {
	offset := 0

	_, size, l, err := thrift.Binary.ReadListBegin(buf[offset:])
	offset += l
	if err != nil {
		return offset, err
	}
	_field := make([]string, 0, size)
	for i := 0; i < size; i++ {
		var _elem string
		if v, l, err := thrift.Binary.ReadString(buf[offset:]); err != nil {
			return offset, err
		} else {
			offset += l
			_elem = v
		}

		_field = append(_field, _elem)
	}
	p.Cookies = _field
	return offset, nil
}

// for compatibility
func (p *LoginData) FastWrite(buf []byte) int {
	return 0
}

func (p *LoginData) FastWriteNocopy(buf []byte, w thrift.NocopyWriter) int {
	offset := 0
	if p != nil {
		offset += p.fastWriteField1(buf[offset:], w)
		offset += p.fastWriteField2(buf[offset:], w)
	}
	offset += thrift.Binary.WriteFieldStop(buf[offset:])
	return offset
}

func (p *LoginData) BLength() int {
	l := 0
	if p != nil {
		l += p.field1Length()
		l += p.field2Length()
	}
	l += thrift.Binary.FieldStopLength()
	return l
}

func (p *LoginData) fastWriteField1(buf []byte, w thrift.NocopyWriter) int {
	offset := 0
	offset += thrift.Binary.WriteFieldBegin(buf[offset:], thrift.STRING, 1)
	offset += thrift.Binary.WriteStringNocopy(buf[offset:], w, p.Id)
	return offset
}

func (p *LoginData) fastWriteField2(buf []byte, w thrift.NocopyWriter) int {
	offset := 0
	offset += thrift.Binary.WriteFieldBegin(buf[offset:], thrift.LIST, 2)
	listBeginOffset := offset
	offset += thrift.Binary.ListBeginLength()
	var length int
	for _, v := range p.Cookies {
		length++
		offset += thrift.Binary.WriteStringNocopy(buf[offset:], w, v)
	}
	thrift.Binary.WriteListBegin(buf[listBeginOffset:], thrift.STRING, length)
	return offset
}

func (p *LoginData) field1Length() int {
	l := 0
	l += thrift.Binary.FieldBeginLength()
	l += thrift.Binary.StringLengthNocopy(p.Id)
	return l
}

func (p *LoginData) field2Length() int {
	l := 0
	l += thrift.Binary.FieldBeginLength()
	l += thrift.Binary.ListBeginLength()
	for _, v := range p.Cookies {
		_ = v
		l += thrift.Binary.StringLengthNocopy(v)
	}
	return l
}

func (p *Classroom) FastRead(buf []byte) (int, error) {
	var err error
	var offset int
	var l int
	var fieldTypeId thrift.TType
	var fieldId int16
	var issetBuild bool = false
	var issetLocation bool = false
	var issetCapacity bool = false
	var issetType bool = false
	for {
		fieldTypeId, fieldId, l, err = thrift.Binary.ReadFieldBegin(buf[offset:])
		offset += l
		if err != nil {
			goto ReadFieldBeginError
		}
		if fieldTypeId == thrift.STOP {
			break
		}
		switch fieldId {
		case 1:
			if fieldTypeId == thrift.STRING {
				l, err = p.FastReadField1(buf[offset:])
				offset += l
				if err != nil {
					goto ReadFieldError
				}
				issetBuild = true
			} else {
				l, err = thrift.Binary.Skip(buf[offset:], fieldTypeId)
				offset += l
				if err != nil {
					goto SkipFieldError
				}
			}
		case 2:
			if fieldTypeId == thrift.STRING {
				l, err = p.FastReadField2(buf[offset:])
				offset += l
				if err != nil {
					goto ReadFieldError
				}
				issetLocation = true
			} else {
				l, err = thrift.Binary.Skip(buf[offset:], fieldTypeId)
				offset += l
				if err != nil {
					goto SkipFieldError
				}
			}
		case 3:
			if fieldTypeId == thrift.STRING {
				l, err = p.FastReadField3(buf[offset:])
				offset += l
				if err != nil {
					goto ReadFieldError
				}
				issetCapacity = true
			} else {
				l, err = thrift.Binary.Skip(buf[offset:], fieldTypeId)
				offset += l
				if err != nil {
					goto SkipFieldError
				}
			}
		case 4:
			if fieldTypeId == thrift.STRING {
				l, err = p.FastReadField4(buf[offset:])
				offset += l
				if err != nil {
					goto ReadFieldError
				}
				issetType = true
			} else {
				l, err = thrift.Binary.Skip(buf[offset:], fieldTypeId)
				offset += l
				if err != nil {
					goto SkipFieldError
				}
			}
		default:
			l, err = thrift.Binary.Skip(buf[offset:], fieldTypeId)
			offset += l
			if err != nil {
				goto SkipFieldError
			}
		}
	}

	if !issetBuild {
		fieldId = 1
		goto RequiredFieldNotSetError
	}

	if !issetLocation {
		fieldId = 2
		goto RequiredFieldNotSetError
	}

	if !issetCapacity {
		fieldId = 3
		goto RequiredFieldNotSetError
	}

	if !issetType {
		fieldId = 4
		goto RequiredFieldNotSetError
	}
	return offset, nil
ReadFieldBeginError:
	return offset, thrift.PrependError(fmt.Sprintf("%T read field %d begin error: ", p, fieldId), err)
ReadFieldError:
	return offset, thrift.PrependError(fmt.Sprintf("%T read field %d '%s' error: ", p, fieldId, fieldIDToName_Classroom[fieldId]), err)
SkipFieldError:
	return offset, thrift.PrependError(fmt.Sprintf("%T field %d skip type %d error: ", p, fieldId, fieldTypeId), err)
RequiredFieldNotSetError:
	return offset, thrift.NewProtocolException(thrift.INVALID_DATA, fmt.Sprintf("required field %s is not set", fieldIDToName_Classroom[fieldId]))
}

func (p *Classroom) FastReadField1(buf []byte) (int, error) {
	offset := 0

	var _field string
	if v, l, err := thrift.Binary.ReadString(buf[offset:]); err != nil {
		return offset, err
	} else {
		offset += l
		_field = v
	}
	p.Build = _field
	return offset, nil
}

func (p *Classroom) FastReadField2(buf []byte) (int, error) {
	offset := 0

	var _field string
	if v, l, err := thrift.Binary.ReadString(buf[offset:]); err != nil {
		return offset, err
	} else {
		offset += l
		_field = v
	}
	p.Location = _field
	return offset, nil
}

func (p *Classroom) FastReadField3(buf []byte) (int, error) {
	offset := 0

	var _field string
	if v, l, err := thrift.Binary.ReadString(buf[offset:]); err != nil {
		return offset, err
	} else {
		offset += l
		_field = v
	}
	p.Capacity = _field
	return offset, nil
}

func (p *Classroom) FastReadField4(buf []byte) (int, error) {
	offset := 0

	var _field string
	if v, l, err := thrift.Binary.ReadString(buf[offset:]); err != nil {
		return offset, err
	} else {
		offset += l
		_field = v
	}
	p.Type = _field
	return offset, nil
}

// for compatibility
func (p *Classroom) FastWrite(buf []byte) int {
	return 0
}

func (p *Classroom) FastWriteNocopy(buf []byte, w thrift.NocopyWriter) int {
	offset := 0
	if p != nil {
		offset += p.fastWriteField1(buf[offset:], w)
		offset += p.fastWriteField2(buf[offset:], w)
		offset += p.fastWriteField3(buf[offset:], w)
		offset += p.fastWriteField4(buf[offset:], w)
	}
	offset += thrift.Binary.WriteFieldStop(buf[offset:])
	return offset
}

func (p *Classroom) BLength() int {
	l := 0
	if p != nil {
		l += p.field1Length()
		l += p.field2Length()
		l += p.field3Length()
		l += p.field4Length()
	}
	l += thrift.Binary.FieldStopLength()
	return l
}

func (p *Classroom) fastWriteField1(buf []byte, w thrift.NocopyWriter) int {
	offset := 0
	offset += thrift.Binary.WriteFieldBegin(buf[offset:], thrift.STRING, 1)
	offset += thrift.Binary.WriteStringNocopy(buf[offset:], w, p.Build)
	return offset
}

func (p *Classroom) fastWriteField2(buf []byte, w thrift.NocopyWriter) int {
	offset := 0
	offset += thrift.Binary.WriteFieldBegin(buf[offset:], thrift.STRING, 2)
	offset += thrift.Binary.WriteStringNocopy(buf[offset:], w, p.Location)
	return offset
}

func (p *Classroom) fastWriteField3(buf []byte, w thrift.NocopyWriter) int {
	offset := 0
	offset += thrift.Binary.WriteFieldBegin(buf[offset:], thrift.STRING, 3)
	offset += thrift.Binary.WriteStringNocopy(buf[offset:], w, p.Capacity)
	return offset
}

func (p *Classroom) fastWriteField4(buf []byte, w thrift.NocopyWriter) int {
	offset := 0
	offset += thrift.Binary.WriteFieldBegin(buf[offset:], thrift.STRING, 4)
	offset += thrift.Binary.WriteStringNocopy(buf[offset:], w, p.Type)
	return offset
}

func (p *Classroom) field1Length() int {
	l := 0
	l += thrift.Binary.FieldBeginLength()
	l += thrift.Binary.StringLengthNocopy(p.Build)
	return l
}

func (p *Classroom) field2Length() int {
	l := 0
	l += thrift.Binary.FieldBeginLength()
	l += thrift.Binary.StringLengthNocopy(p.Location)
	return l
}

func (p *Classroom) field3Length() int {
	l := 0
	l += thrift.Binary.FieldBeginLength()
	l += thrift.Binary.StringLengthNocopy(p.Capacity)
	return l
}

func (p *Classroom) field4Length() int {
	l := 0
	l += thrift.Binary.FieldBeginLength()
	l += thrift.Binary.StringLengthNocopy(p.Type)
	return l
}
