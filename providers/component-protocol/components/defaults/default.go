// Copyright (c) 2021 Terminus, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package defaults

import (
	"fmt"
	"reflect"
	"regexp"

	"github.com/erda-project/erda-infra/providers/component-protocol/cptype"
)

// FieldActualImplRef .
const (
	FieldActualImplRef = "ActualImplRef"
)

// DefaultImpl .
type DefaultImpl struct {
	// ActualImplRef inject by framework according to field FieldActualImplRef
	ActualImplRef cptype.IComponent
}

func (d *DefaultImpl) RegisterCompStdOps() (opFuncs map[cptype.OperationKey]cptype.OperationFunc) {
	opFuncs = make(map[cptype.OperationKey]cptype.OperationFunc)
	ve := reflect.ValueOf(d.ActualImplRef)
	te := reflect.TypeOf(d.ActualImplRef)
	for i := 0; i < ve.NumMethod(); i++ {
		vm := ve.Method(i)
		fName := te.Method(i).Name
		if !regexp.MustCompile(`^Register.*Op$`).MatchString(fName) {
			continue
		}
		var args []reflect.Value
		for i := 0; i < vm.Type().NumIn(); i++ {
			args = append(args, reflect.ValueOf(vm.Type().In(i)))
		}
		fmt.Printf("%#v\n", args[0])
		fmt.Println("canInterface: ", args[0].CanInterface())
		fmt.Println("implements: ", args[0].Type().Implements(reflect.TypeOf((*cptype.IOperation)(nil)).Elem()))
		fmt.Println("pkgPath: ", args[0].Type().PkgPath())
		op, _ := args[0].Interface().(cptype.IOperation)
		opFuncs[op.OpKey()] = func(sdk *cptype.SDK) cptype.IStdStructuredPtr {
			vm.Call(args)
			return nil
		}
	}
	return
}

// RegisterRenderingOp .
func (d *DefaultImpl) RegisterRenderingOp() (opFunc cptype.OperationFunc) {
	return d.ActualImplRef.RegisterInitializeOp()
}
