// Copyright 2018 gf Author(https://gitee.com/johng/gf). All Rights Reserved.
//
// This Source Code Form is subject to the terms of the MIT License.
// If a copy of the MIT was not distributed with this file,
// You can obtain one at https://gitee.com/johng/gf.

// 包方法操作

package gtimer_test

import (
    "gitee.com/johng/gf/g/container/garray"
    "gitee.com/johng/gf/g/os/gtimer"
    "gitee.com/johng/gf/g/util/gtest"
    "testing"
    "time"
)


func New() *gtimer.Timer {
    return gtimer.New(10, 10*time.Millisecond)
}

func TestWheel_Add_Close(t *testing.T) {
    gtest.Case(t, func() {
        wheel  := New()
        array  := garray.New(0, 0)
        //fmt.Println("start", time.Now())
        wheel.Add(time.Second, func() {
            //fmt.Println("entry1", time.Now())
            array.Append(1)
        })
        wheel.Add(time.Second, func() {
            //fmt.Println("entry2", time.Now())
            array.Append(1)
        })
        wheel.Add(2*time.Second, func() {
            //fmt.Println("entry3", time.Now())
            array.Append(1)
        })
        time.Sleep(1300*time.Millisecond)
        gtest.Assert(array.Len(), 2)
        time.Sleep(1300*time.Millisecond)
        gtest.Assert(array.Len(), 5)
        wheel.Close()
        time.Sleep(1200*time.Millisecond)
        fixedLength := array.Len()
        time.Sleep(1200*time.Millisecond)
        gtest.Assert(array.Len(), fixedLength)
    })
}

func TestWheel_Singleton(t *testing.T) {
   gtest.Case(t, func() {
       wheel := New()
       array := garray.New(0, 0)
       wheel.AddSingleton(time.Second, func() {
           array.Append(1)
           time.Sleep(10*time.Second)
       })
       time.Sleep(1200*time.Millisecond)
       gtest.Assert(array.Len(), 1)

       time.Sleep(1200*time.Millisecond)
       gtest.Assert(array.Len(), 1)
   })
}

func TestWheel_Once(t *testing.T) {
   gtest.Case(t, func() {
       wheel  := New()
       array  := garray.New(0, 0)
       wheel.AddOnce(time.Second, func() {
           array.Append(1)
       })
       wheel.AddOnce(time.Second, func() {
           array.Append(1)
       })
       time.Sleep(1200*time.Millisecond)
       gtest.Assert(array.Len(), 2)
       time.Sleep(1200*time.Millisecond)
       gtest.Assert(array.Len(), 2)
       wheel.Close()
       time.Sleep(1200*time.Millisecond)
       fixedLength := array.Len()
       time.Sleep(1200*time.Millisecond)
       gtest.Assert(array.Len(), fixedLength)
   })
}

func TestWheel_DelayAdd(t *testing.T) {
   gtest.Case(t, func() {
       wheel := New()
       array := garray.New(0, 0)
       wheel.DelayAdd(time.Second, time.Second, func() {
           array.Append(1)
       })
       time.Sleep(1200*time.Millisecond)
       gtest.Assert(array.Len(), 0)
       time.Sleep(1200*time.Millisecond)
       gtest.Assert(array.Len(), 1)
   })
}

func TestWheel_DelayAdd_Singleton(t *testing.T) {
   gtest.Case(t, func() {
       wheel := New()
       array := garray.New(0, 0)
       wheel.DelayAddSingleton(time.Second, time.Second, func() {
           array.Append(1)
           time.Sleep(10*time.Second)
       })
       time.Sleep(1200*time.Millisecond)
       gtest.Assert(array.Len(), 0)

       time.Sleep(1200*time.Millisecond)
       gtest.Assert(array.Len(), 1)
   })
}

func TestWheel_DelayAdd_Once(t *testing.T) {
   gtest.Case(t, func() {
       wheel := New()
       array := garray.New(0, 0)
       wheel.DelayAddOnce(time.Second, time.Second, func() {
           array.Append(1)
       })
       time.Sleep(1200*time.Millisecond)
       gtest.Assert(array.Len(), 0)

       time.Sleep(1200*time.Millisecond)
       gtest.Assert(array.Len(), 1)

       time.Sleep(1200*time.Millisecond)
       gtest.Assert(array.Len(), 1)
   })
}

func TestWheel_ExitJob(t *testing.T) {
   gtest.Case(t, func() {
       wheel := New()
       array := garray.New(0, 0)
       wheel.Add(time.Second, func() {
           array.Append(1)
           gtimer.Exit()
       })
       time.Sleep(1200*time.Millisecond)
       gtest.Assert(array.Len(), 1)
   })
}