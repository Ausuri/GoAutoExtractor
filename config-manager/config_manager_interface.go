package configmanager

type configManagerInterface interface {
	getSetting(settingName string) (any, error)
}

type configLocationInterface interface {
	getPaths() *configFilePaths
}

type Primitive interface {
	~int |
		~int8 |
		~int16 |
		~int32 |
		~int64 |
		~uint |
		~uint8 |
		~uint16 |
		~uint32 |
		~uint64 |
		~uintptr |
		~float32 |
		~float64 |
		~complex64 |
		~complex128 |
		~string |
		~bool
}
