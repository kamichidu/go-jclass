package jclass

import (
	"github.com/kamichidu/go-jclass/jvms"
	"reflect"
	"testing"
)

func TestJavaClassNames(t *testing.T) {
	cases := []struct {
		Filename      string
		PackageName   string
		SimpleName    string
		CanonicalName string
		Name          string
	}{
		{"./testdata/Constants.class", "", "Constants", "Constants", "Constants"},
		{"./testdata/String.class", "java.lang", "String", "java.lang.String", "java.lang.String"},
		{"./testdata/Map$Entry.class", "java.util", "Entry", "java.util.Map.Entry", "java.util.Map$Entry"},
	}
	for _, c := range cases {
		class, err := NewJavaClassFromFilename(c.Filename)
		if err != nil {
			t.Fatalf("NewJavaClassFromFilename: %s", err)
		}

		if class.PackageName() != c.PackageName {
			t.Errorf("PackageName %v, wants %v", class.PackageName(), c.PackageName)
		}
		if class.SimpleName() != c.SimpleName {
			t.Errorf("SimpleName %v, wants %v", class.SimpleName(), c.SimpleName)
		}
		if class.CanonicalName() != c.CanonicalName {
			t.Errorf("CanonicalName %v, wants %v", class.CanonicalName(), c.CanonicalName)
		}
		if class.Name() != c.Name {
			t.Errorf("Name %v, wants %v", class.Name(), c.Name)
		}
	}
}

func TestJavaClassConstructors(t *testing.T) {
	type constructor struct {
		AccessFlags    uint16
		ParameterTypes []string
		// TODO: Throws
	}
	cases := []struct {
		Filename     string
		Constructors []constructor
	}{
		{"./testdata/String.class", []constructor{
			// public java.lang.String();
			{jvms.ACC_PUBLIC, []string{}},
			// public java.lang.String(java.lang.String);
			{jvms.ACC_PUBLIC, []string{"java.lang.String"}},
			// public java.lang.String(char[]);
			{jvms.ACC_PUBLIC, []string{"char[]"}},
			// public java.lang.String(char[], int, int);
			{jvms.ACC_PUBLIC, []string{"char[]", "int", "int"}},
			// public java.lang.String(int[], int, int);
			{jvms.ACC_PUBLIC, []string{"int[]", "int", "int"}},
			// public java.lang.String(byte[], int, int, int);
			{jvms.ACC_PUBLIC, []string{"byte[]", "int", "int", "int"}},
			// public java.lang.String(byte[], int);
			{jvms.ACC_PUBLIC, []string{"byte[]", "int"}},
			// public java.lang.String(byte[], int, int, java.lang.String) throws java.io.UnsupportedEncodingException;
			{jvms.ACC_PUBLIC, []string{"byte[]", "int", "int", "java.lang.String"}},
			// public java.lang.String(byte[], int, int, java.nio.charset.Charset);
			{jvms.ACC_PUBLIC, []string{"byte[]", "int", "int", "java.nio.charset.Charset"}},
			// public java.lang.String(byte[], java.lang.String) throws java.io.UnsupportedEncodingException;
			{jvms.ACC_PUBLIC, []string{"byte[]", "java.lang.String"}},
			// public java.lang.String(byte[], java.nio.charset.Charset);
			{jvms.ACC_PUBLIC, []string{"byte[]", "java.nio.charset.Charset"}},
			// public java.lang.String(byte[], int, int);
			{jvms.ACC_PUBLIC, []string{"byte[]", "int", "int"}},
			// public java.lang.String(byte[]);
			{jvms.ACC_PUBLIC, []string{"byte[]"}},
			// public java.lang.String(java.lang.StringBuffer);
			{jvms.ACC_PUBLIC, []string{"java.lang.StringBuffer"}},
			// public java.lang.String(java.lang.StringBuilder);
			{jvms.ACC_PUBLIC, []string{"java.lang.StringBuilder"}},
		}},
	}
	for _, c := range cases {
		class, err := NewJavaClassFromFilename(c.Filename)
		if err != nil {
			t.Fatalf("NewJavaClassFromFilename: %s", err)
		}

		ctors := class.Constructors()
		if len(ctors) != len(c.Constructors) {
			t.Errorf("Length %v, wants %v", len(ctors), len(c.Constructors))
			continue
		}
		for _, ctor := range ctors {
			found := false
			for _, other := range c.Constructors {
				if ctor.methodInfo.AccessFlags != other.AccessFlags {
					continue
				}
				if !reflect.DeepEqual(ctor.ParameterTypes(), other.ParameterTypes) {
					continue
				}
				found = true
				break
			}
			if !found {
				t.Errorf("Not found constructor: %s %s", ctor.AccessFlags, ctor.ParameterTypes())
			}
		}

	}
}

func TestJavaClassDeclaredConstructors(t *testing.T) {
	type constructor struct {
		AccessFlags    uint16
		ParameterTypes []string
		// TODO: Throws
	}
	cases := []struct {
		Filename     string
		Constructors []constructor
	}{
		{"./testdata/String.class", []constructor{
			// public java.lang.String();
			{jvms.ACC_PUBLIC, []string{}},
			// public java.lang.String(java.lang.String);
			{jvms.ACC_PUBLIC, []string{"java.lang.String"}},
			// public java.lang.String(char[]);
			{jvms.ACC_PUBLIC, []string{"char[]"}},
			// public java.lang.String(char[], int, int);
			{jvms.ACC_PUBLIC, []string{"char[]", "int", "int"}},
			// public java.lang.String(int[], int, int);
			{jvms.ACC_PUBLIC, []string{"int[]", "int", "int"}},
			// public java.lang.String(byte[], int, int, int);
			{jvms.ACC_PUBLIC, []string{"byte[]", "int", "int", "int"}},
			// public java.lang.String(byte[], int);
			{jvms.ACC_PUBLIC, []string{"byte[]", "int"}},
			// public java.lang.String(byte[], int, int, java.lang.String) throws java.io.UnsupportedEncodingException;
			{jvms.ACC_PUBLIC, []string{"byte[]", "int", "int", "java.lang.String"}},
			// public java.lang.String(byte[], int, int, java.nio.charset.Charset);
			{jvms.ACC_PUBLIC, []string{"byte[]", "int", "int", "java.nio.charset.Charset"}},
			// public java.lang.String(byte[], java.lang.String) throws java.io.UnsupportedEncodingException;
			{jvms.ACC_PUBLIC, []string{"byte[]", "java.lang.String"}},
			// public java.lang.String(byte[], java.nio.charset.Charset);
			{jvms.ACC_PUBLIC, []string{"byte[]", "java.nio.charset.Charset"}},
			// public java.lang.String(byte[], int, int);
			{jvms.ACC_PUBLIC, []string{"byte[]", "int", "int"}},
			// public java.lang.String(byte[]);
			{jvms.ACC_PUBLIC, []string{"byte[]"}},
			// public java.lang.String(java.lang.StringBuffer);
			{jvms.ACC_PUBLIC, []string{"java.lang.StringBuffer"}},
			// public java.lang.String(java.lang.StringBuilder);
			{jvms.ACC_PUBLIC, []string{"java.lang.StringBuilder"}},
			// 	      java.lang.String(char[], boolean);
			{0x0, []string{"char[]", "boolean"}},
		}},
	}
	for _, c := range cases {
		class, err := NewJavaClassFromFilename(c.Filename)
		if err != nil {
			t.Fatalf("NewJavaClassFromFilename: %s", err)
		}

		ctors := class.DeclaredConstructors()
		if len(ctors) != len(c.Constructors) {
			t.Errorf("Length %v, wants %v", len(ctors), len(c.Constructors))
			continue
		}
		for _, ctor := range ctors {
			found := false
			for _, other := range c.Constructors {
				if ctor.methodInfo.AccessFlags != other.AccessFlags {
					continue
				}
				if !reflect.DeepEqual(ctor.ParameterTypes(), other.ParameterTypes) {
					continue
				}
				found = true
				break
			}
			if !found {
				t.Errorf("Not found constructor: %s %s", ctor.AccessFlags, ctor.ParameterTypes())
			}
		}

	}
}

func TestJavaClassMethods(t *testing.T) {
	type method struct {
		AccessFlags    uint16
		ReturnType     string
		Name           string
		ParameterTypes []string
		// TODO: Throws
	}
	cases := []struct {
		Filename string
		Methods  []method
	}{
		{"./testdata/Object.class", []method{
			// public final native java.lang.Class<?> getClass();
			{jvms.ACC_PUBLIC | jvms.ACC_FINAL | jvms.ACC_NATIVE, "java.lang.Class", "getClass", []string{}},
			// public native int hashCode();
			{jvms.ACC_PUBLIC | jvms.ACC_NATIVE, "int", "hashCode", []string{}},
			// public boolean equals(java.lang.Object);
			{jvms.ACC_PUBLIC, "boolean", "equals", []string{"java.lang.Object"}},
			// public java.lang.String toString();
			{jvms.ACC_PUBLIC, "java.lang.String", "toString", []string{}},
			// public final native void notify();
			{jvms.ACC_PUBLIC | jvms.ACC_FINAL | jvms.ACC_NATIVE, "void", "notify", []string{}},
			// public final native void notifyAll();
			{jvms.ACC_PUBLIC | jvms.ACC_FINAL | jvms.ACC_NATIVE, "void", "notifyAll", []string{}},
			// public final native void wait(long) throws java.lang.InterruptedException;
			{jvms.ACC_PUBLIC | jvms.ACC_FINAL | jvms.ACC_NATIVE, "void", "wait", []string{"long"}},
			// public final void wait(long, int) throws java.lang.InterruptedException;
			{jvms.ACC_PUBLIC | jvms.ACC_FINAL, "void", "wait", []string{"long", "int"}},
			// public final void wait() throws java.lang.InterruptedException;
			{jvms.ACC_PUBLIC | jvms.ACC_FINAL, "void", "wait", []string{}},
		}},
	}
	for _, c := range cases {
		class, err := NewJavaClassFromFilename(c.Filename)
		if err != nil {
			t.Fatalf("NewJavaClassFromFilename: %s", err)
		}

		methods := class.Methods()
		if len(methods) != len(c.Methods) {
			t.Errorf("Length %v, wants %v", len(methods), len(c.Methods))
			continue
		}
		for _, method := range methods {
			found := false
			for _, other := range c.Methods {
				if method.methodInfo.AccessFlags != other.AccessFlags {
					continue
				}
				if method.ReturnType() != other.ReturnType {
					continue
				}
				if method.Name() != other.Name {
					continue
				}
				if !reflect.DeepEqual(method.ParameterTypes(), other.ParameterTypes) {
					continue
				}
				found = true
				break
			}
			if !found {
				t.Errorf("Not found method: %s %s %s(%s)", method.AccessFlags, method.ReturnType(), method.Name(), method.ParameterTypes())
			}
		}
	}
}
func TestJavaClassDeclaredMethods(t *testing.T) {
	type method struct {
		AccessFlags    uint16
		ReturnType     string
		Name           string
		ParameterTypes []string
		// TODO: Throws
	}
	cases := []struct {
		Filename string
		Methods  []method
	}{
		{"./testdata/Object.class", []method{
			// private static native void registerNatives();
			{jvms.ACC_PRIVATE | jvms.ACC_STATIC | jvms.ACC_NATIVE, "void", "registerNatives", []string{}},
			// public final native java.lang.Class<?> getClass();
			{jvms.ACC_PUBLIC | jvms.ACC_FINAL | jvms.ACC_NATIVE, "java.lang.Class", "getClass", []string{}},
			// public native int hashCode();
			{jvms.ACC_PUBLIC | jvms.ACC_NATIVE, "int", "hashCode", []string{}},
			// public boolean equals(java.lang.Object);
			{jvms.ACC_PUBLIC, "boolean", "equals", []string{"java.lang.Object"}},
			// protected native java.lang.Object clone() throws java.lang.CloneNotSupportedException;
			{jvms.ACC_PROTECTED | jvms.ACC_NATIVE, "java.lang.Object", "clone", []string{}},
			// public java.lang.String toString();
			{jvms.ACC_PUBLIC, "java.lang.String", "toString", []string{}},
			// public final native void notify();
			{jvms.ACC_PUBLIC | jvms.ACC_FINAL | jvms.ACC_NATIVE, "void", "notify", []string{}},
			// public final native void notifyAll();
			{jvms.ACC_PUBLIC | jvms.ACC_FINAL | jvms.ACC_NATIVE, "void", "notifyAll", []string{}},
			// public final native void wait(long) throws java.lang.InterruptedException;
			{jvms.ACC_PUBLIC | jvms.ACC_FINAL | jvms.ACC_NATIVE, "void", "wait", []string{"long"}},
			// public final void wait(long, int) throws java.lang.InterruptedException;
			{jvms.ACC_PUBLIC | jvms.ACC_FINAL, "void", "wait", []string{"long", "int"}},
			// public final void wait() throws java.lang.InterruptedException;
			{jvms.ACC_PUBLIC | jvms.ACC_FINAL, "void", "wait", []string{}},
			// protected void finalize() throws java.lang.Throwable;
			{jvms.ACC_PROTECTED, "void", "finalize", []string{}},
		}},
	}
	for _, c := range cases {
		class, err := NewJavaClassFromFilename(c.Filename)
		if err != nil {
			t.Fatalf("NewJavaClassFromFilename: %s", err)
		}

		methods := class.DeclaredMethods()
		if len(methods) != len(c.Methods) {
			t.Errorf("Length %v, wants %v", len(methods), len(c.Methods))
			continue
		}
		for _, method := range methods {
			found := false
			for _, other := range c.Methods {
				if method.methodInfo.AccessFlags != other.AccessFlags {
					continue
				}
				if method.ReturnType() != other.ReturnType {
					continue
				}
				if method.Name() != other.Name {
					continue
				}
				if !reflect.DeepEqual(method.ParameterTypes(), other.ParameterTypes) {
					continue
				}
				found = true
				break
			}
			if !found {
				t.Errorf("Not found method: %s %s %s(%s)", method.AccessFlags, method.ReturnType(), method.Name(), method.ParameterTypes())
			}
		}
	}
}
