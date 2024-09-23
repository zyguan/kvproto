package util

type DeprecatedField struct{}

func (d DeprecatedField) Marshal() ([]byte, error)                  { return nil, nil }
func (d *DeprecatedField) MarshalTo(data []byte) (n int, err error) { return 0, nil }
func (d *DeprecatedField) Size() int                                { return 0 }
func (d *DeprecatedField) Unmarshal(data []byte) error              { return nil }
