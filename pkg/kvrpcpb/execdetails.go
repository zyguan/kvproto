package kvrpcpb

type DeprecatedField struct{}

func (d DeprecatedField) Marshal() ([]byte, error)                  { return nil, nil }
func (d *DeprecatedField) MarshalTo(data []byte) (n int, err error) { return 0, nil }
func (d *DeprecatedField) Size() int                                { return 0 }
func (d *DeprecatedField) Unmarshal(data []byte) error              { return nil }

type LazyExecDetails struct {
	data    []byte
	details *ExecDetailsV2
}

func (d *LazyExecDetails) Marshal() ([]byte, error) {
	return d.details.Marshal()
}

func (d *LazyExecDetails) MarshalTo(data []byte) (n int, err error) {
	return d.details.MarshalTo(data)
}

func (d *LazyExecDetails) Size() int {
	return d.details.Size()
}

func (d *LazyExecDetails) Unmarshal(data []byte) error {
	d.data = data
	return nil
}

func (d *LazyExecDetails) Details() *ExecDetailsV2 {
	if d.details == nil {
		d.details = &ExecDetailsV2{}
		d.details.Unmarshal(d.data)
	}
	return d.details
}

func (d *LazyExecDetails) SetDetails(details *ExecDetailsV2) {
	d.details = details
}
