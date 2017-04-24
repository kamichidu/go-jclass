package jclass

type JavaConstructor struct {
	*JavaMethod

	declaringClass *JavaClass
}

func newJavaConstructor(class *JavaClass, method *JavaMethod) *JavaConstructor {
	return &JavaConstructor{method, class}
}

func (self *JavaConstructor) Name() string {
	return self.declaringClass.CanonicalName()
}
