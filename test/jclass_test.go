package test

import (
    "github.com/kamichidu/go-jclass"
    "reflect"
    "testing"
)

func Test(t *testing.T) {
    jc, err := jclass.NewJClassWithFilename("String.class")
    if err != nil {
        t.Error(err)
    }

    if jc.GetClassName() != "java/lang/String" {
        t.Errorf("Name %s", jc.GetClassName())
    }
    // if !reflect.DeepEqual(jc.GetInterfaces(), []string{"java/io/Serializable", "java/lang/Comparable<java.lang.String>", "java.lang.CharSequence"}) {
    if !reflect.DeepEqual(jc.GetInterfaces(), []string{"java/io/Serializable", "java/lang/Comparable", "java/lang/CharSequence"}) {
        t.Errorf("Interfaces: %v", jc.GetInterfaces())
    }
}

// javap -cp ./classes/ java.lang.String
// ------------------------------
// Compiled from "String.java"
// public final class java.lang.String implements java.io.Serializable, java.lang.Comparable<java.lang.String>, java.lang.CharSequence {
//   public static final java.util.Comparator<java.lang.String> CASE_INSENSITIVE_ORDER;
//   public java.lang.String();
//   public java.lang.String(java.lang.String);
//   public java.lang.String(char[]);
//   public java.lang.String(char[], int, int);
//   public java.lang.String(int[], int, int);
//   public java.lang.String(byte[], int, int, int);
//   public java.lang.String(byte[], int);
//   public java.lang.String(byte[], int, int, java.lang.String) throws java.io.UnsupportedEncodingException;
//   public java.lang.String(byte[], int, int, java.nio.charset.Charset);
//   public java.lang.String(byte[], java.lang.String) throws java.io.UnsupportedEncodingException;
//   public java.lang.String(byte[], java.nio.charset.Charset);
//   public java.lang.String(byte[], int, int);
//   public java.lang.String(byte[]);
//   public java.lang.String(java.lang.StringBuffer);
//   public java.lang.String(java.lang.StringBuilder);
//   java.lang.String(char[], boolean);
//   public int length();
//   public boolean isEmpty();
//   public char charAt(int);
//   public int codePointAt(int);
//   public int codePointBefore(int);
//   public int codePointCount(int, int);
//   public int offsetByCodePoints(int, int);
//   void getChars(char[], int);
//   public void getChars(int, int, char[], int);
//   public void getBytes(int, int, byte[], int);
//   public byte[] getBytes(java.lang.String) throws java.io.UnsupportedEncodingException;
//   public byte[] getBytes(java.nio.charset.Charset);
//   public byte[] getBytes();
//   public boolean equals(java.lang.Object);
//   public boolean contentEquals(java.lang.StringBuffer);
//   public boolean contentEquals(java.lang.CharSequence);
//   public boolean equalsIgnoreCase(java.lang.String);
//   public int compareTo(java.lang.String);
//   public int compareToIgnoreCase(java.lang.String);
//   public boolean regionMatches(int, java.lang.String, int, int);
//   public boolean regionMatches(boolean, int, java.lang.String, int, int);
//   public boolean startsWith(java.lang.String, int);
//   public boolean startsWith(java.lang.String);
//   public boolean endsWith(java.lang.String);
//   public int hashCode();
//   public int indexOf(int);
//   public int indexOf(int, int);
//   public int lastIndexOf(int);
//   public int lastIndexOf(int, int);
//   public int indexOf(java.lang.String);
//   public int indexOf(java.lang.String, int);
//   static int indexOf(char[], int, int, java.lang.String, int);
//   static int indexOf(char[], int, int, char[], int, int, int);
//   public int lastIndexOf(java.lang.String);
//   public int lastIndexOf(java.lang.String, int);
//   static int lastIndexOf(char[], int, int, java.lang.String, int);
//   static int lastIndexOf(char[], int, int, char[], int, int, int);
//   public java.lang.String substring(int);
//   public java.lang.String substring(int, int);
//   public java.lang.CharSequence subSequence(int, int);
//   public java.lang.String concat(java.lang.String);
//   public java.lang.String replace(char, char);
//   public boolean matches(java.lang.String);
//   public boolean contains(java.lang.CharSequence);
//   public java.lang.String replaceFirst(java.lang.String, java.lang.String);
//   public java.lang.String replaceAll(java.lang.String, java.lang.String);
//   public java.lang.String replace(java.lang.CharSequence, java.lang.CharSequence);
//   public java.lang.String[] split(java.lang.String, int);
//   public java.lang.String[] split(java.lang.String);
//   public static java.lang.String join(java.lang.CharSequence, java.lang.CharSequence...);
//   public static java.lang.String join(java.lang.CharSequence, java.lang.Iterable<? extends java.lang.CharSequence>);
//   public java.lang.String toLowerCase(java.util.Locale);
//   public java.lang.String toLowerCase();
//   public java.lang.String toUpperCase(java.util.Locale);
//   public java.lang.String toUpperCase();
//   public java.lang.String trim();
//   public java.lang.String toString();
//   public char[] toCharArray();
//   public static java.lang.String format(java.lang.String, java.lang.Object...);
//   public static java.lang.String format(java.util.Locale, java.lang.String, java.lang.Object...);
//   public static java.lang.String valueOf(java.lang.Object);
//   public static java.lang.String valueOf(char[]);
//   public static java.lang.String valueOf(char[], int, int);
//   public static java.lang.String copyValueOf(char[], int, int);
//   public static java.lang.String copyValueOf(char[]);
//   public static java.lang.String valueOf(boolean);
//   public static java.lang.String valueOf(char);
//   public static java.lang.String valueOf(int);
//   public static java.lang.String valueOf(long);
//   public static java.lang.String valueOf(float);
//   public static java.lang.String valueOf(double);
//   public native java.lang.String intern();
//   public int compareTo(java.lang.Object);
//   static {};
// }
