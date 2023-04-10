/*
 * Copyright 2021 ByteDance Inc.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package encoder

import (
    `encoding`
    `encoding/json`
    `reflect`
)

var (
    byteType                 = reflect.TypeOf(byte(0))
    jsonNumberType           = reflect.TypeOf(json.Number(""))
    jsonUnsupportedValueType = reflect.TypeOf(new(json.UnsupportedValueError))
)

var (
    errorType                 = reflect.TypeOf((*error)(nil)).Elem()
    jsonMarshalerType         = reflect.TypeOf((*json.Marshaler)(nil)).Elem()
    encodingTextMarshalerType = reflect.TypeOf((*encoding.TextMarshaler)(nil)).Elem()
)

func isSimpleByte(vt reflect.Type) bool {
    if vt.Kind() != byteType.Kind() {
        return false
    } else {
        return !isEitherMarshaler(vt) && !isEitherMarshaler(reflect.PtrTo(vt))
    }
}

func isEitherMarshaler(vt reflect.Type) bool {
    return vt.Implements(jsonMarshalerType) || vt.Implements(encodingTextMarshalerType)
}
